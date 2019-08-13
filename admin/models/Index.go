package models

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"qpgame/admin/common"
	"qpgame/admin/validations"
	"qpgame/common/utils"
	"qpgame/config"
	db "qpgame/models"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
)

// 登录接口
func Login(ctx *iris.Context) (interface{}, error) {
	// 提交的信息验证
	loginVali := &validations.UserLoginValidation{}
	row := map[string]string{}
	messages, result := loginVali.Validate(ctx, &row)
	if !result {
		return nil, Error{What: messages} //Result{Message: messages, Data: nil}, false
	}

	key := (*ctx).Params().GetString("platform") + "-" + row["role_id"]
	menus, ok := common.AdminRoleMenus[key]
	if !ok {
		return nil, Error{What: "登录失败: 当前用户的操作权限不足"}
	}

	id, _ := strconv.Atoi(row["id"])                   //后台用户id
	roleId, _ := strconv.Atoi(row["role_id"])          //后台用户角色id
	maxManual, mErr := strconv.Atoi(row["max_manual"]) //最大操作金额
	if mErr != nil {
		maxManual = 0
	}

	// 开发阶段暂时把ip相关的验证去掉, 切勿删除以下代码
	//allowIps := "," + row["login_ip"] + ","
	//if strings.Compare(allowIps, "") == 0 {
	//	return nil, Error{What: "登录失败: 不允许在此IP登录"}
	//}
	//currentIp := "," + utils.GetIp((*ctx).Request()) + ","
	//if !strings.Contains(allowIps, currentIp) {
	//	return nil, Error{What: "登录失败: 不允许在此IP(" + currentIp + ")登录"}
	//}

	// 检查是否已登录, 如果已登录，检查登录是不是过期
	platform := (*ctx).Params().Get("platform") //平台标识
	//loginToken := getLoginedToken(platform, &row)               //已登录的token
	//if loginToken != "" && !loginHasExpired(loginToken, &row) { //判断是否有已登录的token并且是否过期
	//	return nil, Error{What: "登录失败: 用户已在其他地方登录"}
	//}

	//生成token
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   row["id"],   //用户密码
		"name": row["name"], //用户名称
		//"exp":     config.JwtTokenExp, //后台过期时间要短一些
		"created": utils.GetNowTime(), //此token的创建时间
	}).SignedString([]byte(config.TokenKey))
	if err != nil {
		return nil, Error{What: "登录失败: 生成用户TOKEN时发生错误"}
	}

	//写用户登录日志
	data := map[string]string{
		"admin_id":   row["id"],
		"admin_name": row["name"],
		"ip":         utils.GetIp((*ctx).Request()),
		"login_time": strconv.Itoa(utils.GetNowTime()),
	}
	insertId, err := createRecord(ctx, &data, "admin_login_logs")
	if err != nil || insertId <= 0 {
		return nil, Error{What: "登录失败: 保存登录日志时发生错误"}
	}

	common.StoreLoginedToken(platform, id, token)
	return common.LoginInfo{
		Info: common.UserInfo{
			Id:              id,
			Name:            row["name"],
			RealName:        row["real_name"],
			MaxManualAmount: maxManual,
			RoleId:          roleId,
		},
		Token: token,
		Menus: menus,
	}, nil
}

// 登录重置
func LoginReset(ctx *iris.Context) error {
	postData := utils.GetPostData(ctx)
	username := postData.Get("username")
	if strings.Compare(username, "") == 0 {
		return Error{What: "用户名不能为空"}
	}
	rows, err := db.MyEngine[(*ctx).Params().Get("platform")].Sql("select token from admins where name='" + username + "'").QueryString()
	if err != nil || len(rows) == 0 {
		return Error{What: "查询统计信息失败"}
	}
	tokenStr := rows[0]["token"]
	if tokenStr != "" {
		platform := (*ctx).Params().Get("platform")
		err := common.RemoveLoginedToken(platform, tokenStr)
		if err != nil {
			return err
		}
	}
	return nil
}

// 退出接口
func Logout(ctx *iris.Context) error {
	tokenStr := (*ctx).GetHeader("Authorization")
	if strings.Compare(tokenStr, "") == 0 {
		return Error{What: "用户还未登录"}
	}
	tokenStr = strings.Replace(tokenStr, "bearer ", "", 7)
	platform := (*ctx).Params().Get("platform")
	err := common.RemoveLoginedToken(platform, tokenStr)
	if err != nil {
		return err
	}
	return nil
}

// 得到网站配置/需要从数据库当中读取
func GetConfig(ctx *Context) (interface{}, error) {
	conf := common.ApiConfig{ //默认的相关配置

	}
	origin := (*ctx).GetHeader("Origin")
	sql := "SELECT code, name, admin_address, admin_api_address FROM platform WHERE admin_address = '" + strings.ReplaceAll(origin, "http://", "https://") + "'"
	rows, err := db.MyEngineMainDb.SQL(sql).QueryString()
	if err != nil {
		return conf, err
	}
	if len(rows) == 0 {
		return conf, errors.New("没有相关的配置")
	}

	row := rows[0]
	conf.Code = row["code"]
	conf.Name = row["name"]
	conf.Url = row["admin_address"]
	conf.ApiUrl = row["admin_api_address"]
	return conf, nil
}

// 唯一字符串
func GUID() string {
	b := make([]byte, 48)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", md5.Sum([]byte(base64.URLEncoding.EncodeToString(b))))
}

// 获取验证码
func Verify(ctx *Context) (string, string) {
	idKey, baseCodes := utils.CreateCaptcha()
	platform := (*ctx).Params().Get("platform") //平台标识号
	guid := GUID()
	key := platform + "-" + guid
	keyTime := utils.GetNowTime()

	// 删除所有过期验证码
	for key, val := range common.AdminVerifyCodes {
		createdCode, _ := strconv.Atoi(val[1])
		if keyTime-createdCode > 30 { //如果大于30秒
			delete(common.AdminVerifyCodes, key)
		}
	}
	common.AdminVerifyCodes[key] = []string{idKey, strconv.Itoa(keyTime)} //将当前保存验证码保存
	return baseCodes, guid
}

// 是否已经处于登录状态
func getLoginedToken(platform string, row *map[string]string) string {
	key := platform + "-" + (*row)["id"]
	token, has := common.AdminTokens[key]
	if !has {
		return (*row)["token"]
	}
	return token
}

// 登录是否过期
func loginHasExpired(tokenString string, row *map[string]string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.TokenKey), nil
	})
	if err != nil { //解析的token不正确
		return true
	}

	claim, _ := token.Claims.(jwt.MapClaims) //解码
	if claim["id"] != (*row)["id"] {         //原登录id与现登录id不一致
		return true
	}

	_, idErr := strconv.Atoi(claim["id"].(string)) //拿到用戶編號
	if idErr != nil {                              //没有找到此管理员信息
		return true
	}

	created, _ := claim["created"].(float64)     //已登录的token创建时间
	current := float64(utils.GetNowTime())       //当前时间
	seconds := config.AdminTokenExpire.Seconds() //token失效的时间
	if current-created > seconds {               //如果当前时间 - token创建时间 > token保存时长, 即: 旧的token已经失效
		return true
	}

	//如果旧的token还没有失效, 则返回false, 表示旧的登录还没有过期
	return false
}
