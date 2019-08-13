package xorm

type SilverMerchantOsLogs struct {
	Id         int    `xorm:"not null pk autoincr INT(11)"`
	MerchantId int    `xorm:"not null comment('silver_merchant_users表id') INT(11)"`
	Content    string `xorm:"not null default '' comment('操作内容') VARCHAR(255)"`
	Created    int    `xorm:"default 0 comment('操作时间') INT(11)"`
	Ip         string `xorm:"comment('操作ip') VARCHAR(50)"`
	City       string `xorm:"comment('所在城市') VARCHAR(50)"`
}
