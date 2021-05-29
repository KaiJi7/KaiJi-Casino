package lowerResponse

import (
	"KaiJi-Casino/internal/pkg/db/collection"
	"KaiJi-Casino/internal/pkg/strategy"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Strategy struct {
	Id        *primitive.ObjectID `json:"id" bson:"_id"`
	GamblerId *primitive.ObjectID `json:"gambler_id" bson:"gambler_id"`
	Name      strategy.Name       `json:"name" bson:"name"`
	//Description string              `json:"description" bson:"description"`
}

func New(data collection.StrategyData) strategy.Strategy {
	return Strategy{
		Id:        data.Id,
		GamblerId: data.GamblerId,
		Name:      data.Name,
	}
}

func (s Strategy) MakeDecision(gambles []collection.Gambling) []collection.Decision {
	decisions := make([]collection.Decision, 0)
	for _, gamble := range gambles {
		if !gamble.Betable() {
			log.Info("unbetable gamble: ", gamble.Id.Hex())
			continue
		}

		decision := collection.Decision{
			StrategyId: s.Id,
			GambleId:   gamble.Id,
			Bet:        gamble.SortedOdds()[0].Bet,
			Put:        1,
		}
		decisions = append(decisions, decision)
	}
	return decisions
}

func (s Strategy) OnWin(decision collection.Decision) {

}

func (s Strategy) OnLose(decision collection.Decision) {

}

func (s Strategy) OnTie(decision collection.Decision) {

}
