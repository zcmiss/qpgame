/*
	这里存放非api缓存
	建议: 此文件只存放健名称, 另一个文件internal.go存放键/值的操作
*/
package ramcache

import (
	"encoding/json"
	"fmt"
	"os"
	"qpgame/models"
	"qpgame/models/beans"
	"qpgame/models/xorm"
	"strconv"
	"sync"

	libXorm "github.com/go-xorm/xorm"
)

// 保存的所有的表-字段映射信息
var TableFields = make(map[string]map[string]int)

func loadTableProxyChessLevels(platform string, engine *libXorm.Engine) bool {
	proxyChessLevels := make([]xorm.ProxyChessLevels, 0)
	errRes := engine.Desc("level").Find(&proxyChessLevels)
	if errRes != nil {
		fmt.Println(errRes)
		return false
	}
	TableProxyChessLevels.Store(platform, proxyChessLevels)
	return true
}

func loadTableProxyRealLevels(platform string, engine *libXorm.Engine) bool {
	proxyRealLevels := make([]xorm.ProxyRealLevels, 0)
	errRes := engine.Desc("level").Find(&proxyRealLevels)
	if errRes != nil {
		fmt.Println(errRes)
		return false
	}
	TableProxyRealLevels.Store(platform, proxyRealLevels)
	return true
}

//加载主数据库的platform表
var MainTablePlatform sync.Map

//加载活动列表
func loadTableActivities(platform string, engine *libXorm.Engine) bool {
	activities := make([]xorm.Activities, 0)
	errRes := engine.Find(&activities)
	if errRes != nil {
		fmt.Println(errRes)
		return false
	}
	TableActivities.Store(platform, activities)
	return true
}

//加载活动分类表
func loadTableActivityClasses(platform string, engine *libXorm.Engine) bool {
	activityClasses := make([]xorm.ActivityClasses, 0)
	errRes := engine.Find(&activityClasses)
	if errRes != nil {
		fmt.Println(errRes)
		return false
	}
	TableActivityClasses.Store(platform, activityClasses)
	return true
}

//加载用户相关字段到缓存
//一个用户1k,10万用户为50M空间
func loadTableUsers(platform string, engine *libXorm.Engine) bool {
	//启动项目加载所有用户手机号到缓存里,用于注册的时候判断该号码是否已经注册
	res, errRes := engine.Query("SELECT unique_code,phone,user_name,id,token,token_created,user_group_id,wx_open_id FROM users")
	if errRes != nil {
		fmt.Println(errRes)
		return false
	}
	upMap := make(map[int]beans.UserProfile)
	woiMap := make(map[string]beans.WxOpenId)
	pmap := make(map[string][]string)
	umap := make(map[string][]string)
	uniqueMap := make(map[string][]string)
	for _, p := range res {
		uid, _ := strconv.Atoi(string(p["id"]))
		iTokenCreated, _ := strconv.Atoi(string(p["token_created"]))
		upMap[uid] = beans.UserProfile{
			Phone:        string(p["phone"]),
			Username:     string(p["user_name"]),
			Token:        string(p["token"]),
			TokenCreated: iTokenCreated,
			UserGroupId:  string(p["user_group_id"]),
			UniqueCode:   string(p["unique_code"]),
			WxOpenId:     string(p["wx_open_id"]),
		}
		if string(p["unique_code"]) != "" {
			uniqueMap[string(p["unique_code"])] = []string{
				string(p["id"]),
				string(p["token"]),
				string(p["token_created"]),
				string(p["user_group_id"]),
			}
		}
		if string(p["phone"]) != "" {
			pmap[string(p["phone"])] = []string{
				string(p["id"]),
				string(p["token"]),
				string(p["token_created"]),
				string(p["user_group_id"]),
			}
		}
		umap[string(p["user_name"])] = []string{
			string(p["id"]),
			string(p["token"]),
			string(p["token_created"]),
			string(p["user_group_id"]),
		}
		wxOpenId := string(p["wx_open_id"])
		if "" != wxOpenId {
			woiMap[wxOpenId] = beans.WxOpenId{
				UserId: uid,
			}
		}
	}
	UserIdCard.Store(platform, upMap)
	PhoneNumAndToken.Store(platform, pmap)
	UserNameAndToken.Store(platform, umap)
	UniqueCodeAndToken.Store(platform, uniqueMap)
	WxOpenIdIndex.Store(platform, woiMap)
	return true
}

