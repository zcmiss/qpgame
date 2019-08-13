package frontEndControllers

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/shopspring/decimal"
	"html"
	"math/rand"
	"qpgame/app/fund"
	"qpgame/common/services"
	"qpgame/common/utils"
	"qpgame/config"
	models2 "qpgame/models"
	"qpgame/models/xorm"
	"qpgame/ramcache"
	"regexp"
	"strconv"
	"strings"
)

type UserController struct {
	platform string
	ctx      iris.Context
}

//构造函数
func NewUserController(ctx iris.Context) *UserController {
	obj := new(UserController)
	obj.platform = ctx.Params().Get("platform")
	obj.ctx = ctx
	return obj
}

/**
 * @api {post} api/v1/register 注册会员（手机号）
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人:wenzhen</span><br/><br/>
 * 注册账号<br>
 * 业务描述:注册新会员</br>
 * @apiVersion 1.0.0
 * @apiName     register
 * @apiGroup    user
 * @apiPermission PC客户端
 * @apiParam (客户端请求参数) {string} phone    	手机号码
 * @apiParam (客户端请求参数) {string} password		密码
 * @apiParam (客户端请求参数) {string} vcode		手机验证码
 * @apiParam (客户端请求参数) {string} loginfrom	登录来源 IOS Android
 * @apiParam (客户端请求参数) {string} parentid		上级代理id
 *
 * @apiError (请求失败返回) {int}      code            错误代码
 * @apiError (请求失败返回) {string}   clientMsg       提示信息
 * @apiError (请求失败返回) {string}   internalMsg     错误代码
 * @apiError (请求失败返回) {float}    timeConsumed   后台耗时
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
 * @apiSuccess (返回结果)  {json}  	  data            返回数据
 * @apiSuccess (返回结果)  {float}   timeConsumed    后台耗时
 *
 *
 * @apiSuccessExample {json} 响应结果
 *
 *{
 *    "clientMsg": "注册成功",
 *    "code": 200,
 *    "data": {
 *        "Phone": "13412345678",
 *        "Username": "a13412345678",
 *        "Name": "",
 *        "Email": "",
 *        "Createtime": 1553478826,
 *        "Birthday": 0,
 *        "Mtype": 1,
 *        "Sex": 1,
 *        "Path": "",
 *        "Viplevel": 1,
 *        "Qq": "",
 *        "Weixin": "",
 *        "Status": 1,
 *        "Isdummy": 0,
 *        "Token": "",
 *        "TokenCreated": 0,
 *        "Safeboxpasswd": "",
 *        "Parentid": 1
 *    },
 *    "internalMsg": "注册成功",
 *    "timeConsumed": 597859
 *} *
 */
func (cthis *UserController) Register() {
	ctx := cthis.ctx

	if !utils.RequiredParamPost(&ctx, []string{"phone", "vcode", "password", "loginfrom"}) {
		return
	}

	// @updated by aTian 限制当前IP每天的注册数量
	sIp := utils.GetIp(ctx.Request())
	engine := models2.MyEngine[cthis.platform]
	conf, _ := ramcache.TableConfigs.Load(cthis.platform)
	cfg := conf.(map[string]interface{})
	iRegNumIp := int64(cfg["register_number_ip"].(float64))
	iFromTime, iToTime := utils.GetDatetimeRange(0, 1)
	var userBean xorm.Users
	iIpRegTotal, _ := engine.Where("reg_ip=? and created between ? and ?", sIp, iFromTime, iToTime).Count(&userBean)
	if iIpRegTotal >= iRegNumIp {
		utils.ResFaiJSON2(&ctx, "", "当前IP今天注册账号的数量已达上限")
		return
	}

	phone := ctx.FormValue("phone")
	password := ctx.FormValue("password")
	loginfrom := ctx.FormValue("loginfrom")
	code := ctx.FormValue("vcode")
	parentid := ctx.FormValue("parentid")
	iParentId, _ := strconv.Atoi(parentid)
	phoNT, _ := ramcache.PhoneNumAndToken.Load(cthis.platform)
	phoNumT := phoNT.(map[string][]string)
	if _, ok := phoNumT[phone]; ok {
		utils.ResFaiJSON(&ctx, "", "手机号码已存在", config.NOTGETDATA)
		return
	}
	vcode, err := strconv.Atoi(code)
	if err != nil {
		utils.ResFaiJSON2(&ctx, err.Error(), "无效的验证码")
		return
	}
	phonCCTemp := make(map[string][2]int)
	phonCC, existPhonCC := ramcache.PhoneCheckCode.Load(cthis.platform)
	if !existPhonCC {
		phonCC = phonCCTemp
	}
	if _, ok := phonCC.(map[string][2]int)[phone]; !ok {
		utils.ResFaiJSON(&ctx, "", "无效的验证码", config.NOTGETDATA)
		return
	}
	phonecode := xorm.PhoneCodes{Phone: phone}
	models2.MyEngine[cthis.platform].Where("phone = ?", phone).Limit(1).Desc("created").Get(&phonecode)
	//判断缓存和数据库里的验证码是否和当前输入的一致
	if phonCC.(map[string][2]int)[phone][0] == vcode && utils.GetNowTime() < phonCC.(map[string][2]int)[phone][1] ||
		(phonecode.Code == code && utils.GetNowTime() < phonecode.Created) {
		var user = xorm.DefaultUser()
		user.Phone = phone
		user.UserName = "a" + phone
		user.Password = utils.MD5(password)
		user.ParentId = iParentId
		if loginfrom == "IOS" {
			user.MobileType = 2
		} else {
			user.MobileType = 1
		}
		session := models2.MyEngine[cthis.platform].NewSession()
		defer session.Close()
		err := session.Begin()
		_, err = session.Insert(&user)
		if err != nil {
			utils.ResFaiJSON2(&ctx, err.Error(), "注册失败")
			session.Rollback()
			return
		}
		//注册成功之后，将用户手机号添加到缓存并将验证码缓存清空
		userid := strconv.Itoa(user.Id)
		delete(phonCC.(map[string][2]int), phone)
		ramcache.PhoneCheckCode.Store(cthis.platform, phonCC)
		_, err = session.Insert(xorm.Accounts{UserId: user.Id, Updated: utils.GetNowTime()})
		if err != nil {
			utils.ResFaiJSON2(&ctx, err.Error(), "注册失败")
			session.Rollback()
			return
		}
		user.UserName = "a" + strconv.Itoa(user.Id) + phone[7:]
		token, _ := utils.GenerateToken(&user)
		userNT, _ := ramcache.UserNameAndToken.Load(cthis.platform)
		userNameT := userNT.(map[string][]string)
		if _, ok := phoNumT[user.UserName]; ok {
			session.Rollback()
			utils.ResFaiJSON(&ctx, "用户名重复，请重试", "系统繁忙，请稍后再试", config.NOTGETDATA)
			return
		}
		user.Token = token
		user.TokenCreated = utils.GetNowTime()
		user.LastLoginTime = utils.GetNowTime()
		//登录成功以后将token缓存到本地
		tokentime := strconv.Itoa(user.TokenCreated)

		//开始事务
		_, err = session.ID(user.Id).Update(user)
		if err != nil {
			session.Rollback()
			utils.ResFaiJSON(&ctx, err.Error(), "登录失败1", config.NOTGETDATA)
			return
		}
		loginlog := xorm.UserLoginLogs{UserId: user.Id, Ip: utils.GetIp(ctx.Request()), LoginTime: utils.GetNowTime(), LoginFrom: loginfrom}
		_, err = session.InsertOne(loginlog)
		if err != nil {
			session.Rollback()
			utils.ResFaiJSON(&ctx, err.Error(), "登录失败2", config.NOTGETDATA)
			return
		}

		pErr := services.PromotionAward(cthis.platform, session, iParentId)
		if pErr != nil {
			utils.ResFaiJSON(&ctx, "1906251756", pErr.Error(), config.NOTGETDATA)
			return
		}

		err = services.BindPhoneAward(cthis.platform, session, user.Id)
		if err != nil {
			utils.ResFaiJSON(&ctx, "1906261656", err.Error(), config.NOTGETDATA)
			return
		}
		var innerMsg string
		innerMsg, err = services.ActivityAward(cthis.platform, session, 1, user.Id, sIp)
		if err != nil {
			session.Rollback()
			utils.ResFaiJSON(&ctx, innerMsg, err.Error(), config.NOTGETDATA)
			return
		}

		err = session.Commit()
		phoNumT[phone] = []string{userid, token, tokentime, "1"}
		ramcache.PhoneNumAndToken.Store(cthis.platform, phoNumT)

		userNameT[user.UserName] = []string{userid, token, tokentime, "1"}
		ramcache.UserNameAndToken.Store(cthis.platform, userNameT)
		utils.UpdateUserIdCard(cthis.platform, user.Id, map[string]interface{}{
			"Username":     user.UserName,
			"Token":        user.Token,
			"TokenCreated": tokentime, // 注意，这里要用字符串，否则会提示登录过期
		}) // user data ok
		utils.ResSuccJSON(&ctx, "注册成功", "注册成功", config.SUCCESSRES, user)
	} else {
		utils.ResFaiJSON(&ctx, "", "无效的验证码", config.NOTGETDATA)
	}
}

/**
 * @api {post} api/v1/registerUserName 注册会员(用户名)
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人:wenzhen</span><br/><br/>
 * 注册账号<br>
 * 业务描述:注册新会员</br>
 * @apiVersion 1.0.0
 * @apiName     registerUserName
 * @apiGroup    user
 * @apiPermission PC客户端
 * @apiParam (客户端请求参数) {string} username	用户名必须以字母开头，只能包含数字，字母，下划线，不允许输入特殊符号，长度5-15位
 * @apiParam (客户端请求参数) {string} password	密码
 * @apiParam (客户端请求参数) {string} vcode		验证码
 * @apiParam (客户端请求参数) {string} codeKey		获取图形验证码时候返回的key
 * @apiParam (客户端请求参数) {string} loginfrom	登录来源 IOS Android
 * @apiParam (客户端请求参数) {string} parentid	上级代理id
 *
 * @apiError (请求失败返回) {int}      code            错误代码
 * @apiError (请求失败返回) {string}   clientMsg       提示信息
 * @apiError (请求失败返回) {string}   internalMsg	错误代码
 * @apiError (请求失败返回) {float}    timeConsumed	后台耗时
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
 * @apiSuccess (返回结果)  {json}  	  data            返回数据
 * @apiSuccess (返回结果)  {float}   timeConsumed    后台耗时
 *
 *
 * @apiSuccessExample {json} 响应结果
 *
 *{
 *    "clientMsg": "注册成功",
 *    "code": 200,
 *    "data": {
 *        "Id": 130,
 *        "Phone": "",
 *        "Password": "b0baee9d279d34fa1dfd71aadb908c3f",
 *        "UserName": "a11111",
 *        "Name": "",
 *        "Email": "",
 *        "Created": 1558197749,
 *        "Birthday": "",
 *        "MobileType": 1,
 *        "Sex": 1,
 *        "Path": "",
 *        "VipLevel": 1,
 *        "Qq": "",
 *        "Wechat": "",
 *        "Status": 1,
 *        "ProxyStatus": 0,
 *        "IsDummy": 0,
 *        "Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NTg0NTY5NDksInBob25lIjoiIiwic3ViIjoxMzB9.eCxt3VGrfEe11Mnp62bR0dKZvnEwg-vszSjiHoCdPVw",
 *        "TokenCreated": 1558197749,
 *        "SafePassword": "",
 *        "UserGroupId": "",
 *        "ParentId": 0,
 *        "LastLoginTime": 1558197749,
 *        "LastPlatformId": 0,
 *        "GroupSize": 0
 *    },
 *    "internalMsg": "注册成功",
 *    "timeConsumed": 1097676
 *} *
 */
