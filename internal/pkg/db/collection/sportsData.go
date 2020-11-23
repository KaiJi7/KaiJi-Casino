package collection

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	GambleTypeTotalPoint  = "total_point"
	GambleTypeSpreadPoint = "spread_point"
	GambleTypeOriginal    = "original"
)

type SportsData struct {
	Id         primitive.ObjectID `json:"_id" bson:"_id"`
	GameTime   *time.Time         `json:"game_time,omitempty" bson:"game_time,omitempty"`
	GambleId   *string            `json:"gamble_id,omitempty" bson:"gamble_id,omitempty"`
	GameType   string             `json:"game_type" bson:"game_type"`
	Guest      TeamInfo           `json:"guest" bson:"guest"`
	Host       TeamInfo           `json:"host" bson:"host"`
	GambleInfo *GambleInfo        `json:"gamble_info,omitempty" bson:"gamble_info,omitempty"`
}

type TeamInfo struct {
	Name  string `json:"name" bson:"name"`
	Score int    `json:"score,omitempty" bson:"score,omitempty"`
}

type GambleInfo struct {
	TotalPoint  *PointBetInfo `json:"total_point,omitempty" bson:"total_point,omitempty"`
	SpreadPoint *SideBetInfo  `json:"spread_point,omitempty" bson:"spread_point,omitempty"`
	Original    *SideBetInfo  `json:"original,omitempty" bson:"original,omitempty"`
}

type PointBetInfo struct {
	Threshold *float64 `json:"threshold,omitempty" bson:"threshold,omitempty"`
	Response  *struct {
		Under float64 `json:"under,omitempty" bson:"under,omitempty"`
		Over  float64 `json:"over,omitempty" bson:"over,omitempty"`
	} `json:"response,omitempty" bson:"response,omitempty"`
	Judgement  *string `json:"judgement,omitempty" bson:"judgement,omitempty"`
	Prediction *struct {
		Under *Vote `json:"under,omitempty" bson:"under,omitempty"`
		Over  *Vote `json:"over,omitempty" bson:"over,omitempty"`
		Major *bool `json:"major,omitempty" bson:"major,omitempty"`
	} `json:"prediction,omitempty" bson:"prediction,omitempty"`
}

type SideBetInfo struct {
	Guest    float64 `json:"guest,omitempty" bson:"guest,omitempty"`
	Host     float64 `json:"host,omitempty" bson:"host,omitempty"`
	Response *struct {
		Guest float64 `json:"guest,omitempty" bson:"guest,omitempty"`
		Host  float64 `json:"host,omitempty" bson:"host,omitempty"`
	} `json:"response,omitempty" bson:"response,omitempty"`
	Judgement  string `json:"judgement,omitempty" bson:"judgement,omitempty"`
	Prediction struct {
		Guest *Vote `json:"guest,omitempty" bson:"guest,omitempty"`
		Host  *Vote `json:"host,omitempty" bson:"host,omitempty"`
		Major bool  `json:"major,omitempty" bson:"major,omitempty"`
	} `json:"prediction,omitempty" bson:"prediction,omitempty"`
}

type Vote struct {
	Percentage float64 `json:"percentage,omitempty" bson:"percentage,omitempty"`
	Population int     `json:"population,omitempty" bson:"population,omitempty"`
}
