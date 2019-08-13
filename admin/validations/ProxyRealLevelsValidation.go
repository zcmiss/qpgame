package validations

import (
	"qpgame/common/mvc"
)

type ProxyRealLevelsValidation struct{}

// 添加/修改动作数据验证
func (self ProxyRealLevelsValidation) Validate(ctx *Context) (string, bool) {
	return mvc.NewValidation(ctx).
		IsNumeric("level", "等级例子(1-10)必须为数字").
		StringLength("name", "等级名称长度在2-30之间", 2, 30).
		IsNumeric("team_total_low", "团队起始资金单位万/天必须为数字").
		IsNumeric("team_total_limit", "团队起始资金单位万封顶单位万/天必须为数字").
		IsNumeric("commission", "万/返佣必须为数字").
		Validate()
}
