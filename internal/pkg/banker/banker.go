package banker

import (
	log "github.com/sirupsen/logrus"
	"sync"
)

const (
	WinnerBanker  Winner = "banker"  // banker win, gambler lose
	WinnerGambler Winner = "gambler" // gambler win, banker lose
	WinnerTie     Winner = "tie"
)

type Winner string

type banker struct{}

var (
	once     sync.Once
	instance *banker
)

func New() *banker {
	once.Do(func() {
		instance = &banker{}
		log.Debug("banker initialized")
	})
	return instance
}
