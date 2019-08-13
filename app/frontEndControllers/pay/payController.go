package pay

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/go-xorm/builder"
	goxorm "github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	"github.com/shopspring/decimal"
	"html/template"
	"math/rand"
	"net/url"
	"qpgame/app/fund"
	"qpgame/common/services"
	"qpgame/common/utils"
	"qpgame/config"
	"qpgame/models"
	"qpgame/models/xorm"
	"qpgame/ramcache"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

type PayController struct {
	platform string
	ctx      iris.Context
}

//构造函数
func NewPayController(ctx iris.Context) *PayController {
	obj := new(PayController)
	obj.platform = ctx.Params().Get("platform")
	obj.ctx = ctx
	return obj
}

/**
 * @api {post} api/auth/v1/payAction 第三方支付订单生成接口
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人:Pling把妹</span><br/><br/>
 * 充值页面充值按钮<br>
 * 业务描述:生成支付订单</br>
 * @apiVersion 1.0.0
 * @apiName     api_auth_v1_payAction
 * @apiGroup    finance_module
 * @apiPermission ios,android客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 Bearer token
 * @apiParam (客户端请求参数) {decimal}    amount          充值金额
 * @apiParam (客户端请求参数) {string}     bank_address    支付编码
 * @apiParam (客户端请求参数) {int}        credential_id   支付证书ID
 * @apiParam (客户端请求参数) {int=4,5,6}  addr_type       支付类型:4.第三方二维码支付,5.第三方 wap支付,6.第三方h5支付 支付类型
  * @apiSuccess (返回结果)  {int}     code            200
 * @apiSuccess (返回结果)  {string}  clientMsg       提示信息
 * @apiSuccess (返回结果)  {string}  internalMsg     提示信息
 * @apiSuccess (返回结果)  {json}  	  data            返回数据支付信息
 * @apiSuccess (返回结果)  {float}   timeConsumed    后台耗时
 * @apiSuccess (data对象字段说明) {object}   list   支付信息
 * @apiSuccess (data对象字段说明) {string} qrcode  地址【addr_type=4时为二维码地址，addr_type=5或6时为网页地址】
 * @apiSuccess (data对象字段说明) {int} app_type   addr_type=4 时扫码软件【1 QQ ；2 微信 ；3 支付宝；4京东；5浏览器】
 * @apiSuccess (data对象字段说明) {int=0,1} isBankTransfer  是否银行卡转账默认0
 * @apiSuccessExample {json} 响应结果
 *  {
 *      "status": 1,
 *      "clientMsg": "",
 *      "internalMsg": "",
 *      "data":  {
 *           "qrcode": "http://desk.ling-pay.com/h5pay/PrePayServlet?payType=JD&orderNum=532018070713255527771392",
 *           "app_type": 2,
			 "isBankTransfer":0,
 *      },
 *      "timeConsumed": 18907
 *  }
*/
func (cthis *PayController) PayAction() {
	ctx := cthis.ctx
	userIdS := ctx.Values().GetString("userid")
	sPayAmount := ctx.PostValue("amount")
	decimal.DivisionPrecision = 3
	dPayAmount, errPayAmount := decimal.NewFromString(sPayAmount)
	if errPayAmount != nil {
		//fmt.Println(existPayAmount)
		utils.ResFaiJSON(&ctx, "不是数字类型", "请填写正确的充值金额", config.PARAMERROR)
		return
	}

	payBankCode := ctx.PostValue("bank_address")
	credentialId, existCredential := ctx.PostValueInt("credential_id")
	if existCredential != nil {
		utils.ResFaiJSON(&ctx, "credential_id必须为整数", "支付证书错误", config.PARAMERROR)
		return
	}
	addrType, existAddrType := ctx.PostValueInt("addr_type")
	if existAddrType != nil || !utils.InArrayInt(addrType, []int{4, 5, 6}) {
		utils.ResFaiJSON(&ctx, "addrType必须为4-6的值", "支付操作出现错误", config.PARAMERROR)
		return
	}
	payCredential, _ := ramcache.TablePayCredential.Load(cthis.platform)
	singleCred, existSingleCred := payCredential.(map[int]xorm.PayCredentials)[credentialId]
	if !existSingleCred {
		utils.ResFaiJSON(&ctx, "credential_id不存在", "支付证书找不到", config.PARAMERROR)
		return
	}
	platForm := singleCred.PlatForm
	status := singleCred.Status
	chargeAmountConf := singleCred.ChargeAmountConf

	if status == 0 {
		utils.ResFaiJSON(&ctx, "已锁定", "该支付证书已停用，请更换支付方式", config.CREDENTIALSTOP)
		return
	}
	financial := NewFinancial(cthis.platform)
	//生成订单
	orderId := utils.CreationOrder("THC", userIdS)
	//添加 检测用户从当前时间往前10分钟之内，若有超过10次未充值成功的
	checkRes := financial.CheckUserByAddChargeRecord(userIdS)
	if checkRes {
		utils.ResFaiJSON(&ctx, "", "短时间内重复充值次数过多,请稍后尝试!", config.CREDENTIALOSTOOMANY)
		return
	}
	//查询该支付证书对应的充值金额配置
	if chargeAmountConf > 0 {
		chargeAmountConfMap := map[int]string{
			1: "0.",  //充值金额配置.开启随机金额小数到角
			2: "0.0", //充值金额配置.开启随机金额小数到分
		}
		chargeAmountConfBefore, isExist := chargeAmountConfMap[chargeAmountConf]
		if isExist {
			rand.Seed(time.Now().UnixNano())
			randVal := rand.Intn(9-1) + 1 //获取到随机金额小数到角或分
			randMoney := chargeAmountConfBefore + strconv.Itoa(randVal)
			dRandMoney, _ := decimal.NewFromString(randMoney)
			dPayAmount = dPayAmount.Add(dRandMoney) //将原来的金额加上对应获取到的小数位,如100+0.5,则需要支付的是100.5
		}
	}
	fPayAmount, _ := dPayAmount.Float64()
	//支付处理参数
	payHandle := NewPay(cthis.platform, payParam{
		platForm:     platForm,
		payAmount:    fPayAmount,
		payBankcode:  payBankCode,
		credentialId: credentialId,
		userIdS:      userIdS,
		addrType:     addrType,
		orderId:      orderId,
	})
	payHandle.httpRequest = ctx.Request() //用于获取ip
	//分发支付处理
	res, resOk := payHandle.PayDispathChannal()
	if !resOk {
		utils.ResFaiJSON(&ctx, "支付错误", res["message"].(string), res["code"].(int16))
		return
	}
	chargeTypes := map[int]string{1: "QQ", 2: "微信", 3: "支付宝", 4: "京东", 5: "网银"}
	addrTypes := map[int]string{4: "扫码", 5: "wap", 6: "H5"}
	ChargeTypeInfo := chargeTypes[res["app_type"].(int)] + addrTypes[addrType] + "支付"
	chargeRecord := new(xorm.ChargeRecords)
	userId, _ := strconv.Atoi(userIdS)
	chargeRecord.UserId = userId
	chargeRecord.OrderId = orderId
	chargeRecord.Amount = dPayAmount.String()
	chargeRecord.Created = int(time.Now().Unix())
	chargeRecord.CredentialId = credentialId
	chargeRecord.ChargeTypeId = res["charge_type"].(int)
	chargeRecord.PayBankCode = payBankCode
	chargeRecord.ChargeTypeInfo = ChargeTypeInfo
	chargeRecord.State = 0
	chargeRecord.IsTppay = 1
	engine := models.MyEngine[cthis.platform]
	affected, err := engine.Insert(chargeRecord)
	if err != nil {
		utils.ResFaiJSON(&ctx, "支付错误", "支付失败", config.PAYACTIONERROR)
		return
	}
	if affected == 0 {
		utils.ResFaiJSON(&ctx, "支付错误", "支付失败", config.PAYACTIONERROR)
		return
	}
	delete(res, "charge_type")
	utils.ResSuccJSON(&ctx, "", "", config.SUCCESSRES, res)
}

