package models

import (
	"strconv"
	"strings"
	"time"

	"qpgame/admin/common"
	"qpgame/admin/validations"
	"qpgame/common/utils"
	db "qpgame/models"
)

// 模型
type SilverMerchantUsers struct{}

var silverMerchantUsersValidation = validations.SilverMerchantUsersValidation{}

// 表名称
func (self *SilverMerchantUsers) GetTableName(ctx *Context) string {
	return "silver_merchant_users"
}

// 得到所有记录-分页
func (self *SilverMerchantUsers) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, user_id, merchant_level, usable_amount, merchant_cash_pledge, total_charge_money, total_auth_amount, "+
			"donate_rate, account, created, status, is_destroy, "+
			"token, last_login_time, merchant_name",
		func(ctx *Context) []string { //获取查询条件
			queries := getQueryFields(ctx, &map[string]string{
				"merchant_level": "=",
				"account":        "%",
				"merchant_name":  "%",
				"is_destroy":     "=",
				"status":         "=",
			})
			queries = append(queries, getQueryFieldByTime(ctx, "created", "created_start", "created_end"))
			return append(queries, getQueryFieldByTime(ctx, "last_login_time", "last_login_start", "last_login_end"))
		},
		func(ctx *Context, row *map[string]string) {
			processDatetime(&[]string{"last_login_time", "token_created", "created"}, row) //添加时间/最后登录时间
			platform := (*ctx).Params().Get("platform")
			(*row)["user_name"] = common.GetUserName(platform, (*row)["user_id"])
		}, nil)
}

// 得到记录详情
func (self *SilverMerchantUsers) GetRecordDetail(ctx *Context) (map[string]string, error) {
	result := map[string]string{}
	id, platform := (*ctx).URLParam("id"), (*ctx).Params().Get("platform")
	engine := db.MyEngine[platform]
	//
	sql := "SELECT id,merchant_name,user_id,merchant_level,usable_amount,merchant_cash_pledge,total_charge_money,total_auth_amount," +
		"donate_rate,account,is_destroy,last_login_time,created,status FROM silver_merchant_users WHERE id=" + id
	rows, err := engine.SQL(sql).QueryString()
	if err != nil || len(rows) < 1 {
		return map[string]string{}, nil
	}
	user := rows[0]
	user["user_name"] = common.GetUserName(platform, user["user_id"])
	processDatetime(&[]string{"last_login_time", "created"}, &user)
	for k, v := range user {
		result["info_"+k] = v
	}

	sql = "SELECT * FROM silver_merchant_bank_cards WHERE merchant_id=" + id
	rows, err = engine.SQL(sql).QueryString()
	if (err == nil) && (len(rows) > 0) {
		for k, v := range rows[0] {
			result["bank_card_"+k] = v
		}
	}

	sql = "SELECT type, SUM(amount) amount,COUNT(DISTINCT order_id) count FROM silver_merchant_capital_flows WHERE merchant_id=" + id + " GROUP BY type"
	rows, err = engine.SQL(sql).QueryString()
	if (err == nil) && (len(rows) > 0) {
		for _, row := range rows {
			switch row["type"] {
			case "1":
				result["flow_charge_amount"] = row["amount"]
				result["flow_charge_count"] = row["count"]
			case "3":
				result["flow_presented_money_amount"] = row["amount"]
			case "2":
				result["flow_user_charge_amount"] = strings.Replace(row["amount"], "-", "", -1)
				result["flow_user_charge_count"] = row["count"]
			}
		}
	}
	return result, nil
}

