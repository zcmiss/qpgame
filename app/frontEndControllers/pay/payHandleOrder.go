package pay

import (
	"encoding/json"
	"net/http"
	"net/url"
	"qpgame/common/utils"
	"qpgame/config"
	"qpgame/models/xorm"
	"qpgame/ramcache"
	"strconv"
	"strings"
	"time"
)

//如果支付成功第三方会在客户端调用这个回调
const ClientPayResultUrlNotifyRouter = "/api/v1/tpPayClientNotify/"

//第三方支付异步回调
const tppayCallBack = "/api/v1/tppayCallBack/"

//支付涉及到参数
type payParam struct {
	platForm     int                 //每个支付的唯一代码码
	payAmount    float64             //支付金额
	payBankcode  string              //支付码
	credentialId int                 //paycreadinal表id
	userIdS      string              //用户id
	addrType     int                 //充值方式1-6
	orderId      string              //订单号
	rowCred      xorm.PayCredentials //支付证书数据
}

type Pay struct {
	platform        string        //运营平台标识
	httpRequest     *http.Request //用于获取ip
	payParams       payParam      //paycredentials表对应Id行数据
	returnClientUrl string        //第三方支付成功同步回调
	notifyUrl       string        //第三方支付成功异步回调
	apiAddress      string        //运营平台接口地址
}

func NewPay(platform string, payParams payParam) *Pay {
	pay := new(Pay)
	pay.platform = platform
	//反射性能很差，以空间换性能
	payC, _ := ramcache.TablePayCredential.Load(platform)
	//获取paycredentials表对应Id数据
	rowCred := payC.(map[int]xorm.PayCredentials)[payParams.credentialId]
	payParams.rowCred = rowCred
	pay.payParams = payParams
	pay.configBackCallbackAddress()
	return pay
}

//配置回调地址
func (cthis *Pay) configBackCallbackAddress() {
	mainTablePlatform := ramcache.GetMainDbPlatform(cthis.platform)
	payBack := mainTablePlatform.PayBackAddress
	cthis.apiAddress = mainTablePlatform.ApiAddress
	//客户端同步地址回调，在第三方支付成功之后会在客户端调用这个接口,有些第三方会调用,有些不会调用
	platFormS := strconv.Itoa(cthis.payParams.platForm) //支付唯一代号
	credentialIdS := strconv.Itoa(cthis.payParams.credentialId)
	//同步回调
	returnClientUrl := payBack + ClientPayResultUrlNotifyRouter + platFormS + "/" + credentialIdS + "/" + cthis.payParams.userIdS
	//异步回调地址
	notifyUrl := payBack + tppayCallBack + platFormS + "/" + credentialIdS + "/" + cthis.payParams.userIdS
	cthis.returnClientUrl = returnClientUrl
	cthis.notifyUrl = notifyUrl
}

//分发支付通道
func (cthis *Pay) PayDispathChannal() (map[string]interface{}, bool) {
	switch cthis.payParams.platForm {

	//蓉银支付
	case 62:
		return cthis.tpPay62YingYingZhiFu()
	//利盈支付
	case 63:
		return cthis.tpPay63LiYingZhiFu()
	//全通云付
	case 67:
		return cthis.tpPay67QuanTongYunFu()
	//联动支付
	case 68:
		return cthis.tpPay68LianDongZhiFu()
	//立信支付
	case 88:
		return cthis.tpPay88LiXinZhiFu()
	//智慧宝支付
	case 89:
		return cthis.tpPay89ZiHuiBaoZhiFu()
	//金砖支付
	case 100:
		return cthis.tpPay100JinZhuan()
	//钉钉支付
	case 103:
		return cthis.tpPay103DingDing()
		//喜多多支付
	case 105:
		return cthis.tpPay105XiDuoDuo()
	}
	//无效平台返回
	return map[string]interface{}{
		"message": "支付失败,该支付未接入",
		"code":    config.PAYACTIONERROR,
	}, false
}

