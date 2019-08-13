package models

import (
	"encoding/json"
	"regexp"
	"strconv"
	"strings"

	"qpgame/admin/validations"
	"qpgame/common/utils"
	"qpgame/models"
	"qpgame/models/xorm"
)

// 模型
type Configs struct{}

//配置相关的校验器
var configValidation = validations.ConfigsValidation{}

//设置相关的字段
var setsFields = []string{
	"withdraw_min_money",      //提现最低金额
	"withdraw_max_money",      //提现最大金额
	"clear_dm_tx_limit",       //清除提现打码限制阙值
	"withdraw_day_limited",    //每天提款次数限制
	"register_number_ip",      //单个ip最多允许注册人数
	"sign_reward",             //签到奖励法：
	"tuiguang_web_url",        //代理推广地址后缀
	"tuiguang_web_domain",     //推广下载地址
	"generalize_award",        //推广分享奖励金额
	"reward_bind",             //绑定手机奖励金额
	"sign_award_switch",       //绑定手机奖励金额
	"bind_phone_award_switch", //是否开启绑定手机奖励
	"report_qq",               //举报QQ号
	//"register_config", 	//注册配置
	//"proxy_charge", 		//代理充值信息
}

// 表名称
func (self *Configs) GetTableName(ctx *Context) string {
	return "configs"
}

// 得到所有记录-分页
func (self *Configs) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, name, value, mark, updated",
		func(ctx *Context) []string { //获取查询条件
			var conditions []string                              //查询条件数组
			if id, err := (*ctx).URLParamInt("id"); err == nil { //按编号查询
				conditions = append(conditions, "id = "+strconv.Itoa(id))
			}
			return conditions
		},
		func(ctx *Context, row *map[string]string) {
			processDatetime(&[]string{"updated"}, row)
		}, nil)
}

// 得到记录详情
func (self *Configs) GetRecordDetail(ctx *Context) (map[string]string, error) {
	return getRecordDetail(ctx, self, "", nil)
}

// 添加记录
func (self *Configs) Save(ctx *Context) (int64, error) {
	return denySave()
}

// 删除记录
func (self *Configs) Delete(ctx *Context) error {
	return denyDelete()
}

// 获取总的邀请奖励金额
func (self *Configs) GetInviteAmount(ctx *Context) float64 {
	platform := (*ctx).Params().Get("platform")
	engine := models.MyEngine[platform]
	result := float64(0)
	rows, _ := engine.SQL("SELECT SUM(amount) amount FROM account_infos WHERE type=19").QueryString()
	if len(rows) > 0 {
		result, _ = strconv.ParseFloat(rows[0]["amount"], 64)
	}
	return result
}

// 系统-设置
func (self *Configs) Sets(ctx *Context) error {
	message, result := configValidation.ValidateSets(ctx)
	if !result {
		return Error{What: message}
	}

	platform := (*ctx).Params().Get("platform")
	data := utils.GetPostData(ctx)
	conn := models.MyEngine[platform]
	for _, field := range setsFields {
		sql := "UPDATE configs SET value = '" + data.Get(field) + "' WHERE name = '" + field + "' LIMIT 1"
		result, err := conn.Exec(sql)
		if err != nil {
			return Error{What: "保存配置信息有误"}
		}
		_, err = result.RowsAffected()
		if err != nil {
			return Error{What: "保存配置信息有误"}
		}
	}
	return nil
}

//系统设置-获取
func (self *Configs) GetSets(ctx *Context) map[string]string {
	fields := strings.Join(setsFields, "','")
	platform := (*ctx).Params().Get("platform")
	conn := models.MyEngine[platform]
	sql := "SELECT name, value FROM configs WHERE name IN ('" + fields + "')"
	rows, err := conn.SQL(sql).QueryString()
	if err != nil || len(rows) == 0 {
		return map[string]string{}
	}

	data := map[string]string{}
	for _, row := range rows {
		data[row["name"]] = row["value"]
	}

	return data
}

