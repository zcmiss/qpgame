package frontEndControllers

import (
	"github.com/kataras/iris"
	"github.com/shopspring/decimal"
	"qpgame/app/fund"
	"qpgame/common/utils"
	"qpgame/config"
	"qpgame/models"
	"qpgame/models/xorm"
	"strconv"
	"time"
)

type ProxyController struct {
	platform string
	ctx      iris.Context
}

//构造函数
func NewProxyController(ctx iris.Context) *ProxyController {
	obj := new(ProxyController)
	obj.platform = ctx.Params().Get("platform")
	obj.ctx = ctx
	return obj
}

/**
* @api {get} api/auth/v1/promotion 推广赚钱接口
* @apiDescription
* <span style="color:lightcoral;">接口负责人:wenzhen</span><br/><br/>
* 推广赚钱<br>
* 业务描述:前端进入推广赚钱需要调用此接口</br>
* @apiVersion 1.0.0
* @apiName     promotion
* @apiGroup    proxy_module
 * @apiPermission iso,android客户端
* @apiHeader (客户端请求头参数) {string} Authorization 		Bearer + 用户登录获得的token
* @apiParam (客户端请求参数) {int}    	recordId    	   	洗码记录编号
* @apiError (请求失败返回) {int}      	code            	错误代码
* @apiError (请求失败返回) {string}   	clientMsg       	提示信息
* @apiError (请求失败返回) {string}  	 internalMsg     	错误代码
* @apiError (请求失败返回) {float}    	time_consumed   	后台耗时
*
* @apiErrorExample {json} 失败返回
*{
*    "clientMsg": "获取失败",
*    "code": 204,
*    "internalMsg": "",
*    "time_consumed": 182507
*} *
*
* @apiSuccess (返回结果)  {int}      code            200
* @apiSuccess (返回结果)  {string}   clientMsg       提示信息
* @apiSuccess (返回结果)  {string}   internalMsg     提示信息
* @apiSuccess (返回结果)  {json}  	  data            返回数据
* @apiSuccess (返回结果)  {float}   time_consumed    后台耗时
* @apiSuccess (data字段说明)  {int}   groupSize    团队成员数量
* @apiSuccess (data字段说明)  {float}   notCashed    可领取佣金金额
* @apiSuccess (data字段说明)  {float}   proxyAmount    累计获取佣金
* @apiSuccess (data字段说明)  {string}   proxyLevel    当前代理等级
* @apiSuccess (data字段说明)  {list}   list    			佣金记录表
* @apiSuccess (list字段说明)  {int}   Created    			结算时间的时间戳
* @apiSuccess (list字段说明)  {string}   CreatedStr    			用于查询详情
* @apiSuccess (list字段说明)  {int}   ProxyType    			代理类型2棋牌5真人
* @apiSuccess (list字段说明)  {string}   TotalCommission    			所得佣金
* @apiSuccess (list字段说明)  {string}   BetAmount    			个人贡献业绩
* @apiSuccess (list字段说明)  {string}   TotalAmount    			团队总贡献业绩
* @apiSuccess (list字段说明)  {int}   Contributions    			贡献人数
* @apiSuccess (list字段说明)  {int}   Status    			状态0未领取1已完成
* @apiSuccess (list字段说明)  {int}   UserId    			用户编号
* @apiSuccess (list字段说明)  {int}   ParentId    			上级代理用户编号

* @apiSuccessExample {json} 响应结果
*
{
    "clientMsg": "数据获取成功",
    "code": 200,
    "data": {
        "groupSize": 4,
        "list": [
            {
                "Id": 17,
                "UserId": 28,
                "ParentId": 0,
                "BetAmount": "93076.000",
                "TotalAmount": "93176.000",
                "Created": 1556199643,
                "Contributions": 1,
                "ProxyType": 2,
                "ProxyLevel": 3,
                "ProxyLevelRate": "0.0070",
                "ProxyLevelName": "代理级",
                "Commission": "651.532",
                "TotalCommission": "651.732",
                "CreatedStr": "20190424",
                "States": 0
            },
            {
                "Id": 15,
                "UserId": 28,
                "ParentId": 0,
                "BetAmount": "500198.000",
                "TotalAmount": "3003251.000",
                "Created": 1556199377,
                "Contributions": 2,
                "ProxyType": 2,
                "ProxyLevel": 10,
                "ProxyLevelRate": "0.0200",
                "ProxyLevelName": "至尊股东级",
                "Commission": "10003.960",
                "TotalCommission": "22016.188",
                "CreatedStr": "20190426",
                "States": 0
            }
        ],
        "notCashed": 22667.92,
        "proxyAmount": 0,
        "proxyLevel": "代理级"
    },
    "internalMsg": "success",
    "timeConsumed": 17952
}
*
*/
func (cthis *ProxyController) Promotion() {
	ctx := cthis.ctx
	userid, _ := (ctx.Values().GetInt("userid"))
	session := models.MyEngine[cthis.platform]
	queryTime := time.Now().AddDate(0, 0, -1).Format("20060102")
	var beans []xorm.ProxyCommissions
	err := session.Where("user_id = ?", userid).Desc("created").Find(&beans)
	sqlstr := "SELECT (SELECT proxy_amount FROM accounts a WHERE a.`user_id`= ?) proxy_amount, (SELECT COUNT(*) FROM users u  WHERE u.parent_id= ?) groupsize FROM DUAL"
	rows, err := session.QueryString(sqlstr, userid, userid)
	if err != nil {
		utils.ResFaiJSON2(&ctx, err.Error(), "获取失败")
		return
	}
	res := rows[0]
	resMap := make(map[string]interface{})
	resMap["groupSize"], _ = strconv.Atoi(res["groupsize"])
	proxyAmount, _ := strconv.ParseFloat(res["proxy_amount"], 64)
	resMap["proxyAmount"] = proxyAmount
	resMap["list"] = beans
	notCashed := decimal.New(0, 0)
	proxyLevel := "会员级"
	for _, b := range beans {
		if b.States == 0 {
			totalCommission, _ := decimal.NewFromString(b.TotalCommission)
			notCashed = notCashed.Add(totalCommission)
		}
		if b.CreatedStr == queryTime {
			proxyLevel = b.ProxyLevelName
		}
	}
	resMap["proxyLevel"] = proxyLevel
	resMap["notCashed"], _ = notCashed.Float64()
	if err != nil {
		utils.ResFaiJSON2(&ctx, err.Error(), "获取失败")
		return
	}
	utils.ResSuccJSON(&ctx, "success", "数据获取成功", config.SUCCESSRES, resMap)
}

