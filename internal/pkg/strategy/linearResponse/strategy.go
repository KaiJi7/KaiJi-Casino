package linearResponse

import (
	"KaiJi-Casino/internal/pkg/strategy/common"
	"fmt"
	"github.com/KaiJi7/common/structs"
	log "github.com/sirupsen/logrus"
	"math"
	"math/rand"
	"time"
)

const (
	defaultSlope = 1.75
)

type Strategy struct {
	structs.StrategyData
	loseCount float64
	loseTimes int
	slope     float64
	lastPut   float64
}

func New(data structs.StrategyData) common.Strategy {
	slope := data.GetProperty("slope").(float64)
	if slope == 0 {
		slope = defaultSlope
	}
	return &Strategy{
		StrategyData: data,
		slope:        slope,
		lastPut:      1,
	}
}

func (s *Strategy) TargetGameType() []structs.GameType {
	return []structs.GameType{structs.GameTypeAll}
}

func (s *Strategy) MakeDecision(gambles []structs.Gambling) []structs.Decision {
	if len(gambles) == 0 {
		return nil
	}

	// filter out original and unknown
	gambles = func() (gs []structs.Gambling) {
		for _, g := range gambles {
			if g.Type == structs.GamblingTypeSpreadPoint || g.Type == structs.GamblingTypeTotalScore {
				gs = append(gs, g)
			}
		}
		return
	}()

	// random pick gamble and odds
	rand.Seed(time.Now().UnixNano())
	gamble := gambles[rand.Intn(len(gambles))]
	odds := gamble.Odds[rand.Intn(len(gamble.Odds))]
	s.lastPut = s.getPut(*odds.Odds)
	return []structs.Decision{
		{
			StrategyId: s.Id,
			GambleId:   gamble.Id,
			Bet:        odds.Bet,
			Put:        s.lastPut,
		},
	}
}

func (s *Strategy) OnWin(decision structs.Decision) {
	s.loseCount = 0
	s.loseTimes = 0
}

func (s *Strategy) OnLose(decision structs.Decision) {
	s.loseCount += s.lastPut
	s.loseTimes += 1
}

func (s *Strategy) OnTie(decision structs.Decision) {
	log.Warn("unhandled on tie")
}

func (s *Strategy) getPut(odds float64) float64 {
	log.Debug(fmt.Sprintf("last put: %f, odds: %f", s.lastPut, odds))

	if s.loseCount == 0 {
		return 1
	}
	
	return math.Ceil((s.loseCount + float64(s.loseTimes)*(s.slope-1)) / (odds - 1))
}
