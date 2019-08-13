package models

import (
	"encoding/json"
	"qpgame/admin/common"
	"qpgame/common/utils"
	db "qpgame/models"
	"strconv"
	"strings"
)

//菜单分级说明
//一级菜单: 主菜单, 展现在后台界面左侧第一次显示
//二级菜单: 子菜单, 做为主菜单的子菜单展示, 点击主菜单后显示/隐藏
//三级菜单: 标签页, 点击后台界面的二级菜单后在右侧展示的多个导航标签, 一个标签对应一个url
//四级菜单: 不在菜单显示的隐藏功能, 如锁定资金, 修改用户密码等，并不做为一个显示的菜单显示给用户，只做为功能使用

// 模型
type AdminNodes struct{}

// 表名称
func (self *AdminNodes) GetTableName(ctx *Context) string {
	return "admin_nodes"
}

// 得到所有记录-分页
func (self *AdminNodes) GetRecords(ctx *Context) (Pager, error) {
	conn := db.MyEngine[(*ctx).Params().Get("platform")]
	sql := "SELECT id, title, name, route, seq, level, type, parent_id, status, method, remark FROM admin_nodes" //生成sql语句
	where := ""                                                                                                  //" WHERE id > " + strconv.Itoa(lastId) + " " //where查询条件
	isUnionAll := strings.Index("admin_nodes", " UNION ALL ") > 0                                                //是否是联合查询
	if !isUnionAll {
		queryConditions := getQueryFields(ctx, &map[string]string{
			"name":   "%",
			"route":  "%",
			"method": "%",
			"status": "=",
			"level":  "=",
		})
		for _, v := range queryConditions {
			if v != "" { //必须不能为空
				where += " AND " + v //对查询语句进行累积
			}
		}
		idQuery := getQueryOfUserId(ctx)
		if idQuery != "" {
			where += " AND " + idQuery
		}
	}

	sql += " WHERE 1 = 1 " + where + " ORDER BY level ASC,seq DESC" //拼装sql语句
	rows, err := conn.SQL(sql).QueryString()
	if err != nil {
		return Pager{}, Error{What: "执行查询失败"}
	}

	rowsCount := len(rows)
	if rowsCount < 1 {
		return newPager(), nil
	}

	for _, r := range rows {
		processOptionsFor("level", "level_name", &map[string]string{"1": "一级菜单", "2": "二级菜单", "3": "三级菜单", "4": "四级菜单"}, &r)
		processOptionsFor("type", "type_name", &map[string]string{"0": "主菜单", "1": "子菜单", "2": "标签栏", "3": "动作"}, &r)
		processDatetime(&[]string{"created", "updated"}, &r)
	}

	tmp := make(map[string]int)
	for _, v := range rows {
		tmp[v["id"]] = 1
	}
	list := make([]map[string]string, 0)
	for _, v := range rows {
		id := v["id"]
		if _, ok := tmp[id]; !ok {
			continue
		}
		if (v["level"] == "1") && (v["parent_id"] == "0") {
			list = append(list, v)
			delete(tmp, id)
			for _, vv := range rows {
				id2 := vv["id"]
				if _, ok := tmp[id2]; (vv["level"] == "2") && (vv["parent_id"] == id) && ok {
					list = append(list, vv)
					delete(tmp, id2)

					for _, vvv := range rows {
						id3 := vvv["id"]
						if _, ok := tmp[id3]; (vvv["level"] == "3") && (vvv["parent_id"] == id2) && ok {
							list = append(list, vvv)
							delete(tmp, id3)

							for _, vvvv := range rows {
								id4 := vvvv["id"]
								if _, ok := tmp[id4]; (vvvv["level"] == "4") && (vvvv["parent_id"] == id3) && ok {
									list = append(list, vvvv)
									delete(tmp, id4)
								}
							}
						}
					}
				}
			}
		}
	}
	sql = "SELECT COUNT(*) AS total FROM admin_nodes WHERE 1 = 1 " + where
	countRows, err := conn.SQL(sql).QueryString()
	if err != nil || len(countRows) == 0 {
		return Pager{}, Error{What: "查询统计信息失败"}
	}
	countRow := countRows[0]
	totalCount, err := strconv.Atoi(countRow["total"])
	if err != nil {
		return Pager{}, Error{What: "获取统计信息失败"}
	}
	pager := Pager{
		PageCount: 1,          //总页数
		Page:      1,          //当前页数
		TotalRows: totalCount, //总记录数
		PageSize:  10000,      //每页记录数
		Rows:      list,       //总记录
	}
	return pager, nil
}

// 得到记录详情
func (self *AdminNodes) GetRecordDetail(ctx *Context) (map[string]string, error) {
	return getRecordDetail(ctx, self, "", nil)
}

// 批量排序
func (self *AdminNodes) Sort(ctx *Context) error {
	postData := utils.GetPostData(ctx)
	listStr := postData.Get("list")
	list := make(map[int64]string)
	err := json.Unmarshal([]byte(listStr), &list)
	if err != nil {
		return err
	}
	mdb := db.MyEngine[(*ctx).Params().Get("platform")]
	tablename := self.GetTableName(ctx)
	for id, sort := range list {
		seq := "0"
		s, _ := strconv.Atoi(sort)
		if s >= 0 {
			seq = strconv.Itoa(s)
		}
		sql := "UPDATE " + tablename + " SET seq=" + seq + " WHERE id=" + strconv.Itoa(int(id))
		_, err := mdb.SQL(sql).QueryString()
		if err != nil {
			return err
		}
	}
	// 更新缓存
	{
		common.LoadAdminRoleMenus()
		common.LoadAdmins()
		common.LoadAdminMenus()
		common.LoadAdminRoles()
	}
	return nil
}

// 添加记录
func (self *AdminNodes) Save(ctx *Context) (int64, error) {
	return saveRecord(ctx, self, func(ctx *Context, data *map[string]string) bool { //由上级节点编号得到本级level
		pidStr, ok := (*data)["parent_id"]
		if !ok { //如果不存在上级节点
			return false
		}
		_, err := strconv.Atoi(pidStr)
		if err != nil {
			return false
		}
		if pidStr == "0" {
			(*data)["level"] = "1"
			return true
		}
		sql := "SELECT level FROM admin_nodes WHERE id = " + pidStr //拿到上级节点的相关信息 //TODO: 需要从缓存当中获取
		result, err := db.MyEngine[(*ctx).Params().Get("platform")].SQL(sql).QueryString()
		if err != nil || len(result) == 0 {
			return false
		}
		row := result[0]
		parentLevel := row["level"]
		pId, err := strconv.Atoi(parentLevel)
		if err != nil {
			return false
		}
		(*data)["level"] = strconv.Itoa(pId + 1)
		return true
	}, nil, nil,
		func(ctx *Context, data *map[string]string, isCreating bool) {
			// 更新缓存
			{
				common.LoadAdminRoleMenus()
				common.LoadAdmins()
				common.LoadAdminMenus()
				common.LoadAdminRoles()
			}
			getSavedFunc("菜单节点", "name")(ctx, data, isCreating)
		})
}

// 删除记录
func (self *AdminNodes) Delete(ctx *Context) error {
	err := deleteRecord(ctx, self, nil, getDeletedFunc("菜单节点"))
	// 更新缓存
	if err == nil {
		common.LoadAdminRoleMenus()
		common.LoadAdmins()
		common.LoadAdminMenus()
		common.LoadAdminRoles()
	}
	return err
}
