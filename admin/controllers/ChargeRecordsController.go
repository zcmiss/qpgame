package controllers

import (
	"qpgame/admin/models"
)

var chargeRecords = models.ChargeRecords{} //模型

type ChargeRecordsController struct{}

/**
 * @api {get} admin/api/auth/v1/charge_records 会员公司入款
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>充值记录列表</strong><br />
 * 业务描述: 充值记录列表</br>
 * @apiVersion 1.0.0
 * @apiName     indexChargeRecords
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
 * @apiParam (客户端请求参数) {string} 	order_id 订单编号
 * @apiParam (客户端请求参数) {int} 	charge_type_id 充值方式
 * @apiParam (客户端请求参数) {string} 	card_number 卡号
 * @apiParam (客户端请求参数) {string} 	created_start 充值时间/开始
 * @apiParam (客户端请求参数) {string} 	created_end  充值时间/结束
 * @apiParam (客户端请求参数) {string} 	credential_id  第三方支付编号
 * @apiParam (客户端请求参数) {string} 	is_tppay 是否第三方
 * @apiParam (客户端请求参数) {string} 	real_name  真实姓名
 * @apiParam (客户端请求参数) {string} 	ip  IP
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
 * @apiSuccess (data-rows每个子对象字段说明) {int}		  id 					记录编号
 * @apiSuccess (data-rows每个子对象字段说明) {int}        user_id                 用户ID
 * @apiSuccess (data-rows每个子对象字段说明) {string}        user_name                 用户名称
 * @apiSuccess (data-rows每个子对象字段说明) {float}      amount                  充值金额
 * @apiSuccess (data-rows每个子对象字段说明) {string}     order_id                充值订单
 * @apiSuccess (data-rows每个子对象字段说明) {int}        charge_type_id          充值方式id
 * @apiSuccess (data-rows每个子对象字段说明) {string}     card_number             卡号
 * @apiSuccess (data-rows每个子对象字段说明) {string}     bank_address            开户银行地址或支付二维码
 * @apiSuccess (data-rows每个子对象字段说明) {int}        created                 添加时间
 * @apiSuccess (data-rows每个子对象字段说明) {int}        state                   公司入款：0 待审核，1 成功，2 失败。线上支付：0待处理，1成功，2失败，3进行中,4退款，5取消，6强制入款
 * @apiSuccess (data-rows每个子对象字段说明) {string}     charge_type             充值类型
 * @apiSuccess (data-rows每个子对象字段说明) {string}     ip                      充值IP
 * @apiSuccess (data-rows每个子对象字段说明) {int}        platform_id             充值platform
 * @apiSuccess (data-rows每个子对象字段说明) {string}     real_name               真实姓名
 * @apiSuccess (data-rows每个子对象字段说明) {int}        bank_type_id            银行转账类型
 * @apiSuccess (data-rows每个子对象字段说明) {int}        bank_charge_time        银行转账时间
 * @apiSuccess (data-rows每个子对象字段说明) {int}        credential_id           第三方支付记录ID
 * @apiSuccess (data-rows每个子对象字段说明) {string}     operator                操作者
 * @apiSuccess (data-rows每个子对象字段说明) {int}        is_tppay                是否第三方支付 0为否；1为是
 * @apiSuccess (data-rows每个子对象字段说明) {int}        charge_card_id          收款银行卡编号
 * @apiSuccess (data-rows每个子对象字段说明) {string}        charge_card_name          收款银行卡信息
 * @apiSuccess (data-rows每个子对象字段说明) {string}     remark                  备注
 * @apiSuccess (data-rows每个子对象字段说明) {string}     updated_last            最后更新时间
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
 *        "charge_type_id": "",
 *        "card_number": "",
 *        "bank_address": "",
 *        "created": "",
 *        "state": "",
 *        "charge_type": "",
 *        "ip": "",
 *        "platform_id": "",
 *        "real_name": "",
 *        "bank_type_id": "",
 *        "bank_charge_time": "",
 *        "credential_id": "",
 *        "operator": "",
 *        "is_tppay": "",
 *        "charge_card_id": "",
 *        "charge_card_name": "",
 *        "remark": "",
 *        "updated_last": ""
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *ChargeRecordsController) Index(ctx *Context) {
	index(ctx, &chargeRecords)
}

/**
 * @api {get} admin/api/auth/v1/charge_records/view 				会员公司入款详情
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>充值记录详情</strong><br />
 * 业务描述: 充值记录详情</br>
 * @apiVersion 1.0.0
 * @apiName     viewChargeRecords
 * @apiGroup    finance
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
 * @apiSuccess (data字段说明) {int}        user_id                 用户ID
 * @apiSuccess (data字段说明) {float}      amount                  充值金额
 * @apiSuccess (data字段说明) {string}     order_id                充值订单
 * @apiSuccess (data字段说明) {int}        charge_type_id          充值方式id
 * @apiSuccess (data字段说明) {string}     card_number             卡号
 * @apiSuccess (data字段说明) {string}     bank_address            开户银行地址或支付二维码
 * @apiSuccess (data字段说明) {int}        state                   公司入款：0 待审核，1 成功，2 失败。线上支付：0待处理，1成功，2失败，3进行中,4退款，5取消，6强制入款
 * @apiSuccess (data字段说明) {string}     screenshot              屏幕截图
 * @apiSuccess (data字段说明) {string}     receipt_screenshot      收据截图
 * @apiSuccess (data字段说明) {string}     charge_type             充值类型
 * @apiSuccess (data字段说明) {string}     ip                      充值IP
 * @apiSuccess (data字段说明) {int}        platform_id             充值platform
 * @apiSuccess (data字段说明) {string}     real_name               真实姓名
 * @apiSuccess (data字段说明) {int}        bank_type_id            银行转账类型
 * @apiSuccess (data字段说明) {int}        bank_charge_time        银行转账时间
 * @apiSuccess (data字段说明) {int}        credential_id           第三方支付记录ID
 * @apiSuccess (data字段说明) {string}     operator                操作者
 * @apiSuccess (data字段说明) {int}        is_tppay                是否第三方支付 0为否；1为是
 * @apiSuccess (data字段说明) {int}        charge_card_id          关联ChargeBankCards.id
 * @apiSuccess (data字段说明) {string}     remark                  备注
 * @apiSuccess (data字段说明) {string}     updated_last            最后更新时间
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
 *        "charge_type_id": "",
 *        "card_number": "",
 *        "bank_address": "",
 *        "state": "",
 *        "screenshot": "",
 *        "receipt_screenshot": "",
 *        "charge_type": "",
 *        "ip": "",
 *        "platform_id": "",
 *        "real_name": "",
 *        "bank_type_id": "",
 *        "bank_charge_time": "",
 *        "credential_id": "",
 *        "operator": "",
 *        "is_tppay": "",
 *        "charge_card_id": "",
 *        "remark": "",
 *        "updated_last": ""
 * 	  },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *ChargeRecordsController) View(ctx *Context) {
	view(ctx, &chargeRecords)
}

/**
 * @api {get} admin/api/auth/v1/charge_records/delete 会员公司入款删除
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>充值记录删除</strong><br />
 * 业务描述: 充值记录删除</br>
 * @apiVersion 1.0.0
 * @apiName     deleteChargeRecords
 * @apiGroup    finance
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
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "记录删除成功",
 *    "code": 200,
 *    "data": {},
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *ChargeRecordsController) Delete(ctx *Context) {
	remove(ctx, &chargeRecords)
}

/**
 * @api {get} admin/api/auth/v1/charge_records/onlines 会员线上入款
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>会员线上入款</strong><br />
 * 业务描述: 会员线上入款</br>
 * @apiVersion 1.0.0
 * @apiName     indexChargeRecordsOnlines
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
 * @apiParam (客户端请求参数) {string} 	order_id 订单编号
 * @apiParam (客户端请求参数) {int} 	charge_type_id 充值方式
 * @apiParam (客户端请求参数) {string} 	card_number 卡号
 * @apiParam (客户端请求参数) {string} 	created_start 充值时间/开始
 * @apiParam (客户端请求参数) {string} 	created_end  充值时间/结束
 * @apiParam (客户端请求参数) {string} 	credential_id  第三方支付编号
 * @apiParam (客户端请求参数) {string} 	is_tppay 是否第三方
 * @apiParam (客户端请求参数) {string} 	real_name  真实姓名
 * @apiParam (客户端请求参数) {string} 	ip  IP
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
 * @apiSuccess (data-rows每个子对象字段说明) {int}		  id 					记录编号
 * @apiSuccess (data-rows每个子对象字段说明) {int}        user_id                 用户ID
 * @apiSuccess (data-rows每个子对象字段说明) {string}        user_name                 用户名称
 * @apiSuccess (data-rows每个子对象字段说明) {float}      amount                  充值金额
 * @apiSuccess (data-rows每个子对象字段说明) {string}     order_id                充值订单
 * @apiSuccess (data-rows每个子对象字段说明) {int}        charge_type_id          充值方式id
 * @apiSuccess (data-rows每个子对象字段说明) {string}     card_number             卡号
 * @apiSuccess (data-rows每个子对象字段说明) {string}     bank_address            开户银行地址或支付二维码
 * @apiSuccess (data-rows每个子对象字段说明) {int}        created                 添加时间
 * @apiSuccess (data-rows每个子对象字段说明) {int}        state                   公司入款：0 待审核，1 成功，2 失败。线上支付：0待处理，1成功，2失败，3进行中,4退款，5取消，6强制入款
 * @apiSuccess (data-rows每个子对象字段说明) {string}     charge_type             充值类型
 * @apiSuccess (data-rows每个子对象字段说明) {string}     ip                      充值IP
 * @apiSuccess (data-rows每个子对象字段说明) {int}        platform_id             充值platform
 * @apiSuccess (data-rows每个子对象字段说明) {string}     real_name               真实姓名
 * @apiSuccess (data-rows每个子对象字段说明) {int}        bank_type_id            银行转账类型
 * @apiSuccess (data-rows每个子对象字段说明) {int}        bank_charge_time        银行转账时间
 * @apiSuccess (data-rows每个子对象字段说明) {int}        credential_id           第三方支付记录ID
 * @apiSuccess (data-rows每个子对象字段说明) {string}     operator                操作者
 * @apiSuccess (data-rows每个子对象字段说明) {int}        is_tppay                是否第三方支付 0为否；1为是
 * @apiSuccess (data-rows每个子对象字段说明) {int}        charge_card_id          关联ChargeBankCards.id
 * @apiSuccess (data-rows每个子对象字段说明) {string}     remark                  备注
 * @apiSuccess (data-rows每个子对象字段说明) {string}     updated_last            最后更新时间
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
 *        "charge_type_id": "",
 *        "card_number": "",
 *        "bank_address": "",
 *        "created": "",
 *        "state": "",
 *        "charge_type": "",
 *        "ip": "",
 *        "platform_id": "",
 *        "real_name": "",
 *        "bank_type_id": "",
 *        "bank_charge_time": "",
 *        "credential_id": "",
 *        "operator": "",
 *        "is_tppay": "",
 *        "charge_card_id": "",
 *        "remark": "",
 *        "updated_last": ""
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *ChargeRecordsController) Onlines(ctx *Context) {
	records, err := chargeRecords.GetOnlines(ctx)
	if err != nil {
		responseFailure(ctx, err.Error(), "获取数据失败")
		return
	}
	responseSuccess(ctx, "获取数据成功", records)
}

/**
 * @api {get} admin/api/auth/v1/charge_records/view_online 				会员线上入款详情
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>充值记录详情</strong><br />
 * 业务描述: 充值记录详情</br>
 * @apiVersion 1.0.0
 * @apiName     viewChargeRecordsOnline
 * @apiGroup    finance
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
 * @apiSuccess (data字段说明) {int}        user_id                 用户ID
 * @apiSuccess (data字段说明) {float}      amount                  充值金额
 * @apiSuccess (data字段说明) {string}     order_id                充值订单
 * @apiSuccess (data字段说明) {int}        charge_type_id          充值方式id
 * @apiSuccess (data字段说明) {string}     card_number             卡号
 * @apiSuccess (data字段说明) {string}     bank_address            开户银行地址或支付二维码
 * @apiSuccess (data字段说明) {int}        state                   公司入款：0 待审核，1 成功，2 失败。线上支付：0待处理，1成功，2失败，3进行中,4退款，5取消，6强制入款
 * @apiSuccess (data字段说明) {string}     screenshot              屏幕截图
 * @apiSuccess (data字段说明) {string}     receipt_screenshot      收据截图
 * @apiSuccess (data字段说明) {string}     charge_type             充值类型
 * @apiSuccess (data字段说明) {string}     ip                      充值IP
 * @apiSuccess (data字段说明) {int}        platform_id             充值platform
 * @apiSuccess (data字段说明) {string}     real_name               真实姓名
 * @apiSuccess (data字段说明) {int}        bank_type_id            银行转账类型
 * @apiSuccess (data字段说明) {int}        bank_charge_time        银行转账时间
 * @apiSuccess (data字段说明) {int}        credential_id           第三方支付记录ID
 * @apiSuccess (data字段说明) {string}     operator                操作者
 * @apiSuccess (data字段说明) {int}        is_tppay                是否第三方支付 0为否；1为是
 * @apiSuccess (data字段说明) {int}        charge_card_id          关联ChargeBankCards.id
 * @apiSuccess (data字段说明) {string}     remark                  备注
 * @apiSuccess (data字段说明) {string}     updated_last            最后更新时间
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
 *        "charge_type_id": "",
 *        "card_number": "",
 *        "bank_address": "",
 *        "state": "",
 *        "screenshot": "",
 *        "receipt_screenshot": "",
 *        "charge_type": "",
 *        "ip": "",
 *        "platform_id": "",
 *        "real_name": "",
 *        "bank_type_id": "",
 *        "bank_charge_time": "",
 *        "credential_id": "",
 *        "operator": "",
 *        "is_tppay": "",
 *        "charge_card_id": "",
 *        "remark": "",
 *        "updated_last": ""
 * 	  },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *ChargeRecordsController) OnlineView(ctx *Context) {
	view(ctx, &chargeRecords)
}

/**
 * @api {get} admin/api/auth/v1/charge_records/delete_online 会员线上入款删除
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>充值记录删除</strong><br />
 * 业务描述: 充值记录删除</br>
 * @apiVersion 1.0.0
 * @apiName     deleteChargeRecordsOnline
 * @apiGroup    finance
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
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "记录删除成功",
 *    "code": 200,
 *    "data": {},
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *ChargeRecordsController) OnlineDelete(ctx *Context) {
	remove(ctx, &chargeRecords)
}

/**
 * @api {get} admin/api/auth/v1/charge_records/allow	会员公司入款审核
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>会员公司入款审核-通过</strong><br />
 * 业务描述: 会员公司入款审核-通过</br>
 * @apiVersion 1.0.0
 * @apiName     viewChargeRecordsAllow
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
func (self *ChargeRecordsController) Allow(ctx *Context) {
	responseResult(ctx, chargeRecords.Allow(ctx), "充值审核处理成功")
}

/**
 * @api {get} admin/api/auth/v1/charge_records/deny	会员公司入款拒绝
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>会员公司入款审核-拒绝</strong><br />
 * 业务描述: 会员公司入款审核-拒绝</br>
 * @apiVersion 1.0.0
 * @apiName     viewChargeRecordsDeny
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
func (self *ChargeRecordsController) Deny(ctx *Context) {
	responseResult(ctx, chargeRecords.Deny(ctx), "已拒绝用户的充值申请")
}

/**
 * @api {get} admin/api/auth/v1/charge_records/forced_deposit	会员公司强制入款
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: xiaoye</span><br/><br/>
 * <strong>会员公司强制入款</strong><br />
 * 业务描述: 会员公司强制入款</br>
 * @apiVersion 1.0.0
 * @apiName     viewChargeRecordsForcedDeposit
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
func (self *ChargeRecordsController) ForcedDeposit(ctx *Context) {
	responseResult(ctx, chargeRecords.ForcedDeposit(ctx), "强制入款成功")
}
