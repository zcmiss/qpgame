package controllers

import (
	"qpgame/admin/models"
	"qpgame/admin/validations"
)

var silverMerchantChargeRecords = models.SilverMerchantChargeRecords{}
var silverMerchantChargeRecordsValidation = validations.SilverMerchantChargeRecordsValidation{}

type SilverMerchantChargeRecordsController struct{}

/**
 * @api {get} admin/api/auth/v1/silver_merchant_charge_records 银商充值记录
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: lyndon</span><br/><br/>
 * <strong>银商充值记录记录</strong><br />
 * 业务描述: 银商充值记录记录</br>
 * @apiVersion 1.0.0
 * @apiName     indexSilverMerchantChargeRecords
 * @apiGroup    silver_merchant
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int}    page        页数
 * @apiParam (客户端请求参数) {int}    page_size   每页记录数
 * @apiParam (客户端请求参数) {int} 	 merchant_id 银商编号
 * @apiParam (客户端请求参数) {int} 	 state       充值状态，0 待审核，1 成功，2 失败
 * @apiParam (客户端请求参数) {string} 	order_id 订单号
 * @apiParam (客户端请求参数) {string} 	created_start 添加时间/开始
 * @apiParam (客户端请求参数) {string} 	created_end   添加时间/结束
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
 * @apiSuccess (返回结果)  {json}	data            返回数据
 * @apiSuccess (返回结果)  {float}  	timeConsumed    后台耗时
 *
 * @apiSuccess (data字段说明) {array}  	rows        数据列表
 * @apiSuccess (data字段说明) {int}    	page		当前页数
 * @apiSuccess (data字段说明) {int}    	page_count	总的页数
 * @apiSuccess (data字段说明) {int}    	total_rows	总记录数
 * @apiSuccess (data字段说明) {int}    	page_size	每页记录数
 *
 * @apiSuccess (data-rows每个子对象字段说明) {int}		  id            记录编号
 * @apiSuccess (data-rows每个子对象字段说明) {float}    amount         充值授权额度金额
 * @apiSuccess (data-rows每个子对象字段说明) {string}   bank_address   充值开户银行地址
 * @apiSuccess (data-rows每个子对象字段说明) {string}   bank_charge_time 银行转账时间
 * @apiSuccess (data-rows每个子对象字段说明) {string}   card_number    充值银行卡号
 * @apiSuccess (data-rows每个子对象字段说明) {string}   created        添加时间
 * @apiSuccess (data-rows每个子对象字段说明) {string}   ip             充值IP
 * @apiSuccess (data-rows每个子对象字段说明) {int}		  merchant_id 	 银商编号
 * @apiSuccess (data-rows每个子对象字段说明) {string}	  merchant_name  银商名称
 * @apiSuccess (data-rows每个子对象字段说明) {string}   operator       操作者
 * @apiSuccess (data-rows每个子对象字段说明) {string}   order_id       充值订单
 * @apiSuccess (data-rows每个子对象字段说明) {string}   real_name      真实姓名
 * @apiSuccess (data-rows每个子对象字段说明) {string}   remark         备注
 * @apiSuccess (data-rows每个子对象字段说明) {int}      state          状态：0待审核，1成功，2失败
 * @apiSuccess (data-rows每个子对象字段说明) {string}   updated          修改时间
 * @apiSuccess (data-rows每个子对象字段说明) {string}   updated_last     最后更新时间
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "获取数据成功",
 *    "code": 200,
 *    "data": {
 *        "rows": [
 *            {
 *                "amount": "23.000",
 *                "bank_address": "中国上海",
 *                "bank_charge_time": "2017-05-25 17:01:13",
 *                "card_number": "65489556848489",
 *                "created": "2017-05-25 17:01:13",
 *                "id": "1",
 *                "ip": "127.0.0.1",
 *                "merchant_id": "1",
 *                "merchant_name": "XXX",
 *                "operator": "yndon",
 *                "order_id": "SDSSDSDF56DS",
 *                "real_name": "胡汉山",
 *                "remark": "测试一下",
 *                "state": "待审核",
 *                "bank_charge_time": "2017-05-25 17:01:13",
 *                "updated": "2019-06-13 10:36:12",
 *                "updated_last": "2019-06-13 10:36:12"
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
func (self *SilverMerchantChargeRecordsController) Index(ctx *Context) {
	index(ctx, &silverMerchantChargeRecords)
}

/**
 * @api {get} admin/api/auth/v1/silver_merchant_charge_records/view 			银商充值记录详情
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: lyndon</span><br/><br/>
 * <strong>银商充值记录详情</strong><br />
 * 业务描述: 银商充值记录详情</br>
 * @apiVersion 1.0.0
 * @apiName     viewSilverMerchantChargeRecords
 * @apiGroup    silver_merchant
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }

 * @apiParam (客户端请求参数) {int} 	id               记录编号
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
 * @apiSuccess (data每个字段说明) {int}      id             记录编号
 * @apiSuccess (data每个字段说明) {float}    amount         充值授权额度金额
 * @apiSuccess (data每个字段说明) {float}    presented_money 充值赠送金额
 * @apiSuccess (data每个字段说明) {string}   bank_address   充值开户银行地址
 * @apiSuccess (data每个字段说明) {string}   bank_charge_time 银行转账时间
 * @apiSuccess (data每个字段说明) {string}   card_number    充值银行卡号
 * @apiSuccess (data每个字段说明) {string}   created        添加时间
 * @apiSuccess (data每个字段说明) {string}   ip             充值IP
 * @apiSuccess (data每个字段说明) {int}	    merchant_id    银商编号
 * @apiSuccess (data每个字段说明) {string}   merchant_name  银商名称
 * @apiSuccess (data每个字段说明) {string}   operator       操作者
 * @apiSuccess (data每个字段说明) {string}   order_id       充值订单
 * @apiSuccess (data每个字段说明) {string}   real_name      真实姓名
 * @apiSuccess (data每个字段说明) {string}   remark         备注
 * @apiSuccess (data每个字段说明) {int}      state          状态：0待审核，1成功，2失败
 * @apiSuccess (data每个字段说明) {string}   updated        修改时间
 * @apiSuccess (data每个字段说明) {string}   updated_last   最后更新时间
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "保存数据成功",
 *    "code": 200,
 *    "data": {
 *                "amount": "23.000",
 *                "presented_money": "0",
 *                "bank_address": "中国上海",
 *                "bank_charge_time": "2017-05-25 17:01:13",
 *                "card_number": "65489556848489",
 *                "created": "2017-05-25 17:01:13",
 *                "id": "1",
 *                "ip": "127.0.0.1",
 *                "merchant_id": "1",
 *                "merchant_name": "XXX",
 *                "operator": "yndon",
 *                "order_id": "SDSSDSDF56DS",
 *                "real_name": "胡汉山",
 *                "remark": "测试一下",
 *                "state": "待审核",
 *                "bank_charge_time": "2017-05-25 17:01:13",
 *                "updated": "2019-06-13 10:36:12",
 *                "updated_last": "2019-06-13 10:36:12"
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *SilverMerchantChargeRecordsController) View(ctx *Context) {
	view(ctx, &silverMerchantChargeRecords)
}

/**
 * @api {post} admin/api/auth/v1/silver_merchant_charge_records/allow 银商充值审核通过
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: lyndon</span><br/><br/>
 * <strong>银商充值记录审核通过</strong><br />
 * 业务描述: 银商充值记录审核通过<br />
 * @apiVersion 1.0.0
 * @apiName     allowSilverMerchantChargeRecords
 * @apiGroup    silver_merchant
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }

 * @apiParam (客户端请求参数) {int} 	id   记录编号
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
func (self *SilverMerchantChargeRecordsController) Allow(ctx *Context) {
	responseResult(ctx, silverMerchantChargeRecords.Allow(ctx), "充值审核处理成功")
}

/**
 * @api {post} admin/api/auth/v1/silver_merchant_charge_records/deny 银商充值拒绝
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: lyndon</span><br/><br/>
 * <strong>银商充值记录审核拒绝</strong><br />
 * 业务描述: 银商充值记录审核拒绝<br />
 * @apiVersion 1.0.0
 * @apiName     denySilverMerchantChargeRecords
 * @apiGroup    silver_merchant
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }

 * @apiParam (客户端请求参数) {int} 	id   记录编号
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
func (self *SilverMerchantChargeRecordsController) Deny(ctx *Context) {
	responseResult(ctx, silverMerchantChargeRecords.Deny(ctx), "充值拒绝处理成功")
}
