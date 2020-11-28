package db

import (
	"KaiJi-Casino/internal/pkg/db/collection"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c *client) CreateGambler(gambler *collection.Gambler) error {
	log.Debug("create gambler: ", gambler.String())

	if _, err := c.Gambler.InsertOne(nil, gambler); err != nil {
		log.Error("fail to insert document: ", err.Error())
		return err
	}

	log.Debug("gambler created: ", gambler.String())
	return nil
}

func (c *client) GetGamblers(filter bson.M, option *options.FindOptions) ([]collection.Gambler, error) {
	log.Debug("get gambler")

	cursor, err := c.Gambler.Find(nil, filter, option)
	if err != nil {
		log.Error("fail to get documents")
		return nil, err
	}
	var gamblers []collection.Gambler
	if err := cursor.All(nil, gamblers); err != nil {
		log.Error("fail to decode document: ", err.Error())
		return nil, err
	}
	return gamblers, nil
}

func (c *client) CountGamblers(filter bson.M) int64 {
	count, err := c.Gambler.CountDocuments(nil, filter)
	if err != nil {
		log.Error("fail to count documents: ", err.Error())
		return -1
	}
	return count
}

func (c client) AppendHistory(gamblerName string, history []collection.GambleHistory) {

}
