package collection

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Simulation struct {
	Id                  *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	GamblerInitialMoney float64             `json:"gambler_initial_money" bson:"gambler_initial_money"`
	StrategySchema map[StrategyName][]struct {
		Quantity   int                    `json:"quantity" bson:"quantity" yaml:"quantity"`
		Properties map[string]interface{} `json:"properties" bson:"properties" yaml:"properties"`
	} `json:"strategy_info" bson:"strategy_schema"`
}

func (s Simulation) String() string {
	return fmt.Sprintf("Id: %s, Initial money: %f", s.Id.Hex(), s.GamblerInitialMoney)
}
