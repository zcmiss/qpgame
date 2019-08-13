package xorm

type PlatformAccounts struct {
	Id       int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	PlatId   int    `json:"plat_id" xorm:"not null comment('平台编号') INT(11)"`
	UserId   int    `json:"user_id" xorm:"not null comment('用户编号') INT(11)"`
	Username string `json:"username" xorm:"not null comment('用户名') VARCHAR(40)"`
	Password string `json:"password" xorm:"not null comment('密码') VARCHAR(40)"`
	Created  int    `json:"created" xorm:"not null comment('创建时间') INT(10)"`
}
