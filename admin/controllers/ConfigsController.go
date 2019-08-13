package controllers

import (
	"qpgame/admin/models"
	"qpgame/admin/validations"
)

var configs = models.Configs{}                          //模型
var configsValidation = validations.ConfigsValidation{} //校验器

type ConfigsController struct{}

/**
 * @api {get} admin/api/auth/v1/configs 系统配置列表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>系统配置列表</strong><br />
 * 业务描述: 系统配置列表</br>
 * @apiVersion 1.0.0
 * @apiName     indexConfigs
 * @apiGroup    system
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int} 	page			页数
 * @apiParam (客户端请求参数) {int} 	page_size		每页记录数
 *
 * @apiError(请求失败返回) {int}      code            错误代码
 * @apiError(请求失败返回) {string}   clientMsg       提示信息
 * @apiError(请求失败返回) {string}   internalMsg     内部错误信息
 * @apiError(请求失败返回) {float}    timeConsumed    后台耗时
 *
 * @apiErrorExample {json} 失败返回
 * {
 *      "code": 204,
 *      "internalMsg": "",
 *      "clientMsg ": 0,
 *      "timeConsumed": 0
 * }
 *
 * @apiSuccess(返回结果)  {int}    	code            200, 成功
 * @apiSuccess(返回结果)  {string} 	clientMsg       提示信息
 * @apiSuccess(返回结果)  {string} 	internalMsg     内部错误信息
 * @apiSuccess(返回结果)  {json}  	data            返回数据
 * @apiSuccess(返回结果)  {float}  	timeConsumed    后台耗时
 *
 * @apiSuccess(data字段说明) {array}  	rows        数据列表
 * @apiSuccess(data字段说明) {int}    	page		当前页数
 * @apiSuccess(data字段说明) {int}    	page_count	总的页数
 * @apiSuccess(data字段说明) {int}    	total_rows	总记录数
 * @apiSuccess(data字段说明) {int}    	page_size	每页记录数
 *
 * @apiSuccess(data-rows每个子对象字段说明) {int}		  id 			记录编号
 * @apiSuccess(data-rows每个子对象字段说明) {string}     name           设置项名称
 * @apiSuccess(data-rows每个子对象字段说明) {string}     value                  设置参数值
 * @apiSuccess(data-rows每个子对象字段说明) {string}     mark                   记号说明(该字段只作数据库字段应用说明,不作程序使用
 * @apiSuccess(data-rows每个子对象字段说明) {int}        updated                更新时间
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "数据获取成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "name": "",
 *        "value": "",
 *        "mark": "",
 *        "updated": ""
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *ConfigsController) Index(ctx *Context) {
	index(ctx, &configs)
}

/**
 * @api {get} admin/api/auth/v1/configs/sets	平台参数获取
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>系统配置获取</strong><br />
 * 业务描述: 系统配置获取 <br />
 * @apiVersion 1.0.0
 * @apiName     sets
 * @apiGroup    system
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
 * @apiSuccess (data字段说明) {float}   	generalize_award 		推广分享奖励金额
 * @apiSuccess (data字段说明) {int}   	withdraw_min_money 		提现最低金额
 * @apiSuccess (data字段说明) {int}   	withdraw_max_money 		提现最高金额
 * @apiSuccess (data字段说明) {int}		clear_tx_limit 			清除提现打码阀值
 * @apiSuccess (data字段说明) {int}   	withdraw_day_limited 	每天提款次数限制
 * @apiSuccess (data字段说明) {string}   tuiguang_web_url 		代理推广地址
 * @apiSuccess (data字段说明) {string}   tuiguang_web_domain		推广下域名
 * @apiSuccess (data字段说明) {int}   sign_reward			签到奖励设置
 * @apiSuccess (data字段说明) {int}   register_number_ip		每IP允许最多注册用户数量
 * @apiSuccess (data字段说明) {int}   reward_bind	绑定手机奖励
 * @apiSuccess (data字段说明) {int}   sign_award_switch	是否开启签到奖励,0否，1是
 * @apiSuccess (data字段说明) {int}   bind_phone_award_switch	 是否开启绑定手机奖励,0否，1是
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "保存数据成功",
 *    "code": 200,
 *    "data": {
 *        "withdraw_min_money": "",
 *        "withdraw_max_money": "",
 *        "clear_tx_limit": "",
 *        "withdraw_day_limited": "",
 *        "tuiguang_web_url": "",
 *        "sign_reward": "",
 *        "register_number_ip": "",
 *        "reward_bind": "",
 *        "sign_award_switch": "0",
 *        "bind_phone_award_switch": "0"
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *ConfigsController) Sets(ctx *Context) {
	responseSuccess(ctx, "", configs.GetSets(ctx))
}

/**
 * @api {post} admin/api/auth/v1/configs/set_sets	平台参数设置
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>系统配置保存</strong><br />
 * 业务描述: 系统配置保存<br />
 * @apiVersion 1.0.0
 * @apiName     setSets
 * @apiGroup    system
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {float} generalize_award 		推广分享奖励金额
 * @apiParam (客户端请求参数) {int}   	withdraw_min_money 		提现最低金额
 * @apiParam (客户端请求参数) {int}   	withdraw_max_money 		提现最高金额
 * @apiParam (客户端请求参数) {int}		clear_tx_limit 			清除提现打码阀值
 * @apiParam (客户端请求参数) {int}   	withdraw_day_limited 	每天提款次数限制
 * @apiParam (客户端请求参数) {string}   	tuiguang_web_url 		代理推广地址
 * @apiParam (客户端请求参数) {string}   	tuiguang_web_domain 		推广下载域名
 * @apiParam (客户端请求参数) {int}   	sign_reward 		签到奖励
 * @apiParam (客户端请求参数) {int}   	register_number_ip	每IP允许最多注册用户数量
 * @apiParam (客户端请求参数) {int}   	reward_bind		绑定手机奖励
 * @apiParam (客户端请求参数) {int}   sign_award_switch	是否开启签到奖励,0否，1是
 * @apiParam (客户端请求参数) {int}   bind_phone_award_switch	 是否开启绑定手机奖励,0否，1是
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
func (self *ConfigsController) SetSets(ctx *Context) {
	responseResult(ctx, configs.Sets(ctx), "配置保存成功")
}

/**
 * @api {get} admin/api/auth/v1/configs/faq	FAQ获取
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>FAQ获取</strong><br />
 * 业务描述: FAQ获取<br />
 * @apiVersion 1.0.0
 * @apiName     faq
 * @apiGroup    system
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }

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
 *      "timeConsumed": 0,
 * }
 *
 * @apiSuccess (返回结果)  {int} 		code            200
 * @apiSuccess (返回结果)  {string} 	clientMsg       提示信息
 * @apiSuccess (返回结果)  {string}  	internalMsg     内部错误信息
 * @apiSuccess (返回结果)  {json}		data            返回数据
 * @apiSuccess (返回结果)  {float}   	timeConsumed    后台耗时
 *
 * @apiSuccess (data字段说明) {string}   	faq FAQ
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
func (self *ConfigsController) Faq(ctx *Context) {
	responseSuccess(ctx, "获取成功", configs.GetFaq(ctx))
}

/**
 * @api {post} admin/api/auth/v1/configs/set_faq	FAQ设置
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>FAQ设置</strong><br />
 * 业务描述: FAQ设置 <br />
 * @apiVersion 1.0.0
 * @apiName     setFaq
 * @apiGroup    system
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {string}   	faq  	FAQ内容
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
func (self *ConfigsController) SetFaq(ctx *Context) {
	responseResult(ctx, configs.Faq(ctx), "配置保存成功")
}

/**
 * @api {get} admin/api/auth/v1/configs/service	客服信息获取
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>客服信息获取</strong><br />
 * 业务描述: 客服信息获取<br />
 * @apiVersion 1.0.0
 * @apiName     service
 * @apiGroup    system
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
 * @apiSuccess (data-字段说明) {array}   	qq	QQ客户联系方式
 * @apiSuccess (data-字段说明) {array}   	wx	微信联系方式

 * @apiSuccess (data-zx字段说明) {string} 	url 在线客服链接地址
 *
 * @apiSuccess (data-qq字段说明) {string} 	url 客户链接地址
 * @apiSuccess (data-qq字段说明) {string} 	name 名称
 * @apiSuccess (data-qq字段说明) {string} 	account 账号
 * @apiSuccess (data-qq字段说明) {string} 	info 说明
 *
 * @apiSuccess (data-wx字段说明) {string} 	url 客户链接地址
 * @apiSuccess (data-wx字段说明) {string} 	name 名称
 * @apiSuccess (data-wx字段说明) {string} 	account 账号
 * @apiSuccess (data-wx字段说明) {string} 	info 说明
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
func (self *ConfigsController) Service(ctx *Context) {
	responseSuccess(ctx, "", configs.GetService(ctx))
}

/**
 * @api {post} admin/api/auth/v1/configs/set_service 客服信息设置
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>客服信息设置</strong><br />
 * 业务描述: 客服信息设置 <br />
 * @apiVersion 1.0.0
 * @apiName     setService
 * @apiGroup    system
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 *
 * @apiParam (zx字段说明) {string} 	url 在线客服链接地址
 *
 * @apiParam (客户端请求参数) {array} 	qq	QQ
 * @apiParam (客户端请求参数) {array} 	wx	微信
 *
 * @apiParam (qq字段说明) {string} 	url 客户链接地址
 * @apiParam (qq字段说明) {string} 	name 名称
 * @apiParam (qq字段说明) {string} 	account 账号
 * @apiParam (qq字段说明) {string} 	info 说明
 *
 * @apiParam (wx字段说明) {string} 	url 客户链接地址
 * @apiParam (wx字段说明) {string} 	name 名称
 * @apiParam (wx字段说明) {string} 	account 账号
 * @apiParam (wx字段说明) {string} 	info 说明
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
func (self *ConfigsController) SetService(ctx *Context) {
	responseResult(ctx, configs.Service(ctx), "配置保存成功")
}

/**
 * @api {get} admin/api/auth/v1/configs/charges	充值配置信息获取
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>充值配置信息获取</strong><br />
 * 业务描述: 充值配置信息获取<br />
 * @apiVersion 1.0.0
 * @apiName     charges
 * @apiGroup    system
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
 * @apiSuccess (data字段说明) {array} proxy_charge_logo		代理充值LOGO图片地址
 * @apiSuccess (data字段说明) {string} info		说明信息
 * @apiSuccess (data字段说明) {string} state		状态
 * @apiSuccess (data字段说明) {array} charge_accounts	充值账号
 *
 * @apiSuccess (data-charge_accounts字段说明) {string} url 图片链接地址
 * @apiSuccess (data-charge_accounts字段说明) {string} name 名称
 * @apiSuccess (data-charge_accounts字段说明) {string} account 账号
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
func (self *ConfigsController) Charges(ctx *Context) {
	responseSuccess(ctx, "", configs.GetCharges(ctx))
}

/**
 * @api {post} admin/api/auth/v1/configs/set_charges 	充值配置信息设置
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>充值配置信息设置</strong><br />
 * 业务描述: 充值配置信息设置<br />
 * @apiVersion 1.0.0
 * @apiName     setCharges
 * @apiGroup    system
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {array} proxy_charge_logo		代理充值LOGO图片地址
 * @apiParam (客户端请求参数) {string} info		说明信息
 * @apiParam (客户端请求参数) {string} state		状态
 * @apiParam (客户端请求参数) {array} charge_accounts	充值账号
 *
 * @apiParam (charge_accounts字段说明) {string} url 图片链接地址
 * @apiParam (charge_accounts字段说明) {string} name 名称
 * @apiParam (charge_accounts字段说明) {string} account 账号
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
func (self *ConfigsController) SetCharges(ctx *Context) {
	responseResult(ctx, configs.Charges(ctx), "配置保存成功")
}

/**
 * @api {get} admin/api/auth/v1/configs/reg	注册配置获取
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>注册配置获取</strong><br />
 * 业务描述: 注册配置获取<br />
 * @apiVersion 1.0.0
 * @apiName     reg
 * @apiGroup    system
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
 * @apiError (请求失败返回) {float}    timeConsumed   	 后台耗时
 *
 * @apiErrorExample {json} 失败返回
 * {
 *      "code": 204,
 *      "internalMsg": "",
 *      "clientMsg ": "",
 *      "timeConsumed": 0,
 * }
 *
 * @apiSuccess (返回结果)  {int} 		code            200
 * @apiSuccess (返回结果)  {string} 	clientMsg       提示信息
 * @apiSuccess (返回结果)  {string}  	internalMsg     内部错误信息
 * @apiSuccess (返回结果)  {json}		data            返回数据
 * @apiSuccess (返回结果)  {float}   	timeConsumed    后台耗时
 *
 * @apiSuccess (data字段说明) {int}   	can_register 	是否允许注册
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
func (self *ConfigsController) Reg(ctx *Context) {
	responseSuccess(ctx, "获取成功", configs.GetReg(ctx))
}

/**
 * @api {post} admin/api/auth/v1/configs/set_reg	注册配置设置
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>注册配置设置</strong><br />
 * 业务描述: 注册配置设置 <br />
 * @apiVersion 1.0.0
 * @apiName     setReg
 * @apiGroup    system
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int} can_register 是否允许注册
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
func (self *ConfigsController) SetReg(ctx *Context) {
	responseResult(ctx, configs.Reg(ctx), "配置保存成功")
}

/**
 * @api {get} admin/api/auth/v1/configs/order_alert 订单提醒获置
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>订单提醒配置获取</strong><br />
 * 业务描述: 订单提醒配置获取<br />
 * @apiVersion 1.0.0
 * @apiName     orderAlert
 * @apiGroup    system
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
 * @apiError (请求失败返回) {float}    timeConsumed   	 后台耗时
 *
 * @apiErrorExample {json} 失败返回
 * {
 *      "code": 204,
 *      "internalMsg": "",
 *      "clientMsg ": "",
 *      "timeConsumed": 0,
 * }
 *
 * @apiSuccess (返回结果)  {int} 		code            200
 * @apiSuccess (返回结果)  {string} 	clientMsg       提示信息
 * @apiSuccess (返回结果)  {string}  	internalMsg     内部错误信息
 * @apiSuccess (返回结果)  {json}		data            返回数据
 * @apiSuccess (返回结果)  {float}   	timeConsumed    后台耗时
 *
 * @apiSuccess (data字段说明) {string}   charge_url 	充值提醒声音地址
 * @apiSuccess (data字段说明) {string}   withdraw_url	提现提醒声音地址
 * @apiSuccess (data字段说明) {int}   	timeout		间隔刷新时间/秒
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
func (self *ConfigsController) OrderAlert(ctx *Context) {
	responseSuccess(ctx, "获取成功", configs.GetOrderAlert(ctx))
}

/**
 * @api {get} admin/api/auth/v1/configs/fund_dama_rate 资金打码量比例列表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: xiaoye</span><br/><br/>
 * <strong>资金打码量比例列表</strong><br />
 * 业务描述: 资金打码量比例列表</br>
 * @apiVersion 1.0.0
 * @apiName     indexConfigs
 * @apiGroup    system
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int} 	page			页数
 * @apiParam (客户端请求参数) {int} 	page_size		每页记录数
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
 * @apiSuccess (返回结果)  {json}  	data            返回数据
 * @apiSuccess (返回结果)  {float}  	timeConsumed    后台耗时
 *
 *
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    ""1": {"dama_rate": 1,"label": "充值"},
 *  	"1000": {"dama_rate": 1,"label": "其他,未列出资金类型"},
 *  	"14": {"dama_rate": 1,"label": "红包收入"},
 *  	"16": {"dama_rate": 1,"label": "VIP晋级礼金"},
 *  	"17": {"dama_rate": 1,"label": "VIP晋级周礼金"},
 *   	"18": {"dama_rate": 1,"label": "VIP晋级月礼金"},
 *   	"19": {"dama_rate": 1,"label": "推广邀请"},
 *  	"3": {"dama_rate": 1,"label": "洗码"},
 *  	"5": {"dama_rate": 1,"label": "赠送彩金"},
 * 		"6": {"dama_rate": 1,"label": "优惠入款"},
 *  	"8": {"dama_rate": 1,"label": "签到奖励"},
 *		"9": {"dama_rate": 1,"label": "活动奖励"}
 * }
 */
