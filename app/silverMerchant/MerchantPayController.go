package silverMerchant

import (
	"github.com/kataras/iris"
)

type MerchantPayController struct {
	platform string
	ctx      iris.Context
}

//构造函数
func NewMerchantPayController(ctx iris.Context) *MerchantPayController {
	obj := new(MerchantPayController)
	obj.platform = ctx.Params().Get("platform")
	obj.ctx = ctx
	return obj
}
