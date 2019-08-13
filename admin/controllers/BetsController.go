package controllers

import (
	"qpgame/admin/models"
)

var bets = models.Bets{} //模型

type BetsController struct{}

/**
 * @api {get} admin/api/auth/v1/bets 注单管理列表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>注单管理列表</strong><br />
 * 业务描述: 注单管理列表</br>
 * @apiVersion 1.0.0
 * @apiName     indexBets
 * @apiGroup    game
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int} 	page			页数
 * @apiParam (客户端请求参数) {int} 	page_size		每页记录数
 * @apiParam (客户端请求参数) {int} 	platform_id		游戏平台
 * @apiParam (客户端请求参数) {int} 	game_id		游戏/平台游戏
 * @apiParam (客户端请求参数) {int} 	user_id		用户编号
 * @apiParam (客户端请求参数) {string} 	user_id		订单编号
 * @apiParam (客户端请求参数) {string} 	time_start 投注时间/开始
 * @apiParam (客户端请求参数) {string} 	time_end 投注时间/结束
 * @apiParam (客户端请求参数) {string} 	accountname 第三方游戏平台账号
 * @apiParam (客户端请求参数) {int} 	rebate_state 	洗码状态,0:未洗,1:已洗
 * @apiParam (客户端请求参数) {int} 	dama_state 	打码状态,0:未更新,1:已更新
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
 * @apiSuccess (data-rows每个子对象字段说明) {float}      amount                  下注金额
 * @apiSuccess (data-rows每个子对象字段说明) {float}      amount_all              总下注金额
 * @apiSuccess (data-rows每个子对象字段说明) {float}      amount_platform         第三方平台抽水
 * @apiSuccess (data-rows每个子对象字段说明) {int}        created                 游戏时间
 * @apiSuccess (data-rows每个子对象字段说明) {int}        game_id                 游戏ID,关联TPPlayType的id
 * @apiSuccess (data-rows每个子对象字段说明) {string}        game_name                 游戏名称
 * @apiSuccess (data-rows每个子对象字段说明) {string}     order_id                订单编号
 * @apiSuccess (data-rows每个子对象字段说明) {int}        platform_id             第三方平台ID
 * @apiSuccess (data-rows每个子对象字段说明) {string}        platform_name             第三方平台名称
 * @apiSuccess (data-rows每个子对象字段说明) {float}      reward                  中奖金额
 * @apiSuccess (data-rows每个子对象字段说明) {int}        user_id                 用户编号
 * @apiSuccess (data-rows每个子对象字段说明) {string}        user_name                 用户名称
 * @apiSuccess (data-rows每个子对象字段说明) {int}        rebate_state			洗码状态
 * @apiSuccess (data-rows每个子对象字段说明) {string}        rebate_state_name	洗码状态名称
 * @apiSuccess (data-rows每个子对象字段说明) {int}        dama_state				打码状态
 * @apiSuccess (data-rows每个子对象字段说明) {string}        dama_state_name		打码状态名称
 * @apiSuccess (data-rows每个子对象字段说明) {string}        account_name		第三方平台游戏账号
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "数据获取成功",
 *    "code": 200,
 *    "data": {
 *        "amount": "",
 *        "amount_all": "",
 *        "amount_platform": "",
 *        "created": "",
 *        "game_id": "",
 *        "game_name": "",
 *        "id": "",
 *        "order_id": "",
 *        "platform_id": "",
 *        "platform_name": "",
 *        "reward": "",
 *        "user_id": "",
 *        "user_name": "",
 *        "gt": "",
 *        "rebate_state": "",
 *        "rebate_state_name": "",
 *        "dama_state": "",
 *        "dama_state_name": "",
 *        "account_name": "",
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *BetsController) Index(ctx *Context) {
	index(ctx, &bets)
}

/**
 * @api {get} admin/api/auth/v1/bets/get_calculated_total 注单统计
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: xiaoye</span><br/><br/>
 * <strong>注单统计</strong><br />
 * 业务描述: 注单统计</br>
 * @apiVersion 1.0.0
 * @apiName     getCalculatedTotalBets
 * @apiGroup    game
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int} 	platform_id		游戏平台
 * @apiParam (客户端请求参数) {int} 	game_id		游戏/平台游戏
 * @apiParam (客户端请求参数) {int} 	user_id		用户编号
 * @apiParam (客户端请求参数) {string} 	user_id		订单编号
 * @apiParam (客户端请求参数) {string} 	time_start 投注时间/开始
 * @apiParam (客户端请求参数) {string} 	time_end 投注时间/结束
 * @apiParam (客户端请求参数) {string} 	accountname 第三方游戏平台账号
 * @apiParam (客户端请求参数) {int} 	rebate_state 	洗码状态,0:未洗,1:已洗
 * @apiParam (客户端请求参数) {int} 	dama_state 	打码状态,0:未更新,1:已更新
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
 *
 * @apiSuccess (data-rows每个子对象字段说明) {float}        total_bet         下注总额
 * @apiSuccess (data-rows每个子对象字段说明) {float}        total_prize			中奖总额
 * @apiSuccess (data-rows每个子对象字段说明) {float}        platform_loss	平台盈利
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "数据获取成功",
 *    "code": 200,
 *    "data": {
 *        "total_bet": "",
 *        "total_prize": "",
 *        "platform_loss": "",
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *BetsController) GetCalculatedTotal(ctx *Context) {
	row, err := bets.GetCalculatedTotal(ctx)
	if err != nil {
		responseFailure(ctx, "", err.Error())
		return
	}

	responseSuccess(ctx, "", row)
}

/**
 * @api {get} admin/api/auth/v1/bets/view 				注单管理详情
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>注单管理详情</strong><br />
 * 业务描述: 注单管理详情</br>
 * @apiVersion 1.0.0
 * @apiName     viewBets
 * @apiGroup    game
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
 * @apiSuccess (data字段说明) {float}      amount                  下注金额
 * @apiSuccess (data字段说明) {float}      amount_all              总下注金额
 * @apiSuccess (data字段说明) {float}      amount_platform         第三方平台抽水
 * @apiSuccess (data字段说明) {int}        created                 游戏时间
 * @apiSuccess (data字段说明) {int}        game_id                 游戏ID,关联TPPlayType的id
 * @apiSuccess (data字段说明) {string}     order_id                订单编号
 * @apiSuccess (data字段说明) {int}        platform_id             第三方平台ID
 * @apiSuccess (data字段说明) {float}      reward                  中奖金额
 * @apiSuccess (data字段说明) {int}        user_id                 用户ID
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "保存数据成功",
 *    "code": 200,
 *    "data": {
 *        "amount": "",
 *        "amount_all": "",
 *        "amount_platform": "",
 *        "created": "",
 *        "game_id": "",
 *        "id": "",
 *        "order_id": "",
 *        "platform_id": "",
 *        "reward": "",
 *        "user_id": ""
 * 	  },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *BetsController) View(ctx *Context) {
	view(ctx, &bets)
}

/**
 * @api {get} admin/api/auth/v1/bets/delete 注单管理删除
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>注单管理删除</strong><br />
 * 业务描述: 注单管理删除</br>
 * @apiVersion 1.0.0
 * @apiName     deleteBets
 * @apiGroup    game
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
func (self *BetsController) Delete(ctx *Context) {
	remove(ctx, &bets)
}
