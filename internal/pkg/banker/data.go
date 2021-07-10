package banker

import (
	"KaiJi-Casino/internal/pkg/db"
	"github.com/KaiJi7/common/structs"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (b banker) GetTodayGames(gameType structs.GameType) []structs.SportsGameInfo {
	log.Debug("get today games")

	filter := bson.M{
		"type": gameType,
		"start_time": bson.M{
			"$gt": time.Now().AddDate(0, 0, -1),
		},
	}

	games, err := db.New().GetGamesInfo(filter, nil)
	if err != nil {
		log.Error("fail to get games info: ", err.Error())
		return nil
	}

	return games
}

func (b banker) GetGames(gameType structs.GameType, begin time.Time, end time.Time) []structs.SportsGameInfo {
	log.Debug("get games")

	var filter bson.M
	if gameType != structs.GameTypeAll {
		filter = bson.M{
			"type": gameType,
			"start_time": bson.M{
				"$gt": begin,
				"$lt": end,
			},
		}
	} else {
		filter = bson.M{
			"start_time": bson.M{
				"$gt": begin,
				"$lt": end,
			},
		}
	}

	games, err := db.New().GetGamesInfo(filter, nil)
	if err != nil {
		log.Error("fail to get games info: ", err.Error())
		return nil
	}

	return games
}

func (b banker) GetGambles(gameId *primitive.ObjectID, betableOnly bool) []structs.Gambling {
	log.Debug("get gambles")

	filter := bson.M{
		"game_id": gameId,
	}

	gambles, err := db.New().GetGambles(filter, nil)
	if err != nil {
		log.Error("fail to get gambles: ", err.Error())
		return nil
	}

	// return gambles with unbetables
	if !betableOnly {
		return gambles
	}

	// filter out betable gambles
	var betableGambles []structs.Gambling
	for _, gamble := range gambles {
		if gamble.Betable() {
			betableGambles = append(betableGambles, gamble)
		}
	}
	return betableGambles
}

func (b banker) GetBettings(gamblingId *primitive.ObjectID) (bets []structs.Betting, err error) {
	log.Debug("get bets")

	filter := bson.M{
		"gambling_id": gamblingId,
	}

	return db.New().GetBets(filter, nil)
}
