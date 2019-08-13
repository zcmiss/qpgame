package config

/**
 * ================ README!!!!! =========================
 *
 * 第三方登录配置信息 *
 * 第三方订单生成请求地址
 * ======================================================
 */
var TppayConfig = map[int]map[string]interface{}{
	62: { //蓉银支付
		"tjurl": "http://103.48.5.57/v1/pay",
	},
	63: { //利盈支付
		"tjurl": map[string]string{
			"1": "http://limeapi.lying-pay.com/v1/cash",
			"2": "http://limeapi.lying-pay.com/v1/query/trade",
		},
	},
	67: { //全通云付
		"tjurl": "https://www.allpasspay.com/hspay/api_node/",
	},
	68: { //联动支付
		"tjurl": "http://pay1.liandzf.com:14562/chargebank.aspx",
	},
	88: { //立信支付
		"tjurl": "https://api.paylixin.com/api/pay",
	},
	89: { //智慧宝支付
		"tjurl": "http://pay.zhihuibao.vip/pay/api.php",
	},
	100: { //金砖支付
		"orderUrl": "https://api.jz-pay.com/v1.0/pay/create_order",
	},
	103: { //钉钉支付
		"tjurl": "http://47.110.245.153:8080/api/pay",
	},
	104: { //七星支付
		"tjurl": "https://openapi.7pay.shysrj.com/gateway/pay",
	},
	105: { //喜多多支付
		"tjurl": "http://gateway.xdd668.cn/paygateway/syPayOrder/pay",
	},
}
