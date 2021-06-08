package db

import (
	"KaiJi-Casino/internal/pkg/db/collection"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c client) CreateStrategy(gamblerId *primitive.ObjectID, name collection.StrategyName, meta *primitive.ObjectID, properties map[string]interface{}) (strategy collection.StrategyData, err error) {
	strategy = collection.StrategyData{
		GamblerId:  gamblerId,
		Name:       name,
		Meta:       meta,
		Properties: properties,
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

func (c client) CreateMetaStrategy(metasStrategy collection.StrategyMeta) (err error) {
	_, err = c.StrategyMeta.InsertOne(nil, metasStrategy)
	return
}

func (c client) GetStrategyData(gamblerId *primitive.ObjectID) (strategy collection.StrategyData, err error) {
	filter := bson.M{
		"gambler_id": gamblerId,
	}
	if err := c.Strategy.FindOne(nil, filter).Decode(&strategy); err != nil {
		log.Error("fail to get strategy: ", gamblerId.Hex(), ". ", err.Error())
	}
	return
}

func (c client) GetStrategyMetaData(name collection.StrategyName) (strategyMeta collection.StrategyMeta, err error) {
	filter := bson.M{
		"name": name,
	}
	if err := c.StrategyMeta.FindOne(nil, filter).Decode(&strategyMeta); err != nil {
		log.Error("fail to get strategy meta: ", name, ". ", err.Error())
		panic(err.Error())
	}
	return
}
