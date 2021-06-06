package banker

import (
	log "github.com/sirupsen/logrus"
	"sync"
)

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
