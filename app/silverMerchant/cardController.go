package silverMerchant

import (
	"github.com/kataras/iris"
	"github.com/shopspring/decimal"
	adminModels "qpgame/admin/models"
	"qpgame/common/mvc"
	"qpgame/common/services"
	"qpgame/common/utils"
	"qpgame/config"
	"qpgame/models"
	"qpgame/models/xorm"
	"strconv"
	"time"
)

type CardController struct {
	platform string
	ctx      iris.Context
}

//构造函数
func NewSilverCardController(ctx iris.Context) *CardController {
	obj := new(CardController)
	obj.platform = ctx.Params().Get("platform")
	obj.ctx = ctx
	return obj
}



/**
 * @api {get} silverMerchant/api/auth/v1/getChargeCard 获取充值银行卡
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: Singh</span><br/><br/>
 * 获取充值银行卡<br>
 * 业务描述:获取充值银行卡信息</br>
 * @apiVersion 1.0.0
 * @apiName     getChargeCard
 * @apiGroup    silver_merchant

 * @apiErrorExample {json} 失败返回
 * {
 * 		"clientMsg": "暂无充值银行卡",
 * 		"code": 204,
 * 		"internalMsg": "Oops",
 * 		"timeConsumed": 5915
 * }
 * @apiSuccess (返回结果)  {int}     code           200
 * @apiSuccess (返回结果)  {string}  clientMsg      提示信息
 * @apiSuccess (返回结果)  {string}  internalMsg    提示信息
 * @apiSuccess (返回结果)  {float}   timeConsumed   后台耗时
 * @apiSuccess (返回结果)  {json}    data           返回数据
 * @apiSuccess (data对象字段说明) {int} id	自增ID
 * @apiSuccess (data对象字段说明) {string} name 银行名
 * @apiSuccess (data对象字段说明) {string} owner 银行卡名字
 * @apiSuccess (data对象字段说明) {string} card_number 银行卡号
 * @apiSuccess (data对象字段说明) {string} bank_address 开户地址
 * @apiSuccess (data对象字段说明) {string} remark 备注
 * @apiSuccess (data对象字段说明) {string} logo logo
 * @apiSuccess (data对象字段说明) {int} mfrom 支付最小额度
 * @apiSuccess (data对象字段说明) {int} mtp 支付最大额度
 * @apiSuccess (data对象字段说明) {int} priority 充值排序
 * @apiSuccess (data对象字段说明) {int} state 状态(0停用,1可用)
 * @apiSuccess (data对象字段说明) {int} created 添加时间
 * @apiSuccessExample {json} 响应结果
 * {
 * "clientMsg": "",
 * "code": 200,
 * "data": {
 * "Id": 3,
 * "Name": "qweqwe",
 * "Owner": "驱蚊器无",
 * "CardNumber": "12312312312312",
 * "BankAddress": "qweqwe",
 * "Remark": "hhhhhhhhhoooooooooooooo",
 * "Logo": "https://s3.ap-northeast-2.amazonaws.com/qpgame/ckyx_280135_axkd_.png",
 * "Mfrom": 1312,
 * "Mto": 123123,
 * "Priority": 2,
 * "State": 1,
 * "Created": 1560845632
 * },
 * "internalMsg": "",
 * "timeConsumed": 5232
 * }
 */
//获取银商充值 的银行卡信息
func (cthis *CardController) GetChargeCard() {
	ctx := cthis.ctx
	chargeCards := cthis.getChargeCards()
	utils.ResSuccJSON(&ctx, "", "", config.SUCCESSRES, chargeCards)
}

func (cthis *CardController) getChargeCards()interface{}{
	ctx := cthis.ctx
	engine := models.MyEngine[cthis.platform]
	chargeCards := new(xorm.SilverMerchantChargeCards)
	act , err :=engine.Where("state = ?",1).Desc("priority").Limit(1).Get(chargeCards)
	if err != nil {
		utils.ResFaiJSON2(&ctx, err.Error(), "查询失败")
		return nil
	}
	if !act {	// 没有记录
		utils.ResFaiJSON2(&ctx, "Oops", "暂无充值银行卡")
		return nil
	}
	return chargeCards
}

