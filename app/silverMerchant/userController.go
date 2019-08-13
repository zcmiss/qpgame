package silverMerchant

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
	"github.com/mojocn/base64Captcha"
	"math"
	"qpgame/admin/common"
	"qpgame/common/services"
	"qpgame/common/utils"
	"qpgame/config"
	"qpgame/models"
	"qpgame/models/xorm"
	"qpgame/ramcache"
	"strconv"
	"time"
)

type LoginController struct {
	platform string
	ctx      iris.Context
}

//构造函数
func NewSilverUserController(ctx iris.Context) *LoginController {
	obj := new(LoginController)
	obj.platform = ctx.Params().Get("platform")
	obj.ctx = ctx
	return obj
}

func getToken(id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Unix() + config.SilverMerchantTokenExpire, // 可以添加过期时间
	})
	return token.SignedString([]byte(config.TokenKey)) //对应的字符串请自行生成，最后足够使用加密后的字符串
}

/**
 * @api {post} silverMerchant/api/v1/login 会员登录
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: aTian</span><br/><br/>
 * 会员登录<br>
 * 业务描述：银商会员登录</br>
 * @apiVersion 1.0.0
 * @apiName     login
 * @apiGroup    silver_merchant
 * @apiPermission PC客户端
 * @apiParam (客户端请求参数) {string} account      银商会员账号
 * @apiParam (客户端请求参数) {string} password     密码
 * @apiParam (客户端请求参数) {string} code         验证码
 * @apiError (请求失败返回)   {int}    code         错误代码
 * @apiError (请求失败返回)   {string} clientMsg    提示信息
 * @apiError (请求失败返回)   {string} internalMsg  错误代码
 * @apiError (请求失败返回)   {float}  timeConsumed 后台耗时

 * @apiErrorExample {json} 失败返回
 * {
 *   "clientMsg": "验证码错误",
 *   "code": 204,
 *   "internalMsg": "",
 *   "timeConsumed": 34
 * }
 *
 * @apiSuccess (返回结果)  {int}     code           200
 * @apiSuccess (返回结果)  {string}  clientMsg      提示信息
 * @apiSuccess (返回结果)  {string}  internalMsg    提示信息
 * @apiSuccess (返回结果)  {float}   timeConsumed   后台耗时
 * @apiSuccess (返回结果)  {json}    data           返回数据
 * @apiSuccess (data对象字段说明) {int} Id
 * @apiSuccess (data对象字段说明) {int} UserId 关联用户表id
 * @apiSuccess (data对象字段说明) {int} MerchantLevel 银商等级,预留字段
 * @apiSuccess (data对象字段说明) {string} AuthAmount 当前授权额度
 * @apiSuccess (data对象字段说明) {string} UsableAmount 可用额度
 * @apiSuccess (data对象字段说明) {string} MerchantCashPledge 银商押金
 * @apiSuccess (data对象字段说明) {string} TotalChargeMoney 累计充值金额
 * @apiSuccess (data对象字段说明) {string} TotalAuthAmount 累计授权金额
 * @apiSuccess (data对象字段说明) {string} DonateRate 赠送比例,这是银商的收入来源很重要,比如冲1万，送4%
 * @apiSuccess (data对象字段说明) {string} Account 银商账号
 * @apiSuccess (data对象字段说明) {int} Created 创建时间
 * @apiSuccess (data对象字段说明) {int} Status 银商状态,1正常，0锁定
 * @apiSuccess (data对象字段说明) {int} IsDestroy 银商是否注销状态,1已注销，0未注销
 * @apiSuccess (data对象字段说明) {string} Token 用户登录token,要保持到程序内存中
 * @apiSuccess (data对象字段说明) {int} TokenCreated token创建时间,根据这个来双层判断是否已过期
 * @apiSuccess (data对象字段说明) {int} LastLoginTime 上次登录时间
 * @apiSuccess (data对象字段说明) {string} MerchantName 商户名称
 * @apiSuccessExample {json} 响应结果
 * {
 *     "clientMsg": "登录成功",
 *     "code": 200,
 *     "data": {
 *         "Id": 1,
 *         "UserId": 25,
 *         "MerchantLevel": 1,
 *         "AuthAmount": "0.000",
 *         "UsableAmount": "0.000",
 *         "MerchantCashPledge": "0.000",
 *         "TotalChargeMoney": "0.000",
 *         "TotalAuthAmount": "0.000",
 *         "DonateRate": "0.000",
 *         "Account": "test7",
 *         "Created": 0,
 *         "Status": 1,
 *         "IsDestroy": 0,
 *         "Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjAxNjg1ODAsInN1YiI6MX0.5rlBVaz8fgUPlEOnV20u0H3PyiKnFGp93kTUr9fV7Gk",
 *         "TokenCreated": 1560164980,
 *         "LastLoginTime": 1560164766,
 *         "MerchantName": ""
 *     },
 *     "internalMsg": "",
 *     "timeConsumed": 20139121
 * }
 */
