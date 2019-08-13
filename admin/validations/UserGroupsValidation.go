package validations

import (
	"qpgame/common/mvc"
)

type UserGroupsValidation struct{}

// 添加/修改动作数据验证
func (self UserGroupsValidation) Validate(ctx *Context) (string, bool) {
	return mvc.NewValidation(ctx).
		StringLength("group_name", "组名称称长度必须在2-50之间", 2, 50).
		StringLength("remark", "备注必须在6-20之间", 2, 50).
		Validate()
}
