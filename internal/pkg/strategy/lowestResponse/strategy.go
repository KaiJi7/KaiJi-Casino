package lowestResponse

import (
	"KaiJi-Casino/internal/pkg/db/collection"
	"KaiJi-Casino/internal/pkg/strategy/common"
	log "github.com/sirupsen/logrus"
)

const defaultLowestOdds = 10.0

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

func (s Strategy) MakeDecision(gambles []collection.Gambling) (decision []collection.Decision) {
	var lowestOdds = defaultLowestOdds
	for _, gamble := range gambles {
		if !gamble.Betable() {
			log.Info("unbetable gamble: ", gamble.Id.Hex())
			continue
		}

		if *gamble.SortedOdds()[0].Odds < lowestOdds {
			lowestOdds = *gamble.SortedOdds()[0].Odds
			decision = []collection.Decision{
				{
					StrategyId: s.Id,
					GambleId:   gamble.Id,
					Bet:        gamble.SortedOdds()[0].Bet,
					Put:        1,
				},
			}
		}
	}
	return
}

func (s Strategy) OnWin(decision collection.Decision) {

}

func (s Strategy) OnLose(decision collection.Decision) {

}

func (s Strategy) OnTie(decision collection.Decision) {

}
