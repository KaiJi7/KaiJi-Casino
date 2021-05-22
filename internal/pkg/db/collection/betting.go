package collection

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	BET_GUEST Bet = "guest"
	BET_HOST  Bet = "host"
	BET_UNDER Bet = "under"
	BET_OVER  Bet = "over"

	BETTING_SOURCE_WILD_MEMBER = "wild_member"
)

type Bet string
type BettingSource string

type Betting struct {
	Id         *primitive.ObjectID `json:"id"`
	GamblingId *primitive.ObjectID `json:"gambling_id"`
	Source     BettingSource       `json:"source"`
	Bet        []struct {
		Side     Bet `json:"side"`
		Quantity int `json:"quantity"`
	} `json:"bet"`
}
