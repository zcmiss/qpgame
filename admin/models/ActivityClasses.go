package models

import (
	"strconv"
	"time"
)

// 模型
type ActivityClasses struct{}

// 表名称
func (self *ActivityClasses) GetTableName(ctx *Context) string {
	return "activity_classes"
}

// 得到所有记录-分页
func (self *ActivityClasses) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, name, status, seq, created, updated",
		func(ctx *Context) []string { //获取查询条件
			return getQueryFields(ctx, &map[string]string{
				"name":   "%",
				"status": "=",
			})
		},
		func(ctx *Context, row *map[string]string) {
			processDatetime(&[]string{"created", "updated"}, row)
			processOptions("status", &statusTypes, row)
		}, nil)
}

// 得到记录详情
func (self *ActivityClasses) GetRecordDetail(ctx *Context) (map[string]string, error) {
	return getRecordDetail(ctx, self, "", nil)
}

// 添加记录
func (self *ActivityClasses) Save(ctx *Context) (int64, error) {
	return saveRecord(ctx, self, nil,
		func(ctx *Context, data *map[string]string) bool { //添加之前处理
			(*data)["created"] = strconv.FormatInt(time.Now().Unix(), 10) //添加时间
			return true
		},
		func(ctx *Context, data *map[string]string) bool { //修改之前处理
			(*data)["updated"] = strconv.FormatInt(time.Now().Unix(), 10) //修改时间
			return true
		}, getSavedFunc("活动分类", "name"))
}

// 删除记录
func (self *ActivityClasses) Delete(ctx *Context) error {
	return deleteRecord(ctx, self, nil, getDeletedFunc("活动分类"))
}