func (self *ConfigsController) FundDamaRate(ctx *Context) {
	responseSuccess(ctx, "获取成功", configs.FundDamaRate(ctx))
}

/**
 * @api {get} admin/api/auth/v1/configs/set_fund_dama_rate	 资金打码量比例列表设置
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: xiaoye</span><br/><br/>
 * <strong>系统配置获取</strong><br />
 * 业务描述: 系统配置获取 <br />
 * @apiVersion 1.0.0
 * @apiName     sets
 * @apiGroup    system
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
 * @apiSuccess (data字段说明) {json}	 1 		充值打码量比率配置
 * @apiSuccess (data字段说明) {json} 	 1000 	其他,未列出资金类型收入打码量比率配置
 * @apiSuccess (data字段说明) {json} 	 14 	红包收入打码量比率配置
 * @apiSuccess (data字段说明) {json}	 16 	VIP晋级礼金打码量比率配置
 * @apiSuccess (data字段说明) {json}	 17	 	VIP晋级周礼金打码量比率配置
 * @apiSuccess (data字段说明) {json} 	 18 	VIP晋级月礼金打码量比率配置
 * @apiSuccess (data字段说明) {json} 	 19 	推广邀请打码量比率配置
 * @apiSuccess (data字段说明) {json}	 3 		洗码打码量比率配置
 * @apiSuccess (data字段说明) {json}	 5 		赠送彩金打码量比率配置
 * @apiSuccess (data字段说明) {json}	 6 		优惠入款打码量比率配置
 * @apiSuccess (data字段说明) {json}	 8 		签到奖励打码量比率配置
 * @apiSuccess (data字段说明) {json}	 9 		活动奖励打码量比率配置
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
func (self *ConfigsController) SetFundDamaRate(ctx *Context) {
	responseSuccess(ctx, "打码量比率配置修改成功", configs.FundDamaRateSet(ctx))
}

/**
 * @api {get} admin/api/auth/v1/configs/com_bank_present_rate	公司入款银行卡转账赠送比例配置获取
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: xiaoye</span><br/><br/>
 * <strong>公司入款银行卡转账赠送比例配置获取</strong><br />
 * 业务描述: 公司入款银行卡转账赠送比例配置获取<br />
 * @apiVersion 1.0.0
 * @apiName     combankpresentrate
 * @apiGroup    system
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
 * @apiError (请求失败返回) {float}    timeConsumed   	 后台耗时
 *
 * @apiErrorExample {json} 失败返回
 * {
 *      "code": 204,
 *      "internalMsg": "",
 *      "clientMsg ": "",
 *      "timeConsumed": 0,
 * }
 *
 * @apiSuccess (返回结果)  {int} 		code            200
 * @apiSuccess (返回结果)  {string} 	clientMsg       提示信息
 * @apiSuccess (返回结果)  {string}  	internalMsg     内部错误信息
 * @apiSuccess (返回结果)  {json}		data            返回数据
 * @apiSuccess (返回结果)  {float}   	timeConsumed    后台耗时
 *
 * @apiSuccess (data字段说明) {int}   	com_bank_present_rate 	赠送彩金比例
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
func (self *ConfigsController) BankPresentRate(ctx *Context) {
	responseSuccess(ctx, "获取成功", configs.BankPresentRate(ctx))
}

/**
 * @api {post} admin/api/auth/v1/configs/set_com_bank_present_rate	公司入款银行卡转账赠送比例配置设置
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: xiaoye</span><br/><br/>
 * <strong>公司入款银行卡转账赠送比例配置设置</strong><br />
 * 业务描述: 公司入款银行卡转账赠送比例配置设置 <br />
 * @apiVersion 1.0.0
 * @apiName     combankpresentrateset
 * @apiGroup    system
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {float} com_bank_present_rate_set 彩金比例配置设置
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
func (self *ConfigsController) SetBankPresentRate(ctx *Context) {
	if err := configs.BankPresentRateSet(ctx); err != nil {
		responseFailure(ctx, err.Error(), "数据保存失败")
		return
	}
	responseSuccess(ctx, "赠送彩金比例配置修改成功", nil)
}

/**
 * @api {get} admin/api/auth/v1/configs/silver_merchant	银商配置信息获取
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: lyndon</span><br/><br/>
 * <strong>银商配置信息获取</strong><br />
 * 业务描述: 银商配置信息获取<br />
 * @apiVersion 1.0.0
 * @apiName     silver_merchant
 * @apiGroup    system
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
 *      "clientMsg ": "",
 *      "timeConsumed": 0
 * }
 *
 * @apiSuccess (返回结果)  {int} 		code            200
 * @apiSuccess (返回结果)  {string} 	    clientMsg       提示信息
 * @apiSuccess (返回结果)  {string}  	internalMsg     内部错误信息
 * @apiSuccess (返回结果)  {json}		data            返回数据
 * @apiSuccess (返回结果)  {float}   	timeConsumed    后台耗时
 *
 * @apiSuccess (data字段说明) {float} cash_pledge		 押金
 * @apiSuccess (data字段说明) {float} min_charge_money	 最低充值金额
 * @apiSuccess (data字段说明) {float} min_transfer_money	 最低给会员充值的金额
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "",
 *    "code": 200,
 *    "data": {
 *        "cash_pledge": 0.15,
 *        "min_charge_money": 0.21,
 *        "min_transfer_money": 15.25
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *ConfigsController) SilverMerchant(ctx *Context) {
	responseSuccess(ctx, "", configs.GetSilverMerchant(ctx))
}

/**
 * @api {post} admin/api/auth/v1/configs/set_silver_merchant 	银商配置信息设置
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: lyndon</span><br/><br/>
 * <strong>银商配置信息设置</strong><br />
 * 业务描述: 银商配置信息设置<br />
 * @apiVersion 1.0.0
 * @apiName     setSilverMerchant
 * @apiGroup    system
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {float} cash_pledge	        押金
 * @apiParam (客户端请求参数) {float} min_charge_money    最低充值金额
 * @apiParam (客户端请求参数) {float} min_transfer_money	最低给会员充值的金额
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
 * @apiSuccess (返回结果)  {string} 	    clientMsg       提示信息
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
func (self *ConfigsController) SetSilverMerchant(ctx *Context) {
	responseResult(ctx, configs.SilverMerchantSet(ctx), "配置保存成功")
}
