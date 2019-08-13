package controllers

import (
	"qpgame/admin/models"
	"qpgame/admin/validations"
)

var redpacketSystems = models.RedpacketSystems{}                          //模型
var redpacketSystemsValidation = validations.RedpacketSystemsValidation{} //校验器

type RedpacketSystemsController struct{}

/**
 * @api {get} admin/api/auth/v1/redpacket_systems 系统红包列表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>系统红包列表</strong><br />
 * 业务描述: 系统红包列表</br>
 * @apiVersion 1.0.0
 * @apiName     indexRedpacketSystems
 * @apiGroup    finance
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int}    page            页数
 * @apiParam (客户端请求参数) {int}   page_size       每页记录数
 * @apiParam (客户端请求参数) {int}   type 	红包分类
 * @apiParam (客户端请求参数) {string}   created_start 	添加时间/开始
 * @apiParam (客户端请求参数) {string}   created_end		添加时间/结束
 * @apiParam (客户端请求参数) {string}   start_time	有效时间/开始
 * @apiParam (客户端请求参数) {string}   end_time	有效时间/结束
 * @apiParam (客户端请求参数) {int}   status 状态
 * @apiParam (客户端请求参数) {int}   calculate_type 发放方式
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
 * @apiSuccess (data-rows每个子对象字段说明) {int}        type                    种类(1节日红包,2.每日幸运红包)
 * @apiSuccess (data-rows每个子对象字段说明) {string}     type_text               种类说明（节日红包/每日幸运红包）
 * @apiSuccess (data-rows每个子对象字段说明) {float}      money                   总金额
 * @apiSuccess (data-rows每个子对象字段说明) {float}      sent_money              已派送金额
 * @apiSuccess (data-rows每个子对象字段说明) {int}        sent_count              红包派送数量
 * @apiSuccess (data-rows每个子对象字段说明) {int}        total                   红包总数量，如果是0就是没有红包个数限制,派完为止
 * @apiSuccess (data-rows每个子对象字段说明) {int}        created                 创建时间
 * @apiSuccess (data-rows每个子对象字段说明) {int}        status                  1正常，2已完结,3.关闭
 * @apiSuccess (data-rows每个子对象字段说明) {int}        end_time                有效期
 * @apiSuccess (data-rows每个子对象字段说明) {int}        start_time              开始时间
 * @apiSuccess (data-rows每个子对象字段说明) {string}     message                 红包消息
 * @apiSuccess (data-rows每个子对象字段说明) {int}        calculate_type          红包发放方式(1.随机红包,2.固定金额红包
 * @apiSuccess (data-rows每个子对象字段说明) {string}     calculate_type_text     红包发放方式说明(随机红包/固定金额红包)
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "数据获取成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "type": "",
 *        "type_text": "",
 *        "money": "",
 *        "sent_money": "",
 *        "sent_count": "",
 *        "total": "",
 *        "created": "",
 *        "status": "",
 *        "end_time": "",
 *        "start_time": "",
 *        "message": "",
 *        "calculate_type": "",
 *        "calculate_type_text": ""
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *RedpacketSystemsController) Index(ctx *Context) {
	index(ctx, &redpacketSystems)
}

/**
 * @api {post} admin/api/v1/redpacket_systems/add	系统红包添加
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>添加/修改系统红包</strong><br />
 * 业务描述: 添加/修改系统红包 <br />
 * <strong><span style="color: red">注意: </span></strong><br />
 * <span style="color:red">修改操作API不再单独列出, 请参考以下</span><br />
 * <span style="color:red">添加: /admin/api/v1/redpacket_systems/add </span> &nbsp;&nbsp; <br />
 * <span style="color:red">修改: /admin/api/v1/redpacket_systems/update </span>
 * @apiVersion 1.0.0
 * @apiName     saveRedpacketSystems
 * @apiGroup    finance
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 *
 * @apiParam (客户端请求参数) {int} 	id    			记录编号,仅修改操作时(即接口: /admin/api/v1/redpacket_systems/update)需要<br />
 *														添加操作(即接口: /admin/api/v1/redpacket_systems/add)不需要此参数<br />
 * 														* 如果提供此编号, 则视为修改记录
 * @apiParam (客户端请求参数) {int}      	type                    种类(1节日红包,2.每日幸运红包
 * @apiParam (客户端请求参数) {float}    	money                   总金额
 * @apiParam (客户端请求参数) {int}      	total                   红包总数量，如果是0就是没有红包个数限制,派完为止
 * @apiParam (客户端请求参数) {int}      	created                 创建时间
 * @apiParam (客户端请求参数) {int}      	status                  1正常，2已完结,3.关闭
 * @apiParam (客户端请求参数) {int}      	end_time                有效期
 * @apiParam (客户端请求参数) {int}      	start_time              开始时间
 * @apiParam (客户端请求参数) {string}   	message                 红包消息
 * @apiParam (客户端请求参数) {int}      	calculate_type          红包发放方式(1.随机红包,2.固定金额红包<br />
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
 * @apiSuccess (返回结果)  {json}  	data            返回数据
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
func (self *RedpacketSystemsController) Save(ctx *Context) {
	save(ctx, &redpacketSystems, &redpacketSystemsValidation)
}

/**
 * @api {get} admin/api/v1/redpacket_systems/view 				系统红包详情
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>系统红包详情</strong><br />
 * 业务描述: 系统红包详情</br>
 * @apiVersion 1.0.0
 * @apiName     viewRedpacketSystems
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
 * @apiSuccess (data字段说明) {int}        type                    种类(1节日红包,2.每日幸运红包
 * @apiSuccess (data字段说明) {float}      money                   总金额
 * @apiSuccess (data字段说明) {float}      sent_money              已派送金额
 * @apiSuccess (data字段说明) {int}        sent_count              红包派送数量
 * @apiSuccess (data字段说明) {int}        total                   红包总数量，如果是0就是没有红包个数限制,派完为止
 * @apiSuccess (data字段说明) {int}        created                 创建时间
 * @apiSuccess (data字段说明) {int}        status                  1正常，2已完结,3.关闭
 * @apiSuccess (data字段说明) {int}        end_time                有效期
 * @apiSuccess (data字段说明) {int}        start_time              开始时间
 * @apiSuccess (data字段说明) {string}     message                 红包消息
 * @apiSuccess (data字段说明) {int}        calculate_type          红包发放方式(1.随机红包,2.固定金额红包
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "保存数据成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "type": "",
 *        "money": "",
 *        "sent_money": "",
 *        "sent_count": "",
 *        "total": "",
 *        "created": "",
 *        "status": "",
 *        "end_time": "",
 *        "start_time": "",
 *        "message": "",
 *        "calculate_type": ""
 * 	  },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *RedpacketSystemsController) View(ctx *Context) {
	view(ctx, &redpacketSystems)
}

/**
 * @api {get} admin/api/v1/redpacket_systems/delete 系统红包删除
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>系统红包删除</strong><br />
 * 业务描述: 系统红包删除</br>
 * @apiVersion 1.0.0
 * @apiName     deleteRedpacketSystems
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
func (self *RedpacketSystemsController) Delete(ctx *Context) {
	remove(ctx, &redpacketSystems)
}
