package xorm

type Admins struct {
	Id            int    `xorm:"not null pk autoincr comment('主键自增') INT(11)"`
	Name          string `xorm:"not null comment('管理员名称') VARCHAR(255)"`
	Email         string `xorm:"not null comment('电子邮件') VARCHAR(255)"`
	Password      string `xorm:"not null comment('密码') VARCHAR(255)"`
	RoleId        int    `xorm:"not null comment('角色') SMALLINT(3)"`
	Created       int    `xorm:"not null comment('创建时间') INT(10)"`
	Updated       int    `xorm:"not null default 0 comment('更新时间') INT(10)"`
	Status        int    `xorm:"not null comment('状态') INT(1)"`
	ChargeAlert   int    `xorm:"not null default 1 comment('后台充值提醒，1开启，0关闭') TINYINT(1)"`
	WithdrawAlert int    `xorm:"not null default 0 comment('后台出款提醒，1开启，0关闭') TINYINT(1)"`
	LoginIp       string `xorm:"not null default '' comment('允许登录IP') VARCHAR(1000)"`
	Permission    int    `xorm:"not null default 1 comment('涉及钱的权限（0、无权限，1、主管权限）') TINYINT(1)"`
	ForceOut      int    `xorm:"not null default 0 comment('是否强制退出,0 无强制退出,1 强制退出') TINYINT(1)"`
	ManualMax     int    `xorm:"not null default 0 comment('最大人工入款金额') INT(10)"`
	IsOtp         int    `xorm:"not null default 0 comment('是否OTP验证登录（0为否，1为是）') TINYINT(1)"`
	IsOtpFirst    int    `xorm:"not null default 1 comment('是否第一次OTP验证登录') TINYINT(1)"`
}
