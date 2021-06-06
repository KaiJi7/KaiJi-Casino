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

type GambleHistory struct {
	Id          *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	DecisionId  *primitive.ObjectID `json:"decision_id" bson:"decision_id"`
	Winner      GambleWinner        `json:"winner"`
	MoneyBefore float64             `json:"money_before" bson:"money_before"`
	MoneyAfter  float64             `json:"money_after" bson:"money_after"`
}
