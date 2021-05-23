package gambler

import "KaiJi-Casino/internal/pkg/db/collection"

type Decision struct {
	Bet  collection.Bet
	Odds float64
}

type Strategy interface {
	MakeDecision(gameData []collection.SportsGameResult) []Decision
	//OnLose()
	//OnWin()
	//OnTie()
}
