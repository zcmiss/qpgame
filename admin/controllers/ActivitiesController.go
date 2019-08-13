package controllers

import (
	"qpgame/admin/models"
	"qpgame/admin/validations"
)

var activities = models.Activities{}                          //模型
var activitiesValidation = validations.ActivitiesValidation{} //校验器

type ActivitiesController struct{}

/**
* @api {get} admin/api/auth/v1/activities 活动管理列表
* @apiDescription
* <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
* <strong>活动管理列表</strong><br />
* 业务描述: 活动管理列表</br>
* @apiVersion 1.0.0
* @apiName     indexActivities
* @apiGroup    content
* @apiPermission PC客户端
* @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
* @apiHeaderExample {json} 请求头示例
* {
*      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
* }
* @apiParam (客户端请求参数) {int} 	page			页数
* @apiParam (客户端请求参数) {int} 	page_size		每页记录数
* @apiParam (客户端请求参数) {int} 	activity_class_id 活动分类
* @apiParam (客户端请求参数) {string} 	title 标题
* @apiParam (客户端请求参数) {string} 	sub_title 子标题
* @apiParam (客户端请求参数) {int} 	status 状态
* @apiParam (客户端请求参数) {string} 	time_start 活动开始时间
* @apiParam (客户端请求参数) {string} 	time_end 活动结束时间
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
* @apiSuccess (data-rows每个子对象字段说明) {int}		id 						记录编号
* @apiSuccess (data-rows每个子对象字段说明) {int}        activity_class_id       活动分类编号
* @apiSuccess (data-rows每个子对象字段说明) {int}        activity_class_name     活动分类名称
* @apiSuccess (data-rows每个子对象字段说明) {string}     title                   标题
* @apiSuccess (data-rows每个子对象字段说明) {string}     sub_title               子标题
* @apiSuccess (data-rows每个子对象字段说明) {string}     icon                    活动图片
* @apiSuccess (data-rows每个子对象字段说明) {int}        created                 创建时间
* @apiSuccess (data-rows每个子对象字段说明) {int}        time_start              开始时间
* @apiSuccess (data-rows每个子对象字段说明) {int}        time_end                结束时间
* @apiSuccess (data-rows每个子对象字段说明) {float}      money                   活动奖励
* @apiSuccess (data-rows每个子对象字段说明) {int}        type                    类型,0:其他，1注册,2:充值
* @apiSuccess (data-rows每个子对象字段说明) {int}        total_ip_limit          限制IP领取总数
* @apiSuccess (data-rows每个子对象字段说明) {int}        day_ip_limit            限制IP当日领取总数
* @apiSuccess (data-rows每个子对象字段说明) {int}        is_repeat               是否可以重复领取(0为否,1为是)
* @apiSuccess (data-rows每个子对象字段说明) {int}        status                  状态
* @apiSuccess (data-rows每个子对象字段说明) {int}        updated                 最後修改時間
* @apiSuccess (data-rows每个子对象字段说明) {int}        is_home_show            是否首页弹出显示(0为否,1为是)
*
* @apiSuccessExample {json} 响应结果
* {
*    "clientMsg": "数据获取成功",
*    "code": 200,
*    "data": {
*        "id": "",
*        "title": "",
 *       "icon": "",
*        "sub_title": "",
*        "content": "",
*        "created": "",
*        "time_start": "",
*        "time_end": "",
*        "money": "",
*        "type": "",
*        "total_ip_limit": "",
*        "day_ip_limit": "",
*        "is_repeat": "",
*        "status": "",
*        "updated": "",
*        "activity_class_id": "",
*        "is_home_show": ""
*    },
*    "internalMsg": "",
*    "timeConsumed": 168
* }
*/
func (self *ActivitiesController) Index(ctx *Context) {
	index(ctx, &activities)
}

