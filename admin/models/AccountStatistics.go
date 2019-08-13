package models

import (
	"qpgame/admin/common"
	"strconv"
)

// 模型
type AccountStatistics struct{}

// 表名称
func (self *AccountStatistics) GetTableName(ctx *Context) string {
	return "account_statistics"
}

// 得到所有记录-分页
func (self *AccountStatistics) GetRecords(ctx *Context) (Pager, error) {
	//读取分页的相关信息，生成排名的起始数字
	page := (*ctx).URLParamIntDefault("page", 1)
	pageSize := (*ctx).URLParamIntDefault("page_size", 20)
	index := (page - 1) * pageSize
	return getRecords(ctx, self,
		"user_id, SUM(charged_amount) AS charged_amount, "+ //充值总额
			"SUM(consumed_amount) AS consumed_amount, SUM(withdraw_amount) AS withdraw_amount, "+ //消费金额,提现总额
			"SUM(bet_amount) AS bet_amount, SUM(reward_amount) AS reward_amount, "+ //抽注金额,中奖金额
			"SUM(wash_amount) AS wash_amount, SUM(proxy_commission) AS proxy_commission ", //洗码金码,代理佣金
		func(ctx *Context) []string { //获取查询条件
			queries := getQueryFieldByDate(ctx, "ymd", "ymd_start", "ymd_end")
			return []string{queries}
		},
		func(ctx *Context, row *map[string]string) {
			index += 1
			(*row)["id"] = strconv.Itoa(index) //序号
			(*row)["user_name"] = common.GetUserName((*ctx).Params().Get("platform"), (*row)["user_id"])
		}, func(ctx *Context) (string, string, int) {
			sortType := (*ctx).URLParam("sort_type")
			fields := map[string]int{
				"charged_amount":   0,
				"consumed_amount":  0,
				"withdraw_amount":  0,
				"bet_amount":       0,
				"reward_amount":    0,
				"wash_amount":      0,
				"proxy_commission": 0,
			}
			_, exists := fields[sortType]
			if exists {
				return "user_id", sortType + " DESC", 0
			}
			return "user_id", "bet_amount DESC", 0
		})
}

// 得到记录详情
func (self *AccountStatistics) GetRecordDetail(ctx *Context) (map[string]string, error) {
	return denyDetail()
}

// 添加记录
func (self *AccountStatistics) Save(ctx *Context) (int64, error) {
	return denySave()
}

// 删除记录
func (self *AccountStatistics) Delete(ctx *Context) error {
	return denyDelete()
}
