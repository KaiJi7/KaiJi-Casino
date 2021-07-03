package lowestResponse

import (
	"KaiJi-Casino/internal/pkg/strategy/common"
	"github.com/KaiJi7/common/structs"
	log "github.com/sirupsen/logrus"
)

const defaultLowestOdds = 10.0

type Strategy struct {
	structs.StrategyData
}

func New(data structs.StrategyData) common.Strategy {
	return Strategy{
		data,
	}
}

func (s Strategy) TargetGameType() []structs.GameType {
	return []structs.GameType{structs.GameTypeAll}
}

func (s Strategy) MakeDecision(gambles []structs.Gambling) (decision []structs.Decision) {
	var lowestOdds = defaultLowestOdds
	for _, gamble := range gambles {
		if !gamble.Betable() {
			log.Info("unbetable gamble: ", gamble.Id.Hex())
			continue
		}

		if *gamble.SortedOdds()[0].Odds < lowestOdds {
			lowestOdds = *gamble.SortedOdds()[0].Odds
			decision = []structs.Decision{
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

func (s Strategy) OnWin(decision structs.Decision) {

}

func (s Strategy) OnLose(decision structs.Decision) {

}

func (s Strategy) OnTie(decision structs.Decision) {

}
