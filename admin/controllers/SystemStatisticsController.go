package controllers

import (
	"qpgame/admin/models"
)

var systemStatistics = models.SystemStatistics{} //模型
//var systemStatisticsValidation = validations.SystemStatisticsValidation{} //校验器

type SystemStatisticsController struct{}

/**
 * @api {get} admin/api/auth/v1/system_statistics 运营统计列表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>运营统计列表</strong><br />
 * 业务描述: 运营统计列表</br>
 * @apiVersion 1.0.0
 * @apiName     indexSystemStatistics
 * @apiGroup    report
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int}    page            页数
 * @apiParam (客户端请求参数) {int}   page_size       每页记录数
 * @apiParam (客户端请求参数) {string}   ymd_start 日期/开始
 * @apiParam (客户端请求参数) {string}   ymd_end 日期/结束
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
 * @apiSuccess (data-rows每个子对象字段说明) {int}        ymd                     统计日期
 * @apiSuccess (data-rows每个子对象字段说明) {float}      charge                  充值总额
 * @apiSuccess (data-rows每个子对象字段说明) {float}      withdraw                提现总额
 * @apiSuccess (data-rows每个子对象字段说明) {float}      deductions              扣除总额
 * @apiSuccess (data-rows每个子对象字段说明) {float}      bet_amount              下注总额
 * @apiSuccess (data-rows每个子对象字段说明) {int}        bet_count               下注总数
 * @apiSuccess (data-rows每个子对象字段说明) {int}        charge_count            充值次数
 * @apiSuccess (data-rows每个子对象字段说明) {int}        charge_user_count       充值人数
 * @apiSuccess (data-rows每个子对象字段说明) {int}        withdraw_count          提现次数
 * @apiSuccess (data-rows每个子对象字段说明) {int}        withdraw_user_count     提现人数
 * @apiSuccess (data-rows每个子对象字段说明) {int}        sale_ratio              销售返点
 * @apiSuccess (data-rows每个子对象字段说明) {float}      winning                 中奖金额
 * @apiSuccess (data-rows每个子对象字段说明) {float}      proxy_ratio             代理返点
 * @apiSuccess (data-rows每个子对象字段说明) {float}      active                  活动奖励
 * @apiSuccess (data-rows每个子对象字段说明) {float}      user_win                团队盈亏
 * @apiSuccess (data-rows每个子对象字段说明) {float}      give_win                派彩损益
 * @apiSuccess (data-rows每个子对象字段说明) {float}      real_win                实际盈亏
 * @apiSuccess (data-rows每个子对象字段说明) {int}        profit                  利润
 * @apiSuccess (data-rows每个子对象字段说明) {int}        first_charge            首充人数
 * @apiSuccess (data-rows每个子对象字段说明) {float}      first_charge_amount     首充金额
 * @apiSuccess (data-rows每个子对象字段说明) {int}        proxy_count             代理总数
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "数据获取成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "ymd": "",
 *        "charge": "",
 *        "withdraw": "",
 *        "deductions": "",
 *        "bet_amount": "",
 *        "bet_count": "",
 *        "charge_count": "",
 *        "charge_user_count": "",
 *        "withdraw_count": "",
 *        "withdraw_user_count": "",
 *        "sale_ratio": "",
 *        "winning": "",
 *        "proxy_ratio": "",
 *        "active": "",
 *        "user_win": "",
 *        "give_win": "",
 *        "real_win": "",
 *        "profit": "",
 *        "reg_user": "",
 *        "bet_new": "",
 *        "deposit_user": "",
 *        "first_charge": "",
 *        "first_charge_amount": "",
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *SystemStatisticsController) Index(ctx *Context) {
	index(ctx, &systemStatistics)
}

/**
 * @api {get} admin/api/auth/v1/system_statistics/view 				运营统计详情
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>运营统计详情</strong><br />
 * 业务描述: 运营统计详情</br>
 * @apiVersion 1.0.0
 * @apiName     viewSystemStatistics
 * @apiGroup    report
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
 * @apiSuccess (data字段说明) {int}        ymd                     统计日期
 * @apiSuccess (data字段说明) {float}      charge                  充值总额
 * @apiSuccess (data字段说明) {float}      withdraw                提现总额
 * @apiSuccess (data字段说明) {float}      deductions              扣除总额
 * @apiSuccess (data字段说明) {float}      bet_amount              下注总额
 * @apiSuccess (data字段说明) {int}        bet_count               下注总数
 * @apiSuccess (data字段说明) {int}        charge_count            充值次数
 * @apiSuccess (data字段说明) {int}        charge_user_count       充值人数
 * @apiSuccess (data字段说明) {int}        withdraw_count          提现次数
 * @apiSuccess (data字段说明) {int}        withdraw_user_count     提现人数
 * @apiSuccess (data字段说明) {int}        sale_ratio              销售返点
 * @apiSuccess (data字段说明) {float}      winning                 中奖金额
 * @apiSuccess (data字段说明) {float}      proxy_ratio             代理返点
 * @apiSuccess (data字段说明) {float}      active                  活动奖励
 * @apiSuccess (data字段说明) {float}      user_win                团队盈亏
 * @apiSuccess (data字段说明) {float}      give_win                派彩损益
 * @apiSuccess (data字段说明) {float}      real_win                实际盈亏
 * @apiSuccess (data字段说明) {int}        profit                  利润
 * @apiSuccess (data字段说明) {int}        first_charge            首充人数
 * @apiSuccess (data字段说明) {float}      first_charge_amount     首充金额
 * @apiSuccess (data字段说明) {int}        proxy_count             代理总数
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "保存数据成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "ymd": "",
 *        "charge": "",
 *        "withdraw": "",
 *        "deductions": "",
 *        "bet_amount": "",
 *        "bet_count": "",
 *        "charge_count": "",
 *        "charge_user_count": "",
 *        "withdraw_count": "",
 *        "withdraw_user_count": "",
 *        "sale_ratio": "",
 *        "winning": "",
 *        "proxy_ratio": "",
 *        "active": "",
 *        "user_win": "",
 *        "give_win": "",
 *        "real_win": "",
 *        "profit": "",
 *        "reg_user": "",
 *        "bet_new": "",
 *        "deposit_user": "",
 *        "first_charge": "",
 *        "first_charge_amount": "",
 *        "proxy_count": "",
 * 	  },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *SystemStatisticsController) View(ctx *Context) {
	view(ctx, &systemStatistics)
}

/**
 * @api {get} admin/api/auth/v1/system_statistics/delete 运营统计删除
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>运营统计删除</strong><br />
 * 业务描述: 运营统计删除</br>
 * @apiVersion 1.0.0
 * @apiName     deleteSystemStatistics
 * @apiGroup    report
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
func (self *SystemStatisticsController) Delete(ctx *Context) {
	remove(ctx, &systemStatistics)
}

/**
 * @api {get} admin/api/auth/v1/system_statistics/counts 			出入款汇总
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>出入款汇总</strong><br />
 * 业务描述: 出入款汇总</br>
 * @apiVersion 1.0.0
 * @apiName     viewSystemStatisticsCounts
 * @apiGroup    finance
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 *
 * @apiParam (客户端请求参数) {string} 	ymd_start    	统计日期/开始
 * @apiParam (客户端请求参数) {string} 	ymd_end    	统计日期/结束
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
 * @apiSuccess (data字段说明) {float}    in_total 				总收入
 * @apiSuccess (data字段说明) {int}      in_total_count			总充值人数(线上入款+公司入款+人工入款)
 * @apiSuccess (data字段说明) {float}    in_company				公司入款
 * @apiSuccess (data字段说明) {int}      in_company_count		公司入款人数
 * @apiSuccess (data字段说明) {float}    in_online				线上入款
 * @apiSuccess (data字段说明) {int}      in_online_count			线上入款人数
 * @apiSuccess (data字段说明) {float}    in_manual				人工入款
 * @apiSuccess (data字段说明) {int}      in_manual_count			人工入款人数
 * @apiSuccess (data字段说明) {float}    in_deduction       		出款扣除
 * @apiSuccess (data字段说明) {int}      in_deduction_count 		出款扣除人数
 * @apiSuccess (data字段说明) {float}    in_first				首次存入金额
 * @apiSuccess (data字段说明) {int}      in_first_count			首次存入人数
 * @apiSuccess (data字段说明) {float}    out_total				总支出
 * @apiSuccess (data字段说明) {int}      out_total_count			总支出人数(线上提现+人工出款+给予优惠+给予反水)
 * @apiSuccess (data字段说明) {float}    out_online				线上出款
 * @apiSuccess (data字段说明) {int}      out_online_count		线上出款人数
 * @apiSuccess (data字段说明) {float}    out_manual				人工出款
 * @apiSuccess (data字段说明) {int}      out_manual_count		人工出款人数
 * @apiSuccess (data字段说明) {float}    out_offer				给予优惠
 * @apiSuccess (data字段说明) {int}      out_offer_count			给予优惠人数
 * @apiSuccess (data字段说明) {float}    out_commission     	 	给予返水
 * @apiSuccess (data字段说明) {int}      out_commission_count	给予返水人数
 * @apiSuccess (data字段说明) {float}    out_first				首次出款
 * @apiSuccess (data字段说明) {int}      out_first_count			首次出款人数
 * @apiSuccess (data字段说明) {float}    total					总计
 * @apiSuccess (data字段说明) {float}    balance					平台实际盈亏
 * @apiSuccess (data字段说明) {float}    remain					可用金额
 * @apiSuccess (data字段说明) {float}    no_settlement			未结算金额
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "保存数据成功",
 *    "code": 200,
 *    "data": {
 *        "in_total": "",
 *        "in_total_count": "",
 *        "in_company": "",
 *        "in_company_count": "",
 *        "in_online": "",
 *        "in_online_count": "",
 *        "in_manual": "",
 *        "in_manual_count": "",
 *        "in_deduction": "",
 *        "in_deduction_count": "",
 *        "in_first": "",
 *        "in_first_count": "",
 *        "out_total": "",
 *        "out_total_count": "",
 *        "out_online": "",
 *        "out_online_count": "",
 *        "out_manual": "",
 *        "out_manual_count": "",
 *        "out_offer": "",
 *        "out_offer_count": "",
 *        "out_commission": "",
 *        "out_commission_count": "",
 *        "out_first": "",
 *        "out_first_count": "",
 *        "total": "",
 *        "balance": "",
 *        "remain": "",
 *        "no_settlement": "",
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *SystemStatisticsController) Counts(ctx *Context) {
	row, err := systemStatistics.Counts(ctx)
	if err != nil {
		responseFailure(ctx, "", err.Error())
		return
	}

	responseSuccess(ctx, "", row)
}

/**
 * @api {get} admin/api/auth/v1/system_statistics/first_charge 首次充值明细
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: lyndon</span><br/><br/>
 * <strong>首次充值明细</strong><br />
 * 业务描述: 首次充值明细</br>
 * @apiVersion 1.0.0
 * @apiName     FirstCharge
 * @apiGroup    report
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int}    page            页数
 * @apiParam (客户端请求参数) {int}   page_size       每页记录数
 * @apiParam (客户端请求参数) {string}   ymd_start 日期/开始
 * @apiParam (客户端请求参数) {string}   ymd_end 日期/结束
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
 * @apiSuccess (data-rows每个子对象字段说明) {int}		  id 					记录编号
 * @apiSuccess (data-rows每个子对象字段说明) {int}        user_id                 用户ID
 * @apiSuccess (data-rows每个子对象字段说明) {string}        user_name                 用户名称
 * @apiSuccess (data-rows每个子对象字段说明) {float}      amount                  充值金额
 * @apiSuccess (data-rows每个子对象字段说明) {string}     order_id                充值订单
 * @apiSuccess (data-rows每个子对象字段说明) {int}        charge_type_id          充值方式id
 * @apiSuccess (data-rows每个子对象字段说明) {string}     card_number             卡号
 * @apiSuccess (data-rows每个子对象字段说明) {string}     bank_address            开户银行地址或支付二维码
 * @apiSuccess (data-rows每个子对象字段说明) {int}        created                 添加时间
 * @apiSuccess (data-rows每个子对象字段说明) {int}        state                   公司入款：0 待审核，1 成功，2 失败。线上支付：0待处理，1成功，2失败，3进行中,4退款，5取消，6强制入款
 * @apiSuccess (data-rows每个子对象字段说明) {string}     charge_type             充值类型
 * @apiSuccess (data-rows每个子对象字段说明) {string}     ip                      充值IP
 * @apiSuccess (data-rows每个子对象字段说明) {int}        platform_id             充值platform
 * @apiSuccess (data-rows每个子对象字段说明) {string}     real_name               真实姓名
 * @apiSuccess (data-rows每个子对象字段说明) {int}        bank_type_id            银行转账类型
 * @apiSuccess (data-rows每个子对象字段说明) {int}        bank_charge_time        银行转账时间
 * @apiSuccess (data-rows每个子对象字段说明) {int}        credential_id           第三方支付记录ID
 * @apiSuccess (data-rows每个子对象字段说明) {string}     operator                操作者
 * @apiSuccess (data-rows每个子对象字段说明) {int}        is_tppay                是否第三方支付 0为否；1为是
 * @apiSuccess (data-rows每个子对象字段说明) {int}        charge_card_id          收款银行卡编号
 * @apiSuccess (data-rows每个子对象字段说明) {string}        charge_card_name          收款银行卡信息
 * @apiSuccess (data-rows每个子对象字段说明) {string}     remark                  备注
 * @apiSuccess (data-rows每个子对象字段说明) {string}     updated_last            最后更新时间
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "数据获取成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "user_id": "",
 *        "user_name": "",
 *        "amount": "",
 *        "order_id": "",
 *        "charge_type_id": "",
 *        "card_number": "",
 *        "bank_address": "",
 *        "created": "",
 *        "state": "",
 *        "charge_type": "",
 *        "ip": "",
 *        "platform_id": "",
 *        "real_name": "",
 *        "bank_type_id": "",
 *        "bank_charge_time": "",
 *        "credential_id": "",
 *        "operator": "",
 *        "is_tppay": "",
 *        "charge_card_id": "",
 *        "charge_card_name": "",
 *        "remark": "",
 *        "updated_last": ""
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *SystemStatisticsController) FirstCharge(ctx *Context) {
	records, err := systemStatistics.CountsFirstCharge(ctx)
	if err != nil {
		responseFailure(ctx, err.Error(), "获取数据失败")
		return
	}
	responseSuccess(ctx, "获取数据成功", records)
}

/**
 * @api {get} admin/api/auth/v1/system_statistics/first_withdraw 首次出款明细
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: lyndon</span><br/><br/>
 * <strong>首次出款明细</strong><br />
 * 业务描述: 首次出款明细</br>
 * @apiVersion 1.0.0
 * @apiName     FirstWithdraw
 * @apiGroup    report
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int}    page            页数
 * @apiParam (客户端请求参数) {int}   page_size       每页记录数
 * @apiParam (客户端请求参数) {string}   ymd_start 日期/开始
 * @apiParam (客户端请求参数) {string}   ymd_end 日期/结束
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
 * @apiSuccess (data-rows每个子对象字段说明) {string}        order_id	订单编号
 * @apiSuccess (data-rows每个子对象字段说明) {int}        user_id                 用户id
 * @apiSuccess (data-rows每个子对象字段说明) {string}        user_name                 用户名称
 * @apiSuccess (data-rows每个子对象字段说明) {float}      amount                  金额 正数为收入，负数为支出
 * @apiSuccess (data-rows每个子对象字段说明) {float}      balance                 变化后的余额
 * @apiSuccess (data-rows每个子对象字段说明) {int}        type                    交易类型
 * @apiSuccess (data-rows每个子对象字段说明) {int}        created                 创建时间
 * @apiSuccess (data-rows每个子对象字段说明) {string}        msg                 流水说明
 * @apiSuccess (data-rows每个子对象字段说明) {float}      charged_amount_old       账变前充值总额
 * @apiSuccess (data-rows每个子对象字段说明) {float}      charged_amount           账变后充值总额
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "数据获取成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "user_id": "",
 *        "user_name": "",
 *        "amount": "",
 *        "balance": "",
 *        "type": "",
 *        "created": "",
 *        "charged_amount_old": "",
 *        "charged_amount": "",
 *        "msg": "",
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *SystemStatisticsController) FirstWithdraw(ctx *Context) {
	records, err := systemStatistics.CountsFirstWithdraw(ctx)
	if err != nil {
		responseFailure(ctx, err.Error(), "获取数据失败")
		return
	}
	responseSuccess(ctx, "获取数据成功", records)
}

/**
 * @api {get} admin/api/auth/v1/system_statistics/back_water 返水明细
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: lyndon</span><br/><br/>
 * <strong>返水明细</strong><br />
 * 业务描述: 返水明细</br>
 * @apiVersion 1.0.0
 * @apiName     backWater
 * @apiGroup    report
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int}    page            页数
 * @apiParam (客户端请求参数) {int}   page_size       每页记录数
 * @apiParam (客户端请求参数) {string}   ymd_start 日期/开始
 * @apiParam (客户端请求参数) {string}   ymd_end 日期/结束
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
 * @apiSuccess (data-rows每个子对象字段说明) {string}        order_id	订单编号
 * @apiSuccess (data-rows每个子对象字段说明) {int}        user_id                 用户id
 * @apiSuccess (data-rows每个子对象字段说明) {string}        user_name                 用户名称
 * @apiSuccess (data-rows每个子对象字段说明) {float}      amount                  金额 正数为收入，负数为支出
 * @apiSuccess (data-rows每个子对象字段说明) {float}      balance                 变化后的余额
 * @apiSuccess (data-rows每个子对象字段说明) {int}        type                    交易类型
 * @apiSuccess (data-rows每个子对象字段说明) {int}        created                 创建时间
 * @apiSuccess (data-rows每个子对象字段说明) {string}        msg                 流水说明
 * @apiSuccess (data-rows每个子对象字段说明) {float}      charged_amount_old       账变前充值总额
 * @apiSuccess (data-rows每个子对象字段说明) {float}      charged_amount           账变后充值总额
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "数据获取成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "user_id": "",
 *        "user_name": "",
 *        "amount": "",
 *        "balance": "",
 *        "type": "",
 *        "created": "",
 *        "charged_amount_old": "",
 *        "charged_amount": "",
 *        "msg": ""
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *SystemStatisticsController) BackWater(ctx *Context) {
	records, err := systemStatistics.CountsBackWater(ctx)
	if err != nil {
		responseFailure(ctx, err.Error(), "获取数据失败")
		return
	}
	responseSuccess(ctx, "获取数据成功", records)
}
