package mainxorm

type Platform struct {
	Id             				int    `xorm:"not null pk autoincr comment('主键自增') INT(11)"`
	Code           				string `xorm:"not null default '' comment('平台号') VARCHAR(20)"`
	Name           				string `xorm:"not null default '' comment('平台名称') VARCHAR(50)"`
	AdminAddress   				string `xorm:"not null default '' comment('后台管理接口地址') VARCHAR(255)"`
	ApiAddress     				string `xorm:"not null default '' comment('接口地址') VARCHAR(255)"`
	PicAddress    		 		string `xorm:"not null default '' comment('图片地址') VARCHAR(255)"`
	PayBackAddress 				string `xorm:"not null default '' comment('支付回调地址') VARCHAR(255)"`
	SilverMerchantAddress 		string `xorm:"not null default '' comment('银商系统h5地址') VARCHAR(100)"`
	SilverMerchantApiAddress 	string `xorm:"not null default '' comment('银商系统api地址') VARCHAR(100)"`
}
