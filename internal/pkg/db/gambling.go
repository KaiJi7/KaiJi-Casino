package db

import (
	"github.com/KaiJi7/common/structs"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c client) GetGambling(gamblingId *primitive.ObjectID) (gambling structs.Gambling, err error) {
	log.Debug("get gambling: ", gamblingId.Hex())
	filter := bson.M{
		"_id": gamblingId,
	}
	err = c.Gambling.FindOne(nil, filter).Decode(&gambling)
	return
}

func (c client) GetGambles(filter bson.M, option *options.FindOptions) (documents []structs.Gambling, err error) {
	log.Debug("get gambles")

	cursor, dbErr := c.Gambling.Find(nil, filter, option)
	if dbErr != nil {
		log.Error("fail to get document: ", dbErr.Error())
		err = dbErr
		return
	}
	if err = cursor.All(nil, &documents); err != nil {
		log.Error("fail to decode document: ", err.Error())
		return
	}
	return
}
