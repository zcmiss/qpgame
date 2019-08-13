package controllers

type SilverMerchantReportsController struct{}

/**
 * @api {get} admin/api/auth/v1/silver_merchant_reports 银商报表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: lyndon</span><br/><br/>
 * <strong>银商报表</strong><br />
 * 业务描述: 银商报表</br>
 * @apiVersion 1.0.0
 * @apiName     indexSilverMerchantReports
 * @apiGroup    silver_merchant
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int}    page           页数
 * @apiParam (客户端请求参数) {int}    page_size      每页记录数
 * @apiParam (客户端请求参数) {string} created_start  时间/开始
 * @apiParam (客户端请求参数) {string} created_end    时间/结束
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
 * @apiSuccess (data-rows每个子对象字段说明) {float}	 amount_in     额度充值总金额
 * @apiSuccess (data-rows每个子对象字段说明) {float}   amount_out    会员充值总金额
 * @apiSuccess (data-rows每个子对象字段说明) {int}   count_in        额度充值银商个数
 * @apiSuccess (data-rows每个子对象字段说明) {int}   count_out       会员充值会员个数
 * @apiSuccess (data-rows每个子对象字段说明) {string}  ymd           日期
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "获取数据成功",
 *    "code": 200,
 *    "data": {
 *        "rows": [
 *            {
 *                "amount_in": "232.000",
 *                "amount_out": "32.000",
 *                "count_in": "23",
 *                "count_out": "123",
 *                "ymd": "2015-12-13",
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
func (self *SilverMerchantReportsController) Index(ctx *Context) {
	records, err := silverMerchantCapitalFlows.GetReports(ctx)
	if err != nil {
		responseFailure(ctx, err.Error(), "获取数据失败")
		return
	}
	responseSuccess(ctx, "获取数据成功", records)
}
