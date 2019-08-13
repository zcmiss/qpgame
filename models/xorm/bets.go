package xorm

import (
	"qpgame/models"
	"strconv"
)

type Bets struct {
	Amount          string `xorm:"not null default 0.00 comment('下注金额') DECIMAL(10,2)"`
	AmountAll       string `xorm:"not null default 0.00 comment('总下注金额') DECIMAL(10,2)"`
	AmountPlatform  string `xorm:"not null default 0.00 comment('第三方平台抽水') DECIMAL(10,2)"`
	Created         int    `xorm:"not null default 0 comment('游戏开始时间') index INT(11)"`
	Ented           int    `xorm:"default 0 comment('游戏结束时间') INT(10)"`
	Started         int    `xorm:"default 0 comment('游戏开始时间') INT(10)"`
	GameCode        string `xorm:"not null comment('游戏编号') VARCHAR(11)"`
	Gt              string `xorm:"comment('游戏类别') VARCHAR(20)"`
	Id              int    `xorm:"not null pk autoincr INT(11)"`
	OrderId         string `xorm:"not null default '' comment('订单编号') unique VARCHAR(80)"`
	PlatformId      int    `xorm:"not null comment('第三方平台ID') index INT(11)"`
	Accountname     string `xorm:"comment('第三方平台账号') VARCHAR(50)"`
	PlatformName    string `xorm:"comment('第三方平台名称') VARCHAR(50)"`
	Reward          string `xorm:"not null default 0.00 comment('中奖金额') DECIMAL(10,2)"`
	UserId          int    `xorm:"not null comment('用户ID') index INT(11)"`
	Name            string `xorm:"<-"`
	GameCategorieId string `xorm:"<-"`
	ParentId        string `xorm:"<-"`
	RebateState     int    `xorm:"not null default 0 comment('洗码状态0未洗1已洗') index INT(2)"`
}

func (bets Bets) TableName() string {
	return "bets_" + strconv.Itoa(bets.UserId%10)
}

func GetBets(platform string, userId int, gameCategorieId string, parentId string) []Bets {
	bets := make([]Bets, 0)
	sql := "SELECT b.*, p.`name`, p.`game_categorie_id`, g.`parent_id` FROM bets_" + strconv.Itoa(userId%10) + " b LEFT JOIN platform_games p ON b.`game_code` = p.service_code LEFT JOIN game_categories g ON p.`game_categorie_id` = g.`id` where g.`parent_id` = ? and b.user_id = ?"
	session := models.MyEngine[platform].SQL(sql, userId, parentId)
	if gameCategorieId != "" {
		session = session.And("p.game_categorie_id = ?", gameCategorieId)
	}
	session.Find(&bets)
	return bets
}
