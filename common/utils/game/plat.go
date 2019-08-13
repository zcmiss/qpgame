package game

import "qpgame/models/xorm"

const userNameConst = "ckyx"
const userPwdConst = "ck20190322"

type Plat interface {
	getGameList(platform string) error
	createPlayer(userid string, platform string, platId int) (xorm.PlatformAccounts, bool)
	getGameUrl(accounts *xorm.PlatformAccounts, gamecode string, ip string) string
	uchips(username string, exId string, amount string) bool
	queryUchips(username string) (string, bool)
}