/**
 * @api {post} admin/api/auth/v1/activities/add	活动管理添加
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>添加/修改活动管理</strong><br />
 * 业务描述: 添加/修改活动管理 <br />
 * <strong><span style="color: red">注意: </span></strong><br />
 * <span style="color:red">修改操作API不再单独列出, 请参考以下</span><br />
 * <span style="color:red">添加: /admin/api/auth/v1/activities/add </span> &nbsp;&nbsp; <br />
 * <span style="color:red">修改: /admin/api/auth/v1/activities/update </span>
 * @apiVersion 1.0.0
 * @apiName     saveActivities
 * @apiGroup    content
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }

 * @apiParam (客户端请求参数) {int} 	id    					记录编号,仅修改操作时(即接口: /admin/api/auth/v1/activities/update)需要<br />
 *															添加操作(即接口: /admin/api/auth/v1/activities/add)不需要此参数<br />
 * 															* 如果提供此编号, 则视为修改记录
 * @apiParam (客户端请求参数) {string}   	title                   标题
 * @apiParam (客户端请求参数) {string}   	sub_title               子标题
 * @apiParam (客户端请求参数) {string}   	icon                    活动图片
 * @apiParam (客户端请求参数) {string}   	content                 内容
 * @apiParam (客户端请求参数) {int}      	time_start              开始时间
 * @apiParam (客户端请求参数) {int}      	time_end                结束时间
 * @apiParam (客户端请求参数) {int}      	type                    类型,0:其他，1注册,2:充值
 * @apiParam (客户端请求参数) {int}      	status                  状态
 * @apiParam (客户端请求参数) {string}   	link                    链接
 * @apiParam (客户端请求参数) {int}      	activity_class_id       活动分类（关联活动分类表）
 * @apiParam (客户端请求参数) {int}      	is_home_show            是否首页弹出显示(0为否,1为是
 * @apiParam (客户端请求参数) {int}      	total_ip_limit          限制IP领取总数
 * @apiParam (客户端请求参数) {int}      	day_ip_limit            限制IP当日领取总数
 * @apiParam (客户端请求参数) {int}      	is_repeat               是否可以重复领取(0为否,1为是)
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
func (self *ActivitiesController) Save(ctx *Context) {
	save(ctx, &activities, &activitiesValidation)
}

/**
 * @api {get} admin/api/auth/v1/activities/view 				活动管理详情
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>活动管理详情</strong><br />
 * 业务描述: 活动管理详情</br>
 * @apiVersion 1.0.0
 * @apiName     viewActivities
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
 * @apiSuccess (data字段说明) {string}     title                   标题
 * @apiSuccess (data字段说明) {string}     sub_title               子标题
 * @apiSuccess (data字段说明) {string}     icon                    活动图片
 * @apiSuccess (data字段说明) {string}     content                 内容
 * @apiSuccess (data字段说明) {int}        created                 创建时间
 * @apiSuccess (data字段说明) {int}        time_start              开始时间
 * @apiSuccess (data字段说明) {int}        time_end                结束时间
 * @apiSuccess (data字段说明) {int}        type                    类型,0:其他，1注册,2:充值
 * @apiSuccess (data字段说明) {float}      money                   活动奖励
 * @apiSuccess (data字段说明) {int}        status                  状态
 * @apiSuccess (data字段说明) {int}        updated                 更新时间
 * @apiSuccess (data字段说明) {int}        activity_class_id       活动分类（关联活动分类表）
 * @apiSuccess (data字段说明) {int}        is_home_show            是否首页弹出显示(0为否,1为是
 * @apiSuccess (data字段说明) {int}        total_ip_limit          限制IP领取总数
 * @apiSuccess (data字段说明) {int}      	  day_ip_limit            限制IP当日领取总数
 * @apiSuccess (data字段说明) {int}        is_repeat               是否可以重复领取(0为否,1为是)
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "保存数据成功",
 *    "code": 200,
 *    "data": {
 *        "id": "",
 *        "title": "",
 *        "sub_title": ""
 *        "icon": "",,
 *        "content": "",
 *        "created": "",
 *        "time_start": "",
 *        "time_end": "",
 *        "type": "",
 *        "status": "",
 *        "updated": "",
 *        "activity_class_id": "",
 *        "is_home_show": "",
 *        "total_ip_limit": "",
 *        "day_ip_limit": "",
 *        "is_repeat": ""
 * 	  },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *ActivitiesController) View(ctx *Context) {
	view(ctx, &activities)
}

/**
 * @api {get} admin/api/auth/v1/activities/delete 活动管理删除
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>活动管理删除</strong><br />
 * 业务描述: 活动管理删除</br>
 * @apiVersion 1.0.0
 * @apiName     deleteActivities
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
func (self *ActivitiesController) Delete(ctx *Context) {
	remove(ctx, &activities)
}
