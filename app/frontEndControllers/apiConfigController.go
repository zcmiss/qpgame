package frontEndControllers

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/kataras/iris"
	"github.com/shopspring/decimal"
	"qpgame/common/utils"
	"qpgame/config"
	"qpgame/models/mainxorm"
	"qpgame/models/xorm"
	"qpgame/ramcache"
	"strconv"
)

type ApiConfigController struct {
	platform string
	ctx      iris.Context
}

//构造函数
func NewApiConfigController(ctx iris.Context) *ApiConfigController {
	obj := new(ApiConfigController)
	obj.platform = ctx.Params().Get("platform")
	obj.ctx = ctx
	return obj
}

/**
 * @api {get} api/v1/getAppVerStatus 获取版本状况
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人:把妹Pling</span><br/><br/>
 * 获取版本信息情况<br>
 * 业务描述:在app启动的时候需要去检测版本是否有更新，如果本接口没有响应，或者没有数据</br>
 * 不做任务处理,忽略即可,如果有数据跟本地版本号进行比对如果小于打回的版本号提示更新</br>
 * @apiVersion 1.0.0
 * @apiName     api_v1_getAppVerStatus
 * @apiGroup    config_module
 * @apiPermission ios,android客户端
 * @apiParam (客户端请求参数) {int} app_type   1安卓,2ios
 * @apiSuccess (返回结果)  {int}      code            -4失败,200正常响应数据,158无数据,空对象
 * @apiSuccess (返回结果)  {string}   clientMsg       提示信息
 * @apiSuccess (返回结果)  {string}   internalMsg     提示信息
 * @apiSuccess (返回结果)  {json}  	  data            返回数据
 * @apiSuccess (返回结果)  {float}   timeConsumed    后台耗时
 * @apiSuccess (data对象字段说明) {string}   description 版本更新描述
 * @apiSuccess (data对象字段说明) {string}   link 更新地址
 * @apiSuccess (data对象字段说明) {int}   update_type 更新类型 1强制更新，2提示更新，3不提示更新
 * @apiSuccess (data对象字段说明) {string}   version 版本号
 * @apiSuccessExample {json} 响应结果
 {
    "clientMsg": "",
    "code": 200,
    "data": {
        "description": "测试的appversion",
        "link": "无",
        "update_type": 1,
        "version": 1
    },
    "internalMsg": "",
    "timeConsumed": 80
}
*/
func (cthis *ApiConfigController) GetAppVerUpdate() {
	ctx := cthis.ctx
	if !utils.RequiredParam(&ctx, []string{"app_type"}) {
		return
	}
	appType, errAppType := ctx.URLParamInt("app_type")
	//不是数字并且不是安卓和ios
	if errAppType != nil || (appType != 1 && appType != 2) {
		utils.ResFaiJSON(&ctx, "app_type参数错误", "更新检测失败", config.PARAMERROR)
		return
	}
	ave, _ := ramcache.TableAppVersions.Load(cthis.platform)
	appVersions := ave.([]xorm.AppVersions)
	var res = make(map[string]interface{})
	for _, v := range appVersions {
		if v.AppType == appType && v.Status == 1 {
			res["version"] = v.Version
			res["description"] = v.Description
			res["link"] = v.Link
			res["update_type"] = v.UpdateType
			utils.ResSuccJSON(&ctx, "", "", config.SUCCESSRES, res)
			return
		}
	}
	utils.ResSuccJSON(&ctx, "", "", config.SUCCESSRES, res)
}

