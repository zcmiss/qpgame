package models

import (
	"qpgame/admin/common"
	"qpgame/admin/validations"
	"qpgame/common/utils"
	db "qpgame/models"
	"qpgame/models/xorm"
	"regexp"
	"strconv"
)

// 模型
type Users struct{}

var usersValidation = validations.UsersValidation{}

// 表名称
func (self *Users) GetTableName(ctx *Context) string {
	return "users"
}

// 得到所有记录-分页
func (self *Users) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, password, user_name, name, email, created, birthday, mobile_type, "+
			"path, vip_level, qq, wechat, status, "+ //    is_dummy, " +
			"token, safe_password, parent_id, last_login_time, last_platform_id",
		func(ctx *Context) []string { //获取查询条件
			queries := getQueryFields(ctx, &map[string]string{
				"parent_id":   "=",
				"name":        "%",
				"user_name":   "%",
				"qq":          "%",
				"wechat":      "%",
				"vip_level":   "=",
				"status":      "=",
				"mobile_type": "=",
			})
			queries = append(queries, getQueryFieldByTime(ctx, "created", "created_start", "created_end"))
			return append(queries, getQueryFieldByTime(ctx, "last_login_time", "last_login_start", "last_login_end"))
		},
		func(ctx *Context, row *map[string]string) {
			processDatetime(&[]string{"last_login_time", "created"}, row) //添加时间/最后登录时间
			processOptions("mobile_type", &appTypes, row)                 //手机型号
			platform := (*ctx).Params().Get("platform")
			(*row)["parent_name"] = common.GetUserName(platform, (*row)["parent_id"])
			(*row)["last_platform_name"] = common.GetGamePlatformName(platform, (*row)["last_platform_id"])
		}, nil)
}

// 得到所有邀请用户记录-分页
func (self *Users) GetInviteRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, user_name, name, created, status, parent_id",
		func(ctx *Context) []string { //获取查询条件
			queries := getQueryFields(ctx, &map[string]string{
				"parent_id": "=",
				"user_id":   "=",
				"name":      "%",
				"user_name": "%",
				"status":    "=",
			})
			queries = append(queries, getQueryFieldByTime(ctx, "created", "created_start", "created_end"))
			return queries
		},
		func(ctx *Context, row *map[string]string) {
			processDatetime(&[]string{"created"}, row)
			userId := (*row)["id"]
			platform := (*ctx).Params().Get("platform")
			engine := db.MyEngine[platform]
			(*row)["parent_name"] = common.GetUserName(platform, (*row)["parent_id"])
			inviteAmount, inviteCount := "0.00", "0"
			rows, _ := engine.SQL("SELECT COUNT(*) count FROM users WHERE parent_id=" + userId).QueryString()
			if (len(rows) > 0) && (rows[0]["count"] != "") {
				inviteCount = rows[0]["count"]
			}
			rows, _ = engine.SQL("SELECT SUM(amount) amount FROM account_infos WHERE type=19 AND user_id=" + userId).QueryString()
			if (len(rows) > 0) && (rows[0]["amount"] != "") {
				inviteAmount = rows[0]["amount"]
			}
			(*row)["invite_amount"] = inviteAmount
			(*row)["invite_count"] = inviteCount
		}, nil)
}

// 得到记录详情
func (self *Users) GetRecordDetail(ctx *Context) (map[string]string, error) {
	return getRecordDetail(ctx, self, "", nil)
}

// 添加记录
func (self *Users) Save(ctx *Context) (int64, error) {
	userId, isAdd := 0, false
	id, err := saveRecord(ctx, self, nil,
		func(ctx *Context, data *map[string]string) bool { //添加之前处理
			isAdd = true
			d := *data
			c := *ctx
			sIp := utils.GetIp(c.Request())
			platform := c.Params().Get("platform")
			engine := db.MyEngine[platform]
			rows, err := engine.SQL("SELECT value FROM configs WHERE name='register_number_ip'").QueryString()
			if (err != nil) || (len(rows) < 1) {
				return false
			}
			iRegNumIp, _ := strconv.Atoi(rows[0]["value"])
			iFromTime, iToTime := utils.GetDatetimeRange(0, 1)
			var userBean xorm.Users
			iIpRegTotal, _ := engine.Where("reg_ip=? and created between ? and ?", sIp, iFromTime, iToTime).Count(&userBean)
			if int(iIpRegTotal) >= iRegNumIp {
				return false
			}
			username := d["user_name"]
			match, _ := regexp.MatchString("(^[a-zA-Z][a-zA-Z0-9_]{4,15}$)", username)
			if !match {
				return false
			}
			password := d["password"]
			safePassword := d["safe_password"]
			loginfrom := d["mobile_type"]
			parentid := d["parent_id"]
			var user = xorm.DefaultUser()
			user.Birthday = d["birthday"]
			user.Sex, _ = strconv.Atoi(d["sex"])
			user.Email = d["email"]
			user.Phone = d["phone"]
			user.Qq = d["qq"]
			user.Wechat = d["wechat"]
			user.UserType, _ = strconv.Atoi(d["user_type"])
			user.UserName = username
			user.Password = utils.MD5(password)
			user.SafePassword = utils.MD5(safePassword)
			user.ParentId, _ = strconv.Atoi(parentid)
			if loginfrom == "2" {
				user.MobileType = 2
			} else {
				user.MobileType = 1
			}
			session := engine.NewSession()
			defer session.Close()
			err = session.Begin()
			_, err = session.Insert(&user)
			if err != nil {
				session.Rollback()
				return false
			}
			_, err = session.Insert(xorm.Accounts{UserId: user.Id, Updated: utils.GetNowTime()})
			if err != nil {
				session.Rollback()
				return false
			}
			// 活动奖励
			cr := ChargeRecords{}
			_, err = (&cr).ActivityAward(platform, session, 1, userId, sIp)
			if err != nil {
				session.Rollback()
				return false
			}
			err = session.Commit()
			userId = user.Id
			return false
		}, nil, getSavedFunc("用户信息", "user_name"))
	if isAdd {
		err = nil
		id = int64(userId)
	}
	return id, err
}

