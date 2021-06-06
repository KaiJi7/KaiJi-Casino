package db

import (
	"KaiJi-Casino/internal/pkg/db/collection"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c client) CreateStrategy(gamblerId *primitive.ObjectID, name collection.StrategyName, description string) (strategy collection.StrategyData, err error) {
	strategy = collection.StrategyData{
		GamblerId:   gamblerId,
		Name:        name,
		Description: description,
	}

	res, dbErr := c.Strategy.InsertOne(nil, strategy)
	if dbErr != nil {
		log.Error("fail to insert strategy: ", dbErr.Error())
		err = dbErr
		return
	}
	id := res.InsertedID.(primitive.ObjectID)
	strategy.Id = &id
	return
}

func (c client) GetStrategy(gamblerId *primitive.ObjectID) (strategy collection.StrategyData, err error) {
	filter := bson.M{
		"gambler_id": gamblerId,
	}
	if err := c.Strategy.FindOne(nil, filter).Decode(&strategy); err != nil {
		log.Error("fail to get strategy: ", gamblerId.Hex(), ". ", err.Error())
	}
	return
}
