package utils

import (
	"qpgame/models"
	"qpgame/models/xorm"
)

func SyncTable() {
	for _, v := range models.MyEngine {
		v.Sync2(
			new(xorm.AccountInfos),
			new(xorm.Accounts),
			new(xorm.AccountStatistics),
			new(xorm.Activities),
			new(xorm.ActivityClasses),
			new(xorm.ActivityRecords),
			new(xorm.AdminAccesses),
			new(xorm.AdminIps),
			new(xorm.AdminLoginLogs),
			new(xorm.AdminLogs),
			new(xorm.AdminNodes),
			new(xorm.AdminRoles),
			new(xorm.Admins),
			new(xorm.AppVersions),
			new(xorm.Blacklists),
			new(xorm.ChargeCards),
			new(xorm.ChargeRecords),
			new(xorm.ChargeTypes),
			new(xorm.Configs),
			new(xorm.ConversionRecords),
			new(xorm.GameCategories),
			new(xorm.Notices),
			new(xorm.PhoneCodes),
			new(xorm.Platforms),
			new(xorm.ProxyChessLevels),
			new(xorm.ProxyRealLevels),
			new(xorm.SystemNotices),
			new(xorm.UserBankCards),
			new(xorm.UserBanks),
			new(xorm.UserFeedbacks),
			new(xorm.UserLoginLogs),
			new(xorm.Users),
			new(xorm.VipLevels))

	}
}
