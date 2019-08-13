package frontEndControllers

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/kataras/iris"
	"github.com/shopspring/decimal"
	"qpgame/app/fund"
	"qpgame/common/utils"
	"qpgame/common/utils/game"
	"qpgame/config"
	"qpgame/models"
	"qpgame/models/xorm"
	"qpgame/ramcache"
	"strconv"
)

type GameController struct {
	platform string
	ctx      iris.Context
}

//构造函数
func NewGameController(ctx iris.Context) *GameController {
	obj := new(GameController)
	obj.platform = ctx.Params().Get("platform")
	obj.ctx = ctx
	return obj
}

func (cthis *GameController) UpdateGameList() {
	ctx := cthis.ctx
	if !utils.RequiredParam(&ctx, []string{"platid"}) {
		return
	}
	platid, _ := strconv.Atoi(ctx.URLParam("platid"))
	if game.GetGameList(cthis.platform, platid) == nil {
		utils.ResSuccJSON(&ctx, "游戏更新成功", "success", config.SUCCESSRES, "{}")
	}
}

/**
 * @api {get} api/v1/getGameList 获取游戏列表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人:wenzhen</span><br/><br/>
 * 获取游戏列表<br>
 * 业务描述:获取游戏列表</br>
 * @apiVersion 1.0.0
 * @apiName     getGameList
 * @apiGroup    game
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
 *
 *
 * @apiSuccessExample {json} 响应结果
 *
 *{
 *    "clientMsg": "游戏列表获取成功",
 *    "code": 200,
 *    "data": {}
 *}
 */
func (cthis *GameController) GetGameList() {
	ctx := cthis.ctx
	var beans []xorm.PlatformGames
	err := models.MyEngine[cthis.platform].Find(&beans)
	if err != nil {
		utils.ResFaiJSON2(&ctx, err.Error(), "游戏列表获取失败")
		return
	}
	if len(beans) == 0 {
		checkNil(&ctx, nil)
	} else {
		checkNil(&ctx, beans)
	}
}

/**
 * @api {get} api/v1/getGameListSub 获取游戏分类列表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人:wenzhen</span><br/><br/>
 * 获取游戏列表<br>
 * 业务描述:获取游戏列表</br>
 * @apiVersion 1.0.0
 * @apiName     getGameListSub
 * @apiGroup    game
 * @apiPermission PC客户端
 * @apiParam (客户端请求参数) {string} game_categorie_id    	游戏分类编号

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
 *
 * @apiSuccessExample {json}  成功响应结果
 *{
 *    "clientMsg": "游戏列表获取成功",
 *    "code": 200,
 *    "data": [
 *        {
 *            "game_categorie_id": 8,
 *            "game_url": "",
 *            "gamecode": "25",
 *            "gt": "slot",
 *            "id": 2577,
 *            "img": "",
 *            "ishot": 0,
 *            "isnew": 0,
 *            "isrecommend": 0,
 *            "ishidden": 0,
 *            "name": "一本万利",
 *            "plat_id": 2,
 *            "service_code": "25",
 *            "small_img": ""
 *        },
 *        {
 *            "game_categorie_id": 8,
 *            "game_url": "",
 *            "gamecode": "5",
 *            "gt": "slot",
 *            "id": 2578,
 *            "img": "",
 *            "ishot": 0,
 *            "isnew": 0,
 *            "isrecommend": 0,
 *            "ishidden": 0,
 *            "name": "发发发",
 *            "plat_id": 2,
 *            "service_code": "5",
 *            "small_img": ""
 *        },
 *        ...
 *    "internalMsg": "",
 *    "timeConsumed": 0
 *}
 */
func (cthis *GameController) GetGameListSub() {
	ctx := cthis.ctx
	tmp, _ := ramcache.GameCache.Load(cthis.platform)
	gmc := tmp.([]map[string]interface{})
	game_categorie_id, err := ctx.URLParamInt("game_categorie_id")
	if err != nil {
		utils.ResFaiJSON(&ctx, err.Error(), "游戏分类编号错误", config.PARAMERROR)
		return
	}
	for _, v := range gmc {
		//如果是第一层
		if v["categoryLevel"] == 1 {
			for _, v2 := range v["categories"].([]map[string]interface{}) {
				if v2["id"] == game_categorie_id {
					utils.ResSuccJSON(&ctx, "", "游戏列表获取成功", config.SUCCESSRES, v2["games"])
					return
				}
			}
		}
	}
	utils.ResSuccJSON(&ctx, "", "该分类下暂无游戏", config.SUCCESSRES, []byte{})
}

/**
 * @api {get} api/v1/GetGameListNewCache 获取游戏分类列表
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人:wenzhen</span><br/><br/>
 * 获取游戏分类列表，如果缓存没变化将打回空数组，无cache_key字段<br>
 * 业务描述:获取游戏列表</br>
 * @apiVersion 1.0.0
 * @apiName     GetGameListNewCache
 * @apiGroup    game
 * @apiPermission PC客户端
 * @apiParam (客户端请求参数) {string} cache_key   缓存md5
 * @apiError (请求失败返回) {int}      code            错误代码
 * @apiError (请求失败返回) {string}   clientMsg       提示信息
 * @apiError (请求失败返回) {string}   internalMsg     开发内部错误代码
 * @apiError (请求失败返回) {float}    timeConsumed   后台耗时
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
 * @apiSuccess (返回结果)  {json}  	  data            json对象
 * @apiSuccess (返回结果)  {float}   timeConsumed    后台耗时
 * @apiSuccess (data数组元素字段说明)  {string}  	  cache_key     md5值
 * @apiSuccess (data数组元素字段说明)  {list}  	  list     数组
 * @apiSuccess (list数组元素字段说明)  {string}  	  btnImg     为选中按钮图片
 * @apiSuccess (list数组元素字段说明)  {int}  	  categoryLevel     分类等级
 * @apiSuccess (list数组元素字段说明)  {array}  	  games     游戏数组，没有下级分类才有此字段
 * @apiSuccess (list数组元素字段说明)  {int}  	  id    分类Id
 * @apiSuccess (list数组元素字段说明)  {int}  	  seq    排序
 * @apiSuccess (list数组元素字段说明)  {int}  	  status    是否可用(0不可用,1可用)
 * @apiSuccess (list数组元素字段说明)  {string}  	  selectedImg    选中图片
 * @apiSuccess (list数组元素字段说明)  {string}  	  name    分类名称
 * @apiSuccess (list数组元素字段说明)  {int}  	  platformStatus    平台状态，0不可用，1可用,2维护中，3敬请期待
 * @apiSuccessExample {json} 成功响应结果
 {
    "clientMsg": "游戏列表获取成功",
    "code": 200,
    "data": [
        {
            "btnImg": "",
            "categoryLevel": 2,
            "games": [
                {
                    "game_categorie_id": 6,
                    "game_url": "https://h5.ppro.98078.net/game?type=h5&gamecode=hlhb&language=zh-cn",
                    "gamecode": "hlhb",
                    "gt": "poker",
                    "id": 2420,
                    "img": "https://static.ppro.98078.net/global/files/images/2019012222072762340.png",
                    "ishot": 1,
                    "isnew": 1,
                    "isrecommend": 1,
                    "ishidden": 0,
                    "name": "欢乐红包",
                    "plat_id": 1,
                    "service_code": "6506",
                    "small_img": "https://static.ppro.98078.net/global/files/images/2019012222073122678.png",
                    "platformStatus": 1
                },
				....
            ],
            "id": 1,
            "name": "热门游戏",
            "platformStatus": 1,
            "selectedImg": "",
            "seq": 1,
            "status": 1
        },
        {
            "btnImg": "",
            "categories": [
                {
                    "btnImg": "https://s3.ap-northeast-2.amazonaws.com/qpgame/ckyx_530136_ahxc_fgqp_.png",
                    "id": 6,
                    "img": "https://s3.ap-northeast-2.amazonaws.com/qpgame/ckyx_221757_fqhi_FG.png",
                    "name": "FG棋牌",
                    "platformStatus": 1,
                    "selectedImg": "https://s3.ap-northeast-2.amazonaws.com/qpgame/ckyx_390125_jaxs_fgqp_.png",
                    "seq": 0,
                    "status": 1
                },
				........
            ],
            "categoryLevel": 1,
            "id": 2,
            "name": "棋牌游戏",
            "platformStatus": 1,
            "selectedImg": "",
            "seq": 2,
            "status": 1
        },
        {
            "btnImg": "",
            "categoryLevel": 2,
            "games": [
                {
                    "game_categorie_id": 3,
                    "game_url": "https://h5.ppro.98078.net/game?type=h5&gamecode=fish_3D&language=zh-cn",
                    "gamecode": "fish_3D",
                    "gt": "fish",
                    "id": 2416,
                    "img": "44444fdsfds",
                    "ishot": 0,
                    "isnew": 1,
                    "isrecommend": 1,
                    "ishidden": 0,
                    "name": "FG捕鱼来了3D",
                    "plat_id": 1,
                    "service_code": "5006",
                    "small_img": "https://s3.ap-northeast-2.amazonaws.com/qpgame/ckyx_10848_usfn_fg.png",
                    "platformStatus": 1
                },
				........
            ],
            "id": 3,
            "name": "捕鱼游戏",
            "platformStatus": 1,
            "selectedImg": "",
            "seq": 3,
            "status": 1
        },
        {
            "btnImg": "",
            "categories": [
                {
                    "btnImg": "https://s3.ap-northeast-2.amazonaws.com/qpgame/ckyx_723543_wzkw_fg_un.png",
                    "id": 7,
                    "img": "https://s3.ap-northeast-2.amazonaws.com/qpgame/ckyx_538251_swhb_FG.png",
                    "name": "FG电子",
                    "platformStatus": 1,
                    "selectedImg": "https://s3.ap-northeast-2.amazonaws.com/qpgame/ckyx_667570_knuf_fg_se.png",
                    "seq": 0,
                    "status": 1
                },
              .........
            ],
            "categoryLevel": 1,
            "id": 4,
            "name": "电子游艺",
            "platformStatus": 1,
            "selectedImg": "",
            "seq": 4,
            "status": 1
        },
        {
            "btnImg": "",
            "categoryLevel": 2,
            "games": [
                {
                    "game_categorie_id": 5,
                    "game_url": "11",
                    "gamecode": "SMG_titaniumLiveGames_Baccarat",
                    "gt": "live",
                    "id": 2614,
                    "img": "",
                    "ishot": 0,
                    "isnew": 0,
                    "isrecommend": 0,
                    "ishidden": 0,
                    "name": "MG视讯",
                    "plat_id": 3,
                    "service_code": "SMG_titaniumLiveGames_Baccarat",
                    "small_img": "https://s3.ap-northeast-2.amazonaws.com/qpgame/ckyx_519873_upmd_MG1.png",
                    "platformStatus": 1
                },
				.........
            ],
            "id": 5,
            "name": "真人视讯",
            "platformStatus": 1,
            "selectedImg": "",
            "seq": 5,
            "status": 1
        },
        {
            "btnImg": "https://s3.ap-northeast-2.amazonaws.com/qpgame/ckyx_244340_wbcj_S3_im.png",
            "categoryLevel": 2,
            "games": [
                {
                    "game_categorie_id": 27,
                    "game_url": "",
                    "gamecode": "UG",
                    "gt": "sport",
                    "id": 3648,
                    "img": "测试8888",
                    "ishot": 0,
                    "isnew": 0,
                    "isrecommend": 0,
                    "ishidden": 0,
                    "name": "联合体育",
                    "plat_id": 11,
                    "service_code": "UnitedGaming",
                    "small_img": "https://s3.ap-northeast-2.amazonaws.com/qpgame/ckyx_641836_ikaa_%201%20.png",
                    "platformStatus": 1
                }
            ],
            "id": 27,
            "name": "体育赛事",
            "platformStatus": 1,
            "selectedImg": "https://s3.ap-northeast-2.amazonaws.com/qpgame/ckyx_590752_yffd_S3_im.png",
            "seq": 0,
            "status": 1
        }
    ],
    "internalMsg": "",
    "timeConsumed": 33
}
*/
func (cthis *GameController) GetGameListNewCache() {
	ctx := cthis.ctx
	cacheKey := ctx.URLParam("cache_key")
	tmp, _ := ramcache.GameCache.Load(cthis.platform)
	gmc := make([]map[string]interface{}, 0)
	var res = make(map[string]interface{})
	for _, v := range tmp.([]map[string]interface{}) {
		ttmap := make(map[string]interface{})
		for kk, vv := range v {
			ttmap[kk] = vv
		}
		//如果是第一层，就清除分类下面的游戏列表
		if v["categoryLevel"] == 1 {
			cmap := make([]map[string]interface{}, 0)
			for _, v2 := range v["categories"].([]map[string]interface{}) {
				tmap := make(map[string]interface{})
				for k, v3 := range v2 {
					if k != "games" {
						tmap[k] = v3
					}
				}
				cmap = append(cmap, tmap)
			}
			ttmap["categories"] = cmap
		}
		gmc = append(gmc, ttmap)
	}
	res["list"] = gmc
	byteCon, _ := json.Marshal(gmc)
	md5hash := fmt.Sprintf("%x", md5.Sum(byteCon))
	//如果没变化就打回空数组
	if cacheKey == md5hash {
		res["list"] = []string{}
	} else {
		res["cache_key"] = md5hash
	}
	utils.ResSuccJSON(&ctx, "", "游戏列表获取成功", config.SUCCESSRES, res)
	return
}

/**
 * @api {get} api/auth/v1/launchGame 获取启动游戏url
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人:wenzhen</span><br/><br/>
 * 启动游戏<br>
 * 业务描述:获取启动游戏url</br>
 * @apiVersion 1.0.0
 * @apiName     launchGame
 * @apiGroup    game
 * @apiPermission iso,android客户端
 * @apiHeader (客户端请求头参数) {string} Authorization Bearer + 用户登录获得的token

 * @apiError (请求失败返回) {int}      code            错误代码
 * @apiError (请求失败返回) {string}   clientMsg       提示信息
 * @apiError (请求失败返回) {string}   internalMsg     错误代码
 * @apiError (请求失败返回) {float}    timeConsumed   后台耗时
 * @apiParam (客户端请求参数) {string} platid    	游戏平台编号 FG=1
 * @apiParam (客户端请求参数) {string} gamecode    	游戏编号
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
 *
 *
 * @apiSuccessExample {json} 响应结果
 *
 *{
 *    "clientMsg": "success",
 *    "code": 200,
 *    "data": "https://h5.ppro.98078.net/game?type=h5&gamecode=ShowHand&language=zh-cn&token=54A493202EFA073B",
 *    "internalMsg": "获取游戏启动地址",
 *    "timeConsumed": 794508
 *}
 */
func (cthis *GameController) LaunchGame() {
	ctx := cthis.ctx
	if !utils.RequiredParam(&ctx, []string{"platid", "gamecode"}) {
		return
	}

	platid, _ := ctx.URLParamInt("platid")
	// @updated by aTian 平台状态不可用时，返回提示信息
	platsConf, _ := ramcache.TablePlatforms.Load(cthis.platform)
	var platStatus = 0
	for _, platConf := range platsConf.([]xorm.Platforms) {
		if platConf.Id == platid {
			platStatus = platConf.Status
			break
		}
	}
	if platStatus != 1 {
		utils.ResFaiJSON(&ctx, "", "该游戏正在维护或者未上线", config.NOTGETDATA)
		return
	}

	userIdS := ctx.Values().GetString("userid")
	userId, _ := ctx.Values().GetInt("userid")
	gamecode := ctx.URLParam("gamecode")
	var plataccounts = xorm.PlatformAccounts{UserId: userId, PlatId: platid}
	var gameurl string
	//首先判断该用户在对应平台有没有对应的游戏账号
	models.MyEngine[cthis.platform].Get(&plataccounts)
	//没有则通过接口创建一个用户
	if plataccounts.Id == 0 {
		pAccs, b := game.CreatePlay(userIdS, cthis.platform, platid)
		if !b {
			utils.ResFaiJSON(&ctx, "", "创建游戏账号失败", config.NOTGETDATA)
			return
		}
		platAcc, _ := ramcache.TablePlatformAccounts.Load(cthis.platform)
		plataccounts = pAccs
		//更新缓存
		platAcc.(map[string]int)[plataccounts.Username] = userId
		ramcache.TablePlatformAccounts.Store(cthis.platform, platAcc)
	}
	//开始获取游戏启动
	gameurl = game.GetGameUrl(&plataccounts, gamecode, utils.GetIp(ctx.Request()), platid, cthis.platform)
	balance := fund.NewUserFundChange(cthis.platform)
	res := balance.BeforeLaunchGameFundChange(userId, platid, &plataccounts) //将玩家余额存进对应平台
	if res["status"] != 1 {
		utils.ResFaiJSON(&ctx, "1905311327", "自动转账失败", config.NOTGETDATA)
		return
	}
	utils.ResSuccJSON(&ctx, "获取游戏启动地址", "success", config.SUCCESSRES, gameurl)
}

/**
 * @api {get} api/auth/v1/exitGame 退出游戏
 * @apiDescription
 * <span style="color:lightcoral;">接口负责人:wenzhen</span><br/><br/>
 * 退出游戏<br>
 * 业务描述:退出游戏</br>
 * @apiVersion 1.0.0
 * @apiName     exitGame
 * @apiGroup    game
 * @apiPermission iso,android客户端
 * @apiHeader (客户端请求头参数) {string} Authorization Bearer + 用户登录获得的token
 * @apiParam (客户端请求参数) {string} platid    	游戏平台编号 FG=1
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
 *
 *
 *{
 *    "clientMsg": "退出成功",
 *    "code": 200,
 *    "data": {
 *        "Id": 2,
 *        "UserId": 28,
 *        "ChargedAmount": "0.000",
 *        "ConsumedAmount": "0.000",
 *        "Balance": "0.000",
 *        "WithdrawAmount": "0.000",
 *        "BalanceCharge": "0.000",
 *        "BalanceLucky": "0.000",
 *        "BalanceSafe": "8200.000",
 *        "BalanceWallet": "1794.11",
 *        "Updated": 1553595076
 *    },
 *    "internalMsg": "退出游戏",
 *    "timeConsumed": 1631129
 *}
 */
func (cthis *GameController) ExitGame() {
	ctx := cthis.ctx
	if !utils.RequiredParam(&ctx, []string{"platid"}) {
		return
	}
	userId, _ := ctx.Values().GetInt("userid")
	platId, _ := ctx.URLParamInt("platid")
	var platAccounts = xorm.PlatformAccounts{UserId: userId, PlatId: platId}
	models.MyEngine[cthis.platform].Get(&platAccounts)
	platBalance, success := game.QueryUchips(platAccounts.Username, platId, cthis.platform)
	if !success {
		utils.ResFaiJSON(&ctx, "1905311328", "自动转账失败", config.NOTGETDATA)
		return
	}
	balance, _ := decimal.NewFromString(platBalance)
	userBalance := fund.NewUserFundChange(cthis.platform)
	amountFloat, _ := balance.Float64()
	res := userBalance.AfterExitGameFundChange(userId, platId, &platAccounts, amountFloat) //将玩家余额存进对应平台
	if res["status"] != 1 {
		utils.ResFaiJSON(&ctx, "1905311329", "自动转账失败", config.NOTGETDATA)
		return
	}
	utils.ResSuccJSON(&ctx, "退出游戏", "退出成功", config.SUCCESSRES, res["accounts"])
}
