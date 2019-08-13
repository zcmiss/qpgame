package silverMerchant

import (
	libXorm "github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	"github.com/shopspring/decimal"
	"html"
	"qpgame/app/fund"
	"qpgame/common/utils"
	"qpgame/config"
	"qpgame/models"
	"qpgame/models/xorm"
	"strconv"
	"time"
)

type MemberPayController struct {
	platform        string
	ctx             iris.Context
	engine          *libXorm.Engine
	dMerchantAmount decimal.Decimal
}

//构造函数
func NewMemberPayController(ctx iris.Context) *MemberPayController {
	obj := new(MemberPayController)
	obj.platform = ctx.Params().Get("platform")
	obj.ctx = ctx
	return obj
}

// 检查银商是否满足给会员充值的条件
func (cthis *MemberPayController) checkMerchantAmount(merchantId int, dChargeToMemberAmount decimal.Decimal) (*xorm.SilverMerchantUsers, bool) {
	ctx := cthis.ctx
	engine := cthis.engine
	var smUser = new(xorm.SilverMerchantUsers)
	has, err := engine.Id(merchantId).Cols("usable_amount", "is_destroy", "status").Get(smUser)
	if err != nil || has == false {
		utils.ResFaiJSON(&ctx, "1906131120", "银商账号不存在", config.NOTGETDATA)
		return nil, false
	}
	if smUser.IsDestroy == 1 || smUser.Status == 0 {
		utils.ResFaiJSON(&ctx, "1906131121", "银商账号已注销或被锁定", config.NOTGETDATA)
		return nil, false
	}
	sUsableAmount := smUser.UsableAmount
	dUsableAmount, _ := decimal.NewFromString(sUsableAmount)
	if dUsableAmount.LessThan(dChargeToMemberAmount) {
		utils.ResFaiJSON(&ctx, "1906131127", "充值失败，银商充值额度不足", config.NOTGETDATA)
		return nil, false
	}
	cthis.dMerchantAmount = dUsableAmount
	return smUser, true
}

/**
 * @api {post} silverMerchant/api/auth/v1/pay 会员充值
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: aTian</span><br/><br/>
 * 会员充值接口<br>
 * 业务描述：创建用户充值记录，资金流水，更新商户账户资金，操作日志处理等</br>
 * @apiVersion 1.0.0
 * @apiName     member_pay
 * @apiGroup    silver_merchant
 * @apiPermission PC客户端
 * @apiParam (客户端请求参数) {string} member_user_id 充值会员ID
 * @apiParam (客户端请求参数) {string} amount         充值金额，正值
 * @apiParam (客户端请求参数) {string} msg            资金流描述说明
 * @apiError (请求失败返回)   {int}    code          错误代码
 * @apiError (请求失败返回)   {string} clientMsg     提示信息
 * @apiError (请求失败返回)   {string} internalMsg   错误代码
 * @apiError (请求失败返回)   {float}  timeConsumed  后台耗时

 * @apiErrorExample {json} 失败返回
 * {
 *   "clientMsg": "充值金额不能小于0",
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
 *     "clientMsg": "充值成功",
 *     "code": 200,
 *     "data": "",
 *     "internalMsg": "",
 *     "timeConsumed": 11958575
 * }
 */
