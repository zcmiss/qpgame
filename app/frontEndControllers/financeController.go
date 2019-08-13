package frontEndControllers

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	goxorm "github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	"qpgame/app/fund"
	"qpgame/common/log"
	"qpgame/common/utils"
	"qpgame/config"
	"qpgame/models"
	"qpgame/models/beans"
	"qpgame/models/xorm"
	"qpgame/ramcache"
	"sort"
	"strconv"
	"strings"
)

type FinanceController struct {
	platform string
	ctx      iris.Context
}

//构造函数
func NewFinanceController(ctx iris.Context) *FinanceController {
	obj := new(FinanceController)
	obj.platform = ctx.Params().Get("platform")
	obj.ctx = ctx
	return obj
}

/**
 * @api {get} api/auth/v1/chargeTypes APP获取充值类型列表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人:Pling把妹</span><br/><br/>
 * 当客户端进入充值页面获取充值类型列表的时候<br>
 * 业务描述:获取充值类型列表,在左边展示支付的类型列表，点击对应的列表之后展示pay_info充值选项</br>
 * 点击充值调用 api/auth/v1/payAction 接口 (第三方支付)
 * 银行卡充值调用 api/auth/v1/payCompanyAction
 * @apiVersion 1.0.0
 * @apiName api_auth_v1_chargeTypes
 * @apiGroup finance_module
 * @apiPermission ios,android客户端
 * @apiHeader (客户端请求头参数) {string} Authorization Bearer + 用户登录获得的token
 * @apiParam (客户端请求参数) {string} cache_key md5的值根据上一次访问接口返回的值，第一次为空。如果内容没变打回空对象
 * @apiSuccess (返回结果)  {int}      code            200
 * @apiSuccess (返回结果)  {string}   clientMsg       提示信息
 * @apiSuccess (返回结果)  {string}   internalMsg     提示信息
 * @apiSuccess (返回结果)  {json}  	  data            返回数据
 * @apiSuccess (返回结果)  {float}   timeConsumed    后台耗时
 * @apiSuccess (data对象字段说明) {array}   list 数据列表(无数据时为空数组list:[])
 * @apiSuccess (data对象字段说明) {string}   cache_key 判断是否缓存的md5值
 * @apiSuccess (data-list元素对象字段说明) {string}   name 支付类型名称
 * @apiSuccess (data-list元素对象字段说明) {array}    charge_numbers 便捷支付金额数组,展示在右边
 * @apiSuccess (data-list元素对象字段说明) {string}   logo 图片地址
 * @apiSuccess (data-list元素对象字段说明) {array}    pay_info 支付平台列表,json对象数组
 * @apiSuccess (data-list元素对象字段说明) {string}   remark 充值说明
 * @apiSuccess (data-list-pay_info元素对象字段说明) {string}    name 支付平台名称
 * @apiSuccess (data-list-pay_info元素对象字段说明) {string}    owner 持卡人
 * @apiSuccess (data-list-pay_info元素对象字段说明) {string}    card_number 收款卡号或收款账号
 * @apiSuccess (data-list-pay_info元素对象字段说明) {string}    title 支付标题
 * @apiSuccess (data-list-pay_info元素对象字段说明) {string}    hint 支付说明,如果addr_type = 7的话hint显示在最上面
 * @apiSuccess (data-list-pay_info元素对象字段说明) {string}    logo charge_type_id=3时，为银行卡图标
 * @apiSuccess (data-list-pay_info元素对象字段说明) {string}    bank_address 地址说明,如果addr_type=2的话就是二维码图片地址
 * @apiSuccess (data-list-pay_info元素对象字段说明) {string}    addr_type 地址类型支付类型: 1.入款转账,2.支付二维码,3.支付地址，4.第三方二维码支付,5.第三方 wap支付,6.第三方h5支付,7.代理支付
 * @apiSuccess (data-list-pay_info元素对象字段说明) {integer}   mfrom 支付额度下限
 * @apiSuccess (data-list-pay_info元素对象字段说明) {integer}   mto 支付额度上限
 * @apiSuccess (data-list-pay_info元素对象字段说明) {integer}   credential_id 支付证书Id
 * @apiSuccess (data-list-pay_info元素对象字段说明) {integer}   charge_cards_id 支付通道id,只有公司入款才用到
 * @apiSuccess (data-list-pay_info元素对象字段说明) {array}     accounts 只有代理充值才有此字段
 * @apiSuccess (data-list-pay_info-accounts数组元素对象字段说明) {string}   accout 聊天账号
 * @apiSuccess (data-list-pay_info-accounts数组元素对象字段说明) {string}   name 名称
 * @apiSuccess (data-list-pay_info-accounts数组元素对象字段说明) {string}   url	账号头像
 * @apiSuccessExample {json} 响应结果
 * {
 *     "clientMsg": "",
 *     "code": 200,
 *     "data": {
 *         "cache_key": "222f830d0ca949b233e21c79e1a0cff8",
 *         "list": [
 *             {
 *                 "charge_numbers": [
 *                     50,
 *                     100,
 *                     200,
 *                     300,
 *                     500,
 *                     1000,
 *                     2000,
 *                     3000,
 *                     5000
 *                 ],
 *                 "logo": "https://s3.ap-northeast-2.amazonaws.com/qpgame/ckyx_823132_xrgv_.png",
 *                 "name": "银联扫码",
 *                 "pay_info": [
 *                     {
 *                         "addr_type": 4,
 *                         "bank_address": "",
 *                         "card_number": "",
 *                         "charge_cards_id": 395,
 *                         "credential_id": 14,
 *                         "hint": "无支付通道吗？",
 *                         "logo": "https://s3.ap-northeast-2.amazonaws.com/qpgame/ckyx_479813_mooe_20171.png",
 *                         "mfrom": 1,
 *                         "mto": 111111111,
 *                         "name": "",
 *                         "owner": "",
 *                         "title": "银联扫码222"
 *                     },
 *                     {
 *                         "addr_type": 2,
 *                         "bank_address": "",
 *                         "card_number": "",
 *                         "charge_cards_id": 393,
 *                         "credential_id": 69,
 *                         "hint": "银联扫码最新配置爱测试",
 *                         "logo": "https://s3.ap-northeast-2.amazonaws.com/qpgame/ckyx_208110_agsy_20170.gif",
 *                         "mfrom": 10,
 *                         "mto": 600000,
 *                         "name": "",
 *                         "owner": "",
 *                         "title": "银联扫码配置"
 *                     }
 *                 ],
 *                 "remark": "12122"
 *             },
 *             {
 *                 "charge_numbers": [
 *                     50,
 *                     100,
 *                     200,
 *                     300,
 *                     500,
 *                     1000,
 *                     2000,
 *                     3000,
 *                     5000
 *                 ],
 *                 "logo": "https://s3.ap-northeast-2.amazonaws.com/qpgame/ckyx_988029_mqcp_.png",
 *                 "name": "微信支付",
 *                 "pay_info": [
 *                     {
 *                         "addr_type": 4,
 *                         "bank_address": "wxpay2",
 *                         "card_number": "",
 *                         "charge_cards_id": 379,
 *                         "credential_id": 67,
 *                         "hint": "限额【10---1000】",
 *                         "logo": "https://s3.ap-northeast-2.amazonaws.com/qpgame/ckyx_935630_wxkh_weixin.png",
 *                         "mfrom": 10,
 *                         "mto": 100,
 *                         "name": "",
 *                         "owner": "",
 *                         "title": "RY支付"
 *                     },
 *                     {
 *                         "addr_type": 4,
 *                         "bank_address": "wxpay2",
 *                         "card_number": "",
 *                         "charge_cards_id": 390,
 *                         "credential_id": 70,
 *                         "hint": "固定金额支付 50 100 200 300",
 *                         "logo": "https://s3.ap-northeast-2.amazonaws.com/qpgame/ckyx_542510_ceci_weixin.png",
 *                         "mfrom": 50,
 *                         "mto": 300,
 *                         "name": "",
 *                         "owner": "",
 *                         "title": "蓉银支付-二维码-微信扫码2"
 *                     },
 *                     {
 *                         "addr_type": 6,
 *                         "bank_address": "wxpaywap",
 *                         "card_number": "",
 *                         "charge_cards_id": 29,
 *                         "credential_id": 67,
 *                         "hint": "支持金额10 20 30 50 100 200",
 *                         "logo": "S3_image_2019_04_17_2019041714252069091.png",
 *                         "mfrom": 10,
 *                         "mto": 200,
 *                         "name": "",
 *                         "owner": "",
 *                         "title": "ry微信wap支付"
 *                     },
 *                     {
 *                         "addr_type": 6,
 *                         "bank_address": "01",
 *                         "card_number": "",
 *                         "charge_cards_id": 31,
 *                         "credential_id": 68,
 *                         "hint": "支持金额10 20 30 50 100 200",
 *                         "logo": "S3_image_2019_04_17_2019041714385945850.png",
 *                         "mfrom": 10,
 *                         "mto": 200,
 *                         "name": "",
 *                         "owner": "",
 *                         "title": "ly微信h5"
 *                     },
 *                     {
 *                         "addr_type": 4,
 *                         "bank_address": "wxpay2",
 *                         "card_number": "",
 *                         "charge_cards_id": 380,
 *                         "credential_id": 70,
 *                         "hint": "限额【10---100】",
 *                         "logo": "https://s3.ap-northeast-2.amazonaws.com/qpgame/ckyx_758302_wrsg_weixin.png",
 *                         "mfrom": 10,
 *                         "mto": 100,
 *                         "name": "",
 *                         "owner": "",
 *                         "title": "RY微信支付"
 *                     }
 *                 ],
 *                 "remark": "扫码步骤：1，点“立即充值”将自动为您截屏并保存到相册，同时打开微信。2，请在微信中打开“扫一扫”。3，在扫一扫中点击右上角，选择“从相册选取二维码”选取截屏的图片。4，输入您欲充值的金额并进行转账。请务必备注上您的会员账户,以便我们能更好的为您充值入款. 如充值未及时到账，请联系在线客服。"
 *             },
 *             {
 *                 "charge_numbers": [
 *                     50,
 *                     100,
 *                     200,
 *                     300,
 *                     500,
 *                     1000,
 *                     2000,
 *                     3000,
 *                     5000
 *                 ],
 *                 "logo": "https://s3.ap-northeast-2.amazonaws.com/qpgame/ckyx_659734_xvud_.png",
 *                 "name": "支付宝支付",
 *                 "pay_info": [
 *                     {
 *                         "addr_type": 6,
 *                         "bank_address": "923",
 *                         "card_number": "",
 *                         "charge_cards_id": 388,
 *                         "credential_id": 106,
 *                         "hint": "喜多多支付-h5支付-支付宝扫码923",
 *                         "logo": "https://s3.ap-northeast-2.amazonaws.com/qpgame/ckyx_665929_oirz_.jpg",
 *                         "mfrom": 1,
 *                         "mto": 50000,
 *                         "name": "",
 *                         "owner": "",
 *                         "title": "喜多多支付-h5支付-支付宝扫码923"
 *                     },
 *                     {
 *                         "addr_type": 6,
 *                         "bank_address": "aliPaySM",
 *                         "card_number": "",
 *                         "charge_cards_id": 389,
 *                         "credential_id": 104,
 *                         "hint": "钉钉支付-h5支付-支付宝扫码",
 *                         "logo": "https://s3.ap-northeast-2.amazonaws.com/qpgame/ckyx_61985_hzwp_.jpg",
 *                         "mfrom": 1,
 *                         "mto": 50000,
 *                         "name": "",
 *                         "owner": "",
 *                         "title": "钉钉支付-h5支付-支付宝扫码"
 *                     },
 *                     {
 *                         "addr_type": 6,
 *                         "bank_address": "ZFB",
 *                         "card_number": "",
 *                         "charge_cards_id": 25,
 *                         "credential_id": 89,
 *                         "hint": "支持金额100-2000",
 *                         "logo": "S3_image_2019_04_17_2019041713560440564.png",
 *                         "mfrom": 100,
 *                         "mto": 2000,
 *                         "name": "",
 *                         "owner": "",
 *                         "title": "lx支付宝"
 *                     },
 *                     {
 *                         "addr_type": 6,
 *                         "bank_address": "ZFB",
 *                         "card_number": "",
 *                         "charge_cards_id": 43,
 *                         "credential_id": 89,
 *                         "hint": "支持金额300-5000",
 *                         "logo": "S3_image_2019_05_02_2019050216281958664.png",
 *                         "mfrom": 300,
 *                         "mto": 5000,
 *                         "name": "",
 *                         "owner": "",
 *                         "title": "lx支付宝一键支付"
 *                     },
 *                     {
 *                         "addr_type": 6,
 *                         "bank_address": "ZFB",
 *                         "card_number": "",
 *                         "charge_cards_id": 45,
 *                         "credential_id": 89,
 *                         "hint": "支持金额100-5000",
 *                         "logo": "https://s3.ap-northeast-2.amazonaws.com/qpgame/ckyx_591803_xrsu_zhifu.png",
 *                         "mfrom": 100,
 *                         "mto": 5000,
 *                         "name": "",
 *                         "owner": "",
 *                         "title": "lx支付宝扫码"
 *                     },
 *                     {
 *                         "addr_type": 5,
 *                         "bank_address": "02",
 *                         "card_number": "",
 *                         "charge_cards_id": 53,
 *                         "credential_id": 68,
 *                         "hint": "支持金额10 20 30 50 100 200",
 *                         "logo": "S3_image_2019_05_03_2019050309560930386.png",
 *                         "mfrom": 10,
 *                         "mto": 200,
 *                         "name": "",
 *                         "owner": "",
 *                         "title": "ly支付宝wap"
 *                     },
 *                     {
 *                         "addr_type": 4,
 *                         "bank_address": "ALIPAY_NATIVE",
 *                         "card_number": "",
 *                         "charge_cards_id": 382,
 *                         "credential_id": 29,
 *                         "hint": "支付宝固码支付",
 *                         "logo": "https://s3.ap-northeast-2.amazonaws.com/qpgame/ckyx_650563_mbig_zhifu.png",
 *                         "mfrom": 10,
 *                         "mto": 1000,
 *                         "name": "",
 *                         "owner": "支付宝固码支付",
 *                         "title": "支付宝固码支付"
 *                     }
 *                 ],
 *                 "remark": "扫码步骤：1，点“立即充值”将自动为您截屏并保存到相册，同时打开支付宝。2，请在支付宝中打开“扫一扫”。3，在扫一扫中点击右上角，选择“从相册选取二维码”选取截屏的图片。4，输入您欲充值的金额并进行转账。如充值未及时到账，请联系在线客服。"
 *             },
 *             {
 *                 "charge_numbers": [
 *                     50,
 *                     100,
 *                     200,
 *                     300,
 *                     500,
 *                     1000,
 *                     2000,
 *                     3000,
 *                     5000
 *                 ],
 *                 "logo": "https://s3.ap-northeast-2.amazonaws.com/qpgame/ckyx_959970_gmee_.png",
 *                 "name": "银行转账",
 *                 "pay_info": [
 *                     {
 *                         "addr_type": 1,
 *                         "bank_address": "深圳",
 *                         "card_number": "123456789",
 *                         "charge_cards_id": 391,
 *                         "credential_id": 1,
 *                         "hint": "请尽情的支付吧",
 *                         "logo": "https://s3.ap-northeast-2.amazonaws.com/qpgame/ckyx_95497_pdpp_20180.png",
 *                         "mfrom": 10,
 *                         "mto": 10000000,
 *                         "name": "中国银行",
 *                         "owner": "哦哦",
 *                         "title": ""
 *                     }
 *                 ],
 *                 "remark": "需要客户填写存入时间，存入金额，存款人姓名，还有选择转账的渠道，包括网银转账，ATM自动柜员机，ATM现金入款，银行柜台转账，手机银行转账，其他的方式"
 *             },
 *             {
 *                 "charge_numbers": [
 *                     50,
 *                     100,
 *                     200,
 *                     300,
 *                     500,
 *                     1000,
 *                     2000,
 *                     5000
 *                 ],
 *                 "logo": "https://s3.ap-northeast-2.amazonaws.com/qpgame/ckyx_692425_jhqq_.png",
 *                 "name": "银联扫码（云闪付）",
 *                 "pay_info": [
 *                     {
 *                         "addr_type": 6,
 *                         "bank_address": "11",
 *                         "card_number": "",
 *                         "charge_cards_id": 32,
 *                         "credential_id": 68,
 *                         "hint": "支持金额10-5000",
 *                         "logo": "S3_image_2019_04_17_2019041714414769785.png",
 *                         "mfrom": 10,
 *                         "mto": 5000,
 *                         "name": "",
 *                         "owner": "",
 *                         "title": "ly银联支付"
 *                     }
 *                 ],
 *                 "remark": "扫码步骤：1、首先我们在手机上下载云闪付，然后点击打开。2、登录云闪付后，接着我们要绑定一张银行卡。3、点“立即充值”将自动为您截屏并保存到相册，同时打开云闪付，点击“扫一扫”。4、然后我们将二维码或条形码放入框内，这时候会自动扫描，几秒钟后就好了。如充值未及时到账，请联系在线客服。"
 *             },
 *             {
 *                 "charge_numbers": [
 *                     50,
 *                     100,
 *                     200,
 *                     300,
 *                     500,
 *                     1000,
 *                     2000,
 *                     3000,
 *                     5000
 *                 ],
 *                 "logo": "https://s3.ap-northeast-2.amazonaws.com/qpgame/ckyx_56394_kyrg_.png",
 *                 "name": "代理充值",
 *                 "pay_info": [
 *                     {
 *                         "addr_type": 4,
 *                         "bank_address": "",
 *                         "card_number": "",
 *                         "charge_cards_id": 392,
 *                         "credential_id": 14,
 *                         "hint": "test哦test哦",
 *                         "logo": "https://s3.ap-northeast-2.amazonaws.com/qpgame/ckyx_143416_yggb_icon.jpg",
 *                         "mfrom": 10,
 *                         "mto": 100000,
 *                         "name": "",
 *                         "owner": "",
 *                         "title": "代理充值哦哦哦"
 *                     }
 *                 ],
 *                 "remark": "测测"
 *             },
 *             {
 *                 "charge_numbers": [],
 *                 "logo": "代理充值logo图片地址",
 *                 "name": "代理充值",
 *                 "pay_info": [
 *                     {
 *                         "accounts": [
 *                             {
 *                                 "account": "123",
 *                                 "name": "财务专员微信号",
 *                                 "url": "头像图片地址"
 *                             }
 *                         ],
 *                         "addr_type": 7,
 *                         "hint": "添加以下官方代理号,可在10秒内完成充值"
 *                     }
 *                 ]
 *             }
 *         ]
 *     },
 *     "internalMsg": "",
 *     "timeConsumed": 998
 * }
 */

