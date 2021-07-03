package common

import (
	"github.com/KaiJi7/common/structs"
)

type Strategy interface {
	MakeDecision(gambles []structs.Gambling) []structs.Decision
	TargetGameType() []structs.GameType

	// for strategies to update their arguments based on gamble result
	OnWin(decision structs.Decision)
	OnLose(decision structs.Decision)
	OnTie(decision structs.Decision)
}
