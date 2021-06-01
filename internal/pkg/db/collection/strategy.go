package collection

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	StrategyNameLowerResponse  StrategyName = "LowerResponse"
	StrategyNameLowestResponse StrategyName = "LowestResponse"
)

type StrategyName string

type StrategyData struct {
	Id          *primitive.ObjectID    `json:"id,omitempty" bson:"_id,omitempty"`
	GamblerId   *primitive.ObjectID    `json:"gambler_id" bson:"gambler_id"`
	Name        StrategyName           `json:"name" bson:"name"`
	Properties  map[string]interface{} `json:"properties" bson:"properties"`
	Description string                 `json:"description" bson:"description"`
}
