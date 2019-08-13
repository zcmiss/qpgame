package validations

import (
	"qpgame/common/mvc"
)

type AppVersionsValidation struct{}

// 添加/修改动作数据验证
func (self AppVersionsValidation) Validate(ctx *Context) (string, bool) {
	return mvc.NewValidation(ctx).
		IsVersion("version", "版本号必须符合规范(如:1.1.1)").
		InNumbers("status", "状态必须在有效范围之内", &[]int64{0, 1}).
		NotNull("description", "请输入版本说明的内容").
		IsUrl("link", "下载地址必须是有效的链接").
		InNumbers("package_type", "包类型必须在有效范围之内", &[]int64{1, 2}).
		InNumbers("app_type", "APP类型必须在有效范围之内", &[]int64{1, 2}).
		InNumbers("update_type", "更新方式必须在有效范围之内", &[]int64{1, 2, 3}).
		Validate()
}