/**
 * @api {post} api/auth/v1/payCompanyAction 创建公司入款银行卡充值订单
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人:pling把妹</span><br/><br/>
 * 我充值页面的,银行卡充值,入款分俩种一种是第三方支付，一种是本平台直接充值，直接充值划为公司入款充值<br>
 * 业务描述:银行卡充值</br>
 * @apiVersion 1.0.0
 * @apiName     api_auth_v1_payCompanyAction
 * @apiGroup    finance_module
 * @apiPermission ios,android客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 Bearer token
 * @apiParam (客户端请求参数) {int} addr_type             充值类型(1,银行卡充值)
 * @apiParam (客户端请求参数) {float64} amount          	充值金额
 * @apiParam (客户端请求参数) {string} bank_name           收款银行
 * @apiParam (客户端请求参数) {string} real_name            收款人
 * @apiParam (客户端请求参数) {string} card_number         收款账号
 * @apiParam (客户端请求参数) {string} bank_address        收款银行开户地
 * @apiParam (客户端请求参数) {int} charge_cards_id     支付通道ID
 * @apiParam (客户端请求参数) {string} info     存款信息(请输入姓名或者卡号后四位)
  * @apiSuccess (返回结果)  {int}     code            200
 * @apiSuccess (返回结果)  {string}  clientMsg       提示信息
 * @apiSuccess (返回结果)  {string}  internalMsg     提示信息
 * @apiSuccess (返回结果)  {json}  	  data            返回数据支付信息
 * @apiSuccess (返回结果)  {float}   time_consumed    后台耗时

 * @apiSuccess (data对象字段说明) {string}   amount   金额说明
 * @apiSuccess (data对象字段说明) {string} orderNum   订单号
 * @apiSuccess (data对象字段说明) {int} payTime   支付时间
 * @apiSuccess (data对象字段说明) {string} payWay  支付方式
 * @apiSuccess (data对象字段说明) {string} status  支付状态

 * @apiSuccessExample {json} 响应结果
 {
    "clientMsg": "",
    "code": 200,
    "data": {
        "amount": "+200",
        "orderNum": "E20190423105848eab3dd",
        "payTime": 1555988328,
        "payWay": "",
        "status": "入款中..."
    },
    "internalMsg": "",
    "timeConsumed": 18495
}
*/

func (cthis *PayController) PayCompanyAction() {
	ctx := cthis.ctx
	userIdS := ctx.Values().GetString("userid")
	userId, _ := strconv.Atoi(userIdS)
	payAmount, existPayAmount := ctx.PostValueFloat64("amount")
	if existPayAmount != nil {
		utils.ResFaiJSON(&ctx, "不是数字类型", "请填写正确的充值金额", config.PARAMERROR)
		return
	}
	bankAddress := ctx.PostValue("bank_address")
	if bankAddress == "" {
		utils.ResFaiJSON(&ctx, "银行卡地址不能为空", "充值失败", config.PARAMERROR)
		return
	}
	bankName := ctx.PostValue("bank_name")
	realName := ctx.PostValue("real_name")
	info := ctx.PostValue("info")
	cardNumber := ctx.PostValue("card_number")
	chargeCardsId, _ := ctx.PostValueInt("charge_cards_id")
	addrType, existAddrType := ctx.PostValueInt("addr_type")
	if existAddrType != nil || addrType != 1 {
		utils.ResFaiJSON(&ctx, "addrType必须为1", "只支持银行卡充值", config.PARAMERROR)
		return
	}
	var chargeCards = new(xorm.ChargeCards)
	engine := models.MyEngine[cthis.platform]
	chCarOk, existChCarOk := engine.Id(chargeCardsId).Get(chargeCards)
	if existChCarOk != nil {
		utils.ResFaiJSON(&ctx, "sql错误", "充值失败", config.INTERNALERROR)
		return
	}
	if !chCarOk {
		utils.ResFaiJSON(&ctx, "", "充值账户不存在!", config.PARAMERROR)
		return
	}
	if chargeCards.Name != bankName {
		utils.ResFaiJSON(&ctx, "", "银行卡名称错误", config.PARAMERROR)
		return
	}
	if chargeCards.Owner != realName {
		utils.ResFaiJSON(&ctx, "", "收款人名称错误", config.PARAMERROR)
		return
	}
	if chargeCards.CardNumber != cardNumber {
		utils.ResFaiJSON(&ctx, "", "收款账号错误", config.PARAMERROR)
		return
	}
	var users = new(xorm.Users)
	engine.Id(userId).Get(users)
	userGrIds := strings.Split(chargeCards.UserGroupIds, ",")
	if chargeCards.UserGroupIds != "1" && !utils.InArrayString(users.UserGroupId, userGrIds) {
		utils.ResFaiJSON(&ctx, "", "当前账号不支持该银行卡充值", config.PARAMERROR)
		return
	}
	financial := NewFinancial(cthis.platform)
	//生成订单前缀公司银行卡
	orderId := utils.CreationOrder("COB", userIdS)
	//添加 检测用户从当前时间往前10分钟之内，若有超过10次未充值成功的
	checkRes := financial.CheckUserByAddChargeRecord(userIdS)
	if checkRes {
		utils.ResFaiJSON(&ctx, "", "短时间内重复充值次数过多,请稍后尝试!", config.CREDENTIALOSTOOMANY)
		return
	}

	ChargeTypeInfo := "银行卡转账支付"
	chargeRecord := new(xorm.ChargeRecords)
	created := int(time.Now().Unix())
	chargeRecord.UserId = userId
	chargeRecord.Amount = floatToString(payAmount)
	chargeRecord.OrderId = orderId
	chargeRecord.ChargeTypeId = chargeCards.ChargeTypeId //银行卡充值charge_types表id
	chargeRecord.CardNumber = cardNumber
	chargeRecord.BankAddress = bankAddress
	chargeRecord.Created = created
	chargeRecord.State = 0
	chargeRecord.Updated = created
	chargeRecord.ChargeTypeInfo = ChargeTypeInfo
	chargeRecord.Ip = utils.GetIp(ctx.Request())
	chargeRecord.PlatformId = 0
	chargeRecord.RealName = realName
	//这个字段暂时保留
	//银行转账类型([ "网银转账", "ATM自动柜员机", "ATM现金入款", "银行柜台转账", "手机银行转账", "其他" ])
	chargeRecord.BankTypeId = 0
	chargeRecord.BankChargeTime = created
	chargeRecord.IsTppay = 0
	chargeRecord.ChargeCardId = chargeCardsId
	chargeRecord.Remark = info
	affected, err := engine.Insert(chargeRecord)
	if err != nil {
		utils.ResFaiJSON(&ctx, "支付错误", "支付失败", config.PAYACTIONERROR)
		return
	}
	if affected == 0 {
		utils.ResFaiJSON(&ctx, "支付错误", "支付失败", config.PAYACTIONERROR)
		return
	}
	var res = make(map[string]interface{})
	res["orderNum"] = orderId
	res["amount"] = "+" + floatToString(payAmount)
	res["status"] = "入款中..."
	res["payTime"] = created
	res["payWay"] = bankName
	utils.ResSuccJSON(&ctx, "", "", config.SUCCESSRES, res)
}

