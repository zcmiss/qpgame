package controllers

import (
	"qpgame/admin/models"
	"qpgame/admin/validations"
)

var chargeCards = models.ChargeCards{}                          //模型
var chargeCardsValidation = validations.ChargeCardsValidation{} //校验器

type ChargeCardsController struct{}

/**
 * @api {get} admin/api/auth/v1/charge_cards/get_user_groups 获取用户分层列表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: lyndon</span><br/><br/>
 * <strong>获取用户分层数据</strong><br />
 * 业务描述: 获取用户分层数据</br>
 * @apiVersion 1.0.0
 * @apiName     getUserGroups
 * @apiGroup    finance
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
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
 * @apiSuccess (data每个子对象字段说明) {int}		  id 					  分层编号
 * @apiSuccess (data每个子对象字段说明) {string}     group_name            分层名称
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *   "clientMsg": "获取数据成功",
 *   "code": 200,
 *   "data": [
 *           {
 *               "group_name": "testpay",
 *               "id": "2",
 *               "is_default": "否",
 *               "remark": "支付测试用户组"
 *           },
 *           {
 *               "group_name": "all",
 *               "id": "1",
 *               "is_default": "是",
 *               "remark": "所有用户组"
 *           }
 *   ],
 *   "internalMsg": "",
 *   "timeConsumed": 4032
 * }
 */
func (self *ChargeCardsController) GetUserGroups(ctx *Context) {
	var userGroups = models.UserGroups{}
	records, err := userGroups.GetRecords(ctx)
	if err != nil {
		responseFailure(ctx, err.Error(), "获取数据失败")
		return
	}
	responseSuccess(ctx, "获取数据成功", records.Rows)
}

