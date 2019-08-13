package validations

import (
	"qpgame/common/mvc"
)

type PayCredentialsValidation struct{}

// 添加/修改动作数据验证
func (self PayCredentialsValidation) Validate(ctx *Context) (string, bool) {
	return mvc.NewValidation(ctx).
		IsNumeric("plat_form", "支付标识必须是数字").
		StringLength("pay_name", "支付方式名称长度必须在2-50之间", 2, 50).
		StringLength("merchant_number", "商户编号长度必须在2-50之间", 2, 50).
		StringLength("private_key", "商户私钥长度必须在2-50之间", 2, 50).
		//StringLength("corporate", "法人长度必须在2-50之间", 2, 50).
		//IsID("id_umber", "法人身份证码必须输入正确").
		//IsBankCardNumber("card_number", "法人银行卡号必须输入正确").
		//IsPhoneNumber("phone_number", "手机号码长度必须输入正确").
		//StringLength("public_key", "证书公钥必须长度必须在2-50之间", 2, 500).
		//StringLength("private_key_file", "私钥文件商户编号必须长度必须在2-500之间", 2, 500).
		//StringLength("credential_key, ", "授权KEY商户编号必须长度必须在2-500之间", 2, 50).
		//StringLength("callback_key", "回调KEY,商户编号必须长度必须在2-50之间", 2, 50).
		IsNumeric("charge_amount_conf", "充值金额配置必须为数字").
		InNumbers("status", "状态必须为数字", &statusTypes).
		Validate()
}
