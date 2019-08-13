package services

import (
	"errors"
	libXorm "github.com/go-xorm/xorm"
	"qpgame/app/fund"
	"qpgame/common/utils"
	"qpgame/config"
	"qpgame/models/xorm"
	"strconv"
)

func RecordActivity(platform string, session *libXorm.Session, activityRecord xorm.ActivityRecords, fAwardMoney float64, rewardName string) (string, error) {
	var affNum int64 = 0
	var err error
	affNum, err = session.Insert(activityRecord)
	if err != nil || affNum == 0 {
		return "活动记录保存失败", errors.New(rewardName + "赠送失败。")
	}
	iUserId := activityRecord.UserId
	sUserId := strconv.Itoa(iUserId)
	info := map[string]interface{}{
		"user_id":     iUserId,
		"transaction": session,
		"type_id":     config.FUNDACTIVITYAWARD,
		"amount":      fAwardMoney,
		"order_id":    utils.CreationOrder("HD", sUserId),
		"msg":         rewardName,
	}
	balanceUpdateRes := fund.NewUserFundChange(platform).BalanceUpdate(info, nil)
	if balanceUpdateRes["status"] != 1 {
		return "资金流水保存失败", errors.New(rewardName + "赠送失败")
	}
	sAward := strconv.FormatFloat(fAwardMoney, 'f', 2, 64)
	content := "恭喜获得" + rewardName + sAward + "，已到账，请查收！"
	noticeEntity := xorm.Notices{
		UserId:  iUserId,
		Title:   rewardName,
		Content: content,
		Status:  1,
		Created: utils.GetNowTime(),
	}
	affNum, err = session.Insert(noticeEntity)
	if err != nil || affNum <= 0 {
		return "无法发送站内信", errors.New(rewardName + "赠送失败。")
	}
	return "", nil
}
