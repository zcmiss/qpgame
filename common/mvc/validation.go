package mvc

import (
	"qpgame/common/utils"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/kataras/iris"
)

/********************************************** 以下, 接口定义 ************************************/
// 数据校验器-实现接口
type IValidator interface {
	Validate(ctx *iris.Context) (string, bool)
}

/********************************************** 以下, 校验证/结构体 基本类型实现 *********************/
// 数据校验器-结构体-实现
type Validation struct {
	ctx           *iris.Context
	post          utils.PostData
	messageLength uint16   //错误信息的长度
	messages      []string //错误信息
}

//创建新的验证器实例
func NewValidation(ctx *iris.Context) *Validation {
	post := utils.GetPostData(ctx)
	return &Validation{ctx: ctx, post: post}
}

// 是否在指定的长度之间
func (self *Validation) StringLength(field string, message string, min int, max int) *Validation {
	value := self.post.Get(field)
	count := utf8.RuneCountInString(string(value)) //现有的字符的数量
	if count < min || count > max {
		self.messages = append(self.messages, message)
	}

	return self
}

// 字符串必须为数字数组，英文逗号拼接
func (self *Validation) IsIntegers(field string, message string) *Validation {
	value := self.post.Get(field)
	match,_ := regexp.Match("^[1-9][0-9]*(,[1-9][0-9]*)*$", []byte(value))
	if !match {
		self.messages = append(self.messages, message)
	}

	return self
}

// 字符串长度必须为固定长度
func (self *Validation) StringLengthEquals(field string, message string, length int) *Validation {
	value := self.post.Get(field)
	count := utf8.RuneCountInString(string(value)) //现有的字符的数量
	if count != length {
		self.messages = append(self.messages, message)
	}

	return self
}

// 字符串必须与设定的字符中相同
func (self *Validation) StringEquals(field string, message string, equal string) *Validation {
	value := self.post.Get(field)
	if strings.Compare(string(value), equal) != 0 {
		self.messages = append(self.messages, message)
	}

	return self
}

// 必须在包含的字符串范围之内
func (self *Validation) InStrings(field string, message string, stringArray *[]string) *Validation {
	value := self.post.Get(field)
	origin := string(value)
	for _, v := range *stringArray {
		if strings.Compare(v, origin) == 0 {
			return self
		}
	}

	self.messages = append(self.messages, message)
	return self
}

// 必须在包含的数字范围之内
func (self *Validation) InNumbers(field string, message string, intArray *[]int64) *Validation {
	value := self.post.GetInt(field)
	for _, v := range *intArray {
		if v == value {
			return self
		}
	}

	self.messages = append(self.messages, message)
	return self
}

// 输入必须是数字
func (self *Validation) IsNumeric(field string, message string) *Validation {
	_, err := strconv.Atoi(self.post.Get(field))
	if err != nil {
		self.messages = append(self.messages, message)
		return self
	}
	return self
}

// 允许为带小数点模式
func (self *Validation) IsDecimal(field string, message string) *Validation {
	value := self.post.Get(field)
	reg := "^[0-9]+\\.?\\d*$"
	matched, err := regexp.MatchString(reg, value)
	if err != nil || !matched { //如果匹配错误或者是没有匹配到，则添加错误信息
		self.messages = append(self.messages, message)
	}
	return self
}

// 输入必须不能为空
func (self *Validation) NotNull(field string, message string) *Validation {
	value := self.post.Get(field)
	if strings.Compare(string(value), "") == 0 {
		self.messages = append(self.messages, message)
	}
	return self
}

// 等于字段值
func (self *Validation) Equals(field string, equalsField string, message string) *Validation {
	value := self.post.Get(field)
	equalsValue := self.post.Get(equalsField)
	if value != equalsValue {
		self.messages = append(self.messages, message)
	}
	return self
}

// 是否是合法的用户名称
func (self *Validation) IsUserName(field string, message string) *Validation {
	value := self.post.Get(field)
	matched, err := regexp.MatchString("^[a-zA-Z]{1}[a-zA-Z0-9_]{4,19}$", value)
	if err != nil || !matched { //如果匹配错误或者是没有匹配到，则添加错误信息
		self.messages = append(self.messages, message)
	}
	return self
}

