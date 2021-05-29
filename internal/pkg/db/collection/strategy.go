package collection

import (
	"KaiJi-Casino/internal/pkg/strategy"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StrategyData struct {
	Id          *primitive.ObjectID    `json:"id" bson:"_id"`
	GamblerId   *primitive.ObjectID    `json:"gambler_id" bson:"gambler_id"`
	Name        strategy.Name          `json:"name" bson:"name"`
	Properties  map[string]interface{} `json:"properties" bson:"properties"`
	Description string                 `json:"description" bson:"description"`
}
