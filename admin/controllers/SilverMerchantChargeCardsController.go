package controllers

import (
	"qpgame/admin/models"
	"qpgame/admin/validations"
)

var silverMerchantChargeCards = models.SilverMerchantChargeCards{}
var silverMerchantChargeCardsValidation = validations.SilverMerchantChargeCardsValidation{}

type SilverMerchantChargeCardsController struct{}

/**
 * @api {get} admin/api/auth/v1/silver_merchant_charge_cards 银商支付方式列表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: lyndon</span><br/><br/>
 * <strong>银商支付方式记录</strong><br />
 * 业务描述: 银商支付方式记录</br>
 * @apiVersion 1.0.0
 * @apiName     indexSilverMerchantChargeCards
 * @apiGroup    silver_merchant
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int}    page           页数
 * @apiParam (客户端请求参数) {int}    page_size      每页记录数
 * @apiParam (客户端请求参数) {string} created_start  添加时间/开始
 * @apiParam (客户端请求参数) {string} created_end    添加时间/结束
 *
 * @apiError (请求失败返回) {int}      code           错误代码
 * @apiError (请求失败返回) {string}   clientMsg      提示信息
 * @apiError (请求失败返回) {string}   internalMsg    内部错误信息
 * @apiError (请求失败返回) {float}    timeConsumed   后台耗时
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
 * @apiSuccess (返回结果)  {json}	data            返回数据
 * @apiSuccess (返回结果)  {float}  	timeConsumed    后台耗时
 *
 * @apiSuccess (data字段说明) {array}  	rows        数据列表
 * @apiSuccess (data字段说明) {int}    	page		当前页数
 * @apiSuccess (data字段说明) {int}    	page_count	总的页数
 * @apiSuccess (data字段说明) {int}    	total_rows	总记录数
 * @apiSuccess (data字段说明) {int}    	page_size	每页记录数
 *
 * @apiSuccess (data-rows每个子对象字段说明) {int}		 id 				记录编号
 * @apiSuccess (data-rows每个子对象字段说明) {string}	 owner              持卡人
 * @apiSuccess (data-rows每个子对象字段说明) {string}	 card_number        卡号
 * @apiSuccess (data-rows每个子对象字段说明) {string}   bank_address      开户地址
 * @apiSuccess (data-rows每个子对象字段说明) {string}   remark            备注
 * @apiSuccess (data-rows每个子对象字段说明) {string}   logo              LOGO
 * @apiSuccess (data-rows每个子对象字段说明) {int}     mfrom              支付最小额度
 * @apiSuccess (data-rows每个子对象字段说明) {int}		 mto 		        支付最大额度
 * @apiSuccess (data-rows每个子对象字段说明) {int}	 priority 		       充值排序
 * @apiSuccess (data-rows每个子对象字段说明) {int}     state               状态(0停用,1可用)
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "获取数据成功",
 *    "code": 200,
 *    "data": {
 *        "rows": [
 *            {
 *                "id": 1,
 *                "owner": "sdf",
 *                "card_number": "23423423423",
 *                "bank_address": "sdfsdf",
 *                "remark": "sdfsdsfsd",
 *                "logo": "http://sdfsdsdfs",
 *                "mfrom": "2",
 *                "mto": "1434",
 *                "priority": "2",
 *                "state": "1"
 *            }
 *        ],
 *        "page": 1,
 *        "page_count": 1,
 *        "total_rows": 1,
 *        "page_size": 20
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 0
 *}
 */
func (self *SilverMerchantChargeCardsController) Index(ctx *Context) {
	index(ctx, &silverMerchantChargeCards)
}

