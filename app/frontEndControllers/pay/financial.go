package pay

import (
	"fmt"
	"net/http"
	"qpgame/common/log"
	"qpgame/models"
	"qpgame/models/xorm"
	"time"
)

type Financial struct {
	platform    string
	HttpRequest *http.Request
}

//构造函数
func NewFinancial(platform string) *Financial {
	finac := new(Financial)
	finac.platform = platform
	return finac
}

//检测用户能否充值,防止恶意刷充值记录
func (cthis *Financial) CheckUserByAddChargeRecord(userIds string) bool {
	//检测用户从当前时间往前10分钟之内，若有超过10次未充值成功的
	checkTime := time.Now().Unix() - 600 //当前时间往前10分钟
	mysqlClient := models.MyEngine[cthis.platform]
	record := xorm.ChargeRecords{}
	total, err := mysqlClient.Where("user_id = "+userIds).Where("state = 0").Where("created >= ?", checkTime).Count(&record)
	if err != nil {
		log.LogPrException(fmt.Sprintf("sql查询出现错误,错误内容为: %v", err))
		return true
	}
	if total >= 10 {
		return true
	}
	return false
}
