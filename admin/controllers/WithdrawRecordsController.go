package controllers

import (
	"qpgame/admin/models"
)

var withdrawRecords = models.WithdrawRecords{} //模型

type WithdrawRecordsController struct{}

/**
 * @api {get} admin/api/auth/v1/withdraw_records 提现申请列表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>提现申请列表</strong><br />
 * 业务描述: 提现申请列表</br>
 * @apiVersion 1.0.0
 * @apiName     indexWithdrawRecords
 * @apiGroup    finance
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int}    page            页数
 * @apiParam (客户端请求参数) {int}   page_size       每页记录数
 * @apiParam (客户端请求参数) {string}   user_id 	用户编号或名称
 * @apiParam (客户端请求参数) {string}   real_name 	真实姓名
 * @apiParam (客户端请求参数) {string}   order_id	订单号码
 * @apiParam (客户端请求参数) {string}   created_start       申请时间/开始
 * @apiParam (客户端请求参数) {string}   created_end       申请时间/结束
 * @apiParam (客户端请求参数) {int}   status       状态
 * @apiParam (客户端请求参数) {string}   card_number 卡号
 * @apiParam (客户端请求参数) {string}   operator 操作人
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
 * @apiSuccess (data-rows每个子对象字段说明) {int}        user_id                 用户编号
 * @apiSuccess (data-rows每个子对象字段说明) {string}        user_name                 用户名称
 * @apiSuccess (data-rows每个子对象字段说明) {string}     order_id                提现订单号
 * @apiSuccess (data-rows每个子对象字段说明) {float}      amount                  提现金额
 * @apiSuccess (data-rows每个子对象字段说明) {string}     real_name               真实姓名
 * @apiSuccess (data-rows每个子对象字段说明) {string}     bank_name               银行名称
 * @apiSuccess (data-rows每个子对象字段说明) {string}     card_number             银行卡号
 * @apiSuccess (data-rows每个子对象字段说明) {string}     bank_address            银行卡地址
 * @apiSuccess (data-rows每个子对象字段说明) {string}     withdraw_type           提现类型比如:在线提款
 * @apiSuccess (data-rows每个子对象字段说明) {int}        status                  0 待审核,1 提现成功,2 退回出款 3 锁定
 * @apiSuccess (data-rows每个子对象字段说明) {int}        created             申请时间
 * @apiSuccess (data-rows每个子对象字段说明) {string}     remark                  备注
 * @apiSuccess (data-rows每个子对象字段说明) {string}     operator 	操作者
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
 *        "order_id": "",
 *        "updated": "",
 *        "status": "",
 *        "created": "",
 *        "card_number": "",
 *        "real_name": "",
 *        "bank_address": "",
 *        "bank_name": "",
 *        "withdraw_type": "",
 *        "remark": "",
 *        "operator": ""
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *WithdrawRecordsController) Index(ctx *Context) {
	index(ctx, &withdrawRecords)
}

/**
 * @api {get} admin/api/auth/v1/withdraw_records/view 				提现申请详情
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>提现申请详情</strong><br />
 * 业务描述: 提现申请详情</br>
 * @apiVersion 1.0.0
 * @apiName     viewWithdrawRecords
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
 * @apiSuccess (data字段说明) {int}        user_id                 关联用户表id
 * @apiSuccess (data字段说明) {float}      amount                  提现金额
 * @apiSuccess (data字段说明) {string}     order_id                提现订单号
 * @apiSuccess (data字段说明) {int}        updated             更新时间
 * @apiSuccess (data字段说明) {int}        status                  0 待审核,1 提现成功,2 退回出款 3 锁定
 * @apiSuccess (data字段说明) {int}        created             创建时间
 * @apiSuccess (data字段说明) {string}     card_number             银行卡号
 * @apiSuccess (data字段说明) {string}     real_name               真实姓名
 * @apiSuccess (data字段说明) {string}     bank_address            银行卡地址
 * @apiSuccess (data字段说明) {string}     bank_name               银行名称
 * @apiSuccess (data字段说明) {string}     withdraw_type           提现类型比如:在线提款
 * @apiSuccess (data字段说明) {string}     remark                  备注
 * @apiSuccess (data字段说明) {string}     operator 	操作者
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "保存数据成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "user_id": "",
 *        "amount": "",
 *        "order_id": "",
 *        "updated": "",
 *        "status": "",
 *        "created": "",
 *        "card_number": "",
 *        "real_name": "",
 *        "bank_address": "",
 *        "bank_name": "",
 *        "withdraw_type": "",
 *        "remark": "",
 *        "auther": ""
 * 	  },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *WithdrawRecordsController) View(ctx *Context) {
	view(ctx, &withdrawRecords)
}

/**
 * @api {get} admin/api/auth/v1/withdraw_records/delete 提现申请删除
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>提现申请删除</strong><br />
 * 业务描述: 提现申请删除</br>
 * @apiVersion 1.0.0
 * @apiName     deleteWithdrawRecords
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
func (self *WithdrawRecordsController) Delete(ctx *Context) {
	remove(ctx, &withdrawRecords)
}

/**
 * @api {get} admin/api/auth/v1/withdraw_records/allow	提现审核通过
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>提现审核通过</strong><br />
 * 业务描述: 提现审核通过</br>
 * @apiVersion 1.0.0
 * @apiName     viewWithdrawRecordsAllow
 * @apiGroup    finance
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }

 * @apiParam (客户端请求参数) {int} 	id    			充值记录编号
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
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "保存数据成功",
 *    "code": 200,
 *    "data": {},
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *WithdrawRecordsController) Allow(ctx *Context) {
	responseResult(ctx, withdrawRecords.Allow(ctx), "提现审核处理成功")
}

/**
 * @api {get} admin/api/auth/v1/withdraw_records/deny	提现申请拒绝
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>提现审核拒绝</strong><br />
 * 业务描述: 提现审核拒绝</br>
 * @apiVersion 1.0.0
 * @apiName     withdrawRecordsDeny
 * @apiGroup    finance
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }

 * @apiParam (客户端请求参数) {int} 	id    			充值记录编号
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
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "保存数据成功",
 *    "code": 200,
 *    "data": {},
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *WithdrawRecordsController) Deny(ctx *Context) {
	responseResult(ctx, withdrawRecords.Deny(ctx), "已拒绝用户提现申请")
}
