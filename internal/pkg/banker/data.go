package banker

import (
	"KaiJi-Casino/internal/pkg/db"
	"KaiJi-Casino/internal/pkg/db/collection"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (b banker) GetTodayGames(gameType collection.GameType) []collection.SportsGameInfo {
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

func (b banker) GetGames(gameType collection.GameType, begin time.Time, end time.Time) []collection.SportsGameInfo {
	log.Debug("get games")

	var filter bson.M
	if gameType != collection.GameTypeAll {
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

func (b banker) GetGambles(gameId *primitive.ObjectID) []collection.Gambling {
	log.Debug("get gambles")

	filter := bson.M{
		"game_id": gameId,
	}

	gambles, err := db.New().GetGambles(filter, nil)
	if err != nil {
		log.Error("fail to get gambles: ", err.Error())
		return nil
	}

	return gambles
}

func (b banker) GetBettings(gamblingId *primitive.ObjectID) (bets []collection.Betting, err error) {
	log.Debug("get bets")

	filter := bson.M{
		"gambling_id": gamblingId,
	}

	return db.New().GetBets(filter, nil)
}
