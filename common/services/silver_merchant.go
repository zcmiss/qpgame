package services

import (
	"errors"
	libXorm "github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	"qpgame/common/utils"
	"qpgame/models/xorm"
)

func SaveOperationLog(ctx iris.Context, session *libXorm.Session, smUserId int, content string) (xorm.SilverMerchantOsLogs, error) {
	var err error

	sIp := utils.GetIp(ctx.Request())
	smOpLog := xorm.SilverMerchantOsLogs{
		MerchantId: smUserId,
		Content:    content,
		Created:    utils.GetNowTime(),
		Ip:         sIp,
		City:       utils.GetIpInfo(sIp),
	}
	var affNum int64
	affNum, err = session.Insert(&smOpLog)
	if err != nil || affNum == 0 {
		return xorm.SilverMerchantOsLogs{}, errors.New("操作日志保存失败")
	}
	return smOpLog, err
}