func (cthis *UserController) RegisterUserName() {
	ctx := cthis.ctx
	if !utils.RequiredParamPost(&ctx, []string{"username", "vcode", "codeKey", "password", "loginfrom"}) {
		return
	}
	// @updated by aTian 限制当前IP每天的注册数量
	sIp := utils.GetIp(ctx.Request())
	engine := models2.MyEngine[cthis.platform]
	conf, _ := ramcache.TableConfigs.Load(cthis.platform)
	cfg := conf.(map[string]interface{})
	iRegNumIp := int64(cfg["register_number_ip"].(float64))
	iFromTime, iToTime := utils.GetDatetimeRange(0, 1)
	var userBean xorm.Users
	iIpRegTotal, _ := engine.Where("reg_ip=? and created between ? and ?", sIp, iFromTime, iToTime).Count(&userBean)
	if iIpRegTotal >= iRegNumIp {
		utils.ResFaiJSON2(&ctx, "", "当前IP今天注册账号的数量已达上限")
		return
	}

	username := ctx.FormValue("username")
	match, _ := regexp.MatchString("(^[a-zA-Z][a-zA-Z0-9_]{4,15}$)", username)
	if !match {
		utils.ResFaiJSON(&ctx, "", "用户名必须以字母开头，只能包含数字，字母，下划线，不允许输入特殊符号，长度5-16位", config.NOTGETDATA)
		return
	}
	password := ctx.FormValue("password")
	loginfrom := ctx.FormValue("loginfrom")
	code := ctx.FormValue("vcode")
	codeKey := ctx.FormValue("codeKey")
	parentid := ctx.FormValue("parentid")
	iParentId, _ := strconv.Atoi(parentid)
	ut, _ := ramcache.UserNameAndToken.Load(cthis.platform)
	utMap := ut.(map[string][]string)
	if _, ok := utMap[username]; ok {
		utils.ResFaiJSON(&ctx, "", "用户名已存在", config.NOTGETDATA)
		return
	}
	PhCC, loadOk := ramcache.UserNameCheckCode.Load(cthis.platform)
	if loadOk == false {
		utils.ResFaiJSON(&ctx, "1906031335", "无效的验证码", config.NOTGETDATA)
		return
	}
	userNameCode, ok := PhCC.(map[string][2]string)[codeKey]
	if !ok {
		utils.ResFaiJSON(&ctx, "1906031336", "无效的验证码", config.NOTGETDATA)
		return
	}
	expTime, _ := strconv.Atoi(userNameCode[1])
	if utils.GetNowTime() > expTime {
		delete(PhCC.(map[string][2]string), codeKey)
		utils.ResFaiJSON(&ctx, "1906031337", "验证码已过期，请刷新重试", config.NOTGETDATA)
		return
	}
	//判断缓存和数据库里的验证码是否和当前输入的一致
	success := userNameCode[0] == code
	if success {
		delete(PhCC.(map[string][2]string), codeKey)
		//kafka推送，通知其它服务器清除缓存
		var user = xorm.DefaultUser()
		user.UserName = username
		user.Password = utils.MD5(password)
		user.ParentId = iParentId
		if loginfrom == "IOS" {
			user.MobileType = 2
		} else {
			user.MobileType = 1
		}
		session := models2.MyEngine[cthis.platform].NewSession()
		defer session.Close()
		err := session.Begin()
		_, err = session.Insert(&user)
		if err != nil {
			utils.ResFaiJSON2(&ctx, err.Error(), "注册失败")
			session.Rollback()
			return
		}
		//注册成功之后，将用户名添加到缓存并将验证码缓存清空
		userid := strconv.Itoa(user.Id)
		_, err = session.Insert(xorm.Accounts{UserId: user.Id, Updated: utils.GetNowTime()})
		if err != nil {
			utils.ResFaiJSON2(&ctx, err.Error(), "注册失败")
			session.Rollback()
			return
		}
		token, _ := utils.GenerateToken(&user)
		user.Token = token
		user.TokenCreated = utils.GetNowTime()
		user.LastLoginTime = utils.GetNowTime()
		//登录成功以后将token缓存到本地
		tokentime := strconv.Itoa(user.TokenCreated)

		//开始事务
		_, err = session.ID(user.Id).Update(user)
		if err != nil {
			session.Rollback()
			utils.ResFaiJSON(&ctx, err.Error(), "登录失败1", config.NOTGETDATA)
			return
		}
		loginlog := xorm.UserLoginLogs{UserId: user.Id, Ip: utils.GetIp(ctx.Request()), LoginTime: utils.GetNowTime(), LoginFrom: loginfrom}
		_, err = session.InsertOne(loginlog)
		if err != nil {
			session.Rollback()
			utils.ResFaiJSON(&ctx, err.Error(), "登录失败2", config.NOTGETDATA)
			return
		}

		pErr := services.PromotionAward(cthis.platform, session, iParentId)
		if pErr != nil {
			utils.ResFaiJSON(&ctx, "1906251818", pErr.Error(), config.NOTGETDATA)
			return
		}

		var innerMsg string
		innerMsg, err = services.ActivityAward(cthis.platform, session, 1, user.Id, sIp)
		if err != nil {
			session.Rollback()
			utils.ResFaiJSON(&ctx, innerMsg, err.Error(), config.NOTGETDATA)
			return
		}

		err = session.Commit()
		utMap[username] = []string{userid, token, tokentime, "1"}
		//ramcache.UserNameAndToken.Store(cthis.platform, phoNumT)
		utils.UpdateUserIdCard(cthis.platform, user.Id, map[string]interface{}{
			"Username":     username,
			"Token":        token,
			"TokenCreated": tokentime, // 注意，这里要用字符串，否则会提示登录过期
		})
		utils.ResSuccJSON(&ctx, "注册成功", "注册成功", config.SUCCESSRES, user)
	} else {
		utils.ResFaiJSON(&ctx, "", "无效的验证码", config.NOTGETDATA)
	}
}

/**
 * @api {post} api/v1/visitorLogin 游客登录
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人:wenzhen</span><br/><br/>
 * 注册账号<br>
 * 业务描述:注册新会员</br>
 * @apiVersion 1.0.0
 * @apiName     visitorLogin
 * @apiGroup    user
 * @apiPermission PC客户端
 * @apiParam (客户端请求参数) {string} uniqueCode    	手机唯一标识
 * @apiParam (客户端请求参数) {string} uniqueKey		加密字符串
 * @apiParam (客户端请求参数) {string} loginfrom	登录来源 IOS Android
 * @apiParam (客户端请求参数) {string} parentid		上级代理id
 *
 * @apiError (请求失败返回) {int}      code            错误代码
 * @apiError (请求失败返回) {string}   clientMsg       提示信息
 * @apiError (请求失败返回) {string}   internalMsg     错误代码
 * @apiError (请求失败返回) {float}    timeConsumed   后台耗时
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
 * @apiSuccess (返回结果)  {json}  	  data            返回数据
 * @apiSuccess (返回结果)  {float}   timeConsumed    后台耗时
 *
 *
 * @apiSuccessExample {json} 响应结果
 *
 *{
 *    "clientMsg": "注册成功",
 *    "code": 200,
 *    "data": {
 *        "Id": 130,
 *        "Phone": "",
 *        "Password": "b0baee9d279d34fa1dfd71aadb908c3f",
 *        "UserName": "a11111",
 *        "Name": "",
 *        "Email": "",
 *        "Created": 1558197749,
 *        "Birthday": "",
 *        "MobileType": 1,
 *        "Sex": 1,
 *        "Path": "",
 *        "VipLevel": 1,
 *        "Qq": "",
 *        "Wechat": "",
 *        "Status": 1,
 *        "ProxyStatus": 0,
 *        "IsDummy": 0,
 *        "Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NTg0NTY5NDksInBob25lIjoiIiwic3ViIjoxMzB9.eCxt3VGrfEe11Mnp62bR0dKZvnEwg-vszSjiHoCdPVw",
 *        "TokenCreated": 1558197749,
 *        "SafePassword": "",
 *        "UserGroupId": "",
 *        "ParentId": 0,
 *        "LastLoginTime": 1558197749,
 *        "LastPlatformId": 0,
 *        "GroupSize": 0
 *    },
 *    "internalMsg": "注册成功",
 *    "timeConsumed": 1097676
 *} *
 */
