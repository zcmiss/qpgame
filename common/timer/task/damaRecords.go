package task

import (
	"fmt"
	"github.com/shopspring/decimal"
	"qpgame/common/utils"
	"qpgame/models"
	"qpgame/models/xorm"
	"strconv"
)

//自动更新用户打码量
func UpdateWithdrawDamaRecords(platform string) {
	jdbc := models.MyEngine[platform]
	//定时器每5分钟执行一次，查询所有未完成打码量的投注记录信息集合，时间必须大于打码量的产生时间
	sqlStr := "SELECT b.id,b.`user_id`,b.`amount` FROM bets_0 b WHERE b.`dama_state` = 0 AND EXISTS(SELECT 1 FROM withdraw_dama_records w WHERE w.`created`<=b.`ented` AND w.`state`=0 AND w.`user_id` = b.`user_id`) UNION ALL SELECT b.id,b.`user_id`,b.`amount` FROM bets_1 b WHERE b.`dama_state` = 0 AND EXISTS(SELECT 1 FROM withdraw_dama_records w WHERE w.`created`<=b.`ented` AND w.`state`=0 AND w.`user_id` = b.`user_id`) UNION ALL SELECT b.id,b.`user_id`,b.`amount` FROM bets_2 b WHERE b.`dama_state` = 0 AND EXISTS(SELECT 1 FROM withdraw_dama_records w WHERE w.`created`<=b.`ented` AND w.`state`=0 AND w.`user_id` = b.`user_id`) UNION ALL SELECT b.id,b.`user_id`,b.`amount` FROM bets_3 b WHERE b.`dama_state` = 0 AND EXISTS(SELECT 1 FROM withdraw_dama_records w WHERE w.`created`<=b.`ented` AND w.`state`=0 AND w.`user_id` = b.`user_id`) UNION ALL SELECT b.id,b.`user_id`,b.`amount` FROM bets_4 b WHERE b.`dama_state` = 0 AND EXISTS(SELECT 1 FROM withdraw_dama_records w WHERE w.`created`<=b.`ented` AND w.`state`=0 AND w.`user_id` = b.`user_id`) UNION ALL SELECT b.id,b.`user_id`,b.`amount` FROM bets_5 b WHERE b.`dama_state` = 0 AND EXISTS(SELECT 1 FROM withdraw_dama_records w WHERE w.`created`<=b.`ented` AND w.`state`=0 AND w.`user_id` = b.`user_id`) UNION ALL SELECT b.id,b.`user_id`,b.`amount` FROM bets_6 b WHERE b.`dama_state` = 0 AND EXISTS(SELECT 1 FROM withdraw_dama_records w WHERE w.`created`<=b.`ented` AND w.`state`=0 AND w.`user_id` = b.`user_id`) UNION ALL SELECT b.id,b.`user_id`,b.`amount` FROM bets_7 b WHERE b.`dama_state` = 0 AND EXISTS(SELECT 1 FROM withdraw_dama_records w WHERE w.`created`<=b.`ented` AND w.`state`=0 AND w.`user_id` = b.`user_id`) UNION ALL SELECT b.id,b.`user_id`,b.`amount` FROM bets_8 b WHERE b.`dama_state` = 0 AND EXISTS(SELECT 1 FROM withdraw_dama_records w WHERE w.`created`<=b.`ented` AND w.`state`=0 AND w.`user_id` = b.`user_id`) UNION ALL SELECT b.id,b.`user_id`,b.`amount` FROM bets_9 b WHERE b.`dama_state` = 0 AND EXISTS(SELECT 1 FROM withdraw_dama_records w WHERE w.`created`<=b.`ented` AND w.`state`=0 AND w.`user_id` = b.`user_id`)"
	updateSqls := ""
	for i := 0; i < 10; i++ {
		updateSql := "UPDATE bets_" + strconv.Itoa(i) + " b SET dama_state=1  WHERE b.`dama_state` = 0 AND EXISTS(SELECT 1 FROM withdraw_dama_records w WHERE w.`created`<=b.`ented` AND w.`state`=0 AND w.`user_id` = b.`user_id`);"
		updateSqls += updateSql
	}
	resMap, err := utils.Query(platform, sqlStr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if len(resMap) == 0 {
		return
	}
	userIds := make([]string, 0)
	userMap := make(map[string]decimal.Decimal)
	//将每个用户的打码量计算总和
	for _, res := range resMap {
		userId := strconv.Itoa(int(res["user_id"].(float64)))
		_, ok := userMap[userId]
		if !ok {
			userMap[userId] = decimal.Zero
		}
		userMap[userId] = userMap[userId].Add(decimal.NewFromFloat(res["amount"].(float64)))
	}
	//统计有多少在此次当中产生投注的用户
	for k, _ := range userMap {
		userIds = append(userIds, k)
	}
	//根据查询到的所有投注记录，获取用户打码量流水详情并计算
	damaIds := make([]int, 0)
	var damas []xorm.WithdrawDamaRecords
	err = jdbc.Where("state = 0").In("user_id", userIds).Find(&damas)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	updateDamaSqls := ""
	for i, dama := range damas {
		damaIds = append(damaIds, dama.Id)
		FinishedProgress, _ := decimal.NewFromString(dama.FinishedProgress)
		FinishedNeeded, _ := decimal.NewFromString(dama.FinishedNeeded)
		//当前剩余打码量
		need := FinishedNeeded.Sub(FinishedProgress)
		//该用户本次投注的总打码量
		totalAmount := userMap[strconv.Itoa(dama.UserId)]
		//如果投注量小于需要的打码量
		if totalAmount.LessThanOrEqual(need) {
			FinishedProgress = FinishedProgress.Add(totalAmount)
			userMap[strconv.Itoa(dama.UserId)] = decimal.Zero
		}
		//如果投注量大于需要的打码量
		if totalAmount.GreaterThan(need) {
			totalAmount = totalAmount.Sub(need)
			FinishedProgress = FinishedProgress.Add(need)
			userMap[strconv.Itoa(dama.UserId)] = totalAmount
			damas[i].State = 1
		}

		if totalAmount.Equal(need) {
			FinishedProgress = FinishedProgress.Add(need)
			userMap[strconv.Itoa(dama.UserId)] = decimal.Zero
			damas[i].State = 1
		}
		damas[i].FinishedProgress = FinishedProgress.String()

		updateDamaSqls += "update withdraw_dama_records w set w.finished_progress = " + damas[i].FinishedProgress + ",updated=" + strconv.Itoa(utils.GetNowTime()) + ",state=" + strconv.Itoa(damas[i].State) + " where id=" + strconv.Itoa(damas[i].Id) + ";"
	}

	//更新用户打码量，并将所有投注记录的打码状态进行更新
	session := jdbc.NewSession()
	err = session.Begin()
	defer session.Close()
	_, err = session.Exec(updateSqls)
	if err != nil {
		session.Rollback()
		fmt.Println(err.Error())
		return
	}
	//更新打码量流水详情表
	_, err = session.Exec(updateDamaSqls)
	if err != nil {
		session.Rollback()
		fmt.Println(err.Error())
		return
	}
	session.Commit()
}