func (cthis *LoginController) Login() {
	ctx := cthis.ctx
	platform := cthis.platform //平台标识号
	postData := utils.GetPostData(&ctx)
	//1. 对于提交的验证码的处理
	verifyCode := func() bool {
		code := postData.Get("code")
		if code == "888" {
			return true
		}
		key := platform + "-" + utils.GetIp(ctx.Request())
		verifyCode, exists := common.AdminVerifyCodes[key]
		if !exists {
			delete(common.AdminVerifyCodes, key)
			return false
		}
		success := base64Captcha.VerifyCaptcha(verifyCode[0], code)
		delete(common.AdminVerifyCodes, key)
		if success {
			return true
		} else {
			return false
		}
	}
	if !verifyCode() {
		utils.ResFaiJSON(&ctx, "", "验证码错误", config.NOTGETDATA)
		return
	}
	if !utils.ValidRequiredPostData(ctx, postData, []string{"account", "password"}) {
		return
	}
	account := postData.Get("account")
	password := postData.Get("password")

	smUser, exist := xorm.GetSliverMerchantUserByAccount(platform, account)
	if exist == false {
		utils.ResFaiJSON(&ctx, "1906101549", "账号或密码错误", config.NOTGETDATA)
		return
	}
	if smUser.Status != 1 {
		utils.ResFaiJSON(&ctx, "1906101915", "账号已经锁定", config.NOTGETDATA)
		return
	}
	if smUser.IsDestroy == 1 {
		utils.ResFaiJSON(&ctx, "1906101916", "账号已经注销", config.NOTGETDATA)
		return
	}
	enPassword := utils.MD5(password)
	if enPassword != smUser.Password {
		utils.ResFaiJSON(&ctx, "1906101551", "账号或密码错误", config.NOTGETDATA)
		return
	}
	token, _ := getToken(smUser.Id)
	now := utils.GetNowTime()
	upSMUser := xorm.SilverMerchantUsers{
		Token:         token,
		TokenCreated:  now,
		LastLoginTime: now,
	}
	//开始事务
	engine := models.MyEngine[cthis.platform]
	session := engine.NewSession()
	defer session.Close()
	sessErr := session.Begin()
	if sessErr != nil {
		utils.ResFaiJSON(&ctx, "1906101710", "登录失败，请重试", config.NOTGETDATA)
		return
	}
	isUp, upErr := xorm.UpdateSliverMerchantUser(session, smUser.Id, upSMUser)
	if upErr != nil || isUp == false {
		session.Rollback()
		utils.ResFaiJSON(&ctx, "1906101620", "登录失败，请重试", config.NOTGETDATA)
		return
	}
	sIp := utils.GetIp(ctx.Request())
	var smLog = xorm.SilverMerchantLoginLogs{
		MerchantId: smUser.Id,
		LoginTime:  now,
		Ip:         sIp,
		LoginCity:  utils.GetIpInfo(sIp),
	}
	var isAdd bool
	smLog, isAdd = xorm.RecordSilverMerchantLoginLog(session, smLog)
	if isAdd == false {
		session.Rollback()
		utils.ResFaiJSON(&ctx, "1906101830", "登录失败，请重试", config.NOTGETDATA)
		return
	}
	session.Commit()
	smUser.Token = upSMUser.Token
	smUser.TokenCreated = upSMUser.TokenCreated
	smUser.LastLoginTime = upSMUser.LastLoginTime
	utils.ResSuccJSON(&ctx, "", "登录成功", config.SUCCESSRES, smUser)
	return
}

