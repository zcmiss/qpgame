package models

import (
	"encoding/json"
	"github.com/shopspring/decimal"
	"qpgame/admin/common"
	"qpgame/app/fund"
	"qpgame/common/services"
	"qpgame/common/utils"
	"qpgame/config"
	"qpgame/models"
	xorm2 "qpgame/models/xorm"
	"qpgame/ramcache"
	"regexp"
	"strconv"
	"time"

	"github.com/go-xorm/xorm"
)

// 模型
type ChargeRecords struct{}

// 表名称
func (self *ChargeRecords) GetTableName(ctx *Context) string {
	return "charge_records"
}

// 得到所有记录-分页
func (self *ChargeRecords) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, user_id, amount, order_id, charge_type_id, card_number, bank_address, created, "+
			"ip, platform_id, real_name, bank_charge_time, credential_id, operator, is_tppay, "+
			"charge_card_id, remark, state",
		func(ctx *Context) []string { //获取查询条件
			queries := getQueryFields(ctx, &map[string]string{
				"order_id":       "%",
				"state":          "=",
				"charge_type_id": "=",
				"card_number":    "%",
				"credential_id":  "%",
				"is_tppay":       "=",
				"real_name":      "%",
				"ip":             "%",
			})
			queries = append(queries, getQueryFieldByTime(ctx, "created", "created_start", "created_end"))
			queries = append(queries, getQueryFieldByTime(ctx, "updated", "updated_start", "updated_end"))
			return append(queries, "`is_tppay`='0'") //公司入款
		},
		func(ctx *Context, row *map[string]string) {
			//processOptions("state", &statusTypes, row)
			processOptions("is_tppay", &yesNo, row)
			processDatetime(&[]string{"bank_charge_time", "created"}, row)
			platform := (*ctx).Params().Get("platform")
			(*row)["user_name"] = common.GetUserName(platform, (*row)["user_id"])
			(*row)["charge_type_name"] = common.GetChargeTypeName(platform, (*row)["charge_type_id"])
		},
		func(ctx *Context) (string, string, int) {
			return "", "id DESC", 0
		})
}

// 得到记录详情
func (self *ChargeRecords) GetRecordDetail(ctx *Context) (map[string]string, error) {
	return getRecordDetail(ctx, self, "", nil)
}

// 添加记录
func (self *ChargeRecords) Save(ctx *Context) (int64, error) {
	return saveRecord(ctx, self, nil,
		func(ctx *Context, data *map[string]string) bool { //添加之前处理
			(*data)["created"] = strconv.FormatInt(time.Now().Unix(), 10) //添加时间
			return true
		},
		func(ctx *Context, data *map[string]string) bool { //修改之前处理
			(*data)["updated"] = strconv.FormatInt(time.Now().Unix(), 10) //修改时间
			return true
		}, getSavedFunc("用户充值", "user_id"))
}

// 删除记录
func (self *ChargeRecords) Delete(ctx *Context) error {
	return denyDelete()
}

// 得到所有记录-分页
func (self *ChargeRecords) GetOnlines(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, user_id, amount, order_id, charge_type_id, card_number, bank_address, created, "+
			"ip, platform_id, real_name, bank_charge_time, credential_id, operator, is_tppay, "+
			"charge_card_id, remark, state",
		func(ctx *Context) []string { //获取查询条件
			queries := getQueryFields(ctx, &map[string]string{
				"order_id":       "%",
				"state":          "=",
				"charge_type_id": "=",
				"card_number":    "%",
				"credential_id":  "%",
				"is_tppay":       "=",
				"real_name":      "%",
				"ip":             "%",
			})
			queries = append(queries, getQueryFieldByTime(ctx, "created", "created_start", "created_end"))
			return append(queries, "`is_tppay`='1'") //线上入款
		},
		func(ctx *Context, row *map[string]string) {
			//processOptions("state", &statusTypes, row)
			processOptions("is_tppay", &yesNo, row)
			processDatetime(&[]string{"bank_charge_time", "created"}, row)
			platform := (*ctx).Params().Get("platform")
			(*row)["user_name"] = common.GetUserName(platform, (*row)["user_id"])
			(*row)["charge_type_name"] = common.GetChargeTypeName(platform, (*row)["charge_type_id"])
			(*row)["charge_card_name"] = common.GetChargeCardName(platform, (*row)["charge_card_id"])
			(*row)["credential_name"] = common.GetThirdPaymentName(platform, (*row)["credential_id"])
		},
		func(ctx *Context) (string, string, int) {
			return "", "id DESC", 0
		})
}

