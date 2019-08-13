package models

import (
	"qpgame/admin/common"
	"strconv"
	"time"
)

// 模型
type Activities struct{}

// 表名称
func (self *Activities) GetTableName(ctx *Context) string {
	return "activities"
}

// 得到所有记录-分页
func (self *Activities) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, title, sub_title, content, created, time_start, time_end, is_home_show, total_ip_limit,day_ip_limit, is_repeat, money, icon, type,"+
			"status, updated, activity_class_id ",
		func(ctx *Context) []string { //获取查询条件
			queries := getQueryFields(ctx, &map[string]string{
				"title":             "%",
				"sub_title":         "%",
				"status":            "=",
				"activity_class_id": "=",
				"is_home_show":      "=",
				"type":              "=",
			})
			queries = append(queries, getQueryFieldByTime(ctx, "created", "time_start", "time_end"))
			return queries
		},
		func(ctx *Context, row *map[string]string) {
			processDatetime(&[]string{"created", "time_start", "time_end"}, row)
			processOptions("status", &statusTypes, row)
			platform := (*ctx).Params().Get("platform")
			(*row)["activity_class_name"] = common.GetActivityClassName(platform, (*row)["activity_class_id"])
			processOptionsFor("is_home_show", "show_home", &yesNo, row)
		}, nil)
}

// 得到记录详情
func (self *Activities) GetRecordDetail(ctx *Context) (map[string]string, error) {
	return getRecordDetail(ctx, self, "",
		func(ctx *Context, row *map[string]string) {
			processDatetime(&[]string{"time_start", "time_end"}, row)
		})
}

// 添加记录
func (self *Activities) Save(ctx *Context) (int64, error) {
	return saveRecord(ctx, self, nil,
		func(ctx *Context, data *map[string]string) bool { //添加之前处理
			(*data)["created"] = strconv.FormatInt(time.Now().Unix(), 10) //添加时间
			(*data)["updated"] = strconv.FormatInt(time.Now().Unix(), 10) //修改时间
			turnDatetimeFields(&[]string{"time_end", "time_start"}, data) //转化时间字段
			(*ctx).Params().Set("log", "添加|添加活动: "+(*data)["title"])      //日志
			return true
		},
		func(ctx *Context, data *map[string]string) bool { //修改之前处理
			(*data)["updated"] = strconv.FormatInt(time.Now().Unix(), 10) //修改时间
			turnDatetimeFields(&[]string{"time_end", "time_start"}, data) //转化时间字段
			(*ctx).Params().Set("log", "修改|修改活动: "+(*data)["title"])      //日志
			return true
		}, getSavedFunc("活动信息", "title"))
}

// 删除记录
func (self *Activities) Delete(ctx *Context) error {
	return deleteRecord(ctx, self, nil, getDeletedFunc("活动信息"))
}