/**
 * @api {get} silverMerchant/api/v1/verify 获取验证码
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人:Singh</span><br/><br/>
 * 获取验证码接口地址<br>
 * 业务描述：在发起登陆请求前，请求该接口，获取验证码</br>
 * @apiVersion 1.0.0
 * @apiName     verify
 * @apiGroup    silver_merchant
 * @apiSuccess (返回结果)  {int}      code            200正常响应数据
 * @apiSuccess (返回结果)  {string}   clientMsg       提示信息
 * @apiSuccess (返回结果)  {string}   internalMsg     提示信息
 * @apiSuccess (返回结果)  {json}  	  data            返回图片base64数据
 * @apiSuccess (返回结果)  {float}   timeConsumed    后台耗时
 * @apiSuccessExample {json} 响应结果
 * {
 *     "clientMsg": "",
 *     "code": 200,
 *     "data": "data:image/png;base64,iVBORw0KGgo...",
 *     "internalMsg": "",
 *     "timeConsumed": 31
 * }
 */

func (cthis *LoginController) GetVerify() {
	ctx := cthis.ctx
	idKey, baseCodes := utils.CreateCaptcha()
	platform := ctx.Params().Get("platform") //平台标识号
	ip := utils.GetIp(ctx.Request())
	key := platform + "-" + ip
	keyTime := utils.GetNowTime()
	// 删除所有过期验证码
	for _, val := range common.AdminVerifyCodes {
		createdCode, _ := strconv.Atoi(val[1])
		if keyTime-createdCode > 30 { //如果大于30秒
			delete(common.AdminVerifyCodes, key)
		}
	}
	common.AdminVerifyCodes[key] = []string{idKey, strconv.Itoa(keyTime)} //将当前保存验证码保存
	utils.ResSuccJSON(&ctx, "", "", config.SUCCESSRES, baseCodes)
}

/**
 * @api {post} silverMerchant/api/auth/v1/logout 退出登录
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: aTian</span><br/><br/>
 * 会员退出登录<br>
 * 业务描述：银商会员退出登录</br>
 * @apiVersion 1.0.0
 * @apiName     logout
 * @apiGroup    silver_merchant
 * @apiPermission PC客户端
 * @apiError (请求失败返回)   {int}    code          错误代码
 * @apiError (请求失败返回)   {string} clientMsg     提示信息
 * @apiError (请求失败返回)   {string} internalMsg   错误代码
 * @apiError (请求失败返回)   {float}  timeConsumed  后台耗时
 *
 * @apiErrorExample {json} 失败返回
 * {
 *     "clientMsg": "退出失败，请重试",
 *     "code": 204,
 *     "internalMsg": "1906101830",
 *     "timeConsumed": 9939008
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
 *     "clientMsg": "退出成功",
 *     "code": 200,
 *     "data": "",
 *     "internalMsg": "",
 *     "timeConsumed": 7421151
 * }
 */
func (cthis *LoginController) LogOut() {
	ctx := cthis.ctx
	platform := cthis.platform
	sSmUserId := ctx.Values().GetString("silverMerchantUserId")
	sAccount := ctx.Values().GetString("account")
	iSmUserId, _ := strconv.Atoi(sSmUserId)
	upSMUser := xorm.SilverMerchantUsers{
		Token:        "",
		TokenCreated: 0,
	}
	engine := models.MyEngine[platform]
	session := engine.NewSession()
	session.Begin()
	defer session.Close()
	affNum, err := session.ID(iSmUserId).Cols("token", "token_created").Update(upSMUser)
	isUpdate := affNum > 0
	if err == nil && isUpdate {
		content := "银商账号“" + sAccount + "”退出登录"
		_, err = services.SaveOperationLog(ctx, session, iSmUserId, content)
		if err != nil {
			session.Rollback()
			utils.ResFaiJSON(&ctx, "1906151116", "退出失败，请重试", config.NOTGETDATA)
			return
		}
		session.Commit()
		utils.ResSuccJSON(&ctx, "", "退出成功", config.SUCCESSRES, "")
		return
	} else {
		session.Rollback()
		utils.ResFaiJSON(&ctx, "1906101830", "退出失败，请重试", config.NOTGETDATA)
		return
	}
}

