package strategy

import (
	"KaiJi-Casino/internal/pkg/db/collection"
)

const (
	LowerResponse  Name = "LowerResponse"
	LowestResponse Name = "LowestResponse"
)

type Name string

type Strategy interface {
	MakeDecision(gambles []collection.Gambling) []collection.Decision

	// for strategies to update their arguments based on gamble result
	OnWin(decision collection.Decision)
	OnLose(decision collection.Decision)
	OnTie(decision collection.Decision)
}
