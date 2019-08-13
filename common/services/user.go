package services

import (
	libXorm "github.com/go-xorm/xorm"
	"github.com/kataras/iris/core/errors"
	"github.com/shopspring/decimal"
	"qpgame/app/fund"
	"qpgame/common/utils"
	"qpgame/config"
	"qpgame/models"
	"qpgame/models/xorm"
	"qpgame/ramcache"
	"strconv"
	"time"
)

func UpgradeLevel(platform string, iUserId int) error {
	iBetNum := iUserId % 10
	sBetNum := strconv.Itoa(iBetNum)
	sUserId := strconv.Itoa(iUserId)
	localTime, _ := time.LoadLocation("Asia/Shanghai")
	zeroTime, _ := time.ParseInLocation("2006-01-02 15:04:05", time.Now().Format("2006-01-02")+" 00:00:00", localTime)
	endTime := strconv.Itoa(int(zeroTime.Unix()))
	sql := "SELECT `user_id`,SUM(`amount`) today_bet_amount,SUM(`reward`) today_balance_lucky FROM bets_" + sBetNum + " WHERE user_id=" + sUserId + " AND `ented` >= " + endTime
	engine := models.MyEngine[platform]
	rows, err := engine.SQL(sql).QueryString()
	if len(rows) != 1 {
		return nil
	}
	if err != nil {
		return err
	}
	row := rows[0]
	dTodayBetAmount, _ := decimal.NewFromString(row["today_bet_amount"])
	dTodayBalanceLucky, _ := decimal.NewFromString(row["today_balance_lucky"])
	dZero := decimal.New(0, 0)
	if dTodayBetAmount.Equal(dZero) && dTodayBalanceLucky.Equal(dZero) {
		return nil
	}
	var affNum int64
	session := engine.NewSession()
	session.Begin()
	defer session.Close()
	_, err = session.Where("user_id=?", iUserId).Cols("today_bet_amount", "today_balance_lucky").ForUpdate().Update(xorm.Accounts{
		TodayBetAmount:    dTodayBetAmount.String(),
		TodayBalanceLucky: dTodayBalanceLucky.String(),
	})
	if err != nil {
		session.Rollback()
		return errors.New("更新等级失败，无法更新当天投注信息")
	}
	var oldUser xorm.Users
	var userOk bool
	userOk, err = session.ID(iUserId).ForUpdate().Get(&oldUser)
	if err != nil || userOk == false {
		session.Rollback()
		return errors.New("更新等级失败，无法获取用户信息")
	}
	var account xorm.Accounts
	var accountOk bool
	accountOk, err = engine.Where("user_id=?", iUserId).Cols("total_bet_amount").Get(&account)
	if err != nil || accountOk == false {
		session.Rollback()
		return errors.New("更新等级失败，无法获取账户信息")
	}
	dTotalBetAmount, _ := decimal.NewFromString(account.TotalBetAmount)
	dBetTotal := dTotalBetAmount.Add(dTodayBetAmount)
	fBetTotal, _ := dBetTotal.Float64()
	vipLevels, _ := ramcache.TableVipLevels.Load(platform)
	vips := vipLevels.([]xorm.VipLevels)
	var iLev int
	for _, vlRow := range vips {
		fValidBetMin := float64(vlRow.ValidBetMin) * 10000
		fValidBetMax := float64(vlRow.ValidBetMax) * 10000
		if fBetTotal >= fValidBetMin && fBetTotal < fValidBetMax {
			iLev = vlRow.Level
			break
		}
	}
	if oldUser.VipLevel == iLev {
		return nil
	}
	affNum, err = session.ID(iUserId).Cols("vip_level").ForUpdate().Update(xorm.Users{
		VipLevel: iLev,
	})
	if err != nil || affNum == 0 {
		session.Rollback()
		return errors.New("更新等级失败，无法更新用户等级")
	}
	for j := oldUser.VipLevel + 1; j <= iLev; j++ {
		//获取对应vip等级的晋级礼金金额
		amount := 0
		for _, v := range vips {
			if v.Level == j {
				amount = v.UpgradeAmount
				break
			}
		}
		sLev := strconv.Itoa(j)
		info := map[string]interface{}{
			"user_id":     iUserId,
			"transaction": session,
			"type_id":     config.FUNDVIPUPLEVEL,
			"amount":      float64(amount),
			"order_id":    utils.CreationOrder("VIP", sUserId),
			"msg":         "VIP" + sLev + "晋级礼金",
		}
		balance := fund.NewUserFundChange(platform)
		balanceUpdateRes := balance.BalanceUpdate(info, nil)
		if balanceUpdateRes["status"] != 1 {
			session.Rollback()
			return errors.New("更新等级失败，无法更新用户资金")
		}
		title := "恭喜升级到VIP " + sLev
		content := "赠送晋级礼金" + strconv.Itoa(amount) + "，已到账，请查收！"
		noticeEntity := xorm.Notices{
			UserId:  iUserId,
			Title:   title,
			Content: content,
			Status:  1,
			Created: utils.GetNowTime(),
		}
		affNum, err = session.Insert(noticeEntity)
		if err != nil || affNum <= 0 {
			session.Rollback()
			return errors.New("更新等级失败，无法发送站内信")
		}
	}
	session.Commit()
	return nil
}

