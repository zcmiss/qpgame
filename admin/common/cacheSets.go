package common

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"qpgame/config"
	db "qpgame/models"
	"strconv"
)

// 保存登录的token
func StoreLoginedToken(platform string, id int, token string) bool {
	idStr := strconv.Itoa(id)
	key := platform + "-" + idStr
	_, exists := AdminTokens[key]
	if exists {
		delete(AdminTokens, key) //如果已经存在，则删除
	}
	AdminTokens[key] = token                                                              //保存到内存当中
	conn := db.MyEngine[platform]                                                         //连接数据库
	sql := "UPDATE admins SET token = '" + token + "' WHERE id = '" + idStr + "' LIMIT 1" //修改用户信息表
	result, err := conn.Exec(sql)
	if err != nil {
		delete(AdminTokens, key)
		return false
	}

	affected, err := result.RowsAffected()
	if err != nil || affected <= 0 {
		delete(AdminTokens, key)
		return false
	}

	return true
}

// 移除用户登录信息
func RemoveLoginedToken(platform string, tokenStr string) error {
	requestToken, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.TokenKey), nil
	})
	if err != nil {
		return Error{What: "解析Token失败: 无效的TOKEN"}
	}
	claim, _ := requestToken.Claims.(jwt.MapClaims)     //解码
	userId, idErr := strconv.Atoi(claim["id"].(string)) //拿到用戶編號
	if idErr != nil {
		return Error{What: "解码Token失败: 无法获取用户编号"}
	}
	idStr := strconv.Itoa(userId)
	key := platform + "-" + idStr //生成key
	delete(AdminTokens, key)      //从当前缓存当中删除

	conn := db.MyEngine[platform] //数据库连接
	sql := "SELECT token FROM admins WHERE id=" + idStr
	rows, err := conn.SQL(sql).QueryString()
	if (err != nil) || (len(rows) == 0) {
		return Error{What: "用户不存在"}
	}
	if rows[0]["token"] == "" {
		return nil
	}
	sql = "UPDATE admins SET token = '' WHERE id = '" + idStr + "' LIMIT 1" //修改token为空
	result, err := conn.Exec(sql)
	if err != nil {
		return Error{What: "用户TOKEN已失效"}
	}
	affected, err := result.RowsAffected()
	if err != nil || affected <= 0 {
		return Error{What: "操作无效, 退出失败"}
	}
	return nil
}

// 加载所有的已登录的token
func LoadAdminTokens() {
	AdminTokens = make(map[string]string)
	for platform := range config.PlatformCPs {
		rows, _ := db.MyEngine[platform].SQL("SELECT id,token FROM admins").QueryString()
		if len(rows) < 1 {
			continue
		}
		for _, row := range rows {
			key := platform + "-" + row["id"]
			AdminTokens[key] = row["token"]
		}
	}
}

// 加载所有平台的api配置
func LoadAdminApiConfigs() {
	AdminApiConfigs = make(map[string]ApiConfig)
	rows, _ := db.MyEngineMainDb.SQL("SELECT code,name,admin_address FROM platform").QueryString()
	if len(rows) < 1 {
		fmt.Println("加载平台信息出错")
		return
	}
	AdminApiConfigs = map[string]ApiConfig{}
	for _, row := range rows {
		key := row["code"]
		AdminApiConfigs[key] = ApiConfig{
			Code:   row["code"],
			Name:   row["name"],
			Url:    row["admin_address"],
			ApiUrl: row["admin_address"],
		}
	}
}

//加载所有后台的菜单
func LoadAdminMenus() {
	AdminMenus = make(map[string][]MenuNode)
	for platform := range config.PlatformCPs {
		rows, _ := db.MyEngine[platform].SQL("SELECT id,title,route FROM admin_nodes").QueryString()
		if len(rows) == 0 {
			continue
		}
		menus := []MenuNode{}
		for _, row := range rows {
			menuId, _ := strconv.Atoi(row["id"])
			menus = append(menus, MenuNode{
				Id:    menuId,
				Title: row["title"],
				Level: 0,
				Url:   row["route"],
				Nodes: nil,
			})
		}
		AdminMenus[platform] = menus
	}
}

// 加载所有的角色信息
// key: 平台标识号 + 角色编号, value: 角色信息
func LoadAdminRoles() {
	AdminRoles = make(map[string]string)
	for platform := range config.PlatformCPs {
		rows, _ := db.MyEngine[platform].SQL("SELECT id,name FROM admin_roles").QueryString()
		if len(rows) < 1 {
			fmt.Println("[棋牌遊戲] 平臺", platform, "後臺沒有任何角色相關信息")
			continue
		}
		for _, row := range rows {
			key := platform + "-" + row["id"]
			AdminRoles[key] = row["name"]
		}
	}
}

