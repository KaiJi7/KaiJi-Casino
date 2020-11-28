package put

import (
	"KaiJi-Casino/internal/pkg/db/collection"
	log "github.com/sirupsen/logrus"
	"sync"
)

// Double bet if lose, otherwise 1.
type fooDouble struct {
	Name string
}

var (
	fooDoubleOnce     sync.Once
	fooDoubleInstance *fooDouble
)

func NewFooDouble() *fooDouble {
	fooDoubleOnce.Do(func() {
		fooDoubleInstance = &fooDouble{
			Name: "FooDouble",
		}
		log.Debug("foo double put strategy initialized")
	})
	return fooDoubleInstance
}

func (f fooDouble) GetUnit(history []collection.GambleHistory) (unit int) {
	log.Debug("get foo double put unit")
	unit = 1
	for i := len(history) - 1; i > 0; i-- {
		if history[i].Win {
			return
		} else {
			unit *= 2
		}
	}
	return
}
