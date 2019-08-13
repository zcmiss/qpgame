package controllers

import (
	"qpgame/admin/models"
)

var redpacketReceives = models.RedpacketReceives{} //模型

type RedpacketReceivesController struct{}

/**
 * @api {get} admin/api/auth/v1/redpacket_receives 红包发放列表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>红包发放列表</strong><br />
 * 业务描述: 红包发放列表</br>
 * @apiVersion 1.0.0
 * @apiName     indexRedpacketReceives
 * @apiGroup    finance
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int}    page            页数
 * @apiParam (客户端请求参数) {int}   page_size       每页记录数
 * @apiParam (客户端请求参数) {int}   user_id 	用户编号
 * @apiParam (客户端请求参数) {int}   redpacket_id 红包编号
 * @apiParam (客户端请求参数) {string}   created_start 抢包时间/开始
 * @apiParam (客户端请求参数) {string}   created_end 抢包时间/结束
 * @apiParam (客户端请求参数) {int}   red_type 红包类型
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
 * @apiSuccess (data-rows每个子对象字段说明) {int}        redpacket_id            红包id
 * @apiSuccess (data-rows每个子对象字段说明) {int}        is_get            是否已领取，0否，1是
 * @apiSuccess (data-rows每个子对象字段说明) {int}        user_id                 抢包用户编号
 * @apiSuccess (data-rows每个子对象字段说明) {string}        user_name                 抢包用户名称
 * @apiSuccess (data-rows每个子对象字段说明) {float}      money                   抢到的金额
 * @apiSuccess (data-rows每个子对象字段说明) {int}        created                 抢红包时间
 * @apiSuccess (data-rows每个子对象字段说明) {int}        red_type                红包类型
 * @apiSuccess (data-rows每个子对象字段说明) {string}        red_type_name                红包类型名称
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "数据获取成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "redpacket_id": "",
 *        "is_get": "0",
 *        "user_id": "",
 *        "user_name": "",
 *        "money": "",
 *        "created": "",
 *        "red_type": "",
 *        "red_type_name": ""
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *RedpacketReceivesController) Index(ctx *Context) {
	index(ctx, &redpacketReceives)
}

/**
 * @api {get} admin/api/v1/redpacket_receives/view 				红包发放详情
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>红包发放详情</strong><br />
 * 业务描述: 红包发放详情</br>
 * @apiVersion 1.0.0
 * @apiName     viewRedpacketReceives
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
 * @apiSuccess (data字段说明) {int}        redpacket_id            红包id
 * @apiSuccess (data字段说明) {int}        is_get                  是否已领取，0否，1是
 * @apiSuccess (data字段说明) {int}        user_id                 抢到红包的用户
 * @apiSuccess (data字段说明) {float}      money                   抢到的金额
 * @apiSuccess (data字段说明) {int}        created                 抢红包时间
 * @apiSuccess (data字段说明) {int}        red_type                红包类型
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "保存数据成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "redpacket_id": "",
 *        "user_id": "",
 *        "money": "",
 *        "created": "",
 *        "red_type": ""
 * 	  },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *RedpacketReceivesController) View(ctx *Context) {
	view(ctx, &redpacketReceives)
}

/**
 * @api {get} admin/api/v1/redpacket_receives/delete 红包发放删除
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>红包发放删除</strong><br />
 * 业务描述: 红包发放删除</br>
 * @apiVersion 1.0.0
 * @apiName     deleteRedpacketReceives
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
func (self *RedpacketReceivesController) Delete(ctx *Context) {
	remove(ctx, &redpacketReceives)
}
