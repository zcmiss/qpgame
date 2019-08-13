package ramcache

import (
	"encoding/json"
	"github.com/buger/jsonparser"
	"qpgame/models"
	"qpgame/models/beans"
	"qpgame/models/xorm"
	"strconv"
)

//活动记录表
func platformDbTableActivities(data []byte) {
	activitiesTable := make([]xorm.Activities, 0)
	var rowMap = make(map[string]interface{})
	json.Unmarshal(data, &rowMap)
	platform := rowMap["platform"]
	val, _ := TableActivities.Load(platform)
	activitiesTable = append(activitiesTable, val.([]xorm.Activities)...)
	oper := rowMap["tableOperation"]
	//获取数据
	rowData := rowMap["data"].(map[string]interface{})
	id := int(rowData["id"].(float64))
	//新数据
	var ele = xorm.Activities{}
	ele.Id = id
	ele.Title = rowData["title"].(string)
	ele.SubTitle = rowData["sub_title"].(string)
	ele.Content = rowData["content"].(string)
	ele.Cover = rowData["cover"].(string)
	ele.Created = int(rowData["created"].(float64))
	ele.TimeStart = int(rowData["time_start"].(float64))
	ele.TimeEnd = int(rowData["time_end"].(float64))
	ele.Type = int(rowData["type"].(float64))
	ele.Status = int(rowData["status"].(float64))
	ele.Updated = int(rowData["updated"].(float64))
	ele.ActivityClassId = int(rowData["activity_class_id"].(float64))
	ele.IsHomeShow = int(rowData["is_home_show"].(float64))
	//新增操作
	if oper == "insert" {
		activitiesTable = append(activitiesTable, ele)
		TableActivities.Store(platform, activitiesTable)
		return
	}

	updateRowTable := make([]xorm.Activities, 0)
	//删除操作
	if oper == "delete" {
		for _, v := range activitiesTable {
			if v.Id != id {
				updateRowTable = append(updateRowTable, v)
			}
		}
	}
	//更新操作
	if oper == "update" {
		for _, v := range activitiesTable {
			if v.Id == id {
				updateRowTable = append(updateRowTable, ele)
			} else {
				updateRowTable = append(updateRowTable, v)
			}
		}
	}
	TableActivities.Store(platform, updateRowTable)
}

//活动分类表
func platformDbTableActivityClasses(data []byte) {
	activityClassesTable := make([]xorm.ActivityClasses, 0)
	var rowMap = make(map[string]interface{})
	json.Unmarshal(data, &rowMap)
	platform := rowMap["platform"]
	val, _ := TableActivityClasses.Load(platform)
	activityClassesTable = append(activityClassesTable, val.([]xorm.ActivityClasses)...)
	oper := rowMap["tableOperation"]
	//获取数据
	rowData := rowMap["data"].(map[string]interface{})
	id := int(rowData["id"].(float64))
	//新数据
	var ele = xorm.ActivityClasses{}
	ele.Id = id
	ele.Created = int(rowData["created"].(float64))
	ele.Name = rowData["name"].(string)
	ele.Status = int(rowData["status"].(float64))
	ele.Seq = int(rowData["seq"].(float64))
	ele.Updated = int(rowData["updated"].(float64))

	//新增操作
	if oper == "insert" {
		activityClassesTable = append(activityClassesTable, ele)
		TableActivityClasses.Store(platform, activityClassesTable)
		return
	}

	updateRowTable := make([]xorm.ActivityClasses, 0)
	//删除操作
	if oper == "delete" {
		for _, v := range activityClassesTable {
			if v.Id != id {
				updateRowTable = append(updateRowTable, v)
			}
		}
	}
	//更新操作
	if oper == "update" {
		for _, v := range activityClassesTable {
			if v.Id == id {
				updateRowTable = append(updateRowTable, ele)
			} else {
				updateRowTable = append(updateRowTable, v)
			}
		}
	}
	TableActivityClasses.Store(platform, updateRowTable)
}

