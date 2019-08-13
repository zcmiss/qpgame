package controllers

import (
	"qpgame/admin/models"
)

var accountStatistics = models.AccountStatistics{} //模型

type AccountStatisticsController struct{}

/**
 * @api {get} admin/api/auth/v1/account_statistics 用户输羸排名
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>账户统计列表</strong><br />
 * 业务描述: 账户统计列表</br>
 * @apiVersion 1.0.0
 * @apiName     indexAccountStatistic
 * @apiGroup    report
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int} 	page			页数
 * @apiParam (客户端请求参数) {string} 	page_size	每页记录数
 * @apiParam (客户端请求参数) {string} 	ymd_start 	统计日期/开始, 默认当天
 * @apiParam (客户端请求参数) {string} 	ymd_end 	统计日期/开始, 默认当天 <br /><span style="color:red">特别注意: 请额外提供 昨天, 今天, 上周, 本周, 上月, 本月 等搜索按钮</span>
 * @apiParam (客户端请求参数) {string} 	sort_type	排序方式，下拉列表, 可选项有以下几种:<br />
 *"charged_amount": 充值金额 <br />
 * "consumed_amount": 消费金额 <br />
 * "withdraw_amount": 提现金额 <br />
 * "bet_amount": 投注金额, 默认<br />
 * "reward_amount": 中奖金额<br />
 * "wash_amount": 洗码金额<br />
 * "proxy_commission": 代理佣金<br />
 *
 * @apiError (请求失败返回) {int}      code            错误代码
 * @apiError (请求失败返回) {string}   clientMsg       提示信息
 * @apiError (请求失败返回) {string}   internalMsg     内部错误信息
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
 * @apiSuccess (返回结果)  {int}    	code            200, 成功
 * @apiSuccess (返回结果)  {string} 	clientMsg       提示信息
 * @apiSuccess (返回结果)  {string} 	internalMsg     内部错误信息
 * @apiSuccess (返回结果)  {json}	data        	返回数据
 * @apiSuccess (返回结果)  {float}  	timeConsumed    后台耗时
 *
 * @apiSuccess (data字段说明) {array}  	rows        数据列表
 * @apiSuccess (data字段说明) {int}    	page		当前页数
 * @apiSuccess (data字段说明) {int}    	page_count	总的页数
 * @apiSuccess (data字段说明) {int}    	total_rows	总记录数
 * @apiSuccess (data字段说明) {int}    	page_size	每页记录数
 *
 * @apiSuccess (data-rows每个子对象字段说明) {int}		id 					编号/排序, 从1开始，表示第几, 如1表示第一
 * @apiSuccess (data-rows每个子对象字段说明) {int}        user_id 			用户编号
 * @apiSuccess (data-rows每个子对象字段说明) {string}        user_name 			用户名称
 * @apiSuccess (data-rows每个子对象字段说明) {int}        charge_amount 		充值金额
 * @apiSuccess (data-rows每个子对象字段说明) {int}        consumed_amount		消费金额
 * @apiSuccess (data-rows每个子对象字段说明) {int}        withdraw_amount		提现金额
 * @apiSuccess (data-rows每个子对象字段说明) {float}      bet_amount			投注金额
 * @apiSuccess (data-rows每个子对象字段说明) {int}        reward_amount		中奖金额
 * @apiSuccess (data-rows每个子对象字段说明) {float}      wash_amount			洗码金额
 * @apiSuccess (data-rows每个子对象字段说明) {int}        proxy_commission 	代理佣金
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "数据获取成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "user_id": "",
 *        "user_name": "",
 *        "charge_amount": "",
 *        "consumed_amount": "",
 *        "withdraw_amount": "",
 *        "bet_amount": "",
 *        "reward_amount": "",
 *        "wash_amount": "",
 *        "proxy_commission": "",
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *AccountStatisticsController) Index(ctx *Context) {
	index(ctx, &accountStatistics)
}
