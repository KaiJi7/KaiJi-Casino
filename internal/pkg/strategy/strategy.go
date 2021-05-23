package strategy

import "KaiJi-Casino/internal/pkg/db/collection"

type Decision struct {
	Bet  collection.Bet
	Put float64
}

type Strategy interface {
	MakeDecision(gambles []collection.Gambling) []Decision
}