//app版本管理表
func platformDbTableAppVersions(data []byte) {
	appVersionTable := make([]xorm.AppVersions, 0)
	var rowMap = make(map[string]interface{})
	json.Unmarshal(data, &rowMap)
	platform := rowMap["platform"]
	val, _ := TableAppVersions.Load(platform)
	appVersionTable = append(appVersionTable, val.([]xorm.AppVersions)...)
	oper := rowMap["tableOperation"]
	//获取数据
	rowData := rowMap["data"].(map[string]interface{})
	id := int(rowData["id"].(float64))
	//新数据
	var ele = xorm.AppVersions{}
	ele.Id = id
	ele.Created = int(rowData["created"].(float64))
	ele.Status = int(rowData["status"].(float64))
	ele.Updated = int(rowData["updated"].(float64))
	ele.Version = rowData["version"].(string)
	ele.Description = rowData["description"].(string)
	ele.Link = rowData["link"].(string)
	ele.PackageType = int(rowData["package_type"].(float64))
	ele.AppType = int(rowData["app_type"].(float64))
	ele.UpdateType = int(rowData["update_type"].(float64))
	//新增操作
	if oper == "insert" {
		appVersionTable = append(appVersionTable, ele)
		TableAppVersions.Store(platform, appVersionTable)
		return
	}

	updateRowTable := make([]xorm.AppVersions, 0)
	//删除操作
	if oper == "delete" {
		for _, v := range appVersionTable {
			if v.Id != id {
				updateRowTable = append(updateRowTable, v)
			}
		}
	}
	//更新操作
	if oper == "update" {
		for _, v := range appVersionTable {
			if v.Id == id {
				updateRowTable = append(updateRowTable, ele)
			} else {
				updateRowTable = append(updateRowTable, v)
			}
		}
	}
	TableAppVersions.Store(platform, updateRowTable)
}

//投注采集key表
func platformDbTableBetsKey(data []byte) {
	bmap := make(map[string]string)
	var rowMap = make(map[string]interface{})
	json.Unmarshal(data, &rowMap)
	platform := rowMap["platform"]
	val, _ := TableBetsKey.Load(platform)
	oper := rowMap["tableOperation"]
	//获取数据
	rowData := rowMap["data"].(map[string]interface{})
	//新数据
	key := strconv.Itoa(int(rowData["plat_id"].(float64))) + "-" + rowData["gt"].(string)
	//复制map
	for k, v := range val.(map[string]string) {
		bmap[k] = v
	}
	//新增操作
	if oper == "insert" {
		bmap[key] = rowData["search_key"].(string)
	}
	//删除操作
	if oper == "delete" {
		delete(bmap, key)
	}
	//更新操作
	if oper == "update" {
		bmap[key] = rowData["search_key"].(string)
	}
	TableBetsKey.Store(platform, bmap)
}

//充值账号表
func platformDbTableChargeCards(data []byte) {
	chargeCardsTable := make([]xorm.ChargeCards, 0)
	var rowMap = make(map[string]interface{})
	json.Unmarshal(data, &rowMap)
	platform := rowMap["platform"]
	val, _ := TableChargeCards.Load(platform)
	chargeCardsTable = append(chargeCardsTable, val.([]xorm.ChargeCards)...)
	oper := rowMap["tableOperation"]
	//获取数据
	rowData := rowMap["data"].(map[string]interface{})
	id := int(rowData["id"].(float64))
	//新数据
	var ele = xorm.ChargeCards{}
	ele.Id = id
	ele.Name = rowData["name"].(string)
	ele.Owner = rowData["owner"].(string)
	ele.CardNumber = rowData["card_number"].(string)
	ele.BankAddress = rowData["bank_address"].(string)
	ele.Remark = rowData["remark"].(string)
	ele.Logo = rowData["logo"].(string)
	ele.Hint = rowData["hint"].(string)
	ele.Title = rowData["title"].(string)
	ele.UserGroupIds = rowData["user_group_ids"].(string)
	ele.QrCode = rowData["qr_code"].(string)
	ele.ChargeTypeId = int(rowData["charge_type_id"].(float64))
	ele.Created = int(rowData["created"].(float64))
	ele.State = int(rowData["state"].(float64))
	ele.Mfrom = int(rowData["mfrom"].(float64))
	ele.Mto = int(rowData["mto"].(float64))
	ele.AmountLimit = int(rowData["amount_limit"].(float64))
	ele.AddrType = int(rowData["addr_type"].(float64))
	ele.CredentialId = int(rowData["credential_id"].(float64))
	ele.Priority = int(rowData["priority"].(float64))
	//新增操作
	if oper == "insert" {
		chargeCardsTable = append(chargeCardsTable, ele)
		TableChargeCards.Store(platform, chargeCardsTable)
		return
	}

	updateRowTable := make([]xorm.ChargeCards, 0)
	//删除操作
	if oper == "delete" {
		for _, v := range chargeCardsTable {
			if v.Id != id {
				updateRowTable = append(updateRowTable, v)
			}
		}
	}
	//更新操作
	if oper == "update" {
		for _, v := range chargeCardsTable {
			if v.Id == id {
				updateRowTable = append(updateRowTable, ele)
			} else {
				updateRowTable = append(updateRowTable, v)
			}
		}
	}
	TableChargeCards.Store(platform, updateRowTable)
}

