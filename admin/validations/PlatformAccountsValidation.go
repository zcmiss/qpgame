package validations

import (
	"qpgame/common/mvc"
)

type PlatformAccountsValidation struct{}

// 添加/修改动作数据验证
func (self PlatformAccountsValidation) Validate(ctx *Context) (string, bool) {
	return mvc.NewValidation(ctx).
		IsNumeric("plat_id", "平台编号必须为数字").
		IsNumeric("user_id", "用户编号必须为数字").
		StringLength("username", "用户名长度在5-20之间", 5, 20).
		StringLength("password", "密码长度在5-20之间", 5, 20).
		Validate()
}
