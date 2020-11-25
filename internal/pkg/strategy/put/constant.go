package put

import (
	"KaiJi-Casino/internal/pkg/db/collection"
	log "github.com/sirupsen/logrus"
	"sync"
)

type constant struct {
	Name string
}

var (
	constantOnce     sync.Once
	constantInstance *constant
)

func NewConstant() *constant {
	constantOnce.Do(func() {
		constantInstance = &constant{
			Name: "Constant",
		}
		log.Debug("constant put strategy initialized")
	})
	return constantInstance
}

func (c constant) GetUnit(history []collection.GambleHistory) int {
	log.Debug("get constant put unit")
	return 1
}