package xorm

import (
	"github.com/go-xorm/xorm"
	"qpgame/models"
)

type SilverMerchantUsers struct {
	Id                 int    `xorm:"not null pk autoincr INT(11)"`
	UserId             int    `xorm:"not null default 0 comment('关联用户表id') INT(11)"`
	MerchantLevel      int    `xorm:"not null default 1 comment('银商等级,预留字段') INT(10)"`
	AuthAmount         string `xorm:"not null default 0.000 comment('当前授权额度') DECIMAL(12,3)"`
	UsableAmount       string `xorm:"not null default 0.000 comment('可用额度') DECIMAL(12,3)"`
	MerchantCashPledge string `xorm:"not null default 0.000 comment('银商押金') DECIMAL(12,3)"`
	TotalChargeMoney   string `xorm:"not null default 0.000 comment('累计充值金额') DECIMAL(12,3)"`
	TotalAuthAmount    string `xorm:"not null default 0.000 comment('累计授权金额') DECIMAL(12,3)"`
	DonateRate         string `xorm:"not null default 0.000 comment('赠送比例,这是银商的收入来源很重要,比如冲1万，送4%') DECIMAL(12,3)"`
	Account            string `xorm:"not null default '' comment('银商账号') VARCHAR(20)"`
	Password           string `xorm:"not null comment('银商密码') CHAR(40)" json:"-"`
	Created            int    `xorm:"not null comment('创建时间') INT(10)"`
	Status             int    `xorm:"not null default 1 comment('银商状态,1正常，0锁定') TINYINT(1)"`
	IsDestroy          int    `xorm:"not null default 0 comment('银商是否注销状态,1已注销，0未注销') TINYINT(1)"`
	Token              string `xorm:"not null default '' comment('用户登录token,要保持到程序内存中') VARCHAR(200)"`
	TokenCreated       int    `xorm:"not null comment('token创建时间,根据这个来双层判断是否已过期') INT(10)"`
	LastLoginTime      int    `xorm:"not null default 0 comment('上次登录时间') INT(10)"`
	MerchantName       string `xorm:"not null default '' comment('商户名称') VARCHAR(20)"`
	WebCustomerUrl     string `xorm:"-" comment('在线客服')` // 非表字段
	CashPledge         string `xorm:"-" comment('押金额度')` // 非表字段
}

func GetSliverMerchantUser(platform string, id int) (SilverMerchantUsers, bool) {
	var smUser = SilverMerchantUsers{}
	exist, _ := models.MyEngine[platform].ID(id).Get(&smUser)
	return smUser, exist
}

func GetSliverMerchantUserByAccount(platform string, account string) (SilverMerchantUsers, bool) {
	var smUser = SilverMerchantUsers{}
	exist, _ := models.MyEngine[platform].Where("account=?", account).Get(&smUser)
	return smUser, exist
}

func UpdateSliverMerchantUser(session *xorm.Session, uid int, smUser SilverMerchantUsers) (bool, error) {
	affNum, err := session.ID(uid).Update(smUser)
	isUpdate := affNum > 0
	return isUpdate, err
}