// 打码量比率-设置
func (self *Configs) FundDamaRateSet(ctx *Context) error {
	conn := models.MyEngine[(*ctx).Params().Get("platform")]
	post := utils.GetPostData(ctx)
	tmp := map[string]string{
		"1":    post.Get("1"),
		"1000": post.Get("1000"),
		"16":   post.Get("16"),
		"17":   post.Get("17"),
		"19":   post.Get("19"),
		"18":   post.Get("18"),
		"14":   post.Get("14"),
		"3":    post.Get("3"),
		"5":    post.Get("5"),
		"6":    post.Get("6"),
		"8":    post.Get("8"),
		"9":    post.Get("9"),
	}

	data := make(map[string]map[string]interface{})
	for k, v := range tmp {
		item := make(map[string]interface{})
		err := json.Unmarshal([]byte(v), &item)
		if err != nil {
			return err
		}
		dama_rate, _ := strconv.ParseFloat(item["dama_rate"].(string), 64)
		item["dama_rate"] = dama_rate
		data[k] = item
	}

	bs, err := json.Marshal(data)
	if err != nil {
		return err
	}

	sql := "UPDATE configs SET value = '" + string(bs) + "' WHERE name =  'fund_dama_rate'"
	_, err = conn.Exec(sql)
	if err != nil {
		return Error{What: "不能保存打码量比率"}
	}
	return nil
}

//打码量比率获取
func (self *Configs) FundDamaRate(ctx *Context) map[string]interface{} {
	conn := models.MyEngine[(*ctx).Params().Get("platform")]
	data := map[string]interface{}{}
	sql := "SELECT value FROM configs WHERE name = 'fund_dama_rate' "
	rows, err := conn.SQL(sql).QueryString()
	if err != nil || len(rows) <= 0 {
		return data
	}
	row := rows[0]
	json.Unmarshal([]byte(row["value"]), &data)
	if err != nil {
		return data
	}
	return data
}

//在线客服-设置
func (self *Configs) Service(ctx *Context) error {
	post := utils.GetPostData(ctx)

	data := ServiceAccount{
		Wx: []ServiceInfo{},
		Qq: []ServiceInfo{},
	}
	err := json.Unmarshal(*post.Data, &data)
	if err != nil {
		return Error{What: "保存在线客服信息失败"}
	}
	qqBytes, qqErr := json.Marshal(data.Qq)
	wxBytes, wxErr := json.Marshal(data.Wx)
	if qqErr != nil || wxErr != nil {
		return Error{What: "解析客服信息失败"}
	}
	conn := models.MyEngine[(*ctx).Params().Get("platform")]
	sql := "UPDATE configs SET value = '" + string(qqBytes) + "' WHERE name = 'qq_customer' LIMIT 1" //qq客服
	_, err = conn.Exec(sql)
	if err != nil {
		return Error{What: "不能保存QQ客服信息"}
	}
	sql = "UPDATE configs SET value = '" + string(wxBytes) + "' WHERE name = 'weixin_customer' LIMIT 1" //微信客服
	_, err = conn.Exec(sql)
	if err != nil {
		return Error{What: "不能保存微信客服信息"}
	}
	row := post.Get("web_customer_url")
	if len(row) <= 0 {
		row = ""
	}
	sql = "UPDATE configs SET value = '" + row + "' WHERE name = 'web_customer_url' LIMIT 1" //在线客服
	_, err = conn.Exec(sql)
	if err != nil {
		return Error{What: "不能保存在线客服信息"}
	}
	return nil
}

//在线客服-获取
func (self *Configs) GetService(ctx *Context) map[string]interface{} {
	var data = make(map[string]interface{})

	conn := models.MyEngine[(*ctx).Params().Get("platform")]
	//sql := "SELECT value FROM configs WHERE name = 'qq_customer' LIMIT 1"
	//rows, err := conn.SQL(sql).QueryString()
	//if err != nil || len(rows) <= 0 {
	//	return data
	//}
	//
	//infoQq := []ServiceInfo{}
	//row := rows[0]
	//json.Unmarshal([]byte(row["value"]), &infoQq)
	//
	//sql = "SELECT value FROM configs WHERE name = 'weixin_customer' LIMIT 1"
	//rows, err = conn.SQL(sql).QueryString()
	//if err != nil || len(rows) <= 0 {
	//	return data
	//}
	//row = rows[0]
	//
	//infoWx := []ServiceInfo{}
	//json.Unmarshal([]byte(row["value"]), &infoWx)
	//
	//
	//sql = "SELECT value FROM configs WHERE name = 'web_customer_url'"
	//var configBean xorm.Configs
	//exist, err := conn.SQL(sql).Get(&configBean)
	//if err != nil || exist == false {
	//	return nil
	//}

	var configBeans []xorm.Configs
	err := conn.Where("`name` IN ('qq_customer', 'weixin_customer', 'web_customer_url')").Cols("`value`,`name`").Find(&configBeans)

	if err != nil {
		return nil
	}

	for _, configBean := range configBeans {
		if configBean.Name == "qq_customer" {
			jsonStr := configBean.Value
			var result []ServiceInfo
			json.Unmarshal([]byte(jsonStr), &result)
			data["qq"] = result
		} else if configBean.Name == "weixin_customer" {
			jsonStr := configBean.Value
			var result []ServiceInfo
			json.Unmarshal([]byte(jsonStr), &result)
			data["wx"] = result
		} else if configBean.Name == "web_customer_url" {
			data["web_customer_url"] = configBean.Value
		}
	}
	//data["qq"] = infoQq
	//data["wx"] = infoWx
	return data
}

