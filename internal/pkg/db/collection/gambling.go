package collection

import (
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sort"
)

const (
	GamblingTypeUnknown     GamblingType = "unknown"
	GamblingTypeOriginal    GamblingType = "original"
	GamblingTypeSpreadPoint GamblingType = "spread_point"
	GamblingTypeTotalScore  GamblingType = "total_score"
)

type GamblingType string

type Gambling struct {
	Id         *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Type       GamblingType        `json:"type" bson:"type"`
	GameId     *primitive.ObjectID `json:"game_id" bson:"game_id"`
	Odds       []Odds              `json:"odds" bson:"odds"`
	Properties []struct {
		Name  string      `json:"name" bson:"name"`
		Value interface{} `json:"value" bson:"value"`
	} `json:"properties" bson:"properties"`
}

type Odds struct {
	Bet  Bet      `json:"bet" bson:"bet"`
	Odds *float64 `json:"odds" bson:"odds"`
}

func (g Gambling) GetOdds(bet Bet) float64 {
	for _, odd := range g.Odds {
		if odd.Bet == bet {
			return *odd.Odds
		}
	}
	log.Warn("no odds info for bet: ", bet)
	return 0
}

func (g Gambling) GetProperty(property string) interface{} {
	for _, p := range g.Properties {
		if p.Name == property {
			return p.Value
		}
	}

	log.Warn("no property: ", property)
	return nil
}

func (g Gambling) Betable() bool {
	for _, odd := range g.Odds {
		if odd.Odds == nil {
			return false
		}
	}
	return true
}

func (g Gambling) SortedOdds() []Odds {
	sort.SliceStable(g.Odds, func(i, j int) bool {
		return *g.Odds[i].Odds < *g.Odds[j].Odds
	})
	return g.Odds
}