//加载游戏平台账号表,10万用户50M空间
func loadTablePlatformAccounts(platform string, engine *libXorm.Engine) bool {
	res, errRes := engine.Query("SELECT username,user_id FROM platform_accounts GROUP BY username,user_id")
	if errRes != nil {
		fmt.Println(errRes)
		return false
	}
	amap := make(map[string]int)
	for _, p := range res {
		amap[string(p["username"])], _ = strconv.Atoi(string(p["user_id"]))
	}
	TablePlatformAccounts.Store(platform, amap)
	return true
}

//加载游戏平台账号表,10万用户50M空间
func UpdateTablePlatformAccounts(username string, platform string, engine *libXorm.Engine) bool {
	res, errRes := engine.Query("SELECT username,user_id FROM platform_accounts where username = '" + username + "' GROUP BY username,user_id")
	if errRes != nil {
		fmt.Println(errRes)
		return false
	}
	amap, _ := TablePlatformAccounts.Load(platform)
	for _, p := range res {
		amap.(map[string]int)[string(p["username"])], _ = strconv.Atoi(string(p["user_id"]))
	}
	TablePlatformAccounts.Store(platform, amap)
	return true
}

//获取投注的第三方key,10K 空间
func loadTableBetsKey(platform string, engine *libXorm.Engine) bool {
	res, errRes := engine.Query("SELECT id,plat_id,gt,search_key from bets_key")
	if errRes != nil {
		fmt.Println(errRes)
		return false
	}
	bmap := make(map[string]string)
	for _, p := range res {
		bmap[string(p["plat_id"])+"-"+string(p["gt"])] = string(p["search_key"])
	}
	TableBetsKey.Store(platform, bmap)
	return true
}

//游戏分类表到内存中
func loadTableGameCategories(platform string, engine *libXorm.Engine) bool {
	gameCategories := make([]xorm.GameCategories, 0)
	//可以修改为select g.*,p.status platform_status from game_categories g left join platforms p on g.platform_id = p.id where g.status = 1
	sqlString := "select * from(select g.*,ifnull(p.`status`,1) platform_status from game_categories g "
	sqlString += "left join platforms p on g.platform_id = p.id) a where a.status=1"
	errRes := engine.SQL(sqlString).Find(&gameCategories)
	if errRes != nil {
		fmt.Println(errRes)
		return false
	}
	TableGameCategories.Store(platform, gameCategories)
	return true
}

//平台列表加载到内存中 10M内存,搜索条件修改之后要修改kafka那边逻辑
func loadTablePlatforms(platform string, engine *libXorm.Engine) bool {
	platforms := make([]xorm.Platforms, 0)
	errRes := engine.SQL("SELECT * FROM platforms").Find(&platforms)
	if errRes != nil {
		fmt.Println(errRes)
		return false
	}
	TablePlatforms.Store(platform, platforms)
	ReloadGamePlatformApiConfig(platform, platforms)
	return true
}

//平台游戏列表加载到内存中 10M内存,搜索条件修改之后要修改kafka那边逻辑
func loadTablePlatformGames(platform string, engine *libXorm.Engine) bool {
	platformGames := make([]xorm.PlatformGames, 0)
	sqlString := "SELECT g.* ,p.status platform_status FROM platform_games g "
	sqlString += "LEFT JOIN platforms p ON g.`plat_id` = p.id  WHERE g.ishidden=0"
	errRes := engine.SQL(sqlString).Find(&platformGames)
	if errRes != nil {
		fmt.Println(errRes)
		return false
	}
	TablePlatformGames.Store(platform, platformGames)
	return true
}

//JDB平台游戏列表加载到内存中 10M内存,搜索条件修改之后要修改kafka那边逻辑
func loadTablePlatformGamesByJDB(platform string, engine *libXorm.Engine) bool {
	platformGames := make([]xorm.PlatformGames, 0)
	errRes := engine.SQL("SELECT g.* ,p.status platform_status FROM platform_games g LEFT JOIN platforms p ON g.`plat_id` = p.id  WHERE g.ishidden=0 and g.plat_id=9").Find(&platformGames)
	if errRes != nil {
		fmt.Println(errRes)
		return false
	}
	TablePlatformGamesByJDB.Store(platform, platformGames)
	return true
}

