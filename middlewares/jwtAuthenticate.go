package middlewares

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"qpgame/admin/common"
	"qpgame/common/utils"
	"qpgame/config"
	"qpgame/models/beans"
	"qpgame/models/xorm"
	"qpgame/ramcache"

	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
)

// 存储所有登录用户最后一次操作时间，token -> created
var UserLastLoginTimeMap = map[string]float64{}
// 存储所有银商登录用户最后一次操作时间，token -> created
var SilverMerchantLastLoginTimeMap = map[string]float64{}

func JwtAuthenticate() *jwtmiddleware.Middleware {
	return jwtmiddleware.New(jwtmiddleware.Config{//jwt中间件
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) { //这个方法将验证jwt的token
			return []byte(config.TokenKey), nil //自己加密的秘钥
		},
		//设置后，中间件会验证令牌是否使用特定的签名算法进行签名
		//如果签名方法不是常量，则可以使用ValidationKeyGetter回调来实现其他检查
		//加密的方式
		SigningMethod: jwt.SigningMethodHS256,
		ErrorHandler: func(ctx iris.Context, err string) { //验证未通过错误处理方式
			if err == "Required authorization token not found" {
				utils.ResFaiJSON(&ctx, "缺少token", "请登录", config.TOKENEMPTY)
				return
			}
			if err == "Token is expired" { //token过期
				utils.ResFaiJSON(&ctx, "token已过期", "登录已过期", config.TOKENEXPIRED)
				return
			}
			if strings.Contains(err, "signing method but token specified") {
				utils.ResFaiJSON(&ctx, "token签名方式不正确", "请重新登录", config.TOKENEMPTY)
				return
			}
			utils.ResFaiJSON(&ctx, err, "需要登录", config.TOKENEMPTY)
			return
		},
		ContextKey: "jwt",
		Debug: false,
		Expiration: true,
	})
}

//token过滤,这里需要查询数据库，不用redis考虑是因为
//去redis查询后还是要查询数据库,那就直接查数据
func JwtHandler() iris.Handler {
	return func(ctx iris.Context) {
		token := ctx.Values().Get("jwt").(*jwt.Token)
		platform := ctx.Params().Get("platform")
		claim, _ := token.Claims.(jwt.MapClaims)
		iUserId := int(claim["sub"].(float64))
		sUserId := strconv.Itoa(iUserId)
		userIdCard, _ := ramcache.UserIdCard.Load(platform)
		userIdCardMap := userIdCard.(map[int]beans.UserProfile)
		if _, uidExist := userIdCardMap[iUserId]; uidExist == false {
			utils.ResFaiJSON(&ctx, "1906011117", "请重新登录", config.TOKENEMPTY)
			return
		}
		userProfile := userIdCardMap[iUserId]
		phone := userProfile.Phone
		realPhone := phone
		userName := userProfile.Username
		//var phoNuT interface{}
		//if phone == "" {
		//	fmt.Println("@user")
		//	phone = userName
		//	phoNuT, _ = ramcache.UserNameAndToken.Load(platform)
		//} else {
		//	fmt.Println("@phone")
		//	phoNuT, _ = ramcache.PhoneNumAndToken.Load(platform)
		//}
		//tokenCache := phoNuT.(map[string][]string)[phone]
		//if len(tokenCache) < 2 {
		//	utils.ResFaiJSON(&ctx, "1906010000", "请重新登录", config.TOKENEMPTY)
		//	return
		//}
		//cacheId := tokenCache[0]
		//cacheToken := tokenCache[1]
		if userProfile.Token != token.Raw {
			utils.ResFaiJSON(&ctx, "1906030000", "请重新登录", config.TOKENEMPTY)
			return
			//if cacheId == sUserId {
			//	user := models2.Users{Id: iUserId}
			//	_, err := models.MyEngine[platform].Get(&user)
			//	if err != nil {
			//		fmt.Println(err.Error())
			//	}
			//	if user.Token == token.Raw && user.TokenCreated+3600*72 > utils.GetNowTime() {
			//		//用户id
			//		ctx.Values().Set("userid", sUserId)
			//		ctx.Values().Set("phone", realPhone)
			//		ctx.Values().Set("username", userName)
			//		ctx.Next()
			//		return
			//	}
			//	utils.ResFaiJSON(&ctx, "", "请重新登录", config.TOKENEMPTY)
			//	return
			//} else {
			//	//if cachetoken != token.Raw && cacheid == userid {
			//	utils.ResFaiJSON(&ctx, "该token用户不存在", "请重新登录", config.TOKENEMPTY)
			//	return
			//	//}
			//}
		}
		if (userProfile.TokenCreated + 3600*72) < utils.GetNowTime() {
			utils.ResFaiJSON(&ctx, "", "登录已过期", config.TOKENEXPIRED)
			return
		}
		//用户id
		ctx.Values().Set("userid", sUserId)
		ctx.Values().Set("phone", realPhone)
		ctx.Values().Set("username", userName)
		ctx.Next()
	}
}

