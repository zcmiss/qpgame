package validations

import (
	"qpgame/common/mvc"
)

type ConfigsValidation struct{}

// 添加/修改动作数据验证
func (self ConfigsValidation) Validate(ctx *Context) (string, bool) {
	return mvc.NewValidation(ctx).
		StringLength("name", "设置项名称长度在2-100之间", 2, 100).
		NotNull("value", "设置参数值不能为空").
		StringLength("mark", "记号说明(该字段只作数据库字段应用说明,不作程序使用)长度在2-100之间", 2, 100).
		Validate()
}

//系统配置
func (self ConfigsValidation) ValidateSets(ctx *Context) (string, bool) {
	return mvc.NewValidation(ctx).
		IsNumeric("withdraw_min_money", "提现最低金额必须是有效的数字").
		IsNumeric("withdraw_max_money", "提现最大金额必须是有效的数字").
		IsNumeric("clear_dm_tx_limit", "清除提现打码限制阀值必须是有效的数字").
		IsNumeric("withdraw_day_limited", "每天提款次数制制必须是数字").
		//IsNumeric("register_config", "注册配置必须").
		//StringLength("tuiguang_web_url", "代理推广地址后缀长度必须介于2-100", 2, 100).
		IsNumeric("sign_reward", "签到奖励必须是有效的数字").
		IsNumeric("register_number_ip", "单IP允许最多注册人数必须是有效的数字").
		IsNumeric("sign_award_switch", "是否开启签到奖励").
		IsNumeric("bind_phone_award_switch", "是否开启绑定手机奖励").
		IsNumeric("reward_bind", "绑定手机奖励金额必须是有效的数字").
		Validate()
}

//proxy_charge / 代理充值
//proxy_change_logo / 充值logo
//change_accounts / (url/name/account => 头像地址/名称/账号)
//info /
////充值配置
func (self ConfigsValidation) ValidateCharge(ctx *Context) (string, bool) {
	return "", true
}

////客服配置
func (self ConfigsValidation) ValidateService(ctx *Context) (string, bool) {
	return "", true
}

////常见问题
func (self ConfigsValidation) ValidateFaq(ctx *Context) (string, bool) {
	return "", true
}
