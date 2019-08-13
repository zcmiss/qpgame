package validations

import (
	"qpgame/common/mvc"
)

type AdminRolesValidation struct{}

// 添加/修改动作数据验证
func (self AdminRolesValidation) Validate(ctx *Context) (string, bool) {
	return mvc.NewValidation(ctx).
		StringLength("name", "名称长度在2-20之间", 2, 20).
		//IsNumeric("parent_id", "父级编号必须为数字").
		IsNumeric("status", "状态必须为数字").
		//StringLength("remark", "备注长度在2-255之间", 2, 255).
		//NotNull("menu_ids", "角色下属菜单不能为空").
		Validate()
}

// 关于角色菜单的判断
func (self AdminRolesValidation) CheckMenus(ctx *Context) (string, bool) {
	return mvc.NewValidation(ctx).
		IsNumeric("id", "必须提供角色编号").
		NotNull("menu_ids", "必须提供角色菜单编号").
		Validate()
}