func PromotionAward(platform string, session *libXorm.Session, iUserId int) error {
	if iUserId > 0 {
		cnf, _ := ramcache.TableConfigs.Load(platform)
		cnfMap := cnf.(map[string]interface{})
		if generalizeAward, ok := cnfMap["generalize_award"]; ok {
			fGeneralizeAward := generalizeAward.(float64)
			sUserId := strconv.Itoa(iUserId)
			info := map[string]interface{}{
				"user_id":     iUserId,
				"transaction": session,
				"type_id":     config.FUNDGENERALIZEAWAED,
				"amount":      fGeneralizeAward,
				"order_id":    utils.CreationOrder("TG", sUserId),
				"msg":         "推广邀请奖励",
			}
			balanceUpdateRes := fund.NewUserFundChange(platform).BalanceUpdate(info, nil)
			if balanceUpdateRes["status"] != 1 {
				session.Rollback()
				return errors.New("推广邀请奖励赠送失败")
			}
			title := "推广奖励"
			sGeneralizeAward := strconv.FormatFloat(fGeneralizeAward, 'f', 2, 64)
			content := "恭喜获得推广奖励" + sGeneralizeAward + "，已到账，请查收！"
			noticeEntity := xorm.Notices{
				UserId:  iUserId,
				Title:   title,
				Content: content,
				Status:  1,
				Created: utils.GetNowTime(),
			}
			affNum, err := session.Insert(noticeEntity)
			if err != nil || affNum <= 0 {
				session.Rollback()
				return errors.New("推广奖励赠送失败，无法发送站内信")
			}
		}
	}
	return nil
}

func BindPhoneAward(platform string, session *libXorm.Session, iUserId int) error {
	if iUserId > 0 {
		cnf, _ := ramcache.TableConfigs.Load(platform)
		cnfMap := cnf.(map[string]interface{})
		bindPhoneAwardSwitch, bpaSwitchOk := cnfMap["bind_phone_award_switch"]
		rewardBind, rbOk := cnfMap["reward_bind"]
		if bpaSwitchOk && rbOk {
			fBindPhoneAwardSwitch := bindPhoneAwardSwitch.(float64)
			bBindPhoneAwardSwitch := int(fBindPhoneAwardSwitch) == 1
			fRewardBind := rewardBind.(float64)
			if bBindPhoneAwardSwitch {
				title := "绑定手机奖励"
				sUserId := strconv.Itoa(iUserId)
				info := map[string]interface{}{
					"user_id":     iUserId,
					"transaction": session,
					"type_id":     config.FUNDBINDPHONE,
					"amount":      fRewardBind,
					"order_id":    utils.CreationOrder("BD", sUserId),
					"msg":         title,
				}
				balanceUpdateRes := fund.NewUserFundChange(platform).BalanceUpdate(info, nil)
				if balanceUpdateRes["status"] != 1 {
					session.Rollback()
					return errors.New("绑定手机奖励赠送失败")
				}
				sAward := strconv.FormatFloat(fRewardBind, 'f', 2, 64)
				content := "恭喜获得绑定手机奖励" + sAward + "，已到账，请查收！"
				noticeEntity := xorm.Notices{
					UserId:  iUserId,
					Title:   title,
					Content: content,
					Status:  1,
					Created: utils.GetNowTime(),
				}
				affNum, err := session.Insert(noticeEntity)
				if err != nil || affNum <= 0 {
					session.Rollback()
					return errors.New("绑定手机奖励赠送失败，无法发送站内信")
				}
			}
		}
	}
	return nil
}