/**
 * @api {get} admin/api/auth/v1/charge_cards 公司入款账号列表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>充值卡列表</strong><br />
 * 业务描述: 充值卡列表</br>
 * @apiVersion 1.0.0
 * @apiName     indexChargeCards
 * @apiGroup    finance
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int} 	page			页数
 * @apiParam (客户端请求参数) {int} 	page_size		每页记录数
 * @apiParam (客户端请求参数) {int} 	name 银行名称
 * @apiParam (客户端请求参数) {int} 	owner 持卡人
 * @apiParam (客户端请求参数) {int} 	card_number 卡号
 * @apiParam (客户端请求参数) {int} 	created_start 添加时间/开始
 * @apiParam (客户端请求参数) {int} 	created_end 添加时间/结束
 * @apiParam (客户端请求参数) {int} 	title  标题
 * @apiParam (客户端请求参数) {int} 	state 状态,0:锁定,1:正常
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
 * @apiSuccess (data-rows每个子对象字段说明) {string}     name                    银行名称
 * @apiSuccess (data-rows每个子对象字段说明) {string}     owner                   持卡人
 * @apiSuccess (data-rows每个子对象字段说明) {string}     card_number             卡号
 * @apiSuccess (data-rows每个子对象字段说明) {string}     bank_address            开户地址
 * @apiSuccess (data-rows每个子对象字段说明) {int}        charge_type_id          充值类型编号
 * @apiSuccess (data-rows每个子对象字段说明) {string}        charge_type_name          充值类型名称
 * @apiSuccess (data-rows每个子对象字段说明) {int}        created                 添加时间
 * @apiSuccess (data-rows每个子对象字段说明) {string}     remark                  备注
 * @apiSuccess (data-rows每个子对象字段说明) {int}        state                   状态,1:正常,0:锁定
 * @apiSuccess (data-rows每个子对象字段说明) {string}        state_name           状态说明
 * @apiSuccess (data-rows每个子对象字段说明) {string}     hint                    支付提示
 * @apiSuccess (data-rows每个子对象字段说明) {string}     title                   支付标题
 * @apiSuccess (data-rows每个子对象字段说明) {int}        mfrom                   支付额度
 * @apiSuccess (data-rows每个子对象字段说明) {int}        mto                     支付额度
 * @apiSuccess (data-rows每个子对象字段说明) {int}        amount_limit            停用金额
 * @apiSuccess (data-rows每个子对象字段说明) {string}     addr_type               地址类型
 * @apiSuccess (data-rows每个子对象字段说明) {int}        credential_id           支付方式
 * @apiSuccess (data-rows每个子对象字段说明) {int}        priority                使用优先级
 * @apiSuccess (data-rows每个子对象字段说明) {string}     qr_code                 财务收款二维码,图片
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "数据获取成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "name": "",
 *        "owner": "",
 *        "card_number": "",
 *        "bank_address": "",
 *        "charge_type_id": "",
 *        "charge_type_name": "",
 *        "created": "",
 *        "remark": "",
 *        "state": "",
 *        "state_name": "",
 *        "logo": "",
 *        "hint": "",
 *        "title": "",
 *        "mfrom": "",
 *        "mto": "",
 *        "amount_limit": "",
 *        "addr_type": "",
 *        "credential_id": "",
 *        "priority": "",
 *        "qr_code": ""
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *ChargeCardsController) Index(ctx *Context) {
	(*ctx).Params().Set("type", "1,2,3")
	index(ctx, &chargeCards)
}

/**
 * @api {post} admin/api/auth/v1/charge_cards/add	公司入款账号添加
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>添加/修改</strong><br />
 * 业务描述: 添加/修 <br />
 * <strong><span style="color: red">注意: </span></strong><br />
 * <span style="color:red">修改操作API不再单独列出, 请参考以下</span><br />
 * <span style="color:red">添加: /admin/api/auth/v1/charge_cards/add </span> &nbsp;&nbsp; <br />
 * <span style="color:red">修改: /admin/api/auth/v1/charge_cards/update </span>
 * @apiVersion 1.0.0
 * @apiName     saveChargeCards
 * @apiGroup    finance
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }

 * @apiParam (客户端请求参数) {int} 	id    					记录编号,仅修改操作时(即接口: /admin/api/auth/v1/charge_cards/update)需要<br />
 *															添加操作(即接口: /admin/api/auth/v1/charge_cards/add)不需要此参数<br />
 * 															* 如果提供此编号, 则视为修改记录
 * @apiParam (客户端请求参数) {string}   	name                    银行名称
 * @apiParam (客户端请求参数) {string}   	owner                   持卡人
 * @apiParam (客户端请求参数) {string}   	card_number             卡号
 * @apiParam (客户端请求参数) {string}   	bank_address            开户地址
 * @apiParam (客户端请求参数) {int}      	charge_type_id          充值类型编号,列表, 数据来源: 充值分类方式
 * @apiParam (客户端请求参数) {string}   	remark                  备注
 * @apiParam (客户端请求参数) {int}      	state                   状态, 0:正常, 1:锁定
 * @apiParam (客户端请求参数) {string}   	logo                    LOGO
 * @apiParam (客户端请求参数) {string}   	hint                    支付提示
 * @apiParam (客户端请求参数) {string}   	title                   支付标题
 * @apiParam (客户端请求参数) {int}      	mfrom                   支付额度
 * @apiParam (客户端请求参数) {int}      	mto                     支付额度
 * @apiParam (客户端请求参数) {int}      	amount_limit            停用金额
 * @apiParam (客户端请求参数) {string}   	addr_type               地址类型
 * @apiParam (客户端请求参数) {int}      	credential_id           支付方式, 下拉列表, 数据来源: 第三方支付证书
 * @apiParam (客户端请求参数) {int}      	priority                优先级
 * @apiParam (客户端请求参数) {string}   	qr_code                 财务收款二维码,图片
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
func (self *ChargeCardsController) Save(ctx *Context) {
	save(ctx, &chargeCards, &chargeCardsValidation)
}

/**
 * @api {get} admin/api/auth/v1/charge_cards/view 				公司入款账号详情
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>充值卡详情</strong><br />
 * 业务描述: 充值卡详情</br>
 * @apiVersion 1.0.0
 * @apiName     viewChargeCards
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
 * @apiSuccess (返回结果)  {int}		code            200
 * @apiSuccess (返回结果)  {string} 	clientMsg       提示信息
 * @apiSuccess (返回结果)  {string}		internalMsg     内部信息
 * @apiSuccess (返回结果)  {json}		data            返回数据
 * @apiSuccess (返回结果)  {float}		timeConsumed    后台耗时
 *
 * @apiSuccess (data字段说明) {int}		  id 					  记录编号
 * @apiSuccess (data字段说明) {string}     name                    银行名称
 * @apiSuccess (data字段说明) {string}     owner                   持卡人
 * @apiSuccess (data字段说明) {string}     card_number             卡号
 * @apiSuccess (data字段说明) {string}     bank_address            开户地址
 * @apiSuccess (data字段说明) {int}        charge_type_id          充值类型编号
 * @apiSuccess (data字段说明) {int}        created                 添加时间
 * @apiSuccess (data字段说明) {string}     remark                  备注
 * @apiSuccess (data字段说明) {int}        state                   状态, 0:锁定, 1:正常
 * @apiSuccess (data字段说明) {string}     logo                    LOGO
 * @apiSuccess (data字段说明) {string}     hint                    支付提示
 * @apiSuccess (data字段说明) {string}     title                   支付标题
 * @apiSuccess (data字段说明) {int}        mfrom                   支付额度
 * @apiSuccess (data字段说明) {int}        mto                     支付额度
 * @apiSuccess (data字段说明) {int}        amount_limit            停用金额
 * @apiSuccess (data字段说明) {string}     addr_type               地址类型
 * @apiSuccess (data字段说明) {int}        credential_id           支付方式, 数据来源: 第三方支付证书
 * @apiSuccess (data字段说明) {int}        priority                充值排序
 * @apiSuccess (data字段说明) {string}     qr_code                 财务收款二维码,图片
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "保存数据成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "name": "",
 *        "owner": "",
 *        "card_number": "",
 *        "bank_address": "",
 *        "charge_type_id": "",
 *        "created": "",
 *        "remark": "",
 *        "state": "",
 *        "logo": "",
 *        "hint": "",
 *        "title": "",
 *        "mfrom": "",
 *        "mto": "",
 *        "amount_limit": "",
 *        "addr_type": "",
 *        "credential_id": "",
 *        "priority": "",
 *        "qr_code": ""
 * 	  },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *ChargeCardsController) View(ctx *Context) {
	view(ctx, &chargeCards)
}

/**
 * @api {get} admin/api/auth/v1/charge_cards/delete 公司入款账号删除
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>充值卡删除</strong><br />
 * 业务描述: 充值卡删除</br>
 * @apiVersion 1.0.0
 * @apiName     deleteChargeCards
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
func (self *ChargeCardsController) Delete(ctx *Context) {
	remove(ctx, &chargeCards)
}

/**
 * @api {get} admin/api/auth/v1/charge_cards/onlines 线上入款账号
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>会员线上入款账号</strong><br />
 * 业务描述: 会员线上入款账号</br>
 * @apiVersion 1.0.0
 * @apiName     chargeCardsOnlines
 * @apiGroup    finance
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int} 	page			页数
 * @apiParam (客户端请求参数) {int} 	page_size		每页记录数
 * @apiParam (客户端请求参数) {int} 	name 银行名称
 * @apiParam (客户端请求参数) {int} 	owner 持卡人
 * @apiParam (客户端请求参数) {int} 	card_number 卡号
 * @apiParam (客户端请求参数) {int} 	created_start 添加时间/开始
 * @apiParam (客户端请求参数) {int} 	created_end 添加时间/结束
 * @apiParam (客户端请求参数) {int} 	title  标题
 * @apiParam (客户端请求参数) {int} 	state 状态,1:正常,0:锁定
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
 * @apiSuccess (data-rows每个子对象字段说明) {string}     name                    银行名称
 * @apiSuccess (data-rows每个子对象字段说明) {string}     owner                   持卡人
 * @apiSuccess (data-rows每个子对象字段说明) {string}     card_number             卡号
 * @apiSuccess (data-rows每个子对象字段说明) {string}     bank_address            开户地址
 * @apiSuccess (data-rows每个子对象字段说明) {int}        charge_type_id          充值类型编号
 * @apiSuccess (data-rows每个子对象字段说明) {int}        created                 添加时间
 * @apiSuccess (data-rows每个子对象字段说明) {string}     remark                  备注
 * @apiSuccess (data-rows每个子对象字段说明) {int}        state                   状态,0:锁定,1:正常
 * @apiSuccess (data-rows每个子对象字段说明) {string}        state_name              状态说明, 锁定/正常
 * @apiSuccess (data-rows每个子对象字段说明) {string}     logo                    LOGO
 * @apiSuccess (data-rows每个子对象字段说明) {string}     hint                    支付提示
 * @apiSuccess (data-rows每个子对象字段说明) {string}     title                   支付标题
 * @apiSuccess (data-rows每个子对象字段说明) {int}        mfrom                   支付额度
 * @apiSuccess (data-rows每个子对象字段说明) {int}        mto                     支付额度
 * @apiSuccess (data-rows每个子对象字段说明) {int}        amount_limit            停用金额
 * @apiSuccess (data-rows每个子对象字段说明) {string}     addr_type               地址类型
 * @apiSuccess (data-rows每个子对象字段说明) {int}        credential_id           支付方式
 * @apiSuccess (data-rows每个子对象字段说明) {int}        priority                充值排序
 * @apiSuccess (data-rows每个子对象字段说明) {string}     qr_code                 财务收款二维码,图片
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "数据获取成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "name": "",
 *        "owner": "",
 *        "card_number": "",
 *        "bank_address": "",
 *        "charge_type_id": "",
 *        "created": "",
 *        "remark": "",
 *        "state": "",
 *        "state_name": "",
 *        "logo": "",
 *        "hint": "",
 *        "title": "",
 *        "mfrom": "",
 *        "mto": "",
 *        "amount_limit": "",
 *        "addr_type": "",
 *        "credential_id": "",
 *        "priority": "",
 *        "qr_code": ""
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *ChargeCardsController) Onlines(ctx *Context) {
	(*ctx).Params().Set("type", "4,5,6")
	index(ctx, &chargeCards)
}

/**
 * @api {get} admin/api/auth/v1/charge_cards/view_online 				线上入款账号详情
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>充值卡详情</strong><br />
 * 业务描述: 充值卡详情</br>
 * @apiVersion 1.0.0
 * @apiName     viewChargeCardsOnline
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
 * @apiSuccess (返回结果)  {int}		code            200
 * @apiSuccess (返回结果)  {string} 	clientMsg       提示信息
 * @apiSuccess (返回结果)  {string}		internalMsg     内部信息
 * @apiSuccess (返回结果)  {json}		data            返回数据
 * @apiSuccess (返回结果)  {float}		timeConsumed    后台耗时
 *
 * @apiSuccess (data字段说明) {int}		  id 					  记录编号
 * @apiSuccess (data字段说明) {string}     name                    银行名称
 * @apiSuccess (data字段说明) {string}     owner                   持卡人
 * @apiSuccess (data字段说明) {string}     card_number             卡号
 * @apiSuccess (data字段说明) {string}     bank_address            开户地址
 * @apiSuccess (data字段说明) {int}        charge_type_id          充值类型编号, 数据来源,充值分类方式
 * @apiSuccess (data字段说明) {int}        created                 添加时间
 * @apiSuccess (data字段说明) {string}     remark                  备注
 * @apiSuccess (data字段说明) {int}        state                   状态
 * @apiSuccess (data字段说明) {string}     logo                    LOGO
 * @apiSuccess (data字段说明) {string}     hint                    支付提示
 * @apiSuccess (data字段说明) {string}     title                   支付标题
 * @apiSuccess (data字段说明) {int}        mfrom                   支付额度
 * @apiSuccess (data字段说明) {int}        mto                     支付额度
 * @apiSuccess (data字段说明) {int}        amount_limit            停用金额
 * @apiSuccess (data字段说明) {string}     addr_type               AddrType
 * @apiSuccess (data字段说明) {int}        credential_id           支付方式, 数据来源:第三方支付证书
 * @apiSuccess (data字段说明) {int}        priority                充值排序
 * @apiSuccess (data字段说明) {string}     qr_code                 财务收款二维码,图片
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "保存数据成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "name": "",
 *        "owner": "",
 *        "card_number": "",
 *        "bank_address": "",
 *        "charge_type_id": "",
 *        "created": "",
 *        "remark": "",
 *        "state": "",
 *        "logo": "",
 *        "hint": "",
 *        "title": "",
 *        "mfrom": "",
 *        "mto": "",
 *        "amount_limit": "",
 *        "addr_type": "",
 *        "credential_id": "",
 *        "priority": "",
 *        "qr_code": ""
 * 	  },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *ChargeCardsController) OnlineView(ctx *Context) {
	view(ctx, &chargeCards)
}

/**
 * @api {get} admin/api/auth/v1/charge_cards/delete_online 线上入款账号删除
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>充值卡删除</strong><br />
 * 业务描述: 充值卡删除</br>
 * @apiVersion 1.0.0
 * @apiName     deleteChargeCardsOnline
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
func (self *ChargeCardsController) OnlineDelete(ctx *Context) {
	remove(ctx, &chargeCards)
}