func (cthis *FinanceController) ChargeTypesList() {
	ctx := cthis.ctx
	iUserId, _ := ctx.Values().GetInt("userid")
	username := ctx.Values().GetString("username")
	cacheKey := ctx.URLParam("cache_key")
	userIdCard, _ := ramcache.UserIdCard.Load(cthis.platform)
	uicMap := userIdCard.(map[int]beans.UserProfile)
	var res = make(map[string]interface{})
	userEntity, existUser := uicMap[iUserId]
	if !existUser {
		utils.ResFaiJSON(&ctx, "缓存中不存在"+username+"用户信息", "获取数据失败", config.PARAMERROR)
		return
	}
	list := cthis.getChargeTypes(userEntity.UserGroupId)
	byteCon, _ := json.Marshal(list)
	md5hash := fmt.Sprintf("%x", md5.Sum(byteCon))
	if cacheKey == md5hash {
		list = make([]map[string]interface{}, 0)
	} else {
		res["cache_key"] = md5hash
	}
	res["list"] = list
	utils.ResSuccJSON(&ctx, "", "", config.SUCCESSRES, res)
}
func (cthis *FinanceController) getChargeTypes(userGroupIdString string) []map[string]interface{} {
	var list = make([]map[string]interface{}, 0)
	chargeType, _ := ramcache.TableChargeTypes.Load(cthis.platform)
	bankCards := cthis.getBankcards(userGroupIdString)
	//pic 域名地址
	var picurls = ramcache.GetMainDbPlatform(cthis.platform)
	host := picurls.PicAddress
	//过滤之后的数据
	var orderPriorityAsc = make(map[int][]map[string]interface{})

	for _, temp := range chargeType.([]xorm.ChargeTypes) {
		single := make(map[string]interface{})
		id := temp.Id
		priority := temp.Priority
		var priorityArr = make([]map[string]interface{}, 0)
		priArr, existPriArr := orderPriorityAsc[priority]
		if existPriArr {
			priorityArr = priArr
		}
		if temp.State == 1 {
			var chargeMoneys interface{}
			errC := json.Unmarshal([]byte(temp.ChargeNumbers), &chargeMoneys)
			if errC != nil {
				log.LogPrException(fmt.Sprintf("获取充值列表解析字符串数组出错 %v", errC))
			}
			single["name"] = temp.Name
			single["remark"] = temp.Remark
			single["charge_numbers"] = chargeMoneys
			logo := temp.Logo
			if logo != "" && !strings.Contains(logo, "http") {
				logo = host + "/" + logo
			}
			single["logo"] = logo
			bkC, existBkC := bankCards[id]
			if existBkC {
				single["pay_info"] = bkC
			} else {
				single["pay_info"] = make([]map[string]interface{}, 0)
			}
			priorityArr = append(priorityArr, single)
		}
		orderPriorityAsc[priority] = priorityArr
	}
	//priority数组
	var orderPrioriAr = make([]int, 0)
	//把priority放入到数组进行排序
	for prioritySort, _ := range orderPriorityAsc {
		orderPrioriAr = append(orderPrioriAr, prioritySort)
	}
	//对priority进行asc排序
	sort.Ints(orderPrioriAr)
	for _, val := range orderPrioriAr {
		list = append(list, orderPriorityAsc[val]...)
	}
	proxyCharge := ramcache.GetConfigs(cthis.platform, "proxy_charge").(map[string]interface{})
	//如果配置了代理充值
	if int(proxyCharge["state"].(float64)) == 1 {
		singlePC := make(map[string]interface{})
		payInfo := make(map[string]interface{})
		singlePC["charge_numbers"] = []int{}
		singlePC["logo"] = proxyCharge["proxy_charge_logo"]
		singlePC["name"] = "代理充值"
		payInfo["hint"] = proxyCharge["info"]
		payInfo["addr_type"] = 7
		payInfo["accounts"] = proxyCharge["charge_accounts"]
		singlePC["pay_info"] = []map[string]interface{}{payInfo}
		list = append(list, singlePC)
	}
	return list
}