//蓉银支付
func (cthis *Pay) tpPay62YingYingZhiFu() (map[string]interface{}, bool) {
	configPay := config.TppayConfig[cthis.payParams.platForm]
	payParams := cthis.payParams
	var chargeType, appType int
	switch payParams.payBankcode {
	case "wxpay", "wxpay2", "wxpaywap": //微信扫码 微信扫码2 微信wap支付
		appType = 2
		chargeType = 1
	case "alipay2": //支付宝
		appType = 3
		chargeType = 3
	default:
		appType = 3
		chargeType = 1
	}
	orderdata := map[string]string{
		"outOrderNo":   payParams.orderId,
		"tradeAmount":  floatToString(payParams.payAmount),
		"payCode":      payParams.payBankcode,
		"goodsClauses": payParams.rowCred.MerchantNumber,
		"code":         payParams.rowCred.MerchantNumber,
		"notifyUrl":    cthis.notifyUrl,
	}

	//请求数据
	var str, qrcode string
	// 签名
	str = payKsortAddstr(orderdata, 2)
	orderdata["sign"] = utils.Php2GoMd5(str + "key=" + payParams.rowCred.PrivateKey)
	tjurl := configPay["tjurl"].(string)
	rtn := utils.ReqPost(tjurl, orderdata, time.Duration(10*time.Second))
	var bodyJson = make(map[string]interface{})
	json.Unmarshal(rtn, &bodyJson)
	status, existStatus := bodyJson["payState"]

	if existStatus && status == "success" { //成功发起支付
		urlS, isUrl := bodyJson["url"]
		if isUrl {
			if payParams.addrType == 4 {
				qrcode = cthis.apiAddress + "/qrcode/" + url.QueryEscape(utils.Base64UrlSafeEncode([]byte(urlS.(string))))
			} else {
				qrcode = urlS.(string)
			}
			return map[string]interface{}{
				"qrcode":      qrcode,
				"app_type":    appType,
				"charge_type": chargeType,
			}, true
		}
	}
	msg, isMsg := bodyJson["message"]
	if isMsg {
		msg = bodyJson["message"].(string)
	} else {
		msg = "支付失败"
	}
	return map[string]interface{}{
		"message": msg,
		"code":    config.PAYACTIONERROR,
	}, false
}

//利盈支付
func (cthis *Pay) tpPay63LiYingZhiFu() (map[string]interface{}, bool) {
	configPay := config.TppayConfig[cthis.payParams.platForm]
	payParams := cthis.payParams
	var chargeType, appType int
	switch payParams.payBankcode {
	case "05", "06": //QQ钱包扫码 QQWAP
		appType = 1
		chargeType = 1
	case "01", "08", "20": //微信扫码 微信WAP支付 微信小程序
		appType = 2
		chargeType = 1
	case "02", "13", "21": //支付宝
		appType = 3
		chargeType = 3
	case "10", "11", "12": //银行卡 银联扫码 网银快捷
		appType = 5
		chargeType = 3
	default:
		appType = 3
		chargeType = 3
	}

	orderdata := map[string]string{
		"mch_id":       payParams.rowCred.MerchantNumber, //商户uid
		"trade_type":   payParams.payBankcode,
		"out_trade_no": payParams.orderId,                        //订单号
		"total_fee":    floatToString(payParams.payAmount * 100), //价格 浮点转字符串 转int *100 转字符串
		"bank_id":      "",                                       //支付银行
		"notify_url":   url.QueryEscape(cthis.notifyUrl),         //异步通知
		"return_url":   url.QueryEscape(cthis.returnClientUrl),   //成功返回地址
		"time_start":   time.Now().Format("20060102150405"),      //订单号生成时间
		"nonce_str":    utils.Php2GoMd5(payParams.userIdS + payParams.orderId + strconv.Itoa(int(time.Now().Unix()))),
	}
	// 签名
	str := payKsortAddstr(orderdata, 8)
	orderdata["sign"] = strings.ToUpper(utils.Php2GoMd5(str + "key=" + payParams.rowCred.PrivateKey))
	orderdata["notify_url"] = cthis.notifyUrl
	orderdata["return_url"] = cthis.returnClientUrl
	orderdata["body"] = "余额充值"
	orderdata["attach"] = "用户余额充值"
	if payParams.addrType == 4 {
		tjurl := configPay["tjurl"].(map[string]string)["1"]
		data := utils.ReqPost(tjurl, orderdata, time.Duration(10*time.Second))
		var bodyJson = make(map[string]interface{})
		json.Unmarshal(data, &bodyJson)
		status, existStatus := bodyJson["result_code"]

		if existStatus && status == "SUCCESS" {
			qrcode := cthis.apiAddress + "/qrcode/" + url.QueryEscape(bodyJson["code_url"].(string))
			return map[string]interface{}{
				"qrcode":      qrcode,
				"app_type":    appType,
				"charge_type": chargeType,
			}, true
		}
	} else {
		tjurl := configPay["tjurl"].(map[string]string)["1"]
		orderdata["wap_url"] = tjurl
		tpjurl := cthis.apiAddress + "/tpwap"
		qrcode := utils.UrlSplitKeyValueOnlyHttp(tpjurl, orderdata, true, true)
		return map[string]interface{}{
			"qrcode":      qrcode,
			"app_type":    appType,
			"charge_type": chargeType,
		}, true
	}
	return map[string]interface{}{
		"message": "支付失败",
		"code":    config.PAYACTIONERROR,
	}, false
}

