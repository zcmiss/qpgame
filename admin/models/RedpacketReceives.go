package models

import (
	"qpgame/admin/common"
	"strconv"
	"time"
)

// 模型
type RedpacketReceives struct{}

// 表名称
func (self *RedpacketReceives) GetTableName(ctx *Context) string {
	return "redpacket_receives AS r LEFT JOIN redpacket_systems As s ON r.redpacket_id = s.id "
}

// 得到所有记录-分页
func (self *RedpacketReceives) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"r.id, r.redpacket_id, s.message As redpacket_name, r.user_id, r.money, r.created, r.red_type, r.is_get",
		func(ctx *Context) []string { //获取查询条件
			queries := getQueryFields(ctx, &map[string]string{
				"redpacket_id": "%",
				"red_type":     "=",
			})
			return append(queries, getQueryFieldByTime(ctx, "created", "created_start", "created_end"))
		},
		func(ctx *Context, row *map[string]string) { //对于查询出来的每条记录的处理
			(*row)["user_name"] = common.GetUserName((*ctx).Params().Get("platform"), (*row)["user_id"])
			processOptionsFor("red_type", "red_type_name", &RedTypes, row)
		}, nil)
}

// 得到记录详情
func (self *RedpacketReceives) GetRecordDetail(ctx *Context) (map[string]string, error) {
	return getRecordDetail(ctx, self, "",
		func(ctx *Context, row *map[string]string) { //对于查询出来的每条记录的处理
			if createTime, createErr := strconv.ParseInt((*row)["created"], 10, 64); createErr == nil {
				(*row)["created"] = time.Unix(createTime, 0).Format("2006-01-02 15:04:05") //添加时间
			}
		})
}

// 添加记录
func (self *RedpacketReceives) Save(ctx *Context) (int64, error) {
	return denySave()
}

// 删除记录
func (self *RedpacketReceives) Delete(ctx *Context) error {
	return denyDelete()
}
