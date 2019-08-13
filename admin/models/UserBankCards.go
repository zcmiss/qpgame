package models

import (
	"qpgame/models"
	"strconv"
	"time"
)

// 模型
type UserBankCards struct{}

// 表名称
func (self *UserBankCards) GetTableName(ctx *Context) string {
	return "user_bank_cards"
}

// 得到所有记录-分页
func (self *UserBankCards) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, user_id, card_number, address, created, bank_name, name, status, updated, remark",
		func(ctx *Context) []string { //获取查询条件
			var conditions []string                              //查询条件数组
			if id, err := (*ctx).URLParamInt("id"); err == nil { //按编号查询
				conditions = append(conditions, "id = "+strconv.Itoa(id))
			}
			return conditions
		},
		func(ctx *Context, row *map[string]string) {
			processOptions("status", &statusTypes, row)
			processDatetime(&[]string{"created", "updated"}, row)
		}, nil)
}

// 得到记录详情
func (self *UserBankCards) GetRecordDetail(ctx *Context) (map[string]string, error) {
	userIdStr := (*ctx).URLParam("id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return map[string]string{}, Error{What: "错误的用户编号"}
	}

	fields := "id, user_id, card_number, address, created, bank_name, name, status, updated, remark" //要提取的字段
	sql := "SELECT " + fields + " FROM " + self.GetTableName(ctx) + " WHERE user_id = '" + strconv.Itoa(userId) + "' LIMIT 1"
	conn := models.MyEngine[(*ctx).Params().Get("platform")]
	rows, rowsErr := conn.SQL(sql).QueryString()
	if rowsErr != nil || len(rows) <= 0 {
		return map[string]string{}, Error{What: "无法查找到相应的记录"}
	}
	return rows[0], nil
}

// 添加记录
func (self *UserBankCards) Save(ctx *Context) (int64, error) {
	return saveRecord(ctx, self, nil,
		func(ctx *Context, data *map[string]string) bool { //添加之前处理
			conn := models.MyEngine[(*ctx).Params().Get("platform")]
			userId, name := (*data)["user_id"], (*data)["name"]
			_, err := conn.Exec("UPDATE users SET name='" + name + "' WHERE id=" + userId)
			if err != nil {
				return false
			}
			(*data)["created"] = strconv.FormatInt(time.Now().Unix(), 10) //添加时间
			return true
		},
		func(ctx *Context, data *map[string]string) bool { //修改之前处理
			conn := models.MyEngine[(*ctx).Params().Get("platform")]
			userId, name := (*data)["user_id"], (*data)["name"]
			_, err := conn.Exec("UPDATE users SET name='" + name + "' WHERE id=" + userId)
			if err != nil {
				return false
			}
			(*data)["updated"] = strconv.FormatInt(time.Now().Unix(), 10) //修改时间
			return true
		}, getSavedFunc("用户绑卡", "card_number"))
}

// 删除记录
func (self *UserBankCards) Delete(ctx *Context) error {
	return denyDelete()
}
