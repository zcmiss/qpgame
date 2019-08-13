package common

import (
	"strconv"
	"strings"

	"qpgame/config"
	"qpgame/models"
	db "qpgame/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
)

// 得到角色名称
func GetRoleName(platform string, roleId int) string {
	idStr := strconv.Itoa(roleId)
	key := platform + "-" + idStr
	name, exists := AdminRoles[key]
	if exists {
		return name
	}
	row := getSingleRow(platform, "SELECT id, name FROM admin_roes WHERE id = '"+strconv.Itoa(roleId)+"' LIMIT 1")
	if row != nil {
		name = row["name"]
		AdminRoles[key] = name
		return name
	}
	return ""
}

// 得到角色相应的菜单信息
func GetMenusByRole(platform string, roleId int) []MenuNode {
	idStr := strconv.Itoa(roleId)
	key := platform + "-" + idStr
	menus, has := AdminRoleMenus[key]
	if !has {
		return nil
	}

	return menus
}

// 得到当前保存到用户token信息
func GetAdminToken(ctx *iris.Context, id int) string {
	platform := (*ctx).Params().Get("platform") //平台标识
	idStr := strconv.Itoa(id)                   //转换为字符串
	key := platform + "-" + idStr               //生成key
	token, exists := AdminTokens[key]           //检测当前key是否存在
	if exists {                                 //如果存在，直接返回
		return token
	}

	conn := db.MyEngine[platform]                                //数据库连接
	sql := "SELECT token FROM admins WHERE id = '" + idStr + "'" //查询数据库
	rows, err := conn.SQL(sql).QueryString()
	if err != nil {
		return ""
	}

	row := rows[0]
	return row["token"]
}

// 得到api配置
func GetApiConfig(name string) ApiConfig {
	conf := ApiConfig{ //默认的相关配置
		Code:   "CKYX",
		Name:   "棋牌系统",
		Url:    "api.admin-qp.com",
		ApiUrl: "http://api.admin-qp.com/CKYX",
	}
	//先从缓存当中读取
	//如果级存当中没有, 则从数据库读取
	//如果数据库没有，则返回默认值
	return conf
}

// 依据id得到游戏平台名称
func GetGamePlatformName(platform string, id string) string {
	key := platform + "-" + id
	name, exists := GamePlatforms[key]
	if exists {
		return name
	}
	row := getSingleRow(platform, "SELECT name FROM platforms WHERE id = '"+id+"' LIMIT 1")
	if row != nil {
		name = row["name"]
		GamePlatforms[key] = name
		return name
	}
	return ""
}

// 依据id得到平台游戏名称
func GetPlatformGameName(platform string, id string) string {
	key := platform + "-" + id
	name, exists := PlatformGames[key]
	if exists {
		return name
	}
	row := getSingleRow(platform, "SELECT name FROM Platform_games WHERE id = '"+id+"' LIMIT 1")
	if row != nil {
		name = row["name"]
		PlatformGames[key] = name
		return name
	}
	return ""
}

// 依据id得到平台游戏名称
func GetGameNameOfCode(platform string, code string) string {
	key := platform + "-" + code
	name, exists := GameCodes[key]
	if exists {
		return name
	}
	row := getSingleRow(platform, "SELECT service_code AS code, name FROM Platform_games WHERE service_code = '"+code+"' LIMIT 1")
	if row != nil {
		name = row["name"]
		GameCodes[key] = name
		return name
	}
	return ""
}

// 依据id得到游戏分类名称
func GetGameCategoryName(platform string, id string) string {
	key := platform + "-" + id
	name, exists := GameCategories[key]
	if exists {
		return name
	}
	row := getSingleRow(platform, "SELECT name FROM game_categories WHERE id = '"+id+"' LIMIT 1")
	if row != nil {
		name = row["name"]
		GameCategories[key] = name
		return name
	}
	return ""
}

