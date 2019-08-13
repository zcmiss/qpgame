package controllers

import (
	"qpgame/admin/models"
	"qpgame/admin/validations"
)

var adminNodes = models.AdminNodes{}                          //模型
var adminNodesValidation = validations.AdminNodesValidation{} //校验器

type AdminNodesController struct{}

/**
 * @api {get} admin/api/auth/v1/admin_nodes 菜单节点列表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>菜单节点列表</strong><br />
 * 业务描述: 菜单节点列表</br>
 * @apiVersion 1.0.0
 * @apiName     indexAdminNodes
 * @apiGroup    admin
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int} 	page			页数
 * @apiParam (客户端请求参数) {int} 	page_size		每页记录数
 * @apiParam (客户端请求参数) {string} 	name 菜单名称
 * @apiParam (客户端请求参数) {string} 	route 路由
 * @apiParam (客户端请求参数) {string} 	method 方法
 * @apiParam (客户端请求参数) {int} 	status 状态
 * @apiParam (客户端请求参数) {int} 	level 级别
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
 * @apiSuccess (data-rows每个子对象字段说明) {string}     title					节点名称, 如 "管理员列表"
 * @apiSuccess (data-rows每个子对象字段说明) {string}     route                 路由, 即url, 如 /index/list
 * @apiSuccess (data-rows每个子对象字段说明) {string}     method				方法，如, Post, Get
 * @apiSuccess (data-rows每个子对象字段说明) {int}        seq					排序, 在同级节点中的排序, 必须是数字
 * @apiSuccess (data-rows每个子对象字段说明) {int}        parent_id				上级节点编号, 必须是数字
 * @apiSuccess (data-rows每个子对象字段说明) {int}        level                 级别, 0:一级菜单,1:二级菜单,2:三级菜单,3:四级菜单
 * @apiSuccess (data-rows每个子对象字段说明) {string}     level_name            级别名称, 如一级菜单, 二级菜单...
 * @apiSuccess (data-rows每个子对象字段说明) {int}		  status				状态, 1:可用, 0:不可用
 * @apiSuccess (data-rows每个子对象字段说明) {int}		  type                  类型
 * @apiSuccess (data-rows每个子对象字段说明) {string}     type_name             类型名称，如主菜单、子菜单、导航(显示在右侧tab页), 动作(添/修/删/查)
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "数据获取成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "name": "",
 *        "route": "",
 *        "title": "",
 *        "method": "",
 *        "status": "",
 *        "seq": "",
 *        "parent_id": "",
 *        "level": "",
 *		  "level_name":"",
 *		  "type_name":"",
 *        "type": ""
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *AdminNodesController) Index(ctx *Context) {
	index(ctx, &adminNodes)
}

/**
 * @api {post} admin/api/auth/v1/admin_nodes/add	菜单节点添加
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>添加/修改菜单节点</strong><br />
 * 业务描述: 添加/修改菜单节点 <br />
 * <strong><span style="color: red">注意: </span></strong><br />
 * <span style="color:red">修改操作API不再单独列出, 请参考以下</span><br />
 * <span style="color:red">添加: /admin/api/auth/v1/admin_nodes/add </span> &nbsp;&nbsp; <br />
 * <span style="color:red">修改: /admin/api/auth/v1/admin_nodes/update </span>
 * @apiVersion 1.0.0
 * @apiName     saveAdminNodes
 * @apiGroup    admin
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }

 * @apiParam (客户端请求参数) {int} 	id    				记录编号,仅修改操作时(即接口: /admin/api/auth/v1/admin_nodes/update)需要<br />
 *															添加操作(即接口: /admin/api/auth/v1/admin_nodes/add)不需要此参数<br />
 * 															* 如果提供此编号, 则视为修改记录
 * @apiParam (客户端请求参数) {string}   	name            节点名称, 用作节点唯一标识<br />如: adminsList 表示管理员列表, 必须是英文或英文数字混合
 * @apiParam (客户端请求参数) {string}   	route           路由, 即url, 如 /index/list
 * @apiParam (客户端请求参数) {string}   	title           节点名称, 如 "管理员列表"
 * @apiParam (客户端请求参数) {string}   	method          方法，如, Post, Get
 * @apiParam (客户端请求参数) {int}      	status          状态, 0:不可用, 1:可用
 * @apiParam (客户端请求参数) {string}   	remark          备注
 * @apiParam (客户端请求参数) {int}      	seq				排序, 在同级节点中的排序, 必须是数字
 * @apiParam (客户端请求参数) {int}      	parent_id       上级节点编号, 必须是数字
 * @apiParam (客户端请求参数) {int}      	level           级别, 0:一级菜单,1:二级菜单,2:三级菜单,3:四级菜单
 * @apiParam (客户端请求参数) {int}			type            类型 <br />0:主菜单 (标签, 无链接, 仅用作显示)
 *															<br />1:子菜单 (标答,有链接,会在右侧导航栏中打开)
 *															<br />2:导航栏 (在右侧导航栏中显示, 需要去重)
 *															<br /><span style="color:red">注意, 以下情况要区别对待:
 *															<br />* 子菜单下只有一个导航栏, 且显示名称一样<br />  如: 子菜单是 "管理员列表", 显示在右侧的导航栏里也是 "管理员列表"
 *															<br />* 子菜单下属有多个导航栏, 显示名称不一样<br />  如: 子菜单是 "代理管理", 下属导航栏有多个: "代理管理" (展示为左侧子菜单)、"代理等级管理"(未展现在左侧子菜单)</span>
 *															<br />3:动作 (如添加/修改/删除/查看等，不显示出来,但有实际功能)
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
func (self *AdminNodesController) Save(ctx *Context) {
	save(ctx, &adminNodes, &adminNodesValidation)
}

/**
 * @api {post} admin/api/auth/v1/admin_nodes/sort	菜单节点排序
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: lyndon</span><br/><br/>
 * <strong>菜单节点排序</strong><br />
 * 业务描述: 菜单节点排序，可批量 <br />
 * <strong><span style="color: red">注意: </span></strong><br />
 * <span style="color:red">修改操作API不再单独列出, 请参考以下</span><br />
 * <span style="color:red">添加: /admin/api/auth/v1/admin_nodes/add </span> &nbsp;&nbsp; <br />
 * <span style="color:red">修改: /admin/api/auth/v1/admin_nodes/update </span>
 * @apiVersion 1.0.0
 * @apiName     sortAdminNodes
 * @apiGroup    admin
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }

 * @apiParam (客户端请求参数) {array}   	list            批量的 编号&序号列表数据
 * @apiParam (客户端请求参数list的子项) {string}   	id      节点编号
 * @apiParam (客户端请求参数list的子项) {string}   	sort    节点序号
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
func (self *AdminNodesController) Sort(ctx *Context) {
	err := adminNodes.Sort(ctx)
	if err != nil{
		responseFailure(ctx, "", err.Error())
		return
	}
	responseSuccess(ctx, "操作成功", nil)
}

/**
 * @api {get} admin/api/auth/v1/admin_nodes/view 				菜单节点详情
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>菜单节点详情</strong><br />
 * 业务描述: 菜单节点详情</br>
 * @apiVersion 1.0.0
 * @apiName     viewAdminNodes
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
 * @apiSuccess (data字段说明) {string}     name                    节点名称
 * @apiSuccess (data字段说明) {string}     route                   路由
 * @apiSuccess (data字段说明) {string}     title                   说明
 * @apiSuccess (data字段说明) {string}     method                  方法
 * @apiSuccess (data字段说明) {int}        status                  状态
 * @apiSuccess (data字段说明) {string}     remark                  备注
 * @apiSuccess (data字段说明) {int}        seq						排序
 * @apiSuccess (data字段说明) {int}        parent_id               上级编号
 * @apiSuccess (data字段说明) {int}        level                   级别
 * @apiSuccess (data字段说明) {string}     type                    类型
 *
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "保存数据成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "name": "",
 *        "route": "",
 *        "title": "",
 *        "method": "",
 *        "status": "",
 *        "remark": "",
 *        "seq": "",
 *        "parent_id": "",
 *        "level": "",
 *        "type": ""
 * 	  },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *AdminNodesController) View(ctx *Context) {
	view(ctx, &adminNodes)
}

/**
 * @api {get} admin/api/auth/v1/admin_nodes/delete 菜单节点删除
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>菜单节点删除</strong><br />
 * 业务描述: 菜单节点删除</br>
 * @apiVersion 1.0.0
 * @apiName     deleteAdminNodes
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
func (self *AdminNodesController) Delete(ctx *Context) {
	remove(ctx, &adminNodes)
}
