package models

import (
	db "qpgame/models"
	"strconv"
	"strings"
	"time"
)

// 模型
type SystemStatistics struct{}

// 表名称
func (self *SystemStatistics) GetTableName(ctx *Context) string {
	return "system_statistics"
}

// 得到所有记录-分页
func (self *SystemStatistics) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, ymd, charge, withdraw, deductions, bet_amount, bet_count, charge_count, charge_user_count, "+
			"withdraw_count, withdraw_user_count, sale_ratio, winning, proxy_ratio, active, user_win, "+
			"profit, reg_user, bet_new, deposit_user, first_charge_amount, first_charge_user, app_win, "+
			"downline_bet_user, member_user",
		func(ctx *Context) []string { //获取查询条件
			queries := getQueryFieldByDate(ctx, "ymd", "ymd_start", "ymd_end")
			return []string{queries}
		},
		func(ctx *Context, row *map[string]string) { //对于查询出来的每条记录的处理
			ymd := (*row)["ymd"]
			(*row)["ymd"] = ymd[:4] + "-" + ymd[4:6] + "-" + ymd[6:]
		}, func(ctx *Context) (string, string, int) {
			return "", "ymd DESC", -1
		})
}

// 得到记录详情
func (self *SystemStatistics) GetRecordDetail(ctx *Context) (map[string]string, error) {
	return getRecordDetail(ctx, self, "",
		func(ctx *Context, row *map[string]string) { //对于查询出来的每条记录的处理
			if createTime, createErr := strconv.ParseInt((*row)["created"], 10, 64); createErr == nil {
				(*row)["created"] = time.Unix(createTime, 0).Format("2006-01-02 15:04:05") //添加时间
			}
		})
}

// 添加记录
func (self *SystemStatistics) Save(ctx *Context) (int64, error) {
	return denySave()
}

// 删除记录
func (self *SystemStatistics) Delete(ctx *Context) error {
	return denyDelete()
}

// 出入款汇总
func (self *SystemStatistics) Counts(ctx *Context) (map[string]string, error) {
	def := map[string]string{
		"in_total":             "0.000",
		"in_total_count":       "0",
		"in_company":           "0.000",
		"in_company_count":     "0",
		"in_online":            "0.000",
		"in_online_count":      "0",
		"in_manual":            "0.000",
		"in_manual_count":      "0",
		"in_deduction":         "0.000",
		"in_deduction_count":   "0",
		"in_first":             "0.000",
		"in_first_count":       "0",
		"out_total":            "0.000",
		"out_total_count":      "0",
		"out_online":           "0.000",
		"out_online_count":     "0",
		"out_manual":           "0.000",
		"out_manual_count":     "0",
		"out_offer":            "0.000",
		"out_offer_count":      "0",
		"out_commission":       "0.000",
		"out_commission_count": "0",
		"out_first":            "0.000",
		"out_first_count":      "0",
		"total":                "0.000",
		"balance":              "0.000",
		"remain":               "0.000",
		"no_settlement":        "0.000",
	}

	fields := "SUM(charge) in_total,SUM(charge_user_count) in_total_count," + //总入款, 总入款人数
		"SUM(charge_company) in_company,SUM(charge_company_count) in_company_count," + //公司入款, 公司入款人数
		"SUM(charge_online) in_online,SUM(charge_online_count) in_online_count," + //线上入款, 线上入款人数
		"SUM(charge_manual) in_manual,SUM(charge_manual_count) in_manual_count," + //人工入款, 人工入款人数
		"SUM(-withdraw) in_deduction,SUM(withdraw_user_count) in_deduction_count," + //出款扣除, 出款扣除人数
		"SUM(first_charge_amount) in_first,SUM(first_charge_user) in_first_count," + //首次充值, 首次充值人数
		"SUM(-withdraw) out_total,SUM(withdraw_user_count) out_total_count," + //总出款, 总出款人数
		"SUM(withdraw_online) out_online,SUM(withdraw_online_count) out_online_count," + //线上出款, 线上出款人数
		"SUM(withdraw_manual) out_manual,SUM(withdraw_manual_count) out_manual_count," + //人工出款, 人工出款人数
		"SUM(active) out_offer,SUM(active_user_count) out_offer_count," + //活动优惠, 活动优惠人数
		"SUM(sale_ratio) out_commission,SUM(sale_ratio_user_count) out_commission_count," + //给予反水, 给予反水人数
		"SUM(first_withdraw_amount) out_first,SUM(first_withdraw_user) out_first_count," + //首次出款, 首次出款人数
		"SUM(charge+withdraw) total,SUM(app_win) balance" //总计, 平台实际盈亏
	//"SUM(profit) remain,SUM(give_win) no_settlement" //可用金额, 未结算金额
	where := ""
	cond := getQueryFieldByDate(ctx, "ymd", "ymd_start", "ymd_end")
	if cond != "" {
		where += " WHERE " + cond
	}
	conn := db.MyEngine[(*ctx).Params().Get("platform")]
	sql := "SELECT " + fields + " FROM " + self.GetTableName(ctx) + where
	rows, err := conn.SQL(sql).QueryString()
	if err != nil || len(rows) == 0 {
		return def, nil
	}
	row := rows[0]
	for k := range def {
		if row[k] != "" {
			def[k] = row[k]
		}
	}
	sql = "SELECT SUM(a.balance_wallet) remain FROM accounts a,users u WHERE a.user_id=u.id AND u.status=1 AND u.user_type=0"
	rows, err = conn.SQL(sql).QueryString()
	if (err == nil) && (len(rows) > 0) {
		def["remain"] = rows[0]["remain"]
	}
	return def, nil
}

