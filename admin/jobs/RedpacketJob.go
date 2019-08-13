package jobs

import (
	"errors"
	"math/rand"
	"qpgame/admin/common"
	"strconv"
	"strings"
	"time"

	"github.com/go-xorm/xorm"
)

// 红包预派发
type RedpacketJob struct{}

// 时间设置
func (self *RedpacketJob) GetSpec() string {
	return "*/5 * * * * *"
}

// 执行事件
func (self *RedpacketJob) Process(db *xorm.Engine, platform string) {
	// 红包列表
	list := common.Redpackets[platform]
	// 更新红包状态 & 分发记录状态
	list = self.synsStatus(db, list)
	// 更新内存数据
	common.Redpackets[platform] = list
	// 生成红包分发记录 & 更新已分发状态
	self.generateLogs(db, platform, list)
}

// 检验是否已过期
func (self *RedpacketJob) checkIsExpire(rp map[string]string) bool {
	endTime, _ := strconv.Atoi(rp["end_time"])
	return int(time.Now().Unix()) > endTime
}

// 更新红包状态
func (self *RedpacketJob) synsStatus(db *xorm.Engine, list map[string]map[string]string) map[string]map[string]string {
	for id, rp := range list {
		if (rp["has_expired"] != "1") && (rp["status"] != "1") {
			if self.checkIsExpire(rp) {
				list[id]["status"] = "2"
				db.Exec("UPDATE redpacket_systems SET status=2 WHERE id=" + id)
			}
			db.Exec("UPDATE redpacket_receives SET is_get=2 WHERE redpacket_id=" + id + " AND is_get=0")
			list[id]["has_expired"] = "1"
		}
		if rp["status"] == "1" {
			if self.checkIsExpire(rp) {
				list[id]["status"] = "2"
				db.Exec("UPDATE redpacket_systems SET status=2 WHERE id=" + id)
				db.Exec("UPDATE redpacket_receives SET is_get=2 WHERE redpacket_id=" + id + " AND is_get=0")
				list[id]["has_expired"] = "1"
			}
		}
	}
	return list
}

// 检验是否已开始
func (self *RedpacketJob) checkIsRunning(rp map[string]string) bool {
	startTime, _ := strconv.Atoi(rp["start_time"])
	endTime, _ := strconv.Atoi(rp["end_time"])
	now := int(time.Now().Unix())
	return (rp["is_done"] == "0") && (rp["status"] == "1") && (now >= startTime) && (now <= endTime)
}

// 获取可分发红包清单
func (self *RedpacketJob) getAvailableRedpackets(list map[string]map[string]string) map[string]map[string]string {
	result := make(map[string]map[string]string)
	for id, rp := range list {
		if self.checkIsRunning(rp) {
			result[id] = rp
		}
	}
	return result
}

// 高纳德置乱算法
func (self *RedpacketJob) Random(strings []string, length int) ([]string, error) {
	if len(strings) < 1 {
		return make([]string, 0), errors.New("用户编号不能为空")
	}
	if len(strings) == 1 {
		return strings, nil
	}
	if (length <= 0) || (len(strings) <= length) {
		length = len(strings) - 1
	}
	for i := len(strings) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		strings[i], strings[num] = strings[num], strings[i]
	}
	result := make([]string, 0)
	for i := 0; i < length; i++ {
		result = append(result, strings[i])
	}
	return result, nil
}

// 获取用户编号集合
func (self *RedpacketJob) getUserIds(db *xorm.Engine, length int) []string {
	rows, _ := db.SQL("SELECT GROUP_CONCAT(id) ids FROM users WHERE status=1 AND user_type IN(0,1)").QueryString()
	if len(rows) == 0 {
		return make([]string, 0)
	}
	ids, _ := self.Random(strings.Split(rows[0]["ids"], ","), length)
	return ids
}

// 生成派奖金额，红包算法
func (self *RedpacketJob) generateMoney(isRandom bool, amount float64, count int, fixedMoney float64) float64 {
	if count == 1 {
		moneyStr := strconv.FormatFloat(amount, 'f', 3, 64)
		money, _ := strconv.ParseFloat(moneyStr, 64)
		return money
	}
	if amount == float64(0) {
		return amount
	}
	if !isRandom {
		moneyStr := strconv.FormatFloat(fixedMoney, 'f', 3, 64)
		money, _ := strconv.ParseFloat(moneyStr, 64)
		return money
	}
	rand.Seed(time.Now().UnixNano())
	var min = float64(1)
	max := amount / float64(count*2)
	money := rand.Float64()*max + min
	moneyStr := strconv.FormatFloat(money, 'f', 3, 64)
	money, _ = strconv.ParseFloat(moneyStr, 64)
	return money
}

// 生成分发数据
func (self *RedpacketJob) generateLogs(db *xorm.Engine, platform string, data map[string]map[string]string) {
	list := self.getAvailableRedpackets(data)
	currentTime := strconv.Itoa(int(time.Now().Unix()))
	for redpackageId, rp := range list {
		amount, _ := strconv.ParseFloat(rp["money"], 64)
		count, _ := strconv.Atoi(rp["total"])
		calculateType := rp["calculate_type"]
		if (amount < 1) || (count < 1) {
			continue
		}
		fixedMoney := float64(0)
		if calculateType != "1" {
			fixedMoney = amount / float64(count)
		}
		result := make([]string, 0)
		userIds := self.getUserIds(db, count)
		for _, userId := range userIds {
			money := self.generateMoney(calculateType == "1", amount, count, fixedMoney)
			if money > float64(0) {
				count--
				amount -= money
				moneyStr := strconv.FormatFloat(money, 'f', 3, 64)
				result = append(result, "("+redpackageId+","+userId+","+moneyStr+","+currentTime+","+rp["type"]+",0)")
			}
		}
		if len(result) > 0 {
			_, err := db.Exec("INSERT INTO redpacket_receives(redpacket_id,user_id,money,created,red_type,is_get) VALUES " + strings.Join(result, ","))
			if err == nil {
				db.Exec("UPDATE redpacket_systems SET is_done=1 WHERE id=" + redpackageId)
				common.Redpackets[platform][redpackageId]["is_done"] = "1"
			}
		}
	}
}
