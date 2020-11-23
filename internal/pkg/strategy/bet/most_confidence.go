package bet

import (
	"KaiJi-Casino/internal/pkg/db/collection"
	log "github.com/sirupsen/logrus"
	"sync"
)

type mostConfidence struct {
	Name string
}

var (
	mostConfidenceOnce     sync.Once
	mostConfidenceInstance *mostConfidence
)

func NewMostConfidence() *mostConfidence {
	mostConfidenceOnce.Do(func() {
		mostConfidenceInstance = &mostConfidence{
			Name: "MostConfidence",
		}
		log.Debug("most confidence bet strategy initialized")
	})
	return mostConfidenceInstance
}

func (m mostConfidence) GetDecisions(games []collection.SportsData, parameters map[string]interface{}) []collection.BetInfo {
	log.Debug("get most confidence decision")

	confidentlyDecisions := NewConfidenceBase().GetDecisions(games, parameters)
	if len(confidentlyDecisions) == 0 {
		log.Info("no confidently decisions")
		return nil
	}

	decision := confidentlyDecisions[0]
	for _, d := range confidentlyDecisions[1:] {
		if decision.Confidence < d.Confidence {
			decision = d
		}
	}
	return []collection.BetInfo{decision}
}
