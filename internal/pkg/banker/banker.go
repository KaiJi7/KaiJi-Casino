package banker

import (
	"KaiJi-Casino/internal/pkg/cache"
	"KaiJi-Casino/internal/pkg/db"
	"KaiJi-Casino/internal/pkg/db/collection"
	"KaiJi-Casino/internal/pkg/util"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
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

func (b banker) GetGambleInfo(gameId primitive.ObjectID) (collection.SportsData, error) {
	return db.New().GetGambleInfo(gameId)
}

func (b banker) GetGameResult(gameId primitive.ObjectID) *collection.GambleInfo {
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

func (b banker) GetResponse(bet collection.BetInfo) (response float64) {
	log.Debug("get response: ", bet.String())

	c, _ := cache.New().Get(bet.String())
	if c != nil {
		log.Debug("had cached data")
		return util.BytesToFloat64(c)
	}

	data, err := b.GetGambleInfo(bet.GameId)
	if err != nil {
		log.Error("fail to get gamble info: ", err.Error())
		return
	}
	r := reflect.ValueOf(data)
	f := reflect.Indirect(r).FieldByName(bet.GambleType)
	if reflect.ValueOf(f).IsZero() {
		log.Debug("unbetable gamble: ", bet.String())
		return
	}

	gameResp := reflect.Indirect(f).FieldByName("Response")
	sideResp := reflect.Indirect(gameResp).FieldByName(bet.BetSide)
	if reflect.ValueOf(sideResp).IsZero() {
		log.Error("invalid response, ", bet.String())
	}

	response = reflect.Indirect(sideResp).Float()
	_ = cache.New().Set(bet.String(), util.Float64ToBytes(response))
	return
}
