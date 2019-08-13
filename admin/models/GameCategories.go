package models

import (
	db "qpgame/models"
	"strconv"
	"time"
)

// 模型
type GameCategories struct{}

//游戏分类
type GameCategory struct {
	Id         int            `json:"id"`         //编号
	Name       string         `json:"name"`       //名称
	Level      int            `json:"level"`      //分类层级
	Categories []GameCategory `json:"categories"` //子级分类
}

// 表名称
func (self *GameCategories) GetTableName(ctx *Context) string {
	return "game_categories"
}

// 得到所有记录-分页
func (self *GameCategories) GetRecords(ctx *Context) (Pager, error) {
	status, err := (*ctx).URLParamInt("status")
	if err != nil {
		status = -1
	}
	sql := "SELECT category_level, id, img, name, parent_id, seq, status, platform_id FROM game_categories"
	if status > -1 {
		sql += " WHERE status=" + strconv.Itoa(status)
	}
	engine := db.MyEngine[(*ctx).Params().Get("platform")]
	rows, err := engine.SQL(sql).QueryString()
	if (err != nil) || (len(rows) < 1) {
		return Pager{}, err
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
		if v["parent_id"] == "0" {
			list = append(list, v)
			delete(tmp, id)
			for _, vv := range rows {
				id2 := vv["id"]
				if _, ok := tmp[id2]; (vv["parent_id"] == id) && ok {
					vv["name"] = vv["name"]
					list = append(list, vv)
					delete(tmp, id2)
					for _, vvv := range rows {
						id3 := vvv["id"]
						if _, ok := tmp[id3]; (vvv["parent_id"] == id2) && ok {
							vvv["name"] = vvv["name"]
							list = append(list, vvv)
							delete(tmp, id3)
							for _, vvvv := range rows {
								id4 := vvvv["id"]
								if _, ok := tmp[id4]; (vvvv["parent_id"] == id3) && ok {
									vvvv["name"] = vvvv["name"]
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
	for k, v := range list {
		items, err := engine.SQL("SELECT a.name gpname,b.name pname FROM game_categories a, platforms b WHERE a.id=" + v["parent_id"] + " AND b.id=" + v["platform_id"]).QueryString()
		if (err != nil) || (len(items) < 1) {
			continue
		}
		list[k]["parent_name"] = items[0]["gpname"]
		list[k]["platform_name"] = items[0]["pname"]
	}
	for _, v := range rows {
		items, err := engine.SQL("SELECT a.name gpname,b.name pname FROM game_categories a, platforms b WHERE a.id=" + v["parent_id"] + " AND b.id=" + v["platform_id"]).QueryString()
		if (err != nil) || (len(items) < 1) {
			continue
		}
		v["parent_name"] = items[0]["gpname"]
		v["platform_name"] = items[0]["pname"]
		if _, ok := tmp[v["id"]]; ok {
			list = append(list, v)
		}
	}
	return Pager{Rows: list, Page: 1, PageCount: 1, PageSize: len(list), TotalRows: len(list)}, nil
}

// 得到记录详情
func (self *GameCategories) GetRecordDetail(ctx *Context) (map[string]string, error) {
	return getRecordDetail(ctx, self, "", nil)
}

// 添加记录
func (self *GameCategories) Save(ctx *Context) (int64, error) {
	return saveRecord(ctx, self, nil,
		func(ctx *Context, data *map[string]string) bool { //添加之前处理
			(*data)["created"] = strconv.FormatInt(time.Now().Unix(), 10) //添加时间
			return true
		}, nil, getSavedFunc("游戏分类", "name"))
}

// 删除记录
func (self *GameCategories) Delete(ctx *Context) error {
	return denyDelete()
}

// 得到关联的游戏分类列表
func (self *GameCategories) Relations(ctx *Context) ([]GameCategory, error) {
	categoryId := (*ctx).URLParam("id")
	if categoryId == "" {
		categoryId = "0"
	}
	categories := getRelatedGameCategories(ctx, categoryId)
	if len(categories) == 0 {
		return nil, Error{What: "没有任何游戏分类"}
	}
	return categories, nil
}

//得到所有下属分类信息
func getRelatedGameCategories(ctx *Context, id string) []GameCategory {
	var categories []GameCategory
	sql := "SELECT id,name,category_level level FROM game_categories WHERE parent_id=" + id
	conn := db.MyEngine[(*ctx).Params().Get("platform")]
	rows, err := conn.SQL(sql).QueryString()
	if (err != nil) || (len(rows) == 0) {
		return categories
	}
	for _, row := range rows {
		categoryId, _ := strconv.Atoi(row["id"])
		level, _ := strconv.Atoi(row["level"])
		categories = append(categories, GameCategory{
			Id:         categoryId,
			Name:       row["name"],
			Level:      level,
			Categories: getRelatedGameCategories(ctx, row["id"]),
		})
	}
	return categories
}
