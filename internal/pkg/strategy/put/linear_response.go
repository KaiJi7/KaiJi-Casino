package put

import (
	"KaiJi-Casino/internal/pkg/db/collection"
	log "github.com/sirupsen/logrus"
	"math"
	"sync"
)

const (
	defaultAverageNetProfit = 0.5
)

type linearResponse struct {
	Name string
}

var (
	linearResponseOnce     sync.Once
	linearResponseInstance *linearResponse
)

func NewLinearResponse() *linearResponse {
	linearResponseOnce.Do(func() {
		linearResponseInstance = &linearResponse{
			Name: "LinearResponse",
		}
		log.Debug("linear response put strategy initialized")
	})
	return linearResponseInstance
}

func (l linearResponse) GetUnit(history []collection.GambleHistory) (unit int) {
	log.Debug("get linear response put unit")

	var averageDailyNetProfit float64
	loseCount := 0
	putCount := 0
	for i := len(history) - 1; i > 0; i-- {
		if history[i].Win {
			averageDailyNetProfit = getAverageNetProfit(history[:i])
		} else {
			loseCount++
			putCount += history[i].BetInfo.Unit
		}
	}

	totalExpectProfit := averageDailyNetProfit * float64(loseCount)
	unit = int(math.Ceil((totalExpectProfit + float64(putCount)) / averageDailyNetProfit))

	log.Debug("linear response unit: ", unit)

	log.Debug("average daily net profit: ", averageDailyNetProfit)
	log.Debug("lose count: ", loseCount, ". put count: ", putCount)
	log.Debug("total expect net profit: ", totalExpectProfit)
	return
}

func getAverageNetProfit(history []collection.GambleHistory) float64 {
	if len(history) == 0 {
		return defaultAverageNetProfit
	}
	sum := 0.0
	for _, h := range history {
		sum += h.BetInfo.GetResponse()
	}
	return sum / float64(len(history)) - 1
}
