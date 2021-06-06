package common

import (
	"KaiJi-Casino/internal/pkg/db/collection"
)

type Strategy interface {
	MakeDecision(gambles []collection.Gambling) []collection.Decision
	TargetGameType() []collection.GameType

	// for strategies to update their arguments based on gamble result
	OnWin(decision collection.Decision)
	OnLose(decision collection.Decision)
	OnTie(decision collection.Decision)
}