func (cthis *MemberPayController) PayPlayer() {
	ctx := cthis.ctx
	postData := utils.GetPostData(&ctx)
	if !utils.ValidRequiredPostData(ctx, postData, []string{"member_user_id", "amount"}) {
		return
	}
	sMemUserId := postData.Get("member_user_id")
	sMsg := postData.Get("msg")
	iMemUserId, err := strconv.Atoi(sMemUserId)
	if err != nil || iMemUserId <= 0 {
		utils.ResFaiJSON(&ctx, "1906131045", "会员不存在", config.NOTGETDATA)
		return
	}
	platform := cthis.platform
	engine := models.MyEngine[platform]
	sAmount := postData.Get("amount")
	decimal.DivisionPrecision = 3
	dAmount, _ := decimal.NewFromString(sAmount)
	if dAmount.LessThanOrEqual(decimal.New(0, 0)) {
		utils.ResFaiJSON(&ctx, "1906131036", "充值金额不能小于0", config.NOTGETDATA)
		return
	}
	dPayUnit := decimal.New(100, 0)
	if dAmount.LessThan(dPayUnit) {
		utils.ResFaiJSON(&ctx, "1906170900", "至少100元以上的整数", config.NOTGETDATA)
		return
	}
	var has bool
	// 查询用户是否存在
	var memUser = new(xorm.Users)
	has, err = engine.ID(iMemUserId).Cols("user_name").Get(memUser)
	if err != nil || has == false {
		utils.ResFaiJSON(&ctx, "1906131050", "会员不存在", config.NOTGETDATA)
		return
	}
	sSmUserId := ctx.Values().GetString("silverMerchantUserId")
	iSmUserId, _ := strconv.Atoi(sSmUserId)
	cthis.engine = engine
	_, smOk := cthis.checkMerchantAmount(iSmUserId, dAmount)
	if smOk == false {
		return
	}
	iNow := utils.GetNowTime()
	var sOrderId string
	sOrderId = utils.TimestampToDateStr(int64(iNow), "060102150405") + utils.RandString(4, 4)
	chargeBankCards := NewSilverCardController(ctx).checkCard().(*xorm.SilverMerchantBankCards)

	session := engine.NewSession()
	session.Begin()
	defer session.Close()
	dAmount64, _ := dAmount.Float64()
	info := map[string]interface{}{
		"merchant_id":    sSmUserId,  // 银商编号，字符串
		"member_user_id": iMemUserId, // 会员ID
		"order_id":       sOrderId,   // 订单ID
		"type_id":        2,          // 类型,数字: 1 额度充值，2 给棋牌用户充值
		"amount":         dAmount64,  // 操作说明，float64数字
		"msg":            "银商用户充值",   // 操作说明，字符串
		"transaction":    session,    // 外部事务，*goxorm.Session类型
		"remark":         sMsg,
	}

	var callback fund.BalanceUpdateCallback = func(db *libXorm.Session, args ...interface{}) (interface{}, error) {
		//生成一条充值订单，调用用户充值
		userChargeRecord := xorm.ChargeRecords{
			UserId:         iMemUserId,
			Amount:         dAmount.String(),
			OrderId:        sOrderId,
			ChargeTypeId:   3,
			CardNumber:     chargeBankCards.CardNumber,
			BankAddress:    chargeBankCards.Address,
			Created:        int(time.Now().Unix()),
			State:          1,
			Updated:        int(time.Now().Unix()),
			ChargeTypeInfo: "银商充值支付",
			Ip:             utils.GetIp(ctx.Request()),
			RealName:       chargeBankCards.Name,
			BankChargeTime: int(time.Now().Unix()),
			CredentialId:   0,
			IsTppay:        0,
			ChargeCardId:   chargeBankCards.Id,
		}
		res, err := session.Insert(&userChargeRecord)
		if err != nil {
			return res, err
		}
		infoUser := map[string]interface{}{
			"user_id":     iMemUserId,        //用户ID
			"type_id":     config.FUNDCHARGE, //交易类型
			"amount":      dAmount64,         //交易金额
			"order_id":    sOrderId,          //订单号
			"msg":         "用户充值",            //交易信息
			"transaction": session,
		}
		userResult := fund.NewUserFundChange(platform).BalanceUpdate(infoUser, nil)
		if userResult["status"] == 1 {
			err := session.Commit()
			return nil, err
		} else {
			session.Rollback()
		}
		return nil, nil
	}
	result := fund.NewMerchantFundChange(platform).BalanceUpdate(info, callback)

	if result["status"] == 1 {
		utils.ResSuccJSON(&ctx, "", "充值成功", config.SUCCESSRES, "")
		return
	} else {
		utils.ResFaiJSON(&ctx, "1906131705", "充值失败，请重试", config.NOTGETDATA)
		return
	}
}