//全通云付
func (cthis *Pay) tpPay67QuanTongYunFu() (map[string]interface{}, bool) {
	configPay := config.TppayConfig[cthis.payParams.platForm]
	payParams := cthis.payParams
	var chargeType, appType int
	switch payParams.payBankcode {
	case "weixin", "weixincode", "wxwap": //微信扫码支付 微信wap支付 微信wap支付
		appType = 2
		chargeType = 1
	case "bdpay", "yinlian": //银联扫码支付 快捷支付
		appType = 5
		chargeType = 3
	case "jdpay": //京东扫码支付
		appType = 4
		chargeType = 3
	case "tenpaywap", "qqmobile": //QQ支付 QQ扫码
		appType = 1
		chargeType = 1
	case "alipay", "alipaywap", "alipaycode": //支付宝扫码 支付宝WAP支付 支付宝付款码支付
		appType = 3
		chargeType = 3
	default:
		appType = 3
		chargeType = 3
	}
	orderdata := map[string]string{
		"p0_Cmd":          "Buy",                              //业务类型
		"p1_MerId":        payParams.rowCred.MerchantNumber,   //商户uid
		"p2_Order":        payParams.orderId,                  //订单号
		"p3_Amt":          floatToString(payParams.payAmount), //价格
		"p4_Cur":          "CNY",                              //交易币种
		"p5_Pid":          "VIP",                              //商品名称
		"pa_MP":           strconv.Itoa(int(time.Now().Unix())),
		"p6_Pcat":         "virtual",             //商品种类
		"p7_Pdesc":        "desc",                //商品描述
		"p8_Url":          cthis.notifyUrl,       //商户接收支付成功数据的地址
		"pd_FrpId":        payParams.payBankcode, //支付通道编码
		"pr_NeedResponse": "1",
	}
	// 签名
	md5str := payKsortAddstr(orderdata, 1)
	sign := hmacMd5(md5str, payParams.rowCred.PrivateKey)
	orderdata["hmac"] = sign
	tjurl := configPay["tjurl"].(string)
	rtn := utils.ReqPost(tjurl, orderdata, time.Duration(10*time.Second))
	var bodyJson = make(map[string]interface{})
	json.Unmarshal(rtn, &bodyJson)
	status, existStatus := bodyJson["status"]
	if existStatus && status == "0" {
		qrcode := bodyJson["payImg"].(string)
		if payParams.addrType == 4 {
			qrcode = cthis.apiAddress + "/qrcode/" + url.QueryEscape(qrcode)
		}
		return map[string]interface{}{
			"qrcode":      qrcode,
			"app_type":    appType,
			"charge_type": chargeType,
		}, true
	}
	return map[string]interface{}{
		"message": "支付失败",
		"code":    config.PAYACTIONERROR,
	}, false
}

