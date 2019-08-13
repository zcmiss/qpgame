package frontEndControllers

import (
	"github.com/kataras/iris"
	"github.com/shopspring/decimal"
	"qpgame/app/fund"
	"qpgame/common/utils"
	"qpgame/config"
	"qpgame/models"
	"qpgame/models/xorm"
	"qpgame/ramcache"
	"strconv"
)

type BetController struct {
	platform string
	ctx      iris.Context
}

//构造函数
func NewBetController(ctx iris.Context) *BetController {
	obj := new(BetController)
	obj.platform = ctx.Params().Get("platform")
	obj.ctx = ctx
	return obj
}

/**
* @api {get} api/auth/v1/getBetsInfo 投注记录明细
* @apiDescription
* <span style="color:lightcoral;">接口负责人:wenzhen</span><br/><br/>
* 投注记录明细<br>
* 业务描述:投注记录明细</br>
* @apiVersion 1.0.0
* @apiName     betsInfo
* @apiGroup    bets_module
 * @apiPermission iso,android客户端
* @apiHeader (客户端请求头参数) {string} Authorization 		Bearer + 用户登录获得的token
* @apiParam (客户端请求参数) {int}    	pageNum    	   		当前页
* @apiParam (客户端请求参数) {int}    	betTime    	   		派彩时间  0全部1今天2昨天3一个月内
* @apiParam (客户端请求参数) {string}   parentId    		游戏类型编号
* @apiParam (客户端请求参数) {string}   gameCategorieId    	游戏平台编号
* @apiError (请求失败返回) {int}      	code            	错误代码
* @apiError (请求失败返回) {string}   	clientMsg       	提示信息
* @apiError (请求失败返回) {string}  	 internalMsg     	错误代码
* @apiError (请求失败返回) {float}    	timeConsumed   	后台耗时
*
* @apiErrorExample {json} 失败返回
*{
*    "clientMsg": "获取失败",
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
{
    "clientMsg": "获取成功",
    "code": 200,
    "data": [
        {
            "Amount": "5.000",
            "AmountAll": "0.000",
            "AmountPlatform": "0.000",
            "Created": 1555654820,
            "Ented": 1555434631,
            "Started": 0,
            "GameCode": "SC03",
            "Gt": "",
            "Id": 78822,
            "OrderId": "8190416055877116",
            "PlatformId": 5,
            "Accountname": "ED728",
            "PlatformName": "",
            "Reward": "50.000",
            "UserId": 28,
            "Name": "金拉霸",
            "GameCategorieId": "12",
            "ParentId": "4"
        },
		....
        }
    ],
    "internalMsg": "获取成功",
    "timeConsumed": 30916
}
*
*/
func (cthis *BetController) GetBetsInfo() {
	ctx := cthis.ctx
	if !utils.RequiredParam(&ctx, []string{"pageNum", "parentId", "betTime"}) {
		return
	}
	userid, _ := (ctx.Values().GetInt("userid"))
	pageNum, err := strconv.Atoi(ctx.URLParam("pageNum"))
	if err != nil || pageNum < 1 {
		utils.ResFaiJSON(&ctx, "", "参数错误", config.PARAMERROR)
		return
	}
	parentId := ctx.URLParam("parentId")
	gameCategorieId := ctx.URLParam("gameCategorieId")
	betTime, err := ctx.URLParamInt("betTime")
	if err != nil {
		utils.ResFaiJSON(&ctx, "", "派彩时间参数错误", config.PARAMERROR)
		return
	}
	var beans []xorm.Bets
	sql := "SELECT b.*, p.`name`, p.`game_categorie_id`, g.`parent_id` FROM bets_" + strconv.Itoa(userid%10) + " b LEFT JOIN platform_games p ON b.`game_code` = p.service_code and b.platform_id=p.plat_id LEFT JOIN game_categories g ON p.`game_plat_id` = g.`id` where g.`parent_id` = ? and b.user_id = ?"
	if gameCategorieId != "" {
		sql += " and p.game_plat_id = " + gameCategorieId
	}
	started, ended := utils.GetQueryTime(betTime)
	sql += " and b.ented >= " + started + " and b.ented<=" + ended
	sql += " ORDER BY b.`ented` DESC limit " + strconv.Itoa((pageNum-1)*20) + ",20"
	session := models.MyEngine[cthis.platform].SQL(sql, parentId, userid)
	err = session.Find(&beans)
	if err != nil {
		utils.ResFaiJSON2(&ctx, err.Error(), "投注记录获取失败")
	}
	if len(beans) == 0 {
		checkNil(&ctx, nil)
	} else {
		checkNil(&ctx, beans)
	}
}