func (cthis *FinanceController) getBankcards(userGroupIdString string) map[int][]map[string]interface{} {
	/* 不走数据库实现如下sql查询效果
	sqlString:="select name,bank_address,card_number,owner,charge_type_id,logo,id,mfrom,mto,title,hint,addr_type,credential_id,"+
		"qr_code from charge_cards where state = 1"
	if userGroupIdString != "1" {
		sqlString += " and FIND_IN_SET("+userGroupIdString+",user_group_ids"
	}else{
		//不为1的看不到
		sqlString += " and user_group_ids = '1'"
	}
	sqlString += " order by priority asc,addr_type desc"
	*/
	//存放过滤之后的数据
	var orderPriorityAsc = make(map[int]map[int][]xorm.ChargeCards)
	//获取charge_cards表数据
	ChargeCards, _ := ramcache.TableChargeCards.Load(cthis.platform)
	//用户组字符串
	for _, tem := range ChargeCards.([]xorm.ChargeCards) {
		//切割UserLevelID字段为数组
		userGroupIds := strings.Split(tem.UserGroupIds, ",")
		//相同priority容器
		var priorArr = make(map[int][]xorm.ChargeCards)
		priorityMap, existPriorityMap := orderPriorityAsc[tem.Priority]
		//不存在就创建priority容器
		if !existPriorityMap {
			orderPriorityAsc[tem.Priority] = priorArr
			//存在就使用priorityMap
		} else {
			priorArr = priorityMap
		}

		//相同addr_type容器
		var adTyAr = make([]xorm.ChargeCards, 0)
		addTypeMap, existAddTypeMap := priorArr[tem.AddrType]
		//不存在就创建AddrType容器
		if !existAddTypeMap {
			priorArr[tem.AddrType] = adTyAr
			//存在就使用AddrTypeMap
		} else {
			adTyAr = addTypeMap
		}
		//相当于这个sql条件
		// "where state = 1 and (user_group_ids = 1 or FIND_IN_SET("+userGroupIdString+",user_group_ids)"
		//充值账号卡表中如果用户组是1表示所有用户可以看,也就是用户组Id不管是什么都在所有可见里面
		//所以在里面判断的时候一定要判断userGroupIds是否为1，如果为1谁都可以看,不为1的情况才过滤特定的组
		//所以这里要特别小心，容易不好理解
		if userGroupIdString != "1" && tem.State == 1 {
			if tem.UserGroupIds == "1" || utils.InArrayString(userGroupIdString, userGroupIds) {
				adTyAr = append(adTyAr, tem)
			}
		}
		//相当于这个sql条件
		// "where state = 1"
		if userGroupIdString == "1" && tem.State == 1 && tem.UserGroupIds == "1" {
			adTyAr = append(adTyAr, tem)
		}
		priorArr[tem.AddrType] = adTyAr
		orderPriorityAsc[tem.Priority] = priorArr
	}

	//存储 Priority asc，addType des 层级关系
	var orderAscDesc = make(map[int][]int)
	//Priority des 排序数组
	var sortPriority = make([]int, 0)
	for priority, val := range orderPriorityAsc {
		sortPriority = append(sortPriority, priority)
		//addType 排序数组
		var sortAddType = make([]int, 0)
		for adTyA, _ := range val {
			sortAddType = append(sortAddType, adTyA)
		}
		//addType进行desc排序
		sort.Sort(sort.Reverse(sort.IntSlice(sortAddType)))
		orderAscDesc[priority] = sortAddType
	}
	//对priority进行asc排序
	sort.Ints(sortPriority)
	//存放结果集
	var temp = make([]xorm.ChargeCards, 0)
	//数据进行priority asc,addType desc排序
	for _, vvv := range sortPriority {
		addTypeSorted := orderAscDesc[vvv]
		for _, v1 := range addTypeSorted {
			temp = append(temp, orderPriorityAsc[vvv][v1]...)
		}
	}
	//sqlString:="select name,bank_address,card_number,owner,charge_type_id,logo," +
	//	"id,mfrom,mto,title,hint,addr_type,credential_id,qrcode from charge_cards"
	var result = make(map[int][]map[string]interface{}, 0)
	//pic 域名地址
	var picurls = ramcache.GetMainDbPlatform(cthis.platform)
	host := picurls.PicAddress
	for _, BkCar := range temp {
		var single = make(map[string]interface{})
		var singleArray = make([]map[string]interface{}, 0)
		//复制map
		single["id"] = BkCar.Id
		single["name"] = BkCar.Name
		single["bank_address"] = BkCar.BankAddress
		single["card_number"] = BkCar.CardNumber
		single["owner"] = BkCar.Owner
		single["charge_cards_id"] = BkCar.Id
		single["mfrom"] = BkCar.Mfrom
		single["mto"] = BkCar.Mto
		single["title"] = BkCar.Title
		single["hint"] = BkCar.Hint
		single["addr_type"] = BkCar.AddrType
		single["credential_id"] = BkCar.CredentialId
		single["qrcode"] = BkCar.QrCode
		single["logo"] = BkCar.Logo
		typeid := BkCar.ChargeTypeId
		addrType := single["addr_type"]
		typeidArray, existTypeidMap := result[typeid]
		if !existTypeidMap {
			result[typeid] = singleArray
		} else {
			singleArray = typeidArray
		}
		qrcode := single["qrcode"].(string)
		if addrType != 3 {
			//如果不是在线支付，删除元素qrcode;兼容App版本
			delete(single, "qrcode")
		}
		bankAddress := single["bank_address"].(string)
		//公司入款支付二维码
		if addrType == 2 {
			bankAddress = strings.Trim(bankAddress, " \t")
			if bankAddress == "" {
				bankAddress = qrcode
			} else if bankAddress != "" && !strings.Contains(bankAddress, "http") {
				bankAddress = host + "/" + bankAddress
			}
		}
		if addrType == 3 && qrcode != "" {
			bankAddress += "/charge.html?tp=" + cthis.platform + "&id=" + strconv.Itoa(single["id"].(int))
			if !strings.Contains(qrcode, "http") {
				qrcode = host + "/" + qrcode
			}
			single["qrcode"] = qrcode
		}
		single["bank_address"] = bankAddress
		singleArray = append(singleArray, single)
		result[typeid] = singleArray
	}
	return result
}

