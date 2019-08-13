package utils

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/mojocn/base64Captcha"
)

//生成随机数，用于验证码
func GetVcode() string {
	//if true {
	//	return "111111"
	//}
	//6位数验证码
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return vcode
}

func CreateCaptchaNum() (string, string) {
	// @updated by aTian 修改验证码样式
	var config = base64Captcha.ConfigCharacter{
		Height: 40,
		Width:  100,
		//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
		Mode:               base64Captcha.CaptchaModeNumber,
		ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
		ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
		IsUseSimpleFont:    true,
		IsShowHollowLine:   true,
		IsShowNoiseDot:     true,
		IsShowNoiseText:    true,
		IsShowSlimeLine:    false,
		IsShowSineLine:     false,
		CaptchaLen:         4,
	}
	//创建数字验证码.
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	_, cap := base64Captcha.GenerateCaptcha("", config)
	ci := cap.(*base64Captcha.CaptchaImageChar)
	//以base64编码
	base64String := base64Captcha.CaptchaWriteToBase64Encoding(cap)
	return ci.VerifyValue, base64String
}

// 创建验证码(base64code) - 可以有几种方式: 数字/声音/公式
func CreateCaptcha() (string, string) {

	//声音验证码配置
	//var configA = base64Captcha.ConfigAudio{
	//	CaptchaLen: 6,
	//	Language:   "zh",
	//}
	//创建声音验证码
	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
	//idKeyA, capA := base64Captcha.GenerateCaptcha("", configA)
	//以base64编码
	//base64stringA := base64Captcha.CaptchaWriteToBase64Encoding(capA)

	codeCharacter := func() (string, string) { //字符,公式,验证码配置
		current := time.Now().Second() % 4
		codeMode := base64Captcha.CaptchaModeNumber
		//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
		switch current {
		case 0:
			codeMode = base64Captcha.CaptchaModeAlphabet
		case 1:
			codeMode = base64Captcha.CaptchaModeNumberAlphabet
		case 2:
			codeMode = base64Captcha.CaptchaModeArithmetic
		}
		var config = base64Captcha.ConfigCharacter{
			Height:             40,
			Width:              100,
			Mode:               codeMode,
			ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
			ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
			IsShowHollowLine:   false,
			IsShowNoiseDot:     false,
			IsShowNoiseText:    false,
			IsShowSlimeLine:    false,
			IsShowSineLine:     false,
			CaptchaLen:         4,
		}
		//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
		idKey, cap := base64Captcha.GenerateCaptcha("", config)

		//以base64编码
		base64String := base64Captcha.CaptchaWriteToBase64Encoding(cap)
		return idKey, base64String
	}

	codeNum := func() (string, string) { //数字验证码
		var config = base64Captcha.ConfigDigit{
			Height:     40,
			Width:      100,
			MaxSkew:    0.7,
			DotCount:   80,
			CaptchaLen: 4,
		}
		//创建数字验证码.
		//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
		idKey, cap := base64Captcha.GenerateCaptcha("", config)

		//以base64编码
		base64String := base64Captcha.CaptchaWriteToBase64Encoding(cap)
		return idKey, base64String
	}

	if time.Now().Second()%2 == 0 { //随机生成验证码
		return codeNum()
	}
	return codeCharacter()
}