// 活动奖励
func (self *ChargeRecords) ActivityAward(platform string, session *xorm.Session, atype int, iUserId int, ipAddr string) (string, error) {
	if (iUserId > 0) && ((atype == 1) || (atype == 2)) {
		var activity = new(xorm2.Activities)
		iNow := utils.GetNowTime()
		actExist, err := session.Where("`status`=1 AND `type`=? AND time_start<=? AND time_end>?", atype, iNow, iNow).Get(activity)
		if (err == nil) && actExist {
			tip := "充值奖励"
			if atype == 1 {
				tip = "注册奖励"
			}
			decimal.DivisionPrecision = 2
			sMoney := activity.Money
			var dMoney decimal.Decimal
			dMoney, _ = decimal.NewFromString(sMoney)
			fMoney, toFloat := dMoney.Float64()
			if toFloat == false {
				return "活动奖励金额设置错误", Error{What: tip + "赠送失败"}
			}
			if dMoney.GreaterThan(decimal.New(0, 0)) == false {
				return "", nil
			}

			var actRecords []xorm2.ActivityRecords
			session.Cols("id", "created", "ip_addr").Where("state=1 AND user_id=? AND activity_id=?", iUserId, activity.Id).Find(&actRecords)
			iActRecordNum := len(actRecords)
			var actRecord = xorm2.ActivityRecords{
				UserId:  iUserId,
				State:   1,
				Applied: iNow,
				Created: iNow,
				Updated: iNow,
				IpAddr:  ipAddr,
			}
			if iActRecordNum == 0 {
				return services.RecordActivity(platform, session, actRecord, fMoney, tip)
			}

			iIsRepeat := activity.IsRepeat
			bIsRepeat := iIsRepeat == 1
			// 是否允许重复，如果不允许重复，查看record的记录条数，如没有，则记录，否则，返回
			if bIsRepeat == false {
				return "", nil
			}
			// 计算当日的起始时间戳
			iFromTime, iToTime := utils.GetDatetimeRange(0, 1)
			iIpTotalCnt := 0 // 通过Ip统计参与指定活动的总数
			iIpDayCnt := 0

			for _, actRecordBean := range actRecords {
				iIpTotalCnt++
				iCreated := int64(actRecordBean.Created)
				if iFromTime < iCreated && iCreated <= iToTime {
					iIpDayCnt++
				}
			}
			iTotalIpLimit := activity.TotalIpLimit
			iDayIpLimit := activity.DayIpLimit
			// 判断当前Ip领取某一活动的总次数是否超出限制
			if iIpTotalCnt >= iTotalIpLimit {
				return "", nil
			}
			// 判断当前Ip当日领取某一活动是否超出限制
			if iIpDayCnt >= iDayIpLimit {
				return "", nil
			}
			return services.RecordActivity(platform, session, actRecord, fMoney, tip)
		}
	}
	return "", nil
}

