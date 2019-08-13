package controllers

import (
	"qpgame/admin/models"
	"qpgame/admin/validations"
)

var manualWithdraws = models.ManualWithdraws{}                          //模型
var manualWithdrawsValidation = validations.ManualWithdrawsValidation{} //校验器

type ManualWithdrawsController struct{}

/**
 * @api {get} admin/api/auth/v1/manual_withdraws 人工出款列表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>人工出款列表</strong><br />
 * 业务描述: 人工出款列表</br>
 * @apiVersion 1.0.0
 * @apiName     indexManualWithdraws
 * @apiGroup    finance
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int} 	page			页数
 * @apiParam (客户端请求参数) {int} 	page_size		每页记录数
 * @apiParam (客户端请求参数) {int} 	user_id 用户编号
 * @apiParam (客户端请求参数) {int} 	order 订单编号
 * @apiParam (客户端请求参数) {string} 	deal_start 交易时间/开始
 * @apiParam (客户端请求参数) {string} 	deal_end 交易时间/结束
 * @apiParam (客户端请求参数) {int} 	state 审核状态
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
 * @apiSuccess (data-rows每个子对象字段说明) {int}        user_id					用户编号
 * @apiSuccess (data-rows每个子对象字段说明) {string}        user_name					用户名称
 * @apiSuccess (data-rows每个子对象字段说明) {string}     order                   订单编号
 * @apiSuccess (data-rows每个子对象字段说明) {float}      amount                  取款金额
 * @apiSuccess (data-rows每个子对象字段说明) {float}      quantity                打码量,例如:人工入款时,录入了打码量;人工出款时可以抵消人工入款的打码量
 * @apiSuccess (data-rows每个子对象字段说明) {int}     item                    出款项目
 * @apiSuccess (data-rows每个子对象字段说明) {string}     item_name                    出款项目描述
 * @apiSuccess (data-rows每个子对象字段说明) {string}     comment                 备注
 * @apiSuccess (data-rows每个子对象字段说明) {int}        deal_time               交易日期
 * @apiSuccess (data-rows每个子对象字段说明) {string}     operator                操作人
 * @apiSuccess (data-rows每个子对象字段说明) {int}        state                   审核状态，0为待审核，1为审核通过，2为审核失败
 * @apiSuccess (data-rows每个子对象字段说明) {string}        state_name            审核状态的文字描述
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "数据获取成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "user_id": "",
 *        "user_name": "",
 *        "order": "",
 *        "amount": "",
 *        "quantity": "",
 *        "item": "",
 *        "item_name": "",
 *        "comment": "",
 *        "deal_time": "",
 *        "operator": "",
 *        "state": ""
 *        "state_name": ""
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *ManualWithdrawsController) Index(ctx *Context) {
	index(ctx, &manualWithdraws)
}

/**
 * @api {post} admin/api/auth/v1/manual_withdraws/add	人工出款添加
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>添加/修改人工出款</strong><br />
 * 业务描述: 添加/修改人工出款 <br />
 * <strong><span style="color: red">注意: </span></strong><br />
 * <span style="color:red">修改操作API不再单独列出, 请参考以下</span><br />
 * <span style="color:red">添加: /admin/api/auth/v1/manual_withdraws/add </span> &nbsp;&nbsp; <br />
 * <span style="color:red">修改: /admin/api/auth/v1/manual_withdraws/update </span>
 * @apiVersion 1.0.0
 * @apiName     saveManualWithdraws
 * @apiGroup    finance
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 *
 * @apiParam (客户端请求参数) {int} 	id    			记录编号,仅修改操作时(即接口: /admin/api/auth/v1/manual_withdraws/update)需要<br />
 *														添加操作(即接口: /admin/api/auth/v1/manual_withdraws/add)不需要此参数<br />
 * 														* 如果提供此编号, 则视为修改记录
 * @apiParam (客户端请求参数) {int}      	user_id                 关联用户表id
 * @apiParam (客户端请求参数) {string}   	order                   订单编号
 * @apiParam (客户端请求参数) {float}    	amount                  取款金额
 * @apiParam (客户端请求参数) {float}    	quantity                打码量,例如:人工入款时,录入了打码量;人工出款时可以抵消人工入款的打码量
 * @apiParam (客户端请求参数) {string}   	item                    存款项目
 * @apiParam (客户端请求参数) {string}   	comment                 备注
 * @apiParam (客户端请求参数) {int}      	deal_time               交易日期
 * @apiParam (客户端请求参数) {string}   	operator                操作人
 * @apiParam (客户端请求参数) {int}      	state                   审核状态，0为待审核，1为审核通过，2为审核失败<br />
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
func (self *ManualWithdrawsController) Save(ctx *Context) {
	save(ctx, &manualWithdraws, &manualWithdrawsValidation)
}

/**
 * @api {get} admin/api/auth/v1/manual_withdraws/view 				人工出款详情
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>人工出款详情</strong><br />
 * 业务描述: 人工出款详情</br>
 * @apiVersion 1.0.0
 * @apiName     viewManualWithdraws
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
 * @apiSuccess (返回结果)  {json}		data            返回数据
 * @apiSuccess (返回结果)  {float}   	timeConsumed    后台耗时
 *
 * @apiSuccess (data字段说明) {int}		  id 					  记录编号
 * @apiSuccess (data字段说明) {int}        user_id                 关联用户表id
 * @apiSuccess (data字段说明) {string}     order                   订单编号
 * @apiSuccess (data字段说明) {float}      amount                  取款金额
 * @apiSuccess (data字段说明) {float}      quantity                打码量,例如:人工入款时,录入了打码量;人工出款时可以抵消人工入款的打码量
 * @apiSuccess (data字段说明) {string}     item                    存款项目
 * @apiSuccess (data字段说明) {string}     comment                 备注
 * @apiSuccess (data字段说明) {int}        deal_time               交易日期
 * @apiSuccess (data字段说明) {string}     operator                操作人
 * @apiSuccess (data字段说明) {int}        state                   审核状态，0为待审核，1为审核通过，2为审核失败
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "保存数据成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "user_id": "",
 *        "order": "",
 *        "amount": "",
 *        "quantity": "",
 *        "item": "",
 *        "comment": "",
 *        "deal_time": "",
 *        "operator": "",
 *        "state": ""
 * 	  },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *ManualWithdrawsController) View(ctx *Context) {
	view(ctx, &manualWithdraws)
}

/**
 * @api {get} admin/api/auth/v1/manual_withdraws/delete 人工出款删除
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>人工出款删除</strong><br />
 * 业务描述: 人工出款删除</br>
 * @apiVersion 1.0.0
 * @apiName     deleteManualWithdraws
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
func (self *ManualWithdrawsController) Delete(ctx *Context) {
	remove(ctx, &manualWithdraws)
}

/**
 * @api {get} admin/api/auth/v1/manual_withdraws/allow	人工出款审核通过
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>人工出款审核-通过</strong><br />
 * 业务描述: 人工出款审核-通过</br>
 * @apiVersion 1.0.0
 * @apiName     manualWithdrawsAllow
 * @apiGroup    finance
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
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "保存数据成功",
 *    "code": 200,
 *    "data": {},
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *ManualWithdrawsController) Allow(ctx *Context) {
	responseResult(ctx, manualWithdraws.Allow(ctx), "人工出款审核处理成功")
}

/**
 * @api {get} admin/api/auth/v1/manual_withdraws/deny	人工出款审核拒绝
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong> 人工出款审核-拒绝</strong><br />
 * 业务描述: 人工出款审核-拒绝</br>
 * @apiVersion 1.0.0
 * @apiName     manualWithdrawsDeny
 * @apiGroup    finance
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
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "保存数据成功",
 *    "code": 200,
 *    "data": {},
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *ManualWithdrawsController) Deny(ctx *Context) {
	responseResult(ctx, manualWithdraws.Deny(ctx), "已拒绝人工出款申请")
}
