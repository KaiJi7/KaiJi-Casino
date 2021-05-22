package collection

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	GAMBLING_TYPE_UNKNOWN      GamblingType = "unknown"
	GAMBLING_TYPE_ORIGINAL     GamblingType = "original"
	GAMBLING_TYPE_SPREAD_POINT GamblingType = "spread_point"
	GAMBLING_TYPE_TOTAL_SCORE  GamblingType = "total_score"
)

type GamblingType string

type Gambling struct {
	Id     *primitive.ObjectID `json:"id" bson:"_id"`
	Type   GamblingType        `json:"type" bson:"type"`
	GameId *primitive.ObjectID `json:"game_id" bson:"game_id"`
	Odds   []struct {
		Bet  Bet     `json:"bet" bson:"bet"`
		Odds float64 `json:"odds" bson:"odds"`
	} `json:"odds" bson:"odds"`
	Properties []struct {
		Name  string      `json:"name" bson:"name"`
		Value interface{} `json:"value" bson:"value"`
	} `json:"properties" bson:"properties"`
}