/**
 * @api {post} api/auth/v1/payCompanyActionQrCan 创建公司入款二维码充值订单
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人:pling把妹</span><br/><br/>
 * 充值页面,addr_type=2 公司入款二维码充值,入款分俩种一种是第三方支付，一种是本平台直接充值，直接充值划为公司入款充值<br>
 * 业务描述:二维码充值，调用addr_type=2的时候,订单创建成功之后，当用户关闭二维码页面就显示结果</br>
 * @apiVersion 1.0.0
 * @apiName     api_auth_v1_payCompanyActionQrCan
 * @apiGroup    finance_module
 * @apiPermission ios,android客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 Bearer token
 * @apiParam (客户端请求参数) {float64} amount          	充值金额
 * @apiParam (客户端请求参数) {int} charge_cards_id     支付通道ID
  * @apiSuccess (返回结果)  {int}     code            200
 * @apiSuccess (返回结果)  {string}  clientMsg       提示信息
 * @apiSuccess (返回结果)  {string}  internalMsg     提示信息
 * @apiSuccess (返回结果)  {json}  	  data            返回数据支付信息
 * @apiSuccess (返回结果)  {float}   time_consumed   后台耗时
 * @apiSuccess (data对象字段说明) {string}   amount   金额说明
 * @apiSuccess (data对象字段说明) {string} orderNum   订单号
 * @apiSuccess (data对象字段说明) {int} payTime   支付时间
 * @apiSuccess (data对象字段说明) {string} payWay  支付方式
 * @apiSuccess (data对象字段说明) {string} status  支付状态
 * @apiSuccessExample {json} 响应结果
 {
    "clientMsg": "",
    "code": 200,
    "data": {
        "amount": "+100",
        "orderNum": "COBQR201905031548052c939c",
        "payTime": 1556869685,
        "payWay": "公司入款二维码支付",
        "status": "入款中..."
    },
    "internalMsg": "",
    "timeConsumed": 22478
}
*/

func (cthis *PayController) PayCompanyActionQrCan() {
	ctx := cthis.ctx
	userIdS := ctx.Values().GetString("userid")
	userId, _ := strconv.Atoi(userIdS)
	payAmount, existPayAmount := ctx.PostValueFloat64("amount")
	if existPayAmount != nil {
		utils.ResFaiJSON(&ctx, "不是数字类型", "请填写正确的充值金额", config.PARAMERROR)
		return
	}
	chargeCardsId, _ := ctx.PostValueInt("charge_cards_id")
	var chargeCards = new(xorm.ChargeCards)
	engine := models.MyEngine[cthis.platform]
	chCarOk, existChCarOk := engine.Id(chargeCardsId).Get(chargeCards)
	if existChCarOk != nil {
		utils.ResFaiJSON(&ctx, "sql错误", "充值失败", config.INTERNALERROR)
		return
	}
	if !chCarOk {
		utils.ResFaiJSON(&ctx, "该账号不存在", "充值账户不存在!", config.PARAMERROR)
		return
	}

	if chargeCards.AddrType != 2 {
		utils.ResFaiJSON(&ctx, "addrType必须为2", "只支持公司入款二维码扫描充值", config.PARAMERROR)
		return
	}

	var users = new(xorm.Users)
	engine.Id(userId).Get(users)
	userGrIds := strings.Split(chargeCards.UserGroupIds, ",")
	if chargeCards.UserGroupIds != "1" && !utils.InArrayString(users.UserGroupId, userGrIds) {
		utils.ResFaiJSON(&ctx, "", "当前账号不支持该充值方式", config.PARAMERROR)
		return
	}
	financial := NewFinancial(cthis.platform)
	//生成订单前缀公司银行卡
	orderId := utils.CreationOrder("COBQR", userIdS)
	//添加 检测用户从当前时间往前10分钟之内，若有超过10次未充值成功的
	checkRes := financial.CheckUserByAddChargeRecord(userIdS)
	if checkRes {
		utils.ResFaiJSON(&ctx, "", "短时间内重复充值次数过多,请稍后尝试!", config.CREDENTIALOSTOOMANY)
		return
	}

	ChargeTypeInfo := "公司入款二维码支付"
	chargeRecord := new(xorm.ChargeRecords)
	created := int(time.Now().Unix())
	chargeRecord.UserId = userId
	chargeRecord.Amount = floatToString(payAmount)
	chargeRecord.OrderId = orderId
	chargeRecord.ChargeTypeId = chargeCards.ChargeTypeId
	chargeRecord.CardNumber = chargeCards.CardNumber
	chargeRecord.BankAddress = chargeCards.BankAddress
	chargeRecord.Created = created
	chargeRecord.State = 0
	chargeRecord.Updated = created
	chargeRecord.ChargeTypeInfo = ChargeTypeInfo
	chargeRecord.Ip = utils.GetIp(ctx.Request())
	chargeRecord.PlatformId = 0
	chargeRecord.RealName = chargeCards.Owner
	//这个字段暂时保留
	//银行转账类型([ "网银转账", "ATM自动柜员机", "ATM现金入款", "银行柜台转账", "手机银行转账", "其他" ])
	chargeRecord.BankTypeId = 0
	chargeRecord.BankChargeTime = created
	chargeRecord.IsTppay = 0
	chargeRecord.ChargeCardId = chargeCardsId
	chargeRecord.Remark = chargeCards.Remark
	affected, err := engine.Insert(chargeRecord)
	if err != nil {
		utils.ResFaiJSON(&ctx, "支付错误", "支付失败", config.PAYACTIONERROR)
		return
	}
	if affected == 0 {
		utils.ResFaiJSON(&ctx, "支付错误", "支付失败", config.PAYACTIONERROR)
		return
	}
	var res = make(map[string]interface{})
	res["orderNum"] = orderId
	res["amount"] = "+" + floatToString(payAmount)
	res["status"] = "入款中..."
	res["payTime"] = created
	res["payWay"] = "公司入款二维码支付"
	utils.ResSuccJSON(&ctx, "", "", config.SUCCESSRES, res)
}

