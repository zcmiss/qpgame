package controllers

import (
	"qpgame/admin/models"
)

var adminLoginLogs = models.AdminLoginLogs{} //模型

type AdminLoginLogsController struct{}

/**
 * @api {get} admin/api/auth/v1/admin_login_logs 后台登录日志
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>后台登录日志列表</strong><br />
 * 业务描述: 后台登录日志列表</br>
 * @apiVersion 1.0.0
 * @apiName     indexAdminLoginLogs
 * @apiGroup    admin
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int} 	page			页数
 * @apiParam (客户端请求参数) {int} 	page_size		每页记录数
 * @apiParam (客户端请求参数) {int} 	admin_id 	用户编号
 * @apiParam (客户端请求参数) {int} 	admin_name	用户名称
 * @apiParam (客户端请求参数) {string} 	login_time_start 登录时间/开始
 * @apiParam (客户端请求参数) {string} 	login_time_end 	登录时间/结束
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
 * @apiSuccess (data-rows每个子对象字段说明) {int}		  id 					编号
 * @apiSuccess (data-rows每个子对象字段说明) {string}     admin_id                管理员编号
 * @apiSuccess (data-rows每个子对象字段说明) {string}     admin_name              管理员编号
 * @apiSuccess (data-rows每个子对象字段说明) {string}     login_time              登录时间
 * @apiSuccess (data-rows每个子对象字段说明) {string}     ip                      登录IP
 * @apiSuccess (data-rows每个子对象字段说明) {string}     ip_info                 IP地址信息
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "数据获取成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "admin_id": "",
 *        "admin_name": "",
 *        "login_time": "",
 *        "ip": ""
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *AdminLoginLogsController) Index(ctx *Context) {
	index(ctx, &adminLoginLogs)
}

/**
 * @api {get} admin/api/auth/v1/admin_login_logs/view 				后台登录日志详情
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>后台登录日志详情</strong><br />
 * 业务描述: 后台登录日志详情</br>
 * @apiVersion 1.0.0
 * @apiName     viewAdminLoginLogs
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
 * @apiSuccess (data字段说明) {int}		  id 					  编号
 * @apiSuccess (data字段说明) {string}     admin_id                管理员编号
 * @apiSuccess (data字段说明) {string}     admin_name                管理员名称
 * @apiSuccess (data字段说明) {string}     login_time              登录时间
 * @apiSuccess (data字段说明) {string}     ip                      登录IP
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "保存数据成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "admin_id": "",
 *        "admin_name": "",
 *        "login_time": "",
 *        "ip": ""
 * 	  },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *AdminLoginLogsController) View(ctx *Context) {
	view(ctx, &adminLoginLogs)
}

/**
 * @api {get} admin/api/auth/v1/admin_login_logs/delete 后台登录日志删除
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>后台登录日志删除</strong><br />
 * 业务描述: 后台登录日志删除</br>
 * @apiVersion 1.0.0
 * @apiName     deleteAdminLoginLogs
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
func (self *AdminLoginLogsController) Delete(ctx *Context) {
	remove(ctx, &adminLoginLogs)
}
