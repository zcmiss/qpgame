package models

import (
	"strconv"
	"time"
)

// 模型
type ProxyRealLevels struct{}

// 表名称
func (self *ProxyRealLevels) GetTableName(ctx *Context) string {
	return "proxy_real_levels"
}

// 得到所有记录-分页
func (self *ProxyRealLevels) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, level, name, team_total_low, team_total_limit, commission, created",
		func(ctx *Context) []string { //获取查询条件
			queries := getQueryFields(ctx, &map[string]string{
				"name": "%",
			})
			return append(queries, getQueryFieldByTime(ctx, "created", "time_start", "time_end"))
		}, nil, nil)
}

// 得到记录详情
func (self *ProxyRealLevels) GetRecordDetail(ctx *Context) (map[string]string, error) {
	return getRecordDetail(ctx, self, "", nil)
}

// 添加记录
func (self *ProxyRealLevels) Save(ctx *Context) (int64, error) {
	return saveRecord(ctx, self, nil,
		func(ctx *Context, data *map[string]string) bool { //添加之前处理
			(*data)["created"] = strconv.FormatInt(time.Now().Unix(), 10) //添加时间
			return true
		}, nil, getSavedFunc("视讯等级", "name"))
}

// 删除记录
func (self *ProxyRealLevels) Delete(ctx *Context) error {
	return denyDelete()
}