//平台客户端版本表 10M内存
func loadTableAppVersions(platform string, engine *libXorm.Engine) bool {
	appVersions := make([]xorm.AppVersions, 0)
	errRes := engine.Find(&appVersions)
	if errRes != nil {
		fmt.Println(errRes)
		return false
	}
	TableAppVersions.Store(platform, appVersions)
	return true
}

//平台客户端配置表 30M内存
func loadTableConfigs(platform string, engine *libXorm.Engine) bool {
	cfgs := make([]xorm.Configs, 0)
	errRes := engine.Cols("name", "value").Find(&cfgs)
	if errRes != nil {
		fmt.Println(errRes.Error())
		return false
	}
	var cfg = make(map[string]interface{})
	for _, v := range cfgs {
		var tempV interface{}
		err := json.Unmarshal([]byte(v.Value), &tempV)
		if err != nil {
			tempV = v.Value
		}
		cfg[v.Name] = tempV
	}
	TableConfigs.Store(platform, cfg)
	return true
}

//平台客户端系统通知表,只取前30条 50M内存
func loadTableSystemNotices(platform string, engine *libXorm.Engine) bool {
	systemNotices := make([]xorm.SystemNotices, 0)
	session := engine.Cols("content", "title", "created")
	//永远只显示前30条记录
	errRes := session.Where("status = ?", 1).Desc("created").Limit(30).Find(&systemNotices)
	if errRes != nil {
		fmt.Println(errRes)
		return false
	}
	TableSystemNotices.Store(platform, systemNotices)

	return true
}

//会有等级说明表1M空间
func loadTableVipLevels(platform string, engine *libXorm.Engine) bool {
	vipLevels := make([]xorm.VipLevels, 0)
	errRes := engine.Asc("level").Find(&vipLevels)
	if errRes != nil {
		fmt.Println(errRes)
		return false
	}
	TableVipLevels.Store(platform, vipLevels)
	return true
}

//加载充值类型表
func loadTableChargeTypes(platform string, engine *libXorm.Engine) bool {
	chargeTypes := make([]xorm.ChargeTypes, 0)
	errRes := engine.Asc("priority").Find(&chargeTypes)
	if errRes != nil {
		fmt.Println(errRes)
		return false
	}
	TableChargeTypes.Store(platform, chargeTypes)

	return true
}

//加载charge_cards表
func loadTableChargeCards(platform string, engine *libXorm.Engine) bool {
	chargeCards := make([]xorm.ChargeCards, 0)
	errRes := engine.Find(&chargeCards)
	if errRes != nil {
		fmt.Println(errRes)
		return false
	}
	TableChargeCards.Store(platform, chargeCards)

	return true
}

//加载pay_credential表
func loadTablePayCredential(platform string, engine *libXorm.Engine) bool {
	payCredentials := make([]xorm.PayCredentials, 0)
	errRes := engine.Find(&payCredentials)
	if errRes != nil {
		fmt.Println(errRes)
		return false
	}
	//下标为id的数据
	var payCredMap = make(map[int]xorm.PayCredentials)
	for _, v := range payCredentials {
		id := v.Id
		payCredMap[id] = v
	}
	TablePayCredential.Store(platform, payCredMap)
	return true
}

//加载bets_key表
func loadBetsSearchType(platform string, engine *libXorm.Engine) bool {
	gameCategories := make([]xorm.GameCategories, 0)
	errRes := engine.SQL("select * from(select g.*,ifnull(p.`status`,1) platform_status from game_categories g left join platforms p on g.platform_id = p.id)a").Find(&gameCategories)
	if errRes != nil {
		fmt.Println(errRes)
		return false
	}
	//下标为id的数据
	var typeMap = make(map[int]map[string]interface{})
	for _, v := range gameCategories {
		id := v.Id
		tmap := make(map[string]interface{})
		if v.ParentId == 0 && v.Id > 1 {
			tmap["name"] = v.Name
			smaps := make([]map[string]string, 0)
			for _, v2 := range gameCategories {
				if v2.ParentId == v.Id {
					smap := make(map[string]string)
					smap["name"] = v2.Name
					smap["value"] = strconv.Itoa(v2.Id)
					smaps = append(smaps, smap)
				}
			}
			tmap["searchList"] = smaps
			typeMap[id] = tmap
		}

	}
	BetsSearchType.Store(platform, typeMap)
	return true
}

