package models

import (
	"qpgame/admin/common"
	"qpgame/config"
	"strconv"
)

// 模型
type AccountInfos struct{}

// 表名称
func (self *AccountInfos) GetTableName(ctx *Context) string {
	return "account_infos"
}

// 得到所有记录-分页
func (self *AccountInfos) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self, "id, user_id, amount, balance, `type`, created, msg, order_id, charged_amount, charged_amount_old",
		func(ctx *Context) []string { //获取查询条件
			var queries []string
			queries = getQueryFields(ctx, &map[string]string{
				"type": "in", //交易类型
			})
			return append(queries, getQueryFieldByTime(ctx, "created", "created_start", "created_end"))
		},
		func(ctx *Context, row *map[string]string) {
			(*row)["user_name"] = common.GetUserName((*ctx).Params().Get("platform"), (*row)["user_id"])
			typeId, _ := strconv.Atoi((*row)["type"])
			typeName := config.GetFundChangeInfoByTypeId(typeId)
			(*row)["type_name"] = typeName
		}, nil)
}

// 得到记录详情
func (self *AccountInfos) GetRecordDetail(ctx *Context) (map[string]string, error) {
	return getRecordDetail(ctx, self, "", nil)
}

// 添加记录
func (self *AccountInfos) Save(ctx *Context) (int64, error) {
	return denySave()
}

// 删除记录
func (self *AccountInfos) Delete(ctx *Context) error {
	return denyDelete()
}
