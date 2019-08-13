package validations

import (
	"qpgame/common/mvc"
)

type AdminNodesValidation struct{}

// 添加/修改动作数据验证
func (self AdminNodesValidation) Validate(ctx *Context) (string, bool) {
	return mvc.NewValidation(ctx).
		StringLength("name", "节点名称长度在2-20之间", 2, 20).
		StringLength("route", "路由长度在2-100之间", 2, 100).
		StringLength("title", "说明长度在2-50之间", 2, 50).
		StringLength("method", "请输入方法的内容", 2, 200).
		IsNumeric("status", "状态必须为数字").
		StringLength("remark", "备注长度在2-255之间", 0, 255).
		IsNumeric("seq", "排序必须为数字").
		IsNumeric("parent_id", "上级编号必须为数字").
		//IsNumeric("level", "级别必须为数字").
		IsNumeric("type", "类型必须为数字").
		Validate()
}