/**
 * @api {get} silverMerchant/api/auth/v1/getSilverCardInfo 银商获取银行卡信息接口
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: Singh</span><br/><br/>
 * 获取绑定银行卡<br>
 * 业务描述:银商获取银行卡绑定信息接口</br>
 * @apiVersion 1.0.0
 * @apiName     getSilverCardInfo
 * @apiGroup    silver_merchant

 * @apiErrorExample {json} 失败返回
 * {
 *     "clientMsg": "请尽快绑定银行卡",
 *     "code": 204,
 *     "internalMsg": "Oops",
 *     "timeConsumed": 9863
 * }
 * @apiSuccess (返回结果)  {int}     code           200
 * @apiSuccess (返回结果)  {string}  clientMsg      提示信息
 * @apiSuccess (返回结果)  {string}  internalMsg    提示信息
 * @apiSuccess (返回结果)  {float}   timeConsumed   后台耗时
 * @apiSuccess (返回结果)  {json}    data           返回数据
 * @apiSuccess (data对象字段说明) {int} id	自增ID
 * @apiSuccess (data对象字段说明) {string} MerchantId 银商UserID
 * @apiSuccess (data对象字段说明) {string} CardNumber 银行卡号
 * @apiSuccess (data对象字段说明) {int} Created 添加时间
 * @apiSuccess (data对象字段说明) {string} Address 银行卡地址
 * @apiSuccess (data对象字段说明) {string} BankName 银行名称
 * @apiSuccess (data对象字段说明) {string} Name 姓名
 * @apiSuccess (data对象字段说明) {int} Status 状态
 * @apiSuccess (data对象字段说明) {int} Updated 修改时间
 * @apiSuccess (data对象字段说明) {string} Remark 备注
 * @apiSuccessExample {json} 响应结果
 * {
 *     "clientMsg": "",
 *     "code": 200,
 *     "data": {
 *         "Id": 2,
 *         "MerchantId": 2,
 *         "CardNumber": "6222022213146561684",
 *         "Address": "兰州七里河土门墩支行",
 *         "Created": 1560252752,
 *         "BankName": "中国工商银行",
 *         "Name": "陈海冰",
 *         "Status": 1,
 *         "Updated": 1560252752,
 *         "Remark": ""
 *     },
 *     "internalMsg": "",
 *     "timeConsumed": 9924
 * }
 */
//银商获取银行卡绑定信息接口
func (cthis *CardController)GetSilverCardInfo(){
	ctx := cthis.ctx
	chargeBankCards := cthis.checkCard().(*xorm.SilverMerchantBankCards)
	if chargeBankCards.Id == 0 {
		utils.ResFaiJSON(&ctx, "Oops", "请尽快绑定银行卡", config.NOTGETDATA)
		return
	}
	utils.ResSuccJSON(&ctx, "", "", config.SUCCESSRES, chargeBankCards)
}

//获取银商绑定的银行卡
func (cthis *CardController)checkCard()interface{}{
	ctx := cthis.ctx
	sSmUserId := ctx.Values().GetString("silverMerchantUserId")
	merchantId, _ := strconv.Atoi(sSmUserId)
	engine := models.MyEngine[cthis.platform]
	chargeBankCards := new(xorm.SilverMerchantBankCards)
	engine.Where("merchant_id = ?",merchantId).And("status = 1").Desc("id").Limit(1).Get(chargeBankCards)
	return chargeBankCards
}