// 加载所有的管理员信息 key: 平台标识 + 管理员编号, value: 管理员信息
func LoadAdmins() {
	Admins = make(map[string]AdminUser)
	for platform := range config.PlatformCPs {
		rows, _ := db.MyEngine[platform].SQL("SELECT id,name,password,status,role_id FROM admins").QueryString()
		if len(rows) < 1 {
			fmt.Println("[棋牌游戏] 错误: 无法加载平台", platform, "任何后台用户信息!")
			continue
		}
		for _, row := range rows {
			key := platform + "-" + row["id"]
			adminId, _ := strconv.Atoi(row["id"])
			roleId, _ := strconv.Atoi(row["role_id"])
			status, _ := strconv.Atoi(row["status"])
			Admins[key] = AdminUser{
				Id:       adminId,
				RoleId:   roleId,
				Status:   status,
				Name:     row["name"],
				Password: row["password"],
			}
		}
	}
}

// 加载所有的菜单信息
// key: 平台标识号 + 角色编号, value : 菜单信息
func LoadAdminRoleMenus() {
	AdminRoleMenus = make(map[string][]MenuNode)
	for platform := range config.PlatformCPs {
		for _, roleId := range getRoleIds(platform) {
			key := platform + "-" + strconv.Itoa(roleId)
			AdminRoleMenus[key] = getMenusByRoleId(platform, roleId)
		}
	}
}

// 加载所有的游戏分类信息
func LoadGameCategories() {
	for platform := range config.PlatformCPs {
		rows, _ := db.MyEngine[platform].SQL("SELECT id,name FROM game_categories").QueryString()
		for _, row := range rows {
			key := platform + "-" + row["id"]
			GameCategories[key] = row["name"]
		}
	}
}

// 加载所有的游戏平台编号/名称
func LoadGamePlatforms() {
	for platform := range config.PlatformCPs {
		rows, _ := db.MyEngine[platform].SQL("SELECT id,name FROM platforms").QueryString()
		for _, row := range rows {
			key := platform + "-" + row["id"]
			GamePlatforms[key] = row["name"]
		}
	}
}

// 加载所有的平台游戏编号/名称
func LoadPlatformGames() {
	for platform := range config.PlatformCPs {
		rows, _ := db.MyEngine[platform].SQL("SELECT id,name FROM platform_games").QueryString()
		for _, row := range rows {
			key := platform + "-" + row["id"]
			GameCategories[key] = row["name"]
		}
	}
}

// 加载所有的平台游戏编号/名称
func LoadGameCodes() {
	for platform := range config.PlatformCPs {
		rows, _ := db.MyEngine[platform].SQL("SELECT service_code code,name FROM platform_games").QueryString()
		for _, row := range rows {
			key := platform + "-" + row["code"]
			GameCodes[key] = row["name"]
		}
	}
}

// 加载活动分类
func LoadActivityClasses() {
	for platform := range config.PlatformCPs {
		rows, _ := db.MyEngine[platform].SQL("SELECT id,name FROM activity_classes").QueryString()
		for _, row := range rows {
			key := platform + "-" + row["id"]
			ActivityClasses[key] = row["name"]
		}
	}
}

// 加载全部的活动信息
// key: 平台名称-
func LoadActivities() {
	for platform := range config.PlatformCPs {
		rows, _ := db.MyEngine[platform].SQL("SELECT id,title FROM activities").QueryString()
		for _, row := range rows {
			key := platform + "-" + row["id"]
			Activities[key] = row["title"]
		}
	}
}

//加载全部的前台用户ID/名称
func LoadFrontendUsers() {
	for platform := range config.PlatformCPs {
		rows, _ := db.MyEngine[platform].SQL("SELECT id,user_name name FROM users").QueryString()
		users := map[string]string{}
		for _, row := range rows {
			users[row["id"]] = row["name"]
		}
		FrontendUsers[platform] = users
	}
}

//加载充会值类型
func LoadChargetTypes() {
	for platform := range config.PlatformCPs {
		rows, _ := db.MyEngine[platform].SQL("SELECT id,name FROM charge_types").QueryString()
		for _, row := range rows {
			key := platform + "-" + row["id"]
			ChargeTypes[key] = row["name"]
		}
	}
}

//加载充会值类型
func LoadChargetCards() {
	for platform := range config.PlatformCPs {
		rows, _ := db.MyEngine[platform].SQL("SELECT id,name,owner FROM charge_cards").QueryString()
		for _, row := range rows {
			key := platform + "-" + row["id"]
			ChargeCards[key] = row["name"] + "|" + row["owner"]
		}
	}
}

//加载充会值类型
func LoadThirdPayments() {
	for platform := range config.PlatformCPs {
		rows, _ := db.MyEngine[platform].SQL("SELECT id,pay_name name FROM pay_credentials").QueryString()
		for _, row := range rows {
			key := platform + "-" + row["id"]
			ThirdPayments[key] = row["name"]
		}
	}
}

//加载红包数据
func LoadRedpackets() {
	for platform := range config.PlatformCPs {
		Redpackets[platform] = make(map[string]map[string]string)
		rows, _ := db.MyEngine[platform].SQL("SELECT id,type,money,total,status,end_time,start_time,calculate_type,is_done FROM redpacket_systems").QueryString()
		for _, row := range rows {
			Redpackets[platform][row["id"]] = row
		}
	}
}
