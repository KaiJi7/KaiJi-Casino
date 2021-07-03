package strategy

import (
	"KaiJi-Casino/internal/pkg/db"
	"KaiJi-Casino/internal/pkg/strategy/common"
	"KaiJi-Casino/internal/pkg/strategy/confidenceBase"
	"KaiJi-Casino/internal/pkg/strategy/lowerResponse"
	"KaiJi-Casino/internal/pkg/strategy/lowestResponse"
	"KaiJi-Casino/internal/pkg/strategy/mostConfidence"
	"fmt"
	"github.com/KaiJi7/common/structs"
	log "github.com/sirupsen/logrus"
)

var NameMap = map[structs.StrategyName]func(data structs.StrategyData) common.Strategy{
	structs.StrategyNameLowerResponse:  lowerResponse.New,
	structs.StrategyNameLowestResponse: lowestResponse.New,
	structs.StrategyNameConfidenceBase: confidenceBase.New,
	structs.StrategyNameMostConfidence: mostConfidence.New,
}

func GenStrategy(strategyData structs.StrategyData) (strategy common.Strategy, err error) {
	log.Debug("get strategy: ", strategyData.Id.Hex())

	if generator, exist := NameMap[strategyData.Name]; !exist {
		log.Error("unsupported strategy: ", strategyData.Name)
		err = fmt.Errorf("unsupported strategy")
		return
	} else {
		strategy = generator(strategyData)
	}
	return
}

func CreateMetaStrategy(metaStrategy structs.StrategyMeta) (err error) {
	log.Debug("create meta strategy: ", metaStrategy.Name)

	if err = db.New().CreateMetaStrategy(metaStrategy); err != nil {
		log.Error("fail to create meta strategy: ", err.Error())
	}
	return
}
