package cache

import (
	"github.com/allegro/bigcache"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

var (
	once     sync.Once
	instance *bigcache.BigCache
)

func New() *bigcache.BigCache{
	once.Do(func() {
		cache, err := bigcache.NewBigCache(bigcache.DefaultConfig(7 * time.Minute))
		if err != nil {
			log.Panic("fail to initial cache: ", err.Error())
			panic(err)
		}
		instance = cache
		log.Debug("cache initialized")
	})
	return instance
}
