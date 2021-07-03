package lowerResponse

import (
	"KaiJi-Casino/internal/pkg/strategy/common"
	"github.com/KaiJi7/common/structs"
	log "github.com/sirupsen/logrus"
)

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

func (s Strategy) MakeDecision(gambles []structs.Gambling) []structs.Decision {
	decisions := make([]structs.Decision, 0)
	for _, gamble := range gambles {
		if !gamble.Betable() {
			log.Info("unbetable gamble: ", gamble.Id.Hex())
			continue
		}

		decision := structs.Decision{
			StrategyId: s.Id,
			GambleId:   gamble.Id,
			Bet:        gamble.SortedOdds()[0].Bet,
			Put:        1,
		}
		decisions = append(decisions, decision)
	}
	return decisions
}

func (s Strategy) OnWin(decision structs.Decision) {

}

func (s Strategy) OnLose(decision structs.Decision) {

}

func (s Strategy) OnTie(decision structs.Decision) {

}
