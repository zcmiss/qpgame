/**
展示相关的控制器，例如公告，优惠活动等
*/
package frontEndControllers

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/kataras/iris"
	"qpgame/app/fund"
	"qpgame/common/utils"
	"qpgame/config"
	"qpgame/models"
	"qpgame/models/xorm"
	"qpgame/ramcache"
	"strconv"
	"strings"
)

type ShowController struct {
	platform string
	ctx      iris.Context
}

//构造函数
func NewShowController(ctx iris.Context) *ShowController {
	obj := new(ShowController)
	obj.platform = ctx.Params().Get("platform")
	obj.ctx = ctx
	return obj
}

/**
 * @api {get} api/v1/getActivity 获取优惠活动信息
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人:wenzhen</span><br/><br/>
 * 获取优惠活动<br>
 * 业务描述:获取优惠活动</br>
 * @apiVersion 1.0.0
 * @apiName     getActivity
 * @apiGroup    showinfo
 * @apiPermission PC客户端
 * @apiParam (客户端请求参数) {int} activityClassId    	活动分类ID

 * @apiError (请求失败返回) {int}      code            错误代码
 * @apiError (请求失败返回) {string}   clientMsg       提示信息
 * @apiError (请求失败返回) {string}   internalMsg     错误代码
 * @apiError (请求失败返回) {float}    timeConsumed   后台耗时
 *
 * @apiErrorExample {json} 失败返回
 * {
 *      "code": 204,
 *      "internalMsg": "",
 *      "clientMsg ": 0,
 *      "timeConsumed": 0
 * }
 *
 * @apiSuccess (返回结果)  {int}      code            200
 * @apiSuccess (返回结果)  {string}   clientMsg       提示信息
 * @apiSuccess (返回结果)  {string}   internalMsg     提示信息
 * @apiSuccess (返回结果)  {array}  	  data         返回数据
 * @apiSuccess (返回结果)  {float}   timeConsumed    后台耗时
* @apiSuccess (data-list数组元素对象字段说明)   {int}  	  Id				活动编号
* @apiSuccess (data-list数组元素对象字段说明)   {string}  	  Title				活动标题
* @apiSuccess (data-list数组元素对象字段说明)   {string}  	  SubTitle			活动子标题
* @apiSuccess (data-list数组元素对象字段说明)   {string}  	  Content			活动内容
* @apiSuccess (data-list数组元素对象字段说明)   {string}  	  Icon				活动图标
* @apiSuccess (data-list数组元素对象字段说明)   {int}  	  Created			创建时间
* @apiSuccess (data-list数组元素对象字段说明)   {int}  	  TimeStart			活动开始时间时间戳
* @apiSuccess (data-list数组元素对象字段说明)   {int}  	  TimeEnd			活动结束时间时间戳
* @apiSuccess (data-list数组元素对象字段说明)   {int}  	  Type				活动类型
* @apiSuccess (data-list数组元素对象字段说明)   {int}  	  Status			活动状态
* @apiSuccess (data-list数组元素对象字段说明)   {int}  	  Updated			最后更新时间时间戳
* @apiSuccess (data-list数组元素对象字段说明)   {int}  	  IpLimit			IP限制总次数
* @apiSuccess (data-list数组元素对象字段说明)   {int}  	  IpLimitDay		IP当天限制次数
* @apiSuccess (data-list数组元素对象字段说明)   {string}  	  AwardCondition	活动奖励条件
* @apiSuccess (data-list数组元素对象字段说明)   {int}  	  TotalAward			总奖励次数
* @apiSuccess (data-list数组元素对象字段说明)   {int}  	  ActivityClassId		活动分类编号
* @apiSuccess (data-list数组元素对象字段说明)   {int}  	  IsMultipleApply		是否允许多次申请(0为否，1为是)
* @apiSuccess (data-list数组元素对象字段说明)   {int}  	  IsHomeShow			是否首页弹出显示(0为否,1为是)
 *
 * @apiSuccessExample {json} 响应结果
 *
 {
    "clientMsg": "获取成功",
    "code": 200,
    "data": [
        {
            "Id": 1,
            "Title": "活动标题",
            "SubTitle": "子标题，这时子标题",
            "Content": "活动说明的内容",
            "Icon": "",
            "Created": 0,
            "TimeStart": 0,
            "TimeEnd": 0,
            "Type": 0,
            "Status": 1,
            "Updated": 0,
            "IpLimit": 0,
            "IpLimitDay": 0,
            "AwardCondition": "",
            "TotalAward": 3,
            "ActivityClassId": 1,
            "IsMultipleApply": 0,
            "IsHomeShow": 0
        }
    ],
    "internalMsg": "获取成功",
    "timeConsumed": 16956
}
*/
func (cthis *ShowController) GetActivity() {
	ctx := cthis.ctx
	if !utils.RequiredParam(&ctx, []string{"activityClassId"}) {
		return
	}
	activityClassId := ctx.URLParam("activityClassId")
	var beans []xorm.Activities
	err := models.MyEngine[cthis.platform].Where("status=1 and activity_class_id=?", activityClassId).Find(&beans)
	if err != nil {
		utils.ResFaiJSON2(&ctx, err.Error(), "优惠活动列表获取失败")
	}
	if len(beans) == 0 {
		checkNil(&ctx, nil)
	} else {
		checkNil(&ctx, beans)
	}
}

