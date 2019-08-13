package task

import (
	"encoding/json"
	"github.com/kataras/iris/core/errors"
	"github.com/shopspring/decimal"
	"qpgame/app/fund"
	"qpgame/common/utils/game"
	"qpgame/models"
	"qpgame/models/xorm"
	"qpgame/ramcache"
	"strconv"

	goxorm "github.com/go-xorm/xorm"
)

// 处理游戏平台成功转入余额后，未转入到用户余额的异常
func ExceptionTasks(platform string) {
	tet2, _ := ramcache.TableExceptionTasks.Load(platform)
	if tet2 == nil {
		return
	}
	tet2Entity := tet2.(map[string]xorm.ExceptionTasks)
	engine := models.MyEngine[platform]
	for idx, tetItem := range tet2Entity {
		taskContent := tetItem.TaskContent
		jo := make(map[string]interface{})
		err := json.Unmarshal([]byte(taskContent), &jo)
		_, ok := jo["order_id"]
		if err == nil && ok {
			iUserId := tetItem.UserId
			iPlatId := tetItem.PlatId
			jo["user_id"] = iUserId
			jo["type_id"] = int(jo["type_id"].(float64))
			num, numOk := jo["retry_times"] // 重试次数
			tryNum := 1
			if numOk {
				tryNum = int(num.(float64)) + 1
			}
			jo["retry_times"] = tryNum
			userBalance := fund.NewUserFundChange(platform)
			retryFlag := false
			if tetItem.Flag == 1 {
				callback := func(session *goxorm.Session, args ...interface{}) (interface{}, error) {
					affNum, err := session.ID(tetItem.Id).Delete(xorm.ExceptionTasks{})
					if err != nil || affNum == 0 {
						session.Rollback()
						return nil, errors.New("删除已完成任务失败")
					}
					delete(tet2Entity, idx)
					return nil, nil
				}
				res := userBalance.BalanceUpdate(jo, callback)
				if res["status"] != 1 {
					retryFlag = true
				}
			} else if tetItem.Flag == 0 {
				platAccount := jo["plat_account"].(string)
				callback := func(session *goxorm.Session, args ...interface{}) (interface{}, error) {
					user := xorm.Users{Id: iUserId, LastPlatformId: 0}
					accountsId := args[0].(int)
					infoAmount := args[1].(decimal.Decimal)
					_, err := session.Cols("last_platform_id").Where("id = ?", iUserId).Update(user)
					if err != nil {
						return nil, err
					}
					if game.Uchips(platAccount, strconv.Itoa(accountsId), infoAmount.Mul(decimal.New(-1, 0)).String(), iPlatId, platform) {
						affNum, err := session.ID(tetItem.Id).Delete(xorm.ExceptionTasks{})
						if err != nil || affNum == 0 {
							session.Rollback()
							return nil, errors.New("删除已完成任务失败")
						}
						delete(tet2Entity, idx)
						return nil, nil
					} else {
						return nil, errors.New("平台余额转出失败")
					}
				}
				res := userBalance.BalanceUpdate(jo, callback)
				if res["status"] != 1 {
					retryFlag = true
				}
			}
			if retryFlag {
				// 尝试3次转换用户余额，仍然失败则修改flag，人工查验
				joStr, joErr := json.Marshal(jo)
				tetItem.TaskContent = string(joStr)
				if tryNum < 3 {
					if joErr == nil {
						tet2Entity[idx] = tetItem
					}
				} else {
					tetItem.Flag = tetItem.Flag - 10
					engine.ID(tetItem.Id).Update(tetItem)
					delete(tet2Entity, idx)
				}
			}
		}
	}
}