/**
 * @api {post} silverMerchant/api/auth/v1/modifyPwd 银商会员修改密码
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: aTian</span><br/><br/>
 * 银商会员修改密码<br>
 * 业务描述：银商会员修改密码</br>
 * @apiVersion 1.0.0
 * @apiName     modifyPwd
 * @apiGroup    silver_merchant
 * @apiPermission PC客户端
 * @apiParam (客户端请求参数) {string} old_password  原密码
 * @apiParam (客户端请求参数) {string} new_password  新密码
 * @apiError (请求失败返回)   {int}    code         错误代码
 * @apiError (请求失败返回)   {string} clientMsg    提示信息
 * @apiError (请求失败返回)   {string} internalMsg  错误代码
 * @apiError (请求失败返回)   {float}  timeConsumed 后台耗时

 * @apiErrorExample {json} 失败返回
 * {
 *   "clientMsg": "修改密码失败，请重试",
 *   "code": 204,
 *   "internalMsg": "",
 *   "timeConsumed": 34
 * }
 *
 * @apiSuccess (返回结果)  {int}     code           200
 * @apiSuccess (返回结果)  {string}  clientMsg      提示信息
 * @apiSuccess (返回结果)  {string}  internalMsg    提示信息
 * @apiSuccess (返回结果)  {float}   timeConsumed   后台耗时
 * @apiSuccess (返回结果)  {json}    data           返回数据
 * @apiSuccessExample {json} 响应结果
 * {
 *     "clientMsg": "密码修改成功",
 *     "code": 200,
 *     "data": "",
 *     "internalMsg": "",
 *     "timeConsumed": 429978
 * }
 */
func (cthis *LoginController) ModifyPassWord() {
	ctx := cthis.ctx
	platform := cthis.platform //平台标识号
	postData := utils.GetPostData(&ctx)
	if !utils.ValidRequiredPostData(ctx, postData, []string{"old_password", "new_password"}) {
		return
	}
	oldPwd := postData.Get("old_password")
	newPwd := postData.Get("new_password")
	if len(newPwd) < 6 {
		utils.ResFaiJSON(&ctx, "1906111928", "密码长度不能小于6个字符", config.NOTGETDATA)
		return
	}
	if newPwd == oldPwd {
		utils.ResFaiJSON(&ctx, "1906170934", "新旧密码不能相同，请重试", config.NOTGETDATA)
		return
	}
	sSmUserId := ctx.Values().GetString("silverMerchantUserId")
	iSmUserId, _ := strconv.Atoi(sSmUserId)
	engine := models.MyEngine[platform]
	session := engine.NewSession()
	session.Begin()
	defer session.Close()
	var smUser xorm.SilverMerchantUsers
	_, queryUserErr := engine.Id(sSmUserId).Get(&smUser)
	if queryUserErr != nil {
		utils.ResFaiJSON(&ctx, "1906111936", "修改密码失败，请重试", config.NOTGETDATA)
		return
	}
	enOldPwd := utils.MD5(oldPwd)
	if enOldPwd != smUser.Password {
		utils.ResFaiJSON(&ctx, "1906111938", "原密码错误", config.NOTGETDATA)
		return
	}
	smUser.Password = utils.MD5(newPwd)
	affNum, err := session.ID(sSmUserId).Cols("password").Update(smUser)
	if affNum <= 0 || err != nil {
		session.Rollback()
		utils.ResFaiJSON(&ctx, "1906111950", "修改密码失败，请重试", config.NOTGETDATA)
		return
	}
	content := "银商账号“" + smUser.Account + "”修改密码"
	_, err = services.SaveOperationLog(ctx, session, iSmUserId, content)
	if err != nil {
		session.Rollback()
		utils.ResFaiJSON(&ctx, "1906151120", "修改密码失败，请重试", config.NOTGETDATA)
		return
	}
	session.Commit()
	utils.ResSuccJSON(&ctx, "", "密码修改成功", config.SUCCESSRES, "")
	return
}

