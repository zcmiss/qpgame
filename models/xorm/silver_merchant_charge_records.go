package xorm

import (
	"time"
)

type SilverMerchantChargeRecords struct {
	Id             		int       `xorm:"not null pk autoincr INT(11)"`
	MerchantId     		int       `xorm:"not null comment('silver_merchant_users表Id') INT(11)"`
	Amount         		string    `xorm:"not null default 0.000 comment('充值授权额度金额') DECIMAL(12,3)"`
	PresentedMoney      string    `xorm:"not null default 0.000 comment('充值赠送金额') DECIMAL(12,3)"`
	OrderId        		string    `xorm:"not null default '' comment('充值订单') CHAR(21)"`
	BankName       		string    `xorm:"not null default '' comment('银行名字') VARCHAR(20)"`
	BankAddress    		string    `xorm:"not null default '' comment('充值开户银行地址') VARCHAR(100)"`
	CardNumber     		string    `xorm:"not null default '' comment('充值卡号') VARCHAR(50)"`
	Created        		int       `xorm:"not null default 0 comment('添加时间') INT(10)"`
	State          		int       `xorm:"not null default 0 comment('0 待审核，1 成功，2 失败。') TINYINT(1)"`
	Updated        		int       `xorm:"not null default 1495702873 comment('修改时间') INT(10)"`
	Ip             		string    `xorm:"not null default '' comment('充值IP') VARCHAR(16)"`
	RealName       		string    `xorm:"not null default '' comment('真实姓名') VARCHAR(30)"`
	BankChargeTime 		int       `xorm:"not null default 0 comment('银行转账时间') INT(10)"`
	Operator       		string    `xorm:"not null default '' comment('操作者') VARCHAR(30)"`
	Remark         		string    `xorm:"not null default '' comment('备注') VARCHAR(100)"`
	UpdatedLast    		time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('最后更新时间') TIMESTAMP"`
}
