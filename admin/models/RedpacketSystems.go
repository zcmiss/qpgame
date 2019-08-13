package models

import (
	"qpgame/admin/common"
	"strconv"
	"time"
)

// 模型
type RedpacketSystems struct{}

// 红包类型
var redpacketTypes = map[string]string{
	"1": "节日红包",
	"2": "每日幸运红包",
}

// 红包发放类型
var redpacketCalculateTypes = map[string]string{
	"1": "随机金额",
	"2": "固定金额",
}

// 表名称
func (self *RedpacketSystems) GetTableName(ctx *Context) string {
	return "redpacket_systems"
}

// 得到所有记录-分页
func (self *RedpacketSystems) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, type, money, sent_money, sent_count, total, created, status, "+
			"start_time, message, end_time, calculate_type",
		func(ctx *Context) []string { //获取查询条件
			queries := getQueryFields(ctx, &map[string]string{
				"type":           "=",
				"status":         "=",
				"calculate_type": "=",
			})
			queries = append(queries, getQueryFieldByTime(ctx, "created", "created_start", "created_end"))
			queries = append(queries, getQueryFieldByTimes(ctx, "start_time", "end_time"))
			return queries
		},
		func(ctx *Context, row *map[string]string) { //对于查询出来的每条记录的处理
			(*row)["type_text"] = redpacketTypes[(*row)["type"]]
			(*row)["calculate_type_text"] = redpacketCalculateTypes[(*row)["calculate_type"]]
			processDatetime(&[]string{"created", "start_time", "end_time"}, row)
		}, nil)
}

// 得到记录详情
func (self *RedpacketSystems) GetRecordDetail(ctx *Context) (map[string]string, error) {
	return getRecordDetail(ctx, self, "",
		func(ctx *Context, row *map[string]string) { //对于查询出来的每条记录的处理
			processDatetime(&[]string{"created", "start_time", "end_time"}, row)
		})
}

// 添加记录
func (self *RedpacketSystems) Save(ctx *Context) (int64, error) {
	id, err := saveRecord(ctx, self, nil,
		func(ctx *Context, data *map[string]string) bool { //添加之前处理
			turnDateFields(&[]string{"end_time", "start_time"}, data)
			(*data)["created"] = strconv.FormatInt(time.Now().Unix(), 10) //添加时间
			return true
		}, func(ctx *Context, data *map[string]string) bool { //更新之前处理
			turnDateFields(&[]string{"end_time", "start_time"}, data)
			(*data)["created"] = strconv.FormatInt(time.Now().Unix(), 10) //添加时间
			return true
		}, getSavedFunc("系统红包", "message"))
	if (id > 0) && (err == nil) {
		common.LoadRedpackets()
	}
	return id, err
}

// 删除记录
func (self *RedpacketSystems) Delete(ctx *Context) error {
	err := deleteRecord(ctx, self, nil, getDeletedFunc("系统红包"))
	if err == nil {
		common.LoadRedpackets()
	}
	return err
}