// 首次存入明细
func (self *SystemStatistics) CountsFirstCharge(ctx *Context) (Pager, error) {
	dateFrom, dateTo := (*ctx).URLParam("date_from"), (*ctx).URLParam("date_to")
	where := ""
	if dateFrom != "" || dateTo != "" {
		if dateFrom != "" {
			where += " AND SUBSTRING_INDEX(FROM_UNIXTIME(a.created),' ',1)>='" + dateFrom + "'"
		}
		if dateTo != "" {
			where = where + " AND SUBSTRING_INDEX(FROM_UNIXTIME(a.created),' ',1)<='" + dateTo + "'"
		}
	}
	engine := db.MyEngine[(*ctx).Params().Get("platform")]
	page := (*ctx).URLParamIntDefault("page", 1)
	if page < 1 {
		page = 1
	}
	limit := (*ctx).URLParamIntDefault("page_size", 20)
	if limit < 1 {
		limit = 20
	}
	offset := (page - 1) * limit
	sql := "SELECT a.*,u.user_name FROM account_infos a,users u WHERE a.type=1 AND a.user_id=u.id " + where + " AND concat(a.user_id,'_',a.created)IN(SELECT DISTINCT concat(user_id,'_',min(created)) FROM account_infos WHERE type=1 group by user_id) group by a.user_id order by a.created desc "
	rows, err := engine.SQL(sql).QueryString()
	totalCount := len(rows)
	sql = sql + "LIMIT " + strconv.Itoa(offset) + "," + strconv.Itoa(limit)
	rows, err = engine.SQL(sql).QueryString()
	if err != nil {
		return Pager{}, nil
	}
	pageCount := 0 //总的页数
	if totalCount > 0 {
		num := totalCount % limit
		if num == 0 {
			pageCount = totalCount / limit
		} else {
			pageCount = totalCount/limit + 1
		}
	}
	for k, v := range rows {
		created, _ := strconv.Atoi(v["created"])
		rows[k]["created"] = time.Unix(int64(created), 0).Format("2006-01-02 15:04:05")
	}
	pager := Pager{
		PageCount: pageCount,  //总页数
		Page:      page,       //当前页数
		TotalRows: totalCount, //总记录数
		PageSize:  limit,      //每页记录数
		Rows:      rows,       //总记录
	}
	return pager, nil
}

