package bet

import (
	"KaiJi-Casino/internal/pkg/db/collection"
	log "github.com/sirupsen/logrus"
	"sync"
)

type lowestResponse struct {
	Name string
}

var (
	lowestResponseOnce     sync.Once
	lowestResponseInstance *lowestResponse
)

func NewLowestResponse() *lowestResponse {
	lowestResponseOnce.Do(func() {
		lowestResponseInstance = &lowestResponse{
			Name: "LowestResponse",
		}
		log.Debug("lowest response bet strategy initialized")
	})
	return lowestResponseInstance
}

func (l lowestResponse) GetDecisions(games []collection.SportsData, _ map[string]interface{}) []collection.BetInfo {
	log.Debug("get lowest response decision")

	lowerDecision := NewLowerResponse().GetDecisions(games, nil)
	if len(lowerDecision) == 0 {
		log.Info("no valid gamble")
		return nil
	}
	decision := lowerDecision[0]
	for _, d := range lowerDecision[1:] {
		if d.Response < decision.Response {
			decision = d
		}
	}

	return []collection.BetInfo{decision}
}
