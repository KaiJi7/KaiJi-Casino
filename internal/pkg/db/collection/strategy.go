package collection

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	StrategyNameLowerResponse  StrategyName = "lowerResponse"
	StrategyNameLowestResponse StrategyName = "lowestResponse"
	StrategyNameConfidenceBase StrategyName = "confidenceBase"
	StrategyNameMostConfidence StrategyName = "mostConfidence"
)

type StrategyName string

type StrategyData struct {
	Id         *primitive.ObjectID    `json:"id,omitempty" bson:"_id,omitempty"`
	Meta       *primitive.ObjectID    `json:"meta" bson:"meta"`
	GamblerId  *primitive.ObjectID    `json:"gambler_id" bson:"gambler_id"`
	Name       StrategyName           `json:"name" bson:"name"`
	Properties map[string]interface{} `json:"properties" bson:"properties"`
	//Description string                 `json:"description" bson:"description"`
}

type StrategyMeta struct {
	Id          *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        StrategyName        `json:"name" bson:"name"`
	Description string              `json:"description" bson:"description"`
	Properties  []struct {
		Name string `json:"name" bson:"name"`
		Type string `json:"type" bson:"type"` // int, float, string
	} `json:"properties" bson:"properties"`
}