// 首次出款明细
func (self *SystemStatistics) CountsFirstWithdraw(ctx *Context) (Pager, error) {
	dateFrom, dateTo := (*ctx).URLParam("date_from"), (*ctx).URLParam("date_to")
	where := ""
	if dateFrom != "" || dateTo != "" {
		if dateFrom != "" {
			where = " AND SUBSTRING_INDEX(FROM_UNIXTIME(a.created),' ',1)>='" + dateFrom + "'"
		}
		if dateTo != "" {
			where = where + " AND SUBSTRING_INDEX(FROM_UNIXTIME(a.created),' ',1)<='" + dateTo + "'"
		}
	}
	engine := db.MyEngine[(*ctx).Params().Get("platform")]
	page := (*ctx).URLParamIntDefault("page", 1)
	if page < 1 {
		page = 1
	}
	limit := (*ctx).URLParamIntDefault("page_size", 20)
	if limit < 1 {
		limit = 20
	}
	offset := (page - 1) * limit
	sql := "SELECT a.*,u.user_name FROM account_infos a,users u WHERE a.type=2 AND a.user_id=u.id " + where + " AND concat(a.user_id,'_',a.created)IN(SELECT DISTINCT concat(user_id,'_',min(created)) FROM account_infos WHERE type=2 group by user_id) group by a.user_id order by a.created desc "
	rows, err := engine.SQL(sql).QueryString()
	totalCount := len(rows)
	sql = sql + "LIMIT " + strconv.Itoa(offset) + "," + strconv.Itoa(limit)
	rows, err = engine.SQL(sql).QueryString()
	if err != nil {
		return Pager{}, nil
	}
	pageCount := 0 //总的页数
	if totalCount > 0 {
		num := totalCount % limit
		if num == 0 {
			pageCount = totalCount / limit
		} else {
			pageCount = totalCount/limit + 1
		}
	}
	for k, v := range rows {
		created, _ := strconv.Atoi(v["created"])
		amount := strings.Replace(v["amount"], "-", "", 0)
		rows[k]["amount"] = amount
		rows[k]["created"] = time.Unix(int64(created), 0).Format("2006-01-02 15:04:05")
	}
	pager := Pager{
		PageCount: pageCount,  //总页数
		Page:      page,       //当前页数
		TotalRows: totalCount, //总记录数
		PageSize:  limit,      //每页记录数
		Rows:      rows,       //总记录
	}
	return pager, nil
}

// 返水明细
func (self *SystemStatistics) CountsBackWater(ctx *Context) (Pager, error) {
	dateFrom, dateTo := (*ctx).URLParam("date_from"), (*ctx).URLParam("date_to")
	where := ""
	if dateFrom != "" || dateTo != "" {
		if dateFrom != "" {
			where = " AND SUBSTRING_INDEX(FROM_UNIXTIME(a.created),' ',1)>='" + dateFrom + "'"
		}
		if dateTo != "" {
			where = where + " AND SUBSTRING_INDEX(FROM_UNIXTIME(a.created),' ',1)<='" + dateTo + "'"
		}
	}
	engine := db.MyEngine[(*ctx).Params().Get("platform")]
	page := (*ctx).URLParamIntDefault("page", 1)
	if page < 1 {
		page = 1
	}
	limit := (*ctx).URLParamIntDefault("page_size", 20)
	if limit < 1 {
		limit = 20
	}
	offset := (page - 1) * limit
	sql := "SELECT a.*,u.user_name FROM account_infos a,users u WHERE a.type=3 AND a.user_id=u.id " + where
	rows, err := engine.SQL(sql).QueryString()
	totalCount := len(rows)
	sql = sql + " LIMIT " + strconv.Itoa(offset) + "," + strconv.Itoa(limit)
	rows, err = engine.SQL(sql).QueryString()
	if err != nil {
		return Pager{}, nil
	}
	pageCount := 0 //总的页数
	if totalCount > 0 {
		num := totalCount % limit
		if num == 0 {
			pageCount = totalCount / limit
		} else {
			pageCount = totalCount/limit + 1
		}
	}
	for k, v := range rows {
		created, _ := strconv.Atoi(v["created"])
		rows[k]["created"] = time.Unix(int64(created), 0).Format("2006-01-02 15:04:05")
	}
	pager := Pager{
		PageCount: pageCount,  //总页数
		Page:      page,       //当前页数
		TotalRows: totalCount, //总记录数
		PageSize:  limit,      //每页记录数
		Rows:      rows,       //总记录
	}
	return pager, nil
}
