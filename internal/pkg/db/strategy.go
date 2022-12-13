package db

import (
	"github.com/KaiJi7/common/structs"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c *client) CreateStrategy(gamblerId *primitive.ObjectID, name structs.StrategyName, meta *primitive.ObjectID, properties map[string]interface{}) (strategy structs.StrategyData, err error) {
	strategy = structs.StrategyData{
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

func (c *client) CreateMetaStrategy(metasStrategy structs.StrategyMeta) (err error) {
	filter := bson.M{"name": metasStrategy.Name}
	update := bson.M{"$set": metasStrategy}
	opts := options.Update().SetUpsert(true)
	_, err = c.StrategyMeta.UpdateOne(nil, filter, update, opts)
	return
}

func (c *client) GetStrategyData(gamblerId *primitive.ObjectID) (strategy structs.StrategyData, err error) {
	filter := bson.M{
		"gambler_id": gamblerId,
	}
	if err := c.Strategy.FindOne(nil, filter).Decode(&strategy); err != nil {
		log.Error("fail to get strategy: ", gamblerId.Hex(), ". ", err.Error())
	}
	return
}

func (c *client) GetStrategyMetaData(name structs.StrategyName) (strategyMeta structs.StrategyMeta, err error) {
	filter := bson.M{
		"name": name,
	}
	if err := c.StrategyMeta.FindOne(nil, filter).Decode(&strategyMeta); err != nil {
		log.Error("fail to get strategy meta: ", name, ". ", err.Error())
		panic(err.Error())
	}
	return
}