/**
 * @api {get} api/auth/v1/redPacketQuery 红包查询接口
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人:pling把妹</span><br/><br/>
 * 红包查询接口<br>
 * 业务描述:在app启动进入首页的时候请求,切记如果查询失败，或者没有内容就不要做任何展示处理</br>
 * 打回的是红包列表，如果是多个，就要一个一个领取
 * @apiVersion 1.0.0
 * @apiName     api_auth_v1_redPacketQuery
 * @apiGroup    finance_module
 * @apiPermission ios,android客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 Bearer token
 * @apiSuccess (返回结果)  {int}     code            200成功
 * @apiSuccess (返回结果)  {string}  clientMsg       提示信息
 * @apiSuccess (返回结果)  {string}  internalMsg     提示信息
 * @apiSuccess (返回结果)  {json}  	  data            数据对象
 * @apiSuccess (返回结果)  {float}   time_consumed    后台耗时
 * @apiSuccess (data对象字段说明) {array}   list   记录对象数组
 * @apiSuccess (data-list元素对象字段说明)  {int}  	  created    时间戳
 * @apiSuccess (data-list元素对象字段说明)  {string}  	  charge_type_info    充值类型说明
 * @apiSuccess (data-list元素对象字段说明)  {decimal}  	  amount    充值金额
 * @apiSuccessExample {json} 响应结果
 * {
 *     "clientMsg": "",
 *     "code": 200,
 *     "data": {
 *         "list": [
 *             {
 *                 "id": 1
 *             }
 *         ]
 *     },
 *     "internalMsg": "",
 *     "timeConsumed": 7040
 * }
 */

