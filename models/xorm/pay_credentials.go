package xorm

type PayCredentials struct {
	Id               int    `xorm:"not null pk autoincr INT(10)"`
	PlatForm         int    `xorm:"not null default 0 comment('支付标识,所有平台有共同标识') TINYINT(3)"`
	PayName          string `xorm:"not null default '' comment('平台名称') VARCHAR(20)"`
	MerchantNumber   string `xorm:"not null default '' comment('商户号') VARCHAR(32)"`
	PrivateKey       string `xorm:"not null default '' comment('商户私钥') VARCHAR(1024)"`
	Corporate        string `xorm:"not null default '' comment('法人') VARCHAR(20)"`
	IdUmber          string `xorm:"not null default '' comment('法人身份证号') VARCHAR(18)"`
	CardNumber       string `xorm:"not null default '' comment('银行卡号') VARCHAR(36)"`
	PhoneNumber      string `xorm:"not null default '0' comment('手机号码') VARCHAR(11)"`
	PublicKey        string `xorm:"not null default '' comment('证书公钥') VARCHAR(1024)"`
	PrivateKeyFile   string `xorm:"not null default '' comment('私钥文件') VARCHAR(150)"`
	CredentialKey    string `xorm:"not null default '' VARCHAR(2048)"`
	CallbackKey      string `xorm:"not null default '' VARCHAR(150)"`
	ChargeAmountConf int    `xorm:"not null default 0 comment('充值金额配置:0 关闭随机金额小数,1 开启随机金额小数到角,2 开启随机金额小数到分') TINYINT(1)"`
	Status           int    `xorm:"not null default 1 comment('是否弃用,0弃用,1启用') TINYINT(1)"`
}
