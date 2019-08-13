package controllers

import (
	"qpgame/admin/models"
)

var proxyStatistics = models.ProxyStatistics{} //模型

type ProxyStatisticsController struct{}

/**
 * @api {get} admin/api/auth/v1/proxy_statistics 代理统计列表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>代理统计列表</strong><br />
 * 业务描述: 代理统计列表</br>
 * @apiVersion 1.0.0
 * @apiName     indexProxyStatistics
 * @apiGroup    report
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int}    page            页数
 * @apiParam (客户端请求参数) {int}   page_size       每页记录数
 * @apiParam (客户端请求参数) {string}   ymd_start 日期/开始
 * @apiParam (客户端请求参数) {string}   ymd_end 日期/结束
 * @apiParam (客户端请求参数) {string}   user_id 用户编号
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
 * @apiSuccess (返回结果)  {json}		data            返回数据
 * @apiSuccess (返回结果)  {float}  	timeConsumed    后台耗时
 *
 * @apiSuccess (data字段说明) {array}  	rows        数据列表
 * @apiSuccess (data字段说明) {int}    	page		当前页数
 * @apiSuccess (data字段说明) {int}    	page_count	总的页数
 * @apiSuccess (data字段说明) {int}    	total_rows	总记录数
 * @apiSuccess (data字段说明) {int}    	page_size	每页记录数
 *
 * @apiSuccess (data-rows每个子对象字段说明) {int}		  id					编号
 * @apiSuccess (data-rows每个子对象字段说明) {int}        ymd                     统计日期
 * @apiSuccess (data-rows每个子对象字段说明) {int}        user_id                 用户编号
 * @apiSuccess (data-rows每个子对象字段说明) {string}        user_name                 用户名称
 * @apiSuccess (data-rows每个子对象字段说明) {int}        level                   层级
 * @apiSuccess (data-rows每个子对象字段说明) {float}      charge                  充值总额
 * @apiSuccess (data-rows每个子对象字段说明) {float}      withdraw                提现总额
 * @apiSuccess (data-rows每个子对象字段说明) {float}      deductions              扣除总额
 * @apiSuccess (data-rows每个子对象字段说明) {float}      bet_amount              下注总额
 * @apiSuccess (data-rows每个子对象字段说明) {int}        bet_count               下注总数
 * @apiSuccess (data-rows每个子对象字段说明) {int}        charge_count            充值次数
 * @apiSuccess (data-rows每个子对象字段说明) {int}        charge_user_count       充值人数
 * @apiSuccess (data-rows每个子对象字段说明) {int}        withdraw_count          提现次数
 * @apiSuccess (data-rows每个子对象字段说明) {int}        withdraw_user_count     提现人数
 * @apiSuccess (data-rows每个子对象字段说明) {int}        sale_ratio              销售返点
 * @apiSuccess (data-rows每个子对象字段说明) {float}      winning                 中奖金额
 * @apiSuccess (data-rows每个子对象字段说明) {float}      proxy_ratio             代理返点
 * @apiSuccess (data-rows每个子对象字段说明) {float}      active                  活动奖励
 * @apiSuccess (data-rows每个子对象字段说明) {float}      user_win                团队盈亏
 * @apiSuccess (data-rows每个子对象字段说明) {float}      give_win                派彩损益
 * @apiSuccess (data-rows每个子对象字段说明) {float}      real_win                实际盈亏
 * @apiSuccess (data-rows每个子对象字段说明) {int}        first_charge            首充人数
 * @apiSuccess (data-rows每个子对象字段说明) {float}      first_charge_amount     首充金额
 * @apiSuccess (data-rows每个子对象字段说明) {int}        proxy_count             代理总数
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "数据获取成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "ymd": "",
 *        "user_id": "",
 *        "user_name": "",
 *        "level": "",
 *        "charge": "",
 *        "withdraw": "",
 *        "deductions": "",
 *        "bet_amount": "",
 *        "bet_count": "",
 *        "charge_count": "",
 *        "charge_user_count": "",
 *        "withdraw_count": "",
 *        "withdraw_user_count": "",
 *        "sale_ratio": "",
 *        "winning": "",
 *        "proxy_ratio": "",
 *        "active": "",
 *        "user_win": "",
 *        "give_win": "",
 *        "real_win": "",
 *        "profit": "",
 *        "reg_user": "",
 *        "bet_new": "",
 *        "deposit_user": "",
 *        "first_charge": "",
 *        "first_charge_amount": "",
 *        "proxy_count": "",
 *        "member_user": "",
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *ProxyStatisticsController) Index(ctx *Context) {
	index(ctx, &proxyStatistics)
}

/**
 * @api {get} admin/api/v1/proxy_statistics/view 				代理统计详情
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>代理统计详情</strong><br />
 * 业务描述: 代理统计详情</br>
 * @apiVersion 1.0.0
 * @apiName     viewProxyStatistics
 * @apiGroup    report
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 *
 * @apiParam (客户端请求参数) {int} 	id    			编号
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
 *      "clientMsg ": "",
 *      "timeConsumed": 0
 * }
 *
 * @apiSuccess (返回结果)  {int} 		code            200
 * @apiSuccess (返回结果)  {string} 	clientMsg       提示信息
 * @apiSuccess (返回结果)  {string}  	internalMsg     内部信息
 * @apiSuccess (返回结果)  {json}  	data            返回数据
 * @apiSuccess (返回结果)  {float}   	timeConsumed    后台耗时
 *
 * @apiSuccess (data字段说明) {int}		  id 					  编号
 * @apiSuccess (data字段说明) {int}        ymd                     统计日期
 * @apiSuccess (data字段说明) {int}        user_id                 用户编号
 * @apiSuccess (data字段说明) {int}        level                   层级
 * @apiSuccess (data字段说明) {float}      charge                  充值总额
 * @apiSuccess (data字段说明) {float}      withdraw                提现总额
 * @apiSuccess (data字段说明) {float}      deductions              扣除总额
 * @apiSuccess (data字段说明) {float}      bet_amount              下注总额
 * @apiSuccess (data字段说明) {int}        bet_count               下注总数
 * @apiSuccess (data字段说明) {int}        charge_count            充值次数
 * @apiSuccess (data字段说明) {int}        charge_user_count       充值人数
 * @apiSuccess (data字段说明) {int}        withdraw_count          提现次数
 * @apiSuccess (data字段说明) {int}        withdraw_user_count     提现人数
 * @apiSuccess (data字段说明) {int}        sale_ratio              销售返点
 * @apiSuccess (data字段说明) {float}      winning                 中奖金额
 * @apiSuccess (data字段说明) {float}      proxy_ratio             代理返点
 * @apiSuccess (data字段说明) {float}      user_win                团队盈亏
 * @apiSuccess (data字段说明) {int}        first_charge            首充人数
 * @apiSuccess (data字段说明) {float}      first_charge_amount     首充金额
 * @apiSuccess (data字段说明) {int}        proxy_count             代理总数
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "保存数据成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "ymd": "",
 *        "user_id": "",
 *        "level": "",
 *        "charge": "",
 *        "withdraw": "",
 *        "deductions": "",
 *        "bet_amount": "",
 *        "bet_count": "",
 *        "charge_count": "",
 *        "charge_user_count": "",
 *        "withdraw_count": "",
 *        "withdraw_user_count": "",
 *        "sale_ratio": "",
 *        "winning": "",
 *        "proxy_ratio": "",
 *        "active": "",
 *        "user_win": "",
 *        "give_win": "",
 *        "real_win": "",
 *        "profit": "",
 *        "reg_user": "",
 *        "bet_new": "",
 *        "deposit_user": "",
 *        "first_charge": "",
 *        "first_charge_amount": "",
 *        "proxy_count": "",
 *        "downline_bet_user": "",
 *        "member_user": "",
 *        "created": ""
 * 	  },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *ProxyStatisticsController) View(ctx *Context) {
	view(ctx, &proxyStatistics)
}

/**
 * @api {get} admin/api/v1/proxy_statistics/delete 代理统计删除
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>代理统计删除</strong><br />
 * 业务描述: 代理统计删除</br>
 * @apiVersion 1.0.0
 * @apiName     deleteProxyStatistics
 * @apiGroup    report
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 *
 * @apiParam (客户端请求参数) {int} 	id    			记录编号
 *
 * @apiError (请求失败返回) {int}      code            错误代码
 * @apiError (请求失败返回) {string}   clientMsg       提示信息
 * @apiError (请求失败返回) {string}   internalMsg     内部错误信息
 * @apiError (请求失败返回) {float}    timeConsumed   	 后台耗时
 *
 * @apiErrorExample {json} 失败返回
 * {
 *      "code": 204,
 *      "internalMsg": "",
 *      "clientMsg ": "",
 *      "timeConsumed": 0
 * }
 *
 * @apiSuccess (返回结果)  {int} 		code            200
 * @apiSuccess (返回结果)  {string} 	clientMsg       提示信息
 * @apiSuccess (返回结果)  {string}  	internalMsg     内部信息
 * @apiSuccess (返回结果)  {json}  	data            返回数据
 * @apiSuccess (返回结果)  {float}   	timeConsumed   	后台耗时
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "记录删除成功",
 *    "code": 200,
 *    "data": {},
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *ProxyStatisticsController) Delete(ctx *Context) {
	remove(ctx, &proxyStatistics)
}
