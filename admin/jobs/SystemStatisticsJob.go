package jobs

import (
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/go-xorm/xorm"
)

// 运营报表统计
type SystemStatisticsJob struct{}

// 时间设置
func (self *SystemStatisticsJob) GetSpec() string {
	return "0 10 3 * * *"
}

// 处理作务
func (self *SystemStatisticsJob) Process(db *xorm.Engine, platform string) {
	yesterday := time.Now().AddDate(0, 0, -1)
	self.CountByDate(db, &yesterday)
}

// 按年月日进行处理
func (self *SystemStatisticsJob) CountByDate(db *xorm.Engine, day *time.Time) {
	ymd, timeStart, timeEnd := getCountTimes(day)
	//开始后续的统计处理
	timeStartStr := strconv.Itoa(timeStart) //开始时间
	timeEndStr := strconv.Itoa(timeEnd)     //结束时间
	data := self.GetStatisticsData(db, timeStartStr, timeEndStr, ymd)
	if data != nil {
		self.SaveData(db, data.(map[string]string))
	}
}

// 获取统计数据
func (self *SystemStatisticsJob) GetStatisticsData(db *xorm.Engine, timeStart string, timeEnd string, ymd int) interface{} {
	// 默认的要写入数据表的数据结构
	record := map[string]interface{}{
		"ymd": ymd, //统计日期

		"bet_amount": float64(0.), //投注总额
		"bet_count":  0,           //投注次数
		"winning":    float64(0),  //中奖总额

		"charge":            float64(0), //充值总额
		"charge_count":      0,          //充值次数
		"charge_user_count": 0,          //充值人数

		"withdraw":            float64(0), //提现总额
		"withdraw_count":      0,          //提现次数
		"withdraw_user_count": 0,          //提现人数

		"profit": float64(0), //利润

		"sale_ratio":            float64(0), // 销售返点
		"sale_ratio_user_count": 0,          // 销售返点人数
		"active":                float64(0), //活动奖励
		"active_user_count":     0,          // 活动奖励人数
		"proxy_ratio":           float64(0), //佣金总额

		"first_charge_amount": float64(0), //首充金额
		"first_charge_user":   0,          //首充人数

		"reg_user": 0, // 新增注册人数

		"charge_company":        float64(0), // 公司入款金额
		"charge_company_count":  0,          // 公司入款人数
		"charge_online":         float64(0), // 线上入款金额
		"charge_online_count":   0,          // 线上入款人数
		"charge_manual":         float64(0), // 手工入款金额
		"charge_manual_count":   0,          // 手工入款人数
		"withdraw_online":       float64(0), // 线上出款金额
		"withdraw_online_count": 0,          // 线上出款人数
		"withdraw_manual":       float64(0), // 手工出款金额
		"withdraw_manual_count": 0,          // 手工出款人数

		"app_win":    float64(0), // 实际盈亏
		"deductions": float64(0), // 总扣除
	}
	//
	between := timeStart + " AND " + timeEnd
	// 投注/中奖统计
	sql := make([]string, 0)
	for i := 0; i < 10; i++ {
		sql = append(sql, "(SELECT SUM(amount) amount,COUNT(id) count,SUM(reward) reward FROM bets_"+strconv.Itoa(i)+" WHERE created BETWEEN "+between+")")
	}
	rows, _ := db.SQL("SELECT SUM(a.amount) amount,SUM(a.count) count,SUM(a.reward) reward FROM (" + strings.Join(sql, "UNION") + ") a").QueryString()
	if len(rows) > 0 {
		record["bet_amount"], _ = strconv.ParseFloat(rows[0]["amount"], 64)
		record["bet_count"], _ = strconv.Atoi(rows[0]["count"])
		record["winning"], _ = strconv.ParseFloat(rows[0]["reward"], 64)
	}
	// 充值/提现/洗码
	rows, _ = db.SQL("SELECT type,COUNT(DISTINCT user_id) ucount,SUM(amount) amount,COUNT(id) count FROM account_infos WHERE type IN(1,2,3,9,10) AND created BETWEEN " + between + " GROUP BY type").QueryString()
	if len(rows) > 0 {
		for _, v := range rows {
			switch v["type"] {
			case "1": // 充值
				record["charge"], _ = strconv.ParseFloat(v["amount"], 64)
				record["charge_count"], _ = strconv.Atoi(v["count"])
				record["charge_user_count"], _ = strconv.Atoi(v["ucount"])
			case "2": // 提现
				record["withdraw"], _ = strconv.ParseFloat(v["amount"], 64)
				record["withdraw_count"], _ = strconv.Atoi(v["count"])
				record["withdraw_user_count"], _ = strconv.Atoi(v["ucount"])
			case "3": // 洗码
				record["sale_ratio"], _ = strconv.ParseFloat(v["amount"], 64)
				record["sale_ratio_user_count"], _ = strconv.Atoi(v["count"])
			case "9": // 活动奖励
				record["active"], _ = strconv.ParseFloat(v["amount"], 64)
				record["active_user_count"], _ = strconv.Atoi(v["count"])
			case "10": // 提现失败返还
				amount, _ := strconv.ParseFloat(v["amount"], 64)
				record["withdraw"] = record["withdraw"].(float64) - amount
			}
		}
		record["profit"] = record["charge"].(float64) - record["withdraw"].(float64)
	}
	// 代理返点
	rows, _ = db.SQL("SELECT SUM(commission) amount FROM proxy_commissions WHERE created BETWEEN " + between).QueryString()
	if len(rows) > 0 {
		record["proxy_ratio"], _ = strconv.ParseFloat(rows[0]["amount"], 64)
	}
	// 首次充值
	rows, _ = db.SQL("SELECT COUNT(DISTINCT user_id) count,SUM(amount) amount FROM account_infos WHERE concat(user_id,'_',created)IN(SELECT concat(user_id, '_', min(created)) FROM account_infos WHERE type=1 GROUP BY user_id) AND type=1 AND created BETWEEN " + between).QueryString()
	if len(rows) > 0 {
		record["first_charge_amount"], _ = strconv.ParseFloat(rows[0]["amount"], 64)
		record["first_charge_user"], _ = strconv.Atoi(rows[0]["count"])
	}
	// 首次提现
	rows, _ = db.SQL("SELECT COUNT(DISTINCT user_id) count,SUM(-amount) amount FROM account_infos WHERE concat(user_id,'_',created)IN(SELECT concat(user_id, '_', min(created)) FROM account_infos WHERE type=2 GROUP BY user_id) AND type=2 AND created BETWEEN " + between).QueryString()
	if len(rows) > 0 {
		record["first_withdraw_amount"], _ = strconv.ParseFloat(rows[0]["amount"], 64)
		record["first_withdraw_user"], _ = strconv.Atoi(rows[0]["count"])
	}
	// 新注册人数
	rows, _ = db.SQL("SELECT COUNT(DISTINCT id) count FROM users WHERE created BETWEEN " + between).QueryString()
	if len(rows) > 0 {
		record["reg_user"], _ = strconv.Atoi(rows[0]["count"])
	}
	// 公司入款/线上入款
	rows, _ = db.SQL("SELECT is_tppay,COUNT(DISTINCT user_id) count,SUM(amount) amount FROM charge_records WHERE state=1 AND updated BETWEEN " + between + " GROUP BY is_tppay").QueryString()
	if len(rows) > 0 {
		for _, v := range rows {
			switch v["is_tppay"] {
			case "0":
				record["charge_company"], _ = strconv.ParseFloat(v["amount"], 64)
				record["charge_company_count"], _ = strconv.Atoi(v["count"])
			case "1":
				record["charge_online"], _ = strconv.ParseFloat(v["amount"], 64)
				record["charge_online_count"], _ = strconv.Atoi(v["count"])
			}
		}
	}
	// 手工入款
	rows, _ = db.SQL("SELECT COUNT(DISTINCT user_id) count,SUM(amount) amount FROM manual_charges WHERE (audit=0 OR (audit=1 AND state=1)) AND deal_time BETWEEN " + between).QueryString()
	if len(rows) > 0 {
		record["charge_manual"], _ = strconv.ParseFloat(rows[0]["amount"], 64)
		record["charge_manual_count"], _ = strconv.Atoi(rows[0]["count"])
	}
	// 线上出款
	rows, _ = db.SQL("SELECT COUNT(DISTINCT user_id) count,SUM(amount) amount FROM withdraw_records WHERE status=1 AND updated BETWEEN " + between).QueryString()
	if len(rows) > 0 {
		record["withdraw_online"], _ = strconv.ParseFloat(rows[0]["amount"], 64)
		record["withdraw_online_count"], _ = strconv.Atoi(rows[0]["count"])
	}
	// 手工出款
	rows, _ = db.SQL("SELECT COUNT(DISTINCT user_id) count,SUM(amount) amount FROM manual_withdraws WHERE state=1 AND deal_time BETWEEN " + between).QueryString()
	if len(rows) > 0 {
		record["withdraw_manual"], _ = strconv.ParseFloat(rows[0]["amount"], 64)
		record["withdraw_manual_count"], _ = strconv.Atoi(rows[0]["count"])
	}
	// 实际盈亏
	rows, _ = db.SQL("SELECT SUM(amount) amount FROM account_infos WHERE created BETWEEN " + between).QueryString()
	if len(rows) > 0 {
		record["app_win"], _ = strconv.ParseFloat(rows[0]["amount"], 64)
	}
	r := map[string]string{}
	for k, v := range record {
		if reflect.TypeOf(v).String() == "int" {
			r[k] = strconv.Itoa(v.(int))
		} else if reflect.TypeOf(v).String() == "float64" {
			r[k] = strconv.FormatFloat(v.(float64), 'f', 3, 64)
		} else if v == "" {
			r[k] = "0"
		}
	}
	return r
}

func (self *SystemStatisticsJob) SaveData(db *xorm.Engine, data map[string]string) {
	ymd, _ := strconv.Atoi(data["ymd"])
	if ymd <= 0 {
		return
	}
	saveRecord(db, &data, "system_statistics")
}