/**
 * @api {get} api/v1/vipDetail 个人中心vip详情
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人:Pling把妹</span><br/><br/>
 * 获取vip详情<br>
 * 业务描述:个人中心的vip详情，这个接口返回html文本，请用webview打开</br>
 * @apiVersion 1.0.0
 * @apiName     api_v1_vipDetail
 * @apiGroup    showinfo
 * @apiPermission ios,android客户端
 * @apiSuccessExample {json} 响应结果
	html文本
*/
func (cthis *ShowController) GetVipDetail() {
	ctx := cthis.ctx
	tvipL, _ := ramcache.TableVipLevels.Load(cthis.platform)
	vipLevels := tvipL.([]xorm.VipLevels)
	ctx.ViewData("Title", "VIP详情")
	ctx.ViewLayout("layout.html")
	lastVip := vipLevels[len(vipLevels)-1]
	ctx.ViewData("viplevels", vipLevels)
	ctx.ViewData("WeeklyAmount", lastVip.WeeklyAmount)
	ctx.ViewData("MonthAmount", lastVip.MonthAmount)
	ctx.ViewData("UpgradeAmountTotal", lastVip.UpgradeAmountTotal)
	ctx.Gzip(true)
	ctx.View("vip_detail.html")

}

/**
 * @api {get} api/v1/proxyDetail 推广中心代理详情
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人:Pling把妹</span><br/><br/>
 * 获取vip详情<br>
 * 业务描述:推广中心的代理详情，这个接口返回html文本，请用webview打开</br>
 * @apiVersion 1.0.0
 * @apiName     api_v1_proxyDetail
 * @apiGroup    showinfo
 * @apiPermission ios,android客户端
 * @apiSuccessExample {json} 响应结果
	html文本
*/
func (cthis *ShowController) GetProxyDetail() {
	ctx := cthis.ctx
	proChLeve, _ := ramcache.TableProxyChessLevels.Load(cthis.platform)
	proRealLevel, _ := ramcache.TableProxyRealLevels.Load(cthis.platform)
	Chess := proChLeve.([]xorm.ProxyChessLevels)
	Real := proRealLevel.([]xorm.ProxyRealLevels)
	newChess := make([]xorm.ProxyChessLevels, len(Chess))
	for i, j := 0, len(Chess)-1; i < j; i, j = i+1, j-1 {
		newChess[i], newChess[j] = Chess[j], Chess[i]
	}
	newReal := make([]xorm.ProxyRealLevels, len(Real))
	for i, j := 0, len(Real)-1; i < j; i, j = i+1, j-1 {
		newReal[i], newReal[j] = Real[j], Real[i]
	}
	ctx.ViewData("Title", "代理详情")
	ctx.ViewLayout("layout.html")
	ctx.ViewData("Chess", newChess)
	ctx.ViewData("Real", newReal)
	ctx.Gzip(true)
	ctx.View("proxy_detail.html")
}