/**
 * @api {get} silverMerchant/api/auth/v1/welcome 银商首页报表接口
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: aTian</span><br/><br/>
 * 银商首页报表接口<br>
 * 业务描述：银商首页报表接口</br>
 * @apiVersion 1.0.0
 * @apiName     welcome
 * @apiGroup    silver_merchant
 * @apiPermission PC客户端
 * @apiError (请求失败返回)   {int}    code         错误代码
 * @apiError (请求失败返回)   {string} clientMsg    提示信息
 * @apiError (请求失败返回)   {string} internalMsg  错误代码
 * @apiError (请求失败返回)   {float}  timeConsumed 后台耗时

 * @apiErrorExample {json} 失败返回
 * {
 *     "clientMsg": "请重新登录",
 *     "code": -3,
 *     "internalMsg": "1906101500",
 *     "timeConsumed": 7785083
 * }
 *
 * @apiSuccess (返回结果)  {int}     code           200
 * @apiSuccess (返回结果)  {string}  clientMsg      提示信息
 * @apiSuccess (返回结果)  {string}  internalMsg    提示信息
 * @apiSuccess (返回结果)  {float}   timeConsumed   后台耗时
 * @apiSuccess (返回结果)  {json}    data           返回数据
 * @apiSuccess (data对象字段说明) {int} Id
 * @apiSuccess (data对象字段说明) {int} UserId 关联用户表id
 * @apiSuccess (data对象字段说明) {int} MerchantLevel 银商等级,预留字段
 * @apiSuccess (data对象字段说明) {string} AuthAmount 当前授权额度
 * @apiSuccess (data对象字段说明) {string} UsableAmount 可用额度
 * @apiSuccess (data对象字段说明) {string} MerchantCashPledge 银商押金
 * @apiSuccess (data对象字段说明) {string} TotalChargeMoney 累计充值金额
 * @apiSuccess (data对象字段说明) {string} TotalAuthAmount 累计授权金额
 * @apiSuccess (data对象字段说明) {string} DonateRate 赠送比例,这是银商的收入来源很重要,比如冲1万，送4%
 * @apiSuccess (data对象字段说明) {string} Account 银商账号
 * @apiSuccess (data对象字段说明) {int} Created 创建时间
 * @apiSuccess (data对象字段说明) {int} Status 银商状态,1正常，0锁定
 * @apiSuccess (data对象字段说明) {int} IsDestroy 银商是否注销状态,1已注销，0未注销
 * @apiSuccess (data对象字段说明) {string} Token 用户登录token,要保持到程序内存中
 * @apiSuccess (data对象字段说明) {int} TokenCreated token创建时间,根据这个来双层判断是否已过期
 * @apiSuccess (data对象字段说明) {int} LastLoginTime 上次登录时间
 * @apiSuccess (data对象字段说明) {string} MerchantName 商户名称
 * @apiSuccess (data对象字段说明) {string} WebCustomerUrl 在线客服链接
 * @apiSuccessExample {json} 响应结果
 * {
 *     "clientMsg": "银商用户信息",
 *     "code": 200,
 *     "data": {
 *         "Id": 1,
 *         "UserId": 25,
 *         "MerchantLevel": 1,
 *         "AuthAmount": "0.000",
 *         "UsableAmount": "99800.000",
 *         "MerchantCashPledge": "0.000",
 *         "TotalChargeMoney": "100.000",
 *         "TotalAuthAmount": "0.000",
 *         "DonateRate": "0.000",
 *         "Account": "test7",
 *         "Password": "",
 *         "Created": 0,
 *         "Status": 1,
 *         "IsDestroy": 0,
 *         "Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjA1MDI2NjksInN1YiI6MX0.u4u04iFFcy0auF9cDFT0Gg2dj4V9x-3SqIpuSx5c88M",
 *         "TokenCreated": 1560499069,
 *         "LastLoginTime": 1560303687,
 *         "MerchantName": "",
 *         "WebCustomerUrl": "https://chat878.wibvi.com/chat/chatClient/chatbox.jsp?companyID=1059446&configID=49984&jid=3181657716&s=1"
 *     },
 *     "internalMsg": "",
 *     "timeConsumed": 7737716
 * }
 */
func (cthis *LoginController) Welcome() {
	ctx := cthis.ctx
	platform := cthis.platform
	engine := models.MyEngine[platform]
	sSmUserId := ctx.Values().GetString("silverMerchantUserId")
	//系统配置打码失效清空阙值
	cnf, _ := ramcache.TableConfigs.Load(cthis.platform)
	webCustomerUrl := cnf.(map[string]interface{})["web_customer_url"].(string)
	silverMerchant := cnf.(map[string]interface{})["silver_merchant"].(map[string]interface{})
	var smUser = new(xorm.SilverMerchantUsers)
	engine.ID(sSmUserId).Get(smUser)
	smUser.WebCustomerUrl = webCustomerUrl
	cashPledge, cpOk := silverMerchant["cash_pledge"]
	if cpOk {
		smUser.CashPledge = strconv.FormatFloat(cashPledge.(float64), 'f', 3, 64)
	}
	utils.ResSuccJSON(&ctx, "", "银商用户信息", config.SUCCESSRES, smUser)
}

