package controllers

type SilverMerchantUserChargeController struct{}

/**
 * @api {get} admin/api/auth/v1/silver_merchant_user_charge 会员充值记录
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: lyndon</span><br/><br/>
 * <strong>会员充值记录</strong><br />
 * 业务描述: 会员充值记录</br>
 * @apiVersion 1.0.0
 * @apiName     indexSilverMerchantUserCharge
 * @apiGroup    silver_merchant
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 * @apiParam (客户端请求参数) {int}    page           页数
 * @apiParam (客户端请求参数) {int}    page_size      每页记录数
 * @apiParam (客户端请求参数) {int} 	 merchant_id    银商编号
 * @apiParam (客户端请求参数) {string} order_id       订单号
 * @apiParam (客户端请求参数) {string} created_start  添加时间/开始
 * @apiParam (客户端请求参数) {string} created_end    添加时间/结束
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
 * @apiSuccess (data-rows每个子对象字段说明) {int}		 id 				记录编号
 * @apiSuccess (data-rows每个子对象字段说明) {float}	 amount             操作金额
 * @apiSuccess (data-rows每个子对象字段说明) {float}	 balance            变化后的余额
 * @apiSuccess (data-rows每个子对象字段说明) {float}   charged_amount     资金变动之后的额度余额
 * @apiSuccess (data-rows每个子对象字段说明) {float}   charged_amount_old 资金变动之前的额度余额
 * @apiSuccess (data-rows每个子对象字段说明) {string}  created            记录时间
 * @apiSuccess (data-rows每个子对象字段说明) {int}     member_user_id     会员编号
 * @apiSuccess (data-rows每个子对象字段说明) {int}		 merchant_id 		银商编号
 * @apiSuccess (data-rows每个子对象字段说明) {string}	 merchant_name 		银商名称
 * @apiSuccess (data-rows每个子对象字段说明) {string}  order_id           订单号
 * @apiSuccess (data-rows每个子对象字段说明) {string}  user_name          会员名称
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "获取数据成功",
 *    "code": 200,
 *    "data": {
 *        "rows": [
 *            {
 *                "amount": "232.000",
 *                "balance": "32.000",
 *                "charged_amount": "23.000",
 *                "charged_amount_old": "1232.000",
 *                "created": "2015-12-13 17:46:40",
 *                "id": "1",
 *                "member_user_id": "2",
 *                "merchant_id": "1",
 *                "merchant_name": "",
 *                "msg": "转了个帐",
 *                "order_id": "DS2324DSF2342",
 *                "user_name": ""
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
func (self *SilverMerchantUserChargeController) Index(ctx *Context) {
	records, err := silverMerchantCapitalFlows.GetUserChargeRecords(ctx)
	if err != nil {
		responseFailure(ctx, err.Error(), "获取数据失败")
		return
	}
	responseSuccess(ctx, "获取数据成功", records)
}
