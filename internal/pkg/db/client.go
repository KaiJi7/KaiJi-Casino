package db

import (
	"KaiJi-Casino/internal/pkg/configs"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

type client struct {
	*mongo.Client
	SportsData *mongo.Collection
	Gambler    *mongo.Collection
}

var (
	once     sync.Once
	instance *client
)

func New() *client {
	once.Do(func() {
		c, err := mongo.Connect(nil, options.Client().ApplyURI(configs.New().Mongo.ConnectionString))
		if err != nil {
			panic(err)
		}
		db := c.Database(configs.New().Mongo.Db)
		instance = &client{
			Client:     c,
			SportsData: db.Collection("sports_data"),
			Gambler:    db.Collection("gambler"),
		}
		if err := instance.Ping(nil, nil); err != nil {
			panic(err)
		}
		log.Debug("mongo client initialized")
	})
	return instance
}