//联动支付
func (cthis *Pay) tpPay68LianDongZhiFu() (map[string]interface{}, bool) {
	configPay := config.TppayConfig[cthis.payParams.platForm]
	payParams := cthis.payParams
	var chargeType, appType int
	switch payParams.payBankcode {
	case "1006", "1007", "1012", "1004": //微信H5 微信H5D1 微信条码 微信扫码支付
		appType = 2
		chargeType = 1
	case "1014", "1009", "1015": //银联H5 银联扫码支付 银联条码4
		appType = 5
		chargeType = 3
	case "1016", "1017", "1018": //京东H5 京东扫码 京东条码
		appType = 4
		chargeType = 3
	case "1010", "1005", "1013": //QQH5 QQ扫码 QQ条码
		appType = 1
		chargeType = 1
	case "992", "1008", "1011": //支付宝H5 支付宝扫码 支付宝条码
		appType = 3
		chargeType = 3
	case "1025", "1026", "1027": //百度钱包H5 百度钱包二维码 百度钱包条码
		appType = 3
		chargeType = 6
	default:
		appType = 3
		chargeType = 3
	}
	orderdata := map[string]string{
		"parter":      payParams.rowCred.MerchantNumber, //商户uid
		"type":        payParams.payBankcode,
		"value":       utils.NumberFormat(payParams.payAmount, 2, ".", ""), //价格
		"orderid":     payParams.orderId,                                   //订单号
		"callbackurl": cthis.notifyUrl,                                     //商户接收支付成功数据的地址
	}
	// 签名
	md5str := payKsortAddstr(orderdata, 1)
	sign := strings.ToLower(utils.Php2GoMd5(md5str + payParams.rowCred.PrivateKey))
	orderdata["sign"] = sign
	orderdata["attach"] = "余额充值"
	orderdata["hrefbackurl"] = cthis.returnClientUrl
	tjurl := configPay["tjurl"].(string)
	qrcode := utils.UrlSplitKeyValue(tjurl, orderdata, true)
	if payParams.addrType == 4 {
		qrcode = cthis.apiAddress + "/qrcode/" + url.QueryEscape(qrcode)
	}
	return map[string]interface{}{
		"qrcode":      qrcode,
		"app_type":    appType,
		"charge_type": chargeType,
	}, true
}

//立信支付
func (cthis *Pay) tpPay88LiXinZhiFu() (map[string]interface{}, bool) {
	configPay := config.TppayConfig[cthis.payParams.platForm]
	payParams := cthis.payParams
	var chargeType, appType int
	switch payParams.payBankcode {
	case "ZFB", "ZFBWAP": //支付宝扫码支付 支付宝手机
		appType = 3
		chargeType = 2
	case "WX", "WXH5": //微信
		appType = 2
		chargeType = 1
	case "WY": //网银
		appType = 5
		chargeType = 4
	default:
		appType = 3
		chargeType = 2
	}
	postData := map[string]string{
		"mer_num":    payParams.rowCred.MerchantNumber,
		"pay_way":    payParams.payBankcode,
		"money":      floatToString(payParams.payAmount * 100),
		"order_num":  payParams.orderId,
		"goods_name": "在线充值",
		"notify_url": cthis.notifyUrl,
		"return_url": cthis.returnClientUrl,
	}
	if payParams.payBankcode == "WX" || payParams.payBankcode == "WXH5" {
		postData["version"] = "1.0"
	} else {
		postData["version"] = "3.0"
	}
	signStr := `{mer_num}&{pay_way}&{money}&{order_num}&{goods_name}&{notify_url}&{return_url}&{version}`
	signStr = strings.Replace(signStr, "{mer_num}", postData["mer_num"], -1)
	signStr = strings.Replace(signStr, "{pay_way}", postData["pay_way"], 1)
	signStr = strings.Replace(signStr, "{money}", postData["money"], 1)
	signStr = strings.Replace(signStr, "{order_num}", postData["order_num"], 1)
	signStr = strings.Replace(signStr, "{goods_name}", postData["goods_name"], 1)
	signStr = strings.Replace(signStr, "{notify_url}", postData["notify_url"], 1)
	signStr = strings.Replace(signStr, "{return_url}", postData["return_url"], 1)
	signStr = strings.Replace(signStr, "{version}", postData["version"], 1)
	postData["sign"] = strings.ToUpper(utils.Php2GoMd5(signStr + "&" + payParams.rowCred.PrivateKey))
	postData["pattern"] = "form"
	tjurl := configPay["tjurl"].(string)
	postData["wap_url"] = tjurl
	tpjurl := cthis.apiAddress + "/tpwap"
	qrcode := utils.UrlSplitKeyValueOnlyHttp(tpjurl, postData, true, true)
	return map[string]interface{}{
		"qrcode":      qrcode,
		"app_type":    appType,
		"charge_type": chargeType,
	}, true
}