/**
 * @api {get} api/v1/getConfig 获取系统配置
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人:把妹Pling</span><br/><br/>
 * 获取系统配置情况(根据开发情况，会动态新增)<br>
 * 业务描述:在app启动的时候每次都需要去获取这个接口的配置，注意事项获取到数据之后要缓存到本地，下次进来的时候，如果/br>
 * 请求不到数据就用本地的缓存数据，如果正常响应就替换掉本地的缓存,保证无论如何都要启动成功进入首页</br>
 * @apiVersion 1.0.0
 * @apiName     api_v1_getConfig
 * @apiGroup    config_module
 * @apiPermission ios,android客户端
 * @apiSuccess (返回结果)  {int}      code            200正常响应数据
 * @apiSuccess (返回结果)  {string}   clientMsg       提示信息
 * @apiSuccess (返回结果)  {string}   internalMsg     提示信息
 * @apiSuccess (返回结果)  {json}     data            返回数据
 * @apiSuccess (返回结果)  {float}    timeConsumed    后台耗时
 * @apiSuccess (data对象字段说明) {string}   api_address 动态接口地址
 * @apiSuccess (data对象字段说明) {string}   can_register 是否开放注册
 * @apiSuccess (data对象字段说明) {string}   tuiguang_web_url 推广地址
 * @apiSuccess (data对象字段说明) {int}      withdraw_min_money 提现最低金额
 * @apiSuccess (data对象字段说明) {int}      withdraw_max_money 提现最大金额
 * @apiSuccess (data对象字段说明) {bool}     bind_phone_award_switch 绑定手机是否有奖励
 * @apiSuccess (data对象字段说明) {bool}     sign_award_switch 每日签到是否有奖励
 * @apiSuccess (data对象字段说明) {object}   fund_change_type 资金变动类型
 * @apiSuccessExample {json} 响应结果
 * {
 *     "clientMsg": "",
 *     "code": 200,
 *     "data": {
 *         "api_address": "https://qpgametest.jjwx88.com/CKYX",
 *         "bind_phone_award_switch": true,
 *         "can_register": 1,
 *         "fund_change_type": [
 *             {
 *                 "id": "",
 *                 "name": "显示全部"
 *             },
 *             {
 *                 "id": "1",
 *                 "name": "充值"
 *             },
 *             {
 *                 "id": "2",
 *                 "name": "提取现金"
 *             },
 *             {
 *                 "id": "3",
 *                 "name": "洗码"
 *             },
 *             {
 *                 "id": "4",
 *                 "name": "保险箱存取款"
 *             },
 *             {
 *                 "id": "5",
 *                 "name": "赠送彩金"
 *             },
 *             {
 *                 "id": "6",
 *                 "name": "优惠入款"
 *             },
 *             {
 *                 "id": "7",
 *                 "name": "代理佣金提成"
 *             },
 *             {
 *                 "id": "8",
 *                 "name": "签到奖励"
 *             },
 *             {
 *                 "id": "9",
 *                 "name": "活动奖励"
 *             },
 *             {
 *                 "id": "10",
 *                 "name": "提现未通过返还"
 *             },
 *             {
 *                 "id": "11",
 *                 "name": "额度转换-转入"
 *             },
 *             {
 *                 "id": "12",
 *                 "name": "额度转换失败返还"
 *             },
 *             {
 *                 "id": "13",
 *                 "name": "额度转换-转出"
 *             },
 *             {
 *                 "id": "14",
 *                 "name": "红包收入"
 *             },
 *             {
 *                 "id": "15",
 *                 "name": "提现退款"
 *             },
 *             {
 *                 "id": "16",
 *                 "name": "VIP晋级礼金"
 *             },
 *             {
 *                 "id": "17",
 *                 "name": "VIP周工资"
 *             },
 *             {
 *                 "id": "18",
 *                 "name": "VIP月工资"
 *             },
 *             {
 *                 "id": "19",
 *                 "name": "推广邀请奖励"
 *             },
 *             {
 *                 "id": "20",
 *                 "name": "绑定手机奖励"
 *             }
 *         ],
 *         "reward_bind": "20",
 *         "sign_award_switch": true,
 *         "sign_reward": "3",
 *         "tuiguang_web_url": "http://www.ckqp88vip.com/CKYX/app/download?parentid=",
 *         "withdraw_max_money": 1000,
 *         "withdraw_min_money": 1
 *         "report_qq": "10000"
 *     },
 *     "internalMsg": "",
 *     "timeConsumed": 998
 * }
 */
