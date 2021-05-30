package casino

import (
	"KaiJi-Casino/internal/pkg/db"
	"KaiJi-Casino/internal/pkg/db/collection"
	"KaiJi-Casino/internal/pkg/gambler"
	"KaiJi-Casino/internal/pkg/strategy"
	log "github.com/sirupsen/logrus"
)

var Gamblers []gambler.Gambler
//var strategies []collection.StrategyData

func InitGamblers(simulation collection.Simulation) (err error) {
	log.Debug("init gamblers: ", simulation.String())

	for strategyName, count := range simulation.StrategySchema {
		content, exist := strategy.NameMap[strategyName]
		if !exist {
			log.Warn("invalid strategy: ", strategyName)
			return
		}

		for i := 0; i < count; i++ {
			gbl, err := db.New().CreateGambler(simulation.Id, simulation.GamblerInitialMoney)
			if err != nil {
				log.Error("fail to create gambler: ", err.Error())
				return
			}

			if _, err := db.New().CreateStrategy(gbl.Id, strategyName, content.Description); err != nil {
				log.Error("fail to create strategy: ", err.Error())
				return
			}
		}
	}

	Gamblers, err = gambler.GetGamblers(simulation.Id)
	return
}