/**
 * @api {get} api/auth/v1/chargeRecordList 充值记录查询列表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人:pling把妹</span><br/><br/>
 * 充值页面充值记录查询<br>
 * 业务描述:进入充值页面点击充值记录，特别要注意当打回的数据少于一页多少条的时候说明没有数据了，就不要再请求了要提示没有更多数据</br>
 * @apiVersion 1.0.0
 * @apiName     api_auth_v1_chargeRecordList
 * @apiGroup    finance_module
 * @apiPermission ios,android客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 Bearer token
 * @apiParam (客户端请求参数) {int}   pageNum         当前页
 * @apiParam (客户端请求参数) {int}   state           充值状态：公司入款（0 待审核，1 成功，2 失败）；线上支付（0待处理，1成功，2失败，3进行中,4退款，5取消，6强制入款）
 * @apiSuccess (返回结果)  {int}     code            200,158空数组
 * @apiSuccess (返回结果)  {string}  clientMsg       提示信息
 * @apiSuccess (返回结果)  {string}  internalMsg     提示信息
 * @apiSuccess (返回结果)  {json}    data            返回数据支付信息
 * @apiSuccess (返回结果)  {float}   time_consumed    后台耗时
 * @apiSuccess (data对象字段说明) {array}   list   记录对象数组
 * @apiSuccess (data-list元素对象字段说明) {float64} amount   金额
 * @apiSuccess (data-list元素对象字段说明) {string} charge_type_info  支付说明
 * @apiSuccess (data-list元素对象字段说明) {int} created  时间戳
 * @apiSuccess (data-list元素对象字段说明) {string} order_id  订单号
 * @apiSuccess (data-list元素对象字段说明) {int} state  0待处理，1成功，2失败，3进行中,4退款，5取消，6强制入款
 * @apiSuccessExample {json} 响应结果
 * {
 *     "clientMsg": "",
 *     "code": 200,
 *     "data": {
 *         "list": [
 *             {
 *                 "amount": 100,
 *                 "charge_type_info": "公司入款二维码支付",
 *                 "created": 1559632212,
 *                 "is_tppay": 0, // 是否第三方支付 0为否；1为是
 *                 "order_id": "COBQR20190604151012b9",
 *                 "state": 0 // 公司入款：0 待审核，1 成功，2 失败。线上支付：0待处理，1成功，2失败，3进行中,4退款，5取消，6强制入款
 *             }
 *         ]
 *     },
 *     "internalMsg": "",
 *     "timeConsumed": 170624
 * }
 */

func (cthis *PayController) ChargeRecordList() {
	ctx := cthis.ctx
	userIdS := ctx.Values().GetString("userid")
	userId, _ := strconv.Atoi(userIdS)
	page, errPage := ctx.URLParamInt("pageNum")
	chargeStatus := ctx.URLParam("state")
	if errPage != nil || page <= 0 {
		page = 1
	}
	//每页条数
	pageCount := 10
	recordStart := (page - 1) * pageCount
	engine := models.MyEngine[cthis.platform]
	var chargeRecords = make([]xorm.ChargeRecords, 0)

	var whereBuilder = builder.Eq{
		"user_id": userId,
	}

	if chargeStatus != "" {
		whereBuilder = builder.Eq{
			"user_id": userId,
			"state":   chargeStatus,
		}
	}

	queryError := engine.Cols("created", "charge_type_info", "state", "amount", "order_id", "is_tppay").
		Where(whereBuilder).Desc("created").
		Limit(pageCount, recordStart).Find(&chargeRecords)
	var res = make(map[string]interface{})
	if len(chargeRecords) == 0 || queryError != nil {
		res["list"] = []string{}
		utils.ResSuccJSON(&ctx, "", "", config.SUCCESSRES, res)
		return
	}
	var list = make([]map[string]interface{}, 0)
	for _, v := range chargeRecords {
		var ele = make(map[string]interface{})
		ele["created"] = v.Created
		ele["charge_type_info"] = v.ChargeTypeInfo
		ele["state"] = v.State
		ele["is_tppay"] = v.IsTppay
		amount, _ := strconv.ParseFloat(v.Amount, 64)
		ele["amount"] = amount
		ele["order_id"] = v.OrderId
		list = append(list, ele)
	}
	res["list"] = list
	utils.ResSuccJSON(&ctx, "", "", config.SUCCESSRES, res)
}

/**
 * @api {get} api/auth/v1/withDrawRecordsList 提现记录查询列表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人:pling把妹</span><br/><br/>
 * 提现页面提现记录<br>
 * 业务描述:提现页面记录列表</br>
 * @apiVersion 1.0.0
 * @apiName     api_auth_v1_withDrawRecordsList
 * @apiGroup    finance_module
 * @apiPermission ios,android客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 Bearer token
  * @apiSuccess (返回结果)  {int}     code            200,158空数组
 * @apiSuccess (返回结果)  {string}  clientMsg       提示信息
 * @apiSuccess (返回结果)  {string}  internalMsg     提示信息
 * @apiSuccess (返回结果)  {json}  	  data            返回数据支付信息
 * @apiSuccess (返回结果)  {float}   time_consumed    后台耗时

 * @apiSuccess (data对象字段说明) {array}   list   记录对象数组
 * @apiSuccess (data-list元素对象字段说明) {float64} amount   金额
 * @apiSuccess (data-list元素对象字段说明) {string} bank_name  银行名称
 * @apiSuccess (data-list元素对象字段说明) {string} card_number  银行卡号
 * @apiSuccess (data-list元素对象字段说明) {int} created  时间戳
 * @apiSuccess (data-list元素对象字段说明) {string} order_id  订单号
 * @apiSuccess (data-list元素对象字段说明) {int} status  0 待审核,1 提现成功,2 退回出款 3 锁定
 * @apiSuccess (data-list元素对象字段说明) {string} withdraw_type 提现类型说明
 * @apiSuccessExample {json} 响应结果
 {
    "clientMsg": "",
    "code": 200,
    "data": {
        "list": [
            {
                "amount": "100.00",
                "bank_name": "工商银行",
                "card_number": "6222032410000296282",
                "created": 1529471708,
                "order_id": "WD-sfsfsfjdk76sf",
                "status": 1,
                "withdraw_type": "提款到银行卡"
            }
        ]
    },
    "internalMsg": "",
    "timeConsumed": 7114
}
*/

func (cthis *PayController) WithDrawRecordsList() {
	ctx := cthis.ctx
	userIdS := ctx.Values().GetString("userid")
	userId, _ := strconv.Atoi(userIdS)
	page, errPage := ctx.URLParamInt("page")
	if errPage != nil || page <= 0 {
		page = 1
	}
	//每页条数
	pageNum := 10
	recordStart := (page - 1) * pageNum
	engine := models.MyEngine[cthis.platform]
	var withdrawRecords = make([]xorm.WithdrawRecords, 0)
	queryError := engine.Where("user_id = ?", userId).
		Limit(pageNum, recordStart).Find(&withdrawRecords)
	var res = make(map[string]interface{})
	if len(withdrawRecords) == 0 || queryError != nil {
		res["list"] = []string{}
		utils.ResSuccJSON(&ctx, "", "", config.SUCCESSRES, res)
		return
	}
	var list = make([]map[string]interface{}, 0)
	for _, v := range withdrawRecords {
		var ele = make(map[string]interface{})
		ele["created"] = v.Created
		ele["status"] = v.Status
		ele["amount"] = v.Amount
		ele["order_id"] = v.OrderId
		ele["card_number"] = v.CardNumber
		ele["withdraw_type"] = v.WithdrawType
		ele["bank_name"] = v.BankName
		list = append(list, ele)
	}
	res["list"] = list
	utils.ResSuccJSON(&ctx, "", "", config.SUCCESSRES, res)
}

/**
 * @api {post} api/auth/v1/bindingBankCard 银行卡绑定接口
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人:pling把妹</span><br/><br/>
 * 提现页面银行卡绑定<br>
 * 业务描述:提现页面银行卡管理绑定银行卡</br>
 * @apiVersion 1.0.0
 * @apiName     api_auth_v1_bindingBankCard
 * @apiGroup    finance_module
 * @apiPermission ios,android客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 Bearer token
 * @apiParam (客户端请求参数) {string} bank_name          银行卡名称
 * @apiParam (客户端请求参数) {string} name          姓名
 * @apiParam (客户端请求参数) {string} card_number          银行卡卡号
 * @apiParam (客户端请求参数) {string} bank_address          开户行地址
 * @apiSuccess (返回结果)  {int}     code            200成功
 * @apiSuccess (返回结果)  {string}  clientMsg       提示信息
 * @apiSuccess (返回结果)  {string}  internalMsg     提示信息
 * @apiSuccess (返回结果)  {json}  	  data            返回空对象
 * @apiSuccess (返回结果)  {float}   time_consumed    后台耗时
 * @apiSuccessExample {json} 响应结果
 {
    "clientMsg": "银行卡绑定成功",
    "code": 200,
    "data": {},
    "internalMsg": "",
    "time_consumed": 14176
}
*/

func (cthis *PayController) BindingBankCard() {
	ctx := cthis.ctx
	userIdS := ctx.Values().GetString("userid")
	userId, _ := strconv.Atoi(userIdS)
	if !utils.RequiredParamPost(&ctx, []string{"bank_name", "name", "card_number", "bank_address"}) {
		return
	}
	bankName := strings.TrimSpace(ctx.PostValue("bank_name"))
	name := strings.TrimSpace(ctx.PostValue("name"))
	bankCardNumber := strings.TrimSpace(ctx.PostValue("card_number"))
	bankAddress := strings.TrimSpace(ctx.PostValue("bank_address"))
	var userBankCard = new(xorm.UserBankCards)
	engine := models.MyEngine[cthis.platform]
	totalNum, errTotalNum := engine.Cols("id").Where("user_id = ?", userId).Count(userBankCard)
	if errTotalNum != nil {
		utils.ResFaiJSON(&ctx, "数据库操作错误", "绑定失败", config.NOTGETDATA)
		return
	}
	if totalNum >= 1 {
		utils.ResFaiJSON(&ctx, "", "银行卡只能绑定一张", config.NOTGETDATA)
		return
	}
	if utf8.RuneCountInString(name) > 6 {
		utils.ResFaiJSON(&ctx, "", "姓名不能超过6个字符", config.PARAMERROR)
		return
	}
	if utf8.RuneCountInString(bankName) > 20 {
		utils.ResFaiJSON(&ctx, "", "银行名称不能超过20个文字", config.PARAMERROR)
		return
	}

	if utf8.RuneCountInString(bankCardNumber) > 20 {
		utils.ResFaiJSON(&ctx, "", "银行卡号不能超过20位", config.PARAMERROR)
		return
	}
	if utf8.RuneCountInString(bankAddress) > 30 {
		utils.ResFaiJSON(&ctx, "", "开户行地址不能超过30个字符", config.PARAMERROR)
		return
	}
	//如果是第一次绑定要更新用户表姓名
	if totalNum == 0 {
		users := new(xorm.Users)
		users.Name = name
		_, errUpdateRow := engine.Cols("name").Where("id = ?", userId).Update(users)
		if errUpdateRow != nil {
			utils.ResFaiJSON(&ctx, "操作数据库错误", "服务器繁忙,稍后再试", config.INTERNALERROR)
			return
		}
	}

	userBankCard.Name = name
	userBankCard.BankName = bankName
	userBankCard.CardNumber = bankCardNumber
	userBankCard.Address = bankAddress
	userBankCard.Created = int(time.Now().Unix())
	userBankCard.UserId = userId
	userBankCard.Status = 1
	userBankCard.Updated = int(time.Now().Unix())
	_, errInsert := engine.Insert(userBankCard)
	if errInsert != nil || userBankCard.Id == 0 {
		utils.ResFaiJSON(&ctx, "数据库操作插入失败", "绑定失败", config.NOTGETDATA)
		return
	}
	utils.ResSuccJSON(&ctx, "", "银行卡绑定成功", config.SUCCESSRES, make(map[string]interface{}))
}

/**
 * @api {get} api/v1/bankCardsList 银行卡列表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人:pling把妹</span><br/><br/>
 * 提现页面绑定银行卡的银行卡列表选择<br>
 * 业务描述:银行卡列表选择</br>
 * @apiVersion 1.0.0
 * @apiName     api_v1_bankCardsList
 * @apiGroup    finance_module
 * @apiPermission ios,android客户端
 * @apiParam (客户端请求参数) {string} cache_key   缓存md5,第一次为空获取最新内容,第二次请求需要带上此key
 * @apiSuccess (返回结果)  {int}     code            200
 * @apiSuccess (返回结果)  {string}  clientMsg       提示信息
 * @apiSuccess (返回结果)  {string}  internalMsg     提示信息
 * @apiSuccess (返回结果)  {json}  	  data            返回数据支付信息
 * @apiSuccess (返回结果)  {float}   time_consumed    后台耗时
 * @apiSuccess (data json对象字段说明)  {string}   cache_key    缓存md5，如果缓存没变将不存在这个值
 * @apiSuccess (data json对象字段说明)  {array}   list    银行列表元素对象
 * @apiSuccess (data-list数组元素对象字段说明) {int} id 银行卡Id
 * @apiSuccess (data-list数组元素对象字段说明) {string} name 银行卡名称
 * @apiSuccessExample {json} 响应结果
 {
    "clientMsg": "",
    "code": 200,
    "data": {
        "cache_key": "8602d33c0d5e788a8122251ff9de45c1",
        "list": [
            {
                "id": 1,
                "name": "中国银行"
            },
           .....
        ]
    },
    "internalMsg": "",
    "timeConsumed": 180
}
*/

func (cthis *PayController) BankCardsList() {
	ctx := cthis.ctx
	cacheKey := ctx.URLParam("cache_key")
	bankCards, _ := ramcache.TableUserBanks.Load(cthis.platform)
	var cards = make([]map[string]interface{}, 0)
	var list = make([]map[string]interface{}, 0)
	var res = make(map[string]interface{})
	var bankCrds = bankCards.([]xorm.UserBanks)
	for _, v := range bankCrds {
		value := make(map[string]interface{})
		value["id"] = v.Id
		value["name"] = v.Name
		cards = append(cards, value)
	}
	byteJson, _ := json.Marshal(cards)
	cacheKeyMd5 := fmt.Sprintf("%x", md5.Sum(byteJson))
	if cacheKeyMd5 != cacheKey {
		list = cards
		res["cache_key"] = cacheKeyMd5
	}
	res["list"] = list
	utils.ResSuccJSON(&ctx, "", "", config.SUCCESSRES, res)
}

/**
 * @api {get} api/auth/v1/userBankCards 用户银行卡列表信息
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人:pling把妹</span><br/><br/>
 * 在进入提现页面的时候需要显示用户已经绑定的银行卡列表,本地缓存后先展示缓存的，在去请求接口<br>
 * 业务描述:用户银行卡信息，原本打算放入个人信息接口中,为了更好管理还是单独列为一个接口</br>
 *	背景图片就不弄了，用一个公用的图片代替</br>
 * @apiVersion 1.0.0
 * @apiName     api_auth_v1_userBankCards
 * @apiGroup    finance_module
 * @apiPermission ios,android客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 Bearer token
 * @apiParam (客户端请求参数) {string} cache_key   缓存md5,第一次为空获取最新内容,第二次请求需要带上此key
 * @apiSuccess (返回结果)  {int}     code            200
 * @apiSuccess (返回结果)  {string}  clientMsg       提示信息
 * @apiSuccess (返回结果)  {string}  internalMsg     提示信息
 * @apiSuccess (返回结果)  {json}  	  data            返回数据支付信息
 * @apiSuccess (返回结果)  {float}   time_consumed    后台耗时
 * @apiSuccess (data json对象字段说明)  {string}   cache_key    缓存md5，如果缓存没变将不存在这个值
 * @apiSuccess (data json对象字段说明)  {array}   list    银行列表元素对象
 * @apiSuccess (data-list数组元素对象字段说明) {int} user_bank_id 用户银行卡表Id
 * @apiSuccess (data-list数组元素对象字段说明) {string} bank_name 银行卡名称
 * @apiSuccess (data-list数组元素对象字段说明) {string} card_number 银行卡号
 * @apiSuccessExample {json} 响应结果
 {
    "clientMsg": "",
    "code": 200,
    "data": {
        "cache_key": "ad033a3c45627b1a91b8fbc17dc982f5",
        "list": [
            {
                "bank_name": "中国银行",
                "card_number": "123456789",
                "user_bank_id": 3
            }
        ]
    },
    "internalMsg": "",
    "timeConsumed": 2626
}
*/

