package frontEndControllers

import (
	"encoding/json"
	"errors"
	goXorm "github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	"qpgame/common/services"
	"qpgame/common/utils"
	"qpgame/config"
	"qpgame/models"
	"qpgame/models/beans"
	"qpgame/models/xorm"
	"qpgame/ramcache"
	"strconv"
	"time"
)

type WxController struct {
	platform string
	ctx      iris.Context
}

type WxService struct {
	platform  string
	loginFrom string
	ip        string
	engine    *goXorm.Engine
}

var WxAppId = ""
var WxAppSecret = ""

const WxAuthUrl string = "https://api.weixin.qq.com/sns/oauth2/access_token"
const WxUserInfoUrl string = "https://api.weixin.qq.com/sns/userinfo"

//构造函数
func NewWxController(ctx iris.Context) *WxController {
	obj := new(WxController)
	obj.platform = ctx.Params().Get("platform")
	obj.ctx = ctx
	for k, v := range config.PlatformCPs {
		if k == obj.platform {
			wxSet := v.(map[string]interface{})["wx"].(config.WxSet)
			WxAppId = wxSet.AppId
			WxAppSecret = wxSet.AppSecret
		}
	}
	return obj
}

// 获取access_token
func wxGetAccessToken(sCode string) (map[string]string, bool) {
	params := map[string]string{
		"appid":      WxAppId,
		"secret":     WxAppSecret,
		"code":       sCode,
		"grant_type": "authorization_code",
	}
	reqUrl := utils.BuildUrl(WxAuthUrl, params)
	respData := utils.ReqGet(reqUrl, 10*time.Second)
	result := make(map[string]string)
	var ok bool
	json.Unmarshal(respData, &result)
	if _, exist := result["access_token"]; exist {
		ok = true
	} else {
		ok = false
	}
	return result, ok
}

func wxGetUserInfo(accessToken, openId string) (userInfo map[string]string, ok bool) {
	params := map[string]string{
		"access_token": accessToken,
		"openid":       openId,
	}
	reqUrl := utils.BuildUrl(WxUserInfoUrl, params)
	respData := utils.ReqGet(reqUrl, 10*time.Second)
	userInfo = make(map[string]string)
	json.Unmarshal(respData, &userInfo)
	if _, exist := userInfo["openid"]; exist {
		ok = true
	} else {
		ok = false
	}
	return
}

func (service WxService) createNewUser(userBean xorm.Users, iNow int) (xorm.Users, error) {
	session := service.engine.NewSession()
	createErr := session.Begin()
	defer session.Close()
	_, createErr = session.InsertOne(&userBean)
	if createErr != nil {
		session.Rollback()
		return userBean, createErr
	}
	iUserId := userBean.Id
	_, createErr = session.Insert(xorm.Accounts{UserId: iUserId, Updated: iNow})
	if createErr != nil {
		session.Rollback()
		return userBean, createErr
	}
	token, _ := utils.GenerateToken(&userBean)
	userBean.Token = token
	userBean.TokenCreated = iNow
	userBean.LastLoginTime = iNow
	//登录成功以后将token缓存到本地
	sTokenTime := strconv.Itoa(userBean.TokenCreated)
	_, createErr = session.ID(iUserId).Update(userBean)
	if createErr != nil {
		session.Rollback()
		return userBean, createErr
	}
	loginLog := xorm.UserLoginLogs{UserId: iUserId, Ip: service.ip, LoginTime: iNow, LoginFrom: service.loginFrom}
	_, createErr = session.InsertOne(loginLog)
	if createErr != nil {
		session.Rollback()
		//utils.ResFaiJSON(&ctx, createErr.Error(), "绑定微信失败", config.NOTGETDATA)
		return userBean, createErr
	}

	err := services.PromotionAward(service.platform, session, userBean.ParentId)
	if err != nil {
		return userBean, err
	}
	var innerMsg string
	innerMsg, err = services.ActivityAward(service.platform, session, 1, userBean.Id, service.ip)
	if err != nil {
		session.Rollback()
		return userBean, errors.New(innerMsg + "； " + err.Error())
	}

	createErr = session.Commit()
	ut, _ := ramcache.UserNameAndToken.Load(service.platform)
	utMap := ut.(map[string][]string)
	sUserId := strconv.Itoa(iUserId)
	utMap[userBean.UserName] = []string{sUserId, token, sTokenTime, "1"}
	utils.UpdateUserIdCard(service.platform, iUserId, map[string]interface{}{
		"Username":     userBean.UserName,
		"Token":        userBean.Token,
		"TokenCreated": sTokenTime, // 注意，这里要用字符串，否则会提示登录过期
		"WxOpenId":     userBean.WxOpenId,
	})
	return userBean, nil
}