//充值类型表
func platformDbTableChargeTypes(data []byte) {
	var rowMap = make(map[string]interface{})
	json.Unmarshal(data, &rowMap)
	platform := rowMap["platform"].(string)
	loadTableChargeTypes(platform, models.MyEngine[platform])
}

//系统配置
func platformDbTableConfigs(data []byte) {
	cfgs := make(map[string]interface{})
	var rowMap = make(map[string]interface{})
	json.Unmarshal(data, &rowMap)
	platform := rowMap["platform"]
	val, _ := TableConfigs.Load(platform)
	oper := rowMap["tableOperation"]
	//获取数据
	rowData := rowMap["data"].(map[string]interface{})
	key := rowData["name"].(string)
	value := rowData["value"].(string)
	//复制map
	for k, v := range val.(map[string]interface{}) {
		cfgs[k] = v
	}
	//新增或者更新操作
	if oper == "insert" || oper == "update" {
		var tempV interface{}
		err := json.Unmarshal([]byte(value), &tempV)
		if err != nil {
			tempV = value
		}
		cfgs[key] = tempV
	}
	//删除操作
	if oper == "delete" {
		delete(cfgs, key)
	}

	TableConfigs.Store(platform, cfgs)

}

//游戏分类
func platformDbTableGameCategories(data []byte) {
	gameCategoriesTable := make([]xorm.GameCategories, 0)
	var rowMap = make(map[string]interface{})
	json.Unmarshal(data, &rowMap)
	platform := rowMap["platform"]
	val, _ := TableGameCategories.Load(platform)
	gameCategoriesTable = append(gameCategoriesTable, val.([]xorm.GameCategories)...)
	oper := rowMap["tableOperation"]
	//获取数据
	rowData := rowMap["data"].(map[string]interface{})
	id := int(rowData["id"].(float64))
	//新数据
	var ele = xorm.GameCategories{}
	ele.Id = id
	ele.Created = int(rowData["created"].(float64))
	ele.Status = int(rowData["status"].(float64))
	ele.CategoryLevel = int(rowData["category_level"].(float64))
	ele.ParentId = int(rowData["parent_id"].(float64))
	ele.Seq = int(rowData["seq"].(float64))
	ele.Img = rowData["img"].(string)
	ele.Name = rowData["name"].(string)
	ele.Rate = strconv.FormatFloat(rowData["rate"].(float64), 'f', -1, 64)
	ele.BtnSelectedImg = rowData["btn_selected_img"].(string)
	ele.BtnImg = rowData["btn_img"].(string)
	//新增操作
	if oper == "insert" && ele.Status == 1 {
		gameCategoriesTable = append(gameCategoriesTable, ele)
		TableGameCategories.Store(platform, gameCategoriesTable)
		return
	}
	updateRowTable := make([]xorm.GameCategories, 0)
	//删除操作
	if oper == "delete" {
		for _, v := range gameCategoriesTable {
			if v.Id != id {
				updateRowTable = append(updateRowTable, v)
			}
		}
	}
	//更新操作
	if oper == "update" {
		isNoExist := false
		for _, v := range gameCategoriesTable {
			if v.Id == id {
				isNoExist = true
				//如果不是1就丢弃相当于删除
				if ele.Status == 1 {

					updateRowTable = append(updateRowTable, ele)
				}
			} else {
				updateRowTable = append(updateRowTable, v)
			}
		}
		//如果之前status = 0的话，缓存中肯定不存在,当改为1的时候要重新插入
		if !isNoExist && ele.Status == 1 {
			updateRowTable = append(updateRowTable, ele)
		}
	}
	TableGameCategories.Store(platform, updateRowTable)
	handlerGameCategory(platform.(string))

}

