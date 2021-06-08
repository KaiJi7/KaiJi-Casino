package collection

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Decision struct {
	Id         *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	StrategyId *primitive.ObjectID `json:"strategy_id" bson:"strategy_id"`
	GambleId   *primitive.ObjectID `json:"gamble_id" bson:"gamble_id"`
	Bet        Bet                 `json:"bet" bson:"bet"`
	Put        float64             `json:"put" bson:"put"`
}

func (d Decision) String() string {
	return fmt.Sprintf("Bet: %s, Put: %d", d.Bet, &d.Put)
}
