package db

import (
	"KaiJi-Casino/internal/pkg/db/collection"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c *client) GetHistory(filter bson.M, option *options.FindOptions) (history []collection.N_GambleHistory, err error) {
	log.Debug("get gamble history")

	cursor, err := c.GambleHistory.Find(nil, filter, option)
	if err != nil {
		log.Error("fail to get gamble history: ", err.Error())
		return
	}

	if err := cursor.All(nil, history); err != nil {
		log.Error("fail to decode document: ", err.Error())
	}
	return
}
