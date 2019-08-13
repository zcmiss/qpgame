package task

import (
	"qpgame/app/fund"
	"qpgame/common/utils"
	"qpgame/config"
	"qpgame/models"
	"qpgame/models/xorm"
	"qpgame/ramcache"
	"strconv"
)

func VipWeekWage(platform string) {
	vipLevels, _ := ramcache.TableVipLevels.Load(platform)
	vips := vipLevels.([]xorm.VipLevels)
	startLevel := 0
	for i := 0; i <= len(vips); i++ {
		if vips[i].WeeklyAmount > 0 {
			startLevel = vips[i].Level
			break
		}
	}
	users := make([]xorm.Users, 0)
	models.MyEngine[platform].Where("vip_level >= ?", startLevel).Find(&users)
	for _, u := range users {
		go updateVipWeekWage(platform, vips, u)
	}
}

func VipMonthWage(platform string) {
	vipLevels, _ := ramcache.TableVipLevels.Load(platform)
	vips := vipLevels.([]xorm.VipLevels)
	startLevel := 0
	for i := 0; i <= len(vips); i++ {
		if vips[i].MonthAmount > 0 {
			startLevel = vips[i].Level
			break
		}
	}
	users := make([]xorm.Users, 0)
	models.MyEngine[platform].Where("vip_level >= ?", startLevel).Find(&users)
	for _, u := range users {
		go updateVipMonthWage(platform, vips, u)
	}
}

func updateVipWeekWage(platform string, vips []xorm.VipLevels, u xorm.Users) {
	amount := 0
	for _, v := range vips {
		if v.Level == u.VipLevel {
			amount = v.WeeklyAmount
			break
		}
	}
	info := map[string]interface{}{
		"user_id":     u.Id,
		"type_id":     config.FUNDVIPWEEK,
		"amount":      float64(amount),
		"order_id":    utils.CreationOrder("VW", strconv.Itoa(u.Id)),
		"msg":         "VIP" + strconv.Itoa(u.VipLevel) + "周工资",
		"finish_rate": 1.0, //需满足的打码量比例
	}
	balance := fund.NewUserFundChange(platform)
	balance.BalanceUpdate(info, nil)
}

func updateVipMonthWage(platform string, vips []xorm.VipLevels, u xorm.Users) {
	amount := 0
	for _, v := range vips {
		if v.Level == u.VipLevel {
			amount = v.MonthAmount
			break
		}
	}
	info := map[string]interface{}{
		"user_id":     u.Id,
		"type_id":     config.FUNDVIPMONTH,
		"amount":      float64(amount),
		"order_id":    utils.CreationOrder("VM", strconv.Itoa(u.Id)),
		"msg":         "VIP" + strconv.Itoa(u.VipLevel) + "月工资",
		"finish_rate": 1.0, //需满足的打码量比例
	}
	balance := fund.NewUserFundChange(platform)
	balance.BalanceUpdate(info, nil)
}
