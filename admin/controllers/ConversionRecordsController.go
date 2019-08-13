package controllers

import (
	"qpgame/admin/models"
)

var conversionRecords = models.ConversionRecords{} //模型

type ConversionRecordsController struct{}

/**
 * @api {get} admin/api/auth/v1/conversion_records 转换记录列表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>转换记录列表</strong><br />
 * 业务描述: 转换记录列表</br>
 * @apiVersion 1.0.0
 * @apiName     indexConversionRecords
 * @apiGroup    finance
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int} 	page			页数
 * @apiParam (客户端请求参数) {int} 	page_size		每页记录数
 * @apiParam (客户端请求参数) {string} 	user_id		用户编号
 * @apiParam (客户端请求参数) {int} 	platform_id 第三方平台/game_platforms接口
 * @apiParam (客户端请求参数) {int} 	type 转换类型/1:向平台转入/2:从平台转出
 * @apiParam (客户端请求参数) {string} 	app_order_id 订单编号
 * @apiParam (客户端请求参数) {string} 	order_id 第三方订单编号
 * @apiParam (客户端请求参数) {int} 	status 状态/0:处理中/1:成功/2:失败
 * @apiParam (客户端请求参数) {string} 	created_start 转换时间/开始
 * @apiParam (客户端请求参数) {string} 	created_end 转换时间/结束
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
 * @apiSuccess (data-rows每个子对象字段说明) {int}        user_id                 用户ID
 * @apiSuccess (data-rows每个子对象字段说明) {int}        platform_id             第三方平台ID（关联tPPlatform表）
 * @apiSuccess (data-rows每个子对象字段说明) {int}        type                    转换类型 1向平台转入，2从平台转出
 * @apiSuccess (data-rows每个子对象字段说明) {string}     app_order_id            本平台订单号，确保唯一，关联balanceLog的orderID
 * @apiSuccess (data-rows每个子对象字段说明) {string}     order_id                第三方订单号
 * @apiSuccess (data-rows每个子对象字段说明) {float}      amount                  上分金额
 * @apiSuccess (data-rows每个子对象字段说明) {int}        status                  订单状态，0处理中，1成功，2失败
 * @apiSuccess (data-rows每个子对象字段说明) {int}        created                 创建时间
 * @apiSuccess (data-rows每个子对象字段说明) {float}      tp_remain               第三方平台转账后金额
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "数据获取成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "user_id": "",
 *        "platform_id": "",
 *        "type": "",
 *        "app_order_id": "",
 *        "order_id": "",
 *        "amount": "",
 *        "status": "",
 *        "created": "",
 *        "tp_remain": ""
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *ConversionRecordsController) Index(ctx *Context) {
	index(ctx, &conversionRecords)
}

/**
 * @api {get} admin/api/auth/v1/conversion_records/view 				转换记录详情
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>转换记录详情</strong><br />
 * 业务描述: 转换记录详情</br>
 * @apiVersion 1.0.0
 * @apiName     viewConversionRecords
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
 * @apiSuccess (data字段说明) {int}        platform_id             第三方平台ID（关联tPPlatform表）
 * @apiSuccess (data字段说明) {int}        type                    转换类型 1向平台转入，2从平台转出
 * @apiSuccess (data字段说明) {string}     app_order_id            本平台订单号，确保唯一，关联balanceLog的orderID
 * @apiSuccess (data字段说明) {string}     order_id                第三方订单号
 * @apiSuccess (data字段说明) {float}      amount                  上分金额
 * @apiSuccess (data字段说明) {int}        status                  订单状态，0处理中，1成功，2失败
 * @apiSuccess (data字段说明) {int}        created                 创建时间
 * @apiSuccess (data字段说明) {float}      tp_remain               第三方平台转账后金额
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "保存数据成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "user_id": "",
 *        "platform_id": "",
 *        "type": "",
 *        "app_order_id": "",
 *        "order_id": "",
 *        "amount": "",
 *        "status": "",
 *        "created": "",
 *        "tp_remain": ""
 * 	  },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *ConversionRecordsController) View(ctx *Context) {
	view(ctx, &conversionRecords)
}

/**
 * @api {get} admin/api/auth/v1/conversion_records/delete 转换记录删除
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>转换记录删除</strong><br />
 * 业务描述: 转换记录删除</br>
 * @apiVersion 1.0.0
 * @apiName     deleteConversionRecords
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
 * @apiSuccess (返回结果)  {json}  		data            返回数据
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
func (self *ConversionRecordsController) Delete(ctx *Context) {
	remove(ctx, &conversionRecords)
}
