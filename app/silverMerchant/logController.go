package silverMerchant

import (
	"github.com/kataras/iris"
	"html"
	"qpgame/common/utils"
	"qpgame/config"
	"qpgame/models"
	"strconv"
)

type LogController struct {
	platform string
	ctx      iris.Context
}

//构造函数
func NewSilverLogController(ctx iris.Context) *LogController {
	obj := new(LogController)
	obj.platform = ctx.Params().Get("platform")
	obj.ctx = ctx
	return obj
}

/**
 * @api {post} silverMerchant/api/auth/v1/loginLog 登录日志
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: aTian</span><br/><br/>
 * 会员登录日志<br>
 * 业务描述：银商会员登录日志查询接口</br>
 * @apiVersion 1.0.0
 * @apiName     loginLog
 * @apiGroup    silver_merchant
 * @apiPermission PC客户端
 * @apiParam (客户端请求参数) {string} from_time    通过登录时间范围查询：起始时间
 * @apiParam (客户端请求参数) {string} to_time      通过登录时间范围查询：结束时间
 * @apiParam (客户端请求参数) {string} ip           通过登录IP查询
 * @apiParam (客户端请求参数) {string} login_city   通过登录城市查询
 * @apiParam (客户端请求参数) {string} page         页码 (默认1)
 * @apiParam (客户端请求参数) {string} size         每页显示行数 (默认20)
 * @apiParam (客户端请求参数) {string} order_by     排序字段[id(默认), login_time, ip, login_city]
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
 * @apiSuccess (data对象字段说明) {string} order_by 排序字段
 * @apiSuccess (data对象字段说明) {int} page 当前页码
 * @apiSuccess (data对象字段说明) {int} size 每页显示行数
 * @apiSuccess (data对象字段说明) {string} sort 排序方式
 * @apiSuccess (data对象字段说明) {int} total 数据总行数
 * @apiSuccess (data对象字段说明) {array} list 日志列表
 * @apiSuccess (data对象子字段 list 说明) {string} id 登录日志主键
 * @apiSuccess (data对象子字段 list 说明) {string} ip 登录IP
 * @apiSuccess (data对象子字段 list 说明) {string} login_city 登录城市
 * @apiSuccess (data对象子字段 list 说明) {string} login_time 登录时间
 * @apiSuccessExample {json} 响应结果
 * {
 *     "clientMsg": "获取登录日志成功",
 *     "code": 200,
 *     "data": {
 *         "list": [
 *             {
 *                 "account": "test7",
 *                 "id": "1",
 *                 "ip": "127.0.0.1",
 *                 "login_city": "",
 *                 "login_time": "2019-06-13 09:01:48",
 *                 "merchant_name": ""
 *             }
 *         ],
 *         "order_by": "id",
 *         "page": 1,
 *         "size": 20,
 *         "sort": "DESC",
 *         "total": 1
 *     },
 *     "internalMsg": "",
 *     "timeConsumed": 273287
 * }
 */
