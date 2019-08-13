package validations

import (
	"qpgame/common/mvc"
)

type UsersValidation struct{}

// 添加/修改动作数据验证
func (self UsersValidation) Validate(ctx *Context) (string, bool) {
	return mvc.NewValidation(ctx).
		NotNull("phone", "手机号不能为空").
		NotNull("password", "密码不能为空").
		NotNull("name", "用户姓名不能为空").
		InNumbers("mobile_type", "手机类型必须在可选范围之内", &[]int64{0, 1, 2}).
		InNumbers("sex", "性别必须在可选范围之内", &[]int64{0, 1, 2}).
		InNumbers("status", "用户状态必须在可选范围之内", &statusTypes).
		InNumbers("is_dummy", "是否虚拟用户必须在可选范围之内", &statusTypes).
		IsNumeric("parent_id", "上级代理用户Id必须为数字").
		Validate()
}

// 对于修改密码的校验
func (self UsersValidation) CheckPass(ctx *Context) (string, bool) {
	return mvc.NewValidation(ctx).
		IsNumeric("id", "必须提供用户编号").
		StringLength("old_password", "旧密码长度必须在6-20之间", 6, 20).
		StringLength("password", "密码长度必须在6-20之间", 6, 20).
		StringLength("re_password", "确认密码的长度必须在6-20之间", 6, 20).
		Equals("password", "re_password", "两次输入的密码必须一致").
		Validate()
}

// 对于修改安全密码的校验
func (self UsersValidation) CheckSafePass(ctx *Context) (string, bool) {
	return mvc.NewValidation(ctx).
		IsNumeric("id", "必须提供用户编号").
		StringLength("old_password", "旧密码长度必须为4位", 4, 4).
		StringLength("password", "密码长度必须为4位", 4, 4).
		StringLength("re_password", "确认密码的长度必须为4位", 4, 4).
		Equals("password", "re_password", "两次输入的密码必须一致").
		Validate()
}
