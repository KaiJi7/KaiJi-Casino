package collection

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	GameTypeNba GameType = "NBA"
	GameTypeMlb GameType = "MLB"
	GameTypeNpb GameType = "NPB"

	GameTypeAll = "all" // for banker query only
)

type GameType string

type SportsGameResult struct {
	Id    *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
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

type SportsGameInfo struct {
	Id    *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Type  GameType            `json:"type" bson:"type"`
	Guest struct {
		Name  string `json:"name" bson:"name"`
		//Score int    `json:"score" bson:"score"`
	} `json:"guest" bson:"guest"`
	Host struct {
		Name  string `json:"name" bson:"name"`
		//Score int    `json:"score" bson:"score"`
	} `json:"host" bson:"host"`
	StartTime      *time.Time `json:"start_time,omitempty" bson:"start_time,omitempty"`
	StartTimeLocal *time.Time `json:"start_time_local,omitempty" bson:"start_time_local,omitempty"`
	Location       string     `json:"location,omitempty" bson:"location,omitempty"`
}
