package db

import (
	"KaiJi-Casino/internal/pkg/db/collection"
	"KaiJi-Casino/internal/pkg/strategy"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c client) CreateStrategy(gamblerId *primitive.ObjectID, name strategy.Name, description string) (strategy collection.StrategyData, err error) {
	strategy = collection.StrategyData{
		GamblerId:   gamblerId,
		Name:        name,
		Description: description,
	}

	res, err := c.Strategy.InsertOne(nil, strategy)
	if err != nil {
		log.Error("fail to insert strategy: ", err.Error())
		return
	}
	strategy.Id = res.InsertedID.(*primitive.ObjectID)
	return
}

func (c client) GetStrategy(strategyId *primitive.ObjectID) (strategy collection.StrategyData, err error) {
	filter := bson.M{
		"_id": strategyId,
	}
	if err := c.Strategy.FindOne(nil, filter).Decode(&strategy); err != nil {
		log.Error("fail to get strategy: ", strategyId.Hex(), ". ", err.Error())
	}
	return
}