/**
 * @api {get} api/v1/getActivityClass 获取优惠活动分类信息
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人:wenzhen</span><br/><br/>
 * 获取优惠活动分类<br>
 * 业务描述:获取优惠活动分类</br>
 * @apiVersion 1.0.0
 * @apiName     getActivityClass
 * @apiGroup    showinfo
 * @apiPermission PC客户端

 * @apiError (请求失败返回) {int}      code            错误代码
 * @apiError (请求失败返回) {string}   clientMsg       提示信息
 * @apiError (请求失败返回) {string}   internalMsg     错误代码
 * @apiError (请求失败返回) {float}    timeConsumed   后台耗时
 *
 * @apiErrorExample {json} 失败返回
 * {
 *      "code": 204,
 *      "internalMsg": "",
 *      "clientMsg ": 0,
 *      "timeConsumed": 0
 * }
 *
 * @apiSuccess (返回结果)  {int}      code            200
 * @apiSuccess (返回结果)  {string}   clientMsg       提示信息
 * @apiSuccess (返回结果)  {string}   internalMsg     提示信息
 * @apiSuccess (返回结果)  {json}  	  data            返回数据
 * @apiSuccess (返回结果)  {float}   timeConsumed    后台耗时
* @apiSuccess (data-list数组元素对象字段说明)   {int}  	  Id				活动类型编号
* @apiSuccess (data-list数组元素对象字段说明)   {string}  	  Name				活动类型名称

 *
 * @apiSuccessExample {json} 响应结果
 *
{
    "clientMsg": "获取成功",
    "code": 200,
    "data": [
        {
            "Id": 1,
            "Name": "综合活动",
            "Status": 1,
            "Seq": 0,
            "Created": 0,
            "Updated": 0
        },
        {
            "Id": 2,
            "Name": "棋牌活动",
            "Status": 1,
            "Seq": 0,
            "Created": 0,
            "Updated": 0
        },
        {
            "Id": 3,
            "Name": "捕鱼活动",
            "Status": 1,
            "Seq": 0,
            "Created": 0,
            "Updated": 0
        },
        {
            "Id": 4,
            "Name": "电子活动",
            "Status": 1,
            "Seq": 0,
            "Created": 0,
            "Updated": 0
        },
        {
            "Id": 5,
            "Name": "视讯活动",
            "Status": 1,
            "Seq": 0,
            "Created": 0,
            "Updated": 0
        },
        {
            "Id": 6,
            "Name": "体育活动",
            "Status": 1,
            "Seq": 0,
            "Created": 0,
            "Updated": 0
        }
    ],
    "internalMsg": "获取成功",
    "timeConsumed": 15956
}
*/
func (cthis *ShowController) GetActivityClass() {
	ctx := cthis.ctx

	var beans []xorm.ActivityClasses
	err := models.MyEngine[cthis.platform].Where("status=?", 1).Desc("seq").Find(&beans)
	if err != nil {
		utils.ResFaiJSON2(&ctx, err.Error(), "优惠活动列表获取失败")
		return
	}
	if len(beans) == 0 {
		checkNil(&ctx, nil)
	} else {
		checkNil(&ctx, beans)
	}
}

