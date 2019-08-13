package controllers

import (
	"qpgame/admin/models"
	"qpgame/common/mvc"
	"qpgame/common/utils"
	"qpgame/config"

	"github.com/kataras/iris"
)

/**
 * @apiDefine login 登录模块
 *
 */

/**
 * @apiDefine content 内容管理
 *
 */

/**
 * @apiDefine game 游戏管理
 */

/**
 * @apiDefine admin 后台管理
 *
 */

/**
 * @apiDefine finance 财务管理
 *
 */

/**
 * @apiDefine user 用户管理
 *
 */

/**
 * @apiDefine report 报表统计
 *
 */

/**
 * @apiDefine system 系统管理
 *
 */

/**
 * @apiDefine websocket WebSocket
 *
 */
/**
 * @apiDefine silver_merchant 银商管理
 *
 */

// 类型别名定义
type IAdminModel = models.IAdminModel //Model接口
type IValidator = mvc.IValidator      //验证器接口
type Context = iris.Context           //简写类型定义

// 数据列表
func index(ctx *Context, model IAdminModel) {
	records, err := model.GetRecords(ctx)
	if err != nil {
		responseFailure(ctx, err.Error(), "获取数据失败")
		return
	}
	responseSuccess(ctx, "获取数据成功", records)
}

// 修改数据
func save(ctx *Context, model IAdminModel, validator IValidator) {
	messages, result := validator.Validate(ctx)
	if !result { //如果数据校验失败, 则运回错误信息
		responseFailure(ctx, "", messages)
		return
	}
	affectedRows, err := model.Save(ctx)
	if err != nil {
		responseFailure(ctx, err.Error(), err.Error())
		return
	}
	if affectedRows <= 0 {
		responseFailure(ctx, "操作记录数为0", "保存数据失败")
		return
	}
	responseSuccess(ctx, "数据保存成功", nil)
}

// 查看单条数据
func view(ctx *Context, model IAdminModel) {
	record, err := model.GetRecordDetail(ctx)
	if err != nil {
		responseFailure(ctx, err.Error(), "获取详情失败")
		return
	}
	responseSuccess(ctx, "", record)
}

// 删除记录
func remove(ctx *Context, model IAdminModel) {
	result := model.Delete(ctx)
	if result != nil {
		responseFailure(ctx, result.Error(), "数据删除失败")
		return
	}
	responseSuccess(ctx, "数据删除成功", nil)
}

// 输出操作相关的处理结果
func responseResult(ctx *Context, err error, successMessage string) {
	if err != nil {
		responseFailure(ctx, "", err.Error())
		return
	}
	responseSuccess(ctx, successMessage, "")
}

// 输出正确内容的json的简便方法
func responseSuccess(ctx *Context, message string, data interface{}) {
	//utils.ResponseCors(ctx)
	models.WriteAdminLog(ctx) //如果是操作成功的状态，则写入相关的操作日志
	utils.ResSuccJSON(ctx, "", message, config.SUCCESSRES, data)
}

// 输出错误的json内容的简便方法
func responseFailure(ctx *Context, internalErrorMessage string, clientErrorMessage string) {
	//utils.ResponseCors(ctx)
	utils.ResSuccJSON(ctx, internalErrorMessage, clientErrorMessage, config.NOTGETDATA, "")
}
