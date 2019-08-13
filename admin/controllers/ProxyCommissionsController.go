package controllers

import (
	"qpgame/admin/models"
)

var proxyCommissions = models.ProxyCommissions{} //模型

type ProxyCommissionsController struct{}

/**
 * @api {get} admin/api/auth/v1/proxy_commissions 代理佣金列表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>代理佣金列表</strong><br />
 * 业务描述: 代理佣金列表</br>
 * @apiVersion 1.0.0
 * @apiName     indexProxyCommissions
 * @apiGroup    finance
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int}    page            页数
 * @apiParam (客户端请求参数) {int}   page_size       每页记录数
 * @apiParam (客户端请求参数) {int}   user_id 用户编号
 * @apiParam (客户端请求参数) {string}   created_start 时间/开始
 * @apiParam (客户端请求参数) {string}   created_end 	时间/结束
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
 * @apiSuccess (data-rows每个子对象字段说明) {int}		  id 	记录编号
 * @apiSuccess (data-rows每个子对象字段说明) {int}        user_id                 用户编号
 * @apiSuccess (data-rows每个子对象字段说明) {string}        user_name                用户名称
 * @apiSuccess (data-rows每个子对象字段说明) {int}        parent_id               上级代理编号
 * @apiSuccess (data-rows每个子对象字段说明) {string}        parent_name				上级代理名称
 * @apiSuccess (data-rows每个子对象字段说明) {float}      bet_amount              个人总业绩
 * @apiSuccess (data-rows每个子对象字段说明) {float}      total_amount            团队总业绩
 * @apiSuccess (data-rows每个子对象字段说明) {int}        created                 数据生成时间
 * @apiSuccess (data-rows每个子对象字段说明) {int}        contributions           业绩贡献人数，不包括自己
 * @apiSuccess (data-rows每个子对象字段说明) {int}        proxy_type              代理类型2棋牌5真人
 * @apiSuccess (data-rows每个子对象字段说明) {int}        proxy_level             代理等级
 * @apiSuccess (data-rows每个子对象字段说明) {float}      proxy_level_rate        代理等级对应返水率
 * @apiSuccess (data-rows每个子对象字段说明) {string}     proxy_level_name        代理等级名称
 * @apiSuccess (data-rows每个子对象字段说明) {float}      commission              个人总佣金
 * @apiSuccess (data-rows每个子对象字段说明) {float}      total_commission        团队总佣金
 * @apiSuccess (data-rows每个子对象字段说明) {string}     created_str             佣金生成的时间字符串
 * @apiSuccess (data-rows每个子对象字段说明) {int}        states                  佣金领取状态0未领1已领
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "数据获取成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "user_id": "",
 *        "user_name": "",
 *        "parent_id": "",
 *        "parent_name": "",
 *        "bet_amount": "",
 *        "total_amount": "",
 *        "created": "",
 *        "contributions": "",
 *        "proxy_type": "",
 *        "proxy_level": "",
 *        "proxy_level_rate": "",
 *        "proxy_level_name": "",
 *        "commission": "",
 *        "total_commission": "",
 *        "created_str": "",
 *        "states": ""
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *ProxyCommissionsController) Index(ctx *Context) {
	index(ctx, &proxyCommissions)
}

/**
 * @api {get} admin/api/v1/proxy_commissions/view 				代理佣金详情
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>代理佣金详情</strong><br />
 * 业务描述: 代理佣金详情</br>
 * @apiVersion 1.0.0
 * @apiName     viewProxyCommissions
 * @apiGroup    finance
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 *
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
 * @apiSuccess (返回结果)  {json}  	data            返回数据
 * @apiSuccess (返回结果)  {float}   	timeConsumed    后台耗时
 *
 * @apiSuccess (data字段说明) {int}		  id 					  记录编号
 * @apiSuccess (data字段说明) {int}        user_id                 用户编号
 * @apiSuccess (data字段说明) {int}        parent_id               上级代理用户编号
 * @apiSuccess (data字段说明) {float}      bet_amount              个人总业绩
 * @apiSuccess (data字段说明) {float}      total_amount            团队总业绩
 * @apiSuccess (data字段说明) {int}        created                 数据生成时间
 * @apiSuccess (data字段说明) {int}        contributions           业绩贡献人数，不包括自己
 * @apiSuccess (data字段说明) {int}        proxy_type              代理类型2棋牌5真人
 * @apiSuccess (data字段说明) {int}        proxy_level             代理等级
 * @apiSuccess (data字段说明) {float}      proxy_level_rate        代理等级对应返水率
 * @apiSuccess (data字段说明) {string}     proxy_level_name        代理等级名称
 * @apiSuccess (data字段说明) {float}      commission              个人总佣金
 * @apiSuccess (data字段说明) {float}      total_commission        团队总佣金
 * @apiSuccess (data字段说明) {string}     created_str             佣金生成的时间字符串
 * @apiSuccess (data字段说明) {int}        states                  佣金领取状态0未领1已领
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "保存数据成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "user_id": "",
 *        "parent_id": "",
 *        "bet_amount": "",
 *        "total_amount": "",
 *        "created": "",
 *        "contributions": "",
 *        "proxy_type": "",
 *        "proxy_level": "",
 *        "proxy_level_rate": "",
 *        "proxy_level_name": "",
 *        "commission": "",
 *        "total_commission": "",
 *        "created_str": "",
 *        "states": ""
 * 	  },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *ProxyCommissionsController) View(ctx *Context) {
	view(ctx, &proxyCommissions)
}

/**
 * @api {get} admin/api/v1/proxy_commissions/delete 代理佣金删除
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>代理佣金删除</strong><br />
 * 业务描述: 代理佣金删除</br>
 * @apiVersion 1.0.0
 * @apiName     deleteProxyCommissions
 * @apiGroup    finance
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 *
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
 * @apiSuccess (返回结果)  {json}  	data            返回数据
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
func (self *ProxyCommissionsController) Delete(ctx *Context) {
	remove(ctx, &proxyCommissions)
}