/**
 * @api {post} api/auth/v1/getActivityAward 领取活动奖励
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人:atian</span><br/><br/>
 * 领取活动奖励<br>
 * 业务描述:登录用户，根据活动状态、IP、活动类型等条件领取活动奖励。</br>
 * @apiVersion 1.0.0
 * @apiName     GetActivityAward
 * @apiGroup    showinfo
 * @apiPermission ios,android客户端

 * @apiError (请求失败返回) {int}      code            错误代码
 * @apiError (请求失败返回) {string}   clientMsg       提示信息
 * @apiError (请求失败返回) {string}   internalMsg     错误代码
 * @apiError (请求失败返回) {float}    timeConsumed    后台耗时
 *
 * @apiErrorExample {json} 失败返回
 * {
 *     "clientMsg": "您领取该活动的次数已达上限，无法继续领取",
 *     "code": 204,
 *     "internalMsg": "",
 *     "timeConsumed": 685588
 * }
 *
 * @apiSuccess (返回结果)  {int}      code            200
 * @apiSuccess (返回结果)  {string}   clientMsg       提示信息
 * @apiSuccess (返回结果)  {string}   internalMsg     提示信息
 * @apiSuccess (返回结果)  {json}     data            返回数据
 * @apiSuccess (返回结果)  {float}    timeConsumed    后台耗时
 *
 * @apiSuccessExample {json} 响应结果
 * {
 *     "clientMsg": "活动奖励领取成功",
 *     "code": 200,
 *     "data": {
 *         "exception": {},
 *         "msg": "交易成功~",
 *         "status": 1
 *     },
 *     "internalMsg": "success",
 *     "timeConsumed": 20568089
 * }
 *
 */
