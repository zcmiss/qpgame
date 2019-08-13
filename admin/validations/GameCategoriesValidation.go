package validations

import (
	"qpgame/common/mvc"
)

type GameCategoriesValidation struct{}

// 添加/修改动作数据验证
func (self GameCategoriesValidation) Validate(ctx *Context) (string, bool) {
	return mvc.NewValidation(ctx).
		IsNumeric("platform_id", "游戏平台编号必须为数字").
		IsNumeric("category_level", "分类层级(最多三层)必须为数字").
		StringLength("img", "2级分类游戏图片长度在2-500之间", 2, 500).
		StringLength("name", "游戏分类名称或者游戏名称长度在2-50之间", 2, 50).
		IsNumeric("parent_id", "上级游戏分类id必须为数字").
		IsNumeric("seq", "分类排序必须为数字").
		IsNumeric("status", "分类状态,0不可用 1可用必须为数字").
		StringLength("btn_img", "必须提供按钮图片", 2, 200).
		StringLength("btn_selected_img", "必须按供按钮选中图片", 2, 200).
		Validate()
}