/**
* @api {get} api/auth/v1/getAccountInfo 账户资金明细
* @apiDescription
* <span style="color:lightcoral;">接口负责人:wenzhen</span><br/><br/>
* 账户资金明细<br>
* 业务描述:账户资金明细</br>
* @apiVersion 1.0.0
* @apiName     getAccountInfo
* @apiGroup    bets_module
 * @apiPermission iso,android客户端
* @apiHeader (客户端请求头参数) {string} Authorization 		Bearer + 用户登录获得的token
* @apiParam (客户端请求参数) {int}    	pageNum    	   		当前页
* @apiParam (客户端请求参数) {int}    	created    	   		交易创建时间  0全部1今天2昨天3一个月内
* @apiParam (客户端请求参数) {string}   typecode    		交易类型
* @apiError (请求失败返回) {int}      	code            	错误代码
* @apiError (请求失败返回) {string}   	clientMsg       	提示信息
* @apiError (请求失败返回) {string}  	 internalMsg     	错误代码
* @apiError (请求失败返回) {float}    	timeConsumed   	后台耗时
*
* @apiErrorExample {json} 失败返回
*{
*    "clientMsg": "获取失败",
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
*            "Id": 432,
*            "UserId": 28,
*            "Amount": "4336.650",
*            "Balance": "43779.500",
*            "ChargedAmount": "43366.500",
*            "ChargedAmountOld": "39029.850",
*            "OrderId": "XM2019042317013452edc8",
*            "Msg": "用户手动洗码",
*            "Type": 1,
*            "Created": 1556010094
*        },
*        {
*            "Id": 431,
*            "UserId": 28,
*            "Amount": "4336.650",
*            "Balance": "39442.850",
*            "ChargedAmount": "39029.850",
*            "ChargedAmountOld": "34693.200",
*            "OrderId": "XM201904231635118303e0",
*            "Msg": "用户手动洗码",
*            "Type": 1,
*            "Created": 1556008511
*        }*,
*        ...
*    ],
*    "internalMsg": "获取成功",
*    "timeConsumed": 19947
*}
*
*/
func (cthis *BetController) GetAccountInfo() {
	ctx := cthis.ctx
	if !utils.RequiredParam(&ctx, []string{"pageNum", "created"}) {
		return
	}
	userid, _ := (ctx.Values().GetInt("userid"))
	pageNum, err := strconv.Atoi(ctx.URLParam("pageNum"))
	if err != nil || pageNum < 1 {
		utils.ResFaiJSON(&ctx, "", "参数错误", config.PARAMERROR)
		return
	}
	typecode := ctx.URLParam("typecode")
	created, err := ctx.URLParamInt("created")
	if err != nil {
		utils.ResFaiJSON(&ctx, "", "时间参数错误", config.PARAMERROR)
		return
	}

	sql := "select * from account_infos a where a.user_id = " + strconv.Itoa(userid)
	session := models.MyEngine[cthis.platform]
	var beans []xorm.AccountInfos
	if typecode != "" {
		sql += " and a.type = " + typecode
	}
	started, ended := utils.GetQueryTime(created)
	sql += " and a.created >= " + started + " and a.created<=" + ended
	sql += " ORDER BY a.`created` DESC limit " + strconv.Itoa((pageNum-1)*20) + ",20"
	err = session.SQL(sql).Find(&beans)
	if err != nil {
		utils.ResFaiJSON2(&ctx, err.Error(), "资金明细获取失败")
		return
	}
	if len(beans) == 0 {
		checkNil(&ctx, nil)
	} else {
		checkNil(&ctx, beans)
	}
}

