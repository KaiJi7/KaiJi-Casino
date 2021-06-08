package lowerResponse

import (
	"KaiJi-Casino/internal/pkg/db/collection"
	"KaiJi-Casino/internal/pkg/strategy/common"
	log "github.com/sirupsen/logrus"
)

type Strategy struct {
	collection.StrategyData
}

func New(data collection.StrategyData) common.Strategy {
	return Strategy{
		data,
	}
}

func (s Strategy) TargetGameType() []collection.GameType {
	return []collection.GameType{collection.GameTypeAll}
}

func (s Strategy) MakeDecision(gambles []collection.Gambling) []collection.Decision {
	decisions := make([]collection.Decision, 0)
	for _, gamble := range gambles {
		if !gamble.Betable() {
			log.Info("unbetable gamble: ", gamble.Id.Hex())
			continue
		}

		decision := collection.Decision{
			StrategyId: s.Id,
			GambleId:   gamble.Id,
			Bet:        gamble.SortedOdds()[0].Bet,
			Put:        1,
		}
		decisions = append(decisions, decision)
	}
	return decisions
}

func (s Strategy) OnWin(decision collection.Decision) {

}

func (s Strategy) OnLose(decision collection.Decision) {

}

func (s Strategy) OnTie(decision collection.Decision) {

}
