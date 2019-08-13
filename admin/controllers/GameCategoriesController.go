package controllers

import (
	"qpgame/admin/models"
	"qpgame/admin/validations"
)

var gameCategories = models.GameCategories{}                          //模型
var gameCategoriesValidation = validations.GameCategoriesValidation{} //校验器

type GameCategoriesController struct{}

/**
 * @api {get} admin/api/auth/v1/game_categories 游戏分类列表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>游戏分类列表</strong><br />
 * 业务描述: 游戏分类列表</br>
 * @apiVersion 1.0.0
 * @apiName     indexGameCategories
 * @apiGroup    game
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int} 	page			页数
 * @apiParam (客户端请求参数) {int} 	page_size		每页记录数
 * @apiParam (客户端请求参数) {string} 	name 	分类名称
 * @apiParam (客户端请求参数) {int} 	parent_id 	上级分类
 * @apiParam (客户端请求参数) {int} 	status 	状态
 * @apiParam (客户端请求参数) {int} 	platform_id 游戏平台
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
 * @apiSuccess (data-rows每个子对象字段说明) {int}        platform_id 		游戏平台编号
 * @apiSuccess (data-rows每个子对象字段说明) {string}        platform_name 	游戏平台名称
 * @apiSuccess (data-rows每个子对象字段说明) {int}        category_level          分类层级(最多三次
 * @apiSuccess (data-rows每个子对象字段说明) {int}        created                 创建时间
 * @apiSuccess (data-rows每个子对象字段说明) {string}     img                     2级分类游戏图片
 * @apiSuccess (data-rows每个子对象字段说明) {string}     name                    游戏分类名称或者游戏名称
 * @apiSuccess (data-rows每个子对象字段说明) {int}        parent_id               上级游戏分类id
 * @apiSuccess (data-rows每个子对象字段说明) {int}        parent_name               上级游戏分类名称
 * @apiSuccess (data-rows每个子对象字段说明) {int}        seq					分类排序
 * @apiSuccess (data-rows每个子对象字段说明) {int}        status                  分类状态,0不可用 1可用
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "数据获取成功",
 *    "code": 200,
 *    "data": {
 *        "platform_id": "",
 *        "platform_name": "",
 *        "category_level": "",
 *        "created": "",
 *        "id": "",
 *        "img": "",
 *        "name": "",
 *        "parent_id": "",
 *        "parent_name": "",
 *        "seq": "",
 *        "status": ""
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *GameCategoriesController) Index(ctx *Context) {
	index(ctx, &gameCategories)
}

/**
 * @api {post} admin/api/auth/v1/game_categories/add	游戏分类添加
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>添加/修改游戏分类</strong><br />
 * 业务描述: 添加/修改游戏分类 <br />
 * <strong><span style="color: red">注意: </span></strong><br />
 * <span style="color:red">修改操作API不再单独列出, 请参考以下</span><br />
 * <span style="color:red">添加: /admin/api/auth/v1/game_categories/add </span> &nbsp;&nbsp; <br />
 * <span style="color:red">修改: /admin/api/auth/v1/game_categories/update </span>
 * @apiVersion 1.0.0
 * @apiName     saveGameCategories
 * @apiGroup    game
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }

 * @apiParam (客户端请求参数) {int} 	id    					记录编号,仅修改操作时(即接口: /admin/api/auth/v1/game_categories/update)需要<br />
 *															添加操作(即接口: /admin/api/auth/v1/game_categories/add)不需要此参数<br />
 * 															* 如果提供此编号, 则视为修改记录
 * @apiParam (客户端请求参数) {int}      	category_level          分类层级(最多三次
 * @apiParam (客户端请求参数) {int}      	platform_id 	游戏平台编号
 * @apiParam (客户端请求参数) {string}   	img                     2级分类游戏图片
 * @apiParam (客户端请求参数) {string}   	name                    游戏分类名称或者游戏名称
 * @apiParam (客户端请求参数) {int}      	parent_id               上级游戏分类id
 * @apiParam (客户端请求参数) {int}      	seq					分类排序
 * @apiParam (客户端请求参数) {int}      	status                  分类状态,0不可用 1可用
 * @apiParam (客户端请求参数) {string}      	btn_img 			按钮图片
 * @apiParam (客户端请求参数) {string}      	btn_selected_img	按钮选中图片
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
func (self *GameCategoriesController) Save(ctx *Context) {
	save(ctx, &gameCategories, &gameCategoriesValidation)
}

/**
 * @api {get} admin/api/auth/v1/game_categories/view 				游戏分类详情
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>游戏分类详情</strong><br />
 * 业务描述: 游戏分类详情</br>
 * @apiVersion 1.0.0
 * @apiName     viewGameCategories
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
 * @apiSuccess (data字段说明) {int}		  platform_id	游戏平台编号
 * @apiSuccess (data字段说明) {int}        category_level          分类层级(最多三次
 * @apiSuccess (data字段说明) {int}        created                 创建时间
 * @apiSuccess (data字段说明) {string}     img                     2级分类游戏图片
 * @apiSuccess (data字段说明) {string}     name                    游戏分类名称或者游戏名称
 * @apiSuccess (data字段说明) {int}        parent_id               上级游戏分类id
 * @apiSuccess (data字段说明) {int}        seq				分类排序
 * @apiSuccess (data字段说明) {int}        status                  分类状态,0不可用 1可用
 * @apiParam (客户端请求参数) {string}      	btn_img 			按钮图片
 * @apiParam (客户端请求参数) {string}      	btn_selected_img	按钮选中图片
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "保存数据成功",
 *    "code": 200,
 *    "data": {
 *        "platform_id": "",
 *        "category_level": "",
 *        "created": "",
 *        "id": "",
 *        "img": "",
 *        "name": "",
 *        "parent_id": "",
 *        "seq": "",
 *        "status": "",
 *		  "btn_img": "",
 *  	  "btn_selected_img": "",
 * 	  },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *GameCategoriesController) View(ctx *Context) {
	view(ctx, &gameCategories)
}

/**
 * @api {get} admin/api/auth/v1/game_categories/delete 游戏分类删除
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>游戏分类删除</strong><br />
 * 业务描述: 游戏分类删除</br>
 * @apiVersion 1.0.0
 * @apiName     deleteGameCategories
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
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "记录删除成功",
 *    "code": 200,
 *    "data": {},
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *GameCategoriesController) Delete(ctx *Context) {
	remove(ctx, &gameCategories)
}

/**
 * @api {get} admin/api/auth/v1/game_categories/relations 游戏分类关联列表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>游戏分类列表-关联列表</strong><br />
 * 业务描述: 游戏分类列表-关联列表</br>
 * @apiVersion 1.0.0
 * @apiName     indexGameCategoriesRelations
 * @apiGroup    game
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int} 	id		分类编号
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
 * @apiSuccess (data字段说明) {int}  	id        分类编号
 * @apiSuccess (data字段说明) {string}	name		分类名称
 * @apiSuccess (data字段说明) {string}	name		分类层级
 * @apiSuccess (data字段说明) {array}   	categories	子分类, 结构和 data 字段结一样
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "数据获取成功",
 *    "code": 200,
 *    "data": [{
 *        "id": "",
 *        "name": "",
 *        "categories": [],
 *    }],
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *GameCategoriesController) Relations(ctx *Context) {
	categories, err := gameCategories.Relations(ctx)
	if err != nil {
		responseFailure(ctx, "", err.Error())
		return
	}
	responseSuccess(ctx, "", categories)
}
