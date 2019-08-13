package models

import (
	"qpgame/admin/common"
	"qpgame/admin/validations"
	"qpgame/common/utils"
	"qpgame/models"
)

// 模型
type AdminRoles struct{}

var adminRolesValidation = validations.AdminRolesValidation{}

// 表名称
func (self *AdminRoles) GetTableName(ctx *Context) string {
	return "admin_roles"
}

// 得到所有记录-分页
func (self *AdminRoles) GetRecords(ctx *Context) (Pager, error) {
	return getRecords(ctx, self,
		"id, name, status, remark",
		func(ctx *Context) []string { //获取查询条件
			return getQueryFields(ctx, &map[string]string{
				"name":   "%",
				"status": "=",
			})
		}, nil, nil)
}

// 得到记录详情
func (self *AdminRoles) GetRecordDetail(ctx *Context) (map[string]string, error) {
	return getRecordDetail(ctx, self, "", nil)
}

// 添加记录
func (self *AdminRoles) Save(ctx *Context) (int64, error) {
	return saveRecord(ctx, self, nil, nil, nil, getSavedFunc("后台用户角色", "name"))
}

// 删除记录
func (self *AdminRoles) Delete(ctx *Context) error {
	err := deleteRecord(ctx, self, nil, getDeletedFunc("后台用户角色"))
	// 更新缓存
	if err == nil {
		common.LoadAdminRoleMenus()
		common.LoadAdmins()
		common.LoadAdminMenus()
		common.LoadAdminRoles()
	}
	return err
}

// 修改角色所属的菜单
func (self *AdminRoles) SetMenus(ctx *Context) error {
	vali, res := adminRolesValidation.CheckMenus(ctx)
	if !res {
		return Error{What: vali}
	}
	post := utils.GetPostData(ctx)
	id := post.Get("id")
	menuIds := post.Get("menu_ids")
	sql := "UPDATE admin_roles SET menu_ids = '" + menuIds + "' WHERE id = '" + id + "' LIMIT 1"
	conn := models.MyEngine[(*ctx).Params().Get("platform")]
	result, err := conn.Exec(sql)
	if err != nil {
		return Error{What: "修改角色菜单失败"}
	}
	affectedRows, affErr := result.RowsAffected()
	if affErr != nil || affectedRows <= 0 {
		return Error{What: "修改角色菜单失败"}
	}
	// 更新缓存
	{
		common.LoadAdminRoleMenus()
		common.LoadAdmins()
		common.LoadAdminMenus()
		common.LoadAdminRoles()
	}
	return nil
}
