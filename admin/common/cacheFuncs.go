package common

import (
	db "qpgame/models"
	"strconv"
	"strings"
)

// 拿到指定平台的所有角色编号
func getRoleIds(platform string) []int {
	conn := db.MyEngine[platform]
	sql := "SELECT id,name FROM admin_roles WHERE status=1"
	rows, err := conn.SQL(sql).QueryString()
	if err != nil {
		return []int{}
	}
	ids := []int{}
	for _, row := range rows {
		id, _ := strconv.Atoi(row["id"])
		if id < 1 {
			continue
		}
		ids = append(ids, id)
	}
	return ids
}

// 依据角色编号得到相应的菜单
func getSubMenusById(list []map[string]string, id string) []MenuNode {
	nodes := []MenuNode{}
	rows := make([]map[string]string, 0)
	for _, v := range list {
		pid := v["parent_id"]
		if pid == id {
			rows = append(rows, v)
		}
	}
	for _, row := range rows {
		levelStr, idStr, title := row["level"], row["id"], row["title"]
		subMenus := getSubMenusById(list, idStr)
		id, _ := strconv.Atoi(idStr)
		level, _ := strconv.Atoi(levelStr)
		nodes = append(nodes, MenuNode{
			Id:    id,
			Title: title,
			Level: level,
			Nodes: subMenus,
			Url:   row["route"],
		})
	}
	return nodes
}

// 依据角色得到相应菜单权限
func getMenusByRoleId(platform string, roleId int) []MenuNode {
	// 相应平台的数据库连接
	conn := db.MyEngine[platform]
	nodes := []MenuNode{}
	// 列出所有的本角色的菜单编号
	rows, _ := conn.SQL("SELECT menu_ids FROM admin_roles WHERE id=" + strconv.Itoa(roleId)).QueryString()
	if len(rows) < 1 {
		return nodes
	}
	// 此角色下的所有菜单编号
	menuIds := rows[0]["menu_ids"]
	if mids := strings.Split(menuIds, ","); (menuIds == "") || (len(mids) < 1) {
		return nodes
	}
	// 列出所有菜单,主菜单降序排列
	list, _ := conn.SQL("SELECT id,parent_id,level,route,title FROM admin_nodes WHERE id IN(" + menuIds + ") AND status=1 ORDER BY level asc,seq DESC").QueryString()
	for _, v := range list {
		levelStr, idStr, title := v["level"], v["id"], v["title"]
		if levelStr == "1" {
			subMenus := getSubMenusById(list, idStr)
			id, _ := strconv.Atoi(idStr)
			level, _ := strconv.Atoi(levelStr)
			nodes = append(nodes, MenuNode{
				Id:    id,
				Title: title,
				Level: level,
				Url:   "#",
				Nodes: subMenus,
			})
		}
	}
	return nodes
}

//LoadAll: 加载所有需要后台面要的相关配置
func LoadAll() {
	LoadAdminRoleMenus()
	LoadAdmins()
	LoadAdminMenus()
	LoadAdminRoles()
	LoadAdminApiConfigs()
	LoadAdminTokens()
	LoadGameCategories()
	LoadGamePlatforms()
	LoadPlatformGames()
	LoadActivities()
	LoadActivityClasses()
	LoadFrontendUsers()
	LoadChargetTypes()
	LoadChargetCards()
	LoadThirdPayments()
	LoadGameCodes()
	LoadRedpackets() // 加载红包数据
}
