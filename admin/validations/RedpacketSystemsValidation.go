package validations

import (
	"qpgame/common/mvc"
)

type RedpacketSystemsValidation struct{}

// 添加/修改动作数据验证
func (self RedpacketSystemsValidation) Validate(ctx *Context) (string, bool) {
	return mvc.NewValidation(ctx).
		IsNumeric("type", "红包种类必须为数字").
		IsDecimal("money", "总金额必须为数字").
		IsNumeric("total", "红包总数量，如果是0就是没有红包个数限制,派完为止必须为数字").
		IsNumeric("status", "状态必须为数字").
		StringLength("message", "红包消息长度在2-255之间", 2, 255).
		IsNumeric("calculate_type", "红包发放方式必须为数字").
		Validate()
}