/**
* @api {get} api/auth/v1/proxyCommissionsInfo 佣金详情
* @apiDescription
* <span style="color:lightcoral;">接口负责人:wenzhen</span><br/><br/>
* 推广赚钱<br>
* 业务描述:前端进入推广赚钱需要调用此接口</br>
* @apiVersion 1.0.0
* @apiName     proxyCommissionsInfo
* @apiGroup    proxy_module
 * @apiPermission iso,android客户端
* @apiHeader (客户端请求头参数) {string} Authorization 		Bearer + 用户登录获得的token
* @apiParam (客户端请求参数) {string}    	created_str    	 上级数据获取
* @apiError (请求失败返回) {int}      	code            	错误代码
* @apiError (请求失败返回) {string}   	clientMsg       	提示信息
* @apiError (请求失败返回) {string}  	 internalMsg     	错误代码
* @apiError (请求失败返回) {float}    	time_consumed   	后台耗时
*
* @apiErrorExample {json} 失败返回
*{
*    "clientMsg": "获取失败",
*    "code": 204,
*    "internalMsg": "",
*    "time_consumed": 182507
*} *
*
* @apiSuccess (返回结果)  {int}      code            200
* @apiSuccess (返回结果)  {string}   clientMsg       提示信息
* @apiSuccess (返回结果)  {string}   internalMsg     提示信息
* @apiSuccess (返回结果)  {json}  	  data            返回数据
* @apiSuccess (返回结果)  {float}   time_consumed    后台耗时
* @apiSuccess (data字段说明)  {float}   bet_amount    贡献投注量
* @apiSuccess (data字段说明)  {string}   phone    团队成员账号
* @apiSuccess (data字段说明)  {int}   proxy_type    代理类型2棋牌 5真人
* @apiSuccessExample {json} 响应结果
*
{
    "clientMsg": "数据获取成功",
    "code": 200,
    "data": [
        {
            "bet_amount": 500198,
            "user_name": "a13912345678",
            "proxy_type": 2
        },
        {
            "bet_amount": 1000008,
            "user_name": "a15860751424",
            "proxy_type": 2
        },
        {
            "bet_amount": 1503045,
            "user_name": "a13933333333",
            "proxy_type": 2
        }
    ],
    "internalMsg": "success",
    "timeConsumed": 3989
}
*
*/
func (cthis *ProxyController) ProxyCommissionsInfo() {
	ctx := cthis.ctx
	if !utils.RequiredParam(&ctx, []string{"created_str"}) {
		return
	}
	created_str := ctx.URLParam("created_str")
	userid := (ctx.Values().GetString("userid"))
	sqlstr := "SELECT u.user_name, p.proxy_type, p.bet_amount FROM proxy_commissions p LEFT JOIN users u ON p.user_id = u.id WHERE (p.user_id = " + userid + " OR p.parent_id = " + userid + ") AND p.created_str = '" + created_str + "' ORDER BY p.parent_id"
	response, err := utils.Query(cthis.platform, sqlstr)
	if err != nil {
		utils.ResFaiJSON2(&ctx, err.Error(), "获取失败")
		return
	}
	utils.ResSuccJSON(&ctx, "success", "数据获取成功", config.SUCCESSRES, response)
}