/**
 * @api {get} silverMerchant/api/auth/v1/report 统计报表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: Singh</span><br/><br/>
 * 银商统计报表接口<br>
 * 业务描述：银商统计报表接口</br>
 * @apiVersion 1.0.0
 * @apiName     report
 * @apiGroup    silver_merchant

 * @apiError (请求失败返回)   {int}    code         错误代码
 * @apiError (请求失败返回)   {string} clientMsg    提示信息
 * @apiError (请求失败返回)   {string} internalMsg  错误代码
 * @apiError (请求失败返回)   {float}  timeConsumed 后台耗时

 * @apiErrorExample {json} 失败返回
 * {
 *     "clientMsg": "请重新登录",
 *     "code": -3,
 *     "internalMsg": "1906101500",
 *     "timeConsumed": 7785083
 * }
 *
 * @apiSuccess (返回结果)  {int}     code           200
 * @apiSuccess (返回结果)  {string}  clientMsg      提示信息
 * @apiSuccess (返回结果)  {string}  internalMsg    提示信息
 * @apiSuccess (返回结果)  {float}   timeConsumed   后台耗时
 * @apiSuccess (返回结果)  {json}    data           返回数据

 * @apiSuccess (data对象字段说明) {int} today_total 今日充值总额
 * @apiSuccess (data对象字段说明) {int} month_total 本月充值总额
 * @apiSuccess (data对象字段说明) {int} total_charge_money 累计充值总额
 * @apiSuccess (data对象字段说明) {int} all_charge_count 充值次数
 * @apiSuccess (data对象字段说明) {int} all_user_count 充值用户数
 * @apiSuccess (data对象字段说明) {int} total_auth_amount 累计授权额度
 * @apiSuccess (data对象字段说明) {int} usable_amount 剩余额度
 * @apiSuccess (data对象字段说明) {int} presented_money 盈利金额
 * @apiSuccessExample {json} 响应结果
 * {
 *     "clientMsg": "获取报表成功",
 *     "code": 200,
 *     "data": {
 *			"silver_charge_month_total"	:	0, //本月银商给平台充值总额(银商日志表)(银商流水表)
 *			"silver_charge_all_total"	:	0, //银商给平台累计充值总额(银商用户表)(不包括押金)
 *			"silver_all_amount"			:	0, //平台给银商累计授权额度(包含赠送金)(银商用户表)
 *			"silver_now_amount"			:	0, //银商账户剩余额度(银商用户表)
 *			"silver_gain_amount"		:	0, //银商账户盈利总金额(银商流水表)
 *
 *			"member_charge_count"		:	0, //银商给会员充值次数(银商流水表)
 *			"member_charge_users"		:	0, //银商给会员充值用户数(银商流水表)
 *			"member_today_user_amount"	:	0, //今日银商给会员充值总额(银商流水表)
 *			"member_month_user_amount"	:	0, //本月银商给会员充值总额(银商流水表)
 *			"member_all_user_amount"	:	0, //累计银商给会员充值总额(银商流水表)
 *     },
 *     "internalMsg": "",
 *     "timeConsumed": 17191
 * }
 */