//智慧宝支付
func (cthis *Pay) tpPay89ZiHuiBaoZhiFu() (map[string]interface{}, bool) {
	configPay := config.TppayConfig[cthis.payParams.platForm]
	payParams := cthis.payParams
	var chargeType, appType int

	switch payParams.payBankcode {
	case "alipay", "alipayh5": //支付宝扫码支付 支付宝手机
		appType = 3
		chargeType = 3
	default:
		appType = 3
		chargeType = 3
	}
	post_data := map[string]string{
		"shid": payParams.rowCred.MerchantNumber,                    //商户ID
		"bb":   "1.0",                                               //版本
		"zftd": payParams.payBankcode,                               //支付通道
		"ddh":  payParams.orderId,                                   //订单号
		"je":   utils.NumberFormat(payParams.payAmount, 2, ".", ""), //金额(必须保留两位小数点，否则验签失败)
		"ddmc": "在线充值",                                              //订单名称
		"ddbz": "备注",                                                //订单备注
		"ybtz": cthis.notifyUrl,                                     //异步通知地址
		"tbtz": cthis.returnClientUrl,                               //同步通知地址
	}

	str := `shid=(商户ID)&bb=1.0&zftd=（支付通道）&ddh=（订单号）&je=（金额）&ddmc=（订单名称）&ddbz=（备注）&ybtz=（异步通知地址）&tbtz=（同步通知地址）&（用户KEY）`
	str = strings.Replace(str, "(商户ID)", post_data["shid"], -1)
	str = strings.Replace(str, "（支付通道）", post_data["zftd"], -1)
	str = strings.Replace(str, "（订单号）", post_data["ddh"], -1)
	str = strings.Replace(str, "（金额）", post_data["je"], -1)
	str = strings.Replace(str, "（订单名称）", post_data["ddmc"], -1)
	str = strings.Replace(str, "（备注）", post_data["ddbz"], -1)
	str = strings.Replace(str, "（异步通知地址）", post_data["ybtz"], -1)
	str = strings.Replace(str, "（同步通知地址）", post_data["tbtz"], -1)
	str = strings.Replace(str, "（用户KEY）", payParams.rowCred.PrivateKey, -1)
	post_data["sign"] = utils.Php2GoMd5(str)
	var qrcode, tjurl, tpjurl string
	tjurl = configPay["tjurl"].(string)
	post_data["wap_url"] = tjurl
	tpjurl = cthis.apiAddress + "/tpwap"
	qrcode = utils.UrlSplitKeyValueOnlyHttp(tpjurl, post_data, true, true)
	return map[string]interface{}{
		"qrcode":      qrcode,
		"app_type":    appType,
		"charge_type": chargeType,
	}, true
}