func (cthis *UserController) VisitorLogin() {
	ctx := cthis.ctx
	if !utils.RequiredParamPost(&ctx, []string{"uniqueCode", "uniqueKey", "loginfrom"}) {
		return
	}
	uniqueKey := ctx.FormValue("uniqueKey")
	password := "ckgame@20190525"
	loginfrom := ctx.FormValue("loginfrom")
	parentid := ctx.FormValue("parentid")
	iParentId, _ := strconv.Atoi(parentid)
	uniqueCode := ctx.FormValue("uniqueCode")
	success := utils.MD5(uniqueCode+"ckgame") == uniqueKey

	if !success {
		utils.ResFaiJSON(&ctx, "", "唯一标识有误，游客登录失败", config.NOTGETDATA)
		return
	}

	//加载所需要的缓存
	uniqueNT, _ := ramcache.UniqueCodeAndToken.Load(cthis.platform)
	uniqueCodeT := uniqueNT.(map[string][]string)

	userNT, _ := ramcache.UserNameAndToken.Load(cthis.platform)
	userNameT := userNT.(map[string][]string)

	conf, _ := ramcache.TableConfigs.Load(cthis.platform)
	cfg := conf.(map[string]interface{})

	// @updated by aTian 限制当前IP每天的注册数量
	sIp := utils.GetIp(ctx.Request())
	engine := models2.MyEngine[cthis.platform]

	iRegNumIp := int64(cfg["register_number_ip"].(float64))

	//如果已经存在游客用户，直接登录
	uniqueMap, ok := uniqueCodeT[uniqueCode]
	if ok {
		userid := uniqueMap[0]
		user, err := xorm.GetUser(cthis.platform, userid)
		if err != nil {
			utils.ResFaiJSON(&ctx, "", "游客登录失败，请尝试用户名密码登录", config.NOTGETDATA)
			return
		}
		if user.Status == 0 {
			utils.ResFaiJSON(&ctx, "", "当前用户已经锁定", config.NOTGETDATA)
			return
		}
		session := models2.MyEngine[cthis.platform].NewSession()
		if user.Id == 0 || user.UserType == 0 {
			session.Rollback()
			utils.ResFaiJSON(&ctx, "", "游客登录失败，请尝试用户名密码登录", config.NOTGETDATA)
			return
		}
		token, _ := utils.GenerateToken(&user)
		user.Token = token
		user.TokenCreated = utils.GetNowTime()
		user.LastLoginTime = utils.GetNowTime()
		//登录成功以后将token缓存到本地
		tokentime := strconv.Itoa(user.TokenCreated)

		//开始事务
		session.Begin()
		_, err = session.ID(user.Id).Update(user)
		if err != nil {
			session.Rollback()
			utils.ResFaiJSON(&ctx, err.Error(), "登录失败1", config.NOTGETDATA)
			return
		}
		loginlog := xorm.UserLoginLogs{UserId: user.Id, Ip: utils.GetIp(ctx.Request()), LoginTime: utils.GetNowTime(), LoginFrom: loginfrom}
		_, err = session.InsertOne(loginlog)
		if err != nil {
			session.Rollback()
			utils.ResFaiJSON(&ctx, err.Error(), "登录失败2", config.NOTGETDATA)
			return
		}
		err = session.Commit()
		username := user.UserName
		userNameT[username] = []string{userid, token, tokentime, "1"}
		uniqueCodeT[uniqueCode] = []string{userid, token, tokentime, "1"}

		// @updated by aTian 修复游客有手机号时，需要重新登录的问题
		if user.Phone != "" {
			pt, _ := ramcache.PhoneNumAndToken.Load(cthis.platform)
			ptMap := pt.(map[string][]string)
			ptMap[user.Phone] = []string{userid, token, tokentime, "1"}
		}
		utils.UpdateUserIdCard(cthis.platform, user.Id, map[string]interface{}{
			"Username":     user.UserName,
			"Token":        user.Token,
			"TokenCreated": tokentime,
			"UniqueCode":   uniqueCode,
		})
		utils.ResSuccJSON(&ctx, "登录成功", "登录成功", config.SUCCESSRES, user)
		return
	}
	//新游客登注册以后再登录
	username := "yk" + utils.RandString(4, 1)
	iFromTime, iToTime := utils.GetDatetimeRange(0, 1)
	var userBean xorm.Users
	iIpRegTotal, _ := engine.Where("reg_ip=? and created between ? and ?", sIp, iFromTime, iToTime).Count(&userBean)
	if iIpRegTotal >= iRegNumIp {
		utils.ResFaiJSON2(&ctx, "", "当前IP今天注册账号的数量已达上限")
		return
	}

	//判断一下生成的用户名是否重复
	if _, ok := userNameT[username]; ok {
		utils.ResFaiJSON(&ctx, "", "用户名已存在", config.NOTGETDATA)
		return
	}

	var user = xorm.DefaultUser()
	user.UserName = username
	user.RegIp = sIp
	user.UserType = 1
	user.Password = utils.MD5(password)
	user.ParentId = iParentId
	user.UniqueCode = uniqueCode
	if loginfrom == "IOS" {
		user.MobileType = 2
	} else {
		user.MobileType = 1
	}
	session := models2.MyEngine[cthis.platform].NewSession()
	defer session.Close()
	err := session.Begin()
	_, err = session.Insert(&user)
	if err != nil {
		utils.ResFaiJSON2(&ctx, err.Error(), "注册失败")
		session.Rollback()
		return
	}
	//注册成功之后，将用户名添加到缓存并将验证码缓存清空
	userid := strconv.Itoa(user.Id)
	_, err = session.Insert(xorm.Accounts{UserId: user.Id, Updated: utils.GetNowTime()})
	if err != nil {
		utils.ResFaiJSON2(&ctx, err.Error(), "注册失败")
		session.Rollback()
		return
	}
	token, _ := utils.GenerateToken(&user)
	user.Token = token
	user.TokenCreated = utils.GetNowTime()
	user.LastLoginTime = utils.GetNowTime()
	//登录成功以后将token缓存到本地
	tokentime := strconv.Itoa(user.TokenCreated)

	//开始事务
	_, err = session.ID(user.Id).Update(user)
	if err != nil {
		session.Rollback()
		utils.ResFaiJSON(&ctx, err.Error(), "登录失败1", config.NOTGETDATA)
		return
	}
	loginlog := xorm.UserLoginLogs{UserId: user.Id, Ip: utils.GetIp(ctx.Request()), LoginTime: utils.GetNowTime(), LoginFrom: loginfrom}
	_, err = session.InsertOne(loginlog)
	if err != nil {
		session.Rollback()
		utils.ResFaiJSON(&ctx, err.Error(), "登录失败2", config.NOTGETDATA)
		return
	}

	pErr := services.PromotionAward(cthis.platform, session, iParentId)
	if pErr != nil {
		utils.ResFaiJSON(&ctx, "1906251822", pErr.Error(), config.NOTGETDATA)
		return
	}

	var innerMsg string
	innerMsg, err = services.ActivityAward(cthis.platform, session, 1, user.Id, sIp)
	if err != nil {
		session.Rollback()
		utils.ResFaiJSON(&ctx, innerMsg, err.Error(), config.NOTGETDATA)
		return
	}

	err = session.Commit()
	userNameT[username] = []string{userid, token, tokentime, "1"}
	uniqueCodeT[uniqueCode] = []string{userid, token, tokentime, "1"}
	utils.UpdateUserIdCard(cthis.platform, user.Id, map[string]interface{}{
		"Username":     user.UserName,
		"Token":        user.Token,
		"TokenCreated": user.TokenCreated,
		"UniqueCode":   user.UniqueCode,
	})
	utils.ResSuccJSON(&ctx, "注册成功", "注册成功", config.SUCCESSRES, user)
}

/**
 * @api {get} api/v1/getCode 获取图形验证码
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人:wenzhen</span><br/><br/>
 * 获取图片验证码的base64图片地址<br>
 * 业务描述:获取图片验证码的base64图片地址</br>
 * @apiVersion 1.0.0
 * @apiName     getCode
 * @apiGroup    user
 * @apiPermission PC客户端
 * @apiError (请求失败返回) {int}      code            错误代码
 * @apiError (请求失败返回) {string}   clientMsg       提示信息
 * @apiError (请求失败返回) {string}   internalMsg     错误代码
 * @apiError (请求失败返回) {float}    timeConsumed   后台耗时
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
 * @apiSuccess (返回结果)  {json}  	  data            返回数据
 * @apiSuccess (返回结果)  {float}   timeConsumed    后台耗时
 *
 *
 * @apiSuccessExample {json} 响应结果
 *
 *{
 *    "clientMsg": "获取成功",
 *    "code": 200,
 *    "data": {
 *        "baseCodes": "data:image/png;base64,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",
 *        "codeKey": "VMt66VvFJDt3Bi-IpAy-oMY3EU-iirRocp2yaQ1JcKE="
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 7978
 *}
 */
func (cthis *UserController) GetCode() {
	ctx := cthis.ctx
	result := make(map[string]string)
	idKey, baseCodes := utils.CreateCaptchaNum()
	resKey := utils.PswEncrypt(idKey+strconv.Itoa(utils.GetNowTime()), config.CodeKey, config.CodeIv)
	PhccPlat := make(map[string][2]string)
	//如果平台容器还不存在就创建一个map
	PhCC, existPlat := ramcache.UserNameCheckCode.Load(cthis.platform)
	if !existPlat {
		PhCC = PhccPlat
	}
	//缓存验证码60秒过期
	codeArr := [2]string{idKey, strconv.Itoa(utils.GetNowTime() + 60)}
	//需要进行kafka推送
	//代码需要另外补充

	PhCC.(map[string][2]string)[resKey] = codeArr
	ramcache.UserNameCheckCode.Store(cthis.platform, PhCC)
	result["codeKey"] = resKey
	result["baseCodes"] = baseCodes
	utils.ResSuccJSON(&ctx, "", "获取成功", config.SUCCESSRES, result)
}

/**
 * @api {get} api/v1/getVcode 获取手机验证码
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人:wenzhen</span><br/><br/>
 * 发送手机验证码<br>
 * 业务描述:获取手机验证码</br>
 * @apiVersion 1.0.0
 * @apiName     getVcode
 * @apiGroup    user
 * @apiPermission PC客户端
 * @apiParam (客户端请求参数) {string} phone    	手机号码
 * @apiError (请求失败返回) {int}      code            错误代码
 * @apiError (请求失败返回) {string}   clientMsg       提示信息
 * @apiError (请求失败返回) {string}   internalMsg     错误代码
 * @apiError (请求失败返回) {float}    timeConsumed   后台耗时
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
 * @apiSuccess (返回结果)  {json}  	  data            返回数据
 * @apiSuccess (返回结果)  {float}   timeConsumed    后台耗时
 *
 *
 * @apiSuccessExample {json} 响应结果
 *
 *  {
 *      "code": 200,
 *      "internalMsg": "",
 *      "clientMsg": "发送成功",
 *      "data": {
 *          "vocde": 123456
 *      },
 *      "timeConsumed": 45
 *  }
 */
func (cthis *UserController) GetVcode() {
	ctx := cthis.ctx
	if !utils.RequiredParam(&ctx, []string{"phone"}) {
		return
	}
	phoNT, _ := ramcache.PhoneNumAndToken.Load(cthis.platform)
	phoNumT := phoNT.(map[string][]string)
	phone := ctx.URLParam("phone")
	match, _ := regexp.MatchString("(^([1][3,4,5,6,7,8,9])\\d{9}$)", phone)
	if !match {
		utils.ResFaiJSON(&ctx, "", "请填写正确的手机号", config.NOTGETDATA)
		return
	}
	if _, ok := phoNumT[phone]; ok {
		utils.ResFaiJSON(&ctx, "", "手机号码已存在", config.NOTGETDATA)
		return
	}
	vcode, e := utils.SendSms(phone, cthis.platform)
	if e != nil {
		utils.ResFaiJSON2(&ctx, e.Error(), "获取验证码失败")
		return
	}
	var result iris.Map
	//开发环境将验证码返回给前端进行验证
	if config.DevloperDebug {
		result = iris.Map{"vcode": vcode}
	} else {
		result = iris.Map{"vcode": vcode}
	}
	go models2.MyEngine[cthis.platform].Insert(xorm.PhoneCodes{Phone: phone, Code: vcode, Created: utils.GetNowTime() + 180})
	utils.ResSuccJSON(&ctx, "", "验证码已发送", config.SUCCESSRES, result)
}

