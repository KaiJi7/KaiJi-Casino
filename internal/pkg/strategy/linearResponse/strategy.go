package linearResponse

import (
	"KaiJi-Casino/internal/pkg/strategy/common"
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
	lostCount int
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
	s.lostCount = 0
}

func (s *Strategy) OnLose(decision structs.Decision) {
	s.lostCount += 1
}

func (s *Strategy) OnTie(decision structs.Decision) {
	log.Warn("unhandled on tie")
}

func (s *Strategy) getPut(odds float64) float64 {
	if s.lostCount == 0 {
		return 1
	}

	return math.Ceil(float64(s.lastPut*(s.lastPut+4))/2*(odds-1) + s.lastPut*(s.slope-1)/(odds-1))
	//return math.Ceil(float64(s.lostCount*(s.lostCount+1)) * (s.slope - 1) / (2 * (odds - 1)))
}
