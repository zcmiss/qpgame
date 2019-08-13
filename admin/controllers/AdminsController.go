package controllers

import (
	"qpgame/admin/models"
	"qpgame/admin/validations"
)

var admins = models.Admins{}                          //模型
var adminsValidation = validations.AdminsValidation{} //校验器

type AdminsController struct{}

/**
 * @api {get} admin/api/auth/v1/admins 后台用户列表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>后台用户列表</strong><br />
 * 业务描述: 后台用户列表</br>
 * @apiVersion 1.0.0
 * @apiName     indexAdmins
 * @apiGroup    admin
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int} 	page			页数
 * @apiParam (客户端请求参数) {int} 	page_size		每页记录数
 * @apiParam (客户端请求参数) {string} 	name 用户名称
 * @apiParam (客户端请求参数) {int} 	role_id 角色/下拉/从角色列表获取
 * @apiParam (客户端请求参数) {int} 	status 状态
 * @apiParam (客户端请求参数) {int} 	is_otp 是否OTP登录
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
 * @apiSuccess (data-rows每个子对象字段说明) {int}		id 						记录编号
 * @apiSuccess (data-rows每个子对象字段说明) {string}     name                    管理员名称
 * @apiSuccess (data-rows每个子对象字段说明) {string}     email                   电子邮件
 * @apiSuccess (data-rows每个子对象字段说明) {int}        role_id                 角色编号
 * @apiSuccess (data-rows每个子对象字段说明) {int}        role_name               角色名称
 * @apiSuccess (data-rows每个子对象字段说明) {string}     status                  状态,1可用,0禁用
 * @apiSuccess (data-rows每个子对象字段说明) {string}     charge_alert            后台充值提醒，1开启，0关闭
 * @apiSuccess (data-rows每个子对象字段说明) {string}     withdraw_alert          后台出款提醒，1开启，0关闭
 * @apiSuccess (data-rows每个子对象字段说明) {string}     permission              涉及钱的权限, 0:无权限，1:主管权限
 * @apiSuccess (data-rows每个子对象字段说明) {string}     force_out               是否强制退出, 0:无强制退出,1: 强制退出
 * @apiSuccess (data-rows每个子对象字段说明) {int}        manual_max              最大人工入款金额/數字
 * @apiSuccess (data-rows每个子对象字段说明) {string}     is_otp                  是否OTP验证登录, 0:否，1:是
 * @apiSuccess (data-rows每个子对象字段说明) {string}     is_otp_first            是否第一次OTP验证登录
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "数据获取成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "name": "",
 *        "email": "",
 *        "role_id": "",
 *        "created": "",
 *        "updated": "",
 *        "status": "",
 *        "charge_alert": "",
 *        "withdraw_alert": "",
 *        "permission": "",
 *        "force_out": "",
 *        "manual_max": "",
 *        "is_otp": "",
 *        "is_otp_first": ""
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *AdminsController) Index(ctx *Context) {
	index(ctx, &admins)
}

/**
 * @api {post} admin/api/auth/v1/admins/add	后台用户添加
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>添加/修改后台用户</strong><br />
 * 业务描述: 添加/修改后台用户 <br />
 * <strong><span style="color: red">注意: </span></strong><br />
 * <span style="color:red">修改操作API不再单独列出, 请参考以下</span><br />
 * <span style="color:red">添加: /admin/api/auth/v1/admins/add </span> &nbsp;&nbsp; <br />
 * <span style="color:red">修改: /admin/api/auth/v1/admins/update </span>
 * @apiVersion 1.0.0
 * @apiName     saveAdmins
 * @apiGroup    admin
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }

 * @apiParam (客户端请求参数) {int} 	id    					记录编号,仅修改操作时(即接口: /admin/api/auth/v1/admins/update)需要<br />
 *															添加操作(即接口: /admin/api/auth/v1/admins/add)不需要此参数<br />
 * 															* 如果提供此编号, 则视为修改记录
 * @apiParam (客户端请求参数) {string}   	name                    管理员名称
 * @apiParam (客户端请求参数) {string}   	email                   电子邮件
 * @apiParam (客户端请求参数) {string}   	password                密码
 * @apiParam (客户端请求参数) {int}      	role_id                 角色,需要先從管理員角色列表拿到所有角色
 * @apiParam (客户端请求参数) {int}      	status                  状态
 * @apiParam (客户端请求参数) {int}      	charge_alert            后台充值提醒，1开启，0关闭
 * @apiParam (客户端请求参数) {int}      	withdraw_alert          后台出款提醒，1开启，0关闭
 * @apiParam (客户端请求参数) {string}   	login_ip                允许登录IP
 * @apiParam (客户端请求参数) {int}      	permission              涉及钱的权限（0、无权限，1、主管权限）
 * @apiParam (客户端请求参数) {int}      	force_out               是否强制退出,0 无强制退出,1 强制退出
 * @apiParam (客户端请求参数) {int}      	manual_max              最大人工入款金额
 * @apiParam (客户端请求参数) {int}      	is_otp                  是否OTP验证登录（0为否，1为是）
 * @apiParam (客户端请求参数) {int}      	is_otp_first            是否第一次OTP验证登录,0不是,1是
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
 * @apiSuccess (返回结果)  {string}  	internalMsg     内部错误信息
 * @apiSuccess (返回结果)  {json}		data            返回数据
 * @apiSuccess (返回结果)  {float}   	timeConsumed    后台耗时
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "保存数据成功",
 *    "code": 200,
 *    "data": { },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *AdminsController) Save(ctx *Context) {
	save(ctx, &admins, &adminsValidation)
}

/**
 * @api {get} admin/api/auth/v1/admins/view 				后台用户详情
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>后台用户详情</strong><br />
 * 业务描述: 后台用户详情</br>
 * @apiVersion 1.0.0
 * @apiName     viewAdmins
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
 * @apiSuccess (data字段说明) {string}     name                    管理员名称
 * @apiSuccess (data字段说明) {string}     email                   电子邮件
 * @apiSuccess (data字段说明) {int}        role_id                 角色
 * @apiSuccess (data字段说明) {int}        created                 创建时间
 * @apiSuccess (data字段说明) {int}        updated                 更新时间
 * @apiSuccess (data字段说明) {int}        status                  状态
 * @apiSuccess (data字段说明) {int}        charge_alert            后台充值提醒，1开启，0关闭
 * @apiSuccess (data字段说明) {int}        withdraw_alert          后台出款提醒，1开启，0关闭
 * @apiSuccess (data字段说明) {string}     login_ip                允许登录IP
 * @apiSuccess (data字段说明) {int}        permission              涉及钱的权限（0、无权限，1、主管权限）
 * @apiSuccess (data字段说明) {int}        force_out               是否强制退出,0 无强制退出,1 强制退出
 * @apiSuccess (data字段说明) {int}        manual_max              最大人工入款金额
 * @apiSuccess (data字段说明) {int}        is_otp                  是否OTP验证登录（0为否，1为是）
 * @apiSuccess (data字段说明) {int}        is_otp_first            是否第一次OTP验证登录
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "保存数据成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "name": "",
 *        "email": "",
 *        "password": "",
 *        "role_id": "",
 *        "created": "",
 *        "updated": "",
 *        "status": "",
 *        "charge_alert": "",
 *        "withdraw_alert": "",
 *        "login_ip": "",
 *        "permission": "",
 *        "force_out": "",
 *        "manual_max": "",
 *        "is_otp": "",
 *        "is_otp_first": ""
 * 	  },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *AdminsController) View(ctx *Context) {
	view(ctx, &admins)
}

/**
 * @api {get} admin/api/auth/v1/admins/delete 后台用户删除
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>后台用户删除</strong><br />
 * 业务描述: 后台用户删除</br>
 * @apiVersion 1.0.0
 * @apiName     deleteAdmins
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
func (self *AdminsController) Delete(ctx *Context) {
	remove(ctx, &admins)
}

/**
 * @api {post} admin/api/auth/v1/admins/update_password	后台用户密码修改
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>修改后台用户密码</strong><br />
 * 业务描述: 修改后台用户密码<br />
 * @apiVersion 1.0.0
 * @apiName     updateAdminsPass
 * @apiGroup    admin
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }

 * @apiParam (客户端请求参数) {int} 	id    					用户编号
 * @apiParam (客户端请求参数) {string}   	old_password 旧密码
 * @apiParam (客户端请求参数) {string}   	password 密码
 * @apiParam (客户端请求参数) {string}   	re_password 确认密码
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
 * @apiSuccess (返回结果)  {string}  	internalMsg     内部错误信息
 * @apiSuccess (返回结果)  {json}		data            返回数据
 * @apiSuccess (返回结果)  {float}   	timeConsumed    后台耗时
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "密码修改成功",
 *    "code": 200,
 *    "data": { },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *AdminsController) UpdatePassword(ctx *Context) {
	responseResult(ctx, admins.UpdatePassword(ctx), "密码修改成功, 请使用新密码重新登录")
}
