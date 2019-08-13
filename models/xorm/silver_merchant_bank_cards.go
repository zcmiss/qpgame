package xorm

type SilverMerchantBankCards struct {
	Id         int    `xorm:"not null pk autoincr INT(11)"`
	MerchantId int    `xorm:"not null default 0 comment('silver_merchant_users表id') INT(10)"`
	CardNumber string `xorm:"not null default '' comment('银行卡号') VARCHAR(50)"`
	Address    string `xorm:"not null default '' comment('银行卡地址') VARCHAR(100)"`
	Created    int    `xorm:"not null default 0 comment('添加时间') INT(10)"`
	BankName   string `xorm:"not null default '' comment('银行名称') VARCHAR(50)"`
	Name       string `xorm:"not null default '' comment('姓名') VARCHAR(50)"`
	Status     int    `xorm:"not null default 0 comment('状态') TINYINT(1)"`
	Updated    int    `xorm:"not null default 0 comment('修改时间') INT(10)"`
	Remark     string `xorm:"not null default '' comment('备注') VARCHAR(100)"`
}
