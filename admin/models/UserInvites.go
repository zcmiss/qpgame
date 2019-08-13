package models

import (
	"strconv"
)

// 模型
type UserInvites struct{}

// 表名称
func (self *UserInvites) GetTableName(ctx *Context) string {
	return "user_invites"
}

// 得到所有记录-分页
func (self *UserInvites) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, user_id, friend_id, created",
		func(ctx *Context) []string { //获取查询条件
			var conditions []string                              //查询条件数组
			if id, err := (*ctx).URLParamInt("id"); err == nil { //按编号查询
				conditions = append(conditions, "id = "+strconv.Itoa(id))
			}
			return conditions
		}, nil, nil)
}

// 得到记录详情
func (self *UserInvites) GetRecordDetail(ctx *Context) (map[string]string, error) {
	return getRecordDetail(ctx, self, "", nil)
}

// 添加记录
func (self *UserInvites) Save(ctx *Context) (int64, error) {
	return denySave()
}

// 删除记录
func (self *UserInvites) Delete(ctx *Context) error {
	return denyDelete()
}
