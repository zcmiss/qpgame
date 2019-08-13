package controllers

import (
	"qpgame/admin/models"
	"qpgame/admin/validations"
)

var silverMerchantUsers = models.SilverMerchantUsers{}                          //模型
var silverMerchantUsersValidation = validations.SilverMerchantUsersValidation{} //校验器

type SilverMerchantUsersController struct{}

/**
 * @api {get} admin/api/auth/v1/silver_merchant_users 银商用户列表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: lyndon</span><br/><br/>
 * <strong>银商用户列表</strong><br />
 * 业务描述: 银商用户列表</br>
 * @apiVersion 1.0.0
 * @apiName     indexSilverMerchantUsers
 * @apiGroup    silver_merchant
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int}       page            页数
 * @apiParam (客户端请求参数) {int}       page_size       每页记录数
 * @apiParam (客户端请求参数) {string}    title 公告标题
 * @apiParam (客户端请求参数) {int}       merchant_level 银商等级
 * @apiParam (客户端请求参数) {string} 	account 账号
 * @apiParam (客户端请求参数) {string} 	merchant_name 账号
 * @apiParam (客户端请求参数) {int} 	    status 状态,0锁定，1正常
 * @apiParam (客户端请求参数) {int} 	    is_destroy 是否已注销，0否，1是
 * @apiParam (客户端请求参数) {string} 	created_start 添加时间/开始
 * @apiParam (客户端请求参数) {string} 	created_end 添加时间/结束
 * @apiParam (客户端请求参数) {string} 	last_login_start 最后一次登录时间/开始
 * @apiParam (客户端请求参数) {string} 	last_login_end 最后一次登录时间/结束
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
 * @apiSuccess (data-rows每个子对象字段说明) {int}	    id 					  记录编号
 * @apiSuccess (data-rows每个子对象字段说明) {string} account               银商账号
 * @apiSuccess (data-rows每个子对象字段说明) {string} created               创建时间
 * @apiSuccess (data-rows每个子对象字段说明) {float}  donate_rate           赠送比例,这是银商的收入来源很重要,比如冲1万，送4%
 * @apiSuccess (data-rows每个子对象字段说明) {int}    is_destroy            是否已注销,1是，0否
 * @apiSuccess (data-rows每个子对象字段说明) {string} last_login_time       最后一次登录时间
 * @apiSuccess (data-rows每个子对象字段说明) {float}  merchant_cash_pledge  银商押金
 * @apiSuccess (data-rows每个子对象字段说明) {int}    merchant_level        银商等级
 * @apiSuccess (data-rows每个子对象字段说明) {string} merchant_name         商户名称
 * @apiSuccess (data-rows每个子对象字段说明) {int}    status                银商状态,1正常，0锁定
 * @apiSuccess (data-rows每个子对象字段说明) {string} token                 最后一次登录的token
 * @apiSuccess (data-rows每个子对象字段说明) {float}  total_auth_amount     累计授权金额
 * @apiSuccess (data-rows每个子对象字段说明) {float}  total_charge_money    累计充值金额
 * @apiSuccess (data-rows每个子对象字段说明) {float}	usable_amount         可用额度
 * @apiSuccess (data-rows每个子对象字段说明) {int}	    user_id               银商对应的用户表的用户编号
 * @apiSuccess (data-rows每个子对象字段说明) {string}	user_name             银商对应的用户表的用户名称
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *  "clientMsg": "获取数据成功",
 *  "code": 200,
 *  "data": {
 *      "rows": [
 *          {
 *              "account": "test8",
 *              "created": "",
 *              "donate_rate": "0.000",
 *              "id": "3",
 *              "is_destroy": "0",
 *              "last_login_time": "2019-06-11 18:55:15",
 *              "merchant_cash_pledge": "0.000",
 *              "merchant_level": "1",
 *              "merchant_name": "",
 *              "status": "1",
 *              "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjAyNTQxMTUsInN1YiI6M30.7rVhxS_3JxVTmEEm3HBNM3ol0oxQTkA56-6uktZrzdg",
 *              "total_auth_amount": "0.000",
 *              "total_charge_money": "0.000",
 *              "usable_amount": "0.000",
 *              "user_id": "78",
 *              "user_name": "test8"
 *          }
 *      ],
 *      "page": 1,
 *      "page_count": 1,
 *      "total_rows": 300,
 *      "page_size": 20
 *  },
 *  "internalMsg": "",
 *  "timeConsumed": 0
}
*/
func (self *SilverMerchantUsersController) Index(ctx *Context) {
	index(ctx, &silverMerchantUsers)
}