func ActivityAward(platform string, session *libXorm.Session, iActType, iUserId int, sIp string) (string, error) {
	if iUserId > 0 {
		actTypeList := map[int]string{
			1: "注册",
			2: "充值",
		}
		actType, isExist := actTypeList[iActType]
		if isExist {
			var activity = new(xorm.Activities)
			iNow := utils.GetNowTime()
			actExist, _ := session.Where("`status`=1 AND `type`=? AND time_start<=? AND time_end>?", iActType, iNow, iNow).Get(activity)
			rewardName := actType + "奖励"
			var err error
			if actExist {
				decimal.DivisionPrecision = 2
				sMoney := activity.Money
				var dMoney decimal.Decimal
				dMoney, err = decimal.NewFromString(sMoney)
				fMoney, toFloat := dMoney.Float64()
				if err != nil || toFloat == false {
					return "", nil
				}
				if dMoney.GreaterThan(decimal.New(0, 0)) == false {
					return "", nil
				}
				var actRecords []xorm.ActivityRecords
				err = session.Cols("id", "created", "ip_addr").Where("state=1 AND user_id=? AND activity_id=?", iUserId, activity.Id).Find(&actRecords)
				if err != nil {
					return "活动记录获取失败", errors.New(rewardName + "赠送失败。")
				}
				iActRecordNum := len(actRecords)
				var actRecord = xorm.ActivityRecords{
					UserId:  iUserId,
					State:   1,
					Applied: iNow,
					Created: iNow,
					Updated: iNow,
					IpAddr:  sIp,
				}
				if iActRecordNum == 0 {
					return RecordActivity(platform, session, actRecord, fMoney, rewardName)
				}
				iIsRepeat := activity.IsRepeat
				bIsRepeat := iIsRepeat == 1
				// 是否允许重复，如果不允许重复，查看record的记录条数，如没有，则记录，否则，返回
				if bIsRepeat == false {
					return "", errors.New("您已参与过该活动，无法重复参与。")
				}

				// 计算当日的起始时间戳
				iFromTime, iToTime := utils.GetDatetimeRange(0, 1)
				iIpTotalCnt := 0 // 通过Ip统计参与指定活动的总数
				iIpDayCnt := 0

				for _, actRecordBean := range actRecords {
					iIpTotalCnt++
					iCreated := int64(actRecordBean.Created)
					if iFromTime < iCreated && iCreated <= iToTime {
						iIpDayCnt++
					}
				}
				iTotalIpLimit := activity.TotalIpLimit
				iDayIpLimit := activity.DayIpLimit
				// 判断当前Ip领取某一活动的总次数是否超出限制
				if iIpTotalCnt >= iTotalIpLimit {
					return "1906281053", errors.New("领取该活动的次数已达上限，无法继续领取")
				}

				// 判断当前Ip当日领取某一活动是否超出限制
				if iIpDayCnt >= iDayIpLimit {
					return "1906281056", errors.New("今日领取该活动的次数已达上限，请明天再来")
				}
				return RecordActivity(platform, session, actRecord, fMoney, rewardName)
			}
		}
	}
	return "", nil
}
