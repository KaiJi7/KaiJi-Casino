package collection

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	BetGuest Bet = "guest"
	BetHost  Bet = "host"
	BetTie   Bet = "tie"
	BetUnder Bet = "under"
	BetOver  Bet = "over"
	BetEqual Bet = "equal"

	BettingSourceWildMember = "wild_member"
)

type Bet string
type BettingSource string

type Betting struct {
	Id         *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	GamblingId *primitive.ObjectID `json:"gambling_id" bson:"gambling_id"`
	Source     BettingSource       `json:"source" bson:"source"`
	Bet        []struct {
		Side     Bet `json:"side" bson:"side"`
		Quantity int `json:"quantity" bson:"quantity"`
	} `json:"bet" bson:"bet"`
}
