package collection

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	GambleWinnerBanker  GambleWinner = "banker"  // banker win, gambler lose
	GambleWinnerGambler GambleWinner = "gambler" // gambler win, banker lose
	GambleWinnerTie     GambleWinner = "tie"
)

type GambleWinner string

type N_GambleHistory struct {
	Id          *primitive.ObjectID `json:"id"`
	GamblerId   *primitive.ObjectID `json:"gambler_id"`
	GamblingId  *primitive.ObjectID `json:"gambling_id"`
	Bet         Bet                 `json:"bet"`
	Odds        float64             `json:"odds"`
	Winner      GambleWinner        `json:"winner"`
	MoneyBefore float64             `json:"money_before"`
	MoneyAfter  float64             `json:"money_after"`
}