func (cthis *FinanceController) RedPacketQuery() {
	ctx := cthis.ctx
	userIdS := ctx.Values().GetString("userid")
	userId, _ := strconv.Atoi(userIdS)
	redpacketReceive := make([]xorm.RedpacketReceives, 0)
	engine := models.MyEngine[cthis.platform]
	engine.Cols("id").Where("is_get = 0").And("user_id = ?", userId).Find(&redpacketReceive)
	var res = make(map[string]interface{})
	if len(redpacketReceive) == 0 {
		res["list"] = []string{}
		utils.ResSuccJSON(&ctx, "", "", config.SUCCESSRES, res)
		return
	}

	list := make([]map[string]interface{}, 0)
	for _, v := range redpacketReceive {
		var temp = make(map[string]interface{})
		temp["id"] = v.Id
		list = append(list, temp)
	}
	res["list"] = list
	utils.ResSuccJSON(&ctx, "", "", config.SUCCESSRES, res)
}

/**
 * @api {Post} api/auth/v1/redPacketReceiveAction 红包领取接口
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人:pling把妹</span><br/><br/>
 * 红包查询接口<br>
 * 业务描述:红包领取</br>
 * @apiVersion 1.0.0
 * @apiName     api_auth_v1_redPacketReceiveAction
 * @apiGroup    finance_module
 * @apiPermission ios,android客户端
 * @apiHeader (请求头) {string} Authorization 用户令牌  格式为 Bearer token
 * @apiParam (客户端请求参数) {int}    id          红包领取Id
 * @apiSuccess (返回结果)  {int}     code            200成功
 * @apiSuccess (返回结果)  {string}  clientMsg       提示信息
 * @apiSuccess (返回结果)  {string}  internalMsg     提示信息
 * @apiSuccess (返回结果)  {json}  	  data            数据对象
 * @apiSuccess (返回结果)  {float}   time_consumed    后台耗时
 * @apiSuccess (data对象字段说明) {float64}   money   领取金额
 * @apiSuccessExample {json} 响应结果
 * {
 *     "clientMsg": "",
 *     "code": 200,
 *     "data": {
 *         "money": 1.05
 *     },
 *     "internalMsg": "",
 *     "timeConsumed": 32873
 * }
 */