//金砖支付
func (cthis *Pay) tpPay100JinZhuan() (map[string]interface{}, bool) {
	configPay := config.TppayConfig[cthis.payParams.platForm]
	payParams := cthis.payParams
	postData := map[string]string{
		"mch_code":     payParams.rowCred.MerchantNumber,
		"mch_trade_no": payParams.orderId,                    // 订单编号
		"amount":       floatToString(payParams.payAmount),   // 订单金额
		"pay_type":     payParams.payBankcode,                // 支付类型  1: 支付宝跳转扫码
		"timespan":     strconv.Itoa(int(time.Now().Unix())), // 订单时间
		"param":        "金砖支付",                               // 订单名称
		"notify_url":   cthis.notifyUrl,                      // 回调地址
		"return_url":   cthis.returnClientUrl,                // 显示地址
	}
	// 签名
	str := payKsortAddstr(postData, 1)
	str = str + "&key=" + payParams.rowCred.PrivateKey
	postData["sign"] = utils.Php2GoMd5(str)
	orderUrl := configPay["orderUrl"].(string)
	data := utils.ReqPost(orderUrl, postData, time.Duration(10*time.Second))
	var bodyJson = make(map[string]interface{})
	json.Unmarshal(data, &bodyJson)
	success, existStatus := bodyJson["success"]
	if existStatus && success == true {
		result, existData := bodyJson["data"].(map[string]interface{})

		if existData {
			appType := 3 // 支付类型  1: 支付宝扫码|H5支付通用
			chargeTy := 2
			qrcode := result["url"].(string)
			qrcode = cthis.apiAddress + "/qrcode/" + url.QueryEscape(qrcode)
			return map[string]interface{}{
				"qrcode":      qrcode,
				"app_type":    appType,
				"charge_type": chargeTy,
			}, true
		}
	} else {
		msg, existMessage := bodyJson["message"]
		if existMessage {
			return map[string]interface{}{
				"message": msg,
				"code":    config.PAYACTIONERROR,
			}, false
		}
	}
	return map[string]interface{}{
		"message": "支付失败",
		"code":    config.PAYACTIONERROR,
	}, false
}

//钉钉支付
func (cthis *Pay) tpPay103DingDing() (map[string]interface{}, bool) {
	payParams := cthis.payParams
	configPay := config.TppayConfig[payParams.platForm]
	var chargeType, appType int
	switch payParams.payBankcode {
	case "wxPayH5", "wxPaySM": //微信H5 ://微信扫码
		appType = 2
		chargeType = 1
	case "aliPayH5", "aliPaySM": //⽀付宝H5 ://⽀付宝扫码
		appType = 3
		chargeType = 1
	case "qqPayQB": //QQ钱包
		appType = 1
		chargeType = 1
	case "unionPaySM", "unionPayWG", "unionPayKJ": //银联扫码 : //银联⽹关 : //银联快捷
		appType = 5
		chargeType = 4
	case "jdPaySM": //京东扫码
		appType = 4
		chargeType = 3
	default:
		appType = 3
		chargeType = 1
	}
	postData := map[string]string{
		"version":     "v1.0",                                   //接⼝版本号
		"type":        payParams.payBankcode,                    //接⼝类型
		"userId":      payParams.rowCred.MerchantNumber,         //商户编号
		"requestNo":   payParams.orderId,                        //商户流⽔号
		"amount":      floatToString(payParams.payAmount * 100), //订单⾦额
		"callBackURL": cthis.notifyUrl,                          //异步回调地址
		"redirectUrl": cthis.returnClientUrl,                    //前台回调地址
	}

	strSign := payKsortAddstr(postData, 8)
	postData["sign"] = strings.ToLower(utils.Php2GoMd5(strSign + "key=" + payParams.rowCred.PrivateKey))
	postData["attachData"] = "123"
	tjurl := configPay["tjurl"].(string)

	postDataJson,_ := json.Marshal(postData)
	data := utils.CurlPostJson(tjurl, string(postDataJson), time.Duration(10*time.Second))
	var bodyJson = make(map[string]interface{})
	json.Unmarshal(data, &bodyJson)
	message, existMessage := bodyJson["message"]
	postSign, existPostsign := bodyJson["sign"]
	qrcode, exitQrocde := bodyJson["payUrl"]
	if existMessage && message == "000000" && existPostsign && exitQrocde {
		str := payKsortMapInterface(bodyJson, 103)
		str = str + "key=" + payParams.rowCred.PrivateKey
		sign := strings.ToLower(utils.Php2GoMd5(str))
		if sign == postSign {
			return map[string]interface{}{
				"qrcode":      qrcode,
				"app_type":    appType,
				"charge_type": chargeType,
			}, true
		}
	}

	if existMessage {
		return map[string]interface{}{
			"message": message,
			"code":    config.PAYACTIONERROR,
		}, false
	}
	return map[string]interface{}{
		"message": "支付失败",
		"code":    config.PAYACTIONERROR,
	}, false
}

