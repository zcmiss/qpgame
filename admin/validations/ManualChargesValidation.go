package validations

import (
	"qpgame/common/mvc"
)

type ManualChargesValidation struct{}

// 添加/修改动作数据验证
func (self ManualChargesValidation) Validate(ctx *Context) (string, bool) {
	return mvc.NewValidation(ctx).
		//IsNumeric("user_id", "管理用户表ID必须为数字").
		IsDecimal("amount", "充值金额必须为数字").
		IsDecimal("benefits", "优惠金额必须为数字").
		IsDecimal("quantity", "打码量审核必须为数字").
		IsNumeric("audit", "是否流水审核，0为否，1为是必须为数字").
		IsNumeric("item", "存款项目必须是数字").
		StringLength("comment", "备注长度在2-255之间", 2, 255).
		//IsNumeric("deal_time", "交易日期必须为数字").
		//StringLength("operator", "操作人长度在2-255之间", 2, 255).
		//InNumbers("state", "审核状态必须在可选的范围之内", &[]int64{-1, 0, 1, 2}).
		Validate()
}