/**
 * @api {post} silverMerchant/api/auth/v1/searchMember 会员搜索接口
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: Singh</span><br/><br/>
 * 会员充值接口<br>
 * 业务描述：通过user_name 搜索用户id</br>
 * @apiVersion 1.0.0
 * @apiName     searchMember
 * @apiGroup    silver_merchant
 * @apiPermission PC客户端
 * @apiParam (客户端请求参数) {string} user_name 会员用户名
 * @apiError (请求失败返回)   {int}    code          错误代码
 * @apiError (请求失败返回)   {string} clientMsg     提示信息
 * @apiError (请求失败返回)   {string} internalMsg   错误代码
 * @apiError (请求失败返回)   {float}  timeConsumed  后台耗时

 * @apiErrorExample {json} 失败返回
 * {
 *   "clientMsg": "会员不存在",
 *   "code": 204,
 *   "internalMsg": "Oops",
 *   "timeConsumed": 34
 * }
 *
 * @apiSuccess (返回结果)  {int}     code           200
 * @apiSuccess (返回结果)  {string}  clientMsg      提示信息
 * @apiSuccess (返回结果)  {string}  internalMsg    提示信息
 * @apiSuccess (返回结果)  {float}   timeConsumed   后台耗时
 * @apiSuccess (返回结果)  {json}    data           返回数据
 * @apiSuccess (data对象字段说明) {int}   id   用户ID
 * @apiSuccess (data对象字段说明) {string} user_name   会员用户名
 * @apiSuccessExample {json} 响应结果
 *  {
 *     "clientMsg": "查询成功",
 *     "code": 200,
 *     "data": {
 *         "id": 111,
 *         "user_name": "yk111"
 *     },
 *     "internalMsg": "",
 *     "timeConsumed": 5938
 * }
 */
//会员搜索接口(用户名，必须要确保搜索唯一，避免充错到其他会员)
func (cthis *MemberPayController) CheckUserName() {
	ctx := cthis.ctx
	postData := utils.GetPostData(&ctx)
	userName := postData.Get("user_name")
	user, exits := checkUser(cthis.platform, userName)
	if !exits {
		utils.ResFaiJSON(&ctx, "Oops", "会员不存在", config.NOTGETDATA)
		return
	}
	res := map[string]interface{}{
		"id":        user.Id,
		"user_name": user.UserName,
	}
	utils.ResSuccJSON(&ctx, "", "查询成功", config.SUCCESSRES, res)
}

//检查用户是否存在
func checkUser(platform string, userName string) (xorm.Users, bool) {
	var Users = xorm.Users{}
	exist, _ := models.MyEngine[platform].Where("user_name=?", userName).Get(&Users)
	return Users, exist
}