//喜多多支付
func (cthis *Pay) tpPay105XiDuoDuo() (map[string]interface{}, bool) {
	payParams := cthis.payParams
	configPay := config.TppayConfig[payParams.platForm]
	var chargeType, appType int
	var payType string
	switch payParams.payBankcode {
	case "SCANPAY_WEIXIN", "H5_WEIXIN": //微信扫码支付 微信H5
		appType = 2
		chargeType = 1
		payType = "WEIXIN"
	case "SCANPAY_ALIPAY", "ALIPAY_H5": //支付宝码支付 支付宝移动端支付
		appType = 3
		chargeType = 1
		payType = "ALIPAY"
	case "H5_QQ", "SCANPAY_QQ": //QQ钱包 QQ钱包扫码支付
		appType = 1
		chargeType = 1
		payType = "QQPAY"
	case "SCANPAY_UNIONPAY": //银联扫码
		appType = 3
		chargeType = 6
		payType = "UNIONPAY"
	default:
		appType = 3
		chargeType = 1
		payType = "ALIPAY"
	}
	ip := utils.GetIp(cthis.httpRequest)
	postData := map[string]string{
		"merAccount":  payParams.rowCred.CredentialKey,          //商户标识，由系统随机生成
		"customerNo":  payParams.rowCred.MerchantNumber,         //用户编号，由系统生成
		"payType":     payType,                                  //仅支持MD5
		"payTypeCode": payParams.payBankcode,                    //支付类型，参考支付类型表
		"orderNo":     payParams.orderId,                        //商户订单号，由商户自行生成，必须唯一
		"time":        strconv.Itoa(int(time.Now().Unix())),     //时间戳，例如：1510109185，精确到秒，前后误差不超过5分钟
		"payAmount":   floatToString(payParams.payAmount * 100), //支付金额，单位分，必须大于0
		"productCode": "01",                                     //商品类别码，固定值01
		"productName": "VipPay",                                 //商品名称
		"productDesc": "iphone",                                 //商品描述
		"userType":    "0",                                      //用户类型，固定值0
		"payIp":       ip,                                       //用户IP地址
		"returnUrl":   cthis.returnClientUrl,                    //页面通知地址
		"notifyUrl":   cthis.notifyUrl,                          //异步通知地址
		"signType":    "2",                                      //固定值传2
	}
	strSign := payKsortAddstr(postData, 2)
	postData["sign"] = utils.StrSha1(strSign + "key=" + payParams.rowCred.PrivateKey)
	tjurl := configPay["tjurl"].(string)
	postTemp, _ := json.Marshal(postData)
	postMap := map[string]string{
		"merAccount": postData["merAccount"],
		"data":       string(postTemp),
	}
	data := utils.ReqPost(tjurl, postMap, time.Duration(10*time.Second))
	var bodyJson = make(map[string]interface{})
	json.Unmarshal(data, &bodyJson)

	code, existStatus := bodyJson["code"]
	resData, exitsData := bodyJson["data"].(map[string]interface{})
	if existStatus && code == "000000" && exitsData {
		payUrl, exitsPayUrl := resData["payUrl"]
		if exitsPayUrl {
			qrcode := payUrl
			return map[string]interface{}{
				"qrcode":      qrcode,
				"app_type":    appType,
				"charge_type": chargeType,
			}, true
		}
	}
	msg, exitsMsg := bodyJson["msg"]
	if !exitsMsg {
		msg = "支付失败"
	}
	return map[string]interface{}{
		"message": msg,
		"code":    config.PAYACTIONERROR,
	}, false
}