func (cthis *ShowController) GetActivityAward() {
	ctx := cthis.ctx
	if !utils.RequiredParamPost(&ctx, []string{"actId"}) {
		return
	}
	sUserId := ctx.Values().GetString("userid")
	var userBean xorm.Users
	engine := models.MyEngine[cthis.platform]
	engine.Where("id=?", sUserId).Cols("name", "status").Get(&userBean)
	// 用户锁定退出
	if userBean.Status != 1 {
		utils.ResFaiJSON2(&ctx, "", "当前用户已锁定，无法参与活动")
		return
	}
	sActId := ctx.URLParam("actId")
	// 首先判断活动ID是否存在
	var actBean xorm.Activities
	actExist, actErr := engine.Where("id=?", sActId).Cols("type", "status", "total_ip_limit", "day_ip_limit", "money", "is_repeat").Get(&actBean)
	if actErr != nil {
		utils.ResFaiJSON2(&ctx, actErr.Error(), "优惠活动获取失败")
		return
	}
	// 活动不存在退出
	if !actExist {
		utils.ResFaiJSON2(&ctx, "", "优惠活动不存在")
		return
	}
	// 活动失效退出
	if actBean.Status != 1 {
		utils.ResFaiJSON2(&ctx, "", "优惠活动已经失效")
		return
	}

	// 判断用户是否参与该活动
	actRecordExist, actRecordErr := engine.Table(new(xorm.ActivityRecords)).Where("user_id=? and activity_id=?", sUserId, sActId).Exist()
	if actRecordErr != nil {
		utils.ResFaiJSON2(&ctx, actRecordErr.Error(), "获取当前用户参与的活动记录失败")
		return
	}

	iUserId, _ := strconv.Atoi(sUserId)
	sIp := utils.GetIp(ctx.Request())
	if actRecordExist {
		// 已经参与过该活动，根据活动参与条件判断是否可以继续参与
		// 不可以重复领取
		if actBean.IsRepeat != 1 {
			utils.ResFaiJSON2(&ctx, "", "该活动彩金只能领取一次，您已经领取过该活动的奖励")
			return
		}

		// 计算当日的起始时间戳
		iFromTime, iToTime := utils.GetDatetimeRange(0, 1)
		iIpTotalCnt := 0 // 通过Ip统计参与指定活动的总数
		iIpDayCnt := 0

		var actRecordBeans []xorm.ActivityRecords
		actRecordsErr := engine.Where("activity_id=? and ip_addr=?", sActId, sIp).Find(&actRecordBeans)
		if actRecordsErr != nil {
			utils.ResFaiJSON2(&ctx, actRecordsErr.Error(), "获取用户参与的活动记录失败")
			return
		}

		if len(actRecordBeans) > 0 {
			for _, actRecordBean := range actRecordBeans {
				iIpTotalCnt++
				iCreated := int64(actRecordBean.Created)
				if iFromTime < iCreated && iCreated <= iToTime {
					iIpDayCnt++
				}
			}
		}
		// 判断当前Ip领取某一活动的总次数是否超出限制
		if iIpTotalCnt >= actBean.TotalIpLimit {
			utils.ResFaiJSON2(&ctx, "", "领取该活动的次数已达上限，无法继续领取")
			return
		}
		// 判断当前Ip当日领取某一活动是否超出限制
		if iIpDayCnt >= actBean.DayIpLimit {
			utils.ResFaiJSON2(&ctx, "", "今日领取该活动的次数已达上限，请明天再来")
			return
		}
	}

	// 没有参与该活动或者满足领取条件的，记录参与信息并领取相应奖励
	iActivityId, _ := strconv.Atoi(sActId)
	sRemark := strings.TrimSpace(ctx.FormValue("remark"))
	iNow := utils.GetNowTime()
	var actRecordBean = xorm.ActivityRecords{
		UserId:     iUserId,
		ActivityId: iActivityId,
		Remark:     sRemark,
		State:      0,
		Applied:    iNow,
		Created:    iNow,
		Updated:    iNow,
		IpAddr:     sIp,
	}

	// 保存领取记录，并获取奖励
	autoGetAward := func() {
		session := engine.NewSession()
		defer session.Close()
		sessErr := session.Begin()
		if sessErr != nil {
			utils.ResFaiJSON2(&ctx, sessErr.Error(), "活动彩金领取失败")
			return
		}
		actRecordAddNum, _ := session.Insert(actRecordBean)
		if actRecordAddNum <= 0 {
			session.Rollback()
			utils.ResFaiJSON2(&ctx, "", "活动彩金领取失败")
			return
		}
		fAwardMoney, _ := strconv.ParseFloat(actBean.Money, 64)
		if fAwardMoney >= 0 {
			awardInfo := map[string]interface{}{
				"user_id":     iUserId,
				"type_id":     config.FUNDBROKERAGE,
				"amount":      fAwardMoney,
				"order_id":    utils.CreationOrder("YH", sUserId),
				"msg":         "活动彩金领取",
				"finish_rate": 1.0, //需满足的打码量比例
			}

			balance := fund.NewUserFundChange(cthis.platform)
			balanceUpdateRes := balance.BalanceUpdate(awardInfo, nil)
			if balanceUpdateRes["status"] == 1 {
				session.Commit()
				utils.ResSuccJSON(&ctx, "success", "活动彩金领取成功，彩金已经存入您的账户余额", config.SUCCESSRES, balanceUpdateRes)
				return
			} else {
				session.Rollback()
				utils.ResFaiJSON2(&ctx, balanceUpdateRes["msg"].(string), "活动彩金领取失败")
				return
			}
		} else {
			session.Rollback()
			utils.ResFaiJSON2(&ctx, "", "活动彩金领取失败")
			return
		}
	}

	manualGetAward := func() {
		actRecordAddNum, _ := engine.Insert(actRecordBean)
		if actRecordAddNum > 0 {
			utils.ResSuccJSON(&ctx, "success", "参与优惠活动成功, 请联系客服领取彩金", config.SUCCESSRES, make([]interface{}, 0))
			return
		} else {
			utils.ResFaiJSON2(&ctx, "", "活动彩金领取失败")
			return
		}
	}

	switch actBean.Type {
	case 1:
		autoGetAward()
	default:
		manualGetAward()
	}
}

