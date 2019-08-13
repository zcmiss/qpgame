package task

import (
	"fmt"
	"github.com/shopspring/decimal"
	"qpgame/common/log"
	"qpgame/common/utils"
	"qpgame/models"
	"qpgame/models/xorm"
	"qpgame/ramcache"
	"strconv"
	"strings"
	"time"
)

//代理佣金生成
func CommissionPayment(platform string) {

	updateUserSql := "UPDATE users u SET u.`downIds`=IFNULL((SELECT GROUP_CONCAT(id) FROM (SELECT id,parent_id FROM users)a WHERE a.`parent_id`=u.id ),'')"
	jdbc := models.MyEngine[platform]
	//更新用户下级信息
	jdbc.Exec(updateUserSql)
	//获取本次数据产生的时间
	payTime := utils.GetNowTime()
	//本次数据查询时间字符串
	queryTime := time.Now().AddDate(0, 0, -1).Format("20060102")
	//获取昨天0点到23点59分59秒的时间戳
	yesterday := time.Now().Add(time.Hour * -24)
	localTime, _ := time.LoadLocation("Asia/Shanghai")
	st, _ := time.ParseInLocation("20060102", yesterday.Format("20060102"), localTime)
	started := strconv.Itoa(int(st.Unix()))
	ended := strconv.Itoa(int(st.Add(time.Hour * 24).Add(time.Second * -1).Unix()))

	unionAll := ""
	for i := 0; i <= 9; i++ {
		unionAll += "SELECT b.`user_id`,g.`parent_id`,SUM(b.`amount`) totalAmount,u.downIds FROM bets_" + strconv.Itoa(i)
		unionAll += " b LEFT JOIN platform_games p ON b.`game_code`=p.`service_code` AND p.`plat_id`=b.`platform_id` "
		unionAll += "LEFT JOIN game_categories g ON p.`game_categorie_id`=g.`id` LEFT JOIN users u ON b.`user_id` = u.`id` WHERE b.`ented` >= "
		unionAll += started + " AND b.`ented` <= " + ended + " AND g.`parent_id` IN(2,5) GROUP BY b.`user_id`,g.`parent_id` "
		if i != 9 {
			unionAll += "UNION ALL "
		}
	}
	res, err := jdbc.Query(unionAll)
	if err != nil {
		log.DeferRecover()
		log.Log.Error("获取当天投注记录失败，平台编号：", platform, " 报错信息：", err.Error())
		return
	}
	userMap := make(map[string]map[string]interface{})
	//计算个人业绩
	for _, r := range res {
		userId := string(r["user_id"])
		uMap := make(map[string]interface{})
		uMap["user_id"] = userId
		//个人业绩
		uMap["betAmount"], _ = decimal.NewFromString(string(r["totalAmount"]))
		//总业绩
		uMap["totalAmount"], _ = decimal.NewFromString(string(r["totalAmount"]))
		//代理类型2棋牌5真人视讯
		uMap["proxyType"] = string(r["parent_id"])
		//佣金生成时间
		uMap["created"] = payTime
		//佣金生成时间代号
		uMap["created_str"] = queryTime
		//上级用户编号
		uMap["parent_id"] = ""
		//贡献数 不包括自己
		uMap["contributions"] = 0
		//团队成员
		downIds := string(r["downIds"])
		if downIds != "" {
			uMap["downIds"] = strings.Split(downIds, ",")
		}
		userMap[userId+"-"+uMap["proxyType"].(string)] = uMap
	}

	//加载棋牌代理等级表
	cLevels, _ := ramcache.TableProxyChessLevels.Load(platform)
	//加载真人视讯代理等级表
	rLevels, _ := ramcache.TableProxyRealLevels.Load(platform)

	//计算总业绩
	for k, v := range userMap {
		if v["downIds"] != nil {
			for _, s := range v["downIds"].([]string) {
				key := s + "-" + v["proxyType"].(string)
				_, ok := userMap[key]
				if ok {
					v["contributions"] = v["contributions"].(int) + 1
					v["totalAmount"] = v["totalAmount"].(decimal.Decimal).Add(userMap[key]["betAmount"].(decimal.Decimal))
				}
			}
		}
		//如果是棋牌
		if v["proxyType"] == "2" {
			for _, c := range cLevels.([]xorm.ProxyChessLevels) {
				teamTotalLow := decimal.New(int64(c.TeamTotalLow*10000), 0)
				teamTotalLimit := decimal.New(int64(c.TeamTotalLimit*10000), 0)
				vTotalAmount := v["totalAmount"].(decimal.Decimal)
				//大于等于起始资金，小于对应等级的封顶时间，小于其实是可以去掉的，因为ProxyChessLevels是倒序，
				// 为了以防万一意外调整对佣金计算的影响特加上小于的限制更加保险
				if vTotalAmount.GreaterThanOrEqual(teamTotalLow) && vTotalAmount.LessThan(teamTotalLimit) {
					v["proxyLevelName"] = c.Name
					v["proxyLevel"] = c.Level
					commisRate := float64(c.Commission) / 10000
					v["rate"] = commisRate
					v["commission"] = v["betAmount"].(decimal.Decimal).Mul(decimal.NewFromFloat(commisRate))
					v["totalcommission"] = v["totalAmount"].(decimal.Decimal).Mul(decimal.NewFromFloat(commisRate))
					break
				}
			}
		}
		if v["proxyType"] == "5" {
			for _, c := range rLevels.([]xorm.ProxyRealLevels) {
				teamTotalLow := decimal.New(int64(c.TeamTotalLow*10000), 0)
				teamTotalLimit := decimal.New(int64(c.TeamTotalLimit*10000), 0)
				vTotalAmount := v["totalAmount"].(decimal.Decimal)
				if vTotalAmount.GreaterThanOrEqual(teamTotalLow) && vTotalAmount.LessThan(teamTotalLimit) {
					v["proxyLevelName"] = c.Name
					v["proxyLevel"] = c.Level
					commisRate := float64(c.Commission) / 10000
					v["rate"] = commisRate
					v["commission"] = userMap[k]["betAmount"].(decimal.Decimal).Mul(decimal.NewFromFloat(commisRate))
					v["totalcommission"] = userMap[k]["totalAmount"].(decimal.Decimal).Mul(decimal.NewFromFloat(commisRate))
					break
				}
			}
		}
	}
	//计算总佣金
	for _, v := range userMap {
		if v["downIds"] != nil {
			//自己佣金比例
			selfCommisRate := decimal.NewFromFloat(v["rate"].(float64))
			for _, s := range v["downIds"].([]string) {
				key := s + "-" + v["proxyType"].(string)
				_, ok := userMap[key]
				if ok {
					userMap[key]["parent_id"] = v["user_id"]
					//下线佣金比例
					downlineCommisRate := decimal.NewFromFloat(userMap[key]["rate"].(float64))
					totalCommission := v["totalcommission"].(decimal.Decimal)
					//自己返点大于下线返点
					if selfCommisRate.GreaterThanOrEqual(downlineCommisRate) {
						v["totalcommission"] = totalCommission.Sub(userMap[key]["commission"].(decimal.Decimal))
						//下线返点大于上级代理
					} else {
						v["totalcommission"] = totalCommission.Sub(selfCommisRate.Mul(userMap[key]["betAmount"].(decimal.Decimal)))
					}
				}
			}
		}
	}
	//如果有重复的就覆盖替换
	insertField := "insert into proxy_commissions "
	insertField += "(user_id,commission,parent_id,bet_amount,total_amount,total_commission,created,created_str,proxy_level,"
	insertField += "proxy_level_name,proxy_type,states,contributions,proxy_level_rate) values ("
	userMapLen := len(userMap)
	autoIncrease := 0
	for _, v := range userMap {
		//循环计数
		autoIncrease += 1
		insertField += v["user_id"].(string) + ","
		insertField += v["commission"].(decimal.Decimal).String() + ","
		parentId := "0"
		if v["parent_id"].(string) != "" {
			parentId = v["parent_id"].(string)
		}
		insertField += parentId + ","
		insertField += v["betAmount"].(decimal.Decimal).String() + ","
		insertField += v["totalAmount"].(decimal.Decimal).String() + ","
		insertField += v["totalcommission"].(decimal.Decimal).String() + ","
		insertField += strconv.Itoa(v["created"].(int)) + ","
		insertField += "'" + v["created_str"].(string) + "',"
		insertField += strconv.Itoa(v["proxyLevel"].(int)) + ","
		insertField += "'" + v["proxyLevelName"].(string) + "',"
		insertField += v["proxyType"].(string) + ","
		insertField += "0,"
		insertField += strconv.Itoa(v["contributions"].(int)) + ","
		insertField += decimal.NewFromFloat(v["rate"].(float64)).String() + ")"
		if autoIncrease != userMapLen {
			insertField += ",("
		}
	}
	insertField += " ON DUPLICATE KEY UPDATE commission=VALUES(commission),bet_amount=VALUES(bet_amount),total_amount=VALUES(total_amount),"
	insertField += "total_commission=VALUES(total_commission),contributions=VALUES(contributions),proxy_level_rate=VALUES(proxy_level_rate),"
	insertField += "states=VALUES(states),proxy_level_name=VALUES(proxy_level_name),proxy_level=VALUES(proxy_level),created=VALUES(created)"
	_, errInsert := jdbc.Exec(insertField)
	if errInsert != nil {
		fmt.Println(err.Error())
	}
}

//beans := make([]xorm.ProxyCommissions, 0)
//var bean xorm.ProxyCommissions
//bean.UserId, _ = strconv.Atoi(v["user_id"].(string))
//bean.Commission = v["commission"].(decimal.Decimal).String()
//bean.ParentId, _ = strconv.Atoi(v["parent_id"].(string))
//bean.BetAmount = v["betAmount"].(decimal.Decimal).String()
//bean.TotalAmount = v["totalAmount"].(decimal.Decimal).String()
//bean.TotalCommission = v["totalcommission"].(decimal.Decimal).String()
//bean.Created = v["created"].(int)
//bean.CreatedStr = v["created_str"].(string)
//bean.ProxyLevel = v["proxyLevel"].(int)
//bean.ProxyLevelName = v["proxyLevelName"].(string)
//bean.ProxyType, _ = strconv.Atoi(v["proxyType"].(string))
//bean.States = 0
//bean.Contributions = v["contributions"].(int)
//bean.ProxyLevelRate = decimal.NewFromFloat(v["rate"].(float64)).String()
//beans = append(beans, bean)