/**
 * @api {post} silverMerchant/api/auth/v1/bindCard 银商绑定银行卡接口
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: Singh</span><br/><br/>
 * 绑定银商银行卡<br>
 * 业务描述:银商绑定银行卡接口</br>
 * @apiVersion 1.0.0
 * @apiName     bindCard
 * @apiGroup    silver_merchant

 * @apiParam (客户端请求参数) {string} name     姓名
 * @apiParam (客户端请求参数) {string} bank_name     银行名称
 * @apiParam (客户端请求参数) {string} card_number     银行卡号
 * @apiParam (客户端请求参数) {string} address     银行地址
 * @apiError (请求失败返回)   {int}    code         错误代码
 * @apiError (请求失败返回)   {string} clientMsg    提示信息
 * @apiError (请求失败返回)   {string} internalMsg  错误代码
 * @apiError (请求失败返回)   {float}  timeConsumed 后台耗时

 * @apiErrorExample {json} 失败返回
 * {
 *     "clientMsg": "银行卡号不能为空",
 *     "code": 204,
 *     "internalMsg": "Oops",
 *     "timeConsumed": 3085
 * }
 * @apiSuccess (返回结果)  {int}     code           200
 * @apiSuccess (返回结果)  {string}  clientMsg      提示信息
 * @apiSuccess (返回结果)  {string}  internalMsg    提示信息
 * @apiSuccess (返回结果)  {float}   timeConsumed   后台耗时
 * @apiSuccess (返回结果)  {json}    data           返回数据
 * @apiSuccessExample {json} 响应结果
 * {
 *     "clientMsg": "绑定成功",
 *     "code": 200,
 *     "data": "",
 *     "internalMsg": "",
 *     "timeConsumed": 14884
 * }
 */
//银商绑定银行卡接口
func (cthis *CardController)BindCard(){
	ctx := cthis.ctx
	errMessage, pass := mvc.NewValidation(&ctx).
		NotNull("bank_name","银行名称不能为空").
		NotNull("card_number","银行卡号不能为空").
		NotNull("address","银行地址不能为空").
		StringLength("name","请输入你的名字,长度在2-5之间",2,5).
		Validate()
	if !pass {
		utils.ResFaiJSON(&ctx, "Oops", errMessage, config.NOTGETDATA)
		return
	}
	engine := models.MyEngine[cthis.platform]
	session := engine.NewSession()
	session.Begin()
	defer session.Close()

	sSmUserId := ctx.Values().GetString("silverMerchantUserId")
	merchantId, _ := strconv.Atoi(sSmUserId)

	form := utils.GetPostData(&ctx) //提交的post数据
	chargeBankCards := new(xorm.SilverMerchantBankCards)
	chargeBankCards.MerchantId = merchantId
	chargeBankCards.CardNumber = form.Get("card_number")
	chargeBankCards.Address = form.Get("address")
	chargeBankCards.Created = int(time.Now().Unix())
	chargeBankCards.BankName = form.Get("bank_name")
	chargeBankCards.Name = form.Get("name")
	chargeBankCards.Status = 1
	chargeBankCards.Updated = int(time.Now().Unix())
	chargeBankCards.Remark = form.Get("remark")

	_, err := engine.Insert(chargeBankCards)
	if err != nil {
		session.Rollback()
		utils.ResFaiJSON(&ctx, "绑定错误", "银行卡绑定失败", config.NOTGETDATA)
		return
	}
	sLogContent := "绑定银行卡：" + chargeBankCards.CardNumber
	_, err = services.SaveOperationLog(ctx, session, merchantId, sLogContent)
	if err != nil {
		session.Rollback()
		utils.ResFaiJSON(&ctx, "记录错误", "银行卡绑定失败", config.NOTGETDATA)
		return
	}
	err = session.Commit()
	if err != nil {
		utils.ResFaiJSON(&ctx, "提交错误", "银行卡绑定失败", config.NOTGETDATA)
		return
	}

	utils.ResSuccJSON(&ctx, "", "绑定成功", config.SUCCESSRES, "")
	return
}