/**
* @api {get} api/auth/v1/getBetsSearchType 获取投注记录平台查询列表
* @apiDescription
* <span style="color:lightcoral;">接口负责人:wenzhen</span><br/><br/>
* 获取投注记录平台查询列表<br>
* 业务描述:获取投注记录平台查询列表</br>
* @apiVersion 1.0.0
* @apiName     getBetsSearchType
* @apiGroup    bets_module
 * @apiPermission iso,android客户端
* @apiHeader (客户端请求头参数) {string} Authorization 		Bearer + 用户登录获得的token
* @apiError (请求失败返回) {int}      	code            	错误代码
* @apiError (请求失败返回) {string}   	clientMsg       	提示信息
* @apiError (请求失败返回) {string}  	 internalMsg     	错误代码
* @apiError (请求失败返回) {float}    	timeConsumed   	后台耗时
*
* @apiErrorExample {json} 失败返回
*{
*    "clientMsg": "获取失败",
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
*    "data": {
*        "2": {
*            "name": "棋牌游戏",
*            "searchList": [
*                {
*                    "name": "FG棋牌",
*                    "value": "6"
*                },
*                {
*                    "name": "开元棋牌",
*                    "value": "11"
*                }
*            ]
*        },
*        "3": {
*            "name": "捕鱼游戏",
*            "searchList": []
*        },
*        "4": {
*            "name": "电子游艺",
*            "searchList": [
*                {
*                    "name": "FG电子",
*                    "value": "7"
*                },
*                {
*                    "name": "AE电子",
*                    "value": "8"
*                },
*                {
*                    "name": "MG电子",
*                    "value": "9"
*                },
*                {
*                    "name": "AG电子",
*                    "value": "12"
*                }
*            ]
*        },
*        "5": {
*            "name": "真人视讯",
*            "searchList": [
*                {
*                    "name": "AG视讯",
*                    "value": "13"
*                }
*            ]
*        }
*    },
*    "internalMsg": "获取成功",
*    "timeConsumed": 1993
*}
*
*/
func (cthis *BetController) GetBetsSearchType() {
	ctx := cthis.ctx
	data, _ := ramcache.BetsSearchType.Load(cthis.platform)
	utils.ResSuccJSON(&ctx, "获取成功", "获取成功", config.SUCCESSRES, data)
}

