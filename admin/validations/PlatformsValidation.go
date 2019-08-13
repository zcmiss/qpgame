package validations

import (
	"qpgame/common/mvc"
)

type PlatformsValidation struct{}

// 添加/修改动作数据验证
func (self PlatformsValidation) Validate(ctx *Context) (string, bool) {
	return mvc.NewValidation(ctx).
		StringLength("name", "平台名长度在2-30之间", 2, 30).
		StringLength("code", "平台代号长度在2-20之间", 2, 20).
		IsNumeric("status", "状态必须为数字").
		//StringLength("logo", "平台LOGO长度在2-255之间", 2, 255).
		//StringLength("index_logo", "平台首页logo,区别平台logo,有大小尺寸之分长度在2-255之间", 2, 255).
		//StringLength("content", "平台介绍长度在2-255之间", 2, 255).
		//IsNumeric("sort", "排序必须为数字").
		Validate()
}