//暂时用不到
func platformDbTableNotices(data []byte) {}

//第三方支付证书
func platformDbTablePayCredentials(data []byte) {
	var rowMap = make(map[string]interface{})
	json.Unmarshal(data, &rowMap)
	platform := rowMap["platform"].(string)
	loadTablePayCredential(platform, models.MyEngine[platform])
}

//第三方用户游戏账号
func platformDbTablePlatformAccounts(data []byte) {
	var rowMap = make(map[string]interface{})
	json.Unmarshal(data, &rowMap)
	platform := rowMap["platform"].(string)
	loadTablePlatformAccounts(platform, models.MyEngine[platform])
}

//第三方平台游戏
func platformDbTablePlatformGames(data []byte) {
	platformGamesTable := make([]xorm.PlatformGames, 0)
	var rowMap = make(map[string]interface{})
	json.Unmarshal(data, &rowMap)
	platform := rowMap["platform"].(string)
	val, _ := TablePlatformGames.Load(platform)
	platformGamesTable = append(platformGamesTable, val.([]xorm.PlatformGames)...)
	oper := rowMap["tableOperation"]
	//获取数据
	rowData := rowMap["data"].(map[string]interface{})
	id := int(rowData["id"].(float64))
	//新数据
	var ele = xorm.PlatformGames{}
	ele.Id = id
	ele.Img = rowData["img"].(string)
	ele.Name = rowData["name"].(string)
	ele.GameCategorieId = int(rowData["game_categorie_id"].(float64))
	ele.Ishot = int(rowData["ishot"].(float64))
	ele.Isnew = int(rowData["isnew"].(float64))
	ele.Isrecommend = int(rowData["isrecommend"].(float64))
	ele.Ishidden = int(rowData["ishidden"].(float64))
	ele.PlatId = int(rowData["plat_id"].(float64))
	ele.GameUrl = rowData["game_url"].(string)
	ele.Gamecode = rowData["gamecode"].(string)
	ele.Gt = rowData["gt"].(string)
	ele.ServiceCode = rowData["service_code"].(string)
	ele.SmallImg = rowData["small_img"].(string)
	//新增操作
	if oper == "insert" && ele.Ishidden == 0 {
		platformGamesTable = append(platformGamesTable, ele)
		TablePlatformGames.Store(platform, platformGamesTable)
		return
	}
	updateRowTable := make([]xorm.PlatformGames, 0)
	//删除操作
	if oper == "delete" {
		for _, v := range platformGamesTable {
			if v.Id != id {
				updateRowTable = append(updateRowTable, v)
			}
		}
	}
	//更新操作
	if oper == "update" {
		isNoExist := false
		for _, v := range platformGamesTable {
			if v.Id == id {
				isNoExist = true
				//如果不是1就丢弃相当于删除
				if ele.Ishidden == 0 {
					updateRowTable = append(updateRowTable, ele)
				}
			} else {
				updateRowTable = append(updateRowTable, v)
			}
		}
		//如果之前Ishidden = 0的话，缓存中肯定不存在,当改为1的时候要重新插入
		if !isNoExist && ele.Ishidden == 0 {
			updateRowTable = append(updateRowTable, ele)
		}
	}
	TablePlatformGames.Store(platform, updateRowTable)
	handlerGameCategory(platform)
}

