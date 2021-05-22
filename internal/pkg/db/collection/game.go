package collection

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	GAME_TYPE_NBA = "NBA"
	GAME_TYPE_MLB = "MLB"
	GAME_TYPE_NPB = "NPB"
)

type GameType string

type Game struct {
	Id    *primitive.ObjectID `json:"id" bson:"_id"`
	Type  GameType            `json:"type" bson:"type"`
	Guest struct {
		Name  string `json:"name" bson:"name"`
		Score int    `json:"score" bson:"score"`
	} `json:"guest" bson:"guest"`
	Host struct {
		Name  string `json:"name" bson:"name"`
		Score int    `json:"score" bson:"score"`
	} `json:"host" bson:"host"`
	StartTime      *time.Time `json:"start_time,omitempty" bson:"start_time,omitempty"`
	StartTimeLocal *time.Time `json:"start_time_local,omitempty" bson:"start_time_local,omitempty"`
	Location       string     `json:"location,omitempty" bson:"location,omitempty"`
}
