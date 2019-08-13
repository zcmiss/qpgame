package validations

import (
	"qpgame/common/mvc"
)

type SilverMerchantUsersValidation struct{}

// 添加/修改动作数据验证
func (self SilverMerchantUsersValidation) Validate(ctx *Context) (string, bool) {
	return mvc.NewValidation(ctx).
		NotNull("user_id", "用户编号不能为空").
		NotNull("account", "银商账号不能为空").
		NotNull("merchant_name", "商户名称不能为空").
		IsNumeric("merchant_level", "银商等级必须为数字").
		IsDecimal("donate_rate", "赠送比例必须为数字").
		Validate()
}

// 对于修改密码的校验
func (self SilverMerchantUsersValidation) CheckPass(ctx *Context) (string, bool) {
	return mvc.NewValidation(ctx).
		IsNumeric("id", "必须提供用户编号").
		StringLength("old_password", "旧密码长度必须在6-20之间", 6, 20).
		StringLength("password", "密码长度必须在6-20之间", 6, 20).
		StringLength("confirm_password", "确认密码的长度必须在6-20之间", 6, 20).
		Equals("password", "confirm_password", "两次输入的密码必须一致").
		Validate()
}
