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
	Decision      *mongo.Collection
	Game          *mongo.Collection
	Gambling      *mongo.Collection
	Betting       *mongo.Collection
	Gambler       *mongo.Collection
	GambleHistory *mongo.Collection
	Strategy      *mongo.Collection
	Simulation    *mongo.Collection
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
			Client:        c,
			Game:          db.Collection("game"),
			Decision:      db.Collection("decision"),
			Gambling:      db.Collection("gambling"),
			Betting:       db.Collection("betting"),
			Gambler:       db.Collection("gambler"),
			GambleHistory: db.Collection("gamble_history"),
			Strategy:      db.Collection("strategy"),
			Simulation:    db.Collection("simulation"),
		}
		if err := instance.Ping(nil, nil); err != nil {
			panic(err)
		}
		log.Debug("mongo client initialized")
	})
	return instance
}
