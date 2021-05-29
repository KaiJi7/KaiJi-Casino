package gambler

import (
	"KaiJi-Casino/internal/pkg/db"
	"KaiJi-Casino/internal/pkg/db/collection"
	"KaiJi-Casino/internal/pkg/strategy"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Gambler struct {
	collection.GamblerData
	Strategy strategy.Strategy
}

func (g Gambler) MakeDecision(gambles []collection.Gambling) []collection.Decision {
	return nil
}

func (g Gambler) OnBroken() {
	log.Info("gambler: ", g.Id.Hex(), ". was broken.")
	return
}

func LoadGambler(simulationId *primitive.ObjectID) (gamblers []Gambler, err error) {

	gamblersData, err := db.New().ListGambler(simulationId)
	if err != nil {
		log.Error("fail to load gambler: ", err.Error())
		return
	}

	for _, gamblerData := range gamblersData {
		strategyData, err := db.New().GetStrategy(gamblerData.SimulationId)
		if err != nil {
			log.Error("fail to get strategyData: ", err.Error())
			return
		}
		stg, err := strategy.InitStrategy(strategyData.Id)
		if err != nil {
			log.Error("gail to init strategy: ", err.Error())
			return
		}

		gamblers = append(gamblers, Gambler{
			GamblerData: gamblerData,
			Strategy:    stg,
		})
	}
	return
}
