package collection

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Simulation struct {
	Id                  *primitive.ObjectID  `json:"id,omitempty" bson:"_id,omitempty"`
	GamblerInitialMoney float64              `json:"gambler_initial_money"`
	StrategySchema      map[StrategyName]int `json:"strategy_info"`
}

func (s Simulation) String() string {
	return fmt.Sprintf("Id: %s, Initial money: %f, Schema: %v", s.Id.Hex(), s.GamblerInitialMoney, s.StrategySchema)
}
