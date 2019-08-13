package controllers

import (
	"qpgame/admin/models"
	"qpgame/admin/validations"
)

var users = models.Users{}                          //模型
var usersValidation = validations.UsersValidation{} //校验器

type UsersController struct{}

/**
 * @api {get} admin/api/auth/v1/users 用户列表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>注册用户列表</strong><br />
 * 业务描述: 注册用户列表</br>
 * @apiVersion 1.0.0
 * @apiName     indexUsers
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
 * @apiParam (客户端请求参数) {int}    parent_id 上级用户编号
 * @apiParam (客户端请求参数) {string}    name 用户名称
 * @apiParam (客户端请求参数) {string}    phone 手机号码
 * @apiParam (客户端请求参数) {string}    qq QQ号码
 * @apiParam (客户端请求参数) {string}    wechat 微信
 * @apiParam (客户端请求参数) {int}    vip_level VIP等级/从VIP等级接口获取/下拉列表
 * @apiParam (客户端请求参数) {int}    status 状态
 * @apiParam (客户端请求参数) {int}    mobile_type 手机类型/0:其他,1:IOS,2:安卓
 * @apiParam (客户端请求参数) {string}    created_start 	注册时间/开始
 * @apiParam (客户端请求参数) {string}    created_end 	注册时间/结束
 * @apiParam (客户端请求参数) {string}    last_login_start	最后登录时间/开始
 * @apiParam (客户端请求参数) {string}    last_login_end 	最后登录时间/结束
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
 * @apiSuccess (data-rows每个子对象字段说明) {string}     name                    用户姓名
 * @apiSuccess (data-rows每个子对象字段说明) {string}     phone                   手机号
 * @apiSuccess (data-rows每个子对象字段说明) {int}        mobile_type             手机类型, 1:安卓,2:IOS
 * @apiSuccess (data-rows每个子对象字段说明) {int}        sex                     性别:1男,2女
 * @apiSuccess (data-rows每个子对象字段说明) {int}        vip_level               vip等级(1-10)
 * @apiSuccess (data-rows每个子对象字段说明) {string}     qq                      用户QQ号
 * @apiSuccess (data-rows每个子对象字段说明) {string}     wechat                  用户微信号
 * @apiSuccess (data-rows每个子对象字段说明) {int}        status                  用户状态,1正常，0锁定
 * @apiSuccess (data-rows每个子对象字段说明) {int}        created                 创建时间
 * @apiSuccess (data-rows每个子对象字段说明) {int}        parent_id               上级代理编号
 * @apiSuccess (data-rows每个子对象字段说明) {int}        parent_name               上级代理名称
 * @apiSuccess (data-rows每个子对象字段说明) {int}        last_login_time         上次登录时间
 * @apiSuccess (data-rows每个子对象字段说明) {int}        last_platform_id        上次登录的游戏编号
 * @apiSuccess (data-rows每个子对象字段说明) {int}        last_platform_name		上次登录平台名称
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "数据获取成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "phone": "",
 *        "name": "",
 *        "email": "",
 *        "mobile_type": "",
 *        "sex": "",
 *        "path": "",
 *        "vip_level": "",
 *        "qq": "",
 *        "wechat": "",
 *        "status": "",
 *        "is_dummy": "",
 *        "parent_id": "",
 *        "created": "",
 *        "last_login_time": "",
 *        "last_platform_id": "",
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *UsersController) Index(ctx *Context) {
	index(ctx, &users)
}

/**
 * @api {post} admin/api/auth/v1/users/add	用户添加
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>添加/修改注册用户</strong><br />
 * 业务描述: 添加/修改注册用户 <br />
 * <strong><span style="color: red">注意: </span></strong><br />
 * <span style="color:red">修改操作API不再单独列出, 请参考以下</span><br />
 * <span style="color:red">添加: /admin/api/auth/v1/users/add </span> &nbsp;&nbsp; <br />
 * <span style="color:red">修改: /admin/api/auth/v1/users/update </span>
 * @apiVersion 1.0.0
 * @apiName     saveUsers
 * @apiGroup    user
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }

 * @apiParam (客户端请求参数) {int} 	id    					记录编号,仅修改操作时(即接口: /admin/api/auth/v1/users/update)需要<br />
 *															添加操作(即接口: /admin/api/auth/v1/users/add)不需要此参数<br />
 * 															* 如果提供此编号, 则视为修改记录
 * @apiParam (客户端请求参数) {string}   	phone                   手机号
 * @apiParam (客户端请求参数) {string}   	password                密码
 * @apiParam (客户端请求参数) {string}   	safe_password           保险箱密码
 * @apiParam (客户端请求参数) {string}   	name                    用户姓名
 * @apiParam (客户端请求参数) {string}   	email                   邮箱
 * @apiParam (客户端请求参数) {int}      	birthday                生日
 * @apiParam (客户端请求参数) {int}      	mobile_type             手机类型,1:安卓,2:ios
 * @apiParam (客户端请求参数) {int}      	sex                     性别:1男,2女
 * @apiParam (客户端请求参数) {string}   	path                    代理层级id路径例子:(,1,2,4,5,7,
 * @apiParam (客户端请求参数) {int}      	vip_level               vip等级(1-10)
 * @apiParam (客户端请求参数) {string}   	qq                      用户QQ号
 * @apiParam (客户端请求参数) {string}   	wechat                  用户微信号
 * @apiParam (客户端请求参数) {int}      	status                  用户状态,1正常，0锁定
 * @apiParam (客户端请求参数) {int}      	is_dummy                0:正常用户，1：虚拟用户
 * @apiParam (客户端请求参数) {int}      	parent_id               上级代理用户Id
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
func (self *UsersController) Save(ctx *Context) {
	save(ctx, &users, &usersValidation)
}

/**
 * @api {get} admin/api/auth/v1/users/view 				用户详情
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>注册用户详情</strong><br />
 * 业务描述: 注册用户详情</br>
 * @apiVersion 1.0.0
 * @apiName     viewUsers
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
 * @apiSuccess (data字段说明) {string}     phone                   手机号
 * @apiSuccess (data字段说明) {string}     name                    用户姓名
 * @apiSuccess (data字段说明) {string}     email                   邮箱
 * @apiSuccess (data字段说明) {int}        created                 创建时间
 * @apiSuccess (data字段说明) {int}        birthday                生日
 * @apiSuccess (data字段说明) {int}        mobile_type             手机类型,1:安卓,2:IOS
 * @apiSuccess (data字段说明) {int}        sex                     性别:1男,2女
 * @apiSuccess (data字段说明) {string}     path                    代理层级id路径例子:(,1,2,4,5,7,
 * @apiSuccess (data字段说明) {int}        vip_level               vip等级(1-10
 * @apiSuccess (data字段说明) {string}     qq                      用户QQ号
 * @apiSuccess (data字段说明) {string}     wechat                  用户微信号
 * @apiSuccess (data字段说明) {int}        status                  用户状态,1正常，0锁定
 * @apiSuccess (data字段说明) {int}        is_dummy                0:正常用户，1：虚拟用户
 * @apiSuccess (data字段说明) {string}     token                   用户登录token,要保持到程序内存中
 * @apiSuccess (data字段说明) {int}        token_created           token创建时间,根据这个来双层判断是否已过期
 * @apiSuccess (data字段说明) {int}        parent_id               上级代理用户Id
 * @apiSuccess (data字段说明) {int}        last_login_time         上次登录时间
 * @apiSuccess (data字段说明) {int}        last_platform_id        上次登录的游戏平台
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "保存数据成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "phone": "",
 *        "name": "",
 *        "email": "",
 *        "created": "",
 *        "birthday": "",
 *        "mobile_type": "",
 *        "sex": "",
 *        "path": "",
 *        "vip_level": "",
 *        "qq": "",
 *        "wechat": "",
 *        "status": "",
 *        "is_dummy": "",
 *        "token": "",
 *        "token_created": "",
 *        "parent_id": "",
 *        "last_login_time": "",
 *        "last_platform_id": "",
 * 	  },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *UsersController) View(ctx *Context) {
	view(ctx, &users)
}

/**
 * @api {get} admin/api/auth/v1/users/delete 用户删除
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>注册用户删除</strong><br />
 * 业务描述: 注册用户删除</br>
 * @apiVersion 1.0.0
 * @apiName     deleteUsers
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
func (self *UsersController) Delete(ctx *Context) {
	remove(ctx, &users)
}

/**
 * @api {post} admin/api/auth/v1/users/update_password	用户修改密码
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>修改用户密码</strong><br />
 * 业务描述: 修改用户密码<br />
 * @apiVersion 1.0.0
 * @apiName     updateUsersPass
 * @apiGroup    user
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
func (self *UsersController) UpdatePassword(ctx *Context) {
	err := users.UpdatePassword(ctx)
	if err != nil {
		responseFailure(ctx, err.Error(), "密码修改失败")
		return
	}

	responseSuccess(ctx, "密码修改成功", "")
}

/**
 * @api {post} admin/api/auth/v1/users/update_safe_password	用户修改安全密码
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>修改用户安全密码</strong><br />
 * 业务描述: 修改用户安全密码<br />
 * @apiVersion 1.0.0
 * @apiName     updateUsersSafePass
 * @apiGroup    user
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }

 * @apiParam (客户端请求参数) {int} 	id    	用户编号
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
func (self *UsersController) UpdateSafePassword(ctx *Context) {
	err := users.UpdateSafePassword(ctx)
	if err != nil {
		responseFailure(ctx, "密码修改失败", err.Error())
		return
	}
	responseSuccess(ctx, "密码修改成功", "")
}

/**
 * @api {get} admin/api/auth/v1/users/lock	用户功能-锁定
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>用户锁定</strong><br />
 * 业务描述: 用户锁定<br />
 * @apiVersion 1.0.0
 * @apiName     updateUsersLock
 * @apiGroup    user
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }

 * @apiParam (客户端请求参数) {int} 	id    	用户编号
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
func (self *UsersController) Lock(ctx *Context) {
	err := users.Lock(ctx)
	if err != nil {
		responseFailure(ctx, "", err.Error())
		return
	}

	responseSuccess(ctx, "锁定用户成功", "")
}

/**
 * @api {get} admin/api/auth/v1/users/unlock 用户功能-解锁
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>用户解锁</strong><br />
 * 业务描述: 用户解锁<br />
 * @apiVersion 1.0.0
 * @apiName     updateUsersUnlock
 * @apiGroup   user
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }

 * @apiParam (客户端请求参数) {int} 	id    	用户编号
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
func (self *UsersController) Unlock(ctx *Context) {
	err := users.Unlock(ctx)
	if err != nil {
		responseFailure(ctx, "", err.Error())
		return
	}

	responseSuccess(ctx, "解锁用户成功", "")
}

/**
 * @api {get} admin/api/auth/v1/users/lock_proxy	用户功能-代理锁定
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>代理锁定</strong><br />
 * 业务描述: 代理锁定<br />
 * @apiVersion 1.0.0
 * @apiName     updateUsersLockProxy
 * @apiGroup    user
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }

 * @apiParam (客户端请求参数) {int} 	id    	用户编号
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
func (self *UsersController) LockProxy(ctx *Context) {
	err := users.LockProxy(ctx)
	if err != nil {
		responseFailure(ctx, "", err.Error())
		return
	}

	responseSuccess(ctx, "锁定代理成功", "")
}

/**
 * @api {get} admin/api/auth/v1/users/unlock_proxy 用户功能-代理解锁
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>代理解锁</strong><br />
 * 业务描述: 代理解锁<br />
 * @apiVersion 1.0.0
 * @apiName     updateUsersUnlockProxy
 * @apiGroup   user
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }

 * @apiParam (客户端请求参数) {int} 	id    	用户编号
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
func (self *UsersController) UnlockProxy(ctx *Context) {
	err := users.UnlockProxy(ctx)
	if err != nil {
		responseFailure(ctx, "", err.Error())
		return
	}

	responseSuccess(ctx, "解锁代理成功", "")
}

/**
 * @api {get} admin/api/auth/v1/users/query		用户信息查询
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong> 用户信息查询</strong><br />
 * 业务描述: 用户信息查询，根据用户编号(id)来查询用户信息</br>
 * @apiVersion 1.0.0
 * @apiName     UsersQueryUser
 * @apiGroup    finance
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 *
 * @apiParam (客户端请求参数) {string} 	user_name    			用户名称
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
 * @apiSuccess (data字段说明) {int}  	id		用户编号
 * @apiSuccess (data字段说明) {string}  	user_name	用户名称
 * @apiSuccess (data字段说明) {string}  	name		用户姓名
 * @apiSuccess (data字段说明) {float}  	balance_wallet	钱包余额
 * @apiSuccess (data字段说明) {float}  	charged_amount	充值总额
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "用户存在",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "name": "",
 *        "user_name": "",
 *        "balance_wallet": "",
 *        "charged_amount": "",
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *UsersController) QueryUser(ctx *Context) {
	user, err := users.GetQueryUser(ctx)
	if err != nil {
		responseFailure(ctx, "", err.Error())
		return
	}
	responseSuccess(ctx, "", user)
}

/**
 * @api {get} admin/api/auth/v1/users/invite 邀请用户列表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: lyndon</span><br/><br/>
 * <strong>邀请用户列表</strong><br />
 * 业务描述: 邀请用户列表</br>
 * @apiVersion 1.0.0
 * @apiName     inviteUsers
 * @apiGroup    user
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int}    page            页数
 * @apiParam (客户端请求参数) {int}    page_size       每页记录数
 * @apiParam (客户端请求参数) {int}    user_id 用户编号
 * @apiParam (客户端请求参数) {int}    parent_id 上级用户编号
 * @apiParam (客户端请求参数) {string}    name 用户名称
 * @apiParam (客户端请求参数) {int}    status 状态
 * @apiParam (客户端请求参数) {string}    created_start 	注册时间/开始
 * @apiParam (客户端请求参数) {string}    created_end 	注册时间/结束
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
 * @apiSuccess (data字段说明) {float}    	invite_amount	邀请奖励总金额
 *
 * @apiSuccess (data-rows每个子对象字段说明) {int}		   id 					   用户编号
 * @apiSuccess (data-rows每个子对象字段说明) {string}     name                    用户姓名
 * @apiSuccess (data-rows每个子对象字段说明) {int}        parent_id               上级代理编号
 * @apiSuccess (data-rows每个子对象字段说明) {string}     parent_name             上级代理名称
 * @apiSuccess (data-rows每个子对象字段说明) {int}        invite_count            邀请人数
 * @apiSuccess (data-rows每个子对象字段说明) {float}      invite_amount           邀请奖励
 * @apiSuccess (data-rows每个子对象字段说明) {int}        status                  用户状态,1正常，0锁定
 * @apiSuccess (data-rows每个子对象字段说明) {int}        created                 注册时间
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "数据获取成功",
 *    "code": 200,
 *    "data": {
 *        "rows": [
 *            {
 *              "id": "",
 *              "name": "",
 *              "parent_id": "",
 *              "parent_name": "",
 *              "invite_count": "",
 *              "invite_amount": "",
 *              "status": "",
 *              "created": ""
 *            }
 *         ],
 *         "page": 1,
 *         "page_count": 1,
 *         "total_rows": 1,
 *         "page_size": 1,
 *         "invite_amount": 1.987
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *UsersController) Invite(ctx *Context) {
	records, err := users.GetInviteRecords(ctx)
	if err != nil {
		responseFailure(ctx, err.Error(), "获取数据失败")
		return
	}
	result := make(map[string]interface{})
	result["rows"] = records.Rows
	result["page"] = records.Page
	result["page_size"] = records.PageSize
	result["page_count"] = records.PageCount
	result["total_rows"] = records.TotalRows
	configs := models.Configs{}
	result["invite_amount"] = (&configs).GetInviteAmount(ctx)
	responseSuccess(ctx, "获取数据成功", result)
}
