package banker

import (
	"KaiJi-Casino/internal/pkg/db"
	"KaiJi-Casino/internal/pkg/db/collection"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sync"
	"time"
)

const (
	JUDGE_RESULT_BANKER  JudgeResult = "banker"  // banker win, gambler lose
	JUDGE_RESULT_GAMBLER JudgeResult = "gambler" // gambler win, banker lose
	JUDGE_RESULT_TIE     JudgeResult = "tie"
)

type JudgeResult string

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

func (b banker) GetGames(gameType collection.GameType, begin time.Time, end time.Time) (games []collection.SportsGameResult, err error) {
	// TODO: consider result exposure
	log.Debug("get games")

	var filter bson.M
	if gameType != collection.GAME_TYPE_ALL {
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

func (b banker) Judge(gamblingId *primitive.ObjectID, bet collection.Bet, amount int) (result JudgeResult, reward float64, err error) {
	log.Debug("judge")

	gambling, err := db.New().GetGambling(gamblingId)
	if err != nil {
		log.Error("fail to get gambling: ", err.Error())
		return
	}

	res, odds, err := b.gamblingResult(gambling)
	if err != nil {
		log.Error("fail to get gambling result: ", err.Error())
		return
	}

	if res == bet {
		result = JUDGE_RESULT_GAMBLER
		reward = odds * float64(amount)
	} else {
		result = JUDGE_RESULT_BANKER
		reward = 0
	}
	return
}

func (b banker) gamblingResult(gambling collection.Gambling) (result collection.Bet, odds float64, err error) {
	log.Debug("get gambling result: ", gambling.Id)

	game, err := db.New().GetGame(gambling.GameId)
	if err != nil {
		log.Error("fail to get game: ", err.Error())
		return
	}

	// TODO: refine design
	switch gambling.Type {
	case collection.GAMBLING_TYPE_ORIGINAL:
		if game.Guest.Score < game.Host.Score {
			result = collection.BET_HOST
		} else if game.Guest.Score > game.Host.Score {
			result = collection.BET_GUEST
		} else {
			result = collection.BET_TIE
		}

	case collection.GAMBLING_TYPE_SPREAD_POINT:
		gsp := gambling.GetProperty("guest_spread_point").(float64)
		hsp := gambling.GetProperty("host_spread_point").(float64)

		gsc := float64(game.Guest.Score) + gsp
		hsc := float64(game.Host.Score) + hsp

		if gsc < hsc {
			result = collection.BET_HOST
		} else if gsc > hsc {
			result = collection.BET_GUEST
		} else {
			result = collection.BET_TIE
		}

	case collection.GAMBLING_TYPE_TOTAL_SCORE:
		threshold := gambling.GetProperty("threshold").(float64)
		totalScore := float64(game.Guest.Score + game.Host.Score)
		if totalScore < threshold {
			result = collection.BET_UNDER
		} else if totalScore > threshold {
			result = collection.BET_OVER
		} else {
			result = collection.BET_EQUAL
		}
	default:
		log.Warn("unhandled gambling type: ", gambling.Type)
		err = fmt.Errorf("unhandled gambling type: %s", gambling.Type)
		return
	}

	odds = gambling.GetOdds(result)
	return
}
