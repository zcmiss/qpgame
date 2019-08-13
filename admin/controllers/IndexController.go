package controllers

import (
	"fmt"
	"qpgame/admin/models"
)

// 后台登录相关
type IndexController struct{}

// 用于测试或性能测试
func (self *IndexController) Index(ctx *Context) {
	responseSuccess(ctx, "Hello world!", nil)
}

/**
 * @api {get} admin/api/v1/index/verify 获取图片验证码
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>图片验证码</strong><br>
 * 业务描述: 获取图片验证码</br>
 * @apiVersion 1.0.0
 * @apiName    	verify
 * @apiGroup    login
 * @apiPermission PC客户端

 * @apiParam (客户端请求参数) {string} 	key 	随机字符用作唯一标识
 *
 * @apiError (请求失败返回) {int}      code            错误代码
 * @apiError (请求失败返回) {string}   clientMsg       提示信息
 * @apiError (请求失败返回) {string}   internalMsg     内部错误信息
 * @apiError (请求失败返回) {float}    timeConsumed    后台耗时
 * @apiErrorExample {json} 失败返回
 * {
 *      "code": 204,
 *      "internalMsg": "",
 *      "clientMsg ": 0,
 *      "timeConsumed": 0
 * }
 *
 * @apiSuccess (返回结果)  {int}    	code            200, 成功
 * @apiSuccess (返回结果)  {string} 	clientMsg       提示信息
 * @apiSuccess (返回结果)  {string} 	internalMsg     内部错误信息
 * @apiSuccess (返回结果)  {json}  	data            image: base64字符串, key: 验证码唯一标识，登录时传入
 * @apiSuccess (返回结果)  {float}  	timeConsumed    后台耗时
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "数据获取成功",
 *    "code": 200,
 *    "data": {
 *         "image": "图片的base64字符串",
 *         "key": "验证码唯一标识"
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *IndexController) Verify(ctx *Context) {
	image, key := models.Verify(ctx)
	if key == "" {
		responseFailure(ctx, "验证码唯一标识生成失败", "网络异常，请重试")
		return
	}
	result := map[string]string{"image": image, "key": key}
	responseSuccess(ctx, "", result)
}

/**
 * @api {post} admin/api/v1/index/login 用户登录系统
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>后台用户登录</strong><br />
 * 业务描述: 用户登录系统 <br />
 * @apiVersion 1.0.0
 * @apiName     login
 * @apiGroup    login
 * @apiPermission PC客户端

 * @apiParam (客户端请求参数) {string} 	username 	用户名称
 * @apiParam (客户端请求参数) {string}   	password	用户密码
 * @apiParam (客户端请求参数) {string}   	code        验证码
 * @apiParam (客户端请求参数) {string}   	codeKey     验证码唯一标识，从生成验证码接口中获取，随时更新
 *
 * @apiError (请求失败返回) {int}      code            错误代码
 * @apiError (请求失败返回) {string}   clientMsg       提示信息
 * @apiError (请求失败返回) {string}   internalMsg     内部错误信息
 * @apiError (请求失败返回) {float}    timeConsumed   	 后台耗时
 *
 * @apiErrorExample {json} 失败返回
 * {
 *      "code": 204,
 *      "internalMsg": "",
 *      "clientMsg ": "",
 *      "timeConsumed": 0
 * }
 *
 * @apiSuccess (返回结果)  {int} 		code            200
 * @apiSuccess (返回结果)  {string} 	clientMsg       提示信息
 * @apiSuccess (返回结果)  {string}  	internalMsg     内部错误信息
 * @apiSuccess (返回结果)  {json}  	data            返回数据
 * @apiSuccess (返回结果)  {float}   	timeConsumed    后台耗时

 * @apiSuccess (data字段说明) {string}  	token 		用户登录的token
 * @apiSuccess (data字段说明) {object}	info		登录用户的信息
 * @apiSuccess (data字段说明) {array}    	menus		菜单节点信息

 * @apiSuccess (data-info每个子对象字段说明) {int}		  id 			用户编号
 * @apiSuccess (data-info每个子对象字段说明) {int}    	role_id		用户角色编号
 * @apiSuccess (data-info每个子对象字段说明) {string}    name			用户名称
 * @apiSuccess (data-info每个子对象字段说明) {string}    real_name 	真实姓名
 * @apiSuccess (data-info每个子对象字段说明) {int}    	max_manual_amount		最大人工入款金额

 * @apiSuccess (data-menus每个子对象字段说明) {int}		  id		菜单编号
 * @apiSuccess (data-menus每个子对象字段说明) {string}     title		菜单名称
 * @apiSuccess (data-menus每个子对象字段说明) {int}        level		菜单级别(一级菜单/二级菜单/三级菜单...)
 * @apiSuccess (data-menus每个子对象字段说明) {string}      url 		链接地址
 * @apiSuccess (data-menus每个子对象字段说明) {array}     menus 		菜单(多级)
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "保存数据成功",
 *    "code": 200,
 *    "data": {
 *      "token" : "",
 *      "info": {
 *        "id": "",
 *        "name": "",
 *        "real_name": "",
 *        "max_manual_amount: ""
 *      },
 *      "menus": [{
 *        "id": "",
 *        "title": "",
 *        "level": "",
 *        "url": "",
 *        "menus": []
 *      }]
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *IndexController) Login(ctx *Context) {
	result, err := models.Login(ctx)
	if err != nil {
		responseFailure(ctx, "", err.Error())
		return
	}
	responseSuccess(ctx, "登录成功", result)
}

func (self *IndexController) LoginReset(ctx *Context) {
	err := models.LoginReset(ctx)
	if err != nil {
		responseFailure(ctx, "", err.Error())
		return
	}
	responseSuccess(ctx, "操作成功", nil)
}

/**
 * @api {get} admin/api/v1/index/logout 退出系统登录
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>退出系统登录</strong><br />
 * 业务描述: 退出系统登录</br>
 * @apiVersion 1.0.0
 * @apiName    	logout
 * @apiGroup    login
 * @apiPermission PC客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }

 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 token
 * @apiHeaderExample {json} 请求头示例
 * {
 *      Authorization: Bearer CIiwic3ViIjoyNDExMn0.8r4yPplyuQ5KIKLnmiBBoMJ04YXVLOSpeFLCWRyOFC......
 * }
 *
 * @apiError (请求失败返回) {int}      code            错误代码
 * @apiError (请求失败返回) {string}   clientMsg       提示信息
 * @apiError (请求失败返回) {string}   internalMsg     内部错误信息
 * @apiError (请求失败返回) {float}    timeConsumed    后台耗时
 * @apiErrorExample {json} 失败返回
 * {
 *      "code": 204,
 *      "internalMsg": "",
 *      "clientMsg ": 0,
 *      "timeConsumed": 0
 * }
 *
 * @apiSuccess (返回结果)  {int}    	code            200, 成功
 * @apiSuccess (返回结果)  {string} 	clientMsg       提示信息
 * @apiSuccess (返回结果)  {string} 	internalMsg     内部错误信息
 * @apiSuccess (返回结果)  {json}  	data            返回数据
 * @apiSuccess (返回结果)  {float}  	timeConsumed    后台耗时
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "退出登录成功",
 *    "code": 200,
 *    "data": ""
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *IndexController) Logout(ctx *Context) {
	err := models.Logout(ctx)
	if err != nil {
		responseFailure(ctx, "", err.Error())
		return
	}
	responseSuccess(ctx, "用户退出成功", "")
}

/**
 * @api {get} admin/api/v1/index/config 获取网站配置
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
 * <strong>获取网站配置</strong><br>
 * 业务描述: 获取网站配置, 不需要平台识别号</br>
 * @apiVersion 1.0.0
 * @apiName    	config
 * @apiGroup    login
 * @apiPermission PC客户端
 *
 * @apiError (请求失败返回) {int}      code            错误代码
 * @apiError (请求失败返回) {string}   clientMsg       提示信息
 * @apiError (请求失败返回) {string}   internalMsg     内部错误信息
 * @apiError (请求失败返回) {float}    timeConsumed    后台耗时
 * @apiErrorExample {json} 失败返回
 * {
 *      "code": 204,
 *      "internalMsg": "",
 *      "clientMsg ": 0,
 *      "timeConsumed": 0
 * }
 *
 * @apiSuccess (返回结果)  {int}    	code            200, 成功
 * @apiSuccess (返回结果)  {string} 	clientMsg       提示信息
 * @apiSuccess (返回结果)  {string} 	internalMsg     内部错误信息
 * @apiSuccess (返回结果)  {json}  	data            网站配置信息
 * @apiSuccess (返回结果)  {float}  	timeConsumed    后台耗时
 *
 * @apiSuccess (data字段说明) {string}  	code		平台识别号
 * @apiSuccess (data字段说明) {string}	name		平台名称
 * @apiSuccess (data字段说明) {string}    url			后台地址
 * @apiSuccess (data字段说明) {string}    api_url		后台API地址
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *    "clientMsg": "数据获取成功",
 *    "code": 200,
 *    "data": {
 *      "code": "",
 *      "name": "",
 *      "url": "",
 *      "api_url": ""
 *    },
 *    "internalMsg": "",
 *    "timeConsumed": 168
 * }
 */
func (self *IndexController) Config(ctx *Context) {
	conf, err := models.GetConfig(ctx)
	if err == nil {
		responseSuccess(ctx, "获取配置成功", conf)
	} else {
		fmt.Println(err.Error())
		responseFailure(ctx, err.Error(), "获取配置失败")
	}

}