/**
* @api {get} api/auth/v1/receiveCommission 领取佣金
* @apiDescription
* <span style="color:lightcoral;">接口负责人:wenzhen</span><br/><br/>
* 领取佣金<br>
* 业务描述:会员自己点击领取佣金按钮，系统自动将所有未领取的佣金转换到账户余额</br>
* @apiVersion 1.0.0
* @apiName     receiveCommission
* @apiGroup    proxy_module
 * @apiPermission iso,android客户端
* @apiHeader (客户端请求头参数) {string} Authorization 		Bearer + 用户登录获得的token
* @apiError (请求失败返回) {int}      	code            	错误代码
* @apiError (请求失败返回) {string}   	clientMsg       	提示信息
* @apiError (请求失败返回) {string}  	 internalMsg     	错误代码
* @apiError (请求失败返回) {float}    	time_consumed   	后台耗时
*
* @apiErrorExample {json} 失败返回
*{
*    "clientMsg": "获取失败",
*    "code": 204,
*    "internalMsg": "",
*    "time_consumed": 182507
*} *
*
* @apiSuccess (返回结果)  {int}      code            200
* @apiSuccess (返回结果)  {string}   clientMsg       提示信息
* @apiSuccess (返回结果)  {string}   internalMsg     提示信息
* @apiSuccess (返回结果)  {json}  	  data            返回数据
* @apiSuccess (返回结果)  {float}   time_consumed    后台耗时
* @apiSuccessExample {json} 响应结果
*
*{
*    "clientMsg": "佣金领取成功",
*    "code": 200,
*    "data": {
*        "exception": {},
*        "msg": "交易成功~",
*        "status": 1
*    },
*    "internalMsg": "success",
*    "timeConsumed": 436832
*}
*
*/
func (cthis *ProxyController) ReceiveCommission() {
	ctx := cthis.ctx
	userid, _ := (ctx.Values().GetInt("userid"))
	var beans []xorm.ProxyCommissions
	err := models.MyEngine[cthis.platform].Where("user_id = ? and states = 0", userid).Desc("created").Find(&beans)
	if err != nil {
		utils.ResFaiJSON2(&ctx, err.Error(), "佣金信息获取失败")
		return
	}
	if len(beans) == 0 {
		utils.ResFaiJSON2(&ctx, "", "没有可以领取的佣金信息")
		return
	}
	totalCommission, _ := decimal.NewFromString("0")
	ids := make([]int, 0)
	for i, b := range beans {
		tc, _ := decimal.NewFromString(b.TotalCommission)
		totalCommission = totalCommission.Add(tc)
		beans[i].States = 1
		ids = append(ids, b.Id)
	}

	//变更投注记录状态
	session := models.MyEngine[cthis.platform].NewSession()
	defer session.Close()
	err = session.Begin()
	if err != nil {
		utils.ResFaiJSON2(&ctx, err.Error(), "佣金领取失败")
		return
	}
	params := make(map[string]string)
	params["states"] = "1"
	//更新佣金记录表
	updatesql := utils.UpdateBatchSame("proxy_commissions", params, ids)
	_, err = session.Exec(updatesql)
	if err != nil {
		session.Rollback()
		utils.ResFaiJSON2(&ctx, err.Error(), "佣金领取失败")
		return
	}
	amount, _ := totalCommission.Float64()
	info := map[string]interface{}{
		"user_id":     userid,
		"type_id":     config.FUNDBROKERAGE,
		"amount":      amount,
		"order_id":    utils.CreationOrder("YJ", strconv.Itoa(userid)),
		"msg":         "代理佣金领取",
		"finish_rate": 1.0, //需满足的打码量比例
	}
	balance := fund.NewUserFundChange(cthis.platform)
	balanceUpdateRes := balance.BalanceUpdate(info, nil)
	if balanceUpdateRes["status"] == 1 {
		session.Commit()
		utils.ResSuccJSON(&ctx, "success", "佣金领取成功", config.SUCCESSRES, balanceUpdateRes)
		return
	} else {
		session.Rollback()
	}
	utils.ResFaiJSON2(&ctx, balanceUpdateRes["msg"].(string), "佣金领取失败")
}

