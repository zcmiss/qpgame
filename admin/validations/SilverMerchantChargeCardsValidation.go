package validations

import (
	"qpgame/common/mvc"
)

type SilverMerchantChargeCardsValidation struct{}

// 添加/修改动作数据验证
func (self SilverMerchantChargeCardsValidation) Validate(ctx *Context) (string, bool) {
	return mvc.NewValidation(ctx).
		NotNull("name", "银行名称不能为空").
		Validate()
}
