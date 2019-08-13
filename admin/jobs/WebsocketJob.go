package jobs

import (
	"qpgame/admin/websockets"

	"github.com/go-xorm/xorm"
)

// Websocket需要的数据统计
type WebsocketJob struct{}

// 时间设置
func (self *WebsocketJob) GetSpec() string {
	return "*/15 * * * * *"
}

// 处理作务
func (self *WebsocketJob) Process(db *xorm.Engine, platform string) {
	chargeTotal, withdrawTotal, silverMerchantTotal := websockets.GetCountInfo(db)
	for c := range websockets.WSConnecition[platform] {
		data := map[string]int{
			"charge_count":          chargeTotal,
			"withdraw_count":        withdrawTotal,
			"silver_merchant_count": silverMerchantTotal,
		}
		websockets.ResponseData(c, "statistic", data)
	}
}
