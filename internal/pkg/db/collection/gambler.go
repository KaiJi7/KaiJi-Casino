package collection

import "go.mongodb.org/mongo-driver/bson/primitive"

type GamblerData struct {
	Id           *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	SimulationId *primitive.ObjectID `json:"simulation_id" bson:"simulation_id"`
	//StrategyId   *primitive.ObjectID `json:"strategy_id" bson:"strategy_id"`
	MoneyBegin   float64             `json:"money_begin" bson:"money_begin"`
	MoneyCurrent float64             `json:"money_current" bson:"money_current"`
}
