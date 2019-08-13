package validations

import (
	"qpgame/common/mvc"
)

type UserBankCardsValidation struct{}

// 添加/修改动作数据验证
func (self UserBankCardsValidation) Validate(ctx *Context) (string, bool) {
	return mvc.NewValidation(ctx).
		NotNull("user_id", "字段 user_id 不能为空").
		StringLength("card_number", "银行卡号长度在2-50之间", 2, 50).
		StringLength("address", "银行卡地址长度在2-100之间", 2, 100).
		StringLength("bank_name", "银行名称长度在2-50之间", 2, 50).
		StringLength("name", "姓名长度在2-50之间", 2, 50).
		IsNumeric("status", "状态必须为数字").
		StringLength("remark", "备注长度在2-100之间", 2, 100).
		Validate()
}