/**
 * @api {post} silverMerchant/api/auth/v1/chargeSilver 银商充值额度接口
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: Singh</span><br/><br/>
 * 银商充值额度接口<br>
 * 业务描述:银商充值额度接口</br>
 * @apiVersion 1.0.0
 * @apiName     chargeSilver
 * @apiGroup    silver_merchant

 * @apiParam (客户端请求参数) {int} id     充值银行卡ID，获取充值银行卡接口返回的ID
 * @apiParam (客户端请求参数) {int} money  充值金额

 * @apiError (请求失败返回)   {int}    code         错误代码
 * @apiError (请求失败返回)   {string} clientMsg    提示信息
 * @apiError (请求失败返回)   {string} internalMsg  错误代码
 * @apiError (请求失败返回)   {float}  timeConsumed 后台耗时

 * @apiErrorExample {json} 失败返回
 * {
 *     "clientMsg": "该卡已停用，请重新获取",
 *     "code": 204,
 *     "internalMsg": "Oops",
 *     "timeConsumed": 9521
 * }
 * @apiSuccess (返回结果)  {int}     code           200
 * @apiSuccess (返回结果)  {string}  clientMsg      提示信息
 * @apiSuccess (返回结果)  {string}  internalMsg    提示信息
 * @apiSuccess (返回结果)  {float}   timeConsumed   后台耗时
 * @apiSuccess (返回结果)  {json}    data           返回数据
 * @apiSuccess (data对象字段说明) {int} Id	自增ID
 * @apiSuccess (data对象字段说明) {string} MerchantId silver_merchant_users表Id
 * @apiSuccess (data对象字段说明) {string} Amount 充值授权额度金额
 * @apiSuccess (data对象字段说明) {string} OrderId 充值订单
 * @apiSuccess (data对象字段说明) {string} CardNumber 充值卡号
 * @apiSuccess (data对象字段说明) {string} BankName 银行名称
 * @apiSuccess (data对象字段说明) {string} BankAddress 充值开户银行地址
 * @apiSuccess (data对象字段说明) {string} Created 添加时间
 * @apiSuccess (data对象字段说明) {string} State 状态，0待审核，1成功，2失败
 * @apiSuccess (data对象字段说明) {string} Updated 修改时间
 * @apiSuccess (data对象字段说明) {string} Ip 充值IP
 * @apiSuccess (data对象字段说明) {string} RealName 真实姓名
 * @apiSuccess (data对象字段说明) {string} BankChargeTime 银行转账时间
 * @apiSuccess (data对象字段说明) {string} Operator 操作者
 * @apiSuccess (data对象字段说明) {string} Remark 备注
 * @apiSuccess (data对象字段说明) {string} UpdatedLast 最后更新时间

 * @apiSuccessExample {json} 响应结果
 *  {
 *      "clientMsg": "",
 *      "code": 200,
 *      "data": {
 *          "Id": 6,
 *          "MerchantId": 2,
 *          "Amount": "10",
 *          "OrderId": "Y2019061217085470244f",
 *          "BankName": "中国工商银行",
 *          "BankAddress": "兰州七里河土门墩支行",
 *          "CardNumber": "270300 0401023533119",
 *          "Created": 1560330534,
 *          "State": 0,
 *          "Updated": 0,
 *          "Ip": "203.8.24.136",
 *          "RealName": "陈海冰",
 *          "BankChargeTime": 0,
 *          "Operator": "",
 *          "Remark": "",
 *          "UpdatedLast": "0001-01-01T00:00:00Z"
 *      },
 *      "internalMsg": "",
 *      "timeConsumed": 21193
 *  }
 * */
