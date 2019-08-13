package models

import (
	"qpgame/admin/common"
	"strconv"
	"time"
)

// 模型
type Notices struct{}

// 表名称
func (self *Notices) GetTableName(ctx *Context) string {
	return "notices"
}

// 得到所有记录-分页
func (self *Notices) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self, "id, title, user_id, status, created, content",
		func(ctx *Context) []string { //获取查询条件
			queries := getQueryFields(ctx, &map[string]string{
				"status": "=",
				"title":  "%",
			})
			return append(queries, getQueryFieldByTime(ctx, "created", "time_start", "time_end"))
		},
		func(ctx *Context, row *map[string]string) {
			processDatetime(&[]string{"created"}, row)
			processOptions("status", &map[string]string{"0": "已锁定", "1": "正常"}, row)
			(*row)["user_name"] = common.GetUserName((*ctx).Params().Get("platform"), (*row)["user_id"])
		}, nil)
}

// 得到记录详情
func (self *Notices) GetRecordDetail(ctx *Context) (map[string]string, error) {
	return getRecordDetail(ctx, self, "",
		func(ctx *Context, row *map[string]string) {
			(*row)["user_id"] = common.GetUserName((*ctx).Params().Get("platform"), (*row)["user_id"])
			processDatetime(&[]string{"created"}, row)
		})
}

// 添加记录
func (self *Notices) Save(ctx *Context) (int64, error) {
	return saveRecord(ctx, self,
		func(ctx *Context, data *map[string]string) bool {
			(*data)["user_id"] = (*ctx).Params().Get("user_id") //设置真正的user_id
			return true
		},
		func(ctx *Context, data *map[string]string) bool { //添加之前处理
			(*data)["created"] = strconv.FormatInt(time.Now().Unix(), 10) //添加时间
			return true
		}, nil, getSavedFunc("站内公告", "title"))
}

// 删除记录
func (self *Notices) Delete(ctx *Context) error {
	return deleteRecord(ctx, self, nil, getDeletedFunc("站内公告"))
}
