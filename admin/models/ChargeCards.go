package models

import (
	"github.com/kataras/iris/core/errors"
	"strings"

	"qpgame/admin/common"
	"qpgame/models"

	"strconv"
	"time"
)

// 模型
type ChargeCards struct{}

// 表名称
func (self *ChargeCards) GetTableName(ctx *Context) string {
	return "charge_cards"
}

// 得到所有记录-分页
func (self *ChargeCards) GetRecords(ctx *Context) (Pager, error) {
	ctype := (*ctx).Params().Get("type")
	if ctype == "" {
		return Pager{}, errors.New("查询失败")
	}
	records, err := getRecords(ctx, self,
		"id, name, owner, card_number, bank_address, charge_type_id, created, remark, "+
			"logo, hint, title, mfrom, mto, user_group_ids, amount_limit, addr_type, qr_code, credential_id, priority, state",
		func(ctx *Context) []string { //获取查询条件
			queries := getQueryFields(ctx, &map[string]string{
				"name":        "%",
				"title":       "%",
				"state":       "=",
				"owner":       "=",
				"card_number": "%",
			})
			queries = append(queries, "addr_type IN("+ctype+")")
			return append(queries, getQueryFieldByTime(ctx, "created", "created_start", "created_end"))
		}, func(ctx *Context, row *map[string]string) {
			platform := (*ctx).Params().Get("platform")
			(*row)["charge_type_name"] = common.GetChargeTypeName(platform, (*row)["charge_type_id"])
			processOptionsFor("state", "state_name", &statusTypes, row)
		}, nil)
	conn := models.MyEngine[(*ctx).Params().Get("platform")]
	rows, _ := conn.SQL("select id,group_name from user_groups").QueryString()
	userGroups := make(map[string]string)
	if len(rows) > 0 {
		for _, v := range rows {
			userGroups[v["id"]] = string(v["group_name"])
		}
	}
	credentials := make(map[string]string)
	rows, _ = conn.SQL("SELECT id,pay_name FROM pay_credentials").QueryString()
	if len(rows) > 0 {
		for _, v := range rows {
			credentials[v["id"]] = string(v["pay_name"])
		}
	}
	for k, row := range records.Rows {
		var userGroupIds, userGroupNames []string
		if row["user_group_ids"] != "" {
			userGroupIds = strings.Split(row["user_group_ids"], ",")
			for _, id := range userGroupIds {
				if _, ok := userGroups[id]; ok {
					userGroupNames = append(userGroupNames, userGroups[id])
				}
			}
		}
		credentialId, credentialName := row["credential_id"], ""
		if credentialId != "" {
			if _, ok := credentials[credentialId]; ok {
				credentialName = credentials[credentialId]
			}
		}
		records.Rows[k]["user_groups"] = strings.Join(userGroupNames, ",")
		records.Rows[k]["credential_name"] = credentialName
	}
	return records, err
}

// 得到记录详情
func (self *ChargeCards) GetRecordDetail(ctx *Context) (map[string]string, error) {
	record, err := getRecordDetail(ctx, self, "", nil)
	if err == nil {
		record["credential_name"] = ""
		credentialId := record["credential_id"]
		if credentialId != "" {
			sql := "SELECT pay_name FROM pay_credentials WHERE id=" + credentialId
			rows, err := models.MyEngine[(*ctx).Params().Get("platform")].SQL(sql).QueryString()
			if (err == nil) && (len(rows) > 0) {
				record["credential_name"] = rows[0]["pay_name"]
			}
		}
	}
	return record, nil
}

// 添加记录
func (self *ChargeCards) Save(ctx *Context) (int64, error) {
	return saveRecord(ctx, self, nil,
		func(ctx *Context, data *map[string]string) bool { //添加之前处理
			(*data)["created"] = strconv.FormatInt(time.Now().Unix(), 10) //添加时间
			if (*data)["amount_limit"] == "" {
				(*data)["amount_limit"] = "0"
			}
			return true
		}, nil, getSavedFunc("用户绑卡", "user_id"))
}

// 删除记录
func (self *ChargeCards) Delete(ctx *Context) error {
	return deleteRecord(ctx, self, nil, getDeletedFunc("入款删除"))
}