func (cthis *PayController) UserBankCards() {
	ctx := cthis.ctx
	userIdS := ctx.Values().GetString("userid")
	userId, _ := strconv.Atoi(userIdS)
	cacheKey := ctx.URLParam("cache_key")
	var cards = make([]xorm.UserBankCards, 0)
	engine := models.MyEngine[cthis.platform]
	err := engine.Cols("card_number", "bank_name", "id").Where("user_id = ?", userId).Find(&cards)
	if err != nil {
		utils.ResFaiJSON(&ctx, "数据库操作插入失败", "获取银行卡信息失败", config.NOTGETDATA)
		return
	}
	var list = make([]map[string]interface{}, 0)
	var res = make(map[string]interface{})
	res["list"] = list
	if len(cards) == 0 {
		utils.ResSuccJSON(&ctx, "", "", config.SUCCESSRES, res)
		return
	}
	var bankCrds = make([]map[string]interface{}, 0)
	for _, v := range cards {
		value := make(map[string]interface{})
		value["user_bank_id"] = v.Id
		value["bank_name"] = v.BankName
		value["card_number"] = v.CardNumber
		bankCrds = append(bankCrds, value)
	}
	byteJson, _ := json.Marshal(bankCrds)
	cacheKeyMd5 := fmt.Sprintf("%x", md5.Sum(byteJson))
	if cacheKeyMd5 != cacheKey {
		list = bankCrds
		res["cache_key"] = cacheKeyMd5
	}
	res["list"] = list
	utils.ResSuccJSON(&ctx, "", "", config.SUCCESSRES, res)

}

/**
* @api {post} api/auth/v1/withDrawAction 提现接口
* @apiDescription
* <span style="color:lightcoral;">接口负责人:pling把妹</span><br/><br/>
* 提现页面的提现接口<br>
* 业务描述:提现页面点击提现</br>
* @apiVersion 1.0.0
* @apiName     api_auth_v1_withDrawAction
* @apiGroup    finance_module
* @apiPermission ios,android客户端
* @apiHeader (请求头) {string} Authorization 用户令牌  格式为 Bearer token
* @apiParam (客户端请求参数) {int}    amount          提现金额必须为整数
* @apiParam (客户端请求参数) {int}    user_bank_id          用户银行卡id
 * @apiSuccess (返回结果)  {int}     code            200成功
* @apiSuccess (返回结果)  {string}  clientMsg       提示信息
* @apiSuccess (返回结果)  {string}  internalMsg     提示信息
* @apiSuccess (返回结果)  {json}  	  data            返回空对象
* @apiSuccess (返回结果)  {float}   time_consumed    后台耗时
* @apiSuccessExample {json} 响应结果
{
   "clientMsg": "提现成功",
   "code": 200,
   "data": {},
   "internalMsg": "",
   "timeConsumed": 40341
 }
*/

func (cthis *PayController) WithDrawAction() {
	ctx := cthis.ctx
	userIdS := ctx.Values().GetString("userid")
	userId, _ := strconv.Atoi(userIdS)
	amount, _ := ctx.PostValueFloat64("amount")
	userBankId, _ := ctx.PostValueInt("user_bank_id")
	conf, _ := ramcache.TableConfigs.Load(cthis.platform)
	cfg := conf.(map[string]interface{})
	withdrawMinMoney := cfg["withdraw_min_money"].(float64)
	withdrawMaxMoney := cfg["withdraw_max_money"].(float64)
	if amount < withdrawMinMoney {
		utils.ResFaiJSON(&ctx, "", "提现金额不能少于"+floatToString(withdrawMinMoney), config.PARAMERROR)
		return
	}
	if amount > withdrawMaxMoney {
		utils.ResFaiJSON(&ctx, "", "提现金额不能大于"+floatToString(withdrawMaxMoney), config.PARAMERROR)
		return
	}
	engine := models.MyEngine[cthis.platform]
	userBankCards := new(xorm.UserBankCards)
	engine.ID(userBankId).Where("user_id = ?", userId).Get(userBankCards)
	//银行卡有效性检查
	if userBankCards.Id == 0 {
		utils.ResFaiJSON(&ctx, "银行卡Id错误", "不存在该银行卡账号", config.PARAMERROR)
		return
	}

	//账号余额检查
	account := new(xorm.Accounts)
	engine.Cols("balance_wallet", "id").Where("user_id = ?", userId).Get(account)
	if account.Id == 0 {
		utils.ResFaiJSON(&ctx, "数据库查询用户错误", "服务器繁忙,稍后再试", config.INTERNALERROR)
		return
	}
	balanceWallet, _ := strconv.ParseFloat(account.BalanceWallet, 64)
	if balanceWallet < amount {
		utils.ResFaiJSON(&ctx, "", "余额不足,无法提现", config.PARAMERROR)
		return
	}
	//打码量满足条件检查
	wdamaR := new(xorm.WithdrawDamaRecords)
	amountTotal, errTotal := engine.Where("state = 0 and user_id = ?", userId).Sum(wdamaR, "amount")
	if errTotal != nil {
		utils.ResFaiJSON(&ctx, "数据库查询错误", "服务器繁忙,稍后再试", config.INTERNALERROR)
		return
	}
	//未完成打码量总金额
	deciAmount := decimal.NewFromFloat(amountTotal)
	//钱包余额
	deciBalWal := decimal.NewFromFloat(balanceWallet)
	//当前提现金额
	wDAmount := decimal.NewFromFloat(amount)
	//可提现金额
	canWiDr := deciBalWal.Sub(deciAmount)
	if !canWiDr.GreaterThanOrEqual(wDAmount) {
		sCanWithdraw := canWiDr.String()
		if canWiDr.LessThan(decimal.New(0, 0)) {
			sCanWithdraw = "0"
		}
		utils.ResFaiJSON(&ctx, "", "未满足打码量提现失败,最多可提现"+sCanWithdraw, config.INTERNALERROR)
		return
	}
	localTime, _ := time.LoadLocation("Asia/Shanghai")
	today0 := time.Now().Format("2006-01-02")
	//早上零点
	st, _ := time.ParseInLocation("2006-01-02", today0, localTime)
	//当天23点59分
	st2, _ := time.ParseInLocation("2006-01-02 15:04:05", today0+" 23:59:59", localTime)
	timeToday0 := st.Unix()
	timeTodayEnd := st2.Unix()
	//当天提现次数限制
	alreday, _ := engine.Cols("id").
		Where("create_time between ? and ? and user_id = ?", timeToday0, timeTodayEnd, userId).
		Count(&xorm.WithdrawRecords{})
	withdrawDayLimited := int64(cfg["withdraw_day_limited"].(float64))
	if alreday >= withdrawDayLimited {
		utils.ResFaiJSON(&ctx, "", "当天提现次数不能超过三次", config.INTERNALERROR)
		return
	}
	//资金流水日志记录
	balance := fund.NewUserFundChange(cthis.platform) //给用户充值
	orderId := utils.CreationOrder("WD", userIdS)
	info := map[string]interface{}{
		"user_id":  userId,
		"type_id":  config.FUNDWITHDRAW,
		"amount":   -amount,
		"order_id": orderId,
		"msg":      "用户提现",
	}
	//关联回调处理
	callback := func(session *goxorm.Session, args ...interface{}) (interface{}, error) {
		withdrawRecord := new(xorm.WithdrawRecords)
		withdrawRecord.UserId = userId
		withdrawRecord.Amount = decimal.NewFromFloat(amount).String()
		withdrawRecord.OrderId = orderId
		withdrawRecord.Updated = int(time.Now().Unix())
		withdrawRecord.Created = int(time.Now().Unix())
		withdrawRecord.CardNumber = userBankCards.CardNumber
		withdrawRecord.RealName = userBankCards.Name
		withdrawRecord.BankAddress = userBankCards.Address
		withdrawRecord.BankName = userBankCards.BankName
		withdrawRecord.WithdrawType = "在线提款"
		_, errWith := session.InsertOne(withdrawRecord)
		if errWith != nil || withdrawRecord.Id == 0 {
			return nil, errWith
		}
		return nil, nil
	}
	//创建资金流水表
	balanceUpdateRes := balance.BalanceUpdate(info, callback)
	if balanceUpdateRes["status"] != 1 {
		utils.ResFaiJSON(&ctx, "", "提现失败", config.INTERNALERROR)
		return
	}
	utils.ResSuccJSON(&ctx, "", "提现成功", config.SUCCESSRES, make(map[string]interface{}))
}

