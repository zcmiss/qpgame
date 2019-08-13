package controllers

import (
	"qpgame/admin/models"
)

var accountInfos = models.AccountInfos{} //模型

type AccountInfosController struct{}

/**
 * @api {get} admin/api/auth/v1/account_infos 会员账户资金列表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>会员账户资金信息列表</strong><br />
 * 业务描述: 会员账户资金信息列表</br>
 * @apiVersion 1.0.0
 * @apiName     indexAccountInfos
 * @apiGroup    user
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int} 	page			页数
 * @apiParam (客户端请求参数) {int} 	page_size		每页记录数
 * @apiParam (客户端请求参数) {string} 	user_id 	用户编号或名称
 * @apiParam (客户端请求参数) {string} 	created_start	交易时间/开始
 * @apiParam (客户端请求参数) {string} 	created_end 	交易时间/结束
 * @apiParam (客户端请求参数) {int} 	type 交易类型
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
 * @apiSuccess (data-rows每个子对象字段说明) {int}		  id 					  记录编号
 * @apiSuccess (data-rows每个子对象字段说明) {string}        order_id	订单编号
 * @apiSuccess (data-rows每个子对象字段说明) {int}        user_id                 用户id
 * @apiSuccess (data-rows每个子对象字段说明) {string}        user_name                 用户名称
 * @apiSuccess (data-rows每个子对象字段说明) {float}      amount                  金额 正数为收入，负数为支出
 * @apiSuccess (data-rows每个子对象字段说明) {float}      balance                 变化后的余额
 * @apiSuccess (data-rows每个子对象字段说明) {int}        type                    交易类型
 * @apiSuccess (data-rows每个子对象字段说明) {int}        created                 创建时间
 * @apiSuccess (data-rows每个子对象字段说明) {string}        msg                 流水说明
 * @apiSuccess (data-rows每个子对象字段说明) {float}      charged_amount_old       账变前充值总额
 * @apiSuccess (data-rows每个子对象字段说明) {float}      charged_amount           账变后充值总额
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
 *        "balance": "",
 *        "type": "",
 *        "created": "",
 *        "charged_amount_old": "",
 *        "charged_amount": "",
 *        "msg": "",
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *AccountInfosController) Index(ctx *Context) {
	index(ctx, &accountInfos)
}

/**
 * @api {get} admin/api/auth/v1/account_infos/view 				会员账户资金信息详情
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>账户信息详情</strong><br />
 * 业务描述: 账户信息详情</br>
 * @apiVersion 1.0.0
 * @apiName     viewAccountInfos
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
 * @apiSuccess (data字段说明) {int}        user_id                 用户id
 * @apiSuccess (data字段说明) {float}      amount                  金额 正数为收入，负数为支出
 * @apiSuccess (data字段说明) {float}      balance                 变化后的余额
 * @apiSuccess (data字段说明) {int}        type                    交易类型
 * @apiSuccess (data字段说明) {int}        created                 创建时间
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "保存数据成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "user_id": "",
 *        "amount": "",
 *        "balance": "",
 *        "type": "",
 *        "created": ""
 * 	  },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *AccountInfosController) View(ctx *Context) {
	view(ctx, &accountInfos)
}