/**
 * @api {post} api/v1/login 会员登录
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人:wenzhen</span><br/><br/>
 * 会员登录<br>
 * 业务描述:会员登录</br>
 * @apiVersion 1.0.0
 * @apiName     login
 * @apiGroup    user
 * @apiPermission PC客户端
 * @apiParam (客户端请求参数) {string} phone    	手机号或者用户名
 * @apiParam (客户端请求参数) {string} password    	密码
 * @apiParam (客户端请求参数) {string} loginfrom    来源（IOS，ANDROID）
 * @apiError (请求失败返回) {int}      code            错误代码
 * @apiError (请求失败返回) {string}   clientMsg       提示信息
 * @apiError (请求失败返回) {string}   internalMsg     错误代码
 * @apiError (请求失败返回) {float}    timeConsumed   后台耗时
 *
 * @apiErrorExample {json} 失败返回
 * {
 *      "status": 0,
 *      "msg": "",
 *      "code ": 0,
 *      "timeConsumed": 0
 * }
 *
 * @apiSuccess (返回结果)  {int}      code            200
 * @apiSuccess (返回结果)  {string}   clientMsg       提示信息
 * @apiSuccess (返回结果)  {string}   internalMsg     提示信息
 * @apiSuccess (返回结果)  {json}  	  data            返回数据
 * @apiSuccess (返回结果)  {float}   timeConsumed    后台耗时
 *
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "登录成功",
 *    "code": 200,
 *    "data": {
 *        "Phone": "13912345678",
 *        "UserName": "a13912345678",
 *        "Name": "13912345678",
 *        "Email": "",
 *        "Created": 1553582249,
 *        "Birthday": 0,
 *        "MobileType": 1,
 *        "Sex": 1,
 *        "Path": "",
 *        "VipLevel": 1,
 *        "Qq": "",
 *        "Wechat": "",
 *        "Status": 1,
 *        "IsDummy": 0,
 *        "Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjI4ODAwMDAwMDAwMDAwLCJwaG9uZSI6IjEzOTEyMzQ1Njc4Iiwic3ViIjoyOH0.D87HVJst5hUE2mzZA9iFFi6g8CRAn66wDGtM5CAgeRM",
 *        "TokenCreated": 1554358614,
 *        "SafePassword": "81dc9bdb52d04dc20036dbd8313ed055",
 *        "ParentId": 1,
 *        "LastLoginTime": 1554358614,
 *        "LastPlatformId": 0,
 *        "GroupSize": 0
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 977061
 *}
 */
func (cthis *UserController) Login() {
	ctx := cthis.ctx
	if !utils.RequiredParamPost(&ctx, []string{"phone", "password", "loginfrom"}) {
		return
	}
	phone := ctx.FormValue("phone")
	password := ctx.FormValue("password")
	loginfrom := ctx.FormValue("loginfrom")
	var phoNT interface{}
	match, _ := regexp.MatchString("(^[a-zA-Z][a-zA-Z0-9_]{4,15}$)", phone)
	var loginAccountType int // 1为用户名登录， 2为手机号登录
	if match {
		loginAccountType = 1
		phoNT, _ = ramcache.UserNameAndToken.Load(cthis.platform)
	} else {
		loginAccountType = 2
		phoNT, _ = ramcache.PhoneNumAndToken.Load(cthis.platform)
	}
	phoNumT := phoNT.(map[string][]string)
	_, ok := phoNumT[phone]
	if ok {
		var user, err = xorm.GetUser(cthis.platform, phone)
		if err != nil {
			utils.ResFaiJSON2(&ctx, err.Error(), "登录失败111")
			return
		}
		if user.Status == 0 {
			utils.ResFaiJSON(&ctx, "", "当前用户已经锁定", config.NOTGETDATA)
			return
		}

		if user.Password == utils.MD5(password) {
			token, _ := utils.GenerateToken(&user)
			user.Token = token
			user.TokenCreated = utils.GetNowTime()
			user.LastLoginTime = utils.GetNowTime()
			//登录成功以后将token缓存到本地
			userid := strconv.Itoa(user.Id)
			tokentime := strconv.Itoa(user.TokenCreated)
			//开始事务
			session := models2.MyEngine[cthis.platform].NewSession()
			defer session.Close()
			err := session.Begin()
			_, err = session.ID(user.Id).Update(user)
			if err != nil {
				session.Rollback()
				utils.ResFaiJSON(&ctx, err.Error(), "登录失败1", config.NOTGETDATA)
				return
			}
			loginlog := xorm.UserLoginLogs{UserId: user.Id, Ip: utils.GetIp(ctx.Request()), LoginTime: utils.GetNowTime(), LoginFrom: loginfrom}
			_, err = session.InsertOne(loginlog)
			if err != nil {
				session.Rollback()
				utils.ResFaiJSON(&ctx, err.Error(), "登录失败2", config.NOTGETDATA)
				return
			}
			err = session.Commit()
			if err != nil {
				utils.ResFaiJSON(&ctx, err.Error(), "登录失败3", config.NOTGETDATA)
				return
			}
			phoNumT[phone] = []string{userid, token, tokentime, "1"}

			// @update by aTian 20190601
			// 通过登录方式去修改其它登录方式的缓存，否则会导致其它登录方式失败
			if loginAccountType == 1 {
				// 1为通过用户名登录
				if user.Phone != "" {
					pt, _ := ramcache.PhoneNumAndToken.Load(cthis.platform)
					ptMap := pt.(map[string][]string)
					ptMap[user.Phone] = []string{userid, token, tokentime, "1"}
				}
			} else if loginAccountType == 2 {
				// 2为通过手机号登录
				ut, _ := ramcache.UserNameAndToken.Load(cthis.platform)
				utMap := ut.(map[string][]string)
				utMap[user.UserName] = []string{userid, token, tokentime, "1"}
			}

			utils.UpdateUserIdCard(cthis.platform, user.Id, map[string]interface{}{
				"Token":        token,
				"TokenCreated": tokentime,
			})
			utils.ResSuccJSON(&ctx, "", "登录成功", config.SUCCESSRES, user)
		} else {
			utils.ResFaiJSON(&ctx, "", "账号或密码错误", config.NOTGETDATA)
		}
	} else {
		utils.ResFaiJSON(&ctx, "", "账号或密码错误", config.NOTGETDATA)
	}
}

/**
 * @api {post} api/auth/v1/modifyPwd 修改用户密码
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人:wenzhen</span><br/><br/>
 * 修改登录密码<br>
 * 业务描述:修改登录密码</br>
 * @apiVersion 1.0.0
 * @apiName     modifyPwd
 * @apiGroup    user
 * @apiPermission iso,android客户端
 * @apiHeader (客户端请求头参数) {string} Authorization Bearer + 用户登录获得的token
 * @apiParam (客户端请求参数) {string} oldpwd    	原密码
 * @apiParam (客户端请求参数) {string} password    	新密码
 *
 * @apiError (请求失败返回) {int}      code            错误代码
 * @apiError (请求失败返回) {string}   clientMsg       提示信息
 * @apiError (请求失败返回) {string}   internalMsg     错误代码
 * @apiError (请求失败返回) {float}    timeConsumed   后台耗时
 *
 * @apiErrorExample {json} 失败返回
 *  {
 *      "code": 204,
 *      "clientMsg": "失败",
 *      "internalMsg": "",
 *      "timeConsumed": 1785
 *  }
 *
 * @apiSuccess (返回结果)  {int}      code            200
 * @apiSuccess (返回结果)  {string}   clientMsg       提示信息
 * @apiSuccess (返回结果)  {string}   internalMsg     提示信息
 * @apiSuccess (返回结果)  {json}  	  data            返回数据
 * @apiSuccess (返回结果)  {float}   timeConsumed    后台耗时
 * @apiSuccessExample {json} 响应结果
 *
 *  {
 *      "code": 200,
 *      "clientMsg": "修改成功",
 *      "internalMsg": "",
 *      "data": {},
 *      "timeConsumed": 1785
 *  }
 *
 */
func (cthis *UserController) ModifyPwd() {
	ctx := cthis.ctx
	if !utils.RequiredParamPost(&ctx, []string{"oldpwd", "password"}) {
		return
	}
	userid, _ := strconv.Atoi(ctx.Values().GetString("userid"))
	oldpwd := ctx.FormValue("oldpwd")
	password := ctx.FormValue("password")
	engine := models2.MyEngine[cthis.platform]
	var user = new(xorm.Users)
	_, err := engine.ID(userid).Cols("`password`").Get(user)
	if err != nil {
		utils.ResFaiJSON(&ctx, "1906191520", "修改密码失败，请重试", config.NOTGETDATA)
		return
	}
	if user.Password != utils.MD5(oldpwd) {
		utils.ResFaiJSON(&ctx, "", "原密码输入有误，请重新输入", config.NOTGETDATA)
		return
	}
	var affNum int64
	affNum, err = engine.ID(userid).Update(xorm.Users{Id: userid, Password: utils.MD5(password)})
	if err != nil {
		utils.ResFaiJSON(&ctx, "1906191521", "修改密码失败，请重试", config.NOTGETDATA)
		return
	}
	if affNum == 1 {
		utils.ResSuccJSON(&ctx, "", "密码修改成功", config.SUCCESSRES, "")
	} else {
		utils.ResFaiJSON(&ctx, "", "新旧密码不能相同", config.NOTGETDATA)
	}
}

/**
 * @api {get} api/auth/v1/logout 退出登录
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人:wenzhen</span><br/><br/>
 * 退出登录<br>
 * 业务描述:会员退出当前登录状态</br>
 * @apiVersion 1.0.0
 * @apiName     logout
 * @apiGroup    user
 * @apiPermission iso,android客户端
 * @apiHeader (客户端请求头参数) {string} Authorization Bearer + 用户登录获得的token
 * @apiError (请求失败返回) {int}      code            错误代码
 * @apiError (请求失败返回) {string}   clientMsg       提示信息
 * @apiError (请求失败返回) {string}   internalMsg     错误代码
 * @apiError (请求失败返回) {float}    timeConsumed   后台耗时
 *
 * @apiErrorExample {json} 失败返回
 *  {
 *      "code": 204,
 *      "clientMsg": "失败",
 *      "internalMsg": "",
 *      "data": {},
 *      "timeConsumed": 1785
 *  }
 *
 * @apiSuccess (返回结果)  {int}      code            200
 * @apiSuccess (返回结果)  {string}   clientMsg       提示信息
 * @apiSuccess (返回结果)  {string}   internalMsg     提示信息
 * @apiSuccess (返回结果)  {json}  	  data            返回数据
 * @apiSuccess (返回结果)  {float}   timeConsumed    后台耗时
 * @apiSuccessExample {json} 响应结果
 *
 *  {
 *      "code": 200,
 *      "clientMsg": "退出成功",
 *      "internalMsg": "",
 *      "data": {},
 *      "timeConsumed": 1785
 *  }
 *
 */
