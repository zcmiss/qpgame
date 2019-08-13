package xorm

type ChargeCards struct {
	Id           int    `xorm:"not null pk autoincr comment('主键自增') INT(11)"`
	Name         string `xorm:"not null comment('银行名称') VARCHAR(30)"`
	Owner        string `xorm:"not null comment('持卡人') VARCHAR(30)"`
	CardNumber   string `xorm:"comment('卡号') VARCHAR(30)"`
	BankAddress  string `xorm:"not null comment('开户地址') VARCHAR(150)"`
	ChargeTypeId int    `xorm:"not null default 0 comment('充值类型编号') MEDIUMINT(4)"`
	Created      int    `xorm:"not null default 0 comment('添加时间') INT(10)"`
	Remark       string `xorm:"not null default '' comment('备注') VARCHAR(255)"`
	State        int    `xorm:"not null default 1 comment('状态') TINYINT(1)"`
	Logo         string `xorm:"not null default '' comment('LOGO') VARCHAR(100)"`
	Hint         string `xorm:"not null default '' comment('支付提示') VARCHAR(150)"`
	Title        string `xorm:"not null default '' comment('支付标题') VARCHAR(30)"`
	Mfrom        int    `xorm:"not null default 0 comment('支付额度') INT(10)"`
	Mto          int    `xorm:"not null default 0 comment('支付额度') INT(10)"`
	AmountLimit  int    `xorm:"not null default 0 comment('停用金额') INT(10)"`
	AddrType     int    `xorm:"not null default 1 TINYINT(1)"`
	CredentialId int    `xorm:"not null comment('以付方式') INT(10)"`
	UserGroupIds string `xorm:"not null default '1' comment('用户组id') VARCHAR(30)"`
	Priority     int    `xorm:"not null default 1 comment('充值排序') TINYINT(3)"`
	QrCode       string `xorm:"not null comment('财务收款二维码') VARCHAR(150)"`
}