/**
* @api {get} api/auth/v1/getWashCode 获取洗码信息
* @apiDescription
* <span style="color:lightcoral;">接口负责人:wenzhen</span><br/><br/>
* 获取洗码信息<br>
* 业务描述:获取洗码信息</br>
* @apiVersion 1.0.0
* @apiName     getWashCode
* @apiGroup    bets_module
 * @apiPermission iso,android客户端
* @apiHeader (客户端请求头参数) {string} Authorization 		Bearer + 用户登录获得的token
* @apiError (请求失败返回) {int}      	code            	错误代码
* @apiError (请求失败返回) {string}   	clientMsg       	提示信息
* @apiError (请求失败返回) {string}  	 internalMsg     	错误代码
* @apiError (请求失败返回) {float}    	timeConsumed   	后台耗时
*
* @apiErrorExample {json} 失败返回
*{
*    "clientMsg": "获取失败",
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
*    "data": {
*        "totalamount": "788480.34",
*        "washcodes": [
*            {
*                "Id": 0,
*                "RecordId": 0,
*                "Typename": "捕鱼游戏",
*                "Typeid": 3,
*                "Gamename": "FG捕鱼",
*                "Betamount": "543335.900",
*                "Amount": "2988.35"
*            },
*            {
*                "Id": 0,
*                "RecordId": 0,
*                "Typename": "捕鱼游戏",
*                "Typeid": 3,
*                "Gamename": "MG捕鱼",
*                "Betamount": "3171.700",
*                "Amount": "17.44"
*            },
*            {
*                "Id": 0,
*                "RecordId": 0,
*                "Typename": "捕鱼游戏",
*                "Typeid": 3,
*                "Gamename": "乐游捕鱼",
*                "Betamount": "4247.000",
*                "Amount": "23.36"
*            },
*            ...*
*        ]
*    },
*    "internalMsg": "获取成功",
*    "timeConsumed": 250330
*}
*
*/
func (cthis *BetController) GetWashCode() {
	ctx := cthis.ctx
	userid, _ := (ctx.Values().GetInt("userid"))
	var beans []xorm.WashCodeInfos
	sql := "SELECT gp.name type_name, g.`parent_id` type_id, g.`name` game_name, SUM(b.amount) bet_amount, ROUND(g.`rate`*v.wash_code,4) rate, ROUND(SUM(b.amount) * ROUND(g.`rate`*v.wash_code,4),3) amount FROM bets_" + strconv.Itoa(userid%10) + " b LEFT JOIN platform_games p ON b.`game_code` = p.`service_code` AND b.`platform_id`=p.`plat_id` LEFT JOIN game_categories g ON g.`id` = p.`game_plat_id` LEFT JOIN game_categories gp ON g.parent_id = gp.id LEFT JOIN users u ON b.user_id = u.id LEFT JOIN vip_levels v ON u.vip_level = v.id WHERE b.user_id = ? AND b.`rebate_state`=0 GROUP BY gp.name,g.`name`, g.`rate`, g.`parent_id`,v.wash_code"
	err := models.MyEngine[cthis.platform].SQL(sql, userid).Find(&beans)
	if err != nil {
		utils.ResFaiJSON2(&ctx, err.Error(), "洗码信息获取失败")
		return
	}
	totalBetAmount, _ := decimal.NewFromString("0")
	totalAmount, _ := decimal.NewFromString("0")
	for _, b := range beans {
		betAmount, _ := decimal.NewFromString(b.BetAmount)
		amount, _ := decimal.NewFromString(b.Amount)
		totalBetAmount = totalBetAmount.Add(betAmount)
		totalAmount = totalAmount.Add(amount)
	}
	res := make(map[string]interface{})
	res["totalBetAmount"] = totalBetAmount.String()
	res["totalAmount"] = totalAmount.String()
	res["washcodes"] = beans
	if len(beans) == 0 {
		utils.ResSuccJSON(&ctx, "获取成功", "获取成功,但是没有数据", config.SUCCESSRES, make(map[string]interface{}))
		return
	} else {
		checkNil(&ctx, res)
	}
}

