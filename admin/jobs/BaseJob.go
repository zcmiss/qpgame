package jobs

import (
	"qpgame/common/utils"
	"qpgame/config"
	"qpgame/models"
	"strconv"
	"strings"
	"time"

	"github.com/go-xorm/xorm"

	"github.com/robfig/cron"
)

// 加载所有的后台任务
func InitAdminJobs() {
	dispatcher := cron.New() //建立任务调度器
	initJobs(dispatcher)
	dispatcher.Start()
	select {}
}

// 加载预定义的所有任务
func initJobs(c *cron.Cron) {
	initJob(c, websocket)         //Websocket
	initJob(c, accountStatistics) //用户统计
	initJob(c, proxyStatistics)   //代理统计
	initJob(c, systemStatistics)  //运营统计
	initJob(c, redpackets)        //红包分发
}

// 加载单个的任务
func initJob(c *cron.Cron, t AdminCronJob) {
	(*c).AddFunc(t.GetSpec(), func() {
		for platform := range config.PlatformCPs {
			t.Process(models.MyEngine[platform], platform)
		}
	})
}

// 依据分页得到相应的用户编号数组
func getUserIds(db *xorm.Engine, page int, pageSize int, timeLimit int) []string {
	uIds := []string{}
	offset := strconv.Itoa(pageSize * (page - 1))
	sql := "SELECT id FROM users WHERE created <" + strconv.Itoa(timeLimit) + " LIMIT " + offset + ", " + strconv.Itoa(pageSize)
	rows, err := db.SQL(sql).QueryString()
	if err != nil || len(rows) == 0 {
		return nil
	}
	for _, row := range rows {
		if row["id"] != "" {
			uIds = append(uIds, row["id"])
		}
	}
	if len(uIds) == 0 {
		return nil
	}
	return uIds
}

// 依据分页得到相应的代理用户编号数组
func getProxyUserIds(db *xorm.Engine) map[string]string {
	userIds := map[string]string{}
	sql := "SELECT DISTINCT parent_id FROM users WHERE parent_id>0 GROUP BY parent_id ORDER BY parent_id ASC"
	rows, _ := db.SQL(sql).QueryString()
	if len(rows) == 0 {
		return nil
	}
	for _, row := range rows {
		parentId := row["parent_id"]
		ids := []string{}
		items, _ := db.SQL("SELECT id FROM users WHERE id=" + parentId).QueryString()
		if len(items) > 0 {
			ids = append(ids, parentId)
			items, _ := db.SQL("SELECT id FROM users WHERE parent_id=" + parentId).QueryString()
			if len(items) > 0 {
				for _, v := range items {
					ids = append(ids, v["id"])
				}
			}
			userIds[parentId] = strings.Join(ids, ",")
		}
	}
	return userIds
}

// 将用户id数组拆分成10个数组
func getBetSqls(ids []string, callback func(index int, userIds string) string) []string {
	var uIds = []string{"", "", "", "", "", "", "", "", "", ""}
	for _, idStr := range ids {
		id, _ := strconv.Atoi(idStr)
		index := id % 10
		uIds[index] = uIds[index] + "," + idStr
	}
	sqlArr := []string{}
	for k, v := range uIds {
		if len(v) > 1 {
			sqlArr = append(sqlArr, callback(k, v[1:]))
		}
	}
	return sqlArr
}

// 得到需要统计的时间
func getCountTimes(day *time.Time) (int, int, int) {
	ymdStr := (*day).Format("20060102")
	ymd, _ := strconv.Atoi(ymdStr)
	countTime := day.Format("2006-01-02 00:00:00")      //当天的开始时间
	timeStart := int(utils.GetInt64FromTime(countTime)) //搜索开始时间
	timeEnd := int(timeStart + 86399)                   //结束时间
	return ymd, timeStart, timeEnd
}

// 是否已统计过了
func hasCounted(db *xorm.Engine, tableName string, ymd int, userId int) bool {
	sql := "SELECT id FROM " + tableName + " WHERE ymd='" + strconv.Itoa(ymd) + "'"
	if tableName == "proxy_statistics" && userId > 0 {
		sql += " AND user_id=" + strconv.Itoa(userId)
	}
	rows, err := db.SQL(sql).QueryString()
	return err != nil || len(rows) >= 1
}

// 创建/更新记录
func saveRecord(db *xorm.Engine, data *map[string]string, tableName string) int64 {
	fields, values, updateStr := []string{}, []string{}, []string{}
	for k, v := range *data {
		fields = append(fields, k)
		values = append(values, "'"+v+"'")
		if (k != "ymd") && (k != "user_id") {
			updateStr = append(updateStr, k+"='"+v+"'")
		}
	}
	sql := "INSERT INTO " + tableName + " (" + strings.Join(fields, ",") + ") VALUES (" + strings.Join(values, ",") + ") ON DUPLICATE KEY UPDATE " + strings.Join(updateStr, ",")
	result, err := db.Exec(sql)
	if err != nil {
		return 0
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0
	}
	return id
}
