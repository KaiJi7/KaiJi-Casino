package strategy

import (
	"KaiJi-Casino/internal/pkg/db"
	"KaiJi-Casino/internal/pkg/db/collection"
	"KaiJi-Casino/internal/pkg/strategy/lowerResponse"
	"KaiJi-Casino/internal/pkg/strategy/lowestResponse"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var NameMap = map[Name]struct {
	Description string
	Generator   func(data collection.StrategyData) Strategy
}{
	LowerResponse:  {Description: "Bet each games with lower odds.", Generator: lowerResponse.New},
	LowestResponse: {Description: "Bet a game with the lowest odds.", Generator: lowestResponse.New},
}

func InitStrategy(strategyId *primitive.ObjectID) (strategy Strategy, err error) {
	log.Debug("init strategy: ", strategyId.Hex())

	strategyData, err := db.New().GetStrategy(strategyId)
	if err != nil {
		log.Error("fail to init strategy: ", err.Error())
		return
	}

	if content, exist := NameMap[strategyData.Name]; !exist {
		log.Error("unsupported strategy: ", strategyData.Name)
		err = fmt.Errorf("unsupported strategy")
		return
	} else {
		strategy = content.Generator(strategyData)
	}
	return
}
