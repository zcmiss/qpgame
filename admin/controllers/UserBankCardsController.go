package controllers

import (
	"qpgame/admin/models"
	"qpgame/admin/validations"
)

var userBankCards = models.UserBankCards{}                          //模型
var userBankCardsValidation = validations.UserBankCardsValidation{} //校验器

type UserBankCardsController struct{}

/**
 * @api {get} admin/api/auth/v1/user_bank_cards 用户绑卡列表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>银行卡列表</strong><br />
 * 业务描述: 银行卡列表</br>
 * @apiVersion 1.0.0
 * @apiName     indexUserBankCards
 * @apiGroup    user
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int}     page            页数
 * @apiParam (客户端请求参数) {int}    page_size       每页记录数
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
 * @apiSuccess (data-rows每个子对象字段说明) {string}     user_id                 UserId
 * @apiSuccess (data-rows每个子对象字段说明) {string}     card_number             银行卡号
 * @apiSuccess (data-rows每个子对象字段说明) {string}     address                 银行卡地址
 * @apiSuccess (data-rows每个子对象字段说明) {string}     bank_name               银行名称
 * @apiSuccess (data-rows每个子对象字段说明) {string}     name                    姓名
 * @apiSuccess (data-rows每个子对象字段说明) {int}        status                  状态
 * @apiSuccess (data-rows每个子对象字段说明) {string}     remark                  备注
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "数据获取成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "user_id": "",
 *        "card_number": "",
 *        "address": "",
 *        "created": "",
 *        "bank_name": "",
 *        "name": "",
 *        "status": "",
 *        "updated": "",
 *        "remark": ""
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *UserBankCardsController) Index(ctx *Context) {
	index(ctx, &userBankCards)
}

/**
 * @api {post} admin/api/auth/v1/user_bank_cards/add	用户绑卡添加
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>添加/修改银行卡</strong><br />
 * 业务描述: 添加/修改银行卡 <br />
 * <strong><span style="color: red">注意: </span></strong><br />
 * <span style="color:red">修改操作API不再单独列出, 请参考以下</span><br />
 * <span style="color:red">添加: /admin/api/auth/v1/user_bank_cards/add </span> &nbsp;&nbsp; <br />
 * <span style="color:red">修改: /admin/api/auth/v1/user_bank_cards/update </span>
 * @apiVersion 1.0.0
 * @apiName     saveUserBankCards
 * @apiGroup    user
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }

 * @apiParam (客户端请求参数) {int} 	id    					记录编号,仅修改操作时(即接口: /admin/api/auth/v1/user_bank_cards/update)需要<br />
 *															添加操作(即接口: /admin/api/auth/v1/user_bank_cards/add)不需要此参数<br />
 * 															* 如果提供此编号, 则视为修改记录
 * @apiParam (客户端请求参数) {string}   	user_id                 UserId
 * @apiParam (客户端请求参数) {string}   	card_number             银行卡号
 * @apiParam (客户端请求参数) {string}   	address                 银行卡地址
 * @apiParam (客户端请求参数) {int}      	created                 添加时间
 * @apiParam (客户端请求参数) {string}   	bank_name               银行名称
 * @apiParam (客户端请求参数) {string}   	name                    姓名
 * @apiParam (客户端请求参数) {int}      	status                  状态
 * @apiParam (客户端请求参数) {int}      	updated                 修改时间
 * @apiParam (客户端请求参数) {string}   	remark                  备注<br />
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
func (self *UserBankCardsController) Save(ctx *Context) {
	save(ctx, &userBankCards, &userBankCardsValidation)
}

/**
 * @api {get} admin/api/auth/v1/user_bank_cards/view 				用户绑卡详情
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>银行卡详情</strong><br />
 * 业务描述: 银行卡详情</br>
 * @apiVersion 1.0.0
 * @apiName     viewUserBankCards
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
 * @apiSuccess (data字段说明) {string}     user_id                 UserId
 * @apiSuccess (data字段说明) {string}     card_number             银行卡号
 * @apiSuccess (data字段说明) {string}     address                 银行卡地址
 * @apiSuccess (data字段说明) {int}        created                 添加时间
 * @apiSuccess (data字段说明) {string}     bank_name               银行名称
 * @apiSuccess (data字段说明) {string}     name                    姓名
 * @apiSuccess (data字段说明) {int}        status                  状态
 * @apiSuccess (data字段说明) {int}        updated                 修改时间
 * @apiSuccess (data字段说明) {string}     remark                  备注
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "保存数据成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "user_id": "",
 *        "card_number": "",
 *        "address": "",
 *        "created": "",
 *        "bank_name": "",
 *        "name": "",
 *        "status": "",
 *        "updated": "",
 *        "remark": ""
 * 	  },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *UserBankCardsController) View(ctx *Context) {
	view(ctx, &userBankCards)
}

/**
 * @api {get} admin/api/auth/v1/user_bank_cards/delete 用户绑卡删除
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>银行卡删除</strong><br />
 * 业务描述: 银行卡删除</br>
 * @apiVersion 1.0.0
 * @apiName     deleteUserBankCards
 * @apiGroup    user
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }

 * @apiParam (客户端请求参数) {int} 	user_id    		用户编号
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
func (self *UserBankCardsController) Delete(ctx *Context) {
	remove(ctx, &userBankCards)
}
