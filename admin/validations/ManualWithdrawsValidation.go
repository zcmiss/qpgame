package validations

import (
	"qpgame/common/mvc"
)

type ManualWithdrawsValidation struct{}

// 添加/修改动作数据验证
func (self ManualWithdrawsValidation) Validate(ctx *Context) (string, bool) {
	return mvc.NewValidation(ctx).
		IsNumeric("user_id", "关联用户表id必须为数字").
		IsDecimal("amount", "取款金额必须为数字").
		IsDecimal("quantity", "打码量必须为数字").
		IsNumeric("item", "取款项目必须是数字").
		//IsDateTime("deal_time", "交易时间必须合法").
		StringLength("comment", "备注长度在2-255之间", 2, 255).
		//StringLength("operator", "操作人长度在2-255之间", 2, 255).
		//InNumbers("state", "审核状态必须在可选的范围之内", &[]int64{-1, 0, 1, 2}).
		Validate()
}