// 依据id得到活动分类名称
func GetActivityClassName(platform string, id string) string {
	key := platform + "-" + id
	name, exists := ActivityClasses[key]
	if exists {
		return name
	}
	row := getSingleRow(platform, "SELECT name FROM activity_classes WHERE id = '"+id+"' LIMIT 1")
	if row != nil {
		name = row["name"]
		ActivityClasses[key] = name
		return name
	}
	return ""
}

// 依据id得到活动名称
func GetActivityName(platform string, id string) string {
	key := platform + "-" + id
	name, exists := Activities[key]
	if exists {
		return name
	}
	row := getSingleRow(platform, "SELECT title AS name FROM activities WHERE id = '"+id+"' LIMIT 1")
	if row != nil {
		name = row["name"]
		Activities[key] = name
		return name
	}
	return ""
}

// 依据用户编号得到用户名称
func GetUserName(platform string, id string) string {
	name, exists := FrontendUsers[platform][id]
	if exists {
		return name
	}
	sql := "SELECT user_name AS name FROM users WHERE id = '" + id + "' LIMIT 1"
	row := getSingleRow(platform, sql)
	if row == nil {
		return ""
	}

	FrontendUsers[platform][id] = row["name"]
	return row["name"]
}

// 依据用户名称获取id
func GetIdFromUserName(platform string, name string) string {
	users, exists := FrontendUsers[platform]
	if !exists {
		return ""
	}

	for k, v := range users {
		if v != "" && v == name {
			return k
		}
	}

	sql := "SELECT user_name AS name FROM users WHERE user_name = '" + name + "' LIMIT 1"
	row := getSingleRow(platform, sql)
	if row == nil {
		return ""
	}

	FrontendUsers[platform][row["id"]] = row["name"]
	return row["id"]
}

// 得到单条记录, 便于统一读取缓存
func getSingleRow(platform string, sql string) map[string]string {
	conn := models.MyEngine[platform]
	rows, err := conn.SQL(sql).QueryString()
	if err != nil || len(rows) == 0 {
		return nil
	}
	return rows[0]
}

//得到充值类型名称
func GetChargeTypeName(platform string, id string) string {
	key := platform + "-" + id
	name, exists := ChargeTypes[key]
	if exists {
		return name
	}
	row := getSingleRow(platform, "SELECT name FROM charge_types WHERE id = '"+id+"' LIMIT 1")
	if row != nil {
		name = row["name"]
		ChargeTypes[key] = name
		return name
	}
	return ""
}

//得到充值类型名称
func GetChargeCardName(platform string, id string) string {
	key := platform + "-" + id
	name, exists := ChargeCards[key]
	if exists {
		return name
	}
	row := getSingleRow(platform, "SELECT name, owner FROM charge_cards WHERE id = '"+id+"' LIMIT 1")
	if row != nil {
		name = row["name"] + "|" + row["owner"]
		ChargeCards[key] = name
		return name
	}
	return ""
}

//得到支付方式名称
func GetThirdPaymentName(platform string, id string) string {
	key := platform + "-" + id
	name, exists := ThirdPayments[key]
	if exists {
		return name
	}
	row := getSingleRow(platform, "SELECT pay_name AS name FROM pay_credentials WHERE id = '"+id+"' LIMIT 1")
	if row != nil {
		name = row["name"]
		ThirdPayments[key] = name
		return name
	}
	return ""
}

// 得到箮理员相关信息
// map[string]string {
// id: 后台用户编号
// name: 后台用户名称
// }
func GetAdmin(ctx *iris.Context) map[string]string {
	authString := (*ctx).GetHeader("Authorization") //提交過來的授權字符串
	tokenString := strings.Replace(authString, "bearer ", "", 7)
	requestToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.TokenKey), nil
	})
	if err != nil || requestToken == nil {
		return nil
	}
	claim, _ := requestToken.Claims.(jwt.MapClaims) //解码-提交的token
	return map[string]string{
		"id":   claim["id"].(string),
		"name": claim["name"].(string),
	}
}