//银商充值额度接口
func (cthis *CardController)ChargeSilver(){
	ctx := cthis.ctx
	form := utils.GetPostData(&ctx) //提交的post数据

	chargeBankCards := cthis.checkCard().(*xorm.SilverMerchantBankCards)
	if chargeBankCards.Id == 0 {
		utils.ResFaiJSON(&ctx, "Oops", "请先绑定银行卡", config.NOTGETDATA)
		return
	}


	errMessage, pass := mvc.NewValidation(&ctx).
		NotNull("id","银行卡ID不能为空!").
		NotNull("money","充值额度不能为空!").
		IsNumeric("money","充值额度必须为数字类型!").
		Validate()
	if !pass {
		utils.ResFaiJSON(&ctx, "Oops", errMessage, config.NOTGETDATA)
		return
	}


	engine := models.MyEngine[cthis.platform]
	chargeCards := new(xorm.SilverMerchantChargeCards)
	bol,err :=engine.ID(form.Get("id")).And("state=?",1).Get(chargeCards)
	if err != nil || bol == false {
		utils.ResFaiJSON(&ctx, "Oops", "该卡已停用，请重新获取", config.NOTGETDATA)
		return
	}

	//检查被充值银行卡是否存在，状态是否正常
	tempMoney := form.Get("money")
	money, transErr := strconv.Atoi(tempMoney)
	if transErr != nil {
		money = 0
	}

	if chargeCards.Mfrom > money || chargeCards.Mto < money {
		utils.ResFaiJSON(&ctx, "Oops", "充值额度必须在【" + strconv.Itoa(chargeCards.Mfrom) + "~" + strconv.Itoa(chargeCards.Mto) + "】范围内", config.NOTGETDATA)
		return
	}

	var configs = adminModels.Configs{}
	sSmUserId := ctx.Values().GetString("silverMerchantUserId")
	merchantId, _ := strconv.Atoi(sSmUserId)
	chargeUsers,userExits := GetSilverInfo(cthis.platform,sSmUserId)
	if !userExits {
		utils.ResFaiJSON(&ctx, "Oops", "充值错误", config.NOTGETDATA)
		return
	}

	adminSilverConfig := configs.GetSilverMerchant(&ctx)
	cashPledge := adminSilverConfig["cash_pledge"]			//后台配置的银商押金金额

	silverCashType := chargeUsers.MerchantCashPledge 			//银商当前押金额度
	silverCashDec, _ := decimal.NewFromString(silverCashType)
	silverCash,_ := silverCashDec.Float64()
	var msg = ""
	if silverCash > 0 {
		msg = "温馨提示：请把订单号复制到网银的备注中，否则无法正常入款"
	}else{
		money = money + (int(cashPledge))		//如当前押金不足，则补足押金
		msg = "温馨提示：请把订单号复制到网银的备注中，否则无法正常入款。\n说明：该笔订单，其中 " + strconv.Itoa(int(cashPledge)) + "元 是作为保证金充值，其余为可用额度。"

	}
	session := engine.NewSession()
	session.Begin()
	defer session.Close()

	//生成订单号，返回创建订单时间，订单号，被充值银行卡信息
	orderId := utils.CreationOrder("Y", sSmUserId)
	chargeRecord := new(xorm.SilverMerchantChargeRecords)
	chargeRecord.MerchantId = merchantId
	chargeRecord.Amount = strconv.Itoa(money)
	chargeRecord.OrderId = orderId
	chargeRecord.CardNumber = chargeCards.CardNumber
	chargeRecord.BankAddress = chargeCards.BankAddress
	chargeRecord.BankName = chargeCards.Name
	chargeRecord.Created = int(time.Now().Unix())
	chargeRecord.State = 0
	chargeRecord.Ip = utils.GetIp(ctx.Request())
	chargeRecord.RealName = chargeCards.Owner

	_, err = engine.Insert(chargeRecord)
	if err != nil {
		session.Rollback()
		utils.ResFaiJSON(&ctx, "系统错误", "生成订单失败", config.NOTGETDATA)
		return
	}

	sLogContent := "生成充值订单，订单号：" + chargeRecord.OrderId
	_, err = services.SaveOperationLog(ctx, session, chargeRecord.MerchantId, sLogContent)
	if err != nil {
		session.Rollback()
		utils.ResFaiJSON(&ctx, "记录错误", "生成订单失败", config.NOTGETDATA)
		return
	}
	err = session.Commit()
	if err != nil {
		utils.ResFaiJSON(&ctx, "提交错误", "银行卡绑定失败", config.NOTGETDATA)
		return
	}

	utils.ResSuccJSON(&ctx, "", msg, config.SUCCESSRES, chargeRecord)
}