func (cthis *UserController) Logout() {
	ctx := cthis.ctx
	userId := ctx.Values().GetString("userid")
	phone := ctx.Values().GetString("phone")
	username := ctx.Values().GetString("username")
	go models2.MyEngine[cthis.platform].Exec("update users set token='',token_created='0' where id=?", userId)
	ut, _ := ramcache.UserNameAndToken.Load(cthis.platform)
	utMap := ut.(map[string][]string)
	if uerArr, userExist := utMap[username]; userExist && len(uerArr) == 4 {
		utMap[username] = []string{userId, "", "0", uerArr[3]}
	}
	if phone != "" {
		pt, _ := ramcache.PhoneNumAndToken.Load(cthis.platform)
		ptMap := pt.(map[string][]string)
		if phoneArr, phoneExist := ptMap[phone]; phoneExist && len(phoneArr) == 4 {
			ptMap[phone] = []string{userId, "", "0", phoneArr[3]}
		}
	}
	iUserId, _ := strconv.Atoi(userId)
	defer utils.UpdateUserIdCard(cthis.platform, iUserId, map[string]interface{}{
		"Token":        "",
		"TokenCreated": "0",
	})
	utils.ResSuccJSON(&ctx, "", "退出成功", config.SUCCESSRES, make(map[string]interface{}))
}

/**
* @api {post} api/auth/v1/setSafeBoxPwd 设置保险箱密码
* @apiDescription
* <span style="color:lightcoral;">接口负责人:wenzhen</span><br/><br/>
* 设置保险箱密码<br>
* 业务描述:设置保险箱密码</br>
* @apiVersion 1.0.0
* @apiName     setSafeBoxPwd
* @apiGroup    safeBox
 * @apiPermission iso,android客户端
* @apiHeader (客户端请求头参数) {string} Authorization Bearer + 用户登录获得的token
* @apiParam (客户端请求参数) {string} safepassword    	新密码
*
* @apiError (请求失败返回) {int}      code            错误代码
* @apiError (请求失败返回) {string}   clientMsg       提示信息
* @apiError (请求失败返回) {string}   internalMsg     错误代码
* @apiError (请求失败返回) {float}    timeConsumed   后台耗时
*
* @apiErrorExample {json} 失败返回
*  {
*      "code": 204,
*      "clientMsg": "失败",
*      "internalMsg": "",
*      "timeConsumed": 1785
*  }
*
* @apiSuccess (返回结果)  {int}      code            200
* @apiSuccess (返回结果)  {string}   clientMsg       提示信息
* @apiSuccess (返回结果)  {string}   internalMsg     提示信息
* @apiSuccess (返回结果)  {json}  	  data            返回数据
* @apiSuccess (返回结果)  {float}   timeConsumed    后台耗时
* @apiSuccessExample {json} 响应结果
*
*{
*    "clientMsg": "保险箱密码设置成功",
*    "code": 200,
*    "data": "{}",
*    "internalMsg": "",
*    "timeConsumed": 578046
*}
*
*/
func (cthis *UserController) SetSafeBoxPwd() {
	ctx := cthis.ctx
	if !utils.RequiredParamPost(&ctx, []string{"safepassword"}) {
		return
	}
	userid, _ := strconv.Atoi(ctx.Values().GetString("userid"))
	safepassword := ctx.FormValue("safepassword")
	user := xorm.Users{Id: userid, SafePassword: utils.MD5(safepassword)}
	res, err := models2.MyEngine[cthis.platform].Id(userid).Cols("safe_password").Update(user)
	if err != nil {
		fmt.Println(err.Error())
		utils.ResFaiJSON(&ctx, "", "保险密码设置失败", config.NOTGETDATA)
	}
	if res == 1 {
		utils.ResSuccJSON(&ctx, "", "保险箱密码设置成功", config.SUCCESSRES, "{}")
	} else {
		utils.ResFaiJSON(&ctx, "", "保险箱密码设置失败", config.NOTGETDATA)
	}
}

/**
* @api {post} api/auth/v1/intoSafeBox 进入保险箱
* @apiDescription
* <span style="color:lightcoral;">接口负责人:wenzhen</span><br/><br/>
* 进入保险箱<br>
* 业务描述:进入保险箱</br>
* @apiVersion 1.0.0
* @apiName     intoSafeBox
* @apiGroup    safeBox
 * @apiPermission iso,android客户端
* @apiHeader (客户端请求头参数) {string} Authorization Bearer + 用户登录获得的token
* @apiParam (客户端请求参数) {string} safepassword    	保险箱密码
*
* @apiError (请求失败返回) {int}      code            错误代码
* @apiError (请求失败返回) {string}   clientMsg       提示信息
* @apiError (请求失败返回) {string}   internalMsg     错误代码
* @apiError (请求失败返回) {float}    timeConsumed   后台耗时
*
* @apiErrorExample {json} 失败返回
*{
*    "clientMsg": "保险箱密码不正确",
*    "code": 204,
*    "internalMsg": "",
*    "timeConsumed": 649523
*} *
*
* @apiSuccess (返回结果)  {int}      code            200
* @apiSuccess (返回结果)  {string}   clientMsg       提示信息
* @apiSuccess (返回结果)  {string}   internalMsg     提示信息
* @apiSuccess (返回结果)  {json}  	  data            返回数据
* @apiSuccess (返回结果)  {float}   timeConsumed    后台耗时
* @apiSuccessExample {json} 响应结果
*
*{
*    "clientMsg": "进入成功",
*    "code": 200,
*    "data": {
*        "Id": 2,
*        "UserId": 28,
*        "ChargedAmount": "0.000",
*        "ConsumedAmount": "0.000",
*        "Balance": "0.000",
*        "WithdrawAmount": "0.000",
*        "BalanceCharge": "0.000",
*        "BalanceLucky": "0.000",
*        "BalanceSafe": "0.000",
*        "BalanceWallet": "12000.000",
*        "Updated": 1553589822
*    },
*    "internalMsg": "",
*    "timeConsumed": 496886
*}
*
*/
func (cthis *UserController) IntoSafeBox() {
	ctx := cthis.ctx
	if !utils.RequiredParamPost(&ctx, []string{"safepassword"}) {
		return
	}
	userid, _ := strconv.Atoi(ctx.Values().GetString("userid"))
	safepassword := ctx.FormValue("safepassword")
	user, err := xorm.GetUser(cthis.platform, userid)
	if err != nil {
		utils.ResFaiJSON(&ctx, "", "操作出现异常", config.NOTGETDATA)
		return
	}
	if user.Status == 0 {
		utils.ResFaiJSON(&ctx, "", "当前用户已经锁定", config.NOTGETDATA)
		return
	}
	if user.SafePassword == utils.MD5(safepassword) {
		account := xorm.Accounts{UserId: userid}
		res, _ := models2.MyEngine[cthis.platform].Where("user_id = ?", userid).Get(&account)
		if res {
			utils.ResSuccJSON(&ctx, "", "进入成功", config.SUCCESSRES, account)
		} else {
			utils.ResFaiJSON(&ctx, "", "账号信息获取失败", config.NOTGETDATA)
		}
	} else {
		utils.ResFaiJSON(&ctx, "", "保险箱密码不正确", config.NOTGETDATA)
	}
}

/**
 * @api {post} api/auth/v1/safeBoxOperation 保险箱存取款操作
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人:wenzhen</span><br/><br/>
 * 保险箱存取款操作<br>
 * 业务描述:保险箱存取款操作</br>
 * @apiVersion 1.0.0
 * @apiName     safeBoxOperation
 * @apiGroup    safeBox
 * @apiPermission iso,android客户端
 * @apiHeader (客户端请求头参数) {string} Authorization Bearer + 用户登录获得的token
 * @apiParam (客户端请求参数) {int}    amount    	   金额（整数，钱包余额减去保险箱余额，正数为取出，负数为存入）
 *
 * @apiError (请求失败返回) {int}      code            错误代码
 * @apiError (请求失败返回) {string}   clientMsg       提示信息
 * @apiError (请求失败返回) {string}   internalMsg     错误代码
 * @apiError (请求失败返回) {float}    timeConsumed   后台耗时
 *
 * @apiErrorExample {json} 失败返回
 *{
 *    "clientMsg": "金额输入有误，请重新输入",
 *    "code": 204,
 *    "internalMsg": "",
 *    "timeConsumed": 182507
 *} *
 *
 * @apiSuccess (返回结果)  {int}      code            200
 * @apiSuccess (返回结果)  {string}   clientMsg       提示信息
 * @apiSuccess (返回结果)  {string}   internalMsg     提示信息
 * @apiSuccess (返回结果)  {json}  	  data            返回数据
 * @apiSuccess (返回结果)  {float}   timeConsumed    后台耗时
 * @apiSuccessExample {json} 响应结果
 *
 *{
 *    "clientMsg": "操作成功",
 *    "code": 200,
 *    "data": {
 *        "Id": 2,
 *        "UserId": 28,
 *        "ChargedAmount": "0.000",
 *        "ConsumedAmount": "0.000",
 *        "Balance": "0.000",
 *        "WithdrawAmount": "0.000",
 *        "BalanceCharge": "0.000",
 *        "BalanceLucky": "0.000",
 *        "BalanceSafe": "100",
 *        "BalanceWallet": "11900",
 *        "Updated": 1553591581
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 388055
 *}
 *
 */
func (cthis *UserController) SafeBoxOperation() {
	ctx := cthis.ctx
	if !utils.RequiredParamPost(&ctx, []string{"amount"}) {
		return
	}
	amount, err := decimal.NewFromString(ctx.FormValue("amount"))
	if err != nil {
		utils.ResFaiJSON(&ctx, "", "金额输入有误", config.NOTGETDATA)
		return
	}
	userid, _ := strconv.Atoi(ctx.Values().GetString("userid"))
	account := xorm.Accounts{UserId: userid}
	_, aerr := models2.MyEngine[cthis.platform].Where("user_id = ?", userid).Get(&account)
	if aerr != nil {
		panic(aerr.Error())
		utils.ResFaiJSON2(&ctx, aerr.Error(), "获取账号信息失败")
	}
	balanceSafe, _ := decimal.NewFromString(account.BalanceSafe)
	balanceWallet, _ := decimal.NewFromString(account.BalanceWallet)

	var iserr = true
	//如果是正数,为取款操作，要保证输入金额小于等于保险箱余额
	if amount.GreaterThan(decimal.Zero) && amount.LessThanOrEqual(balanceSafe) {
		iserr = false
	}
	//如果是负数,为存款操作，要保证输入金额小于等于钱包余额
	if amount.LessThan(decimal.Zero) && amount.Abs().LessThanOrEqual(balanceWallet) {
		iserr = false
	}

	if !iserr {
		amountf, _ := amount.Float64()
		//执行存入保险箱的操作
		info := map[string]interface{}{
			"user_id":     userid,
			"type_id":     config.FUNDSAFEBOX,
			"amount":      amountf,
			"order_id":    utils.CreationOrder("BXX", strconv.Itoa(userid)),
			"msg":         "保险箱存取款",
			"finish_rate": 1.0, //需满足的打码量比例
		}
		res := fund.NewUserFundChange(cthis.platform).BalanceUpdate(info, nil)
		if res["status"] == 1 {
			utils.ResSuccJSON(&ctx, "", "操作成功", config.SUCCESSRES, res["accounts"])
		}
	} else {
		utils.ResFaiJSON(&ctx, "", "金额输入有误，请重新输入", config.NOTGETDATA)

	}
}

