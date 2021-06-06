package lowestResponse

import (
	"KaiJi-Casino/internal/pkg/db/collection"
	"KaiJi-Casino/internal/pkg/strategy/common"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const defaultLowestOdds = 10.0

type Strategy struct {
	Id        *primitive.ObjectID     `json:"id" bson:"_id"`
	GamblerId *primitive.ObjectID     `json:"gambler_id" bson:"gambler_id"`
	Name      collection.StrategyName `json:"name" bson:"name"`
	//Description string              `json:"description" bson:"description"`
}

func New(data collection.StrategyData) common.Strategy {
	return Strategy{
		Id:        data.Id,
		GamblerId: data.GamblerId,
		Name:      data.Name,
	}
}

func (s Strategy) TargetGameType() []collection.GameType {
	return []collection.GameType{collection.GameTypeAll}
}

func (s Strategy) MakeDecision(gambles []collection.Gambling) (decision []collection.Decision) {
	var lowestOdds = defaultLowestOdds
	for _, gamble := range gambles {
		if !gamble.Betable() {
			log.Info("unbetable gamble: ", gamble.Id.Hex())
			continue
		}

		if *gamble.SortedOdds()[0].Odds < lowestOdds {
			lowestOdds = *gamble.SortedOdds()[0].Odds
			decision = []collection.Decision{
				{
					StrategyId: s.Id,
					GambleId:   gamble.Id,
					Bet:        gamble.SortedOdds()[0].Bet,
					Put:        1,
				},
			}
		}
	}
	return
}

func (s Strategy) OnWin(decision collection.Decision) {

}

func (s Strategy) OnLose(decision collection.Decision) {

}

func (s Strategy) OnTie(decision collection.Decision) {

}