/**
 * @api {get} silverMerchant/api/auth/v1/getChargeSilverList 银商充值记录查询接口
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: Singh</span><br/><br/>
 * 银商充值记录查询接口<br>
 * 业务描述:银商充值记录查询接口</br>
 * @apiVersion 1.0.0
 * @apiName     GetChargeSilverList
 * @apiGroup    silver_merchant

 * @apiParam (客户端请求参数) {string} from_time	查询起始时间
 * @apiParam (客户端请求参数) {string} to_time		查询结束时间
 * @apiParam (客户端请求参数) {int} page_size		分页大小
 * @apiParam (客户端请求参数) {int} page			当前页

 * @apiError (请求失败返回)   {int}    code         错误代码
 * @apiError (请求失败返回)   {string} clientMsg    提示信息
 * @apiError (请求失败返回)   {string} internalMsg  错误代码
 * @apiError (请求失败返回)   {float}  timeConsumed 后台耗时

 * @apiErrorExample {json} 失败返回
 * {
 *     "clientMsg": "暂无记录",
 *     "code": 204,
 *     "internalMsg": "Oops",
 *     "timeConsumed": 4361
 * }
 * @apiSuccess (返回结果)  {int}     code           200
 * @apiSuccess (返回结果)  {string}  clientMsg      提示信息
 * @apiSuccess (返回结果)  {string}  internalMsg    提示信息
 * @apiSuccess (返回结果)  {float}   timeConsumed   后台耗时
 * @apiSuccess (返回结果)  {json}    data           返回数据


 * @apiSuccess (data对象字段说明) {array}   list   记录对象数组
 * @apiSuccess (data对象字段说明) {int}   page   启始返回位置
 * @apiSuccess (data对象字段说明) {int}   page_size   返回条数
 * @apiSuccess (data对象字段说明) {int} total   总数

 * @apiSuccess (data-list元素对象字段说明) {int} Id	自增ID
 * @apiSuccess (data-list元素对象字段说明) {int} MerchantId silver_merchant_users表Id
 * @apiSuccess (data-list元素对象字段说明) {string} Amount 充值授权额度金额
 * @apiSuccess (data-list元素对象字段说明) {string} OrderId 充值订单
 * @apiSuccess (data-list元素对象字段说明) {string} CardNumber 充值卡号
 * @apiSuccess (data-list元素对象字段说明) {string} BankName 银行名称
 * @apiSuccess (data-list元素对象字段说明) {string} BankAddress 充值开户银行地址
 * @apiSuccess (data-list元素对象字段说明) {int} Created 添加时间
 * @apiSuccess (data-list元素对象字段说明) {int} State 状态，0待审核，1成功，2失败
 * @apiSuccess (data-list元素对象字段说明) {int} Updated 修改时间
 * @apiSuccess (data-list元素对象字段说明) {string} Ip 充值IP
 * @apiSuccess (data-list元素对象字段说明) {string} RealName 真实姓名
 * @apiSuccess (data-list元素对象字段说明) {int} BankChargeTime 银行转账时间
 * @apiSuccess (data-list元素对象字段说明) {string} Operator 操作者
 * @apiSuccess (data-list元素对象字段说明) {string} Remark 备注
 * @apiSuccess (data-list元素对象字段说明) {string} UpdatedLast 最后更新时间

 * @apiSuccessExample {json} 响应结果
 * {
 *     "clientMsg": "",
 *     "code": 200,
 *     "data": {
 *         "list": [
 *             {
 *                 "amount": 10,
 *                 "bank_address": "兰州七里河土门墩支行",
 *                 "bank_charge_time": 0,
 *                 "bank_name": "中国工商银行",
 *                 "card_number": "270300 0401023533119",
 *                 "created": 1560322734,
 *                 "id": 2,
 *                 "ip": "203.8.24.136",
 *                 "merchant_id": 2,
 *                 "operator": "",
 *                 "order_id": "YS20190612145854d3497",
 *                 "real_name": "陈海冰",
 *                 "remark": "",
 *                 "state": 0,
 *                 "updated": 0,
 *                 "updated_last": "0000-00-00 00:00:00"
 *             }
 *         ],
 *         "page": "0",
 *         "page_size": "50",
 *         "total": 4
 *     },
 *     "internalMsg": "",
 *     "timeConsumed": 4597
 * }
 *
 **/
