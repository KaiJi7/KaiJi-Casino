package strategy

import (
	"KaiJi-Casino/internal/pkg/db/collection"
)

type Strategy interface {
	MakeDecision(gambles []collection.Gambling) []collection.Decision
}