/**
 * @api {get} api/auth/v1/teamMembers 推广赚钱团队成员
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人:wenzhen</span><br/><br/>
 * 团队成员情况<br>
 * 业务描述:推广赚钱页面的团队成员</br>
 * @apiVersion 1.0.0
 * @apiName     teamMembers
 * @apiGroup    proxy_module
 * @apiPermission ios,android客户端
 * @apiHeader (客户端请求头参数) {string} Authorization Bearer + 用户登录获得的token
 * @apiParam (客户端请求参数) {string} userName   搜索账号
 * @apiParam (客户端请求参数) {int} page   请求第几页
 * @apiSuccess (返回结果)  {int}      code            200
 * @apiSuccess (返回结果)  {string}   clientMsg       提示信息
 * @apiSuccess (返回结果)  {string}   internalMsg     提示信息
 * @apiSuccess (返回结果)  {json}  	  data            返回数据
 * @apiSuccess (返回结果)  {float}   time_consumed    后台耗时
 * @apiSuccess (data json对象字段说明)  {array}   list    成员信息对象数组,无数据返回空数组
 * @apiSuccess (data-list数组元素对象字段说明)  {float}   charged_amount   充值量
 * @apiSuccess (data-list数组元素对象字段说明)  {int}   created   注册时间 时间戳秒
 * @apiSuccess (data-list数组元素对象字段说明)  {int}   last_login_time   最后登录时间戳秒
 * @apiSuccess (data-list数组元素对象字段说明)  {string}   username   用户账号
 * @apiSuccess (data-list数组元素对象字段说明)  {string}   vip_level  成员会员等级
 * @apiSuccessExample {json} 响应结果
{
    "clientMsg": "数据获取成功",
    "code": 200,
    "data": [
        {
            "charged_amount": 0,
            "created": 1553916753,
            "last_login_time": 1556187096,
            "phone": "15860751424",
            "vip_level": 1
        },
        {
            "charged_amount": 16.87,
            "created": 1554358645,
            "last_login_time": 1556181466,
            "phone": "13933333333",
            "vip_level": 1
        },
        {
            "charged_amount": 0,
            "created": 1555378571,
            "last_login_time": 0,
            "phone": "15243639098",
            "vip_level": 1
        },
        {
            "charged_amount": 0,
            "created": 1554728053,
            "last_login_time": 0,
            "phone": "15860751469",
            "vip_level": 1
        }
    ],
    "internalMsg": "success",
    "timeConsumed": 11968
}
*/

