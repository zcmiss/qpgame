package xorm

type SilverMerchantChargeCards struct {
	Id         	int    `xorm:"not null pk autoincr INT(11)"`
	Name       	string `xorm:"not null default '' comment('银行名称') VARCHAR(30)"`
	Owner     	string `xorm:"not null default '' comment('持卡人') VARCHAR(30)"`
	CardNumber 	string `xorm:"not null default '' comment('卡号') VARCHAR(30)"`
	BankAddress string `xorm:"not null default '' comment('开户地址') VARCHAR(150)"`
	Remark     	string `xorm:"not null default '' comment('备注') VARCHAR(255)"`
	Logo     	string `xorm:"not null default '' comment('logo') VARCHAR(300)"`
	Mfrom 		int    `xorm:"not null default 0 comment('支付最小金额') INT(10)"`
	Mto 		int    `xorm:"not null default 0 comment('支付最大金额') INT(10)"`
	Priority 	int    `xorm:"not null default 0 comment('充值排序') INT(3)"`
	State 		int    `xorm:"not null default 0 comment('状态（0停用，1可用）') INT(1)"`
	Created 	int    `xorm:"not null default 0 comment('添加时间') INT(10)"`
}

