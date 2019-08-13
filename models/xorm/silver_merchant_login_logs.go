package xorm

import "github.com/go-xorm/xorm"

type SilverMerchantLoginLogs struct {
	Id         int    `xorm:"not null pk autoincr INT(11)"`
	MerchantId int    `xorm:"not null comment('silver_merchant_users表id') INT(11)"`
	LoginTime  int    `xorm:"not null default 0 comment('登录时间') INT(11)"`
	Ip         string `xorm:"not null comment('登录IP') CHAR(24)"`
	LoginCity  string `xorm:"comment('登录城市') VARCHAR(50)"`
}

func RecordSilverMerchantLoginLog(session *xorm.Session, log SilverMerchantLoginLogs) (SilverMerchantLoginLogs, bool) {
	affNum, err := session.Insert(&log)
	var isAdd bool
	if err != nil {
		isAdd = false
	} else {
		isAdd = affNum > 0
	}
	return log, isAdd
}
