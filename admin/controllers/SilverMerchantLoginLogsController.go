package controllers

import (
	"qpgame/admin/models"
	"qpgame/admin/validations"
)

var silverMerchantLoginLogs = models.SilverMerchantLoginLogs{}
var silverMerchantLoginLogsValidation = validations.SilverMerchantLoginLogsValidation{}

type SilverMerchantLoginLogsController struct{}

/**
 * @api {get} admin/api/auth/v1/silver_merchant_login_logs 银商登录记录
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: lyndon</span><br/><br/>
 * <strong>银商登录记录</strong><br />
 * 业务描述: 银商登录记录</br>
 * @apiVersion 1.0.0
 * @apiName     indexSilverMerchantLoginLogs
 * @apiGroup    silver_merchant
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int}     page            页数
 * @apiParam (客户端请求参数) {int}     page_size       每页记录数
 * @apiParam (客户端请求参数) {int} 	  merchant_id     银商编号
 * @apiParam (客户端请求参数) {string} 	ip               IP
 * @apiParam (客户端请求参数) {string} 	login_city       所在城市
 * @apiParam (客户端请求参数) {string} 	login_time_start 添加时间/开始
 * @apiParam (客户端请求参数) {string} 	login_time_end   添加时间/结束
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
 * @apiSuccess (data-rows每个子对象字段说明) {int}		  id 			    记录编号
 * @apiSuccess (data-rows每个子对象字段说明) {int}		  merchant_id 		银商编号
 * @apiSuccess (data-rows每个子对象字段说明) {string}   merchant_name 		银商名称
 * @apiSuccess (data-rows每个子对象字段说明) {string}   ip          IP
 * @apiSuccess (data-rows每个子对象字段说明) {string}   login_city  所在城市
 * @apiSuccess (data-rows每个子对象字段说明) {string}   login_time  登录时间
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "获取数据成功",
 *    "code": 200,
 *    "data": {
 *        "rows": [
 *            {
 *                "id": "44",
 *                "ip": "127.0.0.1",
 *                "login_city": "中国上海",
 *                "login_time": "",
 *                "merchant_id": "1",
 *                "merchant_name": "XXXXX"
 *            }
 *        ],
 *        "page": 1,
 *        "page_count": 3,
 *        "total_rows": 43,
 *        "page_size": 20
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 0
 *}
 */
func (self *SilverMerchantLoginLogsController) Index(ctx *Context) {
	index(ctx, &silverMerchantLoginLogs)
}