/**
* @api {get} api/auth/v1/washCode 手动洗码
* @apiDescription
* <span style="color:lightcoral;">接口负责人:wenzhen</span><br/><br/>
* 手动洗码<br>
* 业务描述:手动洗码</br>
* @apiVersion 1.0.0
* @apiName     washCode
* @apiGroup    bets_module
 * @apiPermission iso,android客户端
* @apiHeader (客户端请求头参数) {string} Authorization 		Bearer + 用户登录获得的token
* @apiError (请求失败返回) {int}      	code            	错误代码
* @apiError (请求失败返回) {string}   	clientMsg       	提示信息
* @apiError (请求失败返回) {string}  	 internalMsg     	错误代码
* @apiError (请求失败返回) {float}    	timeConsumed   	后台耗时
*
* @apiErrorExample {json} 失败返回
*{
*    "clientMsg": "获取失败",
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
*    "clientMsg": "手动洗码成功",
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
func (cthis *BetController) WashCode() {
	ctx := cthis.ctx
	userid, _ := (ctx.Values().GetInt("userid"))
	sql := "SELECT gp.name type_name, g.`parent_id` type_id, g.`name` game_name, SUM(b.amount) bet_amount, ROUND(g.`rate`*v.wash_code,4) rate, ROUND(SUM(b.amount) * ROUND(g.`rate`*v.wash_code,4),3) amount FROM bets_" + strconv.Itoa(userid%10) + " b LEFT JOIN platform_games p ON b.`game_code` = p.`service_code` AND b.`platform_id`=p.`plat_id` LEFT JOIN game_categories g ON g.`id` = p.`game_plat_id` LEFT JOIN game_categories gp ON g.parent_id = gp.id LEFT JOIN users u ON b.user_id = u.id LEFT JOIN vip_levels v ON u.vip_level = v.id WHERE b.user_id = ? AND b.`rebate_state`=0 GROUP BY gp.name,g.`name`, g.`rate`, g.`parent_id`,v.wash_code"
	var beans []xorm.WashCodeInfos
	err := models.MyEngine[cthis.platform].SQL(sql, userid).Find(&beans)
	if err != nil {
		utils.ResFaiJSON2(&ctx, err.Error(), "洗码信息获取失败")
		return
	}
	totalBetAmount, _ := decimal.NewFromString("0")
	totalAmount, _ := decimal.NewFromString("0")
	for _, b := range beans {
		betAmount, _ := decimal.NewFromString(b.BetAmount)
		amount, _ := decimal.NewFromString(b.Amount)
		totalBetAmount = totalBetAmount.Add(betAmount)
		totalAmount = totalAmount.Add(amount)
	}
	amount, _ := totalAmount.Float64()
	if amount < 1 {
		utils.ResFaiJSON2(&ctx, "洗码金额不足", "洗码金额必须大于1元")
		return
	}
	//变更投注记录状态
	session := models.MyEngine[cthis.platform].NewSession()
	defer session.Close()
	err = session.Begin()
	if err != nil {
		utils.ResFaiJSON2(&ctx, err.Error(), "手动洗码失败")
		return
	}
	record := xorm.WashCodeRecords{TotalBetamount: totalBetAmount.String(), UserId: userid, Amount: totalAmount.String(), Washtime: utils.GetNowTime()}
	_, err = session.Insert(&record)
	if err != nil {
		session.Rollback()
		utils.ResFaiJSON2(&ctx, err.Error(), "手动洗码失败")
		return
	}
	for i, _ := range beans {
		beans[i].RecordId = record.Id
	}
	_, err = session.Insert(&beans)
	if err != nil {
		session.Rollback()
		utils.ResFaiJSON2(&ctx, err.Error(), "手动洗码失败")
		return
	}
	//查询未洗码的投注记录
	bets := make([]xorm.Bets, 0)
	session.SQL("select * from bets_"+strconv.Itoa(userid%10)+" where user_id=? and rebate_state=?", userid, 0).Find(&bets)
	params := make(map[string]string)
	params["rebate_state"] = "1"
	ids := make([]int, 0)
	for _, b := range bets {
		ids = append(ids, b.Id)
	}
	updatesql := utils.UpdateBatchSame("bets_"+strconv.Itoa(userid%10), params, ids)
	_, err = session.Exec(updatesql)
	if err != nil {
		session.Rollback()
		utils.ResFaiJSON2(&ctx, err.Error(), "手动洗码失败")
		return
	}

	info := map[string]interface{}{
		"user_id":     userid,
		"type_id":     config.FUNDXIMA,
		"amount":      amount,
		"order_id":    utils.CreationOrder("XM", strconv.Itoa(userid)),
		"msg":         "用户手动洗码",
		"finish_rate": 1.0, //需满足的打码量比例
	}
	balance := fund.NewUserFundChange(cthis.platform)
	balanceUpdateRes := balance.BalanceUpdate(info, nil)
	if balanceUpdateRes["status"] == 1 {
		session.Commit()
		utils.ResSuccJSON(&ctx, "success", "手动洗码成功", config.SUCCESSRES, balanceUpdateRes)
		return
	} else {
		session.Rollback()
	}
	utils.ResFaiJSON2(&ctx, balanceUpdateRes["msg"].(string), "手动洗码失败")
}

/**
* @api {get} api/auth/v1/getWashCodeRecords 获取洗码记录
* @apiDescription
* <span style="color:lightcoral;">接口负责人:wenzhen</span><br/><br/>
* 获取洗码记录<br>
* 业务描述:获取洗码记录</br>
* @apiVersion 1.0.0
* @apiName     getWashCodeRecords
* @apiGroup    bets_module
 * @apiPermission iso,android客户端
* @apiHeader (客户端请求头参数) {string} Authorization 		Bearer + 用户登录获得的token
* @apiParam (客户端请求参数) {int}    	pageNum    	   		当前页
* @apiError (请求失败返回) {int}      	code            	错误代码
* @apiError (请求失败返回) {string}   	clientMsg       	提示信息
* @apiError (请求失败返回) {string}  	 internalMsg     	错误代码
* @apiError (请求失败返回) {float}    	timeConsumed   	后台耗时

* @apiErrorExample {json} 失败返回
*{
*    "clientMsg": "获取失败",
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
*            "Id": 27,
*            "TotalBetamount": "788480.34",
*            "Amount": "4336.65",
*            "Washtime": 1556008511,
*            "UserId": 28
*        },
*        {
*            "Id": 26,
*            "TotalBetamount": "788480.34",
*            "Amount": "4336.65",
*            "Washtime": 1556008340,
*            "UserId": 28
*        },
*        {
*            "Id": 25,
*            "TotalBetamount": "788480.34",
*            "Amount": "4336.65",
*            "Washtime": 1556006945,
*            "UserId": 28
*        }
*    ],
*    "internalMsg": "获取成功",
*    "timeConsumed": 6981
*}*
*
*/
func (cthis *BetController) GetWashCodeRecords() {
	ctx := cthis.ctx
	if !utils.RequiredParam(&ctx, []string{"pageNum"}) {
		return
	}
	userid, _ := (ctx.Values().GetInt("userid"))
	pageNum, err := strconv.Atoi(ctx.URLParam("pageNum"))
	if err != nil || pageNum < 1 {
		utils.ResFaiJSON(&ctx, "", "参数错误", config.PARAMERROR)
		return
	}
	session := models.MyEngine[cthis.platform]
	var beans []xorm.WashCodeRecords
	err = session.Desc("washtime").Where("user_id = ?", userid).Limit(20, (pageNum-1)*20).Find(&beans)
	if err != nil {
		utils.ResFaiJSON2(&ctx, err.Error(), "洗码记录获取失败")
		return
	}
	if len(beans) == 0 {
		checkNil(&ctx, nil)
	} else {
		checkNil(&ctx, beans)
	}
}