// 通过审核
func (self *ChargeRecords) Allow(ctx *Context) error {
	platform, id := (*ctx).Params().Get("platform"), (*ctx).URLParam("id")
	//获取此充值记录
	tableName, db := self.GetTableName(ctx), models.MyEngine[platform]
	rows, err := db.SQL("SELECT user_id, amount, state, is_tppay FROM " + tableName + " WHERE id=" + id).QueryString()
	if (err != nil) || (len(rows) < 1) {
		return Error{What: "充值记录不存在"}
	}
	record := rows[0]
	if (record["is_tppay"] != "0") || (record["state"] != "0") {
		return Error{What: "充值记录不存在"}
	}
	//充值金额-转化
	chargeAmount, err := strconv.ParseFloat(record["amount"], 64)
	if err != nil {
		return Error{What: "充值记录金额数据错误"}
	}
	// 用户编号
	userIdStr := record["user_id"]
	userId, _ := strconv.Atoi(userIdStr)
	if userId <= 0 {
		return Error{What: "充值记录用户编号数据错误"}
	}
	orderId := utils.CreationOrder("CH", record["user_id"])

	// 充值流水类型
	chargeTypeId, chargeTypeIdStr := config.FUNDCHARGE, strconv.Itoa(config.FUNDCHARGE)
	// 充值赠送流水类型
	presentedTypeId, presentedTypeIdStr := config.FUNDPRESENTER, strconv.Itoa(config.FUNDPRESENTER)

	// 赠送彩金比例
	presentedRate := float64(0)
	// 充值打码量比率
	chargeDamaRate := float64(0)
	// 赠送彩金打码量比率
	presentedDamaRate := float64(0)
	//
	rows, err = db.SQL("SELECT name,value FROM configs WHERE name IN('fund_dama_rate','com_bank_present_rate') ORDER BY name DESC").QueryString()
	if (err != nil) || (len(rows) != 2) {
		return Error{What: "充值配置不存在"}
	}
	presentedRate, _ = strconv.ParseFloat(rows[1]["value"], 64)
	//
	config := make(map[string]interface{})
	json.Unmarshal([]byte(rows[0]["value"]), &config)
	if _, ok := config[chargeTypeIdStr]; ok {
		x := config[chargeTypeIdStr].(map[string]interface{})
		if _, ok := x["dama_rate"]; ok && (x["dama_rate"] != nil) {
			chargeDamaRate = x["dama_rate"].(float64)
		}
		x = config[presentedTypeIdStr].(map[string]interface{})
		if _, ok := x["dama_rate"]; ok && (x["dama_rate"] != nil) {
			presentedDamaRate = x["dama_rate"].(float64)
		}
	}
	presentedAmount := presentedRate * chargeAmount
	// 操作者
	admin := common.GetAdmin(ctx)
	operator, operatorId := admin["name"], admin["id"]
	// 充值部分
	info := map[string]interface{}{
		"user_id":     userId,         //用户编号
		"type_id":     chargeTypeId,   //类型: 充值
		"amount":      chargeAmount,   //row["amount"],     //金额, 要传string类型
		"order_id":    orderId,        //订单编号
		"msg":         "审核通过",         //操作说明
		"finish_rate": chargeDamaRate, //打码倍率
	}
	// 充值回调
	var callback fund.BalanceUpdateCallback = func(db *xorm.Session, args ...interface{}) (interface{}, error) {
		// 更新充值记录状态
		_, err = db.Exec("UPDATE " + tableName + " SET state=1,operator='" + operator + "' WHERE id=" + id)
		if err != nil {
			return nil, Error{What: "审核入款充值操作失败"}
		}
		// 添加充值操作记录
		sqlLog := "INSERT INTO admin_logs (`admin_id`,`admin_name`,`type`,`node`,`content`,`created`)VALUES(" + operatorId + ", '" + operator + "', '审核入款','会员公司审核入款','会员公司审核入款审核通过(用户id:" + userIdStr + ")', " + strconv.Itoa(int(time.Now().Unix())) + ")"
		_, err = db.Exec(sqlLog)
		if err != nil {
			return nil, Error{What: "审核入款充值操作写入日志失败"}
		}
		// 赠送彩金部分
		presentedInfo := map[string]interface{}{
			"user_id":     userId,            //用户编号
			"type_id":     presentedTypeId,   //类型: 赠送彩金,
			"amount":      presentedAmount,   //row["amount"],     //金额, 要传string类型
			"order_id":    orderId,           //订单编号
			"msg":         "彩金赠送",            //操作说明
			"finish_rate": presentedDamaRate, //打码倍率
			"transaction": db,                //事务
		}
		result := fund.NewUserFundChange(platform).BalanceUpdate(presentedInfo,
			func(pdb *xorm.Session, pargs ...interface{}) (i interface{}, e error) {
				// 添加操作记录
				sqlLog := "INSERT INTO notices (user_id,content,`status`,created,title)VALUES(" + strconv.Itoa(userId) + ",'充值赠送礼金,已到账',1," + strconv.Itoa(int(time.Now().Unix())) + ",'充值赠送彩金" + strconv.FormatFloat(presentedAmount, 'f', 3, 64) + "')"
				_, err = db.Exec(sqlLog)
				if err != nil {
					return nil, Error{What: "彩金赠写入站内通知失败"}
				}
				return nil, nil
			})

		if result["status"] == 0 {
			return nil, Error{What: "彩金赠送失败操作失败"}
		}
		// 活动奖励
		sIp := utils.GetIp((*ctx).Request())
		_, err := self.ActivityAward(platform, db, 2, userId, sIp)
		if err != nil {
			return nil, Error{What: err.Error()}
		}
		return nil, nil
	}
	result := fund.NewUserFundChange(platform).BalanceUpdate(info, callback)
	if result["status"] == 0 {
		return Error{What: "审核入款充值操作失败"}
	}
	return nil
}

