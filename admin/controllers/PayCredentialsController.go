package controllers

import (
	"qpgame/admin/models"
	"qpgame/admin/validations"
)

var payCredentials = models.PayCredentials{}                          //模型
var payCredentialsValidation = validations.PayCredentialsValidation{} //校验器

type PayCredentialsController struct{}

/**
 * @api {get} admin/api/auth/v1/pay_credentials 支付证书列表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>支付证书列表</strong><br />
 * 业务描述: 支付证书列表</br>
 * @apiVersion 1.0.0
 * @apiName     indexPayCredentials
 * @apiGroup    finance
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int} 	page			页数
 * @apiParam (客户端请求参数) {int} 	page_size		每页记录数
 * @apiParam (客户端请求参数) {string} 	pay_name 名称
 * @apiParam (客户端请求参数) {string} 	merchant_number 商户编号
 * @apiParam (客户端请求参数) {int} 	status 状态
 * @apiParam (客户端请求参数) {string} 	card_number 银行卡号
 * @apiParam (客户端请求参数) {string} 	phone_number 手机号码
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
 * @apiSuccess (data-rows每个子对象字段说明) {int}        plat_form               支付标识,所有平台有共同标识
 * @apiSuccess (data-rows每个子对象字段说明) {string}     pay_name                平台名称
 * @apiSuccess (data-rows每个子对象字段说明) {string}     merchant_number         商户号
 * @apiSuccess (data-rows每个子对象字段说明) {string}     corporate               法人
 * @apiSuccess (data-rows每个子对象字段说明) {string}     id_umber                法人身份证号
 * @apiSuccess (data-rows每个子对象字段说明) {string}     card_number             银行卡号
 * @apiSuccess (data-rows每个子对象字段说明) {string}     phone_number            手机号码
 * @apiSuccess (data-rows每个子对象字段说明) {int}        charge_amount_conf      充值金额配置:0 关闭随机金额小数,1 开启随机金额小数到角,2 开启随机金额小数到分
 * @apiSuccess (data-rows每个子对象字段说明) {int}        status                  状态,0:锁定,1:正常
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "数据获取成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "plat_form": "",
 *        "pay_name": "",
 *        "merchant_number": "",
 *        "private_key": "",
 *        "corporate": "",
 *        "id_umber": "",
 *        "card_number": "",
 *        "phone_number": "",
 *        "public_key": "",
 *        "private_key_file": "",
 *        "credential_key": "",
 *        "callback_key": "",
 *        "charge_amount_conf": "",
 *        "status": ""
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *PayCredentialsController) Index(ctx *Context) {
	index(ctx, &payCredentials)
}

/**
 * @api {get} admin/api/auth/v1/charge_cards/get_pay_credentials 获取所有支付证书数据
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: lyndon</span><br/><br/>
 * <strong>获取所有支付证书数据</strong><br />
 * 业务描述: 获取所有支付证书数据</br>
 * @apiVersion 1.0.0
 * @apiName     getAllPayCredentials
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
 * @apiSuccess (data每个子对象字段说明) {int}		  id 	记录编号
 * @apiSuccess (data每个子对象字段说明) {int}        plat_form               支付标识,所有平台有共同标识
 * @apiSuccess (data每个子对象字段说明) {string}     pay_name                平台名称
 * @apiSuccess (data每个子对象字段说明) {array}     items         通道列表
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "数据获取成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "plat_form": "",
 *        "pay_name": "",
 *        "merchant_number": "",
 *        "private_key": "",
 *        "corporate": "",
 *        "id_umber": "",
 *        "card_number": "",
 *        "phone_number": "",
 *        "public_key": "",
 *        "private_key_file": "",
 *        "credential_key": "",
 *        "callback_key": "",
 *        "charge_amount_conf": "",
 *        "status": ""
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *PayCredentialsController) All(ctx *Context) {
	result, err := payCredentials.GetAll(ctx)
	if err != nil {
		responseFailure(ctx, "", err.Error())
		return
	}
	responseSuccess(ctx, "", result)
}

/**
 * @api {post} admin/api/auth/v1/pay_credentials/add	支付证书添加
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>添加/修改支付证书</strong><br />
 * 业务描述: 添加/修改支付证书 <br />
 * <strong><span style="color: red">注意: </span></strong><br />
 * <span style="color:red">修改操作API不再单独列出, 请参考以下</span><br />
 * <span style="color:red">添加: /admin/api/auth/v1/pay_credentials/add </span> &nbsp;&nbsp; <br />
 * <span style="color:red">修改: /admin/api/auth/v1/pay_credentials/update </span>
 * @apiVersion 1.0.0
 * @apiName     savePayCredentials
 * @apiGroup    finance
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 *
 * @apiParam (客户端请求参数) {int} 	id    			记录编号,仅修改操作时(即接口: /admin/api/auth/v1/pay_credentials/update)需要<br />
 *														添加操作(即接口: /admin/api/auth/v1/pay_credentials/add)不需要此参数<br />
 * 														* 如果提供此编号, 则视为修改记录
 * @apiParam (客户端请求参数) {int}      	plat_form               支付标识,所有平台有共同标识
 * @apiParam (客户端请求参数) {string}   	pay_name                平台名称
 * @apiParam (客户端请求参数) {string}   	merchant_number         商户号
 * @apiParam (客户端请求参数) {string}   	private_key             商户私钥
 * @apiParam (客户端请求参数) {string}   	corporate               法人
 * @apiParam (客户端请求参数) {string}   	id_umber                法人身份证号
 * @apiParam (客户端请求参数) {string}   	card_number             银行卡号
 * @apiParam (客户端请求参数) {string}   	phone_number            手机号码
 * @apiParam (客户端请求参数) {string}   	public_key              证书公钥
 * @apiParam (客户端请求参数) {string}   	private_key_file        私钥文件
 * @apiParam (客户端请求参数) {string}   	credential_key          CredentialKey
 * @apiParam (客户端请求参数) {string}   	callback_key            CallbackKey
 * @apiParam (客户端请求参数) {int}      	charge_amount_conf      充值金额配置:0 关闭随机金额小数,1 开启随机金额小数到角,2 开启随机金额小数到分
 * @apiParam (客户端请求参数) {int}      	status                  是否弃用,0弃用,1启用<br />
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
func (self *PayCredentialsController) Save(ctx *Context) {
	save(ctx, &payCredentials, &payCredentialsValidation)
}

/**
 * @api {get} admin/api/auth/v1/pay_credentials/view 				支付证书详情
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>支付证书详情</strong><br />
 * 业务描述: 支付证书详情</br>
 * @apiVersion 1.0.0
 * @apiName     viewPayCredentials
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
 * @apiSuccess (data字段说明) {int}        plat_form               支付标识,所有平台有共同标识
 * @apiSuccess (data字段说明) {string}     pay_name                平台名称
 * @apiSuccess (data字段说明) {string}     merchant_number         商户号
 * @apiSuccess (data字段说明) {string}     private_key             商户私钥
 * @apiSuccess (data字段说明) {string}     corporate               法人
 * @apiSuccess (data字段说明) {string}     id_umber                法人身份证号
 * @apiSuccess (data字段说明) {string}     card_number             银行卡号
 * @apiSuccess (data字段说明) {string}     phone_number            手机号码
 * @apiSuccess (data字段说明) {string}     public_key              证书公钥
 * @apiSuccess (data字段说明) {string}     private_key_file        私钥文件
 * @apiSuccess (data字段说明) {string}     credential_key          CredentialKey
 * @apiSuccess (data字段说明) {string}     callback_key            CallbackKey
 * @apiSuccess (data字段说明) {int}        charge_amount_conf      充值金额配置:0 关闭随机金额小数,1 开启随机金额小数到角,2 开启随机金额小数到分
 * @apiSuccess (data字段说明) {int}        status                  是否弃用,0弃用,1启用
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "保存数据成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "plat_form": "",
 *        "pay_name": "",
 *        "merchant_number": "",
 *        "private_key": "",
 *        "corporate": "",
 *        "id_umber": "",
 *        "card_number": "",
 *        "phone_number": "",
 *        "public_key": "",
 *        "private_key_file": "",
 *        "credential_key": "",
 *        "callback_key": "",
 *        "charge_amount_conf": "",
 *        "status": ""
 * 	  },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *PayCredentialsController) View(ctx *Context) {
	view(ctx, &payCredentials)
}

/**
 * @api {get} admin/api/auth/v1/pay_credentials/delete 支付证书删除
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>支付证书删除</strong><br />
 * 业务描述: 支付证书删除</br>
 * @apiVersion 1.0.0
 * @apiName     deletePayCredentials
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
func (self *PayCredentialsController) Delete(ctx *Context) {
	remove(ctx, &payCredentials)
}