/**
 * @api {get} silverMerchant/api/auth/v1/payRecords 银商会员充值记录查询
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: aTian</span><br/><br/>
 * 银商会员充值记录查询<br>
 * 业务描述：银商会员充值记录查询</br>
 * @apiVersion 1.0.0
 * @apiName     pay_records
 * @apiGroup    silver_merchant
 * @apiPermission PC客户端
 * @apiParam (客户端请求参数) {string} from_time    时间范围，创建起始时间：（2012-12-31）
 * @apiParam (客户端请求参数) {string} to_time      时间范围，创建结束时间：（2012-12-31）
 * @apiParam (客户端请求参数) {string} type         1.额度充值,2.会员充值扣款
 * @apiParam (客户端请求参数) {string} order_id     订单号
 * @apiParam (客户端请求参数) {string} min_amount   充值金额范围，小于等于
 * @apiParam (客户端请求参数) {string} max_amount   充值金额范围，大于等于
 * @apiParam (客户端请求参数) {string} msg          资金流描述说明
 * @apiParam (客户端请求参数) {string} member_name  会员名称
 * @apiParam (客户端请求参数) {string} page         页码 (默认1)
 * @apiParam (客户端请求参数) {string} size         每页显示行数 (默认20)
 * @apiParam (客户端请求参数) {string} order_by     排序字段[id(默认), created, amount, balance]
 * @apiParam (客户端请求参数) {string} sort         排序方式[ASC, DESC(默认)]
 * @apiError (请求失败返回)   {int}    code         错误代码
 * @apiError (请求失败返回)   {string} clientMsg    提示信息
 * @apiError (请求失败返回)   {string} internalMsg  错误代码
 * @apiError (请求失败返回)   {float}  timeConsumed 后台耗时

 * @apiErrorExample {json} 失败返回
 * {
 *   "clientMsg": "获取登录日志失败",
 *   "code": 204,
 *   "internalMsg": "1906111353",
 *   "timeConsumed": 34
 * }
 *
 * @apiSuccess (返回结果)  {int}     code           200
 * @apiSuccess (返回结果)  {string}  clientMsg      提示信息
 * @apiSuccess (返回结果)  {string}  internalMsg    提示信息
 * @apiSuccess (返回结果)  {float}   timeConsumed   后台耗时
 * @apiSuccess (返回结果)  {json}    data           返回数据
 * @apiSuccess (data对象字段说明) {string} order_by  排序字段
 * @apiSuccess (data对象字段说明) {int} page 当前页码
 * @apiSuccess (data对象字段说明) {int} size 每页显示行数
 * @apiSuccess (data对象字段说明) {string} sort 排序方式
 * @apiSuccess (data对象字段说明) {int} total 数据总行数
 * @apiSuccess (data对象字段说明) {array} list 充值记录列表
 * @apiSuccess (data对象子字段 list 说明) {string} amount 充值金额
 * @apiSuccess (data对象子字段 list 说明) {string} balance 变化后的余额
 * @apiSuccess (data对象子字段 list 说明) {string} charged_amount 资金变动之后的额度余额
 * @apiSuccess (data对象子字段 list 说明) {string} charged_amount_old 更新之前的额度金额
 * @apiSuccess (data对象子字段 list 说明) {string} created 创建时间
 * @apiSuccess (data对象子字段 list 说明) {string} id 充值记录ID
 * @apiSuccess (data对象子字段 list 说明) {string} member_user_id 会员ID,只有在type=2的情况,1的话默认为0
 * @apiSuccess (data对象子字段 list 说明) {string} merchant_id 银商ID
 * @apiSuccess (data对象子字段 list 说明) {string} msg 资金流描述说明
 * @apiSuccess (data对象子字段 list 说明) {string} order_id 订单号
 * @apiSuccess (data对象子字段 list 说明) {string} type 1.额度充值,2.会员充值扣款,3.额度充值赠送,4.押金
 * @apiSuccess (data对象子字段 list 说明) {string} user_name 会员名
 * @apiSuccessExample {json} 响应结果
 * {
 *     "clientMsg": "获取会员充值记录成功",
 *     "code": 200,
 *     "data": {
 *         "list": [
 *             {
 *                 "amount": "-100.000",
 *                 "balance": "99900.000",
 *                 "charged_amount": "99900.000",
 *                 "charged_amount_old": "100000.000",
 *                 "created": "2019-06-13 17:08:20",
 *                 "id": "4",
 *                 "member_user_id": "1",
 *                 "merchant_id": "1",
 *                 "msg": "hello",
 *                 "order_id": "190613170820KDBM",
 *                 "type": "2",
 *                 "user_name": "1559547860MlDoCE"
 *             }
 *         ],
 *         "order_by": "id",
 *         "page": 1,
 *         "size": 20,
 *         "sort": "DESC",
 *         "total": 1
 *     },
 *     "internalMsg": "",
 *     "timeConsumed": 9198139
 * }
 */
