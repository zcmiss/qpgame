package controllers

import (
	"qpgame/admin/models"
)

var withdrawDamaRecords = models.WithdrawDamaRecords{} //模型

type WithdrawDamaRecordsController struct{}

/**
 * @api {get} admin/api/auth/v1/withdraw_dama_records 提现打码列表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>提现打码列表</strong><br />
 * 业务描述: 提现打码列表</br>
 * @apiVersion 1.0.0
 * @apiName     indexWithdrawDamaRecords
 * @apiGroup    finance
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int} 	page			页数
 * @apiParam (客户端请求参数) {int} 	page_size		每页记录数
 * @apiParam (客户端请求参数) {int} 	user_id 用户编号
 * @apiParam (客户端请求参数) {int} 	fund_type 资金类型/参照下面
 * @apiParam (客户端请求参数) {int} 	state 完成打码状态/0:未完成/1:已完成
 * @apiParam (客户端请求参数) {string} 	created_start 流水时间/开始
 * @apiParam (客户端请求参数) {string} 	created_end 流水时间/结束
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
 * @apiSuccess (data-rows每个子对象字段说明) {int}		  id 	记录编号
 * @apiSuccess (data-rows每个子对象字段说明) {int}        user_id                 用户ID
 * @apiSuccess (data-rows每个子对象字段说明) {string}        user_name                 用户名称
 * @apiSuccess (data-rows每个子对象字段说明) {float}      amount                  原始金额
 * @apiSuccess (data-rows每个子对象字段说明) {int}        fund_type               1.充值,3.洗码,5.赠送彩金,6.优惠入款,9.活动奖励,14.红包收入
 * @apiSuccess (data-rows每个子对象字段说明) {string}        fund_type_name               资金类型(说明文字)
 * @apiSuccess (data-rows每个子对象字段说明) {float}      finish_rate             打码量比例,需要原始资金的多少倍才算完成
 * @apiSuccess (data-rows每个子对象字段说明) {int}        updated                 修改时间
 * @apiSuccess (data-rows每个子对象字段说明) {int}        created                 创建时间
 * @apiSuccess (data-rows每个子对象字段说明) {float}      finished_progress       完成的资金量，产生一次流水都要更新这个金额,也就是实际打码量
 * @apiSuccess (data-rows每个子对象字段说明) {float}      finished_needed         打满量完成金额,打码量比例乘以原始金额
 * @apiSuccess (data-rows每个子对象字段说明) {int}        state                   打码状态, 0:未完成,1:已完成,2:已失效
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "数据获取成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "user_id": "",
 *        "user_name": "",
 *        "amount": "",
 *        "fund_type": "",
 *        "finish_rate": "",
 *        "updated": "",
 *        "created": "",
 *        "finished_progress": "",
 *        "finished_needed": "",
 *        "state": ""
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *WithdrawDamaRecordsController) Index(ctx *Context) {
	index(ctx, &withdrawDamaRecords)
}

/**
 * @api {get} admin/api/auth/v1/withdraw_dama_records/view 				提现打码详情
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>提现打码详情</strong><br />
 * 业务描述: 提现打码详情</br>
 * @apiVersion 1.0.0
 * @apiName     viewWithdrawDamaRecords
 * @apiGroup    finance
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
 * @apiSuccess (data字段说明) {int}		  id 					  记录编号
 * @apiSuccess (data字段说明) {int}        user_id                 用户ID
 * @apiSuccess (data字段说明) {float}      amount                  原始金额
 * @apiSuccess (data字段说明) {int}        fund_type               1.充值,3.洗码,5.赠送彩金,6.优惠入款,9.活动奖励,14.红包收入
 * @apiSuccess (data字段说明) {float}      finish_rate             打码量比例,需要原始资金的多少倍才算完成
 * @apiSuccess (data字段说明) {int}        created                 创建时间
 * @apiSuccess (data字段说明) {float}      finished_progress       完成的资金量，产生一次流水都要更新这个金额,也就是实际打码量
 * @apiSuccess (data字段说明) {float}      finished_needed         打满量完成金额,打码量比例乘以原始金额
 * @apiSuccess (data字段说明) {int}        state                   打码状态, 0:未完成,1:已完成,2:已失效
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "保存数据成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "user_id": "",
 *        "amount": "",
 *        "fund_type": "",
 *        "finish_rate": "",
 *        "updated": "",
 *        "created": "",
 *        "finished_progress": "",
 *        "finished_needed": "",
 *        "state": ""
 * 	  },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *WithdrawDamaRecordsController) View(ctx *Context) {
	view(ctx, &withdrawDamaRecords)
}

/**
 * @api {get} admin/api/auth/v1/withdraw_dama_records/delete 提现打码删除
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>提现打码删除</strong><br />
 * 业务描述: 提现打码删除</br>
 * @apiVersion 1.0.0
 * @apiName     deleteWithdrawDamaRecords
 * @apiGroup    finance
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
func (self *WithdrawDamaRecordsController) Delete(ctx *Context) {
	remove(ctx, &withdrawDamaRecords)
}
