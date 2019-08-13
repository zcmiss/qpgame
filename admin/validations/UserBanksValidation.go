package validations

import (
	"qpgame/common/mvc"
)

type UserBanksValidation struct{}

// 添加/修改动作数据验证
func (self UserBanksValidation) Validate(ctx *Context) (string, bool) {
	return mvc.NewValidation(ctx).
		NotNull("name", "字段 name 不能为空").
		StringLength("logo", "银行图标长度在2-255之间", 2, 255).
		Validate()
}
