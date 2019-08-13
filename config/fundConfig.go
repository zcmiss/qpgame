package config

//资金交易类型
const (
	FUNDCHARGE               = 1  //充值
	FUNDWITHDRAW             = 2  //提现
	FUNDXIMA                 = 3  //洗码
	FUNDSAFEBOX              = 4  //保险箱存取款
	FUNDPRESENTER            = 5  //赠送彩金
	FUNDDISCOUNTS            = 6  //优惠入款
	FUNDBROKERAGE            = 7  //代理佣金提成
	FUNDSIGNIN               = 8  //签到奖励
	FUNDACTIVITYAWARD        = 9  //活动奖励
	FUNDWITHRAWNOPASS        = 10 //提现未通过返还
	FUNDPLATCHANGEIN         = 11 //平台资金转换-转入
	FUNDPLATCHANGEINFAILBACK = 12 //额度转换失败返还
	FUNDPLATCHANGEOUT        = 13 //额度转换-转出
	FUNDREDPACKET            = 14 //红包收入
	FUNDWITHDRAWBACK         = 15 //提现退款
	FUNDVIPUPLEVEL           = 16 //VIP晋级礼金
	FUNDVIPWEEK              = 17 //VIP周工资
	FUNDVIPMONTH             = 18 //VIP月工资
	FUNDGENERALIZEAWAED      = 19 //推广邀请奖励
	FUNDBINDPHONE            = 20 //绑定手机奖励
)

//资金变动类型与说明map
var FundChangeTypeMap = map[int]string{
	FUNDCHARGE:               "充值",
	FUNDWITHDRAW:             "提取现金",
	FUNDXIMA:                 "洗码",
	FUNDSAFEBOX:              "保险箱存取款",
	FUNDPRESENTER:            "赠送彩金",
	FUNDDISCOUNTS:            "优惠入款",
	FUNDBROKERAGE:            "代理佣金提成",
	FUNDSIGNIN:               "签到奖励",
	FUNDACTIVITYAWARD:        "活动奖励",
	FUNDWITHRAWNOPASS:        "提现未通过返还",
	FUNDPLATCHANGEIN:         "额度转换-转入",
	FUNDPLATCHANGEINFAILBACK: "额度转换失败返还",
	FUNDPLATCHANGEOUT:        "额度转换-转出",
	FUNDREDPACKET:            "红包收入",
	FUNDWITHDRAWBACK:         "提现退款",
	FUNDVIPUPLEVEL:           "VIP晋级礼金",
	FUNDVIPWEEK:              "VIP周工资",
	FUNDVIPMONTH:             "VIP月工资",
	FUNDGENERALIZEAWAED:      "推广邀请奖励",
	FUNDBINDPHONE:            "绑定手机奖励",
}

//需要进行提现打码记录的数据创建
var NeedWithDrawDamaReords = []int{
	FUNDCHARGE,
	FUNDXIMA,
	FUNDPRESENTER,
	FUNDDISCOUNTS,
	FUNDACTIVITYAWARD,
	FUNDREDPACKET,
	FUNDSIGNIN,
	FUNDVIPUPLEVEL,
	FUNDVIPWEEK,
	FUNDVIPMONTH,
	FUNDGENERALIZEAWAED,
	FUNDBINDPHONE,
}

//根据资金类型id获取指定中文说明
func GetFundChangeInfoByTypeId(typeId int) string {
	types := FundChangeTypeMap
	if value, keyExist := types[typeId]; keyExist {
		return value
	} else {
		return ""
	}
}