/**
* @api {get} api/auth/v1/getWashCodeInfos 获取洗码记录详情
* @apiDescription
* <span style="color:lightcoral;">接口负责人:wenzhen</span><br/><br/>
* 获取洗码记录详情<br>
* 业务描述:获取洗码记录详情</br>
* @apiVersion 1.0.0
* @apiName     getWashCodeInfos
* @apiGroup    bets_module
 * @apiPermission iso,android客户端
* @apiHeader (客户端请求头参数) {string} Authorization 		Bearer + 用户登录获得的token
* @apiParam (客户端请求参数) {int}    	recordId    	   	洗码记录编号
* @apiError (请求失败返回) {int}      	code            	错误代码
* @apiError (请求失败返回) {string}   	clientMsg       	提示信息
* @apiError (请求失败返回) {string}  	 internalMsg     	错误代码
* @apiError (请求失败返回) {float}    	timeConsumed   	后台耗时
*
* @apiErrorExample {json} 失败返回
*{
*    "clientMsg": "获取失败",
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
*            "Id": 232,
*            "RecordId": 25,
*            "TypeName": "捕鱼游戏",
*            "TypeId": 3,
*            "GameName": "FG捕鱼",
*            "BetAmount": "543335.90",
*            "Amount": "2988.35"
*        },
*        {
*            "Id": 233,
*            "RecordId": 25,
*            "TypeName": "捕鱼游戏",
*            "TypeId": 3,
*            "GameName": "MG捕鱼",
*            "BetAmount": "3171.70",
*            "Amount": "17.44"
*        },
*        ...
*    ],
*    "internalMsg": "获取成功",
*    "timeConsumed": 13963
*}
*
*/
func (cthis *BetController) GetWashCodeInfos() {
	ctx := cthis.ctx
	if !utils.RequiredParam(&ctx, []string{"recordId"}) {
		return
	}
	recordId, err := strconv.Atoi(ctx.URLParam("recordId"))
	session := models.MyEngine[cthis.platform]
	var beans []xorm.WashCodeInfos
	err = session.Where("record_id = ?", recordId).Find(&beans)
	if err != nil {
		utils.ResFaiJSON2(&ctx, err.Error(), "洗码记录获取失败")
		return
	}
	if len(beans) == 0 {
		checkNil(&ctx, nil)
	} else {
		checkNil(&ctx, beans)
	}
}
