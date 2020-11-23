package banker

import (
	"KaiJi-Casino/internal/pkg/db"
	"KaiJi-Casino/internal/pkg/db/collection"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sync"
)

type banker struct{}

var (
	once     sync.Once
	instance *banker
)

func New() *banker {
	once.Do(func() {
		instance = &banker{}
		log.Debug("banker initialized")
	})
	return instance
}

func (b *banker) GetGambleInfo(gameId primitive.ObjectID) (collection.SportsData, error) {
	return db.New().GetGambleInfo(gameId)
}

func (b *banker) GetGameResult(gameId primitive.ObjectID) *collection.GambleInfo {
	log.Debug("get game result: ", gameId.Hex())

	filter := bson.M{
		"_id": gameId,
	}
	games, err := db.New().GetGames(filter, nil)
	if err != nil {
		log.Error("fail to get games: ", err.Error())
		return nil
	}

	if len(games) != 1 {
		log.Error("unexpected match count: ", len(games))
		return nil
	}

	return games[0].GambleInfo
}