// 检查用户是否已经绑定过微信
func (service WxService) checkIsBind(openId string) (userBean xorm.Users, bind bool) {
	woiIdx, _ := ramcache.WxOpenIdIndex.Load(service.platform)
	woiIdxMap := woiIdx.(map[string]beans.WxOpenId)
	var wxOpenId beans.WxOpenId
	wxOpenId, bind = woiIdxMap[openId]
	if bind {
		userId := wxOpenId.UserId
		uic, _ := ramcache.UserIdCard.Load(service.platform)
		uicMap := uic.(map[int]beans.UserProfile)
		userProfile := uicMap[userId]
		userBean = xorm.Users{
			Id:       userId,
			Phone:    userProfile.Phone,
			UserName: userProfile.Username,
		}
	}
	return
}

func (service WxService) doLogin(userBean xorm.Users) (xorm.Users, error) {
	token, _ := utils.GenerateToken(&userBean)
	now := utils.GetNowTime()
	var userUpdateBean = xorm.Users{
		Token:         token,
		TokenCreated:  now,
		LastLoginTime: now,
	}
	var respUserBean xorm.Users
	//开始事务
	session := service.engine.NewSession()
	err := session.Begin()
	defer session.Close()
	_, err = session.ID(userBean.Id).Update(userUpdateBean)
	if err != nil {
		session.Rollback()
		return respUserBean, err
	}
	loginLog := xorm.UserLoginLogs{UserId: userBean.Id, Ip: service.ip, LoginTime: now, LoginFrom: service.loginFrom}
	_, err = session.InsertOne(loginLog)
	if err != nil {
		session.Rollback()
		return respUserBean, err
	}
	err = session.Commit()
	if err != nil {
		return respUserBean, err
	}
	userBean.Password = ""
	userBean.Token = userUpdateBean.Token
	userBean.TokenCreated = userUpdateBean.TokenCreated
	userBean.LastLoginTime = userUpdateBean.LastLoginTime
	utils.UpdateUserIdCard(service.platform, userBean.Id, map[string]interface{}{
		"Token":        userBean.Token,
		"TokenCreated": strconv.Itoa(now), // 注意，这里要用字符串，否则会提示登录过期
	})
	return userBean, nil
}

/**
 * @api {post} api/v1/wxLogin 微信登录
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: aTian</span><br/><br/>
 * 微信登录<br>
 * 业务描述: 微信账号登录，如果之前没有绑定微信，先创建账号再登录</br>
 * @apiVersion 1.0.0
 * @apiName     WxLogin
 * @apiGroup    user
 * @apiPermission iso,android客户端
 * @apiParam (客户端请求参数) {string} code       从微信获取到的code，用于微信OAuth2.0授权登录
 * @apiParam (客户端请求参数) {string} parent_id  上级代理用户Id
 * @apiParam (客户端请求参数) {string} login_from 登录来源 IOS Android
 *
 * @apiError (请求失败返回) {int}      code            错误代码
 * @apiError (请求失败返回) {string}   clientMsg       提示信息
 * @apiError (请求失败返回) {string}   internalMsg     错误代码
 * @apiError (请求失败返回) {float}    timeConsumed    后台耗时
 *
 * @apiErrorExample {json} 失败返回
 * {
 *      "code": 204,
 *      "internalMsg": "",
 *      "clientMsg ": 0,
 *      "timeConsumed": 0
 * }
 *
 * @apiSuccess (返回结果)  {int}      code            200
 * @apiSuccess (返回结果)  {string}   clientMsg       提示信息
 * @apiSuccess (返回结果)  {string}   internalMsg     提示信息
 * @apiSuccess (返回结果)  {json}     data            返回数据
 * @apiSuccess (返回结果)  {float}    timeConsumed    后台耗时
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *     "clientMsg": "登录成功",
 *     "code": 200,
 *     "data": {
 *         "Id": 207,
 *         "Phone": "",
 *         "Password": "",
 *         "UserName": "15592054945OlocX",
 *         "Name": "mj1958",
 *         "Email": "",
 *         "Created": 1559205494,
 *         "Birthday": "",
 *         "MobileType": 1,
 *         "Sex": 1,
 *         "Path": "",
 *         "VipLevel": 1,
 *         "Qq": "",
 *         "Wechat": "",
 *         "Status": 1,
 *         "ProxyStatus": 0,
 *         "UserType": 0,
 *         "Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NTk0NjQ2OTUsInBob25lIjoiIiwic3ViIjoyMDcsInVzZXJOYW1lIjoiMTU1OTIwNTQ5NDVPbG9jWCJ9.h9R-y7HtGNgrU3KdE08vdFUH3zJhZ1IzK3q0JPXTSL4",
 *         "RegIp": "",
 *         "UniqueCode": "",
 *         "TokenCreated": 1559205494,
 *         "SafePassword": "",
 *         "UserGroupId": "",
 *         "ParentId": 0,
 *         "LastLoginTime": 1559205494,
 *         "LastPlatformId": 0,
 *         "GroupSize": 0
 *     },
 *     "internalMsg": "",
 *     "timeConsumed": 8123478
 * }
 */
