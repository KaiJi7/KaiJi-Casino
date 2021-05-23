package db

import (
	"KaiJi-Casino/internal/pkg/db/collection"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c client) GetBets(filter bson.M, option *options.FindOptions) (documents []collection.Betting, err error) {
	log.Debug("get gambles")

	cursor, err := c.Gambling.Find(nil, filter, option)
	if err != nil {
		log.Error("fail to get document: ", err.Error())
		return
	}
	if err := cursor.All(nil, documents); err != nil {
		log.Error("fail to decode document: ", err.Error())
		return
	}
	return
}
