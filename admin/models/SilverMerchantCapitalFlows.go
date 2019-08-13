package models

import (
	"qpgame/admin/common"
	"qpgame/models"
)

// 模型
type SilverMerchantCapitalFlows struct{}

// 流水类型映射
var merchantCapitalFlowTypes = map[string]string{
	"1": "额度充值",
	"2": "会员充值扣款",
	"3": "额度充值赠送",
	"4": "押金扣除",
}

// 表名称
func (self *SilverMerchantCapitalFlows) GetTableName(ctx *Context) string {
	return "silver_merchant_capital_flows"
}

// 得到所有记录-分页
func (self *SilverMerchantCapitalFlows) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, merchant_id, amount, balance, type, order_id, msg, charged_amount, charged_amount_old, member_user_id, created",
		func(ctx *Context) []string { //获取查询条件
			queries := getQueryFields(ctx, &map[string]string{
				"merchant_id":    "=",
				"type":           "=",
				"order_id":       "%",
				"member_user_id": "=",
			})
			queries = append(queries, getQueryFieldByTime(ctx, "created", "created_start", "created_end"))
			return append(queries)
		},
		func(ctx *Context, row *map[string]string) {
			processOptions("type", &merchantCapitalFlowTypes, row)
			processDatetime(&[]string{"created"}, row)
			user := SilverMerchantUsers{}
			platform := (*ctx).Params().Get("platform")
			(*row)["merchant_name"] = user.GetMerchantName(platform, (*row)["merchant_id"])
			(*row)["user_name"] = ""
			if ((*row)["type"] == "2") && ((*row)["member_user_id"] != "") {
				(*row)["user_name"] = common.GetUserName(platform, (*row)["member_user_id"])
			}
		}, nil)
}

// 得到所有会员充值记录-分页
func (self *SilverMerchantCapitalFlows) GetUserChargeRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, merchant_id, amount, balance, order_id, msg, charged_amount, charged_amount_old, member_user_id, created",
		func(ctx *Context) []string { //获取查询条件
			queries := getQueryFields(ctx, &map[string]string{
				"merchant_id":    "=",
				"order_id":       "%",
				"member_user_id": "=",
			})
			queries = append(queries, "type=2")
			queries = append(queries, getQueryFieldByTime(ctx, "created", "created_start", "created_end"))
			return append(queries)
		},
		func(ctx *Context, row *map[string]string) {
			processDatetime(&[]string{"created"}, row)
			user := SilverMerchantUsers{}
			platform := (*ctx).Params().Get("platform")
			(*row)["merchant_name"] = user.GetMerchantName(platform, (*row)["merchant_id"])
			(*row)["user_name"] = ""
			if (*row)["member_user_id"] != "" {
				(*row)["user_name"] = common.GetUserName(platform, (*row)["member_user_id"])
			}
		}, nil)
}

// 得到报表数据
func (self *SilverMerchantCapitalFlows) GetReports(ctx *Context) (Pager, error) {
	timeQuery := getQueryFieldByTime(ctx, "created", "created_start", "created_end")
	if timeQuery == "" {
		timeQuery = "1=1"
	}
	pager, err := getRecords(ctx, self,
		"from_unixtime(created,'%Y-%m-%d') ymd", func(ctx *Context) []string { //获取查询条件
			queries := getQueryFieldByTime(ctx, "created", "created_start", "created_end")
			return []string{queries}
		}, nil, func(ctx *Context) (string, string, int) {
			return "from_unixtime(created,'%Y-%m-%d')", "ymd DESC", 0
		})
	tablename := self.GetTableName(ctx)
	sql := "SELECT from_unixtime(created,'%Y-%m-%d') ymd,COUNT(DISTINCT id) count,SUM(amount) amount FROM silver_merchant_charge_records WHERE state=1 AND " + timeQuery + " GROUP BY ymd ORDER BY ymd DESC"
	rows, _ := models.MyEngine[(*ctx).Params().Get("platform")].SQL(sql).QueryString()
	tmpA, tmpB := make(map[string]map[string]string), make(map[string]map[string]string)
	for _, v := range rows {
		ymd := v["ymd"]
		tmpA[ymd] = v
	}
	sql = "SELECT from_unixtime(created,'%Y-%m-%d') ymd,COUNT(DISTINCT id) count,SUM(-amount) amount FROM " + tablename + " WHERE type=2 AND " + timeQuery + " GROUP BY ymd ORDER BY ymd DESC"
	rows, _ = models.MyEngine[(*ctx).Params().Get("platform")].SQL(sql).QueryString()
	for _, v := range rows {
		ymd := v["ymd"]
		tmpB[ymd] = v
	}
	for k, row := range pager.Rows {
		row["amount_in"] = "0.000"
		row["amount_out"] = "0.000"
		row["count_in"] = "0"
		row["count_out"] = "0"
		ymd := row["ymd"]
		if v, ok := tmpA[ymd]; ok {
			row["amount_in"] = v["amount"]
			row["count_in"] = v["count"]
		}
		if v, ok := tmpB[ymd]; ok {
			row["amount_out"] = v["amount"]
			row["count_out"] = v["count"]
		}
		pager.Rows[k] = row
	}
	return pager, err
}

// 得到记录详情
func (self *SilverMerchantCapitalFlows) GetRecordDetail(ctx *Context) (map[string]string, error) {
	row, err := getRecordDetail(ctx, self, "", nil)
	processDatetime(&[]string{"created"}, &row)
	if _, ok := row["merchant_name"]; ok {
		user := SilverMerchantUsers{}
		platform := (*ctx).Params().Get("platform")
		row["merchant_name"] = user.GetMerchantName(platform, row["merchant_id"])
		row["user_name"] = ""
		if (row["type"] == "2") && (row["member_user_id"] != "") {
			row["user_name"] = common.GetUserName(platform, row["member_user_id"])
		}
		processOptions("type", &merchantCapitalFlowTypes, &row)
	}
	return row, err
}

// 添加记录
func (self *SilverMerchantCapitalFlows) Save(ctx *Context) (int64, error) {
	return denySave()
}

// 删除记录
func (self *SilverMerchantCapitalFlows) Delete(ctx *Context) error {
	return denyDelete()
}
