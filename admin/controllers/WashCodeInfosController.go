package controllers

import (
	"qpgame/admin/models"
)

var washCodeInfos = models.WashCodeInfos{} //模型

type WashCodeInfosController struct{}

/**
 * @api {get} admin/api/auth/v1/wash_code_infos 洗码信息列表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>洗码信息列表</strong><br />
 * 业务描述: 洗码信息列表</br>
 * @apiVersion 1.0.0
 * @apiName     indexWashCodeInfos
 * @apiGroup    finance
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int} 	page			页数
 * @apiParam (客户端请求参数) {int} 	page_size		每页记录数
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
 * @apiSuccess (data-rows每个子对象字段说明) {int}        record_id               洗码记录编号
 * @apiSuccess (data-rows每个子对象字段说明) {string}     type_name                游戏类型名称
 * @apiSuccess (data-rows每个子对象字段说明) {int}        type_id                  游戏类型编号
 * @apiSuccess (data-rows每个子对象字段说明) {string}     game_name                游戏名称
 * @apiSuccess (data-rows每个子对象字段说明) {float}      bet_amount               游戏洗码量
 * @apiSuccess (data-rows每个子对象字段说明) {float}      amount                  洗码金额
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "数据获取成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "record_id": "",
 *        "type_name": "",
 *        "type_id": "",
 *        "game_name": "",
 *        "bet_amount": "",
 *        "amount": ""
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *WashCodeInfosController) Index(ctx *Context) {
	index(ctx, &washCodeInfos)
}

/**
 * @api {get} admin/api/auth/v1/wash_code_infos/view 				洗码信息详情
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>洗码信息详情</strong><br />
 * 业务描述: 洗码信息详情</br>
 * @apiVersion 1.0.0
 * @apiName     viewWashCodeInfos
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
 * @apiSuccess (data字段说明) {int}        record_id               洗码记录编号
 * @apiSuccess (data字段说明) {string}     type_name                游戏类型名称
 * @apiSuccess (data字段说明) {int}        type_id                  游戏类型编号
 * @apiSuccess (data字段说明) {string}     game_name                游戏名称
 * @apiSuccess (data字段说明) {float}      bet_amount               游戏洗码量
 * @apiSuccess (data字段说明) {float}      amount                  洗码金额
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "保存数据成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "record_id": "",
 *        "type_name": "",
 *        "type_id": "",
 *        "game_name": "",
 *        "bet_amount": "",
 *        "amount": ""
 * 	  },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *WashCodeInfosController) View(ctx *Context) {
	view(ctx, &washCodeInfos)
}

/**
 * @api {get} admin/api/auth/v1/wash_code_infos/delete 洗码信息删除
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>洗码信息删除</strong><br />
 * 业务描述: 洗码信息删除</br>
 * @apiVersion 1.0.0
 * @apiName     deleteWashCodeInfos
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
func (self *WashCodeInfosController) Delete(ctx *Context) {
	remove(ctx, &washCodeInfos)
}
