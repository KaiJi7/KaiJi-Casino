package sportsGame

import (
	"KaiJi-Casino/internal/pkg/db/collection"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LowerResponse struct {
	Id          *primitive.ObjectID `json:"id" bson:"_id"`
	GamblerId   *primitive.ObjectID `json:"gambler_id" bson:"gambler_id"`
	Name        string              `json:"name" bson:"name"`
	Description string              `json:"description" bson:"description"`
}

func (l LowerResponse) MakeDecision(gambles []collection.Gambling) []collection.Decision {
	decisions := make([]collection.Decision, 0)
	for _, gamble := range gambles {
		if !gamble.Betable() {
			log.Info("unbetable gamble: ", gamble.Id.Hex())
			continue
		}

		decision := collection.Decision{
			StrategyId: l.Id,
			GambleId:   gamble.Id,
			Bet:        gamble.SortedOdds()[0].Bet,
			Put:        1,
		}
		decisions = append(decisions, decision)
	}
	return decisions
}