/**
 * @api {get} admin/api/auth/v1/silver_merchant_charge_cards/view 			银商支付方式详情
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: lyndon</span><br/><br/>
 * <strong>银商支付方式详情</strong><br />
 * 业务描述: 银商支付方式详情</br>
 * @apiVersion 1.0.0
 * @apiName     viewSilverMerchantChargeCards
 * @apiGroup    silver_merchant
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
 * @apiSuccess (data每个字段说明) {int}		 id 				记录编号
 * @apiSuccess (data每个字段说明) {string}	 owner              持卡人
 * @apiSuccess (data每个字段说明) {string}	 card_number        卡号
 * @apiSuccess (data每个字段说明) {string}   bank_address      开户地址
 * @apiSuccess (data每个字段说明) {string}   remark            备注
 * @apiSuccess (data每个字段说明) {string}   logo              LOGO
 * @apiSuccess (data每个字段说明) {int}     mfrom              支付最小额度
 * @apiSuccess (data每个字段说明) {int}		 mto 		        支付最大额度
 * @apiSuccess (data每个字段说明) {int}	 priority 		       充值排序
 * @apiSuccess (data每个字段说明) {int}     state               状态(0停用,1可用)
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "保存数据成功",
 *    "code": 200,
 *    "data": {
 *                "id": 1,
 *                "owner": "sdf",
 *                "card_number": "23423423423",
 *                "bank_address": "sdfsdf",
 *                "remark": "sdfsdsfsd",
 *                "logo": "http://sdfsdsdfs",
 *                "mfrom": "2",
 *                "mto": "1434",
 *                "priority": "2",
 *                "state": "1"
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *SilverMerchantChargeCardsController) View(ctx *Context) {
	view(ctx, &silverMerchantChargeCards)
}

/**
 * @api {post} admin/api/auth/v1/silver_merchant_charge_cards/add 银商支付方式添加
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: lyndon</span><br/><br/>
 * <strong>添加/修改银商支付方式</strong><br />
 * 业务描述: 添加/修改银商支付方式 <br />
 * <strong><span style="color: red">注意: </span></strong><br />
 * <span style="color:red">修改操作API不再单独列出, 请参考以下</span><br />
 * <span style="color:red">添加: /admin/api/auth/v1/silver_merchant_charge_cards/add </span> &nbsp;&nbsp; <br />
 * <span style="color:red">修改: /admin/api/auth/v1/silver_merchant_charge_cards/update </span>
 * @apiVersion 1.0.0
 * @apiName     saveSilverMerchantChargeCards
 * @apiGroup    silver_merchant
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }

 * @apiParam (客户端请求参数) {int} 	id   记录编号,仅修改操作时(即接口: /admin/api/auth/v1/silver_merchant_charge_cards/update)需要<br />
 *										 添加操作(即接口: /admin/api/auth/v1/silver_merchant_charge_cards/add)不需要此参数<br />
 * 										 * 如果提供此编号, 则视为修改记录
 * @apiParam (客户端请求参数) {string}	 name              银行名称
 * @apiParam (客户端请求参数) {string}	 owner              持卡人
 * @apiParam (客户端请求参数) {string}	 card_number        卡号
 * @apiParam (客户端请求参数) {string}   bank_address      开户地址
 * @apiParam (客户端请求参数) {string}   remark            备注
 * @apiParam (客户端请求参数) {string}   logo              LOGO
 * @apiParam (客户端请求参数) {int}     mfrom              支付最小额度
 * @apiParam (客户端请求参数) {int}		 mto 		        支付最大额度
 * @apiParam (客户端请求参数) {int}	 priority 		       充值排序
 * @apiParam (客户端请求参数) {int}     state               状态(0停用,1可用)
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
 * @apiSuccess (返回结果)  {string}  	internalMsg     内部错误信息
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
func (self *SilverMerchantChargeCardsController) Save(ctx *Context) {
	save(ctx, &silverMerchantChargeCards, &silverMerchantChargeCardsValidation)
}

/**
 * @api {get} admin/api/auth/v1/silver_merchant_charge_cards/delete 银商支付方式删除
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: lyndon</span><br/><br/>
 * <strong>银商支付方式删除</strong><br />
 * 业务描述: 银商支付方式删除</br>
 * @apiVersion 1.0.0
 * @apiName     deleteSilverMerchantChargeCards
 * @apiGroup    silver_merchant
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }

 * @apiParam (客户端请求参数) {int} 	id    			 记录编号
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

 * @apiSuccess (返回结果)  {int} 		code            200
 * @apiSuccess (返回结果)  {string} 	clientMsg           提示信息
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
func (self *SilverMerchantChargeCardsController) Delete(ctx *Context) {
	remove(ctx, &silverMerchantChargeCards)
}