//常见问题-设置
func (self *Configs) Faq(ctx *Context) error {
	data := utils.GetPostData(ctx)
	faq := data.Get("faq")
	if faq == "" {
		return Error{What: "必须输入FAQ内容"}
	}
	conn := models.MyEngine[(*ctx).Params().Get("platform")]
	sql := "UPDATE configs SET value = '" + faq + "' WHERE name = 'faq' LIMIT 1"
	result, err := conn.Exec(sql)
	if err != nil {
		return Error{What: "保存FAQ内容错误"}
	}

	affected, ok := result.RowsAffected()
	if ok != nil || affected <= 0 {
		return Error{What: "保存FAQ操作失败"}
	}
	return nil
}

//常见问题-获取
func (self *Configs) GetFaq(ctx *Context) map[string]string {
	data := map[string]string{
		"faq": "",
	}
	conn := models.MyEngine[(*ctx).Params().Get("platform")]
	sql := "SELECT value FROM configs WHERE name ='faq' LIMIT 1"
	rows, errRows := conn.SQL(sql).QueryString()
	if errRows != nil || len(rows) == 0 {
		return data
	}
	row := rows[0]
	data["faq"] = row["value"]
	return data
}

//充值配置-设置
func (self *Configs) Charges(ctx *Context) error {
	data := ProxyCharge{
		ProxyChargeLogo: "",
		ChargeAccounts:  []ChargeInfo{},
		Info:            "",
		State:           0,
	}
	post := utils.GetPostData(ctx)
	err := json.Unmarshal(*post.Data, &data)
	if err != nil {
		return Error{What: "提交的参数格式不正确"}
	}
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return Error{What: "解析提交的参数失败"}
	}
	sql := "UPDATE configs SET value = '" + string(jsonBytes) + "' WHERE name = 'proxy_charge' LIMIT 1"
	conn := models.MyEngine[(*ctx).Params().Get("platform")]
	result, err := conn.Exec(sql)
	if err != nil {
		return Error{What: "修改代理充值信息失败"}
	}
	affected, ok := result.RowsAffected()
	if ok != nil || affected <= 0 {
		return Error{What: "保存代理充值信息失败"}
	}
	return nil
}

//充值配置-获取
func (self *Configs) GetCharges(ctx *Context) interface{} {
	data := ProxyCharge{
		ProxyChargeLogo: "",
		ChargeAccounts:  []ChargeInfo{},
		Info:            "",
		State:           0,
	}
	conn := models.MyEngine[(*ctx).Params().Get("platform")]
	sql := "SELECT value FROM configs WHERE name = 'proxy_charge' LIMIT 1"
	rows, err := conn.SQL(sql).QueryString()
	if err != nil || len(rows) == 0 {
		return data
	}

	row := rows[0]
	json.Unmarshal([]byte(row["value"]), &data)
	return data
}

//注册配置-设置
func (self *Configs) Reg(ctx *Context) error {
	post := utils.GetPostData(ctx)
	canRegister := post.Get("can_register")
	can, err := strconv.Atoi(canRegister)
	if err != nil {
		return Error{What: "注册配置提交内容有误"}
	}
	conf := RegisterConf{
		CanRegister: can,
	}
	jsonBytes, err := json.Marshal(conf)
	if err != nil {
		return Error{What: "格式化提交的注册配置相关信息失败"}
	}
	sql := "UPDATE configs SET value = '" + string(jsonBytes) + "' WHERE name = 'register_config' LIMIT 1"
	conn := models.MyEngine[(*ctx).Params().Get("platform")]
	result, resErr := conn.Exec(sql)
	if resErr != nil {
		return Error{What: "保存注册配置相关信息失败"}
	}

	affected, ok := result.RowsAffected()
	if ok != nil || affected <= 0 {
		return Error{What: "保存注册配置失败"}
	}
	return nil
}

