package xorm

type WashCodeInfos struct {
	Id        int    `xorm:"not null pk autoincr INT(11)"`
	RecordId  int    `xorm:"not null comment('洗码记录编号') INT(11)"`
	TypeName  string `xorm:"not null comment('游戏类型名称') VARCHAR(20)"`
	TypeId    int    `xorm:"not null comment('游戏类型编号') INT(11)"`
	GameName  string `xorm:"not null comment('游戏名称') VARCHAR(20)"`
	BetAmount string `xorm:"not null comment('游戏洗码量') DECIMAL(12,3)"`
	Amount    string `xorm:"not null comment('洗码金额') DECIMAL(12,3)"`
	Rate      string `xorm:"not null comment('洗码比例') VARCHAR(10)"`
}
