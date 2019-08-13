package controllers

import (
	"qpgame/admin/models"
)

var userLoginLogs = models.UserLoginLogs{} //模型

type UserLoginLogsController struct{}

/**
 * @api {get} admin/api/auth/v1/user_login_logs 登录日志列表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>登录日志列表</strong><br />
 * 业务描述: 登录日志列表</br>
 * @apiVersion 1.0.0
 * @apiName     indexUserLoginLogs
 * @apiGroup    user
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int}     page            页数
 * @apiParam (客户端请求参数) {int}    page_size       每页记录数
 * @apiParam (客户端请求参数) {int}    user_id 用户编号
 * @apiParam (客户端请求参数) {string}    ip IP
 * @apiParam (客户端请求参数) {int}    login_from 	来源/0:其也,1:IOS,2:安卓
 * @apiParam (客户端请求参数) {string}    time_start 登录时间/开始
 * @apiParam (客户端请求参数) {string}    time_end 	登录时间/结束
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
 * @apiSuccess (data-rows每个子对象字段说明) {int}        user_id                 用户编号
 * @apiSuccess (data-rows每个子对象字段说明) {string}        user_name                 用户名称
 * @apiSuccess (data-rows每个子对象字段说明) {int}        login_time              登录时间
 * @apiSuccess (data-rows每个子对象字段说明) {string}     ip                      ip
 * @apiSuccess (data-rows每个子对象字段说明) {string}     addr                    地址
 * @apiSuccess (data-rows每个子对象字段说明) {int}        logout_time             退出时间
 * @apiSuccess (data-rows每个子对象字段说明) {string}     login_from              登陆来源(ios、android
 * @apiSuccess (data-rows每个子对象字段说明) {string}     ip_info ip信息: 国家-省-市
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "数据获取成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "user_id": "",
 *        "user_name": "",
 *        "login_time": "",
 *        "ip": "",
 *        "addr": "",
 *        "logout_time": "",
 *        "login_from": "",
 *        "ip_info": ""
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *UserLoginLogsController) Index(ctx *Context) {
	index(ctx, &userLoginLogs)
}

/**
 * @api {get} admin/api/auth/v1/user_login_logs/view 				登录日志详情
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>登录日志详情</strong><br />
 * 业务描述: 登录日志详情</br>
 * @apiVersion 1.0.0
 * @apiName     viewUserLoginLogs
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
 * @apiSuccess (data字段说明) {int}        login_time              登录时间
 * @apiSuccess (data字段说明) {string}     ip                      ip
 * @apiSuccess (data字段说明) {string}     addr                    地址
 * @apiSuccess (data字段说明) {int}        logout_time             退出时间
 * @apiSuccess (data字段说明) {string}     login_from              登陆来源(ios、android
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "保存数据成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "user_id": "",
 *        "login_time": "",
 *        "ip": "",
 *        "addr": "",
 *        "logout_time": "",
 *        "login_from": ""
 * 	  },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *UserLoginLogsController) View(ctx *Context) {
	view(ctx, &userLoginLogs)
}

/**
 * @api {get} admin/api/auth/v1/user_login_logs/delete 登录日志删除
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>登录日志删除</strong><br />
 * 业务描述: 登录日志删除</br>
 * @apiVersion 1.0.0
 * @apiName     deleteUserLoginLogs
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
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "记录删除成功",
 *    "code": 200,
 *    "data": {},
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *UserLoginLogsController) Delete(ctx *Context) {
	remove(ctx, &userLoginLogs)
}