//平台状态变更，游戏分类进行变更
func platformDbTablePlatforms(data []byte) {
	var rowMap = make(map[string]interface{})
	json.Unmarshal(data, &rowMap)
	platform := rowMap["platform"].(string)
	LoadGameCategory(platform, models.MyEngine[platform])
}

//棋牌代理等级
func platformDbTableProxyChessLevels(data []byte) {
	var rowMap = make(map[string]interface{})
	json.Unmarshal(data, &rowMap)
	platform := rowMap["platform"].(string)
	loadTableProxyChessLevels(platform, models.MyEngine[platform])
}

//真人视讯代理等级
func platformDbTableProxyRealLevels(data []byte) {
	var rowMap = make(map[string]interface{})
	json.Unmarshal(data, &rowMap)
	platform := rowMap["platform"].(string)
	loadTableProxyRealLevels(platform, models.MyEngine[platform])
}

//系统通知
func platformDbTableSystemNotices(data []byte) {
	var rowMap = make(map[string]interface{})
	json.Unmarshal(data, &rowMap)
	platform := rowMap["platform"].(string)
	loadTableSystemNotices(platform, models.MyEngine[platform])
}

//银行卡列表
func platformDbTableUserBanks(data []byte) {
	var rowMap = make(map[string]interface{})
	json.Unmarshal(data, &rowMap)
	platform := rowMap["platform"].(string)
	loadTableUserBanks(platform, models.MyEngine[platform])
}

//用户表
func platformDbTableUsers(data []byte) {
	id, _, _, _ := jsonparser.Get(data, "data", "id")
	token, _, _, _ := jsonparser.Get(data, "data", "token")
	tokenCreated, _, _, _ := jsonparser.Get(data, "data", "token_created")
	userGroupId, _, _, _ := jsonparser.Get(data, "data", "user_group_id")
	platform, _, _, _ := jsonparser.Get(data, "platform")
	operB, _, _, _ := jsonparser.Get(data, "tableOperation")
	phone, _ := jsonparser.GetString(data, "data", "phone")
	userName, _ := jsonparser.GetString(data, "data", "user_name")
	valp, _ := PhoneNumAndToken.Load(string(platform))
	phoneToken := valp.(map[string][]string)
	valu, _ := UserNameAndToken.Load(string(platform))
	userNameToken := valu.(map[string][]string)

	uic, _ := UserIdCard.Load(string(platform))
	uicMap := uic.(map[int]beans.UserProfile)

	oper := string(operB)
	//新数据
	var value = make([]string, 0)
	value = append(value, string(id))
	value = append(value, string(token))
	value = append(value, string(tokenCreated))
	value = append(value, string(userGroupId))
	iUserId, _ := strconv.Atoi(string(id))
	//新增操作
	if oper == "insert" || oper == "update" {
		if phone != "" {
			phoneToken[phone] = value
		}
		if userName != "" {
			userNameToken[userName] = value
		}
		if iUserId > 0 {
			iTokenCreated, _ := strconv.Atoi(string(tokenCreated))
			uicMap[iUserId] = beans.UserProfile{
				Phone:        phone,
				Username:     userName,
				Token:        string(token),
				TokenCreated: iTokenCreated,
				UserGroupId:  string(userGroupId),
			}
		}
	}
	//删除操作
	if oper == "delete" {
		if phone != "" {
			delete(phoneToken, phone)
		}
		if userName != "" {
			delete(userNameToken, userName)
		}
		if iUserId > 0 {
			delete(uicMap, iUserId)
		}
	}
}

//暂时用不到
func platformDbTableUserGroups(data []byte) {

}

//vip等级
func platformDbTableVipLevels(data []byte) {
	var rowMap = make(map[string]interface{})
	json.Unmarshal(data, &rowMap)
	platform := rowMap["platform"].(string)
	loadTableVipLevels(platform, models.MyEngine[platform])
}