// 后台token验证
func AdminJwtHandler() (handler iris.Handler) {
	callback := func(ctx iris.Context) {
		authString := ctx.GetHeader("Authorization") //提交過來的授權字符串
		tokenString := strings.Replace(authString, "bearer ", "", 7)
		requestToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.TokenKey), nil
		})
		if err != nil || requestToken == nil {
			utils.ResFaiJSON(&ctx, "Token不存在, 或者解板token失败", "需要令牌", config.TOKENEMPTY)
			return
		}
		claim, _ := requestToken.Claims.(jwt.MapClaims) //解码-提交的token
		idStr := claim["id"].(string)                   //用户id
		adminId, idErr := strconv.Atoi(idStr)           //转换为int
		if idErr != nil { //转换有误
			utils.ResFaiJSON(&ctx, "解析Token-ID出错", "令牌格式有误", config.TOKENEMPTY)
			return
		}

		savedTokenStr := common.GetAdminToken(&ctx, adminId) //拿到已经保存过的token
		if strings.Compare(requestToken.Raw, savedTokenStr) != 0 { //如果两者一致, 则不往下执行
			utils.ResFaiJSON(&ctx, "Token与保存的数据不一致", "用户已经在其他地方登录", config.TOKENEXPIRED)
			return
		}
		// 登录后最后一次操作时间
		lastOperateTime := claim["created"].(float64) //创建时间
		if _, ok := UserLastLoginTimeMap[tokenString]; ok {
			lastOperateTime = UserLastLoginTimeMap[tokenString]
		}
		currentTime := float64(utils.GetNowTime())   //当前时间
		expired := config.AdminTokenExpire.Seconds() //过期时间
		if currentTime-lastOperateTime > expired { //时间不能过期
			if _, ok := UserLastLoginTimeMap[tokenString]; ok {
				delete(UserLastLoginTimeMap, tokenString)
			}
			utils.ResFaiJSON(&ctx, "解析Token-Created出错", "令牌已经过期", config.TOKENEXPIRED)
			return
		} else {
			UserLastLoginTimeMap[tokenString] = currentTime
		}
		ctx.Next()
	}
	return callback
}

// 银商登录token验证
func SilverMerchantHandler() (handler iris.Handler) {
	return func(ctx iris.Context) {
		token := ctx.Values().Get("jwt").(*jwt.Token)
		platform := ctx.Params().Get("platform")
		claim, _ := token.Claims.(jwt.MapClaims)
		iUserId := int(claim["sub"].(float64))
		// 银商用户不需要写缓存 ，已和LuGer确认
		smUser, exist := xorm.GetSliverMerchantUser(platform, iUserId)
		if exist == false {
			utils.ResFaiJSON(&ctx, "1906101453", "用户权限验证失败", config.TOKENEMPTY)
			return
		}
		tokenString := token.Raw
		if smUser.Token != tokenString {
			utils.ResFaiJSON(&ctx, "1906101500", "请重新登录", config.TOKENEMPTY)
			return
		}
		// 登录后最后一次操作时间
		lastOperateTime := float64(smUser.TokenCreated) //创建时间
		if _, ok := UserLastLoginTimeMap[tokenString]; ok {
			lastOperateTime = UserLastLoginTimeMap[tokenString]
		}
		currentTime := float64(utils.GetNowTime())
		if (currentTime - lastOperateTime) > 3600*72 {
			if _, ok := UserLastLoginTimeMap[tokenString]; ok {
				delete(UserLastLoginTimeMap, tokenString)
			}
			utils.ResFaiJSON(&ctx, "1906101501", "登录已过期", config.TOKENEXPIRED)
			return
		} else {
			SilverMerchantLastLoginTimeMap[tokenString] = currentTime
		}
		//用户id
		ctx.Values().Set("silverMerchantUserId", iUserId)
		ctx.Values().Set("account", smUser.Account)
		ctx.Values().Set("merchantLevel", smUser.MerchantLevel)
		ctx.Next()
	}
}

func GetIdFromClaims(key string, claims jwt.Claims) string {
	v := reflect.ValueOf(claims)
	if v.Kind() == reflect.Map {
		for _, k := range v.MapKeys() {
			value := v.MapIndex(k)
			if fmt.Sprintf("%s", k.Interface()) == key {
				return fmt.Sprintf("%v", value.Interface())
			}
		}
	}
	return ""
}