/**
* @api {get} api/auth/v1/safeBoxInfo 保险箱明细
* @apiDescription
* <span style="color:lightcoral;">接口负责人:wenzhen</span><br/><br/>
* 保险箱存取款明细<br>
* 业务描述:保险箱存取款明细</br>
* @apiVersion 1.0.0
* @apiName     safeBoxInfo
* @apiGroup    safeBox
 * @apiPermission iso,android客户端
* @apiHeader (客户端请求头参数) {string} Authorization Bearer + 用户登录获得的token
* @apiParam (客户端请求参数) {int}    pageNum    	   当前页
* @apiError (请求失败返回) {int}      code            错误代码
* @apiError (请求失败返回) {string}   clientMsg       提示信息
* @apiError (请求失败返回) {string}   internalMsg     错误代码
* @apiError (请求失败返回) {float}    timeConsumed   后台耗时
*
* @apiErrorExample {json} 失败返回
*{
*    "clientMsg": "金额输入有误，请重新输入",
*    "code": 204,
*    "internalMsg": "",
*    "timeConsumed": 182507
*} *
*
* @apiSuccess (返回结果)  {int}      code            200
* @apiSuccess (返回结果)  {string}   clientMsg       提示信息
* @apiSuccess (返回结果)  {string}   internalMsg     提示信息
* @apiSuccess (返回结果)  {json}  	  data            返回数据
* @apiSuccess (返回结果)  {float}   timeConsumed    后台耗时
* @apiSuccessExample {json} 响应结果
*
*{
*    "clientMsg": "获取成功",
*    "code": 200,
*    "data": [
*        {
*            "Id": 81,
*            "UserId": 28,
*            "Amount": "-300.000",
*            "Blance": "0.000",
*            "Type": 3,
*            "Created": 1553594523
*        },
*        {
*            "Id": 80,
*            "UserId": 28,
*            "Amount": "-300.000",
*            "Blance": "300.000",
*            "Type": 3,
*            "Created": 1553594522
*        },
*        {
*            "Id": 79,
*            "UserId": 28,
*            "Amount": "-300.000",
*            "Blance": "600.000",
*            "Type": 3,
*            "Created": 1553594521
*        },
*        {
*            "Id": 78,
*            "UserId": 28,
*            "Amount": "-300.000",
*            "Blance": "900.000",
*            "Type": 3,
*            "Created": 1553594520
*        },
*        {
*            "Id": 77,
*            "UserId": 28,
*            "Amount": "-300.000",
*            "Blance": "1200.000",
*            "Type": 3,
*            "Created": 1553594519
*        },
*        {
*            "Id": 76,
*            "UserId": 28,
*            "Amount": "-300.000",
*            "Blance": "1500.000",
*            "Type": 3,
*            "Created": 1553594518
*        },
*        {
*            "Id": 74,
*            "UserId": 28,
*            "Amount": "-300.000",
*            "Blance": "2100.000",
*            "Type": 3,
*            "Created": 1553594517
*        },
*        {
*            "Id": 75,
*            "UserId": 28,
*            "Amount": "-300.000",
*            "Blance": "1800.000",
*            "Type": 3,
*            "Created": 1553594517
*        }
*    ],
*    "internalMsg": "获取成功",
*    "timeConsumed": 323451
*}
*
*/
func (cthis *UserController) SafeBoxInfo() {
	ctx := cthis.ctx
	if !utils.RequiredParam(&ctx, []string{"pageNum"}) {
		return
	}
	userid, _ := strconv.ParseInt(ctx.Values().GetString("userid"), 10, 64)
	pageNum, _ := strconv.Atoi(ctx.URLParam("pageNum"))
	var beans []xorm.AccountInfos
	err := models2.MyEngine[cthis.platform].Where("type = ?", config.FUNDSAFEBOX).And("user_id = ?", userid).Desc("created").Limit(8, (pageNum-1)*8).Find(&beans)
	if err != nil {
		utils.ResFaiJSON2(&ctx, err.Error(), "保险箱明细获取失败")
	}
	if len(beans) == 0 {
		checkNil(&ctx, nil)
	} else {
		checkNil(&ctx, beans)
	}

}

/**
 * @api {get} api/auth/v1/getAccount 获取用户账号信息
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人:wenzhen</span><br/><br/>
 * 获取用户账号信息<br>
 * 业务描述:获取用户账号信息(余额，保险箱余额等)</br>
 * @apiVersion 1.0.0
 * @apiName     getAccount
 * @apiGroup    user
 * @apiPermission iso,android客户端
 * @apiHeader (客户端请求头参数) {string} Authorization Bearer + 用户登录获得的token
 *
 * @apiError (请求失败返回) {int}      code            错误代码
 * @apiError (请求失败返回) {string}   clientMsg       提示信息
 * @apiError (请求失败返回) {string}   internalMsg     错误代码
 * @apiError (请求失败返回) {float}    timeConsumed   后台耗时
 *
 * @apiErrorExample {json} 失败返回
 *  {
 *      "code": 204,
 *      "clientMsg": "失败",
 *      "internalMsg": "",
 *      "timeConsumed": 1785
 *  }
 *
 * @apiSuccess (返回结果)  {int}      code            200
 * @apiSuccess (返回结果)  {string}   clientMsg       提示信息
 * @apiSuccess (返回结果)  {string}   internalMsg     提示信息
 * @apiSuccess (返回结果)  {json}  	  data            返回数据
 * @apiSuccess (返回结果)  {float}   timeConsumed    后台耗时
 * @apiSuccessExample {json} 响应结果
 *
*{
 *    "clientMsg": "操作成功",
 *    "code": 200,
 *    "data": {
 *        "Id": 2,
 *        "UserId": 28,
 *        "ChargedAmount": "0.000",
 *        "ConsumedAmount": "0.000",
 *        "Balance": "0.000",
 *        "WithdrawAmount": "0.000",
 *        "BalanceCharge": "0.000",
 *        "BalanceLucky": "0.000",
 *        "BalanceSafe": "100",
 *        "BalanceWallet": "11900",
 *        "Updated": 1553591581
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 388055
 *}
 *
*/
func (cthis *UserController) GetAccount() {
	ctx := cthis.ctx
	userid, _ := strconv.Atoi(ctx.Values().GetString("userid"))
	account := xorm.Accounts{UserId: userid}
	res, err := models2.MyEngine[cthis.platform].Where("user_id = ?", userid).Get(&account)
	if err != nil {
		fmt.Println(err.Error())
	}
	if res {
		utils.ResSuccJSON(&ctx, "", "获取成功", config.SUCCESSRES, account)
	} else {
		utils.ResFaiJSON(&ctx, "", "获取失败", config.NOTGETDATA)
	}
}

/**
 * @api {get} api/auth/v1/getUser 获取用户个人信息
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人:wenzhen</span><br/><br/>
 * 获取用户个人信息<br>
 * 业务描述:获取用户个人信息在个人中的个人信息栏目</br>
 * @apiVersion 1.0.0
 * @apiName     api_auth_v1_getUser
 * @apiGroup    user
  * @apiPermission iso,android客户端
 * @apiHeader (客户端请求头参数) {string} Authorization Bearer + 用户登录获得的token
 * @apiSuccess (返回结果)  {int}      code            200,204不能正常拿到数据
 * @apiSuccess (返回结果)  {string}   clientMsg       提示信息
 * @apiSuccess (返回结果)  {string}   internalMsg     提示信息
 * @apiSuccess (返回结果)  {json}  	  data            返回数据
 * @apiSuccess (返回结果)  {float}   timeConsumed    后台耗时
 * @apiSuccess (data字段说明)  {string}   birthday    生日日期
 * @apiSuccess (data字段说明)  {string}   last_platform_id    最后一次登录的游戏平台，不为0的时候需要再次调用退出游戏接口
 * @apiSuccess (data字段说明)  {string}   email    邮箱
 * @apiSuccess (data字段说明)  {string}   name    姓名
 * @apiSuccess (data字段说明)  {string}   phone    账号
 * @apiSuccess (data字段说明)  {string}   qq    qq
 * @apiSuccess (data字段说明)  {string}   sex    性别
 * @apiSuccess (data字段说明)  {float}   total_bet_amount    当前投注额
 * @apiSuccess (data字段说明)  {int}   valid_bet_max    升级到下一vip需要的投注额
 * @apiSuccess (data字段说明)  {int}   vip_level    当前vip等级
 * @apiSuccess (data字段说明)  {string}   wechat    微信号
 * @apiSuccessExample {json} 响应结果
 *
{
    "clientMsg": "获取成功",
    "code": 200,
    "data": {
        "birthday": "1970-01-01",
        "email": "",
        "last_platform_id": 1,
        "name": "13912345678",
        "phone": "13912345678",
        "user_name": "a13912345678",
        "qq": "",
        "sex": 1,
        "total_bet_amount": 0,
        "valid_bet_max": 100000,
        "vip_level": 1,
        "wechat": "",
		"vipList":[
            {
                "Id": 1,
                "Level": 1,
                "Name": "VIP1",
                "ValidBetMin": 0,
                "ValidBetMax": 1,
                "UpgradeAmount": 0,
                "WeeklyAmount": 0,
                "MonthAmount": 0,
                "UpgradeAmountTotal": 0,
                "HasDepositSpeed": 0,
                "HasOwnService": 0,
                "WashCode": "1.00"
            },
			...
			]
    },
    "internalMsg": "",
    "timeConsumed": 2992
}
*/
func (cthis *UserController) GetUser() {
	ctx := cthis.ctx
	userid := ctx.Values().GetString("userid")
	// @update by aTian 更新用户等级，防止客户端获取到升级投注额为负数
	iUserId, _ := strconv.Atoi(userid)
	services.UpgradeLevel(cthis.platform, iUserId)
	sqlString := "SELECT u.phone,u.user_name, u.name, u.sex, u.vip_level, u.birthday, u.email, u.qq, u.wechat,u.user_type, "
	sqlString += "truncate(a.total_bet_amount+a.today_bet_amount,3) total_bet_amount, u.last_platform_id, "
	sqlString += "truncate(v.valid_bet_max * 10000*1.0,3) valid_bet_max FROM users AS u LEFT JOIN accounts AS a ON u.id = a.user_id "
	sqlString += "LEFT JOIN vip_levels v ON u.vip_level = v.level WHERE u.id = " + userid
	response, err := utils.Query(cthis.platform, sqlString, "phone", "name", "total_bet_amount", "valid_bet_max")
	vips, _ := ramcache.TableVipLevels.Load(cthis.platform)
	response[0]["vipList"] = vips
	if err != nil {
		utils.ResFaiJSON(&ctx, "", "获取用户信息失败", config.NOTGETDATA)
		return
	}
	utils.ResSuccJSON(&ctx, "", "获取成功", config.SUCCESSRES, response[0])
}