func (cthis *LogController) LoginLog() {
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

	filterArr := []map[string]string{
		{"key": "merchant_id", "val": sSmUserId},
		{"key": "login_time", "val": sFromTimestamp, "condition": "gt"},
		{"key": "login_time", "val": sToTimestamp, "condition": "lt"},
		{"key": "ip", "val": postData.Get("ip")},
		{"key": "login_city", "val": postData.Get("login_city"), "condition": "like_ab"},
	}
	where := utils.BuildWhere(filterArr)
	countSql := "SELECT COUNT(*) cnt FROM silver_merchant_login_logs" + where
	engine := models.MyEngine[cthis.platform]
	cntRes, cntErr := engine.SQL(countSql).QueryString()
	if cntErr != nil {
		utils.ResFaiJSON(&ctx, "1906111353", "获取登录日志失败", config.NOTGETDATA)
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
		orderBy := " ORDER BY `" + html.EscapeString(sOrderBy) + "` " + html.EscapeString(sSort)

		sql := "SELECT * FROM silver_merchant_login_logs" + where + orderBy + limit
		logList, queryErr := engine.SQL(sql).QueryString()
		if queryErr != nil {
			utils.ResFaiJSON(&ctx, "1906111015", "获取登录日志失败", config.NOTGETDATA)
			return
		}

		var list interface{}
		if len(logList) > 0 {
			for k, log := range logList {
				if sLoginTime, fExist := log["login_time"]; fExist {
					iLoginTime, _ := strconv.Atoi(sLoginTime)
					logList[k]["login_time"] = utils.TimestampToDateStr(int64(iLoginTime), "1")
				}
			}
			list = logList
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
		utils.ResSuccJSON(&ctx, "", "获取登录日志成功", config.SUCCESSRES, data)
	} else {
		utils.ResSuccJSON(&ctx, "", "获取登录日志成功", config.SUCCESSRES, map[string]interface{}{
			"list":     make([]map[string]string, 0),
			"page":     iPage,
			"size":     iSize,
			"total":    total,
			"order_by": sOrderBy,
			"sort":     sSort,
		})
	}
}

/**
 * @api {post} silverMerchant/api/auth/v1/operationLog 操作日志
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: aTian</span><br/><br/>
 * 操作日志<br>
 * 业务描述：银商会员操作日志查询接口</br>
 * @apiVersion 1.0.0
 * @apiName     OperationLog
 * @apiGroup    silver_merchant
 * @apiPermission PC客户端
 * @apiParam (客户端请求参数) {string} from_time    通过创建时间范围查询：起始时间
 * @apiParam (客户端请求参数) {string} to_time      通过创建时间范围查询：结束时间
 * @apiParam (客户端请求参数) {string} ip           通过创建IP查询
 * @apiParam (客户端请求参数) {string} login_city   通过创建城市查询
 * @apiParam (客户端请求参数) {string} page         页码 (默认1)
 * @apiParam (客户端请求参数) {string} size         每页显示行数 (默认20)
 * @apiParam (客户端请求参数) {string} order_by     排序字段[id(默认), created, ip, login_city]
 * @apiParam (客户端请求参数) {string} sort         排序方式[ASC, DESC(默认)]
 * @apiError (请求失败返回)   {int}    code         错误代码
 * @apiError (请求失败返回)   {string} clientMsg    提示信息
 * @apiError (请求失败返回)   {string} internalMsg  错误代码
 * @apiError (请求失败返回)   {float}  timeConsumed 后台耗时

 * @apiErrorExample {json} 失败返回
 * {
 *   "clientMsg": "获取操作日志失败",
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
 * @apiSuccess (data对象字段说明) {string} order_by 排序字段
 * @apiSuccess (data对象字段说明) {int} page 当前页码
 * @apiSuccess (data对象字段说明) {int} size 每页显示行数
 * @apiSuccess (data对象字段说明) {string} sort 排序方式
 * @apiSuccess (data对象字段说明) {int} total 数据总行数
 * @apiSuccess (data对象字段说明) {array} list 日志列表
 * @apiSuccess (data对象子字段 list 说明) {string} id 登录日志主键
 * @apiSuccess (data对象子字段 list 说明) {string} ip 登录IP
 * @apiSuccess (data对象子字段 list 说明) {string} login_city 登录城市
 * @apiSuccess (data对象子字段 list 说明) {string} created 创建时间
 * @apiSuccessExample {json} 响应结果
 * {
 *     "clientMsg": "获取登录日志成功",
 *     "code": 200,
 *     "data": {
 *         "list": [
 *             {
 *                 "account": "test7",
 *                 "id": "1",
 *                 "ip": "127.0.0.1",
 *                 "login_city": "",
 *                 "login_time": "2019-06-13 09:01:48",
 *                 "merchant_name": ""
 *             }
 *         ],
 *         "order_by": "id",
 *         "page": 1,
 *         "size": 20,
 *         "sort": "DESC",
 *         "total": 1
 *     },
 *     "internalMsg": "",
 *     "timeConsumed": 273287
 * }
 */
func (cthis *LogController) OperationLog() {
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

	filterArr := []map[string]string{
		{"prefix": "smol", "key": "merchant_id", "val": sSmUserId},
		{"prefix": "smol", "key": "created", "val": sFromTimestamp, "condition": "gt"},
		{"prefix": "smol", "key": "created", "val": sToTimestamp, "condition": "lt"},
		{"prefix": "smol", "key": "content", "val": postData.Get("content"), "condition": "like_ab"},
		{"prefix": "smol", "key": "ip", "val": postData.Get("ip")},
		{"prefix": "smol", "key": "login_city", "val": postData.Get("login_city"), "condition": "like_ab"},
	}
	where := utils.BuildWhere(filterArr)
	countSql := "SELECT COUNT(*) cnt FROM silver_merchant_os_logs smol LEFT JOIN silver_merchant_users smc ON smc.id=smol.merchant_id" + where
	engine := models.MyEngine[cthis.platform]
	cntRes, cntErr := engine.SQL(countSql).QueryString()
	if cntErr != nil {
		utils.ResFaiJSON(&ctx, "1906141750", "获取操作日志失败", config.NOTGETDATA)
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
		orderBy := " ORDER BY smol.`" + html.EscapeString(sOrderBy) + "` " + html.EscapeString(sSort)
		sql := "SELECT smol.*,smc.merchant_name FROM silver_merchant_os_logs smol LEFT JOIN silver_merchant_users smc ON smc.id=smol.merchant_id" + where + orderBy + limit
		logList, queryErr := engine.SQL(sql).QueryString()
		if queryErr != nil {
			utils.ResFaiJSON(&ctx, "1906141751", "获取操作日志失败", config.NOTGETDATA)
			return
		}
		var list interface{}
		if len(logList) > 0 {
			for k, log := range logList {
				if sCreated, fExist := log["created"]; fExist {
					iCreated, _ := strconv.Atoi(sCreated)
					logList[k]["created"] = utils.TimestampToDateStr(int64(iCreated), "1")
				}
			}
			list = logList
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
		utils.ResSuccJSON(&ctx, "", "获取操作日志成功", config.SUCCESSRES, data)
	} else {
		utils.ResSuccJSON(&ctx, "", "获取操作日志成功", config.SUCCESSRES, map[string]interface{}{
			"list":     make([]map[string]string, 0),
			"page":     iPage,
			"size":     iSize,
			"total":    total,
			"order_by": sOrderBy,
			"sort":     sSort,
		})
	}
}