func (cthis *FinanceController) RedPacketReceiveAction() {
	ctx := cthis.ctx
	userIdS := ctx.Values().GetString("userid")
	userId, _ := strconv.Atoi(userIdS)
	id, errId := ctx.PostValueInt("id")
	if errId != nil || id < 1 {
		utils.ResFaiJSON(&ctx, "id错误", "", config.PARAMERROR)
		return
	}
	engine := models.MyEngine[cthis.platform]
	redPacketReceive := new(xorm.RedpacketReceives)
	engine.Where("is_get = 0").And("id = ?", id).And("user_id = ?", userId).Get(redPacketReceive)
	var res = make(map[string]interface{})
	if redPacketReceive.Id == 0 {
		utils.ResFaiJSON(&ctx, "记录不存在", "无法领取红包", config.PARAMERROR)
		return
	}
	money, _ := strconv.ParseFloat(redPacketReceive.Money, 64)
	balance := fund.NewUserFundChange(cthis.platform) //给用户充值
	info := map[string]interface{}{
		"user_id":     userId,
		"type_id":     config.FUNDREDPACKET,
		"amount":      money,
		"order_id":    utils.CreationOrder("RED", userIdS),
		"msg":         "领取红包",
		"finish_rate": 1.0, //需满足的打码量比例
	}
	//更新记录表
	callback := func(session *goxorm.Session, args ...interface{}) (interface{}, error) {
		redRe := new(xorm.RedpacketReceives)
		redRe.IsGet = 1
		_, errUpdateRow := session.ID(id).Cols("is_get").Update(redRe)

		var redTypeName = "红包"
		switch redPacketReceive.RedType {
		case 1:
			redTypeName = "节日红包"
		case 2:
			redTypeName = "每日幸运红包"
		}
		title := "领取红包通知"
		content := "成功领取" + redTypeName + redPacketReceive.Money + "，请查收！"
		noticeEntity := xorm.Notices{
			UserId:  userId,
			Title:   title,
			Content: content,
			Status:  1,
			Created: utils.GetNowTime(),
		}
		affNum, inErr := session.Insert(noticeEntity)
		if inErr != nil || affNum <= 0 {
			session.Rollback()
			return nil, inErr
		}
		return nil, errUpdateRow
	}
	balanceUpdateRes := balance.BalanceUpdate(info, callback)
	if balanceUpdateRes["status"] == 1 {
		res["money"] = money
		utils.ResSuccJSON(&ctx, "", "", config.SUCCESSRES, res)
		return
	}
	utils.ResFaiJSON(&ctx, "领取出现事物错误", "无法领取红包", config.PARAMERROR)
}
