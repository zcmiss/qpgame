package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"qpgame/config"
	"qpgame/ramcache"
	"strconv"
)

func SendSms(phoneNum string, platform string) (string, error) {
	vcode := GetVcode()
	if config.DevloperDebug == false {
		go httpSend(phoneNum, vcode)
	}
	intvcode, _ := strconv.Atoi(vcode)
	phoneArr := [2]int{intvcode, GetNowTime() + 180}
	PhccPlat := make(map[string][2]int)
	//如果平台容器还不存在就创建一个map
	PhCC, existPlat := ramcache.PhoneCheckCode.Load(platform)
	if !existPlat {
		PhCC = PhccPlat
	}
	//如果发送成功，则将验证码缓存
	PhCC.(map[string][2]int)[phoneNum] = phoneArr
	ramcache.PhoneCheckCode.Store(platform, PhCC)
	return vcode, nil
}

func httpSend(mobile string, code string) {
	// 修改为您的apikey(https://www.yunpian.com)登录官网后获取
	apikey := "96334c7f4cb90f5aca36bf7c6cdd055a"
	// 发送模板编号
	tpl_id := 2801176
	// 发送模板内容
	tpl_value := url.Values{"#code#": {code}}.Encode()
	// 指定模板发送短信url
	url_tpl_sms := "https://sms.yunpian.com/v2/sms/tpl_single_send.json"
	data_tpl_sms := url.Values{"apikey": {apikey}, "mobile": {mobile},
		"tpl_id": {fmt.Sprintf("%d", tpl_id)}, "tpl_value": {tpl_value}}
	httpsPostForm(url_tpl_sms, data_tpl_sms)
}

func httpsPostForm(url string, data url.Values) {
	resp, err := http.PostForm(url, data)

	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(string(body))

}
