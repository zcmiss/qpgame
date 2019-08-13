package controllers

import (
	"qpgame/admin/models"
	"qpgame/admin/validations"
)

var proxyChessLevels = models.ProxyChessLevels{}                          //模型
var proxyChessLevelsValidation = validations.ProxyChessLevelsValidation{} //校验器

type ProxyChessLevelsController struct{}

/**
 * @api {get} admin/api/auth/v1/proxy_chess_levels 棋牌代理等级列表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>代理等级列表</strong><br />
 * 业务描述: 代理等级列表</br>
 * @apiVersion 1.0.0
 * @apiName     indexProxyChessLevels
 * @apiGroup    user
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int} 	page			页数
 * @apiParam (客户端请求参数) {int} 	page_size		每页记录数
 * @apiParam (客户端请求参数) {string}    name       名称
 * @apiParam (客户端请求参数) {string}    time_start 添加时间/开始
 * @apiParam (客户端请求参数) {string}    time_end       添加时间/结束
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
 * @apiSuccess (data-rows每个子对象字段说明) {int}        level                   等级例子(1-10
 * @apiSuccess (data-rows每个子对象字段说明) {string}     name                    等级名称
 * @apiSuccess (data-rows每个子对象字段说明) {int}        team_total_low          团队起始资金单位万/天
 * @apiSuccess (data-rows每个子对象字段说明) {int}        team_total_limit        团队起始资金单位万封顶单位万/天
 * @apiSuccess (data-rows每个子对象字段说明) {int}        commission              万/返佣
 * @apiSuccess (data-rows每个子对象字段说明) {int}        created                 创建时间
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "数据获取成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "level": "",
 *        "name": "",
 *        "team_total_low": "",
 *        "team_total_limit": "",
 *        "commission": "",
 *        "created": ""
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *ProxyChessLevelsController) Index(ctx *Context) {
	index(ctx, &proxyChessLevels)
}

/**
 * @api {post} admin/api/auth/v1/proxy_chess_levels/update	棋牌代理等级修改
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>修改棋牌代理等级</strong><br />
 * 业务描述: 修改棋牌代理等级 <br />
 * @apiVersion 1.0.0
 * @apiName     saveProxyChessLevels
 * @apiGroup    user
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }

 * @apiParam (客户端请求参数) {int} 	id    					记录编号,仅修改操作时(即接口: /admin/api/auth/v1/proxy_chess_levels/update)需要<br />
 * 															* 如果提供此编号, 则视为修改记录
 * @apiParam (客户端请求参数) {int}      	level                   等级例子(1-10
 * @apiParam (客户端请求参数) {string}   	name                    等级名称
 * @apiParam (客户端请求参数) {int}      	team_total_low          团队起始资金单位万/天
 * @apiParam (客户端请求参数) {int}      	team_total_limit        团队起始资金单位万封顶单位万/天
 * @apiParam (客户端请求参数) {int}      	commission              万/返佣
 * @apiParam (客户端请求参数) {int}      	created                 创建时间<br />
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
func (self *ProxyChessLevelsController) Save(ctx *Context) {
	save(ctx, &proxyChessLevels, &proxyChessLevelsValidation)
}

/**
 * @api {get} admin/api/auth/v1/proxy_chess_levels/view 				棋牌代理等级详情
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>代理等级详情</strong><br />
 * 业务描述: 代理等级详情</br>
 * @apiVersion 1.0.0
 * @apiName     viewProxyChessLevels
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
 * @apiSuccess (data字段说明) {int}        level                   等级例子(1-10
 * @apiSuccess (data字段说明) {string}     name                    等级名称
 * @apiSuccess (data字段说明) {int}        team_total_low          团队起始资金单位万/天
 * @apiSuccess (data字段说明) {int}        team_total_limit        团队起始资金单位万封顶单位万/天
 * @apiSuccess (data字段说明) {int}        commission              万/返佣
 * @apiSuccess (data字段说明) {int}        created                 创建时间
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "保存数据成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "level": "",
 *        "name": "",
 *        "team_total_low": "",
 *        "team_total_limit": "",
 *        "commission": "",
 *        "created": ""
 * 	  },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *ProxyChessLevelsController) View(ctx *Context) {
	view(ctx, &proxyChessLevels)
}
