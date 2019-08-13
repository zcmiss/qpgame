package controllers

import (
	"qpgame/admin/models"
)

var adminLogs = models.AdminLogs{} //模型

type AdminLogsController struct{}

/**
 * @api {get} admin/api/auth/v1/admin_logs 操作日志列表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>操作日志列表</strong><br />
 * 业务描述: 操作日志列表</br>
 * @apiVersion 1.0.0
 * @apiName     indexAdminLogs
 * @apiGroup    admin
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int} 	page			页数
 * @apiParam (客户端请求参数) {int} 	page_size		每页记录数
 * @apiParam (客户端请求参数) {int} 	admin_id 用户编号
 * @apiParam (客户端请求参数) {string} 	admin_name 用户名称
 * @apiParam (客户端请求参数) {string} 	time_start 	日志时间/开始
 * @apiParam (客户端请求参数) {string} 	time_end		日志时间/结束

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
 * @apiSuccess (data-rows每个子对象字段说明) {int}		id 				记录编号
 * @apiSuccess (data-rows每个子对象字段说明) {int}     	admin_id                AdminId
 * @apiSuccess (data-rows每个子对象字段说明) {string}     admin_name                后台用户名称
 * @apiSuccess (data-rows每个子对象字段说明) {string}     type                    操作类型
 * @apiSuccess (data-rows每个子对象字段说明) {string}     node                    操作节点
 * @apiSuccess (data-rows每个子对象字段说明) {string}     content                 操作内容
 * @apiSuccess (data-rows每个子对象字段说明) {int}        created                 操作时间
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "数据获取成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "admin_id": "",
 *        "type": "",
 *        "node": "",
 *        "content": "",
 *        "created": ""
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *AdminLogsController) Index(ctx *Context) {
	index(ctx, &adminLogs)
}

/**
 * @api {get} admin/api/auth/v1/admin_logs/view 				操作日志详情
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>操作日志详情</strong><br />
 * 业务描述: 操作日志详情</br>
 * @apiVersion 1.0.0
 * @apiName     viewAdminLogs
 * @apiGroup    admin
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
 * @apiSuccess (data字段说明) {string}     admin_id                AdminId
 * @apiSuccess (data字段说明) {string}     type                    操作类型
 * @apiSuccess (data字段说明) {string}     node                    操作节点ID
 * @apiSuccess (data字段说明) {string}     content                 操作内容
 * @apiSuccess (data字段说明) {int}        created                 操作时间
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "保存数据成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "admin_id": "",
 *        "type": "",
 *        "node": "",
 *        "content": "",
 *        "created": ""
 * 	  },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *AdminLogsController) View(ctx *Context) {
	view(ctx, &adminLogs)
}

/**
 * @api {get} admin/api/auth/v1/admin_logs/delete 操作日志删除
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>操作日志删除</strong><br />
 * 业务描述: 操作日志删除</br>
 * @apiVersion 1.0.0
 * @apiName     deleteAdminLogs
 * @apiGroup    admin
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
func (self *AdminLogsController) Delete(ctx *Context) {
	remove(ctx, &adminLogs)
}
