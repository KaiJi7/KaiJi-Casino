package strategy

import (
	"KaiJi-Casino/internal/pkg/db/collection"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LowerResponse struct {
	GamblerId *primitive.ObjectID
}

func(l LowerResponse) MakeDecision(gambles []collection.Gambling) []Decision {
	decisions := make([]Decision, 0)
	for _, gamble := range gambles {
		if !gamble.Betable() {
			log.Info("unbetable gamble: ", gamble.Id.Hex())
			continue
		}

		decision := Decision{
			Bet: gamble.SortedOdds()[0].Bet,
			Put: 1,
		}
		decisions = append(decisions, decision)
	}
	return decisions
}