/**
 * @api {get} api/v1/getSystemNotice 获取系统公告通知
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人:wenzhen</span><br/><br/>
 * 获取系统公告信息,只打回最新的30条数据,请前端要缓存公告内容,如果30条数据没有变化会打回空数组并且没有cache_key字段<br>
 * 业务描述:请根据cache_key是否存在来判断数据是否有变化</br>
 * @apiVersion 1.0.0
 * @apiName     api_v1_getSystemNotice
 * @apiGroup    showinfo
 * @apiPermission ios,android客户端
 * @apiParam (客户端请求参数) {string} cache_key   缓存md5
 * @apiSuccess (返回结果)  {int}      code            200
 * @apiSuccess (返回结果)  {string}   clientMsg       提示信息
 * @apiSuccess (返回结果)  {string}   internalMsg     提示信息
 * @apiSuccess (返回结果)  {json}  	  data            返回数据
 * @apiSuccess (返回结果)  {float}   timeConsumed    后台耗时
 * @apiSuccess (data对象字段说明) {string}   cache_key md5值,如果数据没有变化或者没有公告内容无此字段
 * @apiSuccess (data对象字段说明) {array}   list 如果没有缓存或者没有数据为空数组
 * @apiSuccess (data-list对象字段说明) {string}   content 内容只有换行一个操作符\n来换行
 * @apiSuccess (data-list对象字段说明) {int}   createTime 时间戳,请前端格式化,减少服务器资源
 * @apiSuccess (data-list对象字段说明) {string}   title 内容标题
 * @apiSuccessExample {json} 响应结果
{
    "clientMsg": "",
    "code": 200,
    "data": {
        "cache_key": "ee43605143e6ddf22badcb7f672bce3b",
        "list": [
            {
                "content": "今日充值首选【网银充值】即送 2% 充值优惠，笔笔存笔笔送，隔天还有神秘彩金相送哦！\\n操作步骤：\\n①下载对应的手机网银控件登陆您的手机银行→②返回APP官网首页点击充值后选择【网银充值】→③点击前往充值按钮进入收款银行信息页面→④复制收款银行信息粘贴到您的手机银行→⑤输入充值金额按照转账步骤完成转账→⑥转账过后返回APP官网提交存款信息→⑦返回官网首页刷新额度等待到账即可！",
                "createTime": 1555324808,
                "title": "今日优惠"
            },
            {
                "content": "充值首选【网银充值】即送 1% 充值优惠，笔笔存笔笔送，隔天还有神秘彩金相送哦！感谢您长期对【383棋牌】支持与信任！祝各位老铁们好运连连！\\n【领取返佣】昨日有投注的会员请注意哦，代理返佣佣金已经结算完成，还未领取的会员，请记得领取哦！",
                "createTime": 1555324729,
                "title": "【CK棋牌】"
            },
            {
                "content": "充值首选【网银充值】即送 2% 充值优惠，笔笔存笔笔送，隔天还有神秘彩金相送哦！\\n操作步骤：\\n①下载对应的手机网银控件登陆您的手机银行→②返回APP官网首页点击充值后选择【网银充值】→③点击前往充值按钮进入收款银行信息页面→④复制收款银行信息粘贴到您的手机银行→⑤输入充值金额按照转账步骤完成转账→⑥转账过后返回APP官网提交存款信息→⑦返回官网首页刷新额度等待到账即可！",
                "createTime": 1555324669,
                "title": "充值推荐"
            }
        ]
    },
    "internalMsg": "",
    "timeConsumed": 49
}
*/
func (cthis *ShowController) GetSystemNotice() {
	ctx := cthis.ctx
	var beans []xorm.SystemNotices
	cacheKey := ctx.URLParam("cache_key")
	tsn, _ := ramcache.TableSystemNotices.Load(cthis.platform)
	beans = tsn.([]xorm.SystemNotices)
	var res = make(map[string]interface{})
	var list = make([]map[string]interface{}, 0)
	for _, v := range beans {
		var element = make(map[string]interface{})
		element["title"] = v.Title
		element["content"] = v.Content
		element["createTime"] = v.Created
		list = append(list, element)
	}
	res["list"] = list
	byteCon, _ := json.Marshal(list)
	md5hash := fmt.Sprintf("%x", md5.Sum(byteCon))
	//如果没变化就打回空数组
	if cacheKey == md5hash {
		res["list"] = []string{}
	} else {
		res["cache_key"] = md5hash
	}
	utils.ResSuccJSON(&ctx, "", "", config.SUCCESSRES, res)
}

/**
 * @api {get} api/v1/getHomeFirstSysNotice 获取首页第一条系统公告
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人:wenzhen</span><br/><br/>
 * app启动的时候进入首页之后去获取最新的系统公告,也就是系统公告的第一条数据,单独成为一个接口的考虑是基于带宽，以及更好的体验<br>
 * 业务描述:我们与383有所不同，首次获取之后需要保存到本地，如果有更新内容就展示给用户，没有更新内容就不展示，否则每次进来都展示</br>
 * 体验非常的差,缓存是否有更新请根据是否存在cache_key来判断
 * @apiVersion 1.0.0
 * @apiName     api_v1_getHomeFirstSysNotice
 * @apiGroup    showinfo
 * @apiPermission ios,android客户端
 * @apiParam (客户端请求参数) {string} cache_key   缓存md5
 * @apiSuccess (返回结果)  {int}      code            200
 * @apiSuccess (返回结果)  {string}   clientMsg       提示信息
 * @apiSuccess (返回结果)  {string}   internalMsg     开发调试提示信息
 * @apiSuccess (返回结果)  {json}  	   data            返回数据
 * @apiSuccess (返回结果)  {float}    timeConsumed    后台耗时
 * @apiSuccess (data对象字段说明) {string}   cache_key md5值,如果数据没有变化或者没有公告内容无此字段
 * @apiSuccess (data对象字段说明) {json}   notice 如果没有缓存或者没有数据为空对象
 * @apiSuccess (data-notice对象字段说明) {string}   content 内容只有换行一个操作符\n来换行
 * @apiSuccess (data-notice对象字段说明) {int}   createTime 时间戳,请前端格式化,减少服务器资源
 * @apiSuccess (data-notice对象字段说明) {string}   title 内容标题
 * @apiSuccessExample {json} 响应结果
{
    "clientMsg": "",
    "code": 200,
    "data": {
        "cache_key": "ee43605143e6ddf22badcb7f672bce3b",
        "notice":{
                "content": "今日充值首选【网银充值】即送 2% 充值优惠，笔笔存笔笔送，隔天还有神秘彩金相送哦！\\n操作步骤：\\n①下载对应的手机网银控件登陆您的手机银行→②返回APP官网首页点击充值后选择【网银充值】→③点击前往充值按钮进入收款银行信息页面→④复制收款银行信息粘贴到您的手机银行→⑤输入充值金额按照转账步骤完成转账→⑥转账过后返回APP官网提交存款信息→⑦返回官网首页刷新额度等待到账即可！",
                "createTime": 1555324808,
                "title": "今日优惠"
		}
    },
    "internalMsg": "",
    "timeConsumed": 49
}
*/
func (cthis *ShowController) GetHomeFirstSysNotice() {
	ctx := cthis.ctx
	var beans []xorm.SystemNotices
	cacheKey := ctx.URLParam("cache_key")
	tsn, _ := ramcache.TableSystemNotices.Load(cthis.platform)
	beans = tsn.([]xorm.SystemNotices)
	var res = make(map[string]interface{})
	var notice = make(map[string]interface{})
	v := beans[0]
	notice["title"] = v.Title
	notice["content"] = v.Content
	notice["createTime"] = v.Created
	byteCon, _ := json.Marshal(notice)
	md5hash := fmt.Sprintf("%x", md5.Sum(byteCon))

	//如果没变化就打回空数组
	if cacheKey == md5hash {
		notice = make(map[string]interface{})
	} else {
		res["cache_key"] = md5hash
	}
	res["notice"] = notice
	utils.ResSuccJSON(&ctx, "", "", config.SUCCESSRES, res)
}

/**
 * @api {get} api/auth/v1/getNotice 获取站内信
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人:wenzhen</span><br/><br/>
 * 获取站内信<br>
 * 业务描述:获取站内信</br>
 * @apiVersion 1.0.0
 * @apiName     api_auth_v1_getNotice
 * @apiGroup    showinfo
  * @apiPermission iso,android客户端
 * @apiHeader (客户端请求头参数) {string} Authorization Bearer + 用户登录获得的token
  * @apiSuccess (返回结果)  {int}      code            200
 * @apiSuccess (返回结果)  {string}   clientMsg       提示信息
 * @apiSuccess (返回结果)  {string}   internalMsg     提示信息
 * @apiSuccess (返回结果)  {json}  	  data            返回数据
 * @apiSuccess (返回结果)  {float}   timeConsumed    后台耗时
 * @apiSuccess (data对象字段说明) {string}   content 内容只有换行一个操作符\n来换行
 * @apiSuccess (data对象字段说明) {int}   createTime 时间戳,请前端格式化,减少服务器资源
 * @apiSuccess (data对象字段说明) {string}   title 内容标题
 * @apiSuccessExample {json} 响应结果
{
    "clientMsg": "",
    "code": 200,
    "data": [
        {
            "content": "今日充值推荐使用手机网银充值，玩家充值成功后即可添加微信客服【qipai383383】获取5%的入款回馈！限时活动，每个会员每天仅限申请一次，还没申请的玩家们抓紧啦！",
            "create": 1555324808,
            "title": "❤充值优惠❤"
        },
        {
            "content": "成功领取彩金:1.8元",
            "create": 1555324729,
            "title": "签到领取成功"
        },
        {
            "content": "【383棋牌限时活动】今日充值推荐使用【银行转账】入款即可获得10%入款优惠，充值后可直接添加微信客服申请，每个会员只限一次，即充即返，充值越高赠送就越高！超高无上限（只限今日）！",
            "create": 1555324669,
            "title": "【383棋牌】"
        }
    ],
    "internalMsg": "",
    "timeConsumed": 7123
}
*/
func (cthis *ShowController) GetNotice() {
	ctx := cthis.ctx
	var beans []xorm.Notices
	userid, _ := strconv.Atoi(ctx.Values().GetString("userid"))
	engine := models.MyEngine[cthis.platform]
	err := engine.Where("status = ? and user_id = ?", 1, userid).Limit(50).Desc("created").Find(&beans)
	var res = make([]map[string]interface{}, 0)
	if err != nil {
		utils.ResFaiJSON2(&ctx, err.Error(), "站内信列表获取失败")
		return
	}

	for _, v := range beans {
		var tempV = make(map[string]interface{})
		tempV["title"] = v.Title
		tempV["content"] = v.Content
		tempV["create"] = v.Created
		res = append(res, tempV)
	}
	utils.ResSuccJSON(&ctx, "", "", config.SUCCESSRES, res)
}

func checkNil(ctx *iris.Context, beans interface{}) {
	if beans == nil {
		utils.ResSuccJSON(ctx, "获取成功", "获取成功,但是没有数据", config.SUCCESSRES, make([]interface{}, 0))
		return
	}
	utils.ResSuccJSON(ctx, "获取成功", "获取成功", config.SUCCESSRES, beans)
}
