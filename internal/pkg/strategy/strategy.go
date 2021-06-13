package strategy

import (
	"KaiJi-Casino/internal/pkg/db"
	"KaiJi-Casino/internal/pkg/db/collection"
	"KaiJi-Casino/internal/pkg/strategy/common"
	"KaiJi-Casino/internal/pkg/strategy/confidenceBase"
	"KaiJi-Casino/internal/pkg/strategy/lowerResponse"
	"KaiJi-Casino/internal/pkg/strategy/lowestResponse"
	"KaiJi-Casino/internal/pkg/strategy/mostConfidence"
	"fmt"
	log "github.com/sirupsen/logrus"
)

var NameMap = map[collection.StrategyName]func(data collection.StrategyData) common.Strategy{
	collection.StrategyNameLowerResponse:  lowerResponse.New,
	collection.StrategyNameLowestResponse: lowestResponse.New,
	collection.StrategyNameConfidenceBase: confidenceBase.New,
	collection.StrategyNameMostConfidence: mostConfidence.New,
}

func GenStrategy(strategyData collection.StrategyData) (strategy common.Strategy, err error) {
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

func CreateMetaStrategy(metaStrategy collection.StrategyMeta) (err error) {
	log.Debug("create meta strategy: ", metaStrategy.Name)

	if err = db.New().CreateMetaStrategy(metaStrategy); err != nil {
		log.Error("fail to create meta strategy: ", err.Error())
	}
	return
}
