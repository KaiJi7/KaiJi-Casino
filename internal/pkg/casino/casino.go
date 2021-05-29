package casino

import (
	"KaiJi-Casino/internal/pkg/db/collection"
	"KaiJi-Casino/internal/pkg/strategy"
	log "github.com/sirupsen/logrus"
)

var gamblers []collection.GamblerData
var strategies []collection.StrategyData

func InitGambler() {

}

func InitStrategy(distribution map[strategy.Name]int) {
	log.Debug("init strategy")

	for strategyName, count := range distribution {
		for i:= 0; i < count; i++ {
			strategies = append(strategies, )
		}
	}
}

func LoadGambler() {

}

func Open() {

}
