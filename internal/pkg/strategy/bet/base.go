package bet

import "KaiJi-Casino/internal/pkg/db/collection"

const (

)

type BaseStrategy interface {
	GetDecisions([]collection.SportsData, map[string]interface{}) []collection.BetInfo
}