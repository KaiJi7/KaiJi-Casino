package banker

import (
	"KaiJi-Casino/internal/pkg/cache"
	"KaiJi-Casino/internal/pkg/db"
	"KaiJi-Casino/internal/pkg/db/collection"
	"KaiJi-Casino/internal/pkg/util"
	"fmt"
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

func (b banker) GetGameResult(gameId primitive.ObjectID) (gambleInfo collection.GambleInfo, err error) {
	log.Debug("get game result: ", gameId.Hex())

	c, _ := cache.New().Get(gameId.Hex())
	if c != nil {
		log.Debug("had cached data")
		obj, err := util.BytesToStruct(c)
		if err != nil {
			log.Error("fail to convert cached data to struct: ", err.Error())
		} else {
			gambleInfo = obj.(collection.GambleInfo)
			return
		}
	}

	filter := bson.M{
		"_id": gameId,
	}
	games, err := db.New().GetGames(filter, nil)
	if err != nil {
		log.Error("fail to get games: ", err.Error())
		return
	}

	if len(games) != 1 {
		log.Error("unexpected match count: ", len(games))
		err = fmt.Errorf("unexpected match count: %d", len(games))
		return
	}

	gambleInfo = *games[0].GambleInfo

	// make cache
	data, err := util.StructToBytes(games[0].GambleInfo)
	if err != nil {
		log.Error("fail to convert gamble info to byte array for cache: ", err.Error())
		return
	}
	_ = cache.New().Set(gameId.Hex(), data)
	return
}

func (b banker) Battle(decision collection.BetInfo) (isGamblerWin bool, err error) {
	gameResult, err := b.GetGameResult(decision.GameId)
	if err != nil {
		log.Error("fail to get game result: ", err.Error())
		return
	}

	switch decision.GambleType {
	case collection.GambleTypeOriginal:
		isGamblerWin = *gameResult.Original.Judgement == decision.BetSide
		return
	case collection.GambleTypeTotalPoint:
		isGamblerWin = *gameResult.TotalPoint.Judgement == decision.BetSide
		return
	case collection.GambleTypeSpreadPoint:
		isGamblerWin = *gameResult.SpreadPoint.Judgement == decision.BetSide
		return
	default:
		log.Warn("no matched gamble type: ", decision.GambleType)
		err = fmt.Errorf("no matched gamble type: %s", decision.GambleType)
		return
	}
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
