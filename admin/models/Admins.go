package models

import (
	"qpgame/admin/common"
	"qpgame/admin/validations"
	"qpgame/common/utils"
	"qpgame/models"
	"strconv"
	"strings"
	"time"
)

// 模型
type Admins struct{}

var adminValidation = validations.AdminsValidation{}

// 表名称
func (self *Admins) GetTableName(ctx *Context) string {
	return "admins"
}

// 得到所有记录-分页
func (self *Admins) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, name, email, role_id, created, updated, status, "+
			"withdraw_alert, permission, force_out, manual_max, is_otp",
		func(ctx *Context) []string { //获取查询条件
			return getQueryFields(ctx, &map[string]string{
				"name":    "%",
				"role_id": "=",
				"status":  "=",
				"is_otp":  "=",
			})
		},
		func(ctx *Context, row *map[string]string) { //对于查询出来的每条记录的处理
			processOptions("charge_alert", &statusTypes, row)
			processOptions("withdraw_alert", &statusTypes, row)
			processOptions("is_otp", &yesNo, row)
			processOptions("is_otp_first", &yesNo, row)
			processOptions("status", &statusTypes, row)
			processOptions("force_out", &yesNo, row)
			processOptions("permission", &adminPermissions, row)
			roleId, err := strconv.Atoi((*row)["role_id"])
			if err != nil {
				(*row)["role_name"] = ""
				return
			}
			platform := (*ctx).Params().Get("platform")
			(*row)["role_name"] = common.GetRoleName(platform, int(roleId))
		}, nil)
}

// 得到记录详情
func (self *Admins) GetRecordDetail(ctx *Context) (map[string]string, error) {
	return getRecordDetail(ctx, self, "", nil)
}

// 添加记录
func (self *Admins) Save(ctx *Context) (int64, error) {
	//如果有密码则需要更改密码
	changePassword := func(ctx *Context, data *map[string]string) {
		post := utils.GetPostData(ctx)
		submitPass := post.Get("password")
		if strings.Compare(submitPass, "") == 0 {
			return
		}
		(*data)["password"] = utils.MD5(submitPass)
	}
	return saveRecord(ctx, self,
		func(ctx *Context, data *map[string]string) bool { //汪厍
			changePassword(ctx, data)
			return true
		},
		func(ctx *Context, data *map[string]string) bool { //添加之前处理
			(*data)["created"] = strconv.FormatInt(time.Now().Unix(), 10) //添加时间
			(*data)["updated"] = strconv.FormatInt(time.Now().Unix(), 10) //修改时间
			return true
		},
		func(ctx *Context, data *map[string]string) bool { //修改之前处理
			(*data)["updated"] = strconv.FormatInt(time.Now().Unix(), 10) //修改时间
			return true
		}, getSavedFunc("后台用户", "name"))
}

// 删除记录
func (self *Admins) Delete(ctx *Context) error {
	return deleteRecord(ctx, self, nil, getDeletedFunc("后台用户"))
}

//修改用户密码
func (self *Admins) UpdatePassword(ctx *Context) error {
	//提交信息的校验
	errMessages, err := adminValidation.CheckPass(ctx)
	if !err {
		return Error{What: errMessages}
	}
	post := utils.GetPostData(ctx)
	pass := utils.MD5(post.Get("old_password")) //旧的密码
	newPass := utils.MD5(post.Get("password"))
	if newPass == pass {
		return Error{What: "新旧密码不能一样"}
	}

	idStr := post.Get("id")
	sql := "SELECT id,password FROM admins WHERE id=" + idStr
	conn := models.MyEngine[(*ctx).Params().Get("platform")]
	rows, errRows := conn.SQL(sql).QueryString()
	if errRows != nil || len(rows) == 0 {
		return Error{What: "后台用户信息不存在或旧密码错误"}
	}
	row := rows[0]
	if row["password"] != pass {
		return Error{What: "旧密码错误"}
	}

	sql = "UPDATE admins SET `password` = '" + newPass + "' WHERE id = '" + idStr + "'"
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
