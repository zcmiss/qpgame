package validations

import (
	"qpgame/common/mvc"
)

type ChargeCardsValidation struct{}

// 添加/修改动作数据验证
func (self ChargeCardsValidation) Validate(ctx *Context) (string, bool) {
	return mvc.NewValidation(ctx).
		//StringLength("name", "银行名称长度在2-30之间", 2, 30).
		//StringLength("owner", "持卡人长度在2-30之间", 2, 30).
		//StringLength("card_number", "卡号长度在2-30之间", 2, 30).
		//StringLength("bank_address", "开户地址长度在2-150之间", 2, 150).
		//IsNumeric("charge_type_id", "充值类型编号必须为数字").
		//StringLength("remark", "备注长度在2-255之间", 2, 255).
		//IsNumeric("state", "状态必须为数字").
		//StringLength("logo", "LOGO长度在2-100之间", 2, 100).
		//StringLength("hint", "支付提示长度在2-150之间", 2, 150).
		//StringLength("title", "支付标题长度在2-30之间", 2, 30).
		//IsNumeric("mfrom", "支付额度必须为数字").
		//IsNumeric("mto", "支付额度必须为数字").
		//IsNumeric("amount_limit", "停用金额必须为数字").
		//NotNull("addr_type", "字段 addr_type 不能为空").
		//IsNumeric("credential_id", "以付方式必须为数字").
		//IsNumeric("priority", "充值排序必须为数字").
		//StringLength("qr_code", "财务收款二维码长度在2-150之间", 2, 150).
		Validate()
}
