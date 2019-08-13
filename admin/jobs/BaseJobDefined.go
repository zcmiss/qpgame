package jobs

import "github.com/go-xorm/xorm"

// 运营报表统计
var systemStatistics = &SystemStatisticsJob{}

// 用户统计
var accountStatistics = &AccountStatisticsJob{}

// 代理统计
var proxyStatistics = &ProxyStatisticsJob{}

// 红包分发
var redpackets = &RedpacketJob{}

// Websocket
var websocket = &WebsocketJob{}

// 每页显示的数量
var pageSize = 500

// 后台的任务的统一格式
type AdminJob struct {
	Spec string
	Fn   func()
}

// 接口: 后台任务处理
type AdminCronJob interface {
	GetSpec() string
	Process(db *xorm.Engine, platform string)
}
