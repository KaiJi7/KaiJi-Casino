package banker

import (
	"KaiJi-Casino/internal/pkg/db"
	"KaiJi-Casino/internal/pkg/db/collection"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (b banker) Judge(gamblingId *primitive.ObjectID, bet collection.Bet, amount int) (result Winner, reward float64, err error) {
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
		result = WinnerGambler
		reward = odds * float64(amount)
	} else {
		result = WinnerBanker
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
	case collection.GamblingTypeOriginal:
		if game.Guest.Score < game.Host.Score {
			result = collection.BetHost
		} else if game.Guest.Score > game.Host.Score {
			result = collection.BetGuest
		} else {
			result = collection.BetTie
		}

	case collection.GamblingTypeSpreadPoint:
		gsp := gambling.GetProperty("guest_spread_point").(float64)
		hsp := gambling.GetProperty("host_spread_point").(float64)

		gsc := float64(game.Guest.Score) + gsp
		hsc := float64(game.Host.Score) + hsp

		if gsc < hsc {
			result = collection.BetHost
		} else if gsc > hsc {
			result = collection.BetGuest
		} else {
			result = collection.BetTie
		}

	case collection.GamblingTypeTotalScore:
		threshold := gambling.GetProperty("threshold").(float64)
		totalScore := float64(game.Guest.Score + game.Host.Score)
		if totalScore < threshold {
			result = collection.BetUnder
		} else if totalScore > threshold {
			result = collection.BetOver
		} else {
			result = collection.BetEqual
		}
	default:
		log.Warn("unhandled gambling type: ", gambling.Type)
		err = fmt.Errorf("unhandled gambling type: %s", gambling.Type)
		return
	}

	odds = gambling.GetOdds(result)
	return
}
