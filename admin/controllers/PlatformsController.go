package controllers

import (
	"qpgame/admin/models"
	"qpgame/admin/validations"
)

var platforms = models.Platforms{}                          //模型
var platformsValidation = validations.PlatformsValidation{} //校验器

type PlatformsController struct{}

/**
 * @api {get} admin/api/auth/v1/platforms 游戏平台列表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>游戏平台列表</strong><br />
 * 业务描述: 游戏平台列表</br>
 * @apiVersion 1.0.0
 * @apiName     indexPlatforms
 * @apiGroup    game
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int} 	page			页数
 * @apiParam (客户端请求参数) {int} 	page_size		每页记录数
 * @apiParam (客户端请求参数) {int} 	status 状态
 * @apiParam (客户端请求参数) {string} 	name 名称
 * @apiParam (客户端请求参数) {string} 	name 代码
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
 * @apiSuccess (data-rows每个子对象字段说明) {string}     name                    平台名
 * @apiSuccess (data-rows每个子对象字段说明) {string}     code                    平台代号
 * @apiSuccess (data-rows每个子对象字段说明) {int}        status                  平台状态，0是锁定 1是正常 2是维护中 3是敬请期待
 * @apiSuccess (data-rows每个子对象字段说明) {string}        status_name                  平台状态说明
 * @apiSuccess (data-rows每个子对象字段说明) {string}     logo                    平台logo
 * @apiSuccess (data-rows每个子对象字段说明) {string}     index_logo              平台首页logo,区别平台logo,有大小尺寸之分
 * @apiSuccess (data-rows每个子对象字段说明) {string}     content                 平台介绍
 * @apiSuccess (data-rows每个子对象字段说明) {int}        sort                    排序
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "数据获取成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "name": "",
 *        "code": "",
 *        "status": "",
 *        "status_name": "",
 *        "logo": "",
 *        "index_logo": "",
 *        "content": "",
 *        "sort": ""
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *PlatformsController) Index(ctx *Context) {
	index(ctx, &platforms)
}

/**
 * @api {post} admin/api/auth/v1/platforms/add	游戏平台添加
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>添加/修改游戏平台</strong><br />
 * 业务描述: 添加/修改游戏平台 <br />
 * <strong><span style="color: red">注意: </span></strong><br />
 * <span style="color:red">修改操作API不再单独列出, 请参考以下</span><br />
 * <span style="color:red">添加: /admin/api/auth/v1/platforms/add </span> &nbsp;&nbsp; <br />
 * <span style="color:red">修改: /admin/api/auth/v1/platforms/update </span>
 * @apiVersion 1.0.0
 * @apiName     savePlatforms
 * @apiGroup    game
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }

 * @apiParam (客户端请求参数) {int} 	id    					记录编号,仅修改操作时(即接口: /admin/api/auth/v1/platforms/update)需要<br />
 *															添加操作(即接口: /admin/api/auth/v1/platforms/add)不需要此参数<br />
 * 															* 如果提供此编号, 则视为修改记录
 * @apiParam (客户端请求参数) {string}   	name                    平台名
 * @apiParam (客户端请求参数) {string}   	code                    平台代号
 * @apiParam (客户端请求参数) {int}      	status                  平台状态，0不可用，1可用
 * @apiParam (客户端请求参数) {string}   	logo                    平台logo
 * @apiParam (客户端请求参数) {string}   	index_logo              平台首页logo,区别平台logo,有大小尺寸之分
 * @apiParam (客户端请求参数) {string}   	content                 平台介绍
 * @apiParam (客户端请求参数) {int}      	sort                    排序<br />
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
func (self *PlatformsController) Save(ctx *Context) {
	save(ctx, &platforms, &platformsValidation)
}

/**
 * @api {get} admin/api/auth/v1/platforms/view 				游戏平台详情
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>游戏平台详情</strong><br />
 * 业务描述: 游戏平台详情</br>
 * @apiVersion 1.0.0
 * @apiName     viewPlatforms
 * @apiGroup    game
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
 * @apiSuccess (data字段说明) {string}     name                    平台名
 * @apiSuccess (data字段说明) {string}     code                    平台代号
 * @apiSuccess (data字段说明) {int}        status                  平台状态，0不可用，1可用
 * @apiSuccess (data字段说明) {string}     logo                    平台logo
 * @apiSuccess (data字段说明) {string}     index_logo              平台首页logo,区别平台logo,有大小尺寸之分
 * @apiSuccess (data字段说明) {string}     content                 平台介绍
 * @apiSuccess (data字段说明) {int}        sort                    排序
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "保存数据成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "name": "",
 *        "code": "",
 *        "status": "",
 *        "logo": "",
 *        "index_logo": "",
 *        "content": "",
 *        "sort": ""
 * 	  },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *PlatformsController) View(ctx *Context) {
	view(ctx, &platforms)
}

/**
 * @api {get} admin/api/auth/v1/platforms/delete 游戏平台删除
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>游戏平台删除</strong><br />
 * 业务描述: 游戏平台删除</br>
 * @apiVersion 1.0.0
 * @apiName     deletePlatforms
 * @apiGroup    game
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
func (self *PlatformsController) Delete(ctx *Context) {
	remove(ctx, &platforms)
}
