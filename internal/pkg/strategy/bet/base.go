package bet

import "KaiJi-Casino/internal/pkg/db/collection"

type BaseStrategy interface {
	GetDecisions([]collection.SportsData, map[string]interface{}) []collection.BetInfo
}