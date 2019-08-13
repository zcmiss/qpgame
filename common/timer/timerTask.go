package timer

import (
	"github.com/robfig/cron"
	"os"
	"qpgame/common/timer/task"
	"qpgame/common/utils/game"
	"qpgame/config"
)

// 定时任务文件
func InitTimerTask() {
	c := cron.New()
	//每1分钟采集一次FG平台的投注记录
	//秒 分 时 日 月 星期
	c.AddFunc("1 */1 * * * *", func() {
		for k, _ := range config.PlatformCPs {
			fg := game.GetFG(k)
			gts := []string{"poker", "fish", "slot", "fruit"}
			for _, v := range gts {
				go fg.GetBets(v)
			}
		}
	})
	//每5分钟采集一次AE平台的投注记录
	c.AddFunc("20 */5 * * * *", func() {
		for k, _ := range config.PlatformCPs {
			ea := game.GetAe(k)
			go ea.GetBets()
		}
	})
	//每2分钟采集一次MG平台的投注记录
	c.AddFunc("10 */5 * * * *", func() {
		for k, _ := range config.PlatformCPs {
			mg := game.GetMg(k)
			go mg.GetBets()
		}
	})
	//每5分钟采集一次开元平台的投注记录
	c.AddFunc("30 */5 * * * *", func() {
		for k, _ := range config.PlatformCPs {
			ky := game.GetKy(k)
			go ky.GetBets()
		}
	})
	//每5分钟采集一次乐游平台的投注记录
	c.AddFunc("40 */5 * * * *", func() {
		for k, _ := range config.PlatformCPs {
			ly := game.GetLy(k)
			go ly.GetBets()
		}
	})

	//每5分钟采集一次新世界平台的投注记录
	c.AddFunc("50 */5 * * * *", func() {
		for k, _ := range config.PlatformCPs {
			nw := game.GetNW(k)
			go nw.GetBets()
		}
	})
	//每10分钟采集一次AG平台的投注记录
	c.AddFunc("0 */10 * * * *", func() {
		for k, _ := range config.PlatformCPs {
			ag := game.GetAg(k)
			go ag.GetBets()
		}
	})

	//每3分钟采集一次VG平台的投注记录
	c.AddFunc("35 */3 * * * *", func() {
		for k, _ := range config.PlatformCPs {
			vg := game.GetVG(k)
			go vg.GetBets()
		}
	})

	//每5分钟采集一次JDB平台的投注记录
	c.AddFunc("0 */5 * * * *", func() {
		for k, _ := range config.PlatformCPs {
			jdb := game.GetJdb(k)
			go jdb.GetBets()
		}
	})

	//每15分钟采集一次OG平台的投注记录
	c.AddFunc("15 */5 * * * *", func() {
		for k, _ := range config.PlatformCPs {
			og := game.GetOg(k)
			go og.GetBets()
		}
	})

	//每5分钟采集一次UG平台的投注记录
	c.AddFunc("25 */5 * * * *", func() {
		for k, _ := range config.PlatformCPs {
			ug := game.GetUg(k)
			go ug.GetBets()
		}
	})

	// 每10秒钟检查一次平台需要处理的，转账异常的任务
	c.AddFunc("*/10 * * * * *", func() {
		for k, _ := range config.PlatformCPs {
			go task.ExceptionTasks(k)
		}
	})
	//启动计划任务
	c.Start()
	//关闭这计划任务, 但是不能关闭已经在执行中的任务.
	defer c.Stop()
	select {}
}

// 初始化代理佣金派发
func InitProxyTimerTask() {
	//这里的定时器必须是平台唯一的,这个定时器会在其中的一台备用服务器上跑
	//测试的时候千万不要注入ISCOMMONTIMERSERVER 环境变量，否则后果很严重
	if os.Getenv("ISCOMMONTIMERSERVER") != "yes" {
		return
	}
	c := cron.New()
	//每天早上7点定时刷新代理佣金信息
	c.AddFunc("0 0 7 * * *", func() {
		for k, _ := range config.PlatformCPs {
			go task.CommissionPayment(k)
		}
	})
	//每5钟执行刷新打码量
	c.AddFunc("0 */5 * * * *", func() {
		for k, _ := range config.PlatformCPs {
			go task.UpdateWithdrawDamaRecords(k)
		}
	})

	//每5钟执行刷新今天投注总量
	c.AddFunc("0 */5 * * * *", func() {
		for k, _ := range config.PlatformCPs {
			go task.UpdateAccountTodayBetAmount(k)
		}
	})

	//每天2点执行刷新投注总额
	c.AddFunc("0 0 2 * * *", func() {
		for k, _ := range config.PlatformCPs {
			go task.UpdateAccountTotalBetAmount(k)
		}
	})

	//每周一5点执行派发VIP周工资
	c.AddFunc("0 0 5 * * 1", func() {
		for k, _ := range config.PlatformCPs {
			go task.VipWeekWage(k)
		}
	})

	//每月1号6点执行派发VIP月工资
	c.AddFunc("0 0 6 1 * *", func() {
		for k, _ := range config.PlatformCPs {
			go task.VipMonthWage(k)
		}
	})
	//启动计划任务
	c.Start()
	//关闭这计划任务, 但是不能关闭已经在执行中的任务.
	defer c.Stop()
	select {}
}
