package models

import (
	"qpgame/admin/common"

	"github.com/kataras/iris"
)

type Context = iris.Context //类型别名
type Error = common.Error

// 实现后台的模型接口
type IAdminModel interface {
	GetTableName(ctx *Context) string                        //得到表名
	GetRecords(ctx *Context) (Pager, error)                  //得到多条记录 (分页信息, 内部错误信息)
	GetRecordDetail(ctx *Context) (map[string]string, error) //得到单条记录的详情 (记录信息, 内部错误信息)
	Save(ctx *Context) (int64, error)                        //保存数据 (操作的记录ID, 内部错误信息)
	Delete(ctx *Context) error                               //删除记录 (操作结果, 内部错误信息)
}

// 分页
type Pager struct {
	Rows      []map[string]string `json:"rows"`       //记录结果集
	Page      int                 `json:"page"`       //当前页数
	PageCount int                 `json:"page_count"` //总计页数
	TotalRows int                 `json:"total_rows"` //总计记录数
	PageSize  int                 `json:"page_size"`  //每页记录数
}

// 默认的分页信息
func newPager() Pager {
	return Pager{
		Rows:      nil, //记录数
		Page:      0,   //当前页
		PageCount: 0,   //总页数
		TotalRows: 0,   //记录总数
		PageSize:  0,   //每页记录数
	}
}

//app更新类型
var appUpdateTypes = map[string]string{
	"1": "强制更新",
	"2": "提示更新",
	"3": "不提示更新",
}

//app类型
var appTypes = map[string]string{
	"0": "其他",
	"1": "安卓",
	"2": "IOS",
}

//app包类型
var appPackageTypes = map[string]string{
	"1": "全量包",
	"2": "增量包",
}

var statusTypes = map[string]string{
	"0": "锁定",
	"1": "正常",
}
var statusType = map[string]string{
	"0": "待处理",
	"1": "成功",
	"2": "失败",
	"3": "进行中",
	"4": "退款",
	"5": "取消",
	"6": "强制入款",
}

var adminPermissions = map[string]string{
	"0": "普通权限",
	"1": "主管权限",
}

//yes-no选项
var yesNo = map[string]string{
	"0": "否",
	"1": "是",
}

//红包类型
var RedTypes = map[string]string{
	"0": "未知红包",
	"1": "节日红包",
	"2": "每日幸运红包",
}

//代理类型
var ProxyTypes = map[string]string{
	"2": "棋牌游戏",
	"5": "真人视讯",
}

/** 关于系统配置里面的一些字段的定义 **/
// 充值信息
type ChargeInfo struct {
	Url     string `json:"url"`
	Name    string `json:"name"`
	Account string `json:"account"`
}

//代理用户充值
type ProxyCharge struct {
	ProxyChargeLogo string       `json:"proxy_charge_logo"`
	ChargeAccounts  []ChargeInfo `json:"charge_accounts"`
	Info            string       `json:"info"`
	State           int          `json:"state"`
}

//客服信息
type ServiceInfo struct {
	Url     string `json:"url"`
	Name    string `json:"name"`
	Account string `json:"account"`
	Info    string `json:"info"`
}

//客服信息
type ServiceAccount struct {
	Wx []ServiceInfo `json:"wx"`
	Qq []ServiceInfo `json:"qq"`
}

//支付配置
type RegisterConf struct {
	CanRegister int `json:"can_register"`
}

//订单提醒配置
type OrderAlert struct {
	ChargeUrl   string `json:"charge_url"`
	WithdrawUrl string `json:"withdraw_url"`
	Timeout     int    `json:"timeout"`
}

//后台用户操作日志
type AdminLog struct {
	//AdminId   int    //管理员编号
	//AdminName string //管理员名称
	Type string //操作类型
	//Node    string //节点名称
	Content string //日志内容
}

// 人工入款项目
var manualChargeItems = map[string]string{
	"1": "人工存入-存款",
	"2": "注册优惠",
	"3": "优惠活动",
	"4": "余额负数冲正",
	"5": "派彩错误",
	"6": "其他错误",
}

// 人工出款项目
var manualWithdrawItems = map[string]string{
	"1": "人工出款-取款",
	"2": "优惠扣除",
	"3": "异常出款-取款",
}
