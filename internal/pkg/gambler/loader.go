package gambler

import (
	"KaiJi-Casino/internal/pkg/db"
	"KaiJi-Casino/internal/pkg/strategy"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetGamblers(simulationId *primitive.ObjectID) (gamblers []Gambler, err error) {
	log.Debug("get gamblers by simulation id: ", simulationId.Hex())

	gamblersData, dbErr := db.New().ListGambler(simulationId)
	if dbErr != nil {
		log.Error("fail to load gambler: ", dbErr.Error())
		err = dbErr
		return
	}

	for _, gamblerData := range gamblersData {
		strategyData, dbErr := db.New().GetStrategyData(gamblerData.Id)
		if dbErr != nil {
			log.Error("fail to get strategyData: ", dbErr.Error())
			err = dbErr
			return
		}
		stg, sErr := strategy.GenStrategy(strategyData)
		if sErr != nil {
			log.Error("gail to init strategy: ", sErr.Error())
			err = sErr
			return
		}

		gamblers = append(gamblers, Gambler{
			GamblerData: gamblerData,
			Strategy:    stg,
		})
	}
	return
}
