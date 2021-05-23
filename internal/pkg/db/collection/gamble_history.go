package collection

import (
	"KaiJi-Casino/internal/pkg/banker"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type N_GambleHistory struct {
	Id          *primitive.ObjectID `json:"id"`
	GamblerId   *primitive.ObjectID `json:"gambler_id"`
	GamblingId  *primitive.ObjectID `json:"gambling_id"`
	Bet         Bet                 `json:"bet"`
	Odds        float64             `json:"odds"`
	Winner      banker.Winner       `json:"winner"`
	MoneyBefore float64             `json:"money_before"`
	MoneyAfter  float64             `json:"money_after"`
}
