package models

import (
	"strconv"
	"time"
)

// 模型
type AppVersions struct{}

// 表名称
func (self *AppVersions) GetTableName(ctx *Context) string {
	return "app_versions"
}

// 得到所有记录-分页
func (self *AppVersions) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, version, status, description, created, link, package_type, app_type, update_type, updated",
		func(ctx *Context) []string { //获取查询条件
			queries := getQueryFields(ctx, &map[string]string{
				"version":      "%",
				"status":       "=",
				"app_type":     "=",
				"package_type": "=",
				"update_type":  "=",
			})
			return append(queries, getQueryFieldByTime(ctx, "created", "time_start", "time_end"))
		},
		func(ctx *Context, row *map[string]string) {
			processOptions("status", &statusTypes, row)
			processOptions("update_type", &appUpdateTypes, row)
			processOptions("app_type", &appTypes, row)
			processOptions("package_type", &appPackageTypes, row)
		}, nil)
}

// 得到记录详情
func (self *AppVersions) GetRecordDetail(ctx *Context) (map[string]string, error) {
	return getRecordDetail(ctx, self, "", nil)
}

// 添加记录
func (self *AppVersions) Save(ctx *Context) (int64, error) {
	return saveRecord(ctx, self,
		func(ctx *Context, data *map[string]string) bool {
			(*data)["updated"] = strconv.FormatInt(time.Now().Unix(), 10) //修改时间
			return true
		},
		func(ctx *Context, data *map[string]string) bool { //添加之前处理
			(*data)["created"] = strconv.FormatInt(time.Now().Unix(), 10) //添加时间
			return true
		}, nil, getSavedFunc("APP版本", "version"))
}

// 删除记录
func (self *AppVersions) Delete(ctx *Context) error {
	return deleteRecord(ctx, self, nil, getDeletedFunc("APP版本"))
}