func (cthis *ApiConfigController) GetConfig() {
	ctx := cthis.ctx
	var res = make(map[string]interface{})
	mtp, _ := ramcache.MainTablePlatform.Load("platform")
	for _, v := range mtp.([]mainxorm.Platform) {
		if v.Code == cthis.platform {
			res["api_address"] = v.ApiAddress
		}
	}
	conf, _ := ramcache.TableConfigs.Load(cthis.platform)
	cfg := conf.(map[string]interface{})
	//是否开放注册
	res["can_register"] = cfg["register_config"].(map[string]interface{})["can_register"]
	//推广地址
	res["tuiguang_web_url"] = cfg["tuiguang_web_url"].(string) + "/?parentid="
	res["withdraw_min_money"] = cfg["withdraw_min_money"]
	res["withdraw_max_money"] = cfg["withdraw_max_money"]
	fSignAwardSwitch := cfg["sign_award_switch"].(float64)
	fSignReward := cfg["sign_reward"].(float64)
	fRewardBind := cfg["reward_bind"].(float64)
	fBindPhoneAwardSwitch := cfg["bind_phone_award_switch"].(float64)
	decimal.DivisionPrecision = 2
	res["sign_reward"] = decimal.NewFromFloat(fSignReward).String()
	res["reward_bind"] = decimal.NewFromFloat(fRewardBind).String()
	res["sign_award_switch"] = int(fSignAwardSwitch) == 1
	res["bind_phone_award_switch"] = int(fBindPhoneAwardSwitch) == 1
	res["report_qq"] = cfg["report_qq"].(string)
	var fundChangeTypeMap = []map[string]interface{}{
		{"id": "", "name": "显示全部"},
		{"id": strconv.Itoa(config.FUNDCHARGE), "name": "充值"},
		{"id": strconv.Itoa(config.FUNDWITHDRAW), "name": "提取现金"},
		{"id": strconv.Itoa(config.FUNDXIMA), "name": "洗码"},
		{"id": strconv.Itoa(config.FUNDSAFEBOX), "name": "保险箱存取款"},
		{"id": strconv.Itoa(config.FUNDPRESENTER), "name": "赠送彩金"},
		{"id": strconv.Itoa(config.FUNDDISCOUNTS), "name": "优惠入款"},
		{"id": strconv.Itoa(config.FUNDBROKERAGE), "name": "代理佣金提成"},
		{"id": strconv.Itoa(config.FUNDSIGNIN), "name": "签到奖励"},
		{"id": strconv.Itoa(config.FUNDACTIVITYAWARD), "name": "活动奖励"},
		{"id": strconv.Itoa(config.FUNDWITHRAWNOPASS), "name": "提现未通过返还"},
		{"id": strconv.Itoa(config.FUNDPLATCHANGEIN), "name": "额度转换-转入"},
		{"id": strconv.Itoa(config.FUNDPLATCHANGEINFAILBACK), "name": "额度转换失败返还"},
		{"id": strconv.Itoa(config.FUNDPLATCHANGEOUT), "name": "额度转换-转出"},
		{"id": strconv.Itoa(config.FUNDREDPACKET), "name": "红包收入"},
		{"id": strconv.Itoa(config.FUNDWITHDRAWBACK), "name": "提现退款"},
		{"id": strconv.Itoa(config.FUNDVIPUPLEVEL), "name": "VIP晋级礼金"},
		{"id": strconv.Itoa(config.FUNDVIPWEEK), "name": "VIP周工资"},
		{"id": strconv.Itoa(config.FUNDVIPMONTH), "name": "VIP月工资"},
		{"id": strconv.Itoa(config.FUNDGENERALIZEAWAED), "name": "推广邀请奖励"},
		{"id": strconv.Itoa(config.FUNDBINDPHONE), "name": "绑定手机奖励"},
	}
	res["fund_change_type"] = fundChangeTypeMap
	utils.ResSuccJSON(&ctx, "", "", config.SUCCESSRES, res)
}

