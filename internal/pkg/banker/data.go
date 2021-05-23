package banker

import (
	"KaiJi-Casino/internal/pkg/db"
	"KaiJi-Casino/internal/pkg/db/collection"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func (b banker) GetGames(gameType collection.GameType, begin time.Time, end time.Time) (games []collection.SportsGameResult, err error) {
	// TODO: consider result exposure
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

	return db.New().GetGames(filter, nil)
}

func (b banker) GetGambles(gameId *primitive.ObjectID) (gambles []collection.Gambling, err error) {
	log.Debug("get gambles")

	filter := bson.M{
		"game_id": gameId,
	}

	return db.New().GetGambles(filter, nil)
}

func (b banker) GetBettings(gamblingId *primitive.ObjectID) (bets []collection.Betting, err error) {
	log.Debug("get bets")

	filter := bson.M{
		"gambling_id": gamblingId,
	}

	return db.New().GetBets(filter, nil)
}
