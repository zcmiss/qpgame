package game

import (
	"qpgame/models/xorm"
)

//根据不同平台调用不同的刷新游戏列表接口
func GetGameList(platform string, platId int) error {
	return getPlat(platId, platform).getGameList(platform)
}

//根据不同平台调用不同创建会员接口
func CreatePlay(userid string, platform string, platId int) (xorm.PlatformAccounts, bool) {
	return getPlat(platId, platform).createPlayer(userid, platform, platId)
}

//根据不同平台调用不同创建会员接口
func GetGameUrl(accounts *xorm.PlatformAccounts, gamecode string, ip string, platId int, platform string) string {
	return getPlat(platId, platform).getGameUrl(accounts, gamecode, ip)
}

//根据不同平台调用不同存取款接口
func Uchips(username string, exId string, amount string, platId int, platform string) bool {
	return getPlat(platId, platform).uchips(username, exId, amount)
}

func QueryUchips(username string, platId int, platform string) (string, bool) {
	return getPlat(platId, platform).queryUchips(username)
}

func ExitGame(username string, exId string, amount string, platId int, platform string) bool {
	return getPlat(platId, platform).uchips(username, exId, amount)
}

//根据平台编号返回对应的平台
func getPlat(platId int, platform string) Plat {
	switch platId {
	case 1:
		return GetFG(platform)
	case 2:
		return GetAe(platform)
	case 3:
		return GetMg(platform)
	case 4:
		return GetKy(platform)
	case 5:
		return GetAg(platform)
	case 6:
		return GetLy(platform)
	case 7:
		return GetNW(platform)
	case 8:
		return GetVG(platform)
	case 9:
		return GetJdb(platform)
	case 10:
		return GetOg(platform)
	case 11:
		return GetUg(platform)
	}
	return GetFG(platform)
}
