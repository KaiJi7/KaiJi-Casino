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

func (c client) GetGames(filter bson.M, option *options.FindOptions) (documents []collection.SportsGameResult, err error) {
	log.Debug("query games from db: ", filter)

	cursor, err := c.Game.Find(nil, filter, option)
	if err != nil {
		log.Error("fail to get document: ", err.Error())
		return
	}
	if err := cursor.All(nil, &documents); err != nil {
		log.Error("fail to decode documents: ", err.Error())
	}
	return
}
