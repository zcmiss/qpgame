package validations

import (
	"qpgame/common/mvc"
)

type VipLevelsValidation struct{}

// 添加/修改动作数据验证
func (self VipLevelsValidation) Validate(ctx *Context) (string, bool) {
	return mvc.NewValidation(ctx).
		IsNumeric("level", "等级例子(1-10)必须为数字").
		StringLength("name", "等级名称长度在2-30之间", 2, 30).
		IsNumeric("valid_bet_min", "有效投注金额区间起点单位万必须为数字").
		IsNumeric("valid_bet_max", "有效投注金额区间封顶单位万必须为数字").
		IsNumeric("upgrade_amount", "晋级礼金必须为数字").
		IsNumeric("weekly_amount", "周礼金必须为数字").
		IsNumeric("month_amount", "月俸禄必须为数字").
		IsNumeric("upgrade_amount_total", "累计晋级礼金必须为数字").
		IsNumeric("has_deposit_speed", "存款加速通道(0不支持,1支持)必须为数字").
		IsNumeric("has_own_service", "专属客服经理(0没有,1有)必须为数字").
		IsDecimal("wash_code", "洗码率必须为合法的数字").
		Validate()
}
