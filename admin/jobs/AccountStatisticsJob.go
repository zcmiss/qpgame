package jobs

import (
	"strconv"
	"strings"
	"time"

	"github.com/go-xorm/xorm"
)

// 用户报表统计
type AccountStatisticsJob struct{}

// 时间设置
func (self *AccountStatisticsJob) GetSpec() string {
	return "0 0 3 * * *"
}

// 处理作务
func (self *AccountStatisticsJob) Process(db *xorm.Engine, platform string) {
	yesterday := time.Now().AddDate(0, 0, -1)
	self.CountByDate(db, &yesterday)
}

// 按年月日进行处理
func (self *AccountStatisticsJob) CountByDate(db *xorm.Engine, day *time.Time) {
	ymd, timeStart, timeEnd := getCountTimes(day)
	timeStartStr, timeEndStr, page := strconv.Itoa(timeStart), strconv.Itoa(timeEnd), 1
	for {
		idArr := getUserIds(db, page, 100, timeEnd)
		if idArr == nil {
			return
		}
		self.CountByUserIds(db, idArr, timeStartStr, timeEndStr, ymd)
		page += 1 //对页数进行累加
	}
}

// 按每个用户进行统计
func (self *AccountStatisticsJob) CountByUserIds(db *xorm.Engine, ids []string, timeStart string, timeEnd string, ymd int) {
	between := " BETWEEN " + timeStart + " AND " + timeEnd
	idWhere := " WHERE user_id IN (" + strings.Join(ids, ",") + ")"
	groupBy := " GROUP BY user_id"
	betQueries := getBetSqls(ids, func(index int, userIds string) string {
		return "SELECT user_id, 'bet_amount' AS `type`, SUM(amount) AS total, SUM(reward) AS reward FROM bets_" + strconv.Itoa(index) +
			" WHERE user_id IN (" + userIds + ") AND created " + between + groupBy
	})
	sqlArr := append(betQueries,
		"SELECT user_id, 'charge_online' AS `type`, SUM(amount) AS total, 0 AS reward FROM charge_records"+idWhere+" AND state = 1 AND created "+between+groupBy,       //充值总额-公司入款+线上入款
		"SELECT user_id, 'charge_manual' AS `type`, SUM(amount) AS total, 0 AS reward FROM manual_charges"+idWhere+" AND state = 1 AND deal_time "+between+groupBy,     //充值总额-人工入款
		"SELECT user_id, 'withdraw_online' AS `type`, SUM(amount) AS total, 0 AS reward FROM withdraw_records"+idWhere+" AND status = 1 AND created "+between+groupBy,  //提现总额-线上提现
		"SELECT user_id, 'withdraw_manual' AS `type`, SUM(amount) AS total, 0 AS reward FROM manual_withdraws"+idWhere+" AND state = 1 AND deal_time "+between+groupBy, //提现总额-人工出款
		"SELECT user_id, 'wash_amount' AS `type`, SUM(amount) AS total, 0 AS reward FROM wash_code_records "+idWhere+" AND washtime "+between+groupBy,                  //洗码
		"SELECT user_id, 'proxy_commission' AS `type`, SUM(commission) AS total, 0 AS reward FROM proxy_commissions"+idWhere+" AND created "+between+groupBy,           //代理佣金
	)
	sql := "(" + strings.Join(sqlArr, ") UNION ALL (") + ")"
	rows, err := db.SQL(sql).QueryString()
	if err != nil || len(rows) == 0 { //表示没有任何记录
		return
	}
	records := map[string]map[string]string{}
	fields := map[string]string{
		"charged_amount":   "0", //充值总额
		"consumed_amount":  "0", //消费总额
		"withdraw_amount":  "0", //提现总额
		"bet_amount":       "0", //投注总额
		"reward_amount":    "0", //中奖总额
		"wash_amount":      "0", //洗码总额
		"proxy_commission": "0", //佣金总额
	}
	ymdStr := strconv.Itoa(ymd)                                               //统计日期
	processRecord := func(row map[string]string, record *map[string]string) { //针对于每个记录进行处理
		rowType := row["type"]
		if rowType == "bet_amount" {
			(*record)["bet_amount"] = row["total"]
			(*record)["reward_amount"] = row["reward"]
		} else if rowType == "charge_online" || rowType == "charge_manual" {
			charged, _ := strconv.ParseFloat((*record)["charged_amount"], 64)
			rowTotal, _ := strconv.ParseFloat(row["total"], 64)
			charged += rowTotal
			(*record)["charged_amount"] = strconv.FormatFloat(charged, 'f', 3, 64)
		} else if rowType == "withdraw_manual" || rowType == "withdraw_online" {
			withdraw, _ := strconv.ParseFloat((*record)["withdraw_amount"], 64)
			rowTotal, _ := strconv.ParseFloat(row["total"], 64)
			withdraw += rowTotal
			(*record)["charged_amount"] = strconv.FormatFloat(withdraw, 'f', 3, 64)
		} else {
			(*record)[row["type"]] = row["total"]
		}
	}
	for _, row := range rows {
		userId := row["user_id"]
		record, has := records[userId]
		if !has {
			record = map[string]string{
				"ymd":     ymdStr,
				"user_id": userId,
			}
			for field, value := range fields {
				record[field] = value
			}
		}
		processRecord(row, &record)
		records[userId] = record
	}

	for _, r := range records {
		r["consumed_amount"] = r["bet_amount"]
		saveRecord(db, &r, "account_statistics")
	}
}
