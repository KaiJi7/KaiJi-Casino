package strategy

import (
	"KaiJi-Casino/internal/pkg/db/collection"
	"KaiJi-Casino/internal/pkg/strategy/common"
	"KaiJi-Casino/internal/pkg/strategy/lowerResponse"
	"KaiJi-Casino/internal/pkg/strategy/lowestResponse"
	"fmt"
	log "github.com/sirupsen/logrus"
)

var NameMap = map[collection.StrategyName]struct {
	Description string
	Generator   func(data collection.StrategyData) common.Strategy
}{
	collection.StrategyNameLowerResponse:  {Description: "Bet each games with lower odds.", Generator: lowerResponse.New},
	collection.StrategyNameLowestResponse: {Description: "Bet a game with the lowest odds.", Generator: lowestResponse.New},
}

func GetStrategy(strategyData collection.StrategyData) (strategy common.Strategy, err error) {
	log.Debug("get strategy: ", strategyData.Id.Hex())

	if content, exist := NameMap[strategyData.Name]; !exist {
		log.Error("unsupported strategy: ", strategyData.Name)
		err = fmt.Errorf("unsupported strategy")
		return
	} else {
		strategy = content.Generator(strategyData)
	}
	return
}
