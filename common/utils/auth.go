package utils

import (
	"github.com/dgrijalva/jwt-go"
	"qpgame/config"
	"qpgame/models/beans"
	"qpgame/models/xorm"
	"qpgame/ramcache"
	"reflect"
	"strconv"
	"time"
)

// 生成/检验token
func GenerateToken(user *xorm.Users) (string, error) {
	//fmt.Println(user.Phone)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		//"phone":    user.Phone,
		//"userName": user.UserName,
		"sub": user.Id,
		//"exp":   config.JwtTokenExp, // 可以添加过期时间
		"exp": time.Now().Unix() + config.JwtTokenExp, // 可以添加过期时间
	})
	return token.SignedString([]byte(config.TokenKey)) //对应的字符串请自行生成，最后足够使用加密后的字符串
}

// 更新用户基本信息缓存
func UpdateUserIdCard(platform string, uid int, user map[string]interface{}) {
	userIdCard, _ := ramcache.UserIdCard.Load(platform)
	userIdCardMap := userIdCard.(map[int]beans.UserProfile)
	userProfile := userIdCardMap[uid]
	for key, val := range user {
		valType := reflect.TypeOf(val).String()
		if valType == "string" {
			switch key {
			case "Username":
				userProfile.Username = val.(string)
			case "Phone":
				userProfile.Phone = val.(string)
			case "Token":
				userProfile.Token = val.(string)
			case "TokenCreated":
				iTokenCreated, _ := strconv.Atoi(val.(string))
				userProfile.TokenCreated = iTokenCreated
			case "UserGroupId":
				userProfile.UserGroupId = val.(string)
			case "UniqueCode":
				userProfile.UniqueCode = val.(string)
			case "WxOpenId":
				userProfile.WxOpenId = val.(string)
			}
		}
	}
	if userProfile.WxOpenId != "" {
		// 如果 wxOpenId不为空时（微信新用户），需要更新相关缓存
		wxOpenId := userProfile.WxOpenId
		woiIdx, _ := ramcache.WxOpenIdIndex.Load(platform)
		woiIdxMap := woiIdx.(map[string]beans.WxOpenId)
		woiIdxMap[wxOpenId] = beans.WxOpenId{
			UserId: uid,
		}
	}
	userIdCardMap[uid] = userProfile
}
