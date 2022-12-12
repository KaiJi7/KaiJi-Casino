package linearResponse

import (
	"KaiJi-Casino/internal/pkg/strategy/common"
	"github.com/KaiJi7/common/structs"
	log "github.com/sirupsen/logrus"
	"math"
)

const (
	defaultSlope = 1.75
)

type Strategy struct {
	structs.StrategyData
	lostCount int
	slope     float64
}

func New(data structs.StrategyData, slope float64) common.Strategy {
	if slope == 0 {
		slope = defaultSlope
	}
	return &Strategy{
		StrategyData: data,
		slope:        slope,
	}
}

func (s *Strategy) TargetGameType() []structs.GameType {
	return []structs.GameType{structs.GameTypeAll}
}

func (s *Strategy) MakeDecision(gambles []structs.Gambling) []structs.Decision {
	//decisions := make([]structs.Decision, 1)
	//for _, gamble := range gambles {
	//	decision := structs.Decision{
	//		StrategyId: s.Id,
	//		GambleId:   gamble.Id,
	//		Bet:        gamble.SortedOdds()[0].Bet,
	//		Put:        s.getPut(),
	//	}
	//	//decisions = append(decisions, decision)
	//	decisions[0] = decision
	//	break
	//}

	//return decisions

	if len(gambles) == 0 {
		return nil
	}

	return []structs.Decision{
		{
			StrategyId: s.Id,
			GambleId:   gambles[0].Id,
			Bet:        gambles[0].SortedOdds()[0].Bet,
			Put:        s.getPut(),
		},
	}
}

func (s *Strategy) OnWin(decision structs.Decision) {
	s.lostCount = 0
}

func (s *Strategy) OnLose(decision structs.Decision) {
	s.lostCount += 1
}

func (s *Strategy) OnTie(decision structs.Decision) {
	log.Warn("unhandled on tie")
}

func (s *Strategy) getPut() float64 {
	return math.Ceil(float64(s.lostCount*(s.lostCount+1)) / (2 * (s.slope - 1)))
}
