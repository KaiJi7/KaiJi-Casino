package collection

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Gambler struct {
	Name           string   `json:"name"`
	InitCapital    int      `json:"init_capital"`
	CurrentCapital int      `json:"current_capital"`
	Strategy       Strategy `json:"strategy"`
	GambleHistory  []GambleHistory
}

func (g *Gambler) String() string {
	return fmt.Sprintf("gambler name: %s, bet strategy: %s, put strategy: %s", g.Name, g.Strategy.Bet.Name, g.Strategy.Put.Name)
}

type Strategy struct {
	Bet struct {
		Name       string                 `json:"name"`
		Parameters map[string]interface{} `json:"parameters,omitempty"`
	} `json:"bet"`
	Put struct {
		Name       string                 `json:"name"`
		Parameters map[string]interface{} `json:"parameters,omitempty"`
	} `json:"put"`
}

type GambleHistory struct {
	BetInfo       BetInfo `json:"bet_info"`
	Win           bool    `json:"win"`
	CapitalBefore int     `json:"capital_before"`
	CapitalAfter  int     `json:"capital_after"`
}

const (
	LocalBanker    = "local"
	NationalBanker = "national"

	BetSideUnder = "under"
	BetSideOver  = "over"
	BetSideGuest = "guest"
	BetSideHost  = "host"
)

type BetInfo struct {
	GameId     primitive.ObjectID `json:"game_id"`
	Banker     string             `json:"banker"`
	GambleType string             `json:"gamble_type"`
	BetSide    string             `json:"bet_side"`
	Unit       int                `json:"unit"`
	Confidence float64            `json:"confidence,omitempty"`
	Response   float64            `json:"response,omitempty"`
}

func (b *BetInfo) GetResponse() float64 {
	// TODO: get from db, make cache

	return 0
}
