package controllers

import (
	"qpgame/admin/models"
	"qpgame/admin/validations"
)

var vipLevels = models.VipLevels{}                          //模型
var vipLevelsValidation = validations.VipLevelsValidation{} //校验器

type VipLevelsController struct{}

/**
 * @api {get} admin/api/auth/v1/vip_levels VIP等级列表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>VIP等级列表</strong><br />
 * 业务描述: VIP等级列表</br>
 * @apiVersion 1.0.0
 * @apiName     indexVipLevels
 * @apiGroup    user
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int}     page            页数
 * @apiParam (客户端请求参数) {int}    page_size       每页记录数
 * @apiParam (客户端请求参数) {string}    name 	名称
 * @apiParam (客户端请求参数) {int}    has_deposit_speed 存款加速通道
 * @apiParam (客户端请求参数) {int}    has_own_service 专属客服经理
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
 * @apiSuccess (data-rows每个子对象字段说明) {int}        valid_bet_min           有效投注金额区间起点单位万
 * @apiSuccess (data-rows每个子对象字段说明) {int}        valid_bet_max           有效投注金额区间封顶单位万
 * @apiSuccess (data-rows每个子对象字段说明) {int}        upgrade_amount          晋级礼金
 * @apiSuccess (data-rows每个子对象字段说明) {int}        weekly_amount           周礼金
 * @apiSuccess (data-rows每个子对象字段说明) {int}        month_amount            月俸禄
 * @apiSuccess (data-rows每个子对象字段说明) {int}        upgrade_amount_total    累计晋级礼金
 * @apiSuccess (data-rows每个子对象字段说明) {int}        has_deposit_speed       存款加速通道(0不支持,1支持
 * @apiSuccess (data-rows每个子对象字段说明) {int}        has_own_service         专属客服经理(0没有,1有
 * @apiSuccess (data-rows每个子对象字段说明) {float}      wash_code               洗码率
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "数据获取成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "level": "",
 *        "name": "",
 *        "valid_bet_min": "",
 *        "valid_bet_max": "",
 *        "upgrade_amount": "",
 *        "weekly_amount": "",
 *        "month_amount": "",
 *        "upgrade_amount_total": "",
 *        "has_deposit_speed": "",
 *        "has_own_service": "",
 *        "wash_code": "",
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *VipLevelsController) Index(ctx *Context) {
	index(ctx, &vipLevels)
}

/**
 * @api {post} admin/api/auth/v1/vip_levels/update	VIP等级修改
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>修改VIP等级</strong><br />
 * 业务描述: 修改VIP等级 <br />
 * <strong><span style="color: red">注意: </span></strong><br />
 * <span style="color:red">修改: /admin/api/auth/v1/vip_levels/update </span>
 * @apiVersion 1.0.0
 * @apiName     saveVipLevels
 * @apiGroup    user
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }

 * @apiParam (客户端请求参数) {int} 	id    					编号,仅修改操作时(即接口: /admin/api/auth/v1/vip_levels/update)需要
 * @apiParam (客户端请求参数) {int}      	level                   等级例子(1-10
 * @apiParam (客户端请求参数) {string}   	name                    等级名称
 * @apiParam (客户端请求参数) {int}      	valid_bet_min           有效投注金额区间起点单位万
 * @apiParam (客户端请求参数) {int}      	valid_bet_max           有效投注金额区间封顶单位万
 * @apiParam (客户端请求参数) {int}      	upgrade_amount          晋级礼金
 * @apiParam (客户端请求参数) {int}      	weekly_amount           周礼金
 * @apiParam (客户端请求参数) {int}      	month_amount            月俸禄
 * @apiParam (客户端请求参数) {int}      	upgrade_amount_total    累计晋级礼金
 * @apiParam (客户端请求参数) {int}      	has_deposit_speed       存款加速通道(0不支持,1支持
 * @apiParam (客户端请求参数) {int}      	has_own_service         专属客服经理(0没有,1有
 * @apiParam (客户端请求参数) {float}    	wash_code               洗码率
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
func (self *VipLevelsController) Save(ctx *Context) {
	save(ctx, &vipLevels, &vipLevelsValidation)
}

/**
 * @api {get} admin/api/auth/v1/vip_levels/view 				VIP等级详情
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>VIP等级详情</strong><br />
 * 业务描述: VIP等级详情</br>
 * @apiVersion 1.0.0
 * @apiName     viewVipLevels
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
 * @apiSuccess (data字段说明) {int}        valid_bet_min           有效投注金额区间起点单位万
 * @apiSuccess (data字段说明) {int}        valid_bet_max           有效投注金额区间封顶单位万
 * @apiSuccess (data字段说明) {int}        upgrade_amount          晋级礼金
 * @apiSuccess (data字段说明) {int}        weekly_amount           周礼金
 * @apiSuccess (data字段说明) {int}        month_amount            月俸禄
 * @apiSuccess (data字段说明) {int}        upgrade_amount_total    累计晋级礼金
 * @apiSuccess (data字段说明) {int}        has_deposit_speed       存款加速通道(0不支持,1支持
 * @apiSuccess (data字段说明) {int}        has_own_service         专属客服经理(0没有,1有
 * @apiSuccess (data字段说明) {float}      wash_code               洗码率
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "保存数据成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "level": "",
 *        "name": "",
 *        "valid_bet_min": "",
 *        "valid_bet_max": "",
 *        "upgrade_amount": "",
 *        "weekly_amount": "",
 *        "month_amount": "",
 *        "upgrade_amount_total": "",
 *        "has_deposit_speed": "",
 *        "has_own_service": "",
 *        "wash_code": "",
 * 	  },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *VipLevelsController) View(ctx *Context) {
	view(ctx, &vipLevels)
}