// 添加记录
func (self *SilverMerchantUsers) Save(ctx *Context) (int64, error) {
	return saveRecord(ctx, self, nil,
		func(ctx *Context, data *map[string]string) bool { //添加之前处理
			if _, ok := (*data)["user_id"]; (!ok) || ((*data)["user_id"] == "") {
				return false
			}
			user_id := (*data)["user_id"]
			conn := db.MyEngine[(*ctx).Params().Get("platform")]
			rows, err := conn.SQL("SELECT id FROM " + self.GetTableName(ctx) + " WHERE user_id=" + user_id).QueryString()
			if (err != nil) || (len(rows) > 0) {
				return false
			}
			(*data)["password"] = utils.MD5((*data)["password"])
			(*data)["token_created"] = "0"
			(*data)["created"] = strconv.FormatInt(time.Now().Unix(), 10) //添加时间
			return true
		}, func(ctx *Context, data *map[string]string) bool { //更新之前处理
			if _, ok := (*data)["user_id"]; (!ok) || ((*data)["user_id"] == "") {
				return false
			}
			if _, ok := (*data)["password"]; ok {
				(*data)["password"] = utils.MD5((*data)["password"])
			}
			(*data)["updated"] = strconv.FormatInt(time.Now().Unix(), 10) //更新时间
			return true
		}, getSavedFunc("银商用户信息", "merchant_name"))
}

// 删除记录
func (self *SilverMerchantUsers) Delete(ctx *Context) error {
	return deleteRecord(ctx, self, nil, getDeletedFunc("银商用户"))
}

//修改银商用户密码
func (self *SilverMerchantUsers) UpdatePassword(ctx *Context) error {
	//提交信息的校验
	errMessages, err := silverMerchantUsersValidation.CheckPass(ctx)
	if !err {
		return Error{What: errMessages}
	}
	post := utils.GetPostData(ctx)
	pass := utils.MD5(post.Get("old_password")) //旧的密码
	idStr := post.Get("id")
	sql := "SELECT id FROM silver_merchant_users WHERE id = '" + idStr + "' AND `password` = '" + pass + "'"
	conn := db.MyEngine[(*ctx).Params().Get("platform")]
	rows, errRows := conn.SQL(sql).QueryString()
	if errRows != nil {
		return Error{What: "银商用户信息不存在或旧密码错误"}
	}
	row := rows[0]
	if row["password"] == pass {
		return Error{What: "新密码与旧密码不能一致"}
	}
	newPass := utils.MD5(post.Get("password"))
	sql = "UPDATE silver_merchant_users SET `password` = '" + newPass + "' WHERE id = '" + idStr + "'"
	result, errUpdate := conn.Exec(sql)
	if errUpdate != nil {
		return Error{What: "保存信息失败"}
	}
	affectedRows, errSave := result.RowsAffected()
	if errSave != nil || affectedRows <= 0 {
		return Error{What: "保存新的密码失败"}
	}
	return nil
}

//修改银商用户状态
func (self *SilverMerchantUsers) changeStatus(ctx *Context, field string, toStatus string, title string) error {
	idStr := (*ctx).URLParam("id")
	_, err := strconv.Atoi(idStr)
	if err != nil {
		return Error{What: "必须提供银商用户编号"}
	}
	conn := db.MyEngine[(*ctx).Params().Get("platform")]
	sql := "UPDATE silver_merchant_users SET " + field + " = " + toStatus + " WHERE id = '" + idStr + "' LIMIT 1"
	result, err := conn.Exec(sql)
	if err != nil {
		return Error{What: "更改" + title + "状态失败"}
	}
	affectedRows, err := result.RowsAffected()
	if err != nil || affectedRows <= 0 {
		return Error{What: "更改" + title + "状态失败"}
	}
	return nil
}

//锁定用户
func (self *SilverMerchantUsers) Lock(ctx *Context) error {
	return self.changeStatus(ctx, "status", "0", "银商用户")
}

//解锁用户
func (self *SilverMerchantUsers) Unlock(ctx *Context) error {
	return self.changeStatus(ctx, "status", "1", "银商用户")
}

// 获取银商名字
func (self *SilverMerchantUsers) GetMerchantName(platform string, id string) string {
	sql := "SELECT merchant_name AS name FROM silver_merchant_users WHERE id = '" + id + "' LIMIT 1"
	conn := db.MyEngine[platform]
	rows, err := conn.SQL(sql).QueryString()
	if (err != nil) || (len(rows) == 0) {
		return ""
	}
	row := rows[0]
	if row == nil {
		return ""
	}
	return row["name"]
}
