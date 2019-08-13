package models

import (
	"qpgame/admin/common"
	"qpgame/common/utils"
	"strings"
)

// 模型
type Blacklists struct{}

// 表名称
func (self *Blacklists) GetTableName(ctx *Context) string {
	path := (*ctx).Path()
	if strings.LastIndex(path, "/update") > 0 ||
		strings.LastIndex(path, "/add") > 0 ||
		strings.LastIndex(path, "/delete") > 0 {
		return "blacklists"
	}
	return "blacklists LEFT JOIN users ON blacklists.user_id = users.id"
}

// 得到所有记录-分页
func (self *Blacklists) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"blacklists.id, user_id, phone, last_login_time, status, qq, wechat",
		nil,
		func(ctx *Context, row *map[string]string) {
			processDatetime(&[]string{"last_login_time"}, row)
			processOptionsFor("status", "status_name", &statusTypes, row)
			(*row)["user_name"] = common.GetUserName((*ctx).Params().Get("platform"), (*row)["user_id"])
		}, nil)
}

// 得到记录详情
func (self *Blacklists) GetRecordDetail(ctx *Context) (map[string]string, error) {
	return getRecordDetail(ctx, self, "", nil)
}

// 添加记录
func (self *Blacklists) Save(ctx *Context) (int64, error) {
	return saveRecord(ctx, self,
		func(ctx *Context, data *map[string]string) bool {
			postData := utils.GetPostData(ctx)
			userId := common.GetIdFromUserName((*ctx).Params().Get("platform"), postData.Get("name"))
			if userId == "" {
				return false
			}
			(*data)["user_id"] = userId
			return true
		}, nil, nil, getSavedFunc("用户黑名单", "user_id"))
}

// 删除记录
func (self *Blacklists) Delete(ctx *Context) error {
	return deleteRecord(ctx, self, nil, getDeletedFunc("用户黑名单"))
}
