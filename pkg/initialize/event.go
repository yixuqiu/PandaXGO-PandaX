package initialize

import (
	"encoding/json"
	"pandax/apps/device/entity"
	"pandax/apps/device/services"
	ruleEntity "pandax/apps/rule/entity"
	"pandax/pkg/cache"
	"pandax/pkg/events"
	"pandax/pkg/global"
	"pandax/pkg/rule_engine"
	"pandax/pkg/tool"
)

// 初始化事件监听
func InitEvents() {
	// 监听规则链改变 更新所有绑定改规则链的产品
	global.EventEmitter.On(events.ProductChainRuleEvent, func(ruleId, codeData string) {
		global.Log.Infof("规则链%s变更", ruleId)
		list := services.ProductModelDao.FindList(entity.Product{
			RuleChainId: ruleId,
		})
		if list != nil {
			var lfData ruleEntity.RuleDataJson
			err := tool.StringToStruct(codeData, &lfData)
			if err != nil {
				global.Log.Error("规则链序列化失败", err)
				return
			}
			code, err := json.Marshal(lfData.LfData.DataCode)
			if err != nil {
				global.Log.Error("规则链序列化失败", err)
				return
			}
			//新建规则链实体
			instance, errs := rule_engine.NewRuleChainInstance(ruleId, code)
			if len(errs) > 0 {
				global.Log.Error("规则链初始化失败", errs[0])
				return
			}
			for _, product := range *list {
				cache.PutProductRule(product.Id, instance)
			}
		}
	})
}