//银商充值记录查询接口
func (cthis *CardController)GetChargeSilverList(){
	ctx := cthis.ctx
	sSmUserId := ctx.Values().GetString("silverMerchantUserId")
	errMessage, pass := mvc.NewValidation(&ctx).
		NotNull("from_time","from_time不能为空!").
		NotNull("to_time","to_time不能为空!").
		NotNull("page","page不能为空!").
		NotNull("page_size","page_size不能为空!").
		Validate()
	if !pass {
		utils.ResFaiJSON(&ctx, "Oops", errMessage, config.NOTGETDATA)
		return
	}

	form := utils.GetPostData(&ctx) //提交的post数据
	fromTime := form.Get("from_time")
	toTime := form.Get("to_time")
	page := form.Get("page")
	page_size := form.Get("page_size")

	sqlCount := "SELECT count(*) as total FROM `silver_merchant_charge_records` WHERE merchant_id = " + sSmUserId +" AND created BETWEEN UNIX_TIMESTAMP('" + fromTime + "') and UNIX_TIMESTAMP('" + toTime + " 23:59:59" + "');"
	response, err := utils.Query(cthis.platform, sqlCount)
	total := response[0]["total"]
	if err != nil || total == 0 {
		utils.ResFaiJSON(&ctx, "Oops", "暂无记录", config.NOTGETDATA)
		return
	}
	sqlstr := "SELECT * FROM `silver_merchant_charge_records` WHERE merchant_id = " + sSmUserId +" AND created BETWEEN UNIX_TIMESTAMP('" + fromTime + "') and UNIX_TIMESTAMP('" + toTime + " 23:59:59" + "') ORDER BY id DESC LIMIT " + page_size + " OFFSET " + page + ";"
	response, err = utils.Query(cthis.platform, sqlstr) 
	if err != nil || len(response) == 0 {
		utils.ResFaiJSON(&ctx, "Oops", "暂无记录", config.NOTGETDATA)
		return
	}
	res := map[string]interface{}{
		"list" : response,
		"total" : total,
		"page" : page,
		"page_size" : page_size,
	}
	utils.ResSuccJSON(&ctx, "", "", config.SUCCESSRES, res)
}

//检查被充值银行卡状态是否正常
func CheckChargeCard(platform string, id string) (xorm.ChargeCards, bool) {
	var chargeCard = xorm.ChargeCards{}
	exist, _ := models.MyEngine[platform].Where("id=?", id).And("addr_type=1").And("state=1") .Get(&chargeCard)
	return chargeCard, exist
}

//获取银商用户详情
func GetSilverInfo(platform string, id string) (xorm.SilverMerchantUsers, bool) {
	var chargeUsers = xorm.SilverMerchantUsers{}
	exist, _ := models.MyEngine[platform].Where("id=?", id).Get(&chargeUsers)
	return chargeUsers, exist
}

