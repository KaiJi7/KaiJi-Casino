package casino

import (
	"KaiJi-Casino/internal/pkg/db"
	"KaiJi-Casino/internal/pkg/db/collection"
	"KaiJi-Casino/internal/pkg/gambler"
	"KaiJi-Casino/internal/pkg/strategy"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sync"
)

var Gamblers []gambler.Gambler
//var strategies []collection.StrategyData

func CreateGamblers(simulation collection.Simulation) (err error) {
	log.Debug("create gamblers: ", simulation.String())

	for strategyName, count := range simulation.StrategySchema {
		content, exist := strategy.NameMap[strategyName]
		if !exist {
			log.Warn("invalid strategy: ", strategyName)
			return
		}

		for i := 0; i < count; i++ {
			gbl, dbErr := db.New().CreateGambler(simulation.Id, simulation.GamblerInitialMoney)
			if dbErr != nil {
				log.Error("fail to create gambler: ", dbErr.Error())
				err = dbErr
				return
			}

			if _, err = db.New().CreateStrategy(gbl.Id, strategyName, content.Description); err != nil {
				log.Error("fail to create strategy: ", err.Error())
				return
			}
		}
	}

	Gamblers, err = gambler.GetGamblers(simulation.Id)
	return
}

func LoadGamblers(simulationId string) (err error){
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