/**
* @api {get} api/auth/v1/runningWaterDetail 提现页面流水详情接口
* @apiDescription
* <span style="color:lightcoral;">接口负责人:pling把妹</span><br/><br/>
* 提现页面流水详情接口<br>
* 业务描述:提现页面流水详情</br>
* @apiVersion 1.0.0
* @apiName     api_auth_v1_runningWaterDetail
* @apiGroup    finance_module
* @apiPermission ios,android客户端
* @apiHeader (请求头) {string} Authorization 用户令牌  格式为 Bearer token
 * @apiSuccess (返回结果)  {int}     code            200成功
* @apiSuccess (返回结果)  {string}  clientMsg       提示信息
* @apiSuccess (返回结果)  {string}  internalMsg     提示信息
* @apiSuccess (返回结果)  {json}  	  data            数据对象
* @apiSuccess (返回结果)  {float}   time_consumed    后台耗时
* @apiSuccess (data对象字段说明) {array}   list   记录对象数组
* @apiSuccess (data-list元素对象字段说明)  {int}  	  created    时间戳
* @apiSuccess (data-list元素对象字段说明)  {decimal}  	  finished_needed    需求打码
* @apiSuccess (data-list元素对象字段说明)  {decimal}  	  finished_progress    实际打码
* @apiSuccess (data-list元素对象字段说明)  {int}  	  state    流水状态 0未完成，1已完成，2已失效
* @apiSuccessExample {json} 响应结果
{
    "clientMsg": "",
    "code": 200,
    "data": {
        "list": [
            {
                "created": 1556190428,
                "finished_needed": 100,
                "finished_progress": 20,
                "state": 3
            }
        ],
    },
    "internalMsg": "",
    "timeConsumed": 2582
}
*/

func (cthis *PayController) RunningWaterDetail() {
	ctx := cthis.ctx
	userIdS := ctx.Values().GetString("userid")
	userId, _ := strconv.Atoi(userIdS)
	page, errPage := ctx.URLParamInt("page")
	if errPage != nil || page <= 0 {
		page = 1
	}
	//每页条数
	pageNum := 10
	recordStart := (page - 1) * pageNum
	wiDrdamaR := make([]xorm.WithdrawDamaRecords, 0)
	var list = make([]map[string]interface{}, 0)
	var res = make(map[string]interface{})
	engine := models.MyEngine[cthis.platform]
	queryError := engine.Where("user_id = ?", userId).
		Limit(pageNum, recordStart).Find(&wiDrdamaR)
	if len(wiDrdamaR) == 0 || queryError != nil {
		res["list"] = []string{}
		utils.ResSuccJSON(&ctx, "", "", config.SUCCESSRES, res)
		return
	}
	for _, v := range wiDrdamaR {
		var temp = make(map[string]interface{})
		temp["created"] = v.Created
		//需求打码
		flFneed, _ := strconv.ParseFloat(v.FinishedNeeded, 64)
		temp["finished_needed"] = flFneed
		//实际打码
		flFPr, _ := strconv.ParseFloat(v.FinishedProgress, 64)
		temp["finished_progress"] = flFPr
		//0未完成，1已完成，2已失效
		temp["state"] = v.State
		list = append(list, temp)
	}
	res["list"] = list
	utils.ResSuccJSON(&ctx, "", "", config.SUCCESSRES, res)
}

/**
* @api {get} api/auth/v1/withdrawFundParticulars 提现页面资金明细接口
* @apiDescription
* <span style="color:lightcoral;">接口负责人:pling把妹</span><br/><br/>
* 提现页面资金明细接口<br>
* 业务描述:提现页面资金明细</br>
* @apiVersion 1.0.0
* @apiName     api_auth_v1_withdrawFundParticulars
* @apiGroup    finance_module
* @apiPermission ios,android客户端
* @apiHeader (请求头) {string} Authorization 用户令牌  格式为 Bearer token
 * @apiSuccess (返回结果)  {int}     code            200成功
* @apiSuccess (返回结果)  {string}  clientMsg       提示信息
* @apiSuccess (返回结果)  {string}  internalMsg     提示信息
* @apiSuccess (返回结果)  {json}  	  data            数据对象
* @apiSuccess (返回结果)  {float}   time_consumed    后台耗时
* @apiSuccess (data对象字段说明) {array}   list   记录对象数组
* @apiSuccess (data-list元素对象字段说明)  {int}  	  created    时间戳
* @apiSuccess (data-list元素对象字段说明)  {string}  	  charge_type_info    充值类型说明
* @apiSuccess (data-list元素对象字段说明)  {decimal}  	  amount    充值金额
* @apiSuccessExample {json} 响应结果
{
    "clientMsg": "",
    "code": 200,
    "data": {
        "list": [
            {
                "amount": 100,
                "charge_type_info": "线上充值",
                "created": 1556329786
            },
            {
                "amount": 100,
                "charge_type_info": "优惠入库",
                "created": 1556332530
            }
        ]
    },
    "internalMsg": "",
    "timeConsumed": 2656
}
*/

