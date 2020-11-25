package put

import "KaiJi-Casino/internal/pkg/db/collection"

type BaseStrategy interface {
	GetUnit([]collection.GambleHistory) int
}