/**
 * @api {post} api/auth/v1/updateUser 修改用户个人信息
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人:wenzhen</span><br/><br/>
 * 修改用户个人信息<br>
 * 业务描述:修改用户个人信息</br>
 * @apiVersion 1.0.0
 * @apiName     updateUser
 * @apiGroup    user
 * @apiPermission iso,android客户端
 * @apiHeader (客户端请求头参数) {string} Authorization Bearer + 用户登录获得的token
 * @apiParam (客户端请求参数) {string}    username      用户名
 * @apiParam (客户端请求参数) {string}    email         邮箱
 * @apiParam (客户端请求参数) {string}    birthday      生日 1970-01-01
 * @apiParam (客户端请求参数) {int}       sex           性别 1男2女
 * @apiParam (客户端请求参数) {string}    qq            QQ号
 * @apiParam (客户端请求参数) {string}    wechat        微信号
 * @apiParam (客户端请求参数) {string}    password      游客修改用户名时，password字段必填
 * @apiError (请求失败返回)   {int}      code           错误代码
 * @apiError (请求失败返回)   {string}   clientMsg      提示信息
 * @apiError (请求失败返回)   {string}   internalMsg    错误代码
 * @apiError (请求失败返回)   {float}    timeConsumed   后台耗时
 *
 * @apiErrorExample {json} 失败返回
 *  {
 *      "code": 204,
 *      "clientMsg": "失败",
 *      "internalMsg": "",
 *      "timeConsumed": 1785
 *  }
 *
 * @apiSuccess (返回结果)  {int}      code            200
 * @apiSuccess (返回结果)  {string}   clientMsg       提示信息
 * @apiSuccess (返回结果)  {string}   internalMsg     提示信息
 * @apiSuccess (返回结果)  {json}  	  data            返回数据
 * @apiSuccess (返回结果)  {float}   timeConsumed    后台耗时
 * @apiSuccessExample {json} 响应结果
 * {
 *     "clientMsg": "修改成功",
 *     "code": 200,
 *     "data": {
 *         "Id": 19,
 *         "Phone": "",
 *         "Password": "96e79218965eb72c92a549dd5a330112",
 *         "UserName": "stopXu",
 *         "Name": "",
 *         "Email": "billatin@gmail.com",
 *         "Created": 0,
 *         "Birthday": "1999-12-31",
 *         "MobileType": 0,
 *         "Sex": 1,
 *         "Path": "",
 *         "VipLevel": 0,
 *         "Qq": "100000",
 *         "Wechat": "17127290868",
 *         "Status": 0,
 *         "ProxyStatus": 0,
 *         "UserType": 0,
 *         "Token": "",
 *         "RegIp": "",
 *         "UniqueCode": "",
 *         "TokenCreated": 0,
 *         "SafePassword": "",
 *         "UserGroupId": "",
 *         "ParentId": 0,
 *         "LastLoginTime": 0,
 *         "LastPlatformId": 0,
 *         "GroupSize": 0,
 *         "WxOpenId": "",
 *         "IsModifiedUname": 0
 *     },
 *     "internalMsg": "",
 *     "timeConsumed": 14251294
 * }
 *
 */
func (cthis *UserController) UpdateUser() {
	ctx := cthis.ctx
	userid, _ := ctx.Values().GetInt("userid")
	sOldUsername := ctx.Values().GetString("username")
	user := xorm.Users{Id: userid}
	// @updated by aTian 修改用户名
	formVals := ctx.FormValues()
	lUsername, bUsernameExist := formVals["username"]
	sUsername := ""
	if bUsernameExist {
		sUsername = lUsername[0]
	}
	sUsername = strings.Trim(sUsername, " \t")
	//name := ctx.FormValue("name")
	// , Email: email, Birthday: birthday, Sex: sex, Qq: qq, Wechat: wechat
	lEmail, bEmailExist := formVals["email"]
	sEmail := ""
	updateFields := make(map[string]string, 0)
	if bEmailExist {
		sEmail = lEmail[0]
		user.Email = sEmail
		updateFields["email"] = sEmail
	}

	lBirthday, bBirthdayExist := formVals["birthday"]
	sBirthday := ""
	if bBirthdayExist {
		sBirthday = lBirthday[0]
		user.Birthday = sBirthday
		updateFields["birthday"] = sBirthday
	}

	lSex, bSexExist := formVals["sex"]
	sSex := ""
	if bSexExist {
		sSex = lSex[0]
		user.Sex, _ = strconv.Atoi(sSex)
		updateFields["sex"] = sSex
	}

	lQq, bQqExist := formVals["qq"]
	sQq := ""
	if bQqExist {
		sQq = lQq[0]
		user.Qq = sQq
		updateFields["qq"] = sQq
	}

	lWeChat, bWeChat := formVals["wechat"]
	sWeChat := ""
	if bWeChat {
		sWeChat = lWeChat[0]
		user.Wechat = sWeChat
		updateFields["wechat"] = sWeChat
	}
	var err error

	engine := models2.MyEngine[cthis.platform]
	ut, _ := ramcache.UserNameAndToken.Load(cthis.platform)
	utMap := ut.(map[string][]string)
	if sUsername != "" {
		match, _ := regexp.MatchString("(^[a-zA-Z][a-zA-Z0-9_]{4,15}$)", sUsername)
		if !match {
			utils.ResFaiJSON(&ctx, "", "用户名必须以字母开头，只能包含数字，字母，下划线，不允许输入特殊符号，长度5-16位", config.NOTGETDATA)
			return
		}
		var dbUser xorm.Users
		var dbUserExist bool
		dbUserExist, err = engine.Id(userid).Cols("user_type", "is_modified_uname").Get(&dbUser)
		if err != nil || dbUserExist == false {
			utils.ResFaiJSON(&ctx, "1906041715", "用户信息不存在", config.NOTGETDATA)
			return
		}
		if dbUser.UserType != 1 {
			utils.ResFaiJSON(&ctx, "1907031817", "只允许游客修改用户名", config.NOTGETDATA)
			return
		}
		cacheUser, cacheUserExist := utMap[sUsername]
		if cacheUserExist && cacheUser[0] != strconv.Itoa(userid) {
			utils.ResFaiJSON(&ctx, "1906041750", "用户名已存在", config.NOTGETDATA)
			return
		}
		sPassword := ctx.FormValue("password")
		iPwdLen := len(sPassword)
		if iPwdLen == 0 {
			utils.ResFaiJSON(&ctx, "1907041005", "游客修改用户名时，将成为正式用户，请填写密码", config.NOTGETDATA)
			return
		} else if iPwdLen < 6 {
			utils.ResFaiJSON(&ctx, "1907041009", "密码长度不可小于6位", config.NOTGETDATA)
			return
		}
		sEnPwd := utils.MD5(sPassword)
		user.UserName = sUsername
		user.Password = sEnPwd
		user.UserType = 0
		_, err = engine.Id(userid).Cols("user_name", "password", "user_type").Update(user)
	} else {
		if len(updateFields) > 0 {
			tmpArr := make([]string, 0)
			for idx, val := range updateFields {
				tmpArr = append(tmpArr, "`"+idx+"`='"+html.EscapeString(val)+"'")
			}
			sql := "UPDATE `users` SET " + strings.Join(tmpArr, ",") + " WHERE id=" + strconv.Itoa(userid)
			_, err = engine.Exec(sql)
		}
		//_, err = engine.Id(userid).Cols("email", "birthday", "sex", "qq", "wechat").Update(user)
	}
	if err == nil {
		if sUsername != "" && sUsername != sOldUsername {
			// 同步到UserIdCard，防止修改后使用用户名登录失败
			utils.UpdateUserIdCard(cthis.platform, userid, map[string]interface{}{
				"Username": sUsername,
			})
			utMapOld := utMap[sOldUsername]
			utMap[sUsername] = utMapOld
			delete(utMap, sOldUsername)
		}
		utils.ResSuccJSON(&ctx, "", "修改成功", config.SUCCESSRES, user)
	} else {
		utils.ResFaiJSON(&ctx, "", "修改失败", config.NOTGETDATA)
	}
}

/**
 * @api {get} api/auth/v1/getUserReport 获取用户个人报表信息
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人:wenzhen</span><br/><br/>
 * 获取用户个人报表信息<br>
 * 业务描述:获取用户个人报表信息</br>
 * @apiVersion 1.0.0
 * @apiName     getUserReport
 * @apiGroup    user
  * @apiPermission iso,android客户端
 * @apiHeader (客户端请求头参数) {string} Authorization Bearer + 用户登录获得的token
* @apiParam (客户端请求参数) {int}    	betTime    	   查询时间  0全部1今天2昨天3一个月内
 * @apiSuccess (返回结果)  {int}      code            200,204不能正常拿到数据
 * @apiSuccess (返回结果)  {string}   clientMsg       提示信息
 * @apiSuccess (返回结果)  {string}   internalMsg     提示信息
 * @apiSuccess (返回结果)  {json}  	  data            返回数据
 * @apiSuccess (返回结果)  {float}   timeConsumed    后台耗时
 * @apiSuccess (返回结果)  {array}   data    		   数据列表
 * @apiSuccess (data字段说明)  {float}   amount       总投注总额
 * @apiSuccess (data字段说明)  {float}   parent_id    游戏类型
 * @apiSuccess (data字段说明)  {float}   reward       总派彩金额
 * @apiSuccess (data字段说明)  {float}   washAmount   返利金额
 * @apiSuccess (data字段说明)  {float}   winAmount    输赢金额
 * @apiSuccessExample {json} 响应结果
 *
{
    "clientMsg": "获取成功",
    "code": 200,
    "data": [
        {
            "amount": 3555,
            "parent_id": 2,
            "reward": 2916.5,
            "washAmount": 16.84,
            "winAmount": -638.5
        },
        {
            "amount": 5.8,
            "parent_id": 3,
            "reward": 1,
            "washAmount": 0.03,
            "winAmount": -4.8
        }
    ],
    "internalMsg": "获取成功",
    "timeConsumed": 14959
}
*/
func (cthis *UserController) GetUserReport() {
	ctx := cthis.ctx
	if !utils.RequiredParam(&ctx, []string{"betTime"}) {
		return
	}
	betTime, err := ctx.URLParamInt("betTime")
	if err != nil {
		utils.ResFaiJSON(&ctx, "", "时间参数错误", config.PARAMERROR)
		return
	}
	userid, _ := ctx.Values().GetInt("userid")
	started, ended := utils.GetQueryTime(betTime)
	sqlString := "SELECT (a.reward - a.amount) winAmount, a.parent_id, a.amount, a.reward, TRUNCATE(IFNULL( (SELECT SUM(wi.amount) FROM wash_code_infos wi LEFT JOIN wash_code_records w ON w.`id` = wi.`record_id` WHERE wi.type_id = a.parent_id AND w.`user_id` = a.user_id AND w.`washtime` >= " + started + " AND w.`washtime`<= " + ended + "), 0 ),3) washAmount FROM (SELECT b.`user_id`, g.`parent_id`, SUM(b.`amount`) amount, SUM(b.`reward`) reward FROM bets_" + strconv.Itoa(userid%10) + " b LEFT JOIN platform_games p ON b.`game_code` = p.service_code AND b.`platform_id` = p.plat_id LEFT JOIN game_categories g ON p.`game_plat_id` = g.`id` WHERE b.`user_id` = " + strconv.Itoa(userid) + " AND b.`ented` >= " + started + " AND b.`ented` <= " + ended + " GROUP BY g.parent_id, b.`user_id`) a GROUP BY a.user_id, a.parent_id, a.amount, a.reward"
	response, err := utils.QueryString(cthis.platform, sqlString)
	if err != nil {
		utils.ResFaiJSON(&ctx, "", "获取用户报表信息失败", config.NOTGETDATA)
		return
	}
	if len(response) == 0 {
		checkNil(&ctx, nil)
	} else {
		checkNil(&ctx, response)
	}
}
func (cthis *UserController) ToDownload() {
	ctx := cthis.ctx
	ctx.View("download/" + cthis.platform + "/appDownload.html")
}