//加载银行卡列表
func loadTableUserBanks(platform string, engine *libXorm.Engine) bool {
	userBanks := make([]xorm.UserBanks, 0)
	errRes := engine.Find(&userBanks)
	if errRes != nil {
		fmt.Println(errRes)
		return false
	}
	TableUserBanks.Store(platform, userBanks)
	return true
}

//加载需要处理的因异常导致的任务
func loadExceptionTasks(platform string, engine *libXorm.Engine) bool {
	exceptionTasks := make(map[string]xorm.ExceptionTasks, 0)
	allExceptionTasks := make([]xorm.ExceptionTasks, 0)
	errRes := engine.Where("flag>-1").Find(&allExceptionTasks)
	if errRes != nil {
		fmt.Println(errRes)
		return false
	}
	for _, et := range allExceptionTasks {
		idx := strconv.Itoa(et.UserId) + "_" + strconv.Itoa(et.PlatId) + "_" + strconv.Itoa(et.Flag)
		exceptionTasks[idx] = et
	}
	TableExceptionTasks.Store(platform, exceptionTasks)
	return true
}

//处理每个游戏分类的层级关系
func handlerGameCategory(platform string) {
	//加载每个平台对应的游戏
	//大分类游戏
	gameCategories := make([]xorm.GameCategories, 0)
	//小分类游戏
	gameCategories2 := make([]xorm.GameCategories, 0)
	gcates, _ := TableGameCategories.Load(platform)
	for _, v := range gcates.([]xorm.GameCategories) {
		if v.ParentId == 0 && v.Status == 1 {
			gameCategories = append(gameCategories, v)
		}
		if v.ParentId != 0 && v.Status == 1 {
			gameCategories2 = append(gameCategories2, v)
		}
	}

	cmaps := make([]map[string]interface{}, 0)
	platGms, _ := TablePlatformGames.Load(platform)
	for _, c := range gameCategories {
		cmap := make(map[string]interface{})
		cmap["id"] = c.Id
		cmap["name"] = c.Name
		cmap["categoryLevel"] = c.CategoryLevel
		cmap["seq"] = c.Seq
		cmap["status"] = c.Status
		cmap["selectedImg"] = c.BtnSelectedImg
		cmap["btnImg"] = c.BtnImg
		cmap["platformStatus"] = c.PlatformStatus
		if c.CategoryLevel == 1 {
			cmaps2 := make([]map[string]interface{}, 0)
			for _, t := range gameCategories2 {
				cmap2 := make(map[string]interface{})
				if t.ParentId == c.Id {
					cmap2["id"] = t.Id
					cmap2["name"] = t.Name
					cmap2["seq"] = t.Seq
					cmap2["status"] = t.Status
					cmap2["img"] = t.Img
					cmap2["selectedImg"] = t.BtnSelectedImg
					cmap2["btnImg"] = t.BtnImg
					cmap2["platformStatus"] = t.PlatformStatus
					games := make([]xorm.PlatformGames, 0)
					for _, gameLIst := range platGms.([]xorm.PlatformGames) {
						if gameLIst.GameCategorieId == t.Id {
							games = append(games, gameLIst)
						}
					}
					cmap2["games"] = games
					cmaps2 = append(cmaps2, cmap2)
				}
			}
			cmap["categories"] = cmaps2

		} else {
			games := make([]xorm.PlatformGames, 0)
			if c.Id == 1 {
				for _, gameLIst := range platGms.([]xorm.PlatformGames) {
					if gameLIst.Ishot == 1 {
						games = append(games, gameLIst)
					}
				}
			} else {
				for _, gameLIst := range platGms.([]xorm.PlatformGames) {
					if gameLIst.GameCategorieId == c.Id {
						games = append(games, gameLIst)
					}
				}
			}
			cmap["games"] = games
		}
		cmaps = append(cmaps, cmap)
	}
	GameCache.Store(platform, cmaps)

	//加载每个平台的采集记录
}