//注册配置-获取
func (self *Configs) GetReg(ctx *Context) RegisterConf {
	data := RegisterConf{
		CanRegister: 0,
	}
	conn := models.MyEngine[(*ctx).Params().Get("platform")]
	sql := "SELECT value FROM configs WHERE name ='register_config' LIMIT 1"
	rows, errRows := conn.SQL(sql).QueryString()
	if errRows != nil || len(rows) == 0 {
		return data
	}
	row := rows[0]

	json.Unmarshal([]byte(row["value"]), &data)
	return data
}

//订单提醒-获取
func (self *Configs) GetOrderAlert(ctx *Context) OrderAlert {
	data := OrderAlert{}
	conn := models.MyEngine[(*ctx).Params().Get("platform")]
	sql := "SELECT value FROM configs WHERE name ='order_alert' LIMIT 1"
	rows, errRows := conn.SQL(sql).QueryString()
	if errRows != nil || len(rows) == 0 {
		return data
	}
	row := rows[0]

	json.Unmarshal([]byte(row["value"]), &data)
	return data
}

//公司入款银行卡转账赠送比例配置-设置
func (self *Configs) BankPresentRateSet(ctx *Context) error {
	post := utils.GetPostData(ctx)
	conn := models.MyEngine[(*ctx).Params().Get("platform")]
	row := post.Get("com_bank_present_rate")
	if match, _ := regexp.MatchString(`^0(\.[0-9]+)?$`, row); !match {
		return Error{What: "输入的数据不合法"}
	}
	sql := "UPDATE configs SET value=" + row + " WHERE name='com_bank_present_rate'"
	_, err := conn.Exec(sql)
	if err != nil {
		return Error{What: "不能保存赠送彩金比例"}
	}
	return nil
}

//公司入款银行卡转账赠送比例配置-获取
func (self *Configs) BankPresentRate(ctx *Context) map[string]string {
	data := make(map[string]string)
	conn := models.MyEngine[(*ctx).Params().Get("platform")]
	sql := "SELECT value com_bank_present_rate FROM configs WHERE name='com_bank_present_rate'"
	rows, err := conn.SQL(sql).QueryString()
	if err != nil || len(rows) <= 0 {
		return data
	}
	return rows[0]
}

// 银商配置-设置
func (self *Configs) SilverMerchantSet(ctx *Context) error {
	data := map[string]interface{}{
		"cash_pledge":        0,
		"min_charge_money":   0,
		"min_transfer_money": 0,
	}
	post := utils.GetPostData(ctx)
	err := json.Unmarshal(*post.Data, &data)
	if err != nil {
		return Error{What: "提交的参数格式不正确"}
	}
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return Error{What: "解析提交的参数失败"}
	}
	sql := "UPDATE configs SET value = '" + string(jsonBytes) + "' WHERE name = 'silver_merchant' LIMIT 1"
	conn := models.MyEngine[(*ctx).Params().Get("platform")]
	result, err := conn.Exec(sql)
	if err != nil {
		return Error{What: "修改银商配置信息失败"}
	}
	affected, ok := result.RowsAffected()
	if ok != nil || affected <= 0 {
		return Error{What: "保存银商配置信息失败"}
	}
	return nil
}

// 银商配置-获取
func (self *Configs) GetSilverMerchant(ctx *Context) map[string]float64 {
	data := map[string]float64{
		"cash_pledge":        0,
		"min_charge_money":   0,
		"min_transfer_money": 0,
	}
	conn := models.MyEngine[(*ctx).Params().Get("platform")]
	sql := "SELECT value FROM configs WHERE name = 'silver_merchant' LIMIT 1"
	rows, err := conn.SQL(sql).QueryString()
	if err != nil || len(rows) == 0 {
		if err == nil {
			sql = "INSERT INTO configs(name,value)VALUES('silver_merchant', '{\"cash_pledge\":0,\"min_charge_money\":0,\"min_transfer_money\":0}')"
			conn.Exec(sql)
		}
		return data
	}
	row := rows[0]
	json.Unmarshal([]byte(row["value"]), &data)
	return data

}