func (cthis *ProxyController) TeamMembers() {
	ctx := cthis.ctx
	if !utils.RequiredParam(&ctx, []string{"page"}) {
		return
	}
	phone := ctx.URLParam("phone")
	page, _ := ctx.URLParamInt("page")
	userid := ctx.Values().GetString("userid")
	sqlstr := "SELECT u.`last_login_time`,u.`created`,u.`user_name`,u.`vip_level`,a.`charged_amount` FROM users u LEFT JOIN accounts a ON u.id = a.user_id WHERE u.parent_id=" + userid
	if phone != "" {
		sqlstr += " and u.user_name='" + phone + "'"
	}
	sqlstr += " ORDER BY u.`last_login_time` DESC limit " + strconv.Itoa((page-1)*20) + ",20"
	response, err := utils.Query(cthis.platform, sqlstr)

	if err != nil {
		utils.ResFaiJSON2(&ctx, err.Error(), "获取失败")
		return
	}
	if phone != "" && response == nil {
		utils.ResFaiJSON(&ctx, "success", "该会员号不存在", config.NOTGETDATA)
		return
	}
	if len(response) == 0 {
		utils.ResSuccJSON(&ctx, "success", "数据获取成功", config.SUCCESSRES, make([]map[string]string, 0))
	} else {
		utils.ResSuccJSON(&ctx, "success", "数据获取成功", config.SUCCESSRES, response)
	}
}

/**
 * @api {get} api/auth/v1/commissionRecord 佣金记录
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人:wenzhen</span><br/><br/>
 * 佣金记录<br>
 * 业务描述:查询佣金记录</br>
 * @apiVersion 1.0.0
 * @apiName     commissionRecord
 * @apiGroup    proxy_module
 * @apiPermission ios,android客户端
 * @apiHeader (客户端请求头参数) {string} Authorization Bearer + 用户登录获得的token
 * @apiParam (客户端请求参数) {string} endTime   	交易结束日期 时间戳
 * @apiParam (客户端请求参数) {string} startTime    交易开始日期 时间戳
 * @apiSuccess (返回结果)  {int}      code            200
 * @apiSuccess (返回结果)  {string}   clientMsg       提示信息
 * @apiSuccess (返回结果)  {string}   internalMsg     提示信息
 * @apiSuccess (返回结果)  {json}  	  data            返回数据
 * @apiSuccess (返回结果)  {float}   time_consumed    后台耗时
 * @apiSuccess (data-list数组元素对象字段说明)  {float}   amount   转出金额
 * @apiSuccess (data-list数组元素对象字段说明)  {float}   balance   转出后余额
 * @apiSuccess (data-list数组元素对象字段说明)  {int}   created   交易时间
 * @apiSuccess (data-list数组元素对象字段说明)  {string}   order_id   订单编号
 * @apiSuccessExample {json} 响应结果
{
    "clientMsg": "数据获取成功",
    "code": 200,
    "data": [
        {
            "amount": 22667.92,
            "balance": 23267.92,
            "created": 1556205828,
            "order_id": "YJ201904252323485914bd"
        }
    ],
    "internalMsg": "success",
    "timeConsumed": 2990
}
*/

func (cthis *ProxyController) CommissionRecord() {
	ctx := cthis.ctx
	endTime := ctx.URLParam("endTime")
	if endTime == "" {
		endTime = strconv.Itoa(utils.GetNowTime())
	}
	startTime := ctx.URLParam("startTime")
	userid := (ctx.Values().GetString("userid"))

	sqlstr := "SELECT a.`created`,a.order_id,a.`amount`,a.`balance` FROM account_infos a WHERE a.`type`=" + strconv.Itoa(config.FUNDBROKERAGE) + " and a.user_id = " + userid + " and created<= " + endTime
	if startTime != "" {
		sqlstr += " and a.created >=" + startTime
	}
	response, err := utils.Query(cthis.platform, sqlstr)
	if err != nil {
		utils.ResFaiJSON2(&ctx, err.Error(), "获取失败")
		return
	}
	utils.ResSuccJSON(&ctx, "success", "数据获取成功", config.SUCCESSRES, response)
}
