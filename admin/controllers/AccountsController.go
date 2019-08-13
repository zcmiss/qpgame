package controllers

import (
	"qpgame/admin/models"
)

var accounts = models.Accounts{} //模型

type AccountsController struct{}

/**
 * @api {get} admin/api/auth/v1/accounts 用户账户列表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>用户账户列表</strong><br />
 * 业务描述: 用户账户列表</br>
 * @apiVersion 1.0.0
 * @apiName     indexAccounts
 * @apiGroup    user
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int} 	page			页数
 * @apiParam (客户端请求参数) {int} 	page_size		每页记录数
 * @apiParam (客户端请求参数) {int} 	user_id 用户编号
 * @apiParam (客户端请求参数) {string} 	time_start 变更时间/开始
 * @apiParam (客户端请求参数) {string} 	time_end 变更时间/结束
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
 * @apiSuccess (返回结果)  {json}  	data            返回数据
 * @apiSuccess (返回结果)  {float}  	timeConsumed    后台耗时
 *
 * @apiSuccess (data字段说明) {array}  	rows        数据列表
 * @apiSuccess (data字段说明) {int}    	page		当前页数
 * @apiSuccess (data字段说明) {int}    	page_count	总的页数
 * @apiSuccess (data字段说明) {int}    	total_rows	总记录数
 * @apiSuccess (data字段说明) {int}    	page_size	每页记录数
 *
 * @apiSuccess (data-rows每个子对象字段说明) {int}		  id 					  记录编号
 * @apiSuccess (data-rows每个子对象字段说明) {int}        user_id                 用户编号
 * @apiSuccess (data-rows每个子对象字段说明) {string}        user_name                 用户名称
 * @apiSuccess (data-rows每个子对象字段说明) {float}      charged_amount          充值总金额
 * @apiSuccess (data-rows每个子对象字段说明) {float}      consumed_amount         消费总金额
 * @apiSuccess (data-rows每个子对象字段说明) {float}      withdraw_amount         提现总金额
 * @apiSuccess (data-rows每个子对象字段说明) {float}      total_bet_amount		  累計打码量
 * @apiSuccess (data-rows每个子对象字段说明) {float}      balance_lucky           总中奖金额
 * @apiSuccess (data-rows每个子对象字段说明) {float}      balance_safe            保险箱余额
 * @apiSuccess (data-rows每个子对象字段说明) {float}      balance_wallet          钱包余额
 * @apiSuccess (data-rows每个子对象字段说明) {int}        updated                 更新时间
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "数据获取成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "user_id": "",
 *        "user_name": "",
 *        "charged_amount": "",
 *        "consumed_amount": "",
 *        "total_bet_amount": "",
 *        "withdraw_amount": "",
 *        "balance_charge": "",
 *        "balance_lucky": "",
 *        "balance_safe": "",
 *        "balance_wallet": "",
 *        "updated": ""
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *AccountsController) Index(ctx *Context) {
	index(ctx, &accounts)
}

/**
 * @api {get} admin/api/auth/v1/accounts/view 				用户账户详情
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>用户账户详情</strong><br />
 * 业务描述: 用户账户详情</br>
 * @apiVersion 1.0.0
 * @apiName     viewAccounts
 * @apiGroup    user
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }

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
 * @apiSuccess (返回结果)  {json}		data            返回数据
 * @apiSuccess (返回结果)  {float}   	timeConsumed    后台耗时
 *
 * @apiSuccess (data字段说明) {int}		  id 					  记录编号
 * @apiSuccess (data字段说明) {int}        user_id                 用户编号
 * @apiSuccess (data字段说明) {float}      charged_amount          充值总金额
 * @apiSuccess (data字段说明) {float}      consumed_amount         消费总金额
 * @apiSuccess (data字段说明) {float}      total_bet_amount	       余额
 * @apiSuccess (data字段说明) {float}      withdraw_amount         提现总金额
 * @apiSuccess (data字段说明) {float}      balance_charge          剩余打码量
 * @apiSuccess (data字段说明) {float}      balance_lucky           总中奖金额
 * @apiSuccess (data字段说明) {float}      balance_safe            保险箱余额
 * @apiSuccess (data字段说明) {float}      balance_wallet          钱包余额
 * @apiSuccess (data字段说明) {int}        updated                 更新时间
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "保存数据成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "user_id": "",
 *        "charged_amount": "",
 *        "consumed_amount": "",
 *        "total_bet_amount": "",
 *        "withdraw_amount": "",
 *        "balance_charge": "",
 *        "balance_lucky": "",
 *        "balance_safe": "",
 *        "balance_wallet": "",
 *        "updated": ""
 * 	  },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *AccountsController) View(ctx *Context) {
	view(ctx, &accounts)
}

/**
 * @api {get} admin/api/auth/v1/accounts/delete 用户账户删除
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>用户账户删除</strong><br />
 * 业务描述: 用户账户删除</br>
 * @apiVersion 1.0.0
 * @apiName     deleteAccounts
 * @apiGroup    user
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }

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
 * @apiSuccess (返回结果)  {json}		data            返回数据
 * @apiSuccess (返回结果)  {float}   	timeConsumed   	后台耗时
 *
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
func (self *AccountsController) Delete(ctx *Context) {
	remove(ctx, &accounts)
}