//异步并行缓存各个平台
func syncCachePlatformTable(platform string, engine *libXorm.Engine, wg *sync.WaitGroup) {
	defer wg.Done()

	//加载活动列表
	if !loadTableActivities(platform, engine) {
		os.Exit(1)
	}
	//加载活动分类表
	if !loadTableActivityClasses(platform, engine) {
		os.Exit(1)
	}

	//加载用户
	if !loadTableUsers(platform, engine) {
		os.Exit(1)
	}
	//加载游戏平台账户
	if !loadTablePlatformAccounts(platform, engine) {
		os.Exit(1)
	}
	//加载投注记录获取的平台key
	if !loadTableBetsKey(platform, engine) {
		os.Exit(1)
	}
	//加载游戏分类表到内存
	if !loadTableGameCategories(platform, engine) {
		os.Exit(1)
	}
	//加载平台游戏列表到内存
	if !loadTablePlatformGames(platform, engine) {
		os.Exit(1)
	}

	//加载平台列表到内存
	if !loadTablePlatforms(platform, engine) {
		os.Exit(1)
	}
	//加载JDB平台游戏列表到内存
	if !loadTablePlatformGamesByJDB(platform, engine) {
		os.Exit(1)
	}

	//加载版本管理表
	if !loadTableAppVersions(platform, engine) {
		os.Exit(1)
	}
	//加载配置表
	if !loadTableConfigs(platform, engine) {
		os.Exit(1)
	}
	//加载系统公告表
	if !loadTableSystemNotices(platform, engine) {
		os.Exit(1)
	}
	//加载vip等级表
	if !loadTableVipLevels(platform, engine) {
		os.Exit(1)
	}
	//加载充值类型表
	if !loadTableChargeTypes(platform, engine) {
		os.Exit(1)
	}
	//加载charge_cards表
	if !loadTableChargeCards(platform, engine) {
		os.Exit(1)
	}
	//加载pay_credential表
	if !loadTablePayCredential(platform, engine) {
		os.Exit(1)
	}
	//加载BetsSearchType
	if !loadBetsSearchType(platform, engine) {
		os.Exit(1)
	}
	//加载TableProxyChessLevels
	if !loadTableProxyChessLevels(platform, engine) {
		os.Exit(1)
	}
	//加载TableProxyRealLevels
	if !loadTableProxyRealLevels(platform, engine) {
		os.Exit(1)
	}
	//加载银行卡列表
	if !loadTableUserBanks(platform, engine) {
		os.Exit(1)
	}
	//加载需要处理的，因异常导致的任务
	if !loadExceptionTasks(platform, engine) {
		os.Exit(1)
	}
	//组装平台游戏类别
	handlerGameCategory(platform)
	fmt.Println("[" + platform + "棋牌游戏]----------------用户信息，游戏分类,系统配置,等等缓存加载完成----------------")
}

func LoadGameCategory(platform string, engine *libXorm.Engine) {
	loadTablePlatforms(platform, engine)
	loadTableGameCategories(platform, engine)
	loadTablePlatformGames(platform, engine)
	handlerGameCategory(platform)
}

//加载表到缓存中
func LoadCache() {
	var wg = sync.WaitGroup{}
	for k, v := range models.MyEngine {
		wg.Add(1)
		go syncCachePlatformTable(k, v, &wg)
	}
	wg.Wait()
}

//加载主库表到缓存里
func MainDbLoadCache() {
	//加载主库platform表
	if !loadMainDbTablePlatform() {
		os.Exit(1)
	}
	fmt.Println("[棋牌游戏]----------------主数据库表缓存加载完毕----------------")

}

// 重载游戏平台配置
func ReloadGamePlatformApiConfig(platform string, gamePlatforms []xorm.Platforms) {
	gamePlatformConfigMap := make(map[string]string)
	for _, gamePlatformRow := range gamePlatforms {
		apiCfgValue := gamePlatformRow.ApiCfgValue
		if apiCfgValue != "" {
			var apiCfgValueJo = make([]map[string]string, 0)
			json.Unmarshal([]byte(apiCfgValue), &apiCfgValueJo)
			var cfgObj = make(map[string]string)
			for _, item := range apiCfgValueJo {
				keyVal, keyExist := item["key"]
				valVal, valExist := item["val"]
				if keyExist && valExist {
					cfgObj[keyVal] = valVal
				}
			}
			cfgBytes, marErr := json.Marshal(cfgObj)
			if marErr != nil {
				apiCfgValue = ""
			} else {
				apiCfgValue = string(cfgBytes)
			}
		}
		gamePlatformConfigMap[gamePlatformRow.Code] = apiCfgValue
	}
	GamePlatformAPiConfigs.Store(platform, gamePlatformConfigMap)
}
