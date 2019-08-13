package pay

import (
	"qpgame/common/utils"
	"qpgame/models/xorm"
)

//专门用于回调处理
//100金砖支付回调处理
func tpPay100JinZhuanCallback(data map[string]string, rowCred xorm.PayCredentials) map[string]interface{} {
	postSign, exitSign := data["sign"]
	mchTradeNo, exitMchTradeNo := data["mch_trade_no"]
	amount, exitAmount := data["amount"]
	if !exitMchTradeNo {
		mchTradeNo = ""
	}
	if !exitAmount {
		amount = ""
	}
	if exitSign {
		str := payKsortAddstr(data, 1)
		str = str + "&key=" + rowCred.PrivateKey
		sign := utils.Php2GoMd5(str)
		if sign == postSign {
			return map[string]interface{}{
				"status":  1,
				"amount":  amount,
				"orderId": mchTradeNo,
				"sign":    "success",
			}
		}
	}
	return map[string]interface{}{
		"status":  2,
		"amount":  amount,
		"orderId": mchTradeNo,
		"sign":    "success",
	}

}
