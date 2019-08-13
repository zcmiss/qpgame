package controllers

import (
	"qpgame/admin/models"
	"qpgame/admin/validations"
)

var userGroups = models.UserGroups{}                          //模型
var userGroupsValidation = validations.UserGroupsValidation{} //校验器

type UserGroupsController struct{}

/**
 * @api {get} admin/api/auth/v1/user_groups 用户分组列表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>用户分组列表</strong><br />
 * 业务描述: 用户分组列表</br>
 * @apiVersion 1.0.0
 * @apiName     indexUserGroups
 * @apiGroup    user
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int}     page            页数
 * @apiParam (客户端请求参数) {int}    page_size       每页记录数
 * @apiParam (客户端请求参数) {string}    group_name 名称
 * @apiParam (客户端请求参数) {int}    is_default 是否默认
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
 * @apiSuccess (data-rows每个子对象字段说明) {int}		id				记录编号
 * @apiSuccess (data-rows每个子对象字段说明) {string}   group_name		组名称
 * @apiSuccess (data-rows每个子对象字段说明) {string}   remark			备注
 * @apiSuccess (data-rows每个子对象字段说明) {int}		is_default		是否默认
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "数据获取成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "name": "",
 *        "logo": ""
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *UserGroupsController) Index(ctx *Context) {
	index(ctx, &userGroups)
}

/**
 * @api {post} admin/api/auth/v1/user_groups/add	用户分组添加
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>添加/修改用户分组</strong><br />
 * 业务描述: 添加/修改用户分组 <br />
 * <strong><span style="color: red">注意: </span></strong><br />
 * <span style="color:red">修改操作API不再单独列出, 请参考以下</span><br />
 * <span style="color:red">添加: /admin/api/auth/v1/user_groups/add </span> &nbsp;&nbsp; <br />
 * <span style="color:red">修改: /admin/api/auth/v1/user_groups/update </span>
 * @apiVersion 1.0.0
 * @apiName     saveUserGroups
 * @apiGroup    user
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }

 * @apiParam (客户端请求参数) {int}			id    		记录编号,仅修改操作时(即接口: /admin/api/auth/v1/user_groups/update)需要<br />
 *														添加操作(即接口: /admin/api/auth/v1/user_groups/add)不需要此参数<br />
 * 														* 如果提供此编号, 则视为修改记录
 * @apiParam (客户端请求参数) {string}   	group_name  组名称
 * @apiParam (客户端请求参数) {string}   	remark		备注
 * @apiParam (客户端请求参数) {string}   	is_default	是否默认
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
func (self *UserGroupsController) Save(ctx *Context) {
	save(ctx, &userGroups, &userGroupsValidation)
}

/**
 * @api {get} admin/api/auth/v1/user_groups/view 				用户分组详情
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>用户分组详情</strong><br />
 * 业务描述: 用户分组详情</br>
 * @apiVersion 1.0.0
 * @apiName     viewUserGroups
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
 * @apiSuccess (data字段说明) {int}		id				记录编号
 * @apiSuccess (data字段说明) {string}  group_name      组名称
 * @apiSuccess (data字段说明) {string}  remark			备注
 * @apiSuccess (data字段说明) {int}		is_default		是否默认
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "保存数据成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "name": "",
 *        "logo": ""
 * 	  },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *UserGroupsController) View(ctx *Context) {
	view(ctx, &userGroups)
}

/**
 * @api {get} admin/api/auth/v1/user_groups/delete 用户分组删除
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>用户分组删除</strong><br />
 * 业务描述: 用户分组删除</br>
 * @apiVersion 1.0.0
 * @apiName     deleteUserGroups
 * @apiGroup    user
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }

 * @apiParam (客户端请求参数) {int} 	id    			记录编号
 *
 * @apiError (请求失败返回) {int}      code				错误代码
 * @apiError (请求失败返回) {string}   clientMsg		提示信息
 * @apiError (请求失败返回) {string}   internalMsg		内部错误信息
 * @apiError (请求失败返回) {float}    timeConsumed		后台耗时
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
func (self *UserGroupsController) Delete(ctx *Context) {
	remove(ctx, &userGroups)
}
