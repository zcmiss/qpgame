package controllers

import (
	"qpgame/admin/models"
	"qpgame/admin/validations"
)

var platformGames = models.PlatformGames{}                          //模型
var platformGamesValidation = validations.PlatformGamesValidation{} //校验器

type PlatformGamesController struct{}

/**
 * @api {get} admin/api/auth/v1/platform_games 平台游戏列表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>平台游戏列表</strong><br />
 * 业务描述: 平台游戏列表</br>
 * @apiVersion 1.0.0
 * @apiName     indexPlatformGames
 * @apiGroup    game
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int} 	page			页数
 * @apiParam (客户端请求参数) {int} 	page_size		每页记录数
 * @apiParam (客户端请求参数) {int} 	page_size		每页记录数
 * @apiParam (客户端请求参数) {int} 	game_categorie_id 游戏分类编号
 * @apiParam (客户端请求参数) {string} 	name 游戏名称
 * @apiParam (客户端请求参数) {int} 	ishot		是否热门
 * @apiParam (客户端请求参数) {int} 	isnew 		是否新游戏
 * @apiParam (客户端请求参数) {int} 	plat_id 	游戏平台
 * @apiParam (客户端请求参数) {int} 	has_img 	是否有图片,0:无图片,1:有图片
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
 * @apiSuccess (data-rows每个子对象字段说明) {int}		  id 					  记录
 * @apiSuccess (data-rows每个子对象字段说明) {int}        game_categorie_id       游戏分类编号
 * @apiSuccess (data-rows每个子对象字段说明) {int}        category_name      游戏分类
 * @apiSuccess (data-rows每个子对象字段说明) {string}     game_url                游戏地址
 * @apiSuccess (data-rows每个子对象字段说明) {string}     gamecode                游戏编码
 * @apiSuccess (data-rows每个子对象字段说明) {string}     gt                      游戏分类
 * @apiSuccess (data-rows每个子对象字段说明) {string}     img                     方形游戏图片资源
 * @apiSuccess (data-rows每个子对象字段说明) {int}        ishot                   是否热门
 * @apiSuccess (data-rows每个子对象字段说明) {int}        isnew                   是否新游戏
 * @apiSuccess (data-rows每个子对象字段说明) {int}        isrecommend             是否推荐游戏
 * @apiSuccess (data-rows每个子对象字段说明) {string}     name                    游戏名称
 * @apiSuccess (data-rows每个子对象字段说明) {int}        plat_id                 平台编号
 * @apiSuccess (data-rows每个子对象字段说明) {int}        platform_name                 平台名称
 * @apiSuccess (data-rows每个子对象字段说明) {string}        service_code              游戏代码
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "数据获取成功",
 *    "code": 200,
 *    "data": {
 *        "game_categorie_id": "",
 *        "game_category_name": "",
 *        "game_url": "",
 *        "gamecode": "",
 *        "gt": "",
 *        "id": "",
 *        "img": "",
 *        "ishot": "",
 *        "isnew": "",
 *        "isrecommend": "",
 *        "name": "",
 *        "plat_id": "",
 *        "platform_name": "",
 *        "service_code": "",
 *        "small_img": "",
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *PlatformGamesController) Index(ctx *Context) {
	index(ctx, &platformGames)
}

/**
 * @api {post} admin/api/auth/v1/platform_games/add	平台游戏添加
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>添加/修改平台游戏</strong><br />
 * 业务描述: 添加/修改平台游戏 <br />
 * <strong><span style="color: red">注意: </span></strong><br />
 * <span style="color:red">修改操作API不再单独列出, 请参考以下</span><br />
 * <span style="color:red">添加: /admin/api/auth/v1/platform_games/add </span> &nbsp;&nbsp; <br />
 * <span style="color:red">修改: /admin/api/auth/v1/platform_games/update </span>
 * @apiVersion 1.0.0
 * @apiName     savePlatformGames
 * @apiGroup    game
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }

 * @apiParam (客户端请求参数) {int} 	id    					记录编号,仅修改操作时(即接口: /admin/api/auth/v1/platform_games/update)需要<br />
 *															添加操作(即接口: /admin/api/auth/v1/platform_games/add)不需要此参数<br />
 * 															* 如果提供此编号, 则视为修改记录
 * @apiParam (客户端请求参数) {int}      	game_categorie_id       游戏分类
 * @apiParam (客户端请求参数) {int}      	plat_game_id       游戏平台分类
 * @apiParam (客户端请求参数) {string}   	game_url                游戏地址
 * @apiParam (客户端请求参数) {string}   	gamecode                游戏编码
 * @apiParam (客户端请求参数) {string}   	gt                      游戏分类
 * @apiParam (客户端请求参数) {string}   	img                     方形游戏图片资源
 * @apiParam (客户端请求参数) {int}      	ishot                   是否热门
 * @apiParam (客户端请求参数) {int}      	isnew                   是否新游戏
 * @apiParam (客户端请求参数) {int}      	isrecommend             是否推荐游戏
 * @apiParam (客户端请求参数) {string}   	name                    游戏名称
 * @apiParam (客户端请求参数) {int}      	plat_id                 平台编号
 * @apiParam (客户端请求参数) {int}      	service_code              游戏代码
 * @apiParam (客户端请求参数) {string}   	small_img               圆形游戏图片资源<br />
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
func (self *PlatformGamesController) Save(ctx *Context) {
	save(ctx, &platformGames, &platformGamesValidation)
}

/**
 * @api {get} admin/api/auth/v1/platform_games/view 				平台游戏详情
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>平台游戏详情</strong><br />
 * 业务描述: 平台游戏详情</br>
 * @apiVersion 1.0.0
 * @apiName     viewPlatformGames
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
 * @apiSuccess (data字段说明) {int}        game_categorie_id       游戏分类ID
 * @apiSuccess (data字段说明) {string}     game_url                游戏地址
 * @apiSuccess (data字段说明) {string}     gamecode                游戏编码
 * @apiSuccess (data字段说明) {string}     gt                      游戏分类
 * @apiSuccess (data字段说明) {string}     img                     方形游戏图片资源
 * @apiSuccess (data字段说明) {int}        ishot                   是否热门
 * @apiSuccess (data字段说明) {int}        isnew                   是否新游戏
 * @apiSuccess (data字段说明) {int}        isrecommend             是否推荐游戏
 * @apiSuccess (data字段说明) {string}     name                    游戏名称
 * @apiSuccess (data字段说明) {int}        plat_id                 平台编号
 * @apiSuccess (data字段说明) {string}        service_code              游戏代码
 * @apiSuccess (data字段说明) {string}     small_img               图片
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "保存数据成功",
 *    "code": 200,
 *    "data": {
 *        "game_categorie_id": "",
 *        "game_url": "",
 *        "gamecode": "",
 *        "gt": "",
 *        "id": "",
 *        "img": "",
 *        "ishot": "",
 *        "isnew": "",
 *        "isrecommend": "",
 *        "name": "",
 *        "plat_id": "",
 *        "service_code": "",
 *        "small_img": "",
 * 	  },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *PlatformGamesController) View(ctx *Context) {
	view(ctx, &platformGames)
}

/**
 * @api {get} admin/api/auth/v1/platform_games/delete 平台游戏删除
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>平台游戏删除</strong><br />
 * 业务描述: 平台游戏删除</br>
 * @apiVersion 1.0.0
 * @apiName     deletePlatformGames
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
func (self *PlatformGamesController) Delete(ctx *Context) {
	remove(ctx, &platformGames)
}