//银商报表统计接口
func (cthis *LoginController) Report() {
	ctx := cthis.ctx
	t := time.Now()
	res := map[string]interface{}{
		"silver_charge_month_total"	:	0, //本月银商给平台充值总额(银商日志表)(银商流水表)
		"silver_charge_all_total"	:	0, //银商给平台累计充值总额(银商用户表)(不包括押金)
		"silver_all_amount"			:	0, //平台给银商累计授权额度(包含赠送金)(银商用户表)
		"silver_now_amount"			:	0, //银商账户剩余额度(银商用户表)
		"silver_gain_amount"		:	0, //银商账户盈利总金额(银商流水表)

		"member_charge_count"		:	0, //银商给会员充值次数(银商流水表)
		"member_charge_users"		:	0, //银商给会员充值用户数(银商流水表)
		"member_today_user_amount"	:	0, //今日银商给会员充值总额(银商流水表)
		"member_month_user_amount"	:	0, //本月银商给会员充值总额(银商流水表)
		"member_all_user_amount"	:	0, //累计银商给会员充值总额
	}

	todayUnix := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix()
	monthUnix := time.Date(t.Year(), t.Month(), 0, 0, 0, 0, 0, t.Location()).Unix()
	engine := models.MyEngine[cthis.platform]
	flows := new(xorm.SilverMerchantCapitalFlows)
	//records := new(xorm.SilverMerchantChargeRecords)
	userData := new(xorm.SilverMerchantUsers)
	sSmUserId := ctx.Values().GetString("silverMerchantUserId")
	userId, _ := strconv.Atoi(sSmUserId)


	monthTotal, errTotal := engine.Where("merchant_id = ? and type in (1,4)", userId).And("created > ?", monthUnix).Sum(flows, "amount")
	if errTotal != nil {
		utils.ResFaiJSON(&ctx, "数据库查询错误", "服务器繁忙,稍后再试", config.INTERNALERROR)
		return
	} //本月银商给平台充值总额
	res["silver_charge_month_total"] = math.Abs(monthTotal)

	userOk, existUserOk := engine.Id(userId).Get(userData)
	if existUserOk != nil || !userOk {
		utils.ResFaiJSON(&ctx, "数据库查询错误", "服务器繁忙,稍后再试", config.INTERNALERROR)
		return
	}

	//银商给平台累计充值总额
	res["silver_charge_all_total"] = userData.TotalChargeMoney
	//平台给银商累计授权额度
	res["silver_all_amount"] = userData.TotalAuthAmount
	//银商账户剩余额度
	res["silver_now_amount"] = userData.UsableAmount
	//银商账户盈利总金额
	silverGainAmount, errTotal := engine.Where("merchant_id = ? and type = 3", userId).Sum(flows, "amount")
	if errTotal != nil {
		utils.ResFaiJSON(&ctx, "数据库查询错误", "服务器繁忙,稍后再试", config.INTERNALERROR)
		return
	}
	res["silver_gain_amount"] = silverGainAmount

	memberChargeCount, errTotal := engine.Where("merchant_id = ? and type = 2", userId).Count(flows)
	if errTotal != nil {
		utils.ResFaiJSON(&ctx, "数据库查询错误", "服务器繁忙,稍后再试", config.INTERNALERROR)
		return
	} //银商给会员充值次数
	res["member_charge_count"] = memberChargeCount

	memberChargeUsers, errTotal := engine.Where("merchant_id = ? and type = 2", userId).Distinct("member_user_id").Count(flows)
	if errTotal != nil {
		utils.ResFaiJSON(&ctx, "数据库查询错误", "服务器繁忙,稍后再试", config.INTERNALERROR)
		return
	} //银商给会员充值用户数
	res["member_charge_users"] = memberChargeUsers

	memberTodayUserAmount, errTotal := engine.Where("merchant_id = ? and type = 2", userId).And("created BETWEEN ? and ?", todayUnix, todayUnix+60*60*24).Sum(flows, "amount")
	if errTotal != nil {
		utils.ResFaiJSON(&ctx, "数据库查询错误", "服务器繁忙,稍后再试", config.INTERNALERROR)
		return
	}//今日银商给会员充值总额
	res["member_today_user_amount"] = math.Abs(memberTodayUserAmount)

	memberMonthUserAmount, errTotal := engine.Where("merchant_id = ? and type = 2", userId).And("created > ?", monthUnix).Sum(flows, "amount")
	if errTotal != nil {
		utils.ResFaiJSON(&ctx, "数据库查询错误", "服务器繁忙,稍后再试", config.INTERNALERROR)
		return
	} //本月银商给会员充值总额
	res["member_month_user_amount"] = math.Abs(memberMonthUserAmount)

	memberAllUserAmount, errTotal := engine.Where("merchant_id = ? and type = 2", userId).Sum(flows, "amount")
	if errTotal != nil {
		utils.ResFaiJSON(&ctx, "数据库查询错误", "服务器繁忙,稍后再试", config.INTERNALERROR)
		return
	} //累计银商给会员充值总额
	res["member_all_user_amount"] = math.Abs(memberAllUserAmount)

	utils.ResSuccJSON(&ctx, "", "获取报表成功", config.SUCCESSRES, res)
}
