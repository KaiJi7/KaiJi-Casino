package banker

import (
	log "github.com/sirupsen/logrus"
	"sync"
)

type Banker struct{}

var (
	once     sync.Once
	instance *Banker
)

func New() *Banker {
	once.Do(func() {
		instance = &Banker{}
		log.Debug("banker initialized")
	})
	return instance
}
