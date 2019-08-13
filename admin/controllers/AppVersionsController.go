package controllers

import (
	"qpgame/admin/models"
	"qpgame/admin/validations"
)

var appVersions = models.AppVersions{}                          //模型
var appVersionsValidation = validations.AppVersionsValidation{} //校验器

type AppVersionsController struct{}

/**
 * @api {get} admin/api/auth/v1/app_versions App版本列表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>App版本列表</strong><br />
 * 业务描述: App版本列表</br>
 * @apiVersion 1.0.0
 * @apiName     indexAppVersions
 * @apiGroup    system
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int} 	page			页数
 * @apiParam (客户端请求参数) {int} 	page_size		每页记录数
 * @apiParam (客户端请求参数) {string} 	version 版本号
 * @apiParam (客户端请求参数) {int} 	status 状态
 * @apiParam (客户端请求参数) {int} 	package_type 包类型
 * @apiParam (客户端请求参数) {int} 	app_type APP类型
 * @apiParam (客户端请求参数) {int} 	update_type 更新方式
 * @apiParam (客户端请求参数) {int} 	time_start	添加时间/开始
 * @apiParam (客户端请求参数) {int} 	time_end	添加时间/结束
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
 * @apiSuccess (data-rows每个子对象字段说明) {int}        version               版本号
 * @apiSuccess (data-rows每个子对象字段说明) {int}        status                  版本状态,0:锁定,1:正常
 * @apiSuccess (data-rows每个子对象字段说明) {string}     description             版本说明
 * @apiSuccess (data-rows每个子对象字段说明) {int}        created                 添加时间
 * @apiSuccess (data-rows每个子对象字段说明) {string}     link                    下载地址
 * @apiSuccess (data-rows每个子对象字段说明) {int}        package_type            包类型:1全量包,2增量包
 * @apiSuccess (data-rows每个子对象字段说明) {int}        app_type                APP类型:1安卓,2ios
 * @apiSuccess (data-rows每个子对象字段说明) {int}        update_type             更新类型，1强制更新，2提示更新，3不提示更新
 * @apiSuccess (data-rows每个子对象字段说明) {int}        updated                 修改时间
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "数据获取成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "version": "",
 *        "status": "",
 *        "description": "",
 *        "created": "",
 *        "link": "",
 *        "package_type": "",
 *        "app_type": "",
 *        "update_type": "",
 *        "updated": ""
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *AppVersionsController) Index(ctx *Context) {
	index(ctx, &appVersions)
}

/**
 * @api {post} admin/api/auth/v1/app_versions/add	App版本添加
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>添加/修改App版本</strong><br />
 * 业务描述: 添加/修改App版本 <br />
 * <strong><span style="color: red">注意: </span></strong><br />
 * <span style="color:red">修改操作API不再单独列出, 请参考以下</span><br />
 * <span style="color:red">添加: /admin/api/auth/v1/app_versions/add </span> &nbsp;&nbsp; <br />
 * <span style="color:red">修改: /admin/api/auth/v1/app_versions/update </span>
 * @apiVersion 1.0.0
 * @apiName     saveAppVersions
 * @apiGroup   system
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }

 * @apiParam (客户端请求参数) {int} 	id    					记录编号,仅修改操作时(即接口: /admin/api/auth/v1/app_versions/update)需要<br />
 *															添加操作(即接口: /admin/api/auth/v1/app_versions/add)不需要此参数<br />
 * 															* 如果提供此编号, 则视为修改记录
 * @apiParam (客户端请求参数) {int}      	version                     版本号
 * @apiParam (客户端请求参数) {int}      	status                  版本状态,0:锁定,1:正常
 * @apiParam (客户端请求参数) {string}   	description             版本说明
 * @apiParam (客户端请求参数) {string}   	link                    下载地址
 * @apiParam (客户端请求参数) {int}      	package_type            包类型:1全量包,2增量包
 * @apiParam (客户端请求参数) {int}      	app_type                APP类型:1安卓,2ios
 * @apiParam (客户端请求参数) {int}      	update_type             更新类型，1强制更新，2提示更新，3不提示更新
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
func (self *AppVersionsController) Save(ctx *Context) {
	save(ctx, &appVersions, &appVersionsValidation)
}

/**
 * @api {get} admin/api/auth/v1/app_versions/view 				App版本详情
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>App版本详情</strong><br />
 * 业务描述: App版本详情</br>
 * @apiVersion 1.0.0
 * @apiName     viewAppVersions
 * @apiGroup    system
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
 * @apiSuccess (data字段说明) {int}        version                     版本号
 * @apiSuccess (data字段说明) {int}        status                  版本状态, 0:锁定,1正常
 * @apiSuccess (data字段说明) {string}     description             版本说明
 * @apiSuccess (data字段说明) {int}        created                 添加时间
 * @apiSuccess (data字段说明) {string}     link                    下载地址
 * @apiSuccess (data字段说明) {int}        package_type            包类型:1全量包,2增量包
 * @apiSuccess (data字段说明) {int}        app_type                APP类型:1安卓,2ios
 * @apiSuccess (data字段说明) {int}        update_type             更新类型，1强制更新，2提示更新，3不提示更新
 * @apiSuccess (data字段说明) {int}        updated                 修改时间
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "获取数据成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "version": "",
 *        "status": "",
 *        "description": "",
 *        "created": "",
 *        "link": "",
 *        "package_type": "",
 *        "app_type": "",
 *        "update_type": "",
 *        "updated": ""
 * 	  },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *AppVersionsController) View(ctx *Context) {
	view(ctx, &appVersions)
}

/**
 * @api {get} admin/api/auth/v1/app_versions/delete App版本删除
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>App版本删除</strong><br />
 * 业务描述: App版本删除</br>
 * @apiVersion 1.0.0
 * @apiName     deleteAppVersions
 * @apiGroup   system
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
func (self *AppVersionsController) Delete(ctx *Context) {
	remove(ctx, &appVersions)
}
