package validations

import (
	"qpgame/common/mvc"
)

type SystemNoticesValidation struct{}

// 添加/修改动作数据验证
func (self SystemNoticesValidation) Validate(ctx *Context) (string, bool) {
	return mvc.NewValidation(ctx).
		StringLength("title", "标题长度在2-30之间", 2, 30).
		NotNull("content", "请输入公告内容的内容").
		InNumbers("status", "状态必须在可选的范围之内", &statusTypes).
		Validate()
}
