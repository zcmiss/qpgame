package silverMerchant

import (
	"github.com/kataras/iris"
	"qpgame/common/utils"
	"qpgame/config"
	"qpgame/models/mainxorm"
	"qpgame/ramcache"
)

type ApiConfigController struct {
	platform string
	ctx      iris.Context
}

//构造函数
func NewApiConfigController(ctx iris.Context) *ApiConfigController {
	obj := new(ApiConfigController)
	obj.platform = ctx.Params().Get("platform")
	obj.ctx = ctx
	return obj
}

/**
 * @api {get} silverMerchant/getConfig 获取系统配置
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人:Singh</span><br/><br/>
 * 获取银商Api接口地址<br>
 * 业务描述:在银商后台启动时，传入域名参数，如果匹配则返回银商api接口地址</br>
 * @apiVersion 1.0.0
 * @apiName     getConfig
 * @apiGroup    silver_merchant
 * @apiSuccess (返回结果)  {int}      code            200正常响应数据
 * @apiSuccess (返回结果)  {string}   clientMsg       提示信息
 * @apiSuccess (返回结果)  {string}   internalMsg     提示信息
 * @apiSuccess (返回结果)  {json}  	  data            返回数据
 * @apiSuccess (返回结果)  {float}   timeConsumed    后台耗时
 * @apiSuccess (data对象字段说明) {string}   silver_merchant_api_address 动态接口地址
 * @apiSuccessExample {json} 响应结果
{
    "clientMsg": "",
    "code": 200,
    "data": {
        "code": "CPTEST",
        "name": "CK棋牌",
        "silver_merchant_api_address": "https://api-silver-merchant.admin-qp.com/CKYX"
    },
    "internalMsg": "",
    "timeConsumed": 31
}
*/
func (cthis *ApiConfigController) GetConfig() {
	ctx := cthis.ctx

	res := map[string]interface{}{
		"code" : "CKYX",
		"name" : "CK棋牌",
		"silver_merchant_api_address" : "https://api-silver-merchant.admin-qp.com/CKYX",
	}
	origin := ctx.GetHeader("Origin")
	mtp, _ := ramcache.MainTablePlatform.Load("platform")
	for _, v := range mtp.([]mainxorm.Platform) {
		if v.SilverMerchantAddress == origin{
			res["code"] = v.Code
			res["name"] = v.Name
			res["silver_merchant_api_address"] = v.SilverMerchantApiAddress + "/" + v.Code
			break
		}
	}
	utils.ResSuccJSON(&ctx, "", "", config.SUCCESSRES, res)
}
