package controllers

import (
	"qpgame/admin/models"
)

var washCodeRecords = models.WashCodeRecords{} //模型

type WashCodeRecordsController struct{}

/**
 * @api {get} admin/api/auth/v1/wash_code_records 洗码记录列表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>洗码记录列表</strong><br />
 * 业务描述: 洗码记录列表</br>
 * @apiVersion 1.0.0
 * @apiName     indexWashCodeRecords
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
 * @apiParam (客户端请求参数) {string} 	wash_start 处理时间/开始
 * @apiParam (客户端请求参数) {string} 	wash_end 处理时间/开始
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
 * @apiSuccess (data-rows每个子对象字段说明) {float}      total_betamount         洗码量
 * @apiSuccess (data-rows每个子对象字段说明) {float}      amount                  洗码金额
 * @apiSuccess (data-rows每个子对象字段说明) {int}        washtime                洗码时间
 * @apiSuccess (data-rows每个子对象字段说明) {int}        user_id                 用户编号
 * @apiSuccess (data-rows每个子对象字段说明) {string}        user_name                 用户名称
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "数据获取成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "total_betamount": "",
 *        "amount": "",
 *        "washtime": "",
 *        "user_id": "",
 *        "user_name": ""
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *WashCodeRecordsController) Index(ctx *Context) {
	index(ctx, &washCodeRecords)
}

/**
 * @api {get} admin/api/auth/v1/wash_code_records/view 				洗码记录详情
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>洗码记录详情</strong><br />
 * 业务描述: 洗码记录详情</br>
 * @apiVersion 1.0.0
 * @apiName     viewWashCodeRecords
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
 * @apiSuccess (data字段说明) {float}      total_betamount         洗码量
 * @apiSuccess (data字段说明) {float}      amount                  洗码金额
 * @apiSuccess (data字段说明) {int}        washtime                洗码时间
 * @apiSuccess (data字段说明) {int}        user_id                 用户编号
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "保存数据成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "total_betamount": "",
 *        "amount": "",
 *        "washtime": "",
 *        "user_id": ""
 * 	  },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *WashCodeRecordsController) View(ctx *Context) {
	view(ctx, &washCodeRecords)
}

/**
 * @api {get} admin/api/auth/v1/wash_code_records/delete 洗码记录删除
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>洗码记录删除</strong><br />
 * 业务描述: 洗码记录删除</br>
 * @apiVersion 1.0.0
 * @apiName     deleteWashCodeRecords
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
func (self *WashCodeRecordsController) Delete(ctx *Context) {
	remove(ctx, &washCodeRecords)
}
