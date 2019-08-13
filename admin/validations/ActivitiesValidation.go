package validations

import (
	"qpgame/common/mvc"
)

type ActivitiesValidation struct{}

// 添加/修改动作数据验证
func (self ActivitiesValidation) Validate(ctx *Context) (string, bool) {
	return mvc.NewValidation(ctx).
		StringLength("title", "标题长度在2-50之间", 2, 50).
		StringLength("sub_title", "子标题长度在2-50之间", 2, 50).
		NotNull("content", "请输入活动内容").
		//StringLength("icon", "图标长度在2-100之间", 2, 100).
		IsDateTime("time_start", "请输入正确的开始时间格式").
		IsDateTime("time_end", "请输入正确的结束时间格式").
		IsNumeric("type", "类型必须为数字").
		InNumbers("status", "状态必须在可选的范围之内", &[]int64{0, 1}).
		//StringLength("link", "链接长度在2-150之间", 2, 150).
		//IsNumeric("award_money", "活动奖励金额必须为数字").
		//IsNumeric("award_ratio", "活动奖励百分比必须为数字").
		//IsNumeric("award_max", "奖励上限必须为数字").
		//IsNumeric("ip_limit", "同IP总限制次数必须为数字").
		//IsNumeric("ip_limit_day", "每日同IP限制次数必须为数字").
		//IsNumeric("is_fixed_money", "是否固定额度:0固定，1不固定必须为数字").
		//StringLength("award_condition", "活动奖励条件长度在2-255之间", 2, 255).
		//IsNumeric("total_award", "总奖励次数必须为数字").
		//IsNumeric("award_threshold", "登录、投注、充值量必须为数字").
		IsNumeric("activity_class_id", "活动分类（关联活动分类表）必须为数字").
		//InNumbers("is_show_submit", "是否显示提交按钮值不正确", &[]int64{0, 1}).
		//InNumbers("is_multiple_apply", "是否允许多次值不正确", &[]int64{0, 1}).
		//InNumbers("is_home_show", "是否首页弹出显示值不正确", &[]int64{0, 1}).
		Validate()
}
