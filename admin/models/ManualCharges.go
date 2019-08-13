package models

import (
	"qpgame/admin/common"
	"qpgame/app/fund"
	"qpgame/common/utils"
	"qpgame/config"
	"github.com/go-xorm/xorm"
	"regexp"
	"strconv"
	"time"
)

// 模型
type ManualCharges struct{}

// 表名称
func (self *ManualCharges) GetTableName(ctx *Context) string {
	return "manual_charges"
}

// 得到所有记录-分页
func (self *ManualCharges) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, user_id, `order`, amount, benefits, quantity, audit, item, deal_time, operator,comment, state",
		func(ctx *Context) []string { //获取查询条件
			queries := getQueryFields(ctx, &map[string]string{
				"order_id": "%",
				"audit":    "=",
				"state":    "=",
			})
			return append(queries, getQueryFieldByTime(ctx, "deal_time", "deal_start", "deal_end"))
		},
		func(ctx *Context, row *map[string]string) {
			processDatetime(&[]string{"deal_time"}, row) //时间戳转换成字符串
			(*row)["user_name"] = common.GetUserName((*ctx).Params().Get("platform"), (*row)["user_id"])
			processOptionsFor("state", "state_name", &map[string]string{
				"-1": "无需审核",
				"0":  "待审核",
				"1":  "审核通过",
				"2":  "作废",
			}, row)
			processOptionsFor("item", "item_name", &manualChargeItems, row)
		}, nil)
}

// 得到记录详情
func (self *ManualCharges) GetRecordDetail(ctx *Context) (map[string]string, error) {
	return getRecordDetail(ctx, self, "", nil)
}

// 添加记录
func (self *ManualCharges) Save(ctx *Context) (int64, error) {
	return saveRecord(ctx, self, nil,
		func(ctx *Context, data *map[string]string) bool { //添加之前处理
			post := utils.GetPostData(ctx)
			match, _ := regexp.Match("^[1-9][0-9]*$", []byte(post.Get("user_id")))
			if !match {
				return false
			}
			(*data)["state"] = "0"
			(*data)["order"] = utils.CreationOrder("WH", post.Get("user_id"))
			(*data)["deal_time"] = strconv.FormatInt(time.Now().Unix(), 10) //添加时间
			(*data)["operator"] = common.GetAdmin(ctx)["name"]
			return true
		}, nil, getSavedFunc("人工入款", "user_id"))
}

// 变更人工入款/出款相关信息
func changeManualState(ctx *Context, tableName string, manualType int, toState string) error {
	platform := utils.GetPlatform(ctx) //平台识别号
	idStr := (*ctx).URLParam("id")     //获取提交过来的充值记录id
	match, _ := regexp.Match("^[1-9][0-9]*$", []byte(idStr))
	if !match {
		return Error{What: "错误的记录编号"}
	}
	//获取此充值记录
	conn := utils.GetDbForPlatform(platform)
	sql := "SELECT * FROM " + tableName + " WHERE id=" + idStr
	rows, err := conn.SQL(sql).QueryString()
	if err != nil {
		return Error{What: "相关记录不存在"}
	}
	row := rows[0]
	//必须是未审核状态，才能变更为其他状态
	if row["state"] != "0" {
		return Error{What: "数据异常"}
	}
	uid := row["user_id"]
	// 操作人
	operator := common.GetAdmin(ctx)
	if toState == "1" {
		chargeAmount, _ := strconv.ParseFloat(row["amount"], 64) //充值金额-转化
		userId, _ := strconv.Atoi(uid)                //用户编号-转化
		orderId := row["order"]
		info := map[string]interface{}{
			"user_id":     userId,       //用户编号
			"type_id":     manualType,   //config.FUNDCHARGE, //类型: 充值
			"amount":      chargeAmount, //row["amount"],     //金额, 要传string类型
			"order_id":    orderId,      //订单编号
			"msg":         "后台通过审核", //操作说明
			"finish_rate": 1.0,          //打码倍率
		}
		// 回调
		var callback = func(db *xorm.Session, args ...interface{}) (interface{}, error) {
			// 更新记录状态
			sql = "UPDATE " + tableName + " SET state=1,operator='" + operator["name"] + "' WHERE id=" + idStr
			_, err := db.Exec(sql)
			if err != nil {
				return nil, Error{What: "修改数据失败"}
			}
			// 添加操作记录
			sqlLog := "INSERT INTO admin_logs (admin_id,admin_name,type,node,content,created)VALUES(" + operator["id"] + ",'" + operator["name"] + "','人工入款','人工入款审核','人工入款审核成功(用户id:" + uid + ")'," + strconv.Itoa(int(time.Now().Unix())) + ")"
			_,err = conn.Exec(sqlLog)
			if err != nil {
				return nil, Error{What: "人工入款充值审核操作写入日志失败"}
			}
			return nil, nil
		}
		result := fund.NewUserFundChange(platform).BalanceUpdate(info, callback)
		if result["status"] == 0 {
			return Error{What: result["msg"].(string)}
		}
		return nil
	}
	sql = "UPDATE " + tableName + " SET state=2,operator='" + operator["name"] + "' WHERE id=" + idStr
	_, err = conn.Exec(sql)
	if err != nil {
		return Error{What: "修改数据失败"}
	}
	// 添加操作记录
	sql = "INSERT INTO admin_logs (admin_id,admin_name,type,node,content,created)VALUES(" + operator["id"] + ",'" + operator["name"] + "','人工入款','人工入款审核','人工入款审核失败(用户id:" + uid + ")'," + strconv.Itoa(int(time.Now().Unix())) + ")"
	_, err = conn.Exec(sql)
	if err != nil {
		return Error{What: "人工入款充值审核操作写入日志失败"}
	}
	return nil
}

// 通过人工入款
func (self *ManualCharges) Allow(ctx *Context) error {
	return changeManualState(ctx, self.GetTableName(ctx), config.FUNDCHARGE, "1")
}

// 作废人工入款
func (self *ManualCharges) Deny(ctx *Context) error {
	return changeManualState(ctx, self.GetTableName(ctx), config.FUNDCHARGE, "2")
}

// 删除记录
func (self *ManualCharges) Delete(ctx *Context) error {
	return denyDelete()
}