func (cthis *MemberPayController) PayForMemberRecords() {
	ctx := cthis.ctx
	sSmUserId := ctx.Values().GetString("silverMerchantUserId")
	postData := utils.GetPostData(&ctx)

	sFromDate := postData.Get("from_time")
	sToDate := postData.Get("to_time")
	var sFromTimestamp = ""
	var sToTimestamp = ""
	if sFromDate != "" {
		sFromDate = sFromDate + " 00:00:00"
		iFromTimestamp := utils.GetInt64FromTime(sFromDate)
		sFromTimestamp = strconv.Itoa(int(iFromTimestamp))
	}
	if sToDate != "" {
		sToDate = sToDate + " 23:59:59"
		iToTimestamp := utils.GetInt64FromTime(sToDate)
		sToTimestamp = strconv.Itoa(int(iToTimestamp))
	}
	sType := postData.Get("type")
	if sType == "1" {
		sType = "1,3,4"
	}
	filterArr := []map[string]string{
		{"prefix": "smcf", "key": "merchant_id", "val": sSmUserId},
		{"prefix": "smcf", "key": "created", "val": sFromTimestamp, "condition": "gt"},
		{"prefix": "smcf", "key": "created", "val": sToTimestamp, "condition": "lt"},
		{"prefix": "smcf", "key": "type", "val": sType, "condition": "in"},
		{"prefix": "smcf", "key": "order_id", "val": postData.Get("order_id")},
		{"prefix": "smcf", "key": "amount", "val": postData.Get("min_amount"), "condition": "ge"},
		{"prefix": "smcf", "key": "amount", "val": postData.Get("max_amount"), "condition": "le"},
		{"prefix": "smcf", "key": "msg", "val": postData.Get("msg"), "condition": "like_ab"},
		{"prefix": "u", "key": "user_name", "val": postData.Get("member_name"), "condition": "like_ab"},
	}
	where := utils.BuildWhere(filterArr)
	countSql := "SELECT COUNT(*) cnt FROM silver_merchant_capital_flows smcf LEFT JOIN users u ON smcf.member_user_id=u.id " + where
	engine := models.MyEngine[cthis.platform]
	cntRes, cntErr := engine.SQL(countSql).QueryString()
	if cntErr != nil {
		utils.ResFaiJSON(&ctx, "1906131846", "获取登录日志失败", config.NOTGETDATA)
		return
	}

	var total int
	if len(cntRes) == 1 {
		sTotal := cntRes[0]["cnt"]
		total, _ = strconv.Atoi(sTotal)
	}
	sPage := postData.Get("page")
	iPage, _ := strconv.Atoi(sPage)
	sSize := postData.Get("size")
	iSize, _ := strconv.Atoi(sSize)
	sOrderBy := postData.Get("order_by")
	sSort := postData.Get("sort")
	if sOrderBy == "" {
		sOrderBy = "id"
	}
	if sSort == "" {
		sSort = "DESC"
	}
	iPage, iSize = utils.GetDefaultPageInfo(iPage, iSize)
	if total > 0 {
		limit := utils.BuildLimit(iPage, iSize)
		orderBy := " ORDER BY smcf.`" + html.EscapeString(sOrderBy) + "` " + html.EscapeString(sSort)
		sql := "SELECT smcf.*,u.user_name FROM silver_merchant_capital_flows smcf LEFT JOIN users u ON smcf.member_user_id=u.id " + where + orderBy + limit
		qryList, queryErr := engine.SQL(sql).QueryString()
		if queryErr != nil {
			utils.ResFaiJSON(&ctx, "1906131905", "获取会员充值记录失败", config.NOTGETDATA)
			return
		}
		var list interface{}
		if len(qryList) > 0 {
			for k, row := range qryList {
				if sTime, fExist := row["created"]; fExist {
					iTime, _ := strconv.Atoi(sTime)
					qryList[k]["created"] = utils.TimestampToDateStr(int64(iTime), "1")
				}
			}
			list = qryList
		} else {
			list = make([]map[string]string, 0)
		}

		data := map[string]interface{}{
			"list":     list,
			"page":     iPage,
			"size":     iSize,
			"total":    total,
			"order_by": sOrderBy,
			"sort":     sSort,
		}
		utils.ResSuccJSON(&ctx, "", "获取会员充值记录成功", config.SUCCESSRES, data)
	} else {
		utils.ResSuccJSON(&ctx, "", "获取会员充值记录成功", config.SUCCESSRES, map[string]interface{}{
			"list":     make([]map[string]string, 0),
			"page":     iPage,
			"size":     iSize,
			"total":    total,
			"order_by": sOrderBy,
			"sort":     sSort,
		})
	}
}