func (controller *WxController) WxLogin() {
	ctx := controller.ctx
	if !utils.RequiredParamPost(&ctx, []string{"code", "login_from"}) {
		return
	}
	sCode := ctx.FormValue("code")
	sLoginFrom := ctx.FormValue("login_from")
	sParentId := ctx.FormValue("parent_id")
	iParentId, transErr := strconv.Atoi(sParentId)
	if transErr != nil {
		iParentId = 0
	}
	tokenInfo, tokenOk := wxGetAccessToken(sCode)
	if tokenOk {
		platform := controller.platform
		engine := models.MyEngine[platform]
		var service = &WxService{
			engine:   engine,
			platform: platform,
		}
		sOpenId := tokenInfo["openid"]
		userBean, bind := service.checkIsBind(sOpenId)
		if bind {
			service.ip = utils.GetIp(ctx.Request())
			service.loginFrom = sLoginFrom
			userBean, logErr := service.doLogin(userBean)
			if logErr == nil {
				user := new(xorm.Users)
				engine.ID(userBean.Id).Get(user)
				utils.ResSuccJSON(&ctx, "", "登录成功", config.SUCCESSRES, user)
				return
			} else {
				utils.ResFaiJSON(&ctx, logErr.Error(), "微信登录失败", config.NOTGETDATA)
				return
			}
		} else {
			iNow := utils.GetNowTime()
			sAccessToken := tokenInfo["access_token"]
			sOpenId := tokenInfo["openid"]
			userInfo, getUserInfoOk := wxGetUserInfo(sAccessToken, sOpenId)
			if getUserInfoOk {
				var iMobileType = 1
				if sLoginFrom == "IOS" {
					iMobileType = 2
				}
				var newUserBean = xorm.DefaultUser()
				newUserBean.UserName = "wx" + utils.RandString(5, 2)
				newUserBean.Password = ""
				newUserBean.Name = userInfo["nickname"]
				newUserBean.Phone = ""
				newUserBean.ParentId = iParentId
				newUserBean.MobileType = iMobileType
				newUserBean.WxOpenId = sOpenId
				service.ip = utils.GetIp(ctx.Request())
				service.loginFrom = sLoginFrom
				var createErr error
				newUserBean, createErr = service.createNewUser(newUserBean, iNow)
				if createErr != nil {
					utils.ResFaiJSON(&ctx, createErr.Error(), "绑定微信失败", config.NOTGETDATA)
					return
				} else {
					user := new(xorm.Users)
					engine.ID(newUserBean.Id).Get(user)
					utils.ResSuccJSON(&ctx, "", "登录成功", config.SUCCESSRES, user)
					return
				}
			} else {
				utils.ResFaiJSON(&ctx, "1905301502", "绑定微信失败", config.NOTGETDATA)
				return
			}
		}
	} else {
		var internalMsg = "1905301706"
		var clientMsg = "微信授权失败"
		if _, fieldExist := tokenInfo["errcode"]; fieldExist {
			internalMsg = tokenInfo["errmsg"]
		}
		utils.ResFaiJSON(&ctx, internalMsg, clientMsg, config.NOTGETDATA)
		return
	}
}