func (cthis *PayController) WithdrawFundParticulars() {
	ctx := cthis.ctx
	userIdS := ctx.Values().GetString("userid")
	userId, _ := strconv.Atoi(userIdS)
	page, errPage := ctx.URLParamInt("page")
	if errPage != nil || page <= 0 {
		page = 1
	}
	//每页条数
	pageNum := 10
	recordStart := (page - 1) * pageNum
	accountInfos := make([]xorm.AccountInfos, 0)
	var list = make([]map[string]interface{}, 0)
	var res = make(map[string]interface{})
	engine := models.MyEngine[cthis.platform]
	types := []int{
		config.FUNDCHARGE,        //充值
		config.FUNDXIMA,          //洗码
		config.FUNDPRESENTER,     //赠送彩金
		config.FUNDDISCOUNTS,     //优惠入款
		config.FUNDBROKERAGE,     //代理佣金提成
		config.FUNDACTIVITYAWARD, //活动奖励
		config.FUNDREDPACKET,     //红包收入
	}
	whereIn := "type in ("
	for _, v := range types {
		whereIn += strconv.Itoa(v) + ","
	}
	whereIn = strings.TrimRight(whereIn, ",")
	whereIn += ")"
	queryError := engine.Where("user_id = ?", userId).
		And(whereIn).Limit(pageNum, recordStart).Find(&accountInfos)
	if len(accountInfos) == 0 || queryError != nil {
		res["list"] = []string{}
		utils.ResSuccJSON(&ctx, "", "", config.SUCCESSRES, res)
		return
	}

	for _, v := range accountInfos {
		var temp = make(map[string]interface{})
		temp["created"] = v.Created
		//充值类型说明
		temp["charge_type_info"] = v.Msg
		//充值金额
		amountDecil, _ := decimal.NewFromString(v.Amount)
		//如果小于0就乘以-1,变为正数
		if amountDecil.LessThan(decimal.NewFromFloat(0.0)) {
			amountDecil = amountDecil.Mul(decimal.NewFromFloat(-1))
		}
		amountFl, _ := amountDecil.Float64()
		temp["amount"] = amountFl
		list = append(list, temp)
	}
	res["list"] = list
	utils.ResSuccJSON(&ctx, "", "", config.SUCCESSRES, res)
}

//支付同步回调通知
func (cthis *PayController) TpPayClientNotify() {
	ctx := cthis.ctx
	ctx.Gzip(true)
	ctx.View("paymentSuccess.html")
}

//支付异步回调通知
func (cthis *PayController) TppayCallBack() {
	ctx := cthis.ctx
	userId, _ := ctx.Params().GetInt("userId")
	credentialId, _ := ctx.Params().GetInt("credentialId")
	payNameCode, _ := ctx.Params().GetInt("payNameCode")
	method := ctx.Method()
	var data = make(map[string]string)
	if method == "GET" {
		data = ctx.URLParams()
	}
	if method == "POST" {
		getData := ctx.URLParams()
		postData := ctx.FormValues()
		var temp = make(map[string]string)
		for k, v := range getData {
			temp[k] = v
		}
		for k1, v1 := range postData {
			if len(v1) == 1 {
				temp[k1] = v1[0]
			}
			if len(v1) == 0 {
				temp[k1] = ""
			}
			//如果大于2个实际开发中有遇到的话就要特殊处理了
		}
		data = temp
	}
	result := cthis.checkSign2(data, payNameCode, credentialId)
	result["amount"], _ = strconv.ParseFloat(result["amount"].(string), 64)
	if len(result) == 0 {
		return
	}

	if result["status"] == 1 { //支付成功
		res := cthis.updateChargeRecord(userId, result["orderId"].(string), result["amount"].(float64), 1)
		if !res {
			return
		}
		balance := fund.NewUserFundChange(cthis.platform) //给用户充值
		info := map[string]interface{}{
			"user_id":     userId,
			"type_id":     config.FUNDCHARGE,
			"amount":      result["amount"],
			"order_id":    result["orderId"],
			"msg":         "用户充值",
			"finish_rate": 1.0, //需满足的打码量比例
		}
		//
		balanceUpdateRes := balance.BalanceUpdate(info, nil)
		if balanceUpdateRes["status"] == 1 {
			//充值成功
			ctx.WriteString(result["sign"].(string))

			// 新启个session，不影响充值
			engine := models.MyEngine[cthis.platform]
			session := engine.NewSession()
			session.Begin()
			defer session.Close()
			sIp := utils.GetIp(ctx.Request())
			_, err := services.ActivityAward(cthis.platform, session, 2, userId, sIp)
			if err != nil {
				session.Rollback()
				fmt.Println("充值奖励发放失败", err.Error())
			}
		} else {
			//充值失败
			return
		}
	} else {
		_, exitOrderId := result["orderId"]
		if exitOrderId {
			ctx.WriteString("订单号异常")
			return
		}
		amount, exitAmount := result["amount"]
		if exitAmount {
			amount = 0.00
		}
		status, exitStatus := result["status"]
		if exitStatus {
			status = 2
		}
		cthis.updateChargeRecord(userId, result["orderId"].(string), amount.(float64), status.(int))
	}
}

func (cthis *PayController) checkSign2(data map[string]string, payNameCode int, credentialId int) map[string]interface{} {
	payC, _ := ramcache.TablePayCredential.Load(cthis.platform)
	//获取paycredentials表对应Id数据
	rowCred := payC.(map[int]xorm.PayCredentials)[credentialId]
	var result map[string]interface{}
	switch payNameCode {
	case 100: //金砖支付
		return tpPay100JinZhuanCallback(data, rowCred)
	}
	return result
}

func (cthis *PayController) updateChargeRecord(userId int, orderId string, amount float64, status int) bool {
	var chargeRecord xorm.ChargeRecords
	engine := models.MyEngine[cthis.platform]
	errRow := engine.Cols("state").
		Where("user_id = ", userId).
		Where("order_id = ", orderId).
		Find(&chargeRecord)
	if errRow != nil {
		return false
	}
	if chargeRecord.State == 1 || chargeRecord.State == 6 {
		return false
	}
	chargeRecord.State = status
	chargeRecord.Amount = floatToString(amount)
	chargeRecord.Updated = int(time.Now().Unix())
	affrow, errRow := engine.Cols("state", "amount", "updated").
		Where("user_id = ", userId).
		Where("order_id = ", orderId).
		Update(chargeRecord)
	if errRow != nil {
		return false
	}
	if affrow > 0 {
		return true
	} else {
		return false
	}
}

/**
* 第三方支付处理
 */
func (cthis *PayController) TPWap() {
	ctx := cthis.ctx
	data := ctx.URLParams()
	echoForm, exitsEchoForm := data["echo_form"]
	if exitsEchoForm {
		echoFormTxt, _ := url.QueryUnescape(echoForm)
		data, _ := base64.StdEncoding.DecodeString(echoFormTxt)
		ctx.WriteString(string(data))
		return
	}

	url, exitsWapUrl := data["wap_url"]
	if !exitsWapUrl {
		utils.ResFaiJSON(&ctx, "", "该支付已关闭，请更换支付方式", config.CREDENTIALSTOP)
		return
	}

	ajax, exitsAjax := data["ajax"]
	if !exitsAjax {
		ajax = "0"
	}
	delete(data, "url")

	dataTemp := ""
	dataTemp = dataTemp
	for k, v := range data {
		dataTemp += `<input type="hidden"  name="` + k + `" value="` + v + `" />`
	}

	txt := cthis.unescaped(dataTemp)
	ctx.ViewData("data", txt.(template.HTML))
	ctx.ViewData("url", url)
	ctx.ViewData("ajax", ajax)
	ctx.View("tpwap.html")
	return
}
func (cthis *PayController) unescaped(x string) interface{} { return template.HTML(x) }