// 拒绝公司入款
func (self *ChargeRecords) Deny(ctx *Context) error {
	platform := (*ctx).Params().Get("platform")
	//获取提交过来的充值记录id
	idStr := (*ctx).URLParam("id")
	//获取操作者
	operator := common.GetAdmin(ctx)
	match, _ := regexp.Match("^[1-9][0-9]*$", []byte(idStr))
	if !match {
		return Error{What: "错误的充值记录编号"}
	}
	//获取此充值记录
	conn := models.MyEngine[platform]
	transaction := conn.NewSession()
	defer transaction.Close()
	tableName := self.GetTableName(ctx)
	sql := "SELECT state,is_tppay,user_id FROM " + tableName + " WHERE id=" + idStr
	rows, err := conn.SQL(sql).QueryString()
	if err != nil {
		transaction.Rollback()
		return Error{What: "充值记录异常"}
	}
	row := rows[0]
	if row["is_tppay"] != "0" || row["state"] != "0" {
		transaction.Rollback()
		return Error{What: "充值记录异常"}
	}
	sql = "UPDATE " + tableName + " SET state=2,operator='" + operator["name"] + "' WHERE id=" + idStr
	result, err2 := transaction.Exec(sql)
	if err2 != nil {
		transaction.Rollback()
		return Error{What: "订单操作失败"}
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		transaction.Rollback()
		return Error{What: "拒绝充值操作失败"}
	}
	sql = "INSERT INTO admin_logs(admin_id,admin_name,type,node,content,created)VALUES(" + operator["id"] + ",'" + operator["name"] + "','拒绝入款','会员公司入款','会员公司入款强制入款(用户id:" + row["user_id"] + ")'," + strconv.FormatInt(time.Now().Unix(), 10) + ")"
	_, err = transaction.SQL(sql).QueryString()
	if err != nil {
		return Error{What: "日志写入失败"}
	}
	if transaction.Commit() != nil {
		transaction.Rollback()
		return Error{What: "充值操作处理失败"}
	}
	return nil
}

//强制入款
func (self *ChargeRecords) ForcedDeposit(ctx *Context) error {
	platform, id := (*ctx).Params().Get("platform"), (*ctx).URLParam("id")
	//获取此充值记录
	tableName, db := self.GetTableName(ctx), models.MyEngine[platform]
	rows, err := db.SQL("SELECT user_id, amount, state, is_tppay FROM " + tableName + " WHERE id=" + id).QueryString()
	if (err != nil) || (len(rows) == 0) {
		return Error{What: "充值记录不存在"}
	}
	record := rows[0]
	if (record["is_tppay"] == "0") || (record["state"] == "6") {
		return Error{What: "充值记录不可强制入款操作"}
	}
	//充值金额-转化
	chargeAmount, err := strconv.ParseFloat(record["amount"], 64)
	if err != nil {
		return Error{What: "充值记录金额数据错误"}
	}
	// 用户编号
	userIdStr := record["user_id"]
	userId, _ := strconv.Atoi(userIdStr)
	if userId <= 0 {
		return Error{What: "充值记录用户编号数据错误"}
	}
	orderId := utils.CreationOrder("CH", record["user_id"])
	// 资金操作类型
	typeId := config.FUNDCHARGE
	typeIdStr := strconv.Itoa(typeId)
	// 从缓存里读取打码量比例
	fundDamaRate := float64(0)
	cnf, ok := ramcache.TableConfigs.Load(platform)
	if ok {
		tmp := cnf.(map[string]interface{})["fund_dama_rate"].(map[string]interface{})
		if tmp[typeIdStr] != nil {
			fundDamaRate = tmp[typeIdStr].(map[string]interface{})["dama_rate"].(float64)
		}
	}
	info := map[string]interface{}{
		"user_id":     userId,       //用户编号
		"type_id":     typeId,       //类型: 充值
		"amount":      chargeAmount, //row["amount"],     //金额, 要传string类型
		"order_id":    orderId,      //订单编号
		"msg":         "后台强制入款",     //操作说明
		"finish_rate": fundDamaRate, //打码倍率
	}
	// 操作者
	admin := common.GetAdmin(ctx)
	operator, operatorId := admin["name"], admin["id"]
	// 回调
	var callback fund.BalanceUpdateCallback = func(db *xorm.Session, args ...interface{}) (interface{}, error) {
		// 更新记录状态
		_, err = db.Exec("UPDATE " + tableName + " SET state=1,operator='" + operator + "' WHERE id=" + id)
		if err != nil { //判断是否执行成功
			return nil, Error{What: "强制入款充值操作失败"}
		}
		// 添加操作记录
		sqlLog := "INSERT INTO admin_logs (`admin_id`,`admin_name`,`type`,`node`,`content`,`created`)VALUES(" + operatorId + ", '" + operator + "', '强制入款','会员线上入款','会员线上入款强制入款(用户id:" + userIdStr + ")', " + strconv.Itoa(int(time.Now().Unix())) + ")"
		_, err = db.Exec(sqlLog)
		if err != nil {
			return nil, Error{What: "强制入款充值操作写入日志失败"}
		}
		// 活动奖励
		sIp := utils.GetIp((*ctx).Request())
		_, err := self.ActivityAward(platform, db, 2, userId, sIp)
		if err != nil {
			return nil, Error{What: err.Error()}
		}
		return nil, nil
	}
	result := fund.NewUserFundChange(platform).BalanceUpdate(info, callback)
	if result["status"] == 0 { //如果修改失败
		return Error{What: "强制入款充值操作失败"}
	}
	return nil
}
