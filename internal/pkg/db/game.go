package db

import (
	"KaiJi-Casino/internal/pkg/db/collection"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c client) GetGame(gameId *primitive.ObjectID) (game collection.SportsGameResult, err error) {
	log.Debug("get game: ", gameId.Hex())

	filter := bson.M{
		"_id": gameId,
	}

	err = c.Game.FindOne(nil, filter).Decode(&game)
	return
}

func (c client) GetGamesInfo(filter bson.M, option *options.FindOptions) (documents []collection.SportsGameInfo, err error) {
	log.Debug("query games info from db: ", filter)

	cursor, dbErr := c.Game.Find(nil, filter, option)
	if dbErr != nil {
		log.Error("fail to get document: ", dbErr.Error())
		err = dbErr
		return
	}
	if err = cursor.All(nil, &documents); err != nil {
		log.Error("fail to decode documents: ", err.Error())
	}
	return
}

func (c client) GetGamesResult(filter bson.M, option *options.FindOptions) (documents []collection.SportsGameResult, err error) {
	log.Debug("query games result from db: ", filter)

	cursor, dbErr := c.Game.Find(nil, filter, option)
	if dbErr != nil {
		log.Error("fail to get document: ", dbErr.Error())
		err = dbErr
		return
	}
	if err = cursor.All(nil, &documents); err != nil {
		log.Error("fail to decode documents: ", err.Error())
	}
	return
}