/**
 * @api {get} api/v1/getCustomerService 获取客服页面
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人:把妹Pling</span><br/><br/>
 * 获取客服内容<br>
 * 业务描述:客服页面,这里的内容一般变动很低,一定要缓存到本地,有就更新,如果本地有缓存了先展示缓存的，然后在更新</br>
 * 如果请求内容没有变化,那么data返回一个空对象,请跟进对象里判断是否有cache_key来确定缓存是否有变化,如果有变化</br>
 * data里肯定有cache_key的值</br>
 * @apiVersion 1.0.0
 * @apiName     api_v1_getCustomerService
 * @apiGroup    config_module
 * @apiPermission ios,android客户端
 * @apiParam (客户端请求参数) {string} cache_key   缓存md5
 * @apiSuccess (返回结果)  {int}      code            200正常响应数据
 * @apiSuccess (返回结果)  {string}   clientMsg       提示信息
 * @apiSuccess (返回结果)  {string}   internalMsg     提示信息
 * @apiSuccess (返回结果)  {json}  	  data            返回数据,在缓存没有变化的情况下,返回一个空对象
 * @apiSuccess (返回结果)  {float}   timeConsumed    后台耗时
 * @apiSuccess (data对象字段说明) {string}   web_customer_url 网页客服
 * @apiSuccess (data对象字段说明) {string}   qq_customer qq客服
 * @apiSuccess (data对象字段说明) {string}   weixin_customer 微信客服
 * @apiSuccess (data对象字段说明) {string}   faq 常见问题富文本内容
 * @apiSuccess (data对象字段说明) {string}   cache_key 缓存值的md5,与前端传进来一致的时候data为空对象
 * @apiSuccessExample {json} 响应结果
{
    "clientMsg": "",
    "code": 200,
    "data": {
        "cache_key": "d8398d9837bcce4ea45afa04ede965d9",
		"faq":"",
        "qq_customer": [
            {
                "accout": "121123",
                "info": "客服专员",
                "name": "QQ客服一",
                "url": "头像地址"
            },
            {
                "accout": "1211234",
                "info": "财务咨询",
                "name": "QQ客服二",
                "url": "头像地址"
            },
            {
                "accout": "56793",
                "info": "客服专员",
                "name": "QQ客服三",
                "url": "头像地址"
            }
        ],
        "web_customer_url": "https://szzero.livechatvalue.com/chat/chatClient/chatbox.jsp?companyID=1059446&configID=60224&jid=3181657716&s=1",
        "weixin_customer": [
            {
                "accout": "qipai4564",
                "info": "VIP专员服务",
                "name": "微信客服一",
                "url": "微信二维码地址"
            },
            {
                "accout": "qw98duf",
                "info": "微信活动申请",
                "name": "微信客服二",
                "url": "微信二维码地址"
            },
            {
                "accout": "qp8iun8",
                "info": "财务咨询",
                "name": "微信客服三",
                "url": "微信二维码地址"
            }
        ]
    },
    "internalMsg": "",
    "timeConsumed": 223
}
*/

func (cthis *ApiConfigController) GetCustomerService() {
	ctx := cthis.ctx
	var res = make(map[string]interface{})
	var empty = make(map[string]interface{})
	cacheKey := ctx.URLParam("cache_key")
	tConf, _ := ramcache.TableConfigs.Load(cthis.platform)
	conf := tConf.(map[string]interface{})
	//网页客服
	res["web_customer_url"] = conf["web_customer_url"]
	res["qq_customer"] = conf["qq_customer"]
	res["weixin_customer"] = conf["weixin_customer"]
	res["faq"] = conf["faq"]
	content, _ := json.Marshal(res)
	md5hash := fmt.Sprintf("%x", md5.Sum(content))
	res["cache_key"] = md5hash
	if cacheKey == md5hash {
		res = empty
	}
	utils.ResSuccJSON(&ctx, "", "", config.SUCCESSRES, res)
}
