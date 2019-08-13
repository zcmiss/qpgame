package xorm

import (
	"time"
)

type ChargeRecords struct {
	Id                int       `xorm:"not null pk autoincr comment('主键自增') INT(11)"`
	UserId            int       `xorm:"not null comment('用户ID') INT(11)"`
	Amount            string    `xorm:"not null default 0.00 comment('充值金额') DECIMAL(12,3)"`
	OrderId           string    `xorm:"not null default '' comment('充值订单') CHAR(21)"`
	ChargeTypeId      int       `xorm:"not null default 0 comment('充值方式id') INT(10)"`
	CardNumber        string    `xorm:"not null default '' comment('卡号') VARCHAR(50)"`
	BankAddress       string    `xorm:"not null default '' comment('开户银行地址或支付二维码') VARCHAR(100)"`
	Created           int       `xorm:"not null default 0 comment('添加时间') INT(10)"`
	State             int       `xorm:"not null default 0 comment('公司入款：0 待审核，1 成功，2 失败。线上支付：0待处理，1成功，2失败，3进行中,4退款，5取消，6强制入款') TINYINT(1)"`
	Screenshot        string    `xorm:"not null default '' comment('屏幕截图') VARCHAR(100)"`
	ReceiptScreenshot string    `xorm:"not null default '' comment('收据截图') VARCHAR(100)"`
	Updated           int       `xorm:"not null default 1495702873 comment('修改时间') INT(10)"`
	ChargeTypeInfo    string    `xorm:"not null default '' comment('充值类型说明') VARCHAR(30)"`
	Ip                string    `xorm:"not null default '' comment('充值IP') VARCHAR(16)"`
	PlatformId        int       `xorm:"not null default 0 comment('充值platform') TINYINT(2)"`
	RealName          string    `xorm:"not null default '' comment('真实姓名') VARCHAR(30)"`
	BankTypeId        int       `xorm:"not null default 0 comment('银行转账类型') TINYINT(2)"`
	BankChargeTime    int       `xorm:"not null default 0 comment('银行转账时间') INT(10)"`
	CredentialId      int       `xorm:"not null default 0 comment('第三方支付记录ID') INT(10)"`
	Operator          string    `xorm:"not null default '' comment('操作者') VARCHAR(30)"`
	IsTppay           int       `xorm:"default 0 comment('是否第三方支付 0为否；1为是') TINYINT(1)"`
	ChargeCardId      int       `xorm:"not null default 0 comment('关联ChargeBankCards.id') INT(11)"`
	Remark            string    `xorm:"not null default '' comment('备注') VARCHAR(100)"`
	UpdatedLast       time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('最后更新时间') TIMESTAMP"`
	PayBankCode       string    `xorm:"not null default '' comment('充值IP') VARCHAR(30)"`
}
