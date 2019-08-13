package validations

import (
	"qpgame/common/mvc"
)

type ActivityClassesValidation struct{}

// 添加/修改动作数据验证
func (self ActivityClassesValidation) Validate(ctx *Context) (string, bool) {
	return mvc.NewValidation(ctx).
		StringLength("name", "活动分类名称长度在2-50之间", 2, 50).
		InNumbers("status", "分类状态输入不正确", &[]int64{0, 1}).
		//IsNumeric("seq", "分类排序必须为数字").
		Validate()
}
