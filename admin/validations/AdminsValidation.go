package validations

import (
	"qpgame/common/mvc"
	"qpgame/common/utils"
	"strings"
)

type AdminsValidation struct{}

// 添加/修改动作数据验证
func (self AdminsValidation) Validate(ctx *Context) (string, bool) {

	post := utils.GetPostData(ctx)
	isCreating := post.GetInt("id") == 0 //是否是创建记录, 只有添加时，或时修改时同时修改密码才进行此项验证
	vali := mvc.NewValidation(ctx)

	if isCreating || strings.Compare(post.Get("password"), "") != 0 { //如果是添加或者是修改时提供的有密码, 则进行密码验证
		vali = vali.StringLength("password", "密码长度必须在6-20之间", 6, 20)
	}

	yesNo := []int64{0, 1}
	return vali.
		IsUserName("name", "管理员名称必须符合规范").
		IsMail("email", "必须输入正确格式的电子邮箱").
		IsNumeric("role_id", "角色编号必须为数字").
		InNumbers("status", "状态必须在可选的范围之内", &yesNo).
		InNumbers("charge_alert", "后台充值提醒值必须在可选范围之内", &yesNo).
		InNumbers("withdraw_alert", "后台充值提醒必须在可选范围之内", &yesNo).
		StringLength("login_ip", "允许登录IP长度在5-500之间", 5, 500).
		IsIpAddress("login_ip", "必须输入正确格式的ip地址,多个IP请用,号隔开").
		InNumbers("permission", "涉及钱的权限必须在可选范围之内", &yesNo).
		InNumbers("force_out", "是否强制退出必须在可选范围之内", &yesNo).
		IsNumeric("manual_max", "最大人工入款金额必须为数字").
		InNumbers("is_opt", "是否OTP验证必须在可选范围之内", &yesNo).
		InNumbers("is_otp_first", "是否第一次OTP验证必须在可选范围之内", &yesNo).
		Validate()
}

// 对于修改密码的校验
func (self AdminsValidation) CheckPass(ctx *Context) (string, bool) {
	return mvc.NewValidation(ctx).
		IsNumeric("id", "必须提供管理员编号").
		StringLength("old_password", "旧密码长度必须在6-20之间", 6, 20).
		StringLength("password", "密码长度必须在6-20之间", 6, 20).
		StringLength("re_password", "确认密码的长度必须在6-20之间", 6, 20).
		Equals("password", "re_password", "两次输入的密码必须一致").
		Validate()
}
