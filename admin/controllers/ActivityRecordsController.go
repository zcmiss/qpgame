package controllers

import (
	"qpgame/admin/models"
)

var activityRecords = models.ActivityRecords{} //模型

type ActivityRecordsController struct{}

/**
 * @api {get} admin/api/auth/v1/activity_records 活动申請记录
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>活动申请记录列表</strong><br />
 * 业务描述: 活动申请记录列表</br>
 * @apiVersion 1.0.0
 * @apiName     indexActivityRecords
 * @apiGroup    content
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int} 	page			页数
 * @apiParam (客户端请求参数) {int} 	page_size		每页记录数
 * @apiParam (客户端请求参数) {int} 	user_id 用户编号
 * @apiParam (客户端请求参数) {int} 	activity_id 活动/下拉框/拿出全部活信息
 * @apiParam (客户端请求参数) {string} 	start_time 申请时间/开始
 * @apiParam (客户端请求参数) {string} 	start_time 申请时间/结束
 * @apiParam (客户端请求参数) {int} 	state 处理状态
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
 * @apiSuccess (data-rows每个子对象字段说明) {int}        user_id                 用户编号
 * @apiSuccess (data-rows每个子对象字段说明) {string}     user_name            	用户名称
 * @apiSuccess (data-rows每个子对象字段说明) {int}        activity_id             活动编号
 * @apiSuccess (data-rows每个子对象字段说明) {int}        activity_name           活动名称
 * @apiSuccess (data-rows每个子对象字段说明) {string}     remark                  备注
 * @apiSuccess (data-rows每个子对象字段说明) {int}        state                   是否处理(1:已处理,0:未处理)
 * @apiSuccess (data-rows每个子对象字段说明) {string}     operator                操作者
 * @apiSuccess (data-rows每个子对象字段说明) {int}        applied                 申请时间
 * @apiSuccess (data-rows每个子对象字段说明) {int}        created                 创建时间
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "数据获取成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "user_id": "",
 *        "user_name": "",
 *        "activity_id": "",
 *        "remark": "",
 *        "state": "",
 *        "operator": "",
 *        "applied": "",
 *        "created": "",
 *        "updated": ""
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *ActivityRecordsController) Index(ctx *Context) {
	index(ctx, &activityRecords)
}

/**
 * @api {get} admin/api/auth/v1/activity_records/view 				活动申请记录详情
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>活动记录详情</strong><br />
 * 业务描述: 活动记录详情</br>
 * @apiVersion 1.0.0
 * @apiName     viewActivityRecords
 * @apiGroup    content
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
 * @apiSuccess (data字段说明) {int}        user_id                 用户编号
 * @apiSuccess (data字段说明) {int}        activity_id             活动编号
 * @apiSuccess (data字段说明) {string}     remark                  备注
 * @apiSuccess (data字段说明) {int}        state                   是否处理(1:已处理,0:未处理)
 * @apiSuccess (data字段说明) {string}     operator                操作者
 * @apiSuccess (data字段说明) {int}        applied                 申请时间
 * @apiSuccess (data字段说明) {int}        created                 创建时间
 * @apiSuccess (data字段说明) {int}        updated                 更新时间
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "保存数据成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "user_id": "",
 *        "activity_id": "",
 *        "remark": "",
 *        "state": "",
 *        "operator": "",
 *        "applied": "",
 *        "created": "",
 *        "updated": ""
 * 	  },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *ActivityRecordsController) View(ctx *Context) {
	view(ctx, &activityRecords)
}

/**
 * @api {get} admin/api/auth/v1/activity_records/delete 活动申请记录删除
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>删除活动申请记录</strong><br />
 * 业务描述: 删除活动申请记录</br>
 * @apiVersion 1.0.0
 * @apiName     deleteActivityRecords
 * @apiGroup    content
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
func (self *ActivityRecordsController) Delete(ctx *Context) {
	remove(ctx, &activityRecords)
}
