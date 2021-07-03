package casino

import (
	"KaiJi-Casino/internal/pkg/db"
	"KaiJi-Casino/internal/pkg/gambler"
	"KaiJi-Casino/internal/pkg/strategy"
	"github.com/KaiJi7/common/structs"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sync"
)

var Gamblers []gambler.Gambler

//var strategies []structs.StrategyData

func CreateGamblers(simulation structs.Simulation) (err error) {
	log.Debug("create gamblers: ", simulation.String())

	for strategyName, schema := range simulation.StrategySchema {
		_, exist := strategy.NameMap[strategyName]
		if !exist {
			log.Warn("invalid strategy: ", strategyName)
			return
		}

		for _, s := range schema {
			for i := 0; i < s.Quantity; i++ {
				gbl, dbErr := db.New().CreateGambler(simulation.Id, simulation.GamblerInitialMoney)
				if dbErr != nil {
					log.Error("fail to create gambler: ", dbErr.Error())
					err = dbErr
					return
				}

				meta, dbErr := db.New().GetStrategyMetaData(strategyName)
				if dbErr != nil {
					log.Error("fail to get strategy meta data: ", dbErr.Error())
					err = dbErr
					return
				}

				if _, dbErr = db.New().CreateStrategy(gbl.Id, strategyName, meta.Id, s.Properties); dbErr != nil {
					log.Error("fail to create strategy: ", dbErr.Error())
					err = dbErr
					return
				}
			}
		}
	}

	Gamblers, err = gambler.GetGamblers(simulation.Id)
	return
}

func LoadGamblers(simulationId string) (err error) {
	log.Debug("load gamblers with simulation id: ", simulationId)

	sId, oErr := primitive.ObjectIDFromHex(simulationId)
	if oErr != nil {
		log.Warn("invalid simulationId: ", oErr.Error())
		err = oErr
		return
	}
	if Gamblers, err = gambler.GetGamblers(&sId); err != nil {
		log.Error("fail to get gamblers: ", err.Error())
		return
	}

	log.Debug("gambler loaded, simulation id: ", simulationId)
	return
}

func Start(days int) {
	log.Debug("start casino, days: ", days)

	var wg sync.WaitGroup

	for _, gbl := range Gamblers {
		wg.Add(1)
		go gbl.PlaySince(&wg, days)
	}

	wg.Wait()
	log.Debug("completed")
	//for {
	//	time.Sleep(5 * time.Second)
	//}
}
