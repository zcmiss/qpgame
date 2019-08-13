package xorm

type PhoneCodes struct {
	Id      int    `xorm:"not null pk autoincr comment('主键自增') INT(11)"`
	Phone   string `xorm:"not null comment('手机号') VARCHAR(15)"`
	Code    string `xorm:"not null comment('验证码') VARCHAR(6)"`
	Created int    `xorm:"not null comment('创建时间') INT(15)"`
}