func (cthis *UserController) ToDownloadAPP() {
	ctx := cthis.ctx
	platform := ""
	parentid := ctx.URLParam("parentid")
	host := ctx.Host()
	ramcache.TableConfigs.Range(func(key, value interface{}) bool {
		fmt.Println(key)
		domain := value.(map[string]interface{})["tuiguang_web_url"]
		if domain == "http://"+host+"/download" || domain == "https://"+host+"/download" {
			platform = key.(string)
			return false
		}
		return true
	})
	if platform == "" {
		ctx.WriteString("无权访问!")
		return
	}

	//获取key之后，根据平台编号获取下载链接
	config, _ := ramcache.TableConfigs.Load(platform)
	cMap := config.(map[string]interface{})
	domains := strings.Split(cMap["tuiguang_web_domain"].(string), ",")
	var domain string
	if len(domains) == 1 {
		domain = domains[0]
	} else {
		domain = domains[rand.Intn(len(domains)-1)]
	}
	ctx.ViewData("domainUrl", domain+"?parentid="+parentid)
	ctx.View("download/fxApp.html")
}

func (cthis *UserController) ToHomeDownload() {
	ctx := cthis.ctx
	platform := ""
	host := ctx.Host()
	ramcache.TableConfigs.Range(func(key, value interface{}) bool {
		fmt.Println(key)
		domainsI := value.(map[string]interface{})["tuiguang_web_domain"]
		domains := strings.Split(domainsI.(string), ",")
		for _, domain := range domains {
			if domain == "http://"+host || domain == "https://"+host {
				platform = key.(string)
				return false
			}
		}
		return true
	})
	if platform == "" {
		ctx.WriteString("无权访问!")
		return
	}
	ctx.ViewData("platform", platform)
	ctx.View("download/" + platform + "/appDownload.html")
}

/**
 * @api {post} api/auth/v1/bindPhone 绑定手机号
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: aTian</span><br/><br/>
 * 绑定手机号<br>
 * 业务描述:绑定手机号</br>
 * @apiVersion 1.0.0
 * @apiName     BindPhone
 * @apiGroup    user
 * @apiPermission iso,android客户端
 * @apiParam (客户端请求参数) {string} phone  手机号
 * @apiParam (客户端请求参数) {string} vcode  短信验证码
 * @apiError (请求失败返回) {int}      code            错误代码
 * @apiError (请求失败返回) {string}   clientMsg       提示信息
 * @apiError (请求失败返回) {string}   internalMsg     错误代码
 * @apiError (请求失败返回) {float}    timeConsumed    后台耗时
 *
 * @apiErrorExample {json} 失败返回
 * {
 *     "clientMsg": "无效的验证码",
 *     "code": 204,
 *     "internalMsg": "1905271956",
 *     "timeConsumed": 0
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
 *     "clientMsg": "手机号绑定成功，请重新登录",
 *     "code": 200,
 *     "data": "",
 *     "internalMsg": "",
 *     "timeConsumed": 9631677
 * }
 */
func (cthis *UserController) BindPhone() {
	ctx := cthis.ctx
	if !utils.RequiredParamPost(&ctx, []string{"phone"}) {
		return
	}
	sCode := ctx.FormValue("vcode")
	if sCode == "" {
		utils.ResFaiJSON(&ctx, "1907080957", "请输入验证码", config.NOTGETDATA)
		return
	}
	sOldPhone := ctx.Values().GetString("phone")
	if sOldPhone != "" {
		utils.ResFaiJSON(&ctx, "1906011359", "手机号已绑定，要修改请联系客服", config.NOTGETDATA)
		return
	}
	sNewPhoneNum := ctx.FormValue("phone")
	platform := cthis.platform
	phoneNumAndToken, _ := ramcache.PhoneNumAndToken.Load(platform)
	phoneNumAndTokenMap := phoneNumAndToken.(map[string][]string)
	if _, phoneNumOk := phoneNumAndTokenMap[sNewPhoneNum]; phoneNumOk {
		utils.ResFaiJSON(&ctx, "", "手机号码已存在", config.NOTGETDATA)
		return
	}
	iCode, codeToIntErr := strconv.Atoi(sCode)
	if codeToIntErr != nil {
		utils.ResFaiJSON2(&ctx, codeToIntErr.Error(), "无效的验证码")
		return
	}
	phoneCheckCodeMap, loadOk := ramcache.PhoneCheckCode.Load(platform)
	if loadOk == false {
		utils.ResFaiJSON2(&ctx, "1905271956", "无效的验证码")
		return
	}
	var phoneInfoArr [2]int
	var phoneExist bool
	if phoneInfoArr, phoneExist = phoneCheckCodeMap.(map[string][2]int)[sNewPhoneNum]; !phoneExist {
		utils.ResFaiJSON2(&ctx, "1905271957", "无效的验证码")
		return
	}
	iCachedCode := phoneInfoArr[0]
	iNowTimestamp := utils.GetNowTime()
	iExpiredTimestamp := phoneInfoArr[1]
	if iNowTimestamp > iExpiredTimestamp {
		utils.ResFaiJSON2(&ctx, "1905271958", "验证码已过期")
		return
	}
	if iCachedCode != iCode {
		utils.ResFaiJSON2(&ctx, "1905271959", "无效的验证码")
		return
	}
	engine := models2.MyEngine[cthis.platform]
	session := engine.NewSession()
	session.Begin()
	defer session.Close()
	sUserId := ctx.Values().GetString("userid")
	iUserId, _ := strconv.Atoi(sUserId)
	updateUserBean := xorm.Users{
		Phone: sNewPhoneNum,
	}
	affectedNum, err := session.ID(sUserId).Update(&updateUserBean)
	if err != nil {
		session.Rollback()
		utils.ResFaiJSON2(&ctx, err.Error(), "手机号绑定失败")
		return
	}
	if affectedNum > 0 {
		err = services.BindPhoneAward(platform, session, iUserId)
		if err != nil {
			utils.ResFaiJSON(&ctx, "1906261646", err.Error(), config.NOTGETDATA)
			return
		}
		err = session.Commit()
		if err != nil {
			utils.ResFaiJSON2(&ctx, err.Error(), "手机号绑定失败")
			return
		}
		phoneNumAndTokenMap[sNewPhoneNum] = []string{sUserId, "", "", "1"}
		utils.UpdateUserIdCard(platform, iUserId, map[string]interface{}{
			"Phone": sNewPhoneNum,
		})
		delete(phoneCheckCodeMap.(map[string][2]int), sNewPhoneNum)
		delete(phoneNumAndTokenMap, sOldPhone)
		utils.ResSuccJSON(&ctx, "", "手机号绑定成功", config.SUCCESSRES, "")
		return
	}
	utils.ResFaiJSON2(&ctx, "1905272000", "手机号绑定失败")
	return
}

/**
 * @api {post} api/auth/v1/checkIn 签到
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: aTian</span><br/><br/>
 * 签到<br>
 * 业务描述:会员每日签到</br>
 * @apiVersion 1.0.0
 * @apiName     CheckIn
 * @apiGroup    user
 * @apiPermission iso,android客户端
 * @apiError (请求失败返回) {int}      code            错误代码
 * @apiError (请求失败返回) {string}   clientMsg       提示信息
 * @apiError (请求失败返回) {string}   internalMsg     错误代码
 * @apiError (请求失败返回) {float}    timeConsumed    后台耗时
 *
 * @apiErrorExample {json} 失败返回
 * {
 *     "clientMsg": "您已签到，记得明天再来哟！",
 *     "code": 204,
 *     "internalMsg": "1906261611",
 *     "timeConsumed": 0
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
 *     "clientMsg": "签到成功，记得明天再来哟！",
 *     "code": 200,
 *     "data": "",
 *     "internalMsg": "",
 *     "timeConsumed": 9631677
 * }
 */
func (cthis *UserController) CheckIn() {
	ctx := cthis.ctx
	platform := cthis.platform
	cnf, _ := ramcache.TableConfigs.Load(platform)
	cnfMap := cnf.(map[string]interface{})
	signAwardSwitch, sasOk := cnfMap["sign_award_switch"]
	signAward, saOk := cnfMap["sign_reward"]
	if sasOk && saOk {
		fSignAwardSwitch := signAwardSwitch.(float64)
		fSignAward := signAward.(float64)
		bSignAwardSwitch := int(fSignAwardSwitch) == 1
		if bSignAwardSwitch {
			iUserId, _ := ctx.Values().GetInt("userid")
			iSts, iEts := utils.GetDatetimeRange(0, 1)
			engine := models2.MyEngine[platform]
			isExist, err := engine.Where("user_id=? AND type=? AND created>=? AND created<?", iUserId, config.FUNDSIGNIN, iSts, iEts).Exist(new(xorm.AccountInfos))
			if err != nil {
				utils.ResFaiJSON(&ctx, err.Error(), "签到失败，请重试！", config.NOTGETDATA)
				return
			}
			if isExist {
				utils.ResFaiJSON(&ctx, "1906261825", "您已签到，记得明天再来哟！", config.NOTGETDATA)
				return
			}
			sUserId := strconv.Itoa(iUserId)
			info := map[string]interface{}{
				"user_id":  iUserId,
				"type_id":  config.FUNDSIGNIN,
				"amount":   fSignAward,
				"order_id": utils.CreationOrder("QD", sUserId),
				"msg":      "签到奖励",
			}
			balanceUpdateRes := fund.NewUserFundChange(platform).BalanceUpdate(info, nil)
			if balanceUpdateRes["status"] != 1 {
				utils.ResFaiJSON(&ctx, "1906261815", "签到失败，请重试！", config.NOTGETDATA)
				return
			} else {
				utils.ResSuccJSON(&ctx, "", "签到成功，记得明天再来哟！", config.SUCCESSRES, "")
				return
			}
		}
	}
	utils.ResFaiJSON(&ctx, "", "签到奖励已关闭，请坚持关注！", config.NOTGETDATA)
	return
}
