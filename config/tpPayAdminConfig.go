package config

/*
后台管理线上入款账号添加的时候的配置选择
*/

var PayCredentialsPlat = map[int]string{
	1:   "XINDUI",
	2:   "紫控科技",
	3:   "长城支付",
	4:   "科讯支付",
	5:   "迅银支付",
	6:   "雅付支付",
	7:   "保时捷支付",
	8:   "UUPAY",
	9:   "聚合支付",
	10:  "仁杰支付",
	11:  "汇达支付",
	12:  "高通支付",
	13:  "金阳支付",
	14:  "天泽支付",
	15:  "至尊星付",
	16:  "酷游支付",
	17:  "w付",
	18:  "佰富支付",
	19:  "百盛支付",
	20:  "鑫支付",
	21:  "尚付云支付",
	22:  "艾付",
	23:  "付营通",
	24:  "优米付",
	25:  "星和易通",
	26:  "便利付",
	27:  "仁信支付",
	28:  "如一付",
	29:  "五福支付",
	30:  "e汇支付",
	31:  "易宝支付",
	32:  "易付支付",
	33:  "新艾米森1",
	34:  "新艾米森2",
	35:  "首捷支付",
	36:  "大千支付",
	37:  "易到支付",
	38:  "BB付",
	39:  "商易付",
	40:  "佰钱付",
	41:  "云安付",
	42:  "80卡",
	43:  "会昶支付",
	44:  "80付",
	45:  "凤翼天翔",
	46:  "686充值",
	47:  "明捷付",
	48:  "聚鑫支付",
	49:  "369支付",
	51:  "鑫发支付",
	52:  "今汇通支付",
	53:  "隆发支付",
	54:  "盈支付",
	55:  "乐游支付",
	56:  "A支付",
	57:  "安亿网络",
	58:  "个付",
	59:  "中联支付",
	60:  "万家支付",
	61:  "天润支付",
	62:  "蓉银支付",
	63:  "利盈支付",
	64:  "环球支付",
	65:  "58支付",
	66:  "新聚联支付",
	67:  "全通云付",
	68:  "联动支付",
	69:  "顺优付",
	70:  "钱快付",
	71:  "鼎融支付",
	72:  "锤子支付",
	73:  "闪付支付",
	74:  "亿仟支付",
	75:  "689支付",
	76:  "CK支付",
	77:  "DF点付支付",
	78:  "闪快支付",
	79:  "畅付通支付",
	80:  "德宝莱支付",
	81:  "新宝付支付",
	82:  "豆宝支付",
	83:  "MyPay支付",
	84:  "时代支付",
	85:  "瓜子支付",
	86:  "快闪支付",
	87:  "四海云付",
	88:  "立信支付",
	89:  "智慧宝支付",
	90:  "易富宝支付",
	91:  "星付通支付",
	92:  "百万通支付",
	93:  "99支付",
	94:  "星辰支付",
	95:  "大唐支付",
	96:  "腾云支付",
	97:  "百威支付",
	98:  "开元支付",
	99:  "天瑞支付",
	100: "金砖支付",
	101: "大清支付",
	102: "爱玛支付",
	103: "钉钉支付",
	104: "七星支付",
	105: "喜多多支付",
	106: "雀付支付",
	107: "一加支付",
	108: "汇宏支付",
	109: "GTPAY支付",
	110: "天使支付",
	111: "火箭支付",
	112: "永仁支付",
	113: "龙支付",
	114: "百家付",
}
var PayCredentialsPlayTypes = map[int]map[int][]map[string]string{
	2: {
		4: {
			{"code": "901", "name": "微信公众号"},
			{"code": "902", "name": "微信扫码支付"},
			{"code": "903", "name": "支付宝扫码支付"},
			{"code": "904", "name": "支付宝手机"},
			{"code": "905", "name": "QQ手机支付"},
			{"code": "907", "name": "网银支付"},
			{"code": "908", "name": "QQ扫码支付"},
			{"code": "909", "name": "百度钱包"},
			{"code": "910", "name": "京东支付"},
		},
	},
	3: {
		4: {
			{"code": "1", "name": "支付宝扫码支付"},
			{"code": "2", "name": "微信扫码支付"},
			{"code": "3", "name": "百度钱包支付"},
			{"code": "4", "name": "QQ钱包支付"},
			{"code": "5", "name": "京东钱包支付"},
		},
	},
	4: {
		4: {
			{"code": "1006", "name": "QQ钱包支付"},
			{"code": "1010", "name": "京东钱包支付"},
		},
	},
	5: {
		5: {
			{"code": "993", "name": "QQ钱包wap支付"},
			{"code": "1004", "name": "微信wap支付"},
		},
	},
	6: {
		4: {
			{"code": "0202", "name": "微信扫码支付"},
			{"code": "0302", "name": "支付宝扫码支付"},
			{"code": "0502", "name": "QQ钱包扫码支付"},
			{"code": "0802", "name": "京东钱包扫码支付"},
		},
		5: {
			{"code": "0201", "name": "微信wap支付"},
			{"code": "0301", "name": "支付宝wap支付"},
			{"code": "0501", "name": "QQ钱包wap支付"},
			{"code": "0801", "name": "京东钱包wap支付"},
		},
		6: {
			{"code": "0901", "name": "微信H5支付"},
			{"code": "0303", "name": "支付宝H5支付"},
			{"code": "0503", "name": "QQ钱包H5支付"},
			{"code": "0803", "name": "京东钱包H5支付"},
		},
	},
	7: {
		5: {
			{"code": "2001", "name": "微信wap支付"},
		},
		6: {
			{"code": "2005", "name": "微信H5支付"},
			{"code": "2009", "name": "QQ钱包H5支付"},
		},
	},
	8: {
		4: {
			{"code": "WECHAT_QRCODE_PAY", "name": "微信扫码支付"},
			{"code": "ALIPAY_QRCODE_PAY", "name": "支付宝扫码支付"},
		},
		6: {
			{"code": "QQ_QRCODE_PAY", "name": "QQ扫码支付"},
			{"code": "H5_WECHAT_PAY", "name": "微信H5支付"},
			{"code": "H5_ALI_PAY", "name": "H5 支付宝"},
		},
	},
	9: { //聚合支付
		4: {
			{"code": "10", "name": "微信扫码支付"},
		},
		6: {
			{"code": "11", "name": "支付宝支付"},
			{"code": "12", "name": "快捷支付-PC"},
			{"code": "13", "name": "快捷支付-H5"},
		},
	},
	10: { //仁杰支付
		6: {
			{"code": "1", "name": "仁杰在线支付"},
		},
	},
	11: { //汇达支付
		6: {
			{"code": "1", "name": "网银"},
			{"code": "4", "name": "支付宝扫码"},
			{"code": "5", "name": "微信扫码"},
			{"code": "6", "name": "QQ扫码"},
			{"code": "8", "name": "京东扫码"},
			{"code": "9", "name": "支付宝H5"},
			{"code": "10", "name": "支付宝WAP"},
			{"code": "11", "name": "微信WAP"},
			{"code": "12", "name": "平台快捷支付"},
			{"code": "16", "name": "支付宝收银台"},
		},
	},
	12: { //高通支付
		5: {
			{"code": "WEIXIN", "name": "高通微信支付"},
			{"code": "ALIPAY", "name": "高通支付宝支付"},
			{"code": "QQPAY", "name": "高通QQ钱包支付"},
			{"code": "UNIONPAY", "name": "银联扫码支付"},
			{"code": "JDPAY", "name": "京东扫码支付"},
			{"code": "JDPAYWAP", "name": "京东钱包支付"},
		},
	},
	13: { //金阳支付
		4: {
			{"code": "WEIXIN", "name": "微信支付"},
			{"code": "ALIPAY", "name": "支付宝支付"},
			{"code": "QQPAY", "name": "QQ钱包支付"},
		},
		5: {
			{"code": "WEIXINWAP", "name": "微信支付"},
			{"code": "ALIPAYWAP", "name": "支付宝支付"},
			{"code": "QQPAYWAP", "name": "QQ钱包支付"},
		},
		6: {
			{"code": "WEIXINWAP", "name": "微信支付"},
			{"code": "ALIPAYWAP", "name": "支付宝支付"},
			{"code": "QQPAYWAP", "name": "QQ钱包支付"},
			{"code": "UNIONWAPPAY", "name": "银联H5支付"},
		},
	},
	14: { //天泽支付
		6: {
			{"code": "wxcode", "name": "微信扫码支付"},
			{"code": "alipay", "name": "支付宝扫码支付"},
			{"code": "wxwap", "name": "微信H5支付"},
			{"code": "alipaywap", "name": "支付宝H5支付"},
			{"code": "qqpay", "name": "QQ钱包支付"},
		},
	},
	15: { //至尊星付
		4: {
			{"code": "1", "name": "支付宝支付"},
			{"code": "2", "name": "微信支付"},
			{"code": "3", "name": "QQ钱包支付"},
			{"code": "5", "name": "京东钱包支付"},
		},
		6: {
			{"code": "1", "name": "支付宝支付"},
			{"code": "2", "name": "微信支付"},
			{"code": "3", "name": "QQ钱包支付"},
			{"code": "5", "name": "京东钱包支付"},
		},
	},
	16: { //酷游支付
		4: {
			{"code": "WX-YF", "name": "微信扫码预付"},
			{"code": "WX-JS", "name": "微信扫码支付"},
			{"code": "ALIQR-YF", "name": "支付宝扫码预付"},
			{"code": "ALIQR-JS", "name": "支付宝扫码支付"},
		},
		5: {
			{"code": "EBANK-YF", "name": "网银预付"},
			{"code": "EBANK-JS", "name": "网银支付"},
			{"code": "QUICKPAY-JS", "name": "快捷支付"},
			{"code": "QQWALLET-YF", "name": "QQ钱包预付"},
			{"code": "QQWALLET-JS", "name": "QQ钱包支付"},
			{"code": "WX-YF", "name": "微信扫码预付"},
			{"code": "WX-JS", "name": "微信扫码支付"},
			{"code": "WXWAP-YF", "name": "微信WAP预付"},
			{"code": "WXWAP-JS", "name": "微信WAP支付"},
			{"code": "ALIQR-YF", "name": "支付宝扫码预付"},
			{"code": "ALIQR-JS", "name": "支付宝扫码支付"},
		},
		6: {
			{"code": "ALIPAYMOBILE-YF", "name": "支付宝手机网页预付"},
			{"code": "ALIPAYMOBILE-JS", "name": "支付宝手机网页支付"},
		},
	},
	17: { //w付
		4: {
			{"code": "alipay_scan", "name": "支付宝支付"},
			{"code": "weixin_scan", "name": "微信支付"},
			{"code": "tenpay_scan", "name": "QQ钱包支付"},
		},
		5: {
			{"code": "alipay_scan", "name": "支付宝支付"},
			{"code": "weixin", "name": "微信支付"},
			{"code": "tenpay_scan", "name": "QQ钱包支付"},
			{"code": "b2c", "name": "网银支付"},
		},
		6: {
			{"code": "alipay_h5api", "name": "支付宝支付"},
			{"code": "weixin_h5api", "name": "微信支付"},
			{"code": "qq_h5api", "name": "QQ钱包支付"},
		},
	},
	18: { //佰富
		4: {
			{"code": "ZFB", "name": "支付宝支付"},
			{"code": "WX", "name": "微信支付"},
			{"code": "QQ", "name": "QQ钱包支付"},
			{"code": "JDQB", "name": "京东支付"},
			{"code": "YL", "name": "银联支付"},
			{"code": "KJ", "name": "快捷支付"},
		},
		6: {
			{"code": "ZFB", "name": "支付宝支付"},
			{"code": "WX_WAP", "name": "微信支付"},
			{"code": "QQ", "name": "QQ钱包支付"},
			{"code": "JDQB", "name": "京东支付"},
			{"code": "YL", "name": "银联支付"},
			{"code": "KJ", "name": "快捷支付"},
		},
	},
	19: { //佰盛
		4: {
			{"code": "ALIPAY_QRCODE_PAY", "name": "支付宝支付"},
			{"code": "WECHAT_QRCODE_PAY", "name": "微信支付"},
			{"code": "QQ_QRCODE_PAY", "name": "QQ钱包支付"},
			{"code": "JD_QRCODE_PAY", "name": "京东支付"},
			{"code": "UNIONPAY_QRCODE_PAY", "name": "银联支付"},
			{"code": "ONLINE_BANK_QUICK_PAY", "name": "快捷支付"},
		},
		6: {
			{"code": "ALIPAY_QRCODE_PAY", "name": "支付宝支付"},
			{"code": "WECHAT_QRCODE_PAY", "name": "微信支付"},
			{"code": "QQ_QRCODE_PAY", "name": "QQ钱包支付"},
			{"code": "JD_QRCODE_PAY", "name": "京东支付"},
			{"code": "UNIONPAY_QRCODE_PAY", "name": "银联支付"},
			{"code": "ONLINE_BANK_QUICK_PAY", "name": "快捷支付"},
		},
	},
	20: { //鑫
		4: {
			{"code": "4", "name": "支付宝支付"},
			{"code": "7", "name": "微信支付"},
			{"code": "9", "name": "QQ钱包支付"},
			{"code": "12", "name": "京东支付"},
			{"code": "11", "name": "快捷支付"},
		},
		6: {
			{"code": "4", "name": "支付宝支付"},
			{"code": "7", "name": "微信支付"},
			{"code": "9", "name": "QQ钱包支付"},
			{"code": "12", "name": "京东支付"},
			{"code": "11", "name": "快捷支付"},
		},
	},
	21: { //尚付云
		4: {
			{"code": "1", "name": "支付宝扫码支付"},
			{"code": "2", "name": "微信扫码支付"},
			{"code": "3", "name": "QQ钱包支付"},
			{"code": "4", "name": "银联扫码支付"},
			{"code": "5", "name": "京东钱包支付"},
		},
		6: {
			{"code": "1", "name": "支付宝扫码支付"},
			{"code": "2", "name": "微信扫码支付"},
			{"code": "3", "name": "QQ钱包支付"},
		},
	},
	22: { //艾付
		4: {
			{"code": "ALIPAY", "name": "支付宝扫码支付"},
			{"code": "WECHAT", "name": "微信扫码支付"},
			{"code": "QQSCAN", "name": "QQ钱包支付"},
			{"code": "JDSCAN", "name": "银联扫码支付"},
			{"code": "UNIONPAY", "name": "京东钱包支付"},
		},
		6: {
			{"code": "WECHATWAP", "name": "微信扫码支付"},
			{"code": "QQWAP", "name": "QQ钱包支付"},
		},
	},
	23: { //付营通
		5: {
			{"code": "1", "name": "微信扫码支付"},
			{"code": "2", "name": "支付宝扫码支付"},
		},
	},
	24: { //优米付
		5: {
			{"code": "4", "name": "支付宝扫码支付"},
			{"code": "5", "name": "微信扫码支付"},
			{"code": "6", "name": "QQ钱包支付"},
			{"code": "17", "name": "银联扫码支付"},
			{"code": "8", "name": "京东钱包支付"},
		},
		6: {
			{"code": "9", "name": "支付宝支付"},
			{"code": "13", "name": "微信H5支付"},
			{"code": "15", "name": "QQ钱包支付"},
			{"code": "21", "name": "京东钱包支付"},
		},
	},
	25: { //星和易通
		4: {
			{"code": "1", "name": "支付宝扫码支付"},
			{"code": "2", "name": "微信扫码支付"},
			{"code": "3", "name": "QQ钱包支付"},
			{"code": "4", "name": "银联扫码支付"},
			{"code": "5", "name": "京东钱包支付"},
		},
		6: {
			{"code": "1", "name": "支付宝支付"},
			{"code": "2", "name": "微信H5支付"},
			{"code": "3", "name": "QQ钱包支付"},
		},
	},
	26: { //便利付
		5: {
			{"code": "WEIXIN_NATIVE", "name": "微信扫码支付"},
			{"code": "QQ_NATIVE", "name": "QQ钱包支付"},
			{"code": "KUAIJIE", "name": "快捷支付"},
		},
		6: {
			{"code": "KUAIJIE", "name": "快捷支付"},
			{"code": "WEIXIN_NATIVE", "name": "微信H5支付"},
			{"code": "QQ_NATIVE", "name": "QQ钱包支付"},
		},
	},
	27: { //仁信
		4: {
			{"code": "WEIXIN", "name": "微信扫码支付"},
			{"code": "QQ", "name": "QQ钱包支付"},
			{"code": "ALIPAY", "name": "支付宝"},
		},
		5: {
			{"code": "WEIXINWAP", "name": "微信扫码支付"},
			{"code": "QQ_NATIVE", "name": "QQ钱包支付"},
			{"code": "KUAIJIE", "name": "快捷支付"},
		},
		6: {
			{"code": "KUAIJIE", "name": "快捷支付"},
			{"code": "WEIXIN_NATIVE", "name": "微信H5支付"},
			{"code": "QQWAP", "name": "QQ钱包支付"},
		},
	},
	28: { //如一付
		5: {
			{"code": "21", "name": "微信扫码支付"},
			{"code": "89", "name": "QQ钱包支付"},
			{"code": "32", "name": "快捷支付"},
			{"code": "2", "name": "支付宝"},
		},
		6: {
			{"code": "31", "name": "快捷支付"},
			{"code": "33", "name": "微信H5支付"},
			{"code": "92", "name": "QQ钱包支付"},
			{"code": "36", "name": "支付宝"},
		},
	},
	29: { //五福支付
		4: {
			// {"code":"ALIPAY_NATIVE","name":"支付宝扫码支付"},
			//{"code":"WEIXIN_NATIVE","name":"微信扫码支付"},
			//{"code":"QQ_NATIVE","name":"QQ钱包支付"},
			{"code": "ALIPAY_NATIVE", "name": "支付宝扫码"},
			{"code": "WEIXIN_NATIVE", "name": "微信反扫码"},
			{"code": "UNIONPAY_NATIVE", "name": "快捷支付"},
		},
		6: {
			//{"code":"ALIPAY_H5","name":"支付宝"},
			{"code": "WEIXIN_H5", "name": "微信条形码支付"},
			{"code": "ALIPAY_H5", "name": "支付宝H5"},
			//{"code":"QQ_H5","name":"QQ钱包支付"},
			// {"code":"JD_H5","name":"京东支付"},
			//{"code":"UNIONPAY_H5","name":"银联H5支付"},
			{"code": "0000000", "name": "收银台（显示全部支付产品）"},
			// {"code":"0000001","name":"收银台（仅显示网银支付产品）"},
			// {"code":"0000002","name":"收银台（仅显示快捷支付产品）"},
			//{"code":"0000003","name":"收银台（仅充值卡支付产品）"},
			//{"code":"0000004","name":"收银台（仅显示扫码类支付产品）"},
			//{"code":"0000012","name":"收银台（仅显示显示扫码类支支付产品）"},
		},
	},

	31: { //易宝支付
		5: {
			{"code": "UNIONPAY_NATIVE", "name": "快捷支付"},
		},
		6: {
			{"code": "UNIONPAY_NATIVE", "name": "快捷支付"},
		},
	},
	32: { //易付支付
		4: {
			{"code": "alipay_scan", "name": "支付宝支付"},
			{"code": "alipay", "name": "支付宝电脑端"},
			{"code": "weixin_scan", "name": "微信支付"},
			{"code": "tenpay_scan", "name": "QQ支付"},
			{"code": "jdpay_scan", "name": "JD支付"},
			{"code": "ylpay_scan", "name": "银联扫码"},
		},
		6: {
			{"code": "alipay_h5api", "name": "支付宝支付"},
			{"code": "weixin_h5api", "name": "微信支付"},
			{"code": "qq_h5api", "name": "QQ支付"},
			{"code": "direct_pay", "name": "快捷支付"},
		},
	},
	33: { //新艾米森1
		4: {
			{"code": "40104", "name": "QQ钱包支付"}, //最小2元，最大1000元
			{"code": "60104", "name": "支付宝支付"},
		},
		5: {
			{"code": "50107", "name": "微信支付"},
			{"code": "60107", "name": "支付宝支付"},
		},
	},
	34: { //新艾米森2
		4: {
			{"code": "40104", "name": "QQ钱包支付"}, //最小2元，最大1000元
			{"code": "60104", "name": "支付宝支付"},
		},
		5: {
			{"code": "50107", "name": "微信支付"},
			{"code": "60107", "name": "支付宝支付"},
		},
	},
	35: { //首捷支付
		5: {
			{"code": "weixin", "name": "微信支付"},
			{"code": "alipay", "name": "支付宝支付"}, //最小2元，最大1000元
			{"code": "qq", "name": "QQ钱包支付"},
		},
		6: {
			{"code": "wxwap", "name": "微信支付"},
			{"code": "alipaywap", "name": "支付宝支付"},
			{"code": "qqwap", "name": "QQ钱包支付"},
		},
	},
	36: { //大千支付
		5: {
			{"code": "1", "name": "支付宝支付"},
			{"code": "2", "name": "微信支付"},
			{"code": "3", "name": "QQ钱包支付"},
			{"code": "4", "name": "京东支付"},
			{"code": "5", "name": "快捷支付"},
		},
		6: {
			{"code": "1", "name": "支付宝支付"},
			{"code": "2", "name": "微信支付"},
			{"code": "3", "name": "QQ钱包支付"},
			{"code": "4", "name": "京东支付"},
			{"code": "5", "name": "快捷支付"},
		},
	},
	39: { //商易付
		5: {
			{"code": "8012", "name": "支付宝支付"},
			{"code": "932", "name": "支付宝支付"},
			{"code": "934", "name": "微信支付"},
			{"code": "936", "name": "QQ钱包支付"},
			{"code": "911", "name": "京东支付"},
			{"code": "963", "name": "网银支付"},
			{"code": "2000", "name": "快捷支付"},
		},
		6: {
			{"code": "931", "name": "支付宝支付"},
			{"code": "933", "name": "微信支付"},
			{"code": "935", "name": "QQ钱包支付"},
			{"code": "964", "name": "网银支付"},
			{"code": "2001", "name": "快捷支付"},
		},
	},
	40: { //佰钱付
		4: {
			{"code": "BSM", "name": "银联扫码"},
		},
		6: {
			{"code": "ALIPAYH5", "name": "支付宝支付"},
		},
	},
	41: { //云安付
		4: {
			{"code": "WEIXINPAY", "name": "微信扫码"},
			{"code": "ALIPAY", "name": "支付宝扫码"},
		},
		5: {
			{"code": "ALIPAY", "name": "支付宝支付"},
			{"code": "JDPAY", "name": "京东支付"},
			{"code": "UNIONPAY", "name": "银联支付"},
			{"code": "PAYMODE", "name": "收银台模式"},
		},
		6: {
			{"code": "ALIWAPPAY", "name": "支付宝支付"},
			{"code": "BANKH5", "name": "网银快捷"},
			{"code": "UNIONPAYWAP", "name": "网银快捷"},
		},
	},
	42: { //80卡
		5: {
			{"code": "16801356", "name": "绑卡快捷"},
		},
		6: {
			{"code": "16801353", "name": "快捷H5"},
		},
	},
	43: { //会昶支付
		6: {
			//{"code":"ALIPAYH5","name":"支付宝H5"},
		},
	},
	44: { //80付
		6: {
			{"code": "ALIPAYH5", "name": "支付宝H5"},
		},
	},
	45: { //凤翼天翔
		4: {
			{"code": "alipayai", "name": "支付宝扫码"},
			{"code": "wechatpe", "name": "微信扫码"},
			{"code": "wechat13", "name": "微信扫码13"},
		},
		5: {
			{"code": "alipay", "name": "支付宝wap"},
			{"code": "all", "name": "聚合收银台"},
			{"code": "wechatpe", "name": "微信wap"},
		},
		6: {
			{"code": "alipayai", "name": "支付宝H5"},
		},
	},
	46: { //686支付
		4: {
			{"code": "aliqr", "name": "支付宝扫码"},
			{"code": "wxpay", "name": "微信扫码"},
		},
		6: {
			{"code": "wxh5", "name": "微信h5"},
			{"code": "alipay", "name": "支付宝H5"},
		},
	},
	47: { //明捷付
		4: {
			{"code": "ZFB", "name": "支付宝扫码"},
			{"code": "JD", "name": "京东钱包扫码"},
			{"code": "UNION_WALLET", "name": "银联钱包扫码"},
		},
		5: {
			{"code": "MBANK", "name": "手机银行扫码"},
		},
		6: {
			{"code": "ZFB_WAP", "name": "支付宝WAP"},
			{"code": "WX_WAP", "name": "微信WAP"},
		},
	},
	48: { //聚鑫支付
		6: {
			{"code": "alipay", "name": "支付宝H5"},
			{"code": "941", "name": "支付宝H5-941"},
			{"code": "942", "name": "支付宝H5-942"},
		},
	},
	49: { //369支付
		4: {
			{"code": "alipay", "name": "支付宝扫码"},
		},
		6: {
			{"code": "alipaywap", "name": "支付宝H5"},
		},
	},
	51: { //鑫发支付
		4: {
			{"code": "WX", "name": "微信"},          //微信扫码支付
			{"code": "ZFB", "name": "支付宝"},        //支付宝扫码支付
			{"code": "QQ", "name": "QQ钱包"},        //QQ钱包扫码支付
			{"code": "ZFB_WAP", "name": "支付宝WAP"}, //手机端跳转支付宝支付
		},
		6: {
			{"code": "WX_WAP", "name": "微信WAP"},           //手机端跳转微信支付
			{"code": "WX_H5", "name": "微信H5"},             //手机端跳转微信支付
			{"code": "ZFB_HB", "name": "支付宝红包"},           //支付宝红包支付
			{"code": "ZFB_HB_H5", "name": "支付宝红包H5"},      //手机端跳转支付宝红包支付
			{"code": "QQ_WAP", "name": "QQ钱包WAP"},         //手机端跳转QQ钱包支付
			{"code": "UNION_WALLET", "name": "银联钱包(云闪付)"}, //银联钱包扫码支付
			{"code": "UNION_WAP", "name": "银联WAP"},        //手机端银联在线支付
		},
	},
	52: { //今汇通支付
		5: {
			{"code": "wechat", "name": "新微信"},
			{"code": "wechat1", "name": "微信"},
			{"code": "wechat11", "name": "微信11"},
			{"code": "wechatpe", "name": "微信pe"},
		},
		6: {
			{"code": "alipay", "name": "支付宝H5"},
			{"code": "alipay2", "name": "支付宝企业_H5"},
			{"code": "alipay4", "name": "支付宝5_H5"},
			{"code": "alipay_wap", "name": "支付宝wap"},
			{"code": "alipay7", "name": "支付宝7"},
			{"code": "alipay6", "name": "快乐支付宝"},
		},
	},
	53: { //隆发支付
		4: {
			{"code": "ZFB", "name": "支付宝扫码"},
			{"code": "QQ", "name": "QQ钱包扫码"},
			{"code": "JD", "name": "京东钱包扫码"},
		},
		6: {
			{"code": "ZFB_WAP", "name": "支付宝H5"},
			{"code": "WX_WAP", "name": "微信H5"},
			{"code": "QQ_WAP", "name": "QQ钱包H5"},
		},
	},
	54: { //盈支付
		6: {
			{"code": "904", "name": "支付宝手机"},
		},
	},
	55: { //乐游支付
		6: {
			{"code": "zfbsm", "name": "支付宝扫码"},
		},
		5: {
			{"code": "zfbwap", "name": "支付宝手机"},
		},
	},
	56: { //A支付
		//            4:{
		//                {"code":"70000203","name":"QQ钱包扫码T0"},
		//
		//                {"code":"80000203","name":"JD钱包扫码T0支付"},
		//                {"code":"20000303","name":"支付宝T0扫码支付"},
		//                {"code":"10000103","name":"微信扫码T0"},
		//                {"code":"20000501","name":"支付宝条形码T0支付"},
		//            },
		5: {
			//                {"code":"40000503","name":"快捷支付T0"},
			//                {"code":"50000103","name":"网银B2C支付T0"},
			//                {"code":"10000203","name":"微信WAP支付"},
			//
			//                {"code":"70000204","name":"QQWAPT0"},
			//                {"code":"40000701","name":"快捷WAPT0"},
		},
		6: {
			//               {"code":"80000303","name":"JDH5支付T0"},
			{"code": "40000503", "name": "快捷支付T0"},
			{"code": "50000103", "name": "网银B2C支付T0"},
			{"code": "20000203", "name": "支付宝wap"},
			//                 {"code":"60000103","name":"银联钱包扫码T0支付"},
		},
	},
	57: { //安亿网络
		4: { //扫码
			{"code": "alipay", "name": "支付宝扫码"},
			{"code": "alipayH5", "name": "支付宝h5"},
			{"code": "unionpay", "name": "银联扫码"},
		},
	},
	58: { //个付
		//wx=微信,wxwap=微信WAP,ali=支付宝,aliwap=支付宝WAP,qq=QQ,qqwap=QQWAP
		6: {
			{"code": "wx", "name": "微信"},
			{"code": "wxwap", "name": "微信WAP"},
			{"code": "ali", "name": "支付宝"},
			{"code": "aliwap", "name": "支付宝WAP"},
			{"code": "qq", "name": "QQ"},
			{"code": "qqwap", "name": "QQWAP"},
		},
	},
	59: { //中联支付
		5: {
			{"code": "6006", "name": "网关支付"},
			{"code": "6002", "name": "支付宝扫码"},
		},
		6: {
			{"code": "6004", "name": "微信H5"},
			{"code": "6001", "name": "支付宝H5"},
		},
	},
	60: { //万家支付
		6: {
			{"code": "1", "name": "支付宝"},
			{"code": "2", "name": "微信"},
		},
	},
	61: { //"天润支付",
		4: {
			{"code": "11", "name": "支付宝"},
		},
	},
	62: { //蓉银支付
		4: {
			{"code": "wxpay", "name": "微信扫码"},
			{"code": "wxpay2", "name": "微信扫码2"},
		},
		6: {
			{"code": "wxpaywap", "name": "微信wap支付"},
			{"code": "alipay2", "name": "支付宝"},
		},
	},
	63: { //利盈支付
		4: {
			{"code": "05", "name": "QQ扫码支付"},
		},
		5: {
			{"code": "08", "name": "微信WAP支付"},
			{"code": "02", "name": "支付宝WAP支付"},
			{"code": "06", "name": "QQWAP支付"},
		},
		6: {
			{"code": "01", "name": "微信扫码支付"},

			{"code": "10", "name": "银行卡支付"},

			{"code": "11", "name": "银联扫码"},
			{"code": "12", "name": "银联快捷"},
			{"code": "13", "name": "支付宝支付"},
			{"code": "20", "name": "微信小程序"},
			{"code": "21", "name": "支付宝原生支付"},

			{"code": "22", "name": "支付宝原生扫码支付"},
			{"code": "23", "name": "微信原生扫码支付"},
			{"code": "24", "name": "微信原生WAP支付"},
		},
	},
	64: { //环球支付
		6: {
			{"code": "alipay", "name": "支付宝H5"},
			{"code": "alipay_wap", "name": "支付宝WAP"},
		},
	},
	65: { //58支付
		6: {
			{"code": "0", "name": "支付宝PC端扫码"},
			{"code": "1", "name": "支付宝手机端H5"},
			//               {"code":"2","name":"微信PC端扫码"},
			//                {"code":"3","name":"微信手机端H5"}
		},
	},
	66: { //新聚联
		6: {
			{"code": "902", "name": "微信扫码支付"},
			{"code": "903", "name": "支付宝扫码支付"},
			{"code": "904", "name": "支付宝手机"},
		},
	},
	67: { //云通付
		4: {
			{"code": "alipay", "name": "支付宝扫码支付"},
			{"code": "weixin", "name": "微信扫码支付"},
		},
		6: {
			{"code": "wxwap", "name": "微信WAP支付"},
			{"code": "alipaywap", "name": "支付宝WAP支付"},
		},
	},
	68: { //联动云付
		4: {
			{"code": "1008", "name": "支付宝扫码支付"},
		},
		6: {
			{"code": "992", "name": "支付宝H5支付"},
			{"code": "1006", "name": "微信H5"},
			{"code": "1007", "name": "微信H5D1"},
		},
	},
	69: { //顺优付
		4: {
			{"code": "001", "name": "微信wap"},
			{"code": "002", "name": "微信扫码"},
			{"code": "003", "name": "支付宝扫码"},
			{"code": "006", "name": "支付宝wap"},
			{"code": "012", "name": "银联扫码"},
			// {"code":"018","name":"银联wap"},
		},
		6: {
			{"code": "014", "name": "支付宝h5直连"},
			{"code": "015", "name": "微信h5直连"},
			{"code": "016", "name": "QQh5直连"},
			{"code": "020", "name": "微信扫码直连"},
			{"code": "021", "name": "支付宝扫码直连"},
		},
	},
	70: { //钱快支付
		4: {
			{"code": "alipay", "name": "支付宝扫码支付"},
			{"code": "weixin", "name": "微信扫码支付"},
		},
		6: {
			{"code": "alipaywap", "name": "支付宝WAP支付"},
			{"code": "wxwap", "name": "微信WAP支付"},
		},
	},
	71: { //鼎融支付
		4: {
			{"code": "902", "name": "微信扫码支付"},
			{"code": "903", "name": "支付宝扫码支付"},
			{"code": "908", "name": "QQ（钱包）扫码"},
			{"code": "910", "name": "京东钱包支付"},
		},
		6: {
			{"code": "901", "name": "微信H5支付"},
			{"code": "904", "name": "支付宝H5支付"},
			{"code": "905", "name": "QQ(钱包)H5支付"},
			{"code": "907", "name": "网银支付"},
			{"code": "911", "name": "银联在线支付"},
		},
	},
	72: { //锤子支付
		4: {
			{"code": "wechat_qrcode2", "name": "微信扫码支付"},
			{"code": "alipay_yard", "name": "支付宝扫码支付yard"},
			{"code": "alipay_qrcode2", "name": "支付宝扫码支付qrcode2"},
			{"code": "firm_alipay_qrcode", "name": "支付宝扫码支付qrcode"},
			{"code": "alipay_ttpay_1", "name": "支付宝扫码支付ttpay"},
			{"code": "alipay_online", "name": "支付宝扫码支付online"},
		},
		6: {
			{"code": "alipay_yard", "name": "支付宝H5支付yard"},
			{"code": "alipay_qrcode2", "name": "支付宝H5支付qrcode2"},
			{"code": "firm_alipay_qrcode", "name": "支付宝H5支付qrcode"},
			{"code": "alipay_ttpay_1", "name": "支付宝H5支付ttpay"},
			{"code": "alipay_online", "name": "支付宝H5支付online"},
		},
	},
	73: { //闪付支付
		4: {
			{"code": "1", "name": "微信扫码支付"},
			{"code": "2", "name": "支付宝扫码支付"},
		},
	},
	74: { //亿仟支付
		4: {
			{"code": "903", "name": "支付宝扫码支付"},
		},
		6: {
			{"code": "904", "name": "支付宝H5支付"},
		},
	},
	75: { //689支付
		6: {
			{"code": "903", "name": "支付宝扫码支付"},
			{"code": "922", "name": "微信扫码支付"},
			{"code": "921", "name": "支付宝H5支付"},
			{"code": "923", "name": "微信H5支付"},
		},
	},
	76: {
		4: {
			{"code": "wechat", "name": "微信扫码支付"},
			{"code": "alipay", "name": "支付宝扫码支付"},
			{"code": "alicardpay", "name": "支付宝银行卡支付"},
		},
	},
	77: { //DF点付支付
		6: {
			{"code": "1", "name": "微信扫码支付"},
		},
	},
	78: { //闪快支付
		6: {
			{"code": "902", "name": "微信扫码支付"},
			{"code": "903", "name": "支付宝扫码支付"},
			{"code": "908", "name": "QQ扫码支付"},
			{"code": "905", "name": "QQ手机支付"},
			{"code": "910", "name": "京东支付"},
			{"code": "904", "name": "支付宝H5"},
			{"code": "907", "name": "网银支付"},
		},
	},
	79: { //畅付通支付
		4: {
			{"code": "wechat13", "name": "微信扫码支付"},
			{"code": "alipay13", "name": "支付宝扫码支付"},
		},
	},
	80: { //德宝莱支付
		4: {
			{"code": "100001", "name": "支付宝扫码支付"},
			{"code": "100002", "name": "微信扫码支付"},
		},
		6: {
			{"code": "100001", "name": "支付宝H5支付"},
			{"code": "100002", "name": "微信H5支付"},
		},
	},
	81: { //新宝付支付
		6: {
			{"code": "0", "name": "支付宝PC支付"},
			{"code": "2", "name": "微信PC支付"},
			{"code": "1", "name": "支付宝H5支付"},
			{"code": "3", "name": "微信H5支付"},
		},
	},
	82: { //豆宝支付
		4: {
			{"code": "alipay2qr", "name": "支付宝扫码支付"},
			{"code": "wechat2qr", "name": "微信扫码支付"},
		},
		5: {
			{"code": "alipay2", "name": "支付宝Wap支付"},
			{"code": "wechat2", "name": "微信Wap支付"},
		},
	},
	83: { //MyPay支付
		4: {
			{"code": "80001", "name": "支付宝扫码支付"},
			{"code": "80002", "name": "微信扫码支付"},
			{"code": "80004", "name": "QQ钱包支付"},
			{"code": "80005", "name": "京东扫码支付"},
			{"code": "80006", "name": "百度钱包支付"},
			{"code": "80007", "name": "银联扫码支付"},
		},
		5: {
			{"code": "80009", "name": "微信Wap支付"},
			{"code": "80010", "name": "支付宝Wap支付"},
			{"code": "80011", "name": "QQWap支付"},
		},
		6: {
			{"code": "80003", "name": "网银支付"},
			{"code": "80008", "name": "快捷支付"},
		},
	},
	84: { //时代支付
		6: {
			{"code": "1", "name": "支付宝扫码"},
			{"code": "2", "name": "微信扫码"},
			{"code": "4", "name": "支付宝银行卡转账支付"},
			//{"code":"3","name":"云闪付"},
			{"code": "8", "name": "银联APP扫码"},
		},
	},
	85: { //瓜子支付
		5: {
			{"code": "alipay2", "name": "支付宝Wap支付"},
		},
		6: {
			{"code": "alipay2", "name": "支付宝H5支付"},
		},
	},
	86: { //快闪支付
		5: {
			{"code": "901", "name": "微信公众号"},
			{"code": "904", "name": "支付宝手机"},
			{"code": "905", "name": "QQ手机支付"},
		},
		6: {
			{"code": "902", "name": "微信扫码支付"},
			{"code": "903", "name": "支付宝扫码支付"},
			{"code": "907", "name": "网银支付"},
			{"code": "908", "name": "QQ扫码支付"},
		},
	},
	87: { //四海云付
		6: {
			{"code": "weixin", "name": "微信扫码支付"},
			{"code": "alipay", "name": "支付宝扫码支付"},
			{"code": "alipaywap", "name": "支付宝WAP支付"},
			{"code": "wxwap", "name": "微信wap支付"},
		},
	},
	88: { //立信支付
		6: {
			{"code": "ZFB", "name": "支付宝"},
			{"code": "ZFBWAP", "name": "支付宝WAP"},
			{"code": "WX", "name": "微信"},
			{"code": "WXH5", "name": "微H5"},
			{"code": "WY", "name": "网银"},
		},
	},
	89: { //智慧宝支付
		6: {
			{"code": "alipay", "name": "支付宝扫码"},    //支付宝扫码支付
			{"code": "alipayh5", "name": "支付宝WAP"}, //支付宝手机
		},
	},
	90: { //易富宝支付
		6: {
			{"code": "902", "name": "微信扫码支付"}, //微信扫码支付
			{"code": "903", "name": "支付宝扫码"},  //支付宝扫码支付
			{"code": "904", "name": "支付宝手机"},  //支付宝手机
			{"code": "905", "name": "QQ手机支付"}, //QQ手机支付
			{"code": "907", "name": "网银支付"},   //网银支付
			{"code": "908", "name": "QQ扫码支付"}, //QQ扫码支付
		},
	},
	91: { //星付通支付
		6: {
			{"code": "902", "name": "微信扫码支付"},    //微信扫码支付
			{"code": "903", "name": "支付宝扫码"},     //支付宝扫码支付
			{"code": "904", "name": "支付宝手机"},     //支付宝手机
			{"code": "905", "name": "QQ手机支付"},    //QQ手机支付
			{"code": "907", "name": "网银支付"},      //网银支付
			{"code": "908", "name": "银联扫码（云闪付）"}, //银联扫码（云闪付）
			{"code": "912", "name": "银联H5快捷"},    //银联H5快捷
		},
	},
	92: { //百万通支付
		6: {
			{"code": "wechat", "name": "微信"},  //微信
			{"code": "alipay", "name": "支付宝"}, //支付宝
		},
	},
	93: { //99支付
		6: {
			{"code": "0", "name": "支付宝"},
			{"code": "1", "name": "支付宝WAP"},
			{"code": "2", "name": "微信扫码"},
			{"code": "3", "name": "微信H5"},
			{"code": "5", "name": "QQ"},
			{"code": "6", "name": "QQWAP"},
			{"code": "7", "name": "京东"},
			{"code": "8", "name": "京东WAP"},
			{"code": "9", "name": "银联快捷"},
			{"code": "10", "name": "银联快捷WAP"},
			{"code": "11", "name": "银联网关"},
			{"code": "12", "name": "银联扫码"},
		},
	},
	94: { //星辰支付
		6: {
			{"code": "101", "name": "支付宝扫码"},
			{"code": "103", "name": "微信扫码"},
			{"code": "106", "name": "微信h5"},
			{"code": "104", "name": "银联扫码"},
			{"code": "102", "name": "支付宝h5"},
		},
	},
	95: { //大唐支付
		6: {
			{"code": "alipay", "name": "支付宝"},
			{"code": "wxpay", "name": "微信"},
			{"code": "qqpay", "name": "qq支付"},
			{"code": "quickpay", "name": "快捷支付"},
			{"code": "gatepay", "name": "网关支付"},
			{"code": "unionpay", "name": "银联支付"},
		},
	},
	96: { //腾云支付
		4: {
			{"code": "1", "name": "支付宝"},
			{"code": "2", "name": "微信"},
			{"code": "3", "name": "云闪付"},
		},
	},
	97: { //百威支付
		6: {
			{"code": "wx", "name": "微信"},
			{"code": "al", "name": "支付宝"},
			{"code": "qq", "name": "qq钱包"},
			{"code": "jd", "name": "京东"},
			{"code": "wy", "name": "网银支付"},
			{"code": "kj", "name": "快捷支付"},
			{"code": "yl", "name": "银联二维码"},
		},
	},
	98: { //开元支付
		4: { // 扫码
			{"code": "wxpay", "name": "微信扫码支付"},
			{"code": "alpay", "name": "支付宝扫码支付"},
			{"code": "qqpay", "name": "qq钱包扫码支付"},
			{"code": "jdpay", "name": "京东扫码支付"},
			{"code": "wypay", "name": "网银扫码支付"},
		},
		6: { // H5
			{"code": "wx", "name": "微信H5支付"},
			{"code": "al", "name": "支付宝H5支付"},
			{"code": "qq", "name": "qq钱包H5支付"},
			{"code": "jd", "name": "京东H5支付"},
			{"code": "wy", "name": "网银H5支付"},
		},
	},
	99: { //天瑞支付
		4: { // 扫码
			{"code": "WXP", "name": "微信扫码支付"},
			{"code": "ALP", "name": "支付宝扫码支付"},
		},
		6: { // H5
			{"code": "WXH5", "name": "微信H5支付"},
			{"code": "ALH5", "name": "支付宝H5支付"},
			{"code": "KJP", "name": "快捷支付"},
			{"code": "QQP", "name": "QQ钱包扫码"},
			{"code": "WYP", "name": "网银支付"},
			{"code": "YLP", "name": "银联扫码"},
			{"code": "JDP", "name": "京东扫码"},
			{"code": "JDH5", "name": "京东H5"},
		},
	},
	100: { //金砖支付
		4: { // 扫码
			{"code": "1", "name": "支付宝扫码支付"},
		},
		6: { // H5
			{"code": "1", "name": "支付宝H5支付"},
		},
	},
	101: { //大清支付
		6: {
			{"code": "901", "name": "微信公众号"},
			{"code": "902", "name": "微信扫码支付"},
			{"code": "903", "name": "支付宝扫码支付"},
			{"code": "904", "name": "支付宝手机"},
			{"code": "905", "name": "QQ手机支付"},
			{"code": "907", "name": "网银支付"},
			{"code": "908", "name": "QQ扫码支付"},
			{"code": "909", "name": "百度钱包"},
			{"code": "910", "name": "京东支付"},
		},
	},
	102: { //爱玛支付
		6: {
			{"code": "901", "name": "微信公众号"},
			{"code": "902", "name": "微信扫码支付"},
			{"code": "903", "name": "微信H5"},
			{"code": "904", "name": "支付宝扫码支付"},
			{"code": "905", "name": "支付宝H5"},
			{"code": "907", "name": "网银支付"},
			{"code": "908", "name": "网银快捷"},
			{"code": "909", "name": "银联扫码"},
			{"code": "910", "name": "QQ扫码支付"},
			{"code": "911", "name": "百度钱包"},
			{"code": "912", "name": "京东支付"},
		},
	},
	103: { //钉钉支付
		6: {
			{"code": "aliPayH5", "name": "支付宝H5"},
			{"code": "aliPaySM", "name": "支付宝扫码"},
			{"code": "wxPayH5", "name": "微信H5"},
			{"code": "wxPaySM", "name": "微信扫码"},
			{"code": "unionPaySM", "name": "银联扫码"},
			{"code": "unionPayWG", "name": "银联网关"},
			{"code": "unionPayKJ", "name": "银联快捷"},
			{"code": "jdPaySM", "name": "京东扫码"},
			{"code": "qqPayQB", "name": "QQ钱包"},
		},
	},
	104: { //七星支付
		6: {
			{"code": "ALIPAY_PC", "name": "支付宝码支付"},
			{"code": "ALIPAY_MOBILE", "name": "支付宝移动端支付"},
			{"code": "WECHAT_PC", "name": "微信扫码支付"},
			{"code": "WECHAT_MOBILE", "name": "微信移动端支付"},
			{"code": "UNION", "name": "银联支付"},
		},
	},
	105: { //喜多多支付
		6: {
			{"code": "SCANPAY_UNIONPAY", "name": "银联扫码"},
			{"code": "SCANPAY_WEIXIN", "name": "微信扫码支付"},
			{"code": "SCANPAY_ALIPAY", "name": "支付宝扫码支付"},
			{"code": "SCANPAY_QQ", "name": "QQ钱包扫码支付"},
			{"code": "H5_QQ", "name": "QQH5"},
			{"code": "ALIPAY_H5", "name": "支付宝H5"},
			{"code": "H5_WEIXIN", "name": "微信H5"},
		},
	},
	106: { //雀付支付
		6: {
			{"code": "918", "name": "农信易扫918"},
			{"code": "919", "name": "微信扫码919"},
			{"code": "920", "name": "微信H5920"},
			{"code": "921", "name": "微信扫码921"},
			{"code": "922", "name": "微信H5922"},
			{"code": "923", "name": "支付宝扫码923"},
			{"code": "924", "name": "支付宝H5924"},
		},
	},
	107: { //一加支付
		6: { // H5
			{"code": "wap", "name": "支付宝H5"},
			{"code": "qrcode", "name": "支付宝扫码"},
			{"code": "wxwap", "name": "微信H5"},
			{"code": "wxqrcode", "name": "微信扫码"},
			{"code": "ylkj", "name": "银联快捷"},
			{"code": "ylwg", "name": "银联网关"},
			{"code": "ylsm", "name": "银联扫码"},
			{"code": "qqqb", "name": "QQ钱包"},
		},
	},
	108: { //汇宏支付
		6: { // H5
			{"code": "961", "name": "支付宝H5"},
			{"code": "962", "name": "微信H5"},
		},
	},
	109: { //GTPAY支付
		6: {
			{"code": "26005", "name": "微信H5"},
			{"code": "26001", "name": "支付宝H5"},
			{"code": "26015", "name": "快捷支付"},
			{"code": "26060", "name": "银联支付"},
			{"code": "26085", "name": "网银支付"},
		},
	},
	110: { //天使支付
		6: {
			{"code": "ZFB", "name": "支付宝"},
			{"code": "ZFBH5", "name": "支付宝H5"},
			{"code": "ZFBSM", "name": "支付宝扫码"},
			{"code": "WX", "name": "微信"},
			{"code": "WXYS", "name": "微信原生"},
			{"code": "WXH5", "name": "微信H5"},
		},
	},
	111: { //火箭支付
		6: {
			{"code": "101", "name": "支付宝扫码"},
			{"code": "102", "name": "支付宝h5"},
			{"code": "103", "name": "微信扫码"},
			{"code": "106", "name": "微信h5"},
		},
	},
	112: { // 永仁支付
		6: {
			{"code": "WXWAP", "name": "微信H5"},
			{"code": "ZFBWAP", "name": "支付宝H5"},
			{"code": "QQWAP", "name": "QQ钱包H5"},
			{"code": "UNIONWAP", "name": "银联H5"},
			{"code": "JDWAP", "name": "京东H5"},
		},
	},
	113: { // 龙支付
		6: {
			{"code": "bank", "name": "银行通道"},
			{"code": "alipay", "name": "支付宝"},
			{"code": "wechat", "name": "微信"},
			{"code": "yqb", "name": "壹钱包聚合码"},
		},
	},
	114: { // 百家付
		6: {
			{"code": "wechat", "name": "微信"},
			{"code": "alipay", "name": "支付宝"},
		},
	},
}