// 删除记录
func (self *Users) Delete(ctx *Context) error {
	return denyDelete()
}

//修改用户密码
func (self *Users) UpdatePassword(ctx *Context) error {
	//提交信息的校验
	errMessages, err := usersValidation.CheckPass(ctx)
	if !err {
		return Error{What: errMessages}
	}
	post := utils.GetPostData(ctx)
	pass := utils.MD5(post.Get("old_password"))
	idStr := post.Get("id")
	sql := "SELECT id FROM users WHERE id=" + idStr + " AND password='" + pass + "'"
	conn := db.MyEngine[(*ctx).Params().Get("platform")]
	rows, errRows := conn.SQL(sql).QueryString()
	if (errRows != nil) || (len(rows) < 1) {
		return Error{What: "用户信息不存在或旧密码错误"}
	}
	row := rows[0]
	if row["password"] == pass {
		return Error{What: "新密码与旧密码不能一致"}
	}
	newPass := utils.MD5(post.Get("password"))
	sql = "UPDATE users SET password='" + newPass + "' WHERE id=" + idStr
	result, errUpdate := conn.Exec(sql)
	if errUpdate != nil {
		return Error{What: "保存信息失败"}
	}
	affectedRows, errSave := result.RowsAffected()
	if errSave != nil || affectedRows <= 0 {
		return Error{What: "保存新的密码失败"}
	}
	return nil
}

//修改用户安全密码
func (self *Users) UpdateSafePassword(ctx *Context) error {
	//提交信息的校验
	errMessages, err := usersValidation.CheckSafePass(ctx)
	if !err {
		return Error{What: errMessages}
	}
	post := utils.GetPostData(ctx)
	pass := utils.MD5(post.Get("old_password"))
	idStr := post.Get("id")
	sql := "SELECT id FROM users WHERE id=" + idStr + " AND safe_password='" + pass + "'"
	conn := db.MyEngine[(*ctx).Params().Get("platform")]
	rows, errRows := conn.SQL(sql).QueryString()
	if (errRows != nil) || (len(rows) < 1) {
		return Error{What: "要修改安全密码的用户信息不存在或旧密码错误"}
	}
	row := rows[0]
	if row["password"] == pass {
		return Error{What: "新的安全密码与旧的密码不能一致"}
	}
	newPass := utils.MD5(post.Get("password"))
	sql = "UPDATE users SET safe_password='" + newPass + "' WHERE id=" + idStr
	result, errUpdate := conn.Exec(sql)
	if errUpdate != nil {
		return Error{What: "保存安全密码信息失败"}
	}
	affectedRows, errSave := result.RowsAffected()
	if errSave != nil || affectedRows <= 0 {
		return Error{What: "保存新的安全密码失败"}
	}
	return nil
}

//修改用户/代理状态
func (self *Users) changeStatus(ctx *Context, field string, toStatus string, title string) error {
	idStr := (*ctx).URLParam("id")
	_, err := strconv.Atoi(idStr)
	if err != nil {
		return Error{What: "必须提供用户编号"}
	}
	conn := db.MyEngine[(*ctx).Params().Get("platform")]
	sql := "UPDATE users SET " + field + " = " + toStatus + " WHERE id = '" + idStr + "' LIMIT 1"
	result, err := conn.Exec(sql)
	if err != nil {
		return Error{What: "锁定" + title + "状态失败"}
	}
	affectedRows, err := result.RowsAffected()
	if err != nil || affectedRows <= 0 {
		return Error{What: "锁定" + title + "状态失败"}
	}
	return nil
}

//锁定用户
func (self *Users) Lock(ctx *Context) error {
	return self.changeStatus(ctx, "status", "0", "用户")
}

//解锁用户
func (self *Users) Unlock(ctx *Context) error {
	return self.changeStatus(ctx, "status", "1", "用户")
}

//锁定用户
func (self *Users) LockProxy(ctx *Context) error {
	return self.changeStatus(ctx, "proxy_status", "0", "用户代理")
}

//解锁用户
func (self *Users) UnlockProxy(ctx *Context) error {
	return self.changeStatus(ctx, "proxy_status", "1", "用户代理")
}

//根据user_name/user_id 查询用户
func (self *Users) GetQueryUser(ctx *Context) (map[string]string, error) {
	userName := (*ctx).URLParam("user_name")
	if userName == "" {
		userName = "0"
	}
	sql := "SELECT  u.id, u.name, u.user_type, u.user_name, a.balance_wallet, a.charged_amount " +
		"FROM accounts AS a LEFT JOIN users AS u ON a.user_id = u.id " +
		"WHERE u.user_name = '" + userName + "'"
	userId, err := strconv.Atoi(userName)
	if err == nil {
		sql += " OR u.id = '" + strconv.Itoa(userId) + "'"
	}
	conn := db.MyEngine[(*ctx).Params().Get("platform")]
	rows, err := conn.SQL(sql).QueryString()
	if err != nil || len(rows) == 0 {
		return map[string]string{}, Error{What: "用户不存在"}
	}
	return rows[0], nil
}
