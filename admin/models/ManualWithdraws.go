package models

import (
	"qpgame/admin/common"
	"qpgame/common/utils"
	"qpgame/config"
	"strconv"
	"time"
)

// 模型
type ManualWithdraws struct{}

// 表名称
func (self *ManualWithdraws) GetTableName(ctx *Context) string {
	return "manual_withdraws"
}

// 得到所有记录-分页
func (self *ManualWithdraws) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, user_id, `order`, amount, quantity, item, comment, deal_time, operator, state",
		func(ctx *Context) []string { //获取查询条件
			queries := getQueryFields(ctx, &map[string]string{
				"order": "%",
				"item":  "=",
				"state": "=",
			})
			return append(queries, getQueryFieldByTime(ctx, "deal_time", "deal_start", "deal_end"))
		},
		func(ctx *Context, row *map[string]string) {
			processOptionsFor("state", "state_name", &map[string]string{
				"0": "待审核",
				"1": "审核通过",
				"2": "作废",
			}, row)
			processOptionsFor("item", "item_name", &manualWithdrawItems, row)
			processDatetime(&[]string{"deal_time"}, row) //时间戳转换成字符串
			(*row)["user_name"] = common.GetUserName((*ctx).Params().Get("platform"), (*row)["user_id"])
		}, nil)
}

// 得到记录详情
func (self *ManualWithdraws) GetRecordDetail(ctx *Context) (map[string]string, error) {
	return getRecordDetail(ctx, self, "", nil)
}

// 添加记录
func (self *ManualWithdraws) Save(ctx *Context) (int64, error) {
	return saveRecord(ctx, self, nil,
		func(ctx *Context, data *map[string]string) bool { //添加之前处理
			post := utils.GetPostData(ctx)
			(*data)["state"] = "0"
			(*data)["order"] = utils.CreationOrder("WH", post.Get("user_id"))
			//(*data)["operator"] = common.GetAdmin(ctx)["name"]
			(*data)["deal_time"] = strconv.FormatInt(time.Now().Unix(), 10) //添加时间
			return true
		}, nil, getSavedFunc("人工出款", "user_id"))
}

// 删除记录
func (self *ManualWithdraws) Delete(ctx *Context) error {
	return denyDelete()
}

// 通过人工出款
func (self *ManualWithdraws) Allow(ctx *Context) error {
	return changeManualState(ctx, self.GetTableName(ctx), config.FUNDWITHDRAW, "1")
}

// 作废人工出款
func (self *ManualWithdraws) Deny(ctx *Context) error {
	return changeManualState(ctx, self.GetTableName(ctx), config.FUNDWITHDRAW, "2")
}
