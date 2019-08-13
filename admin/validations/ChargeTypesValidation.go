package validations

import (
	"qpgame/common/mvc"
)

type ChargeTypesValidation struct{}

// 添加/修改动作数据验证
func (self ChargeTypesValidation) Validate(ctx *Context) (string, bool) {
	return mvc.NewValidation(ctx).
		StringLength("name", "充值名称长度在2-30之间", 2, 30).
		StringLength("remark", "充值说明长度在2-512之间", 2, 512).
		IsNumeric("state", "状态必须为数字").
		//IsIntegers("charge_numbers", "必须充值金额选项, 例如:[50,100,300,800,1000,2000,3000,5000]").
		StringLength("logo", "必须上传充值方式的图片", 2, 255).
		StringLength("logo_selected", "必须上传充值方式(选中)的图片", 2, 255).
		IsNumeric("priority", "类型排序必须为数字").
		Validate()
}
