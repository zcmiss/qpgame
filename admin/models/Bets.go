package models

import (
	"qpgame/admin/common"
	"qpgame/models"
	"strconv"
	"strings"
)

// 模型
type Bets struct{}

// 查询字段
var fieldsOfBets = "amount, amount_all, amount_platform, created, id, order_id, " +
	"platform_id, reward, user_id, rebate_state, dama_state, gt, game_code, accountname AS account_name"

// 查询条件
var queryCondOfBets = func(ctx *Context) []string { //获取查询条件
	queries := getQueryFields(ctx, &map[string]string{
		"game_code":    "=", //游戏名称
		"platform_id":  "=", //平台
		"order_id":     "%", //订单编号
		"rebate_state": "=", //洗码状态
		"dama_state":   "=", //打码状态
		"accountname":  "%", //第三方游戏平台账号
	})
	created := getQueryFieldByTime(ctx, "created", "time_start", "time_end")
	if created!=""{
		queries = append(queries, created)
	}

	idQuery := getQueryOfUserId(ctx)
	if idQuery != "" {
		queries = append(queries, idQuery)
	}
	return queries
}

// 表名称
func (self *Bets) GetTableName(ctx *Context) string {
	var sqlArr []string
	where := " WHERE 1 = 1 "
	cond := strings.Join(queryCondOfBets(ctx), " AND ")
	if cond != "" {
		where += " AND " + cond
	}
	for i := 0; i <= 9; i++ {
		sqlArr = append(sqlArr, "SELECT "+fieldsOfBets+" FROM bets_"+strconv.Itoa(i)+" "+where)
	}
	return "(" + strings.Join(sqlArr, " UNION ALL ") + ") AS t"
}

// 得到所有记录-分页
// TODO: 当数据量达到一定规模必须优化获取查询的方式
func (self *Bets) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self, "*", queryCondOfBets,
		func(ctx *Context, row *map[string]string) {
			platform := (*ctx).Params().Get("platform")
			(*row)["platform_name"] = common.GetGamePlatformName(platform, (*row)["platform_id"]) //第三方平台名称
			(*row)["game_name"] = common.GetGameNameOfCode(platform, (*row)["game_code"])         //游戏名称
			(*row)["user_name"] = common.GetUserName(platform, (*row)["user_id"])                 //用户名称
			(*row)["category_name"] = common.GetGameCategoryName(platform, (*row)["gt"])          //游戏分类名称
			processOptionsFor("rebate_state", "rebate_state_name", &map[string]string{
				"0": "未洗",
				"1": "已洗",
			}, row)
			processOptionsFor("dama_state", "dama_state_name", &map[string]string{
				"0": "未更新",
				"1": "已更新",
			}, row)
		}, nil)
}

// 统计下注总额
func (self *Bets) GetCalculatedTotal(ctx *Context) (map[string]string, error)  {
	data := map[string]string{
		"total_bet":"0",
		"total_prize":"0",
		"platform_loss":"0",
	}
	conn := models.MyEngine[(*ctx).Params().Get("platform")]
	var sqlArr []string
	where := " WHERE 1 = 1 "
	cond := strings.Join(queryCondOfBets(ctx), " AND ")
	if cond != "" {
		where += " AND " + cond
	}
	for i := 0; i <= 9; i++ {
		sqlArr = append(sqlArr, "SELECT "+fieldsOfBets+" FROM bets_"+strconv.Itoa(i)+" "+where)
	}
	sql :="(" + strings.Join(sqlArr, " UNION ALL ") + ") AS t"
	sqlAll :="SELECT SUM(amount) AS total_bet,SUM(reward) AS total_prize,(SUM(reward)-SUM(amount)) AS platform_loss FROM  "+sql
	rows, err := conn.SQL(sqlAll).QueryString()
	if err != nil || len(rows) == 0 {
		return data, nil
	}

	row := rows[0]
	for k, _ := range data {
		if row[k] != "" {
			data[k] = row[k]
		}
	}
	return data, nil
}
// 得到记录详情
func (self *Bets) GetRecordDetail(ctx *Context) (map[string]string, error) {
	return getRecordDetail(ctx, self, "", nil)
}

// 添加记录
func (self *Bets) Save(ctx *Context) (int64, error) {
	return denySave()
}

// 删除记录
func (self *Bets) Delete(ctx *Context) error {
	return denyDelete()
}