// 是否是合法的电子邮件地址
func (self *Validation) IsMail(field string, message string) *Validation {
	value := self.post.Get(field)
	reg := "^[a-zA-Z0-9\\._]+@[\\w_\\-0-9]+(\\.[\\w_\\-0-9]+)+$"
	matched, err := regexp.MatchString(reg, value)
	if err != nil || !matched { //如果匹配错误或者是没有匹配到，则添加错误信息
		self.messages = append(self.messages, message)
	}
	return self
}

// 是否是日期格式
func (self *Validation) IsDate(field string, message string) *Validation {
	value := self.post.Get(field)
	matched, err := regexp.MatchString("^\\d{4}\\-\\d{2}\\-\\d{2}$", value)
	if err != nil || !matched {
		self.messages = append(self.messages, message)
	}
	return self
}

// 是否是日期时间格式
func (self *Validation) IsDateTime(field string, message string) *Validation {
	value := self.post.Get(field)
	matched, err := regexp.MatchString("^\\d{4}\\-\\d{2}\\-\\d{2} \\d{1,2}:\\d{1,2}:\\d{1,2}$", value)
	if err != nil || !matched {
		self.messages = append(self.messages, message)
	}
	return self
}

// 是否是url地址
func (self *Validation) IsUrl(field string, message string) *Validation {
	value := self.post.Get(field)
	matched, err := regexp.MatchString("^https?://[a-zA-Z0-9\\-_]+(\\.[a-zA-Z0-9\\-_]+)+$", value)
	if err != nil || !matched {
		self.messages = append(self.messages, message)
	}
	return self
}

// 是否是合法的版本号
func (self *Validation) IsVersion(field string, message string) *Validation {
	value := self.post.Get(field)
	matched, err := regexp.MatchString("^\\d{1,2}\\.\\d{1,2}\\.\\d{1,3}$", value)
	if err != nil || !matched {
		self.messages = append(self.messages, message)
	}
	return self
}

// 是否是身份证号
func (self *Validation) IsID(field string, message string) *Validation {
	value := self.post.Get(field)
	matched, err := regexp.MatchString("^\\d{17}[0-9A-Z]{1}$", value)
	if err != nil || !matched {
		self.messages = append(self.messages, message)
	}
	return self
}

// 是否是手机号码
func (self *Validation) IsPhoneNumber(field string, message string) *Validation {
	value := self.post.Get(field)
	matched, err := regexp.MatchString("^1[3-9]{1}[0-9]{1}\\d{8}$", value)
	if err != nil || !matched {
		self.messages = append(self.messages, message)
	}
	return self
}

// 是否是银行卡号
func (self *Validation) IsBankCardNumber(field string, message string) *Validation {
	value := self.post.Get(field)
	matched, err := regexp.MatchString("^\\d{12, 20}$", value)
	if err != nil || !matched {
		self.messages = append(self.messages, message)
	}
	return self
}

// 是否是合法的ip地址, 要考虑同时输入多个ip的情况, 兼容ipv4/ipv6
func (self *Validation) IsIpAddress(field string, message string) *Validation {
	value := self.post.Get(field)
	reg := "^(\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}(\\.\\d{1,3}\\.\\d{1,3})?)+(\\,\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}(\\.\\d{1,3}\\.\\d{1,3})?)?$"
	matched, err := regexp.MatchString(reg, value)
	if err != nil || !matched {
		self.messages = append(self.messages, message)
		return self
	}

	isIpAddress := func(ip string) bool {
		numbers := strings.Split(ip, ".")
		if len(numbers) < 4 {
			self.messages = append(self.messages, message)
			return false
		}
		matched, err := regexp.MatchString("^0\\d+$", numbers[0])
		if err != nil || matched { //不能以0开头
			return false
		}
		for _, number := range numbers {
			ipNumber, ipErr := strconv.Atoi(number)
			if ipErr != nil || ipNumber > 255 {
				return false
			}
		}
		return true
	}
	ips := []string{value}
	if strings.Index(value, ",") > 0 { //如果包括,号, 则将ip地址修正为多个ip的数组
		ips = strings.Split(value, ",")
	}

	for _, ip := range ips { //对所有ip进行判断
		if !isIpAddress(ip) {
			self.messages = append(self.messages, message)
			return self
		}
	}

	return self
}

// 返回错误信息及校验结果
func (self *Validation) Validate() (string, bool) {
	if len(self.messages) == 0 {
		return "", true
	}

	return strings.Join(self.messages, ","), false
}
