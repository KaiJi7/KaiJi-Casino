package db

import (
	"KaiJi-Casino/internal/pkg/db/collection"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// get games with the whole data
func (c *client) GetGames(filter bson.M, option *options.FindOptions) (documents []collection.SportsData, err error) {
	log.Debug("query games from db: ", filter)

	cursor, err := c.SportsData.Find(nil, filter, option)
	if err != nil {
		log.Error("fail to get document: ", err.Error())
		return
	}
	if err := cursor.All(nil, &documents); err != nil {
		log.Error("fail to decode documents: ", err.Error())
	}
	return
}

// get gamble info only, without result and judgement
func (c *client) GetGambleInfo(gameId primitive.ObjectID) (sportsData collection.SportsData, err error) {
	log.Debug("get gamble info: ", gameId.Hex())

	filter := bson.M{
		"_id": gameId,
	}
	opt := options.FindOne().SetProjection(
		bson.M{
			"gamble_info.total_point.judgement":         0,
			"gamble_info.total_point.prediction.major":  0,
			"gamble_info.spread_point.judgement":        0,
			"gamble_info.spread_point.prediction.major": 0,
			"gamble_info.original.judgement":            0,
			"gamble_info.original.prediction.major":     0,
		},
	)

	if err := c.SportsData.FindOne(nil, filter, opt).Decode(&sportsData); err != nil {
		log.Error("fail to decode document: ", err.Error())
	}
	return
}

func (c *client) CountGames(filter bson.M) int64 {
	count, err := c.SportsData.CountDocuments(nil, filter)
	if err != nil {
		log.Error("fail to count documents: ", err.Error())
		return -1
	}
	return count
}
