package controllers

import (
	"qpgame/admin/models"
	"qpgame/admin/validations"
)

var chargeTypes = models.ChargeTypes{}                          //模型
var chargeTypesValidation = validations.ChargeTypesValidation{} //校验器

type ChargeTypesController struct{}

/**
 * @api {get} admin/api/auth/v1/charge_types 充值类型列表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>充值类型列表</strong><br />
 * 业务描述: 充值类型列表</br>
 * @apiVersion 1.0.0
 * @apiName     indexChargeTypes
 * @apiGroup    finance
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int} 	page			页数
 * @apiParam (客户端请求参数) {int} 	page_size		每页记录数
 * @apiParam (客户端请求参数) {int} 	name 类型名称
 * @apiParam (客户端请求参数) {int} 	state 状态
 * @apiParam (客户端请求参数) {int} 	created_start 添加时间/开始
 * @apiParam (客户端请求参数) {int} 	created_end 添加时间/结束
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
 * @apiSuccess (data-rows每个子对象字段说明) {int}		  id 					编号
 * @apiSuccess (data-rows每个子对象字段说明) {string}     name                    充值名称
 * @apiSuccess (data-rows每个子对象字段说明) {string}     remark                  充值说明
 * @apiSuccess (data-rows每个子对象字段说明) {int}        state                   状态(0禁用,1启用
 * @apiSuccess (data-rows每个子对象字段说明) {string}     charge_numbers          充值金额选项例子:(50,100,300,800,1000,2000,3000,5000
 * @apiSuccess (data-rows每个子对象字段说明) {string}     logo                    保存在s3上的地址
 * @apiSuccess (data-rows每个子对象字段说明) {string}     updated                 最后更新时间
 * @apiSuccess (data-rows每个子对象字段说明) {int}        priority                类型排序
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "数据获取成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "name": "",
 *        "remark": "",
 *        "state": "",
 *        "charge_numbers": "",
 *        "created": "",
 *        "logo": "",
 *        "updated": "",
 *        "priority": ""
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *ChargeTypesController) Index(ctx *Context) {
	index(ctx, &chargeTypes)
}

/**
 * @api {post} admin/api/auth/v1/charge_types/add	充值类型添加
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>添加/修改充值类型</strong><br />
 * 业务描述: 添加/修改充值类型 <br />
 * <strong><span style="color: red">注意: </span></strong><br />
 * <span style="color:red">修改操作API不再单独列出, 请参考以下</span><br />
 * <span style="color:red">添加: /admin/api/auth/v1/charge_types/add </span> &nbsp;&nbsp; <br />
 * <span style="color:red">修改: /admin/api/auth/v1/charge_types/update </span>
 * @apiVersion 1.0.0
 * @apiName     saveChargeTypes
 * @apiGroup    finance
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }

 * @apiParam (客户端请求参数) {int} 	id    					记录编号,仅修改操作时(即接口: /admin/api/auth/v1/charge_types/update)需要<br />
 *															添加操作(即接口: /admin/api/auth/v1/charge_types/add)不需要此参数<br />
 * 															* 如果提供此编号, 则视为修改记录
 * @apiParam (客户端请求参数) {string}   	name                    充值名称
 * @apiParam (客户端请求参数) {string}   	remark                  充值说明
 * @apiParam (客户端请求参数) {int}      	state                   状态(0禁用,1启用
 * @apiParam (客户端请求参数) {string}   	charge_numbers          充值金额选项例子:[50,100,300,800,1000,2000,3000,5000]
 * @apiParam (客户端请求参数) {string}   	logo                    保存在s3上的地址
 * @apiParam (客户端请求参数) {int}      	priority                类型排序<br />
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
func (self *ChargeTypesController) Save(ctx *Context) {
	save(ctx, &chargeTypes, &chargeTypesValidation)
}

/**
 * @api {get} admin/api/auth/v1/charge_types/view 				充值类型详情
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>充值类型详情</strong><br />
 * 业务描述: 充值类型详情</br>
 * @apiVersion 1.0.0
 * @apiName     viewChargeTypes
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
 * @apiSuccess (data字段说明) {int}		  id 					  记录编号
 * @apiSuccess (data字段说明) {string}     name                    充值名称
 * @apiSuccess (data字段说明) {string}     remark                  充值说明
 * @apiSuccess (data字段说明) {int}        state                   状态(0禁用,1启用
 * @apiSuccess (data字段说明) {string}     charge_numbers          充值金额选项例子:(50,100,300,800,1000,2000,3000,5000
 * @apiSuccess (data字段说明) {int}        created                 创建时间
 * @apiSuccess (data字段说明) {string}     logo                    保存在s3上的地址
 * @apiSuccess (data字段说明) {string}     updated                 最后更新时间
 * @apiSuccess (data字段说明) {int}        priority                类型排序
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "保存数据成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "name": "",
 *        "remark": "",
 *        "state": "",
 *        "charge_numbers": "",
 *        "created": "",
 *        "logo": "",
 *        "updated": "",
 *        "priority": ""
 * 	  },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *ChargeTypesController) View(ctx *Context) {
	view(ctx, &chargeTypes)
}

/**
 * @api {get} admin/api/auth/v1/charge_types/delete 充值类型删除
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>充值类型删除</strong><br />
 * 业务描述: 充值类型删除</br>
 * @apiVersion 1.0.0
 * @apiName     deleteChargeTypes
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
func (self *ChargeTypesController) Delete(ctx *Context) {
	remove(ctx, &chargeTypes)
}
