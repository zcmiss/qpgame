package validations

import (
	"qpgame/common/mvc"
)

type SilverMerchantBankCardsValidation struct{}

// 添加/修改动作数据验证
func (self SilverMerchantBankCardsValidation) Validate(ctx *Context) (string, bool) {
	return mvc.NewValidation(ctx).
		NotNull("merchant_id", "银商编号不能为空").
		NotNull("card_number", "号号不能为空").
		NotNull("address", "银行卡地址不能为空").
		NotNull("bank_name", "银行名称不能为空").
		NotNull("name", "姓名不能为空").
		Validate()
}