package task

import (
	"fmt"
	"github.com/shopspring/decimal"
	"qpgame/app/fund"
	"qpgame/common/utils"
	"qpgame/config"
	"qpgame/models"
	"qpgame/models/xorm"
	"qpgame/ramcache"
	"strconv"
	"strings"
	"time"
)

func UpdateAccountTodayBetAmount(platform string) {
	jdbc := models.MyEngine[platform]
	localTime, _ := time.LoadLocation("Asia/Shanghai")

	zeroTime, _ := time.ParseInLocation("2006-01-02 15:04:05", time.Now().Format("2006-01-02")+" 00:00:00", localTime)
	endTime := strconv.Itoa(int(zeroTime.Unix()))
	//定时器每5分钟执行一次，查询今天所有的打码量
	sqlStr := "SELECT b.`user_id`,SUM(b.`amount`) today_bet_amount,SUM(b.`reward`) today_balance_lucky  FROM bets_0 b WHERE b.`ented`>=" + endTime + " GROUP BY b.`user_id` UNION ALL SELECT b.`user_id`,SUM(b.`amount`) today_bet_amount,SUM(b.`reward`) today_balance_lucky  FROM bets_1 b WHERE b.`ented`>=" + endTime + " GROUP BY b.`user_id` UNION ALL SELECT b.`user_id`,SUM(b.`amount`) today_bet_amount,SUM(b.`reward`) today_balance_lucky  FROM bets_2 b WHERE b.`ented`>=" + endTime + " GROUP BY b.`user_id` UNION ALL SELECT b.`user_id`,SUM(b.`amount`) today_bet_amount,SUM(b.`reward`) today_balance_lucky  FROM bets_3 b WHERE b.`ented`>=" + endTime + " GROUP BY b.`user_id` UNION ALL SELECT b.`user_id`,SUM(b.`amount`) today_bet_amount,SUM(b.`reward`) today_balance_lucky  FROM bets_4 b WHERE b.`ented`>=" + endTime + " GROUP BY b.`user_id` UNION ALL SELECT b.`user_id`,SUM(b.`amount`) today_bet_amount,SUM(b.`reward`) today_balance_lucky  FROM bets_5 b WHERE b.`ented`>=" + endTime + " GROUP BY b.`user_id` UNION ALL SELECT b.`user_id`,SUM(b.`amount`) today_bet_amount,SUM(b.`reward`) today_balance_lucky  FROM bets_6 b WHERE b.`ented`>=" + endTime + " GROUP BY b.`user_id` UNION ALL SELECT b.`user_id`,SUM(b.`amount`) today_bet_amount,SUM(b.`reward`) today_balance_lucky  FROM bets_7 b WHERE b.`ented`>=" + endTime + " GROUP BY b.`user_id` UNION ALL SELECT b.`user_id`,SUM(b.`amount`) today_bet_amount,SUM(b.`reward`) today_balance_lucky  FROM bets_8 b WHERE b.`ented`>=" + endTime + " GROUP BY b.`user_id` UNION ALL SELECT b.`user_id`,SUM(b.`amount`) today_bet_amount,SUM(b.`reward`) today_balance_lucky  FROM bets_9 b WHERE b.`ented`>=" + endTime + " GROUP BY b.`user_id` "
	updateSqls := ""
	resMap, err := utils.Query(platform, sqlStr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if len(resMap) == 0 {
		return
	}
	userIds := make([]string, 0)
	//根据查询到的所有投注记录，更新用户当日的投注总额
	for _, m := range resMap {
		todatBetAmount := decimal.NewFromFloat(m["today_bet_amount"].(float64))
		todayBalanceLucky := decimal.NewFromFloat(m["today_balance_lucky"].(float64))
		userId := strconv.Itoa(int(m["user_id"].(float64)))
		userIds = append(userIds, userId)
		updateSqls += "update accounts set today_bet_amount = " + todatBetAmount.String() + ",today_balance_lucky = " + todayBalanceLucky.String() + " where user_id = " + userId + ";"
	}
	//获取用户原来的vip等级
	var oldUsers []xorm.Users
	session := jdbc.NewSession()
	defer session.Close()
	session.Begin()
	session.In("id", userIds).ForUpdate().Find(&oldUsers)
	//更新用户打码量，并将所有投注记录的打码状态进行更新
	_, err = session.Exec(updateSqls)
	if err != nil {
		session.Rollback()
		fmt.Println(err.Error())
		return
	}
	// 从accounts表里获取需要更新用户的总打码量
	userIdsStr := strings.Join(userIds, ",")
	qrySql := "SELECT user_id,(today_bet_amount+total_bet_amount) bet_total FROM accounts WHERE user_id IN (" + userIdsStr + ")"
	qryRows, qryErr := jdbc.SQL(qrySql).QueryString()
	if qryErr != nil {
		session.Rollback()
		fmt.Println(qryErr.Error())
		return
	}
	vipLevels, _ := ramcache.TableVipLevels.Load(platform)
	vips := vipLevels.([]xorm.VipLevels)
	var upgradeUsers = make(map[int]int, 0)
	for _, qryRow := range qryRows {
		sBetTotal := qryRow["bet_total"]
		fBetTotal, _ := strconv.ParseFloat(sBetTotal, 64)
		for _, vlRow := range vips {
			fValidBetMin := float64(vlRow.ValidBetMin) * 10000
			fValidBetMax := float64(vlRow.ValidBetMax) * 10000
			if fBetTotal >= fValidBetMin && fBetTotal < fValidBetMax {
				sUserId := qryRow["user_id"]
				iUserId, _ := strconv.Atoi(sUserId)
				upgradeUsers[iUserId] = vlRow.Level
				break
			}
		}
	}
	var oldUsersIdx = make(map[int]xorm.Users, 0)
	for _, oldUser := range oldUsers {
		oldUsersIdx[oldUser.Id] = oldUser
	}
	//如果有会员升级VIP，触发VIP晋级奖金派发
	if len(upgradeUsers) > 0 {
		var affNum int64
		for iUserId, iLevel := range upgradeUsers {
			oldUserItem, oldUserExist := oldUsersIdx[iUserId]
			if oldUserExist {
				if oldUserItem.VipLevel == iLevel {
					continue
				}
				affNum, err = session.ID(iUserId).Cols("vip_level").ForUpdate().Update(xorm.Users{
					VipLevel: iLevel,
				})
				if err != nil || affNum == 0 {
					session.Rollback()
					return
				}
				for j := oldUserItem.VipLevel + 1; j <= iLevel; j++ {
					//获取对应vip等级的晋级礼金金额
					amount := 0
					for _, v := range vips {
						if v.Level == j {
							amount = v.UpgradeAmount
							break
						}
					}
					sLevel := strconv.Itoa(j)
					info := map[string]interface{}{
						"user_id":     oldUserItem.Id,
						"transaction": session,
						"type_id":     config.FUNDVIPUPLEVEL,
						"amount":      float64(amount),
						"order_id":    utils.CreationOrder("VIP", strconv.Itoa(oldUserItem.Id)),
						"msg":         "VIP" + sLevel + "晋级礼金",
						//"finish_rate": 1.0, //需满足的打码量比例
					}
					balance := fund.NewUserFundChange(platform)
					balanceUpdateRes := balance.BalanceUpdate(info, nil)
					if balanceUpdateRes["status"] != 1 {
						session.Rollback()
						return
					}
					title := "恭喜升级到VIP " + sLevel
					content := "赠送晋级礼金" + strconv.Itoa(amount) + "，已到账，请查收！"
					noticeEntity := xorm.Notices{
						UserId:  iUserId,
						Title:   title,
						Content: content,
						Status:  1,
						Created: utils.GetNowTime(),
					}
					affNum, inErr := session.Insert(noticeEntity)
					if inErr != nil || affNum <= 0 {
						session.Rollback()
						return
					}
				}
			}
		}
	}
	session.Commit()
}

func UpdateAccountTotalBetAmount(platform string) {
	jdbc := models.MyEngine[platform]
	localTime, _ := time.LoadLocation("Asia/Shanghai")
	zeroTime, _ := time.ParseInLocation("2006-01-02 15:04:05", time.Now().Format("2006-01-02")+" 00:00:00", localTime)
	endTime := strconv.Itoa(int(zeroTime.Unix()))
	startTime := strconv.Itoa(int(zeroTime.AddDate(0, 0, -1).Unix()))
	//每天凌晨执行一次，查询昨天所有的打码量
	sqlStr := "SELECT b.`user_id`,SUM(b.`amount`) today_bet_amount,SUM(b.`reward`) today_balance_lucky FROM bets_0 b WHERE b.`ented`<" + endTime + " and b.`ented`>=" + startTime + "  GROUP BY b.`user_id` UNION ALL SELECT b.`user_id`,SUM(b.`amount`) today_bet_amount,SUM(b.`reward`) today_balance_lucky FROM bets_1 b WHERE b.`ented`<" + endTime + " and b.`ented`>=" + startTime + "  GROUP BY b.`user_id` UNION ALL SELECT b.`user_id`,SUM(b.`amount`) today_bet_amount,SUM(b.`reward`) today_balance_lucky FROM bets_2 b WHERE b.`ented`<" + endTime + " and b.`ented`>=" + startTime + "  GROUP BY b.`user_id` UNION ALL SELECT b.`user_id`,SUM(b.`amount`) today_bet_amount,SUM(b.`reward`) today_balance_lucky FROM bets_3 b WHERE b.`ented`<" + endTime + " and b.`ented`>=" + startTime + "  GROUP BY b.`user_id` UNION ALL SELECT b.`user_id`,SUM(b.`amount`) today_bet_amount,SUM(b.`reward`) today_balance_lucky FROM bets_4 b WHERE b.`ented`<" + endTime + " and b.`ented`>=" + startTime + "  GROUP BY b.`user_id` UNION ALL SELECT b.`user_id`,SUM(b.`amount`) today_bet_amount,SUM(b.`reward`) today_balance_lucky FROM bets_5 b WHERE b.`ented`<" + endTime + " and b.`ented`>=" + startTime + "  GROUP BY b.`user_id` UNION ALL SELECT b.`user_id`,SUM(b.`amount`) today_bet_amount,SUM(b.`reward`) today_balance_lucky FROM bets_6 b WHERE b.`ented`<" + endTime + " and b.`ented`>=" + startTime + "  GROUP BY b.`user_id` UNION ALL SELECT b.`user_id`,SUM(b.`amount`) today_bet_amount,SUM(b.`reward`) today_balance_lucky FROM bets_7 b WHERE b.`ented`<" + endTime + " and b.`ented`>=" + startTime + "  GROUP BY b.`user_id` UNION ALL SELECT b.`user_id`,SUM(b.`amount`) today_bet_amount,SUM(b.`reward`) today_balance_lucky FROM bets_8 b WHERE b.`ented`<" + endTime + " and b.`ented`>=" + startTime + "  GROUP BY b.`user_id` UNION ALL SELECT b.`user_id`,SUM(b.`amount`) today_bet_amount,SUM(b.`reward`) today_balance_lucky FROM bets_9 b WHERE b.`ented`<" + endTime + " and b.`ented`>=" + startTime + "  GROUP BY b.`user_id`"
	updateSqls := ""
	resMap, err := utils.Query(platform, sqlStr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if len(resMap) == 0 {
		return
	}
	//根据查询到的所有投注记录，更新用户当日的投注总额
	for _, m := range resMap {
		todatBetAmount := decimal.NewFromFloat(m["today_bet_amount"].(float64))
		todayBalanceLucky := decimal.NewFromFloat(m["today_balance_lucky"].(float64))
		userId := strconv.Itoa(int(m["user_id"].(float64)))
		updateSqls += "update accounts set today_bet_amount=0,today_balance_lucky=0, total_bet_amount = total_bet_amount +" + todatBetAmount.String() + ",balance_lucky = balance_lucky +" + todayBalanceLucky.String() + " where user_id = " + userId + ";"
	}
	//更新用户打码量，并将所有投注记录的打码状态进行更新
	_, err = jdbc.Exec(updateSqls)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
