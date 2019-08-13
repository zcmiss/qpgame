package models

import (
	"strconv"
	"time"
)

// 模型
type SystemNotices struct{}

// 表名称
func (self *SystemNotices) GetTableName(ctx *Context) string {
	return "system_notices"
}

// 得到所有记录-分页
func (self *SystemNotices) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, title, content, status, created",
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
		}, nil)
}

// 得到记录详情
func (self *SystemNotices) GetRecordDetail(ctx *Context) (map[string]string, error) {
	return getRecordDetail(ctx, self, "",
		func(ctx *Context, row *map[string]string) {
			processDatetime(&[]string{"created"}, row)
		})
}

// 添加记录
func (self *SystemNotices) Save(ctx *Context) (int64, error) {
	return saveRecord(ctx, self, nil,
		func(ctx *Context, data *map[string]string) bool { //添加之前处理
			(*data)["created"] = strconv.FormatInt(time.Now().Unix(), 10) //添加时间
			return true
		}, nil, getSavedFunc("系统公告", "title"))
}

// 删除记录
func (self *SystemNotices) Delete(ctx *Context) error {
	return deleteRecord(ctx, self, nil, getDeletedFunc("系统公告"))
}