/**
 * @api {post} admin/api/auth/v1/silver_merchant_users/add 银商用户添加
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: lyndon</span><br/><br/>
 * <strong>添加/修改银商用户</strong><br />
 * 业务描述: 添加/修改银商用户 <br />
 * <strong><span style="color: red">注意: </span></strong><br />
 * <span style="color:red">修改操作API不再单独列出, 请参考以下</span><br />
 * <span style="color:red">添加: /admin/api/auth/v1/silver_merchant_users/add </span> &nbsp;&nbsp; <br />
 * <span style="color:red">修改: /admin/api/auth/v1/silver_merchant_users/update </span>
 * @apiVersion 1.0.0
 * @apiName     saveSilverMerchantUsers
 * @apiGroup    silver_merchant
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }

 * @apiParam (客户端请求参数) {int} 	id   记录编号,仅修改操作时(即接口: /admin/api/auth/v1/silver_merchant_users/update)需要<br />
 *										 添加操作(即接口: /admin/api/auth/v1/silver_merchant_users/add)不需要此参数<br />
 * 										 * 如果提供此编号, 则视为修改记录
 * @apiParam (客户端请求参数) {int}     user_id                 用户编号
 * @apiParam (客户端请求参数) {int}     merchant_level          银网用户等级,可选数组，目前为空，默认为1就可以
 * @apiParam (客户端请求参数) {float}   merchant_cash_pledge    银商押金
 * @apiParam (客户端请求参数) {float}   donate_rate             赠送比例
 * @apiParam (客户端请求参数) {string}  account                 银商账号
 * @apiParam (客户端请求参数) {string}  password                银商密码
 * @apiParam (客户端请求参数) {string}  merchant_name           商户名称
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
func (self *SilverMerchantUsersController) Save(ctx *Context) {
	save(ctx, &silverMerchantUsers, &silverMerchantUsersValidation)
}

/**
 * @api {get} admin/api/auth/v1/silver_merchant_users/view 			银商用户详情
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: lyndon</span><br/><br/>
 * <strong>银商用户详情</strong><br />
 * 业务描述: 银商用户详情</br>
 * @apiVersion 1.0.0
 * @apiName     viewSilverMerchantUsers
 * @apiGroup    silver_merchant
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }

 * @apiParam (客户端请求参数) {int} 	id    			  记录编号
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
 * @apiSuccess (data每个字段说明) {int}	    info_id            银商编号
 * @apiSuccess (data每个字段说明) {string} info_account               银商账号
 * @apiSuccess (data每个字段说明) {string} info_created               创建时间
 * @apiSuccess (data每个字段说明) {float}  info_donate_rate           赠送比例,这是银商的收入来源很重要,比如冲1万，送4%
 * @apiSuccess (data每个字段说明) {int}    info_is_destroy            是否已注销,1是，0否
 * @apiSuccess (data每个字段说明) {string} info_last_login_time       最后一次登录时间
 * @apiSuccess (data每个字段说明) {float}  info_merchant_cash_pledge  银商押金
 * @apiSuccess (data每个字段说明) {int}    info_merchant_level        银商等级
 * @apiSuccess (data每个字段说明) {string} info_merchant_name         商户名称
 * @apiSuccess (data每个字段说明) {int}    info_status                银商状态,1正常，0锁定
 * @apiSuccess (data每个字段说明) {string} info_token                 最后一次登录的token
 * @apiSuccess (data每个字段说明) {float}  info_total_auth_amount     累计授权金额
 * @apiSuccess (data每个字段说明) {float}  info_total_charge_money    累计充值金额
 * @apiSuccess (data每个字段说明) {float}	  info_usable_amount         可用额度
 * @apiSuccess (data每个字段说明) {int}	  info_user_id               银商对应的用户表的用户编号
 * @apiSuccess (data每个字段说明) {string}  info_user_name            银商对应的用户表的用户名称

 * @apiSuccess (data每个字段说明) {int}	     bank_card_id            银行卡编号
 * @apiSuccess (data每个字段说明) {string}	 bank_card_address       银行卡地址
 * @apiSuccess (data每个字段说明) {string}	 bank_card_bank_name     银行名称
 * @apiSuccess (data每个字段说明) {string}	 bank_card_card_number   银行卡卡号
 * @apiSuccess (data每个字段说明) {string}	 bank_card_created       银行卡添加时间
 * @apiSuccess (data每个字段说明) {string}  bank_card_name          银行卡用户姓名
 * @apiSuccess (data每个字段说明) {string}  bank_card_remark        银行卡备注
 * @apiSuccess (data每个字段说明) {int}     bank_card_status        银行卡状态：0失效，1正常
 * @apiSuccess (data每个字段说明) {string}  bank_card_updated       银行卡更新时间
 * @apiSuccess (data每个字段说明) {float}  flow_charge_amount        额度充值总额
 * @apiSuccess (data每个字段说明) {int}	  flow_charge_count         额度充值总次数
 * @apiSuccess (data每个字段说明) {float}	  flow_presented_money_amount           额度充值赠送总金额
 * @apiSuccess (data每个字段说明) {float}  flow_user_charge_amount          银商给会员充值总金额
 * @apiSuccess (data每个字段说明) {int}  flow_user_charge_count          银商给会员充值总次数
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "保存数据成功",
 *    "code": 200,
 *    "data": {
 *        "info_account": "test8",
 *        "info_created": "",
 *        "info_donate_rate": "0.000",
 *        "info_id": "3",
 *        "info_is_destroy": "0",
 *        "info_last_login_time": "2019-06-11 18:55:15",
 *        "info_merchant_cash_pledge": "0.000",
 *        "info_merchant_level": "1",
 *        "info_merchant_name": "",
 *        "info_status": "1",
 *        "info_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjAyNTQxMTUsInN1YiI6M30.7rVhxS_3JxVTmEEm3HBNM3ol0oxQTkA56-6uktZrzdg",
 *        "info_total_auth_amount": "0.000",
 *        "info_total_charge_money": "0.000",
 *        "info_usable_amount": "0.000",
 *        "info_user_id": "78",
 *        "info_user_name": "test8",
 *        "bank_card_id": "1",
 *        "bank_card_address": "sdsds",
 *        "bank_card_bank_name": "sdfsdfs",
 *        "bank_card_card_number": "sdfsdfsd",
 *        "bank_card_created": "232",
 *        "bank_card_name": "asda",
 *        "bank_card_remark": "vsda",
 *        "bank_card_status": "2",
 *        "bank_card_updated": "234234",
 *        "flow_charge_amount": "2342.234",
 *        "flow_charge_count": "23",
 *        "flow_presented_money_amount": "2342.234",
 *        "flow_user_charge_amount": "2342.234",
 *        "flow_user_charge_count": "23"
 * 	  },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *SilverMerchantUsersController) View(ctx *Context) {
	view(ctx, &silverMerchantUsers)
}

/**
 * @api {get} admin/api/auth/v1/silver_merchant_users/delete 银商用户删除
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: lyndon</span><br/><br/>
 * <strong>银商用户删除</strong><br />
 * 业务描述: 银商用户删除</br>
 * @apiVersion 1.0.0
 * @apiName     deleteSilverMerchantUsers
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
func (self *SilverMerchantUsersController) Delete(ctx *Context) {
	remove(ctx, &silverMerchantUsers)
}
