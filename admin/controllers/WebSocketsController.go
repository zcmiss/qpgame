package controllers

/**
* @api {ws} ws 充值提现统计信息
* @apiDescription
* <span style="color:lightcoral;">接口负责人: juzi</span><br/><br/>
* <strong>获取待处理的充值与提现统计信息, Websocket访问模式</strong><br />
* 业务描述: 获取待处理的充值与提现统计信息</br>
* 协议: ws 或 wss</br>
* 地址: 协议 + 接口地址 + ws, 如: ws://api.admin-qp.com/ws </br>
* @apiVersion 1.0.0
* @apiName    	websocketAlert
* @apiGroup    websocket
* @apiPermission PC客户端
* @apiParam (客户端请求参数) {string} 	platform	平台识别号
* @apiParam (客户端请求参数) {string} 	type 	类型，包括以下: <br />
* init: 连接之后的获取初始化信息 (PC后台), <br />
* 	<span style="color:red">示例: {"platform": "CKYX", "type":"init", "id": 1}</span><br />
* statistic: 获取信息 (PC后台), <br />
* 	<span style="color:red">示例: {"platform": "CKYX", "type":"statistic"}</span> <br />
* heart: 心跳检测 (PC后台/手机端), <br />
* 	<span style="color:red">示例: {"platform": "CKYX", "type":"heart"}</span> <br />
* charge: 充值通知提醒 (PC后台/手机端), <br />
* 	<span style="color:red">示例: {"platform": "CKYX", "type":"charge"}</span> <br />
* withdraw: 提现通知提醒 (PC后台/手机端), <br />
* 	<span style="color:red">示例: {"platform": "CKYX", "type":"charge"}</span> <br />
* sound_charge: 充值声音开/关 (PC后台), <br />
* 	<span style="color:red">示例: {"platform": "CKYX", "type":"sound_charge", "id": 1} </span><br />
* sound_withdraw: 提现声音开/关 (PC后台), <br />
* 	<span style="color:red">示例: {"platform": "CKYX", "type":"sound_withdraw", "id":1} </span>
* @apiParam (客户端请求参数) {int} 	id	后台用户编号, 在 type 为 init/statistic/sound_charge/sound_withdraw 时需要
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
* @apiSuccess (返回结果)  {string}    	type        类型, 提交类型 => 返回类型 <br />
* 														"heart" => "hearted" <br />
* 														"init" => "init" <br />
* 														"statistic" => "statistic" <br />
* 														"charge" => "charge" <br />
* 														"withdraw" => "withdraw" <br />
* 														"sound_charge" => "sound_charge" <br />
* 														"sound_withdraw" => "sound_withdraw" <br />
* @apiSuccess (返回结果)  {string} 	clientMsg       提示信息
* @apiSuccess (返回结果)  {json}  	data            返回数据
* @apiSuccess (返回结果)  {float}  	timeConsumed    后台耗时
*
* @apiSuccess (data字段说明) {string}  	charge_sound	充值提示声音文件路径 <br /><span style="color:red">(当 {"type":"init"} 时, 有此字段)</span>
* @apiSuccess (data字段说明) {string}  	charge_alert 	充值提示音状态, 0: 未开启, 1:已开启 <br /><span style="color:red">(当 {"type":"init"} 时, 有此字段)</span>
* @apiSuccess (data字段说明) {string}  	withdraw_sound	提现提示声音文件路径 <br /><span style="color:red">(当 {"type":"init"} 时, 有此字段)</span>
* @apiSuccess (data字段说明) {string}  	withdraw_alert	提现提示声音状态, 0: 未开启, 1:已开启 <br /><span style="color:red">(当 {"type":"init"} 时, 有此字段)</span>
* @apiSuccess (data字段说明) {int}  	status 				充值/提现提示声音开启或关闭状态, 0:已关闭, 1:已开启<br />
* <span style="color:red">(当 {"type":"sound_charge"} 或 {"type":"sound_withdraw"}时, 有此字段)</span>
* @apiSuccess (data字段说明) {int}  		charge_count	待处理充值数量<br /><span style="color:red">(当 {"type":"statistic"} 或 {"type":"init"}时, 有此字段)</span>
* @apiSuccess (data字段说明) {int}		withdraw_count	待处理提现数量<br /><span style="color:red">(当 {"type":"statistic"} 或 {"type":"init"}时, 有此字段)</span>
*
* @apiSuccessExample {json} 响应结果
* {
*    "clientMsg": "",
*    "code": 200,
*    "data":  {
*      "charge_count": "1",
*      "withdraw_count": "1",
*    },
*    "internalMsg": "",
*    "timeConsumed": 168
* }
 */
func (self *IndexController) Alert(ctx *Context) {
	responseFailure(ctx, "此方法仅用于WebSocket的API说明", "禁止访问")
}
