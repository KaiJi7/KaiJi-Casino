package banker

import (
	"KaiJi-Casino/internal/pkg/db"
	"fmt"
	"github.com/KaiJi7/common/structs"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Judge struct {
	DecisionId *primitive.ObjectID
	Winner     structs.GambleWinner
	Reward     float64
}

func (b Banker) Judge(decision structs.Decision) (judge Judge, err error) {
	log.Debug("judge")

	gambling, err := db.New().GetGambling(decision.GambleId)
	if err != nil {
		log.Error("fail to get gambling: ", err.Error())
		return
	}

	res, odds, err := b.gamblingResult(gambling)
	if err != nil {
		log.Error("fail to get gambling result: ", err.Error())
		return
	}

	judge.DecisionId = decision.Id

	// TODO: consider tie
	if res == decision.Bet {
		judge.Winner = structs.GambleWinnerGambler
		judge.Reward = odds * decision.Put
	} else {
		judge.Winner = structs.GambleWinnerBanker
		judge.Reward = 0
	}
	return
}

func (b Banker) gamblingResult(gambling structs.Gambling) (result structs.Bet, odds float64, err error) {
	log.Debug("get gambling result: ", gambling.Id)

	game, err := db.New().GetGame(gambling.GameId)
	if err != nil {
		log.Error("fail to get game: ", err.Error())
		return
	}

	// TODO: refine design
	switch gambling.Type {
	case structs.GamblingTypeOriginal:
		if game.Guest.Score < game.Host.Score {
			result = structs.BetHost
		} else if game.Guest.Score > game.Host.Score {
			result = structs.BetGuest
		} else {
			result = structs.BetTie
		}

	case structs.GamblingTypeSpreadPoint:
		gsp := gambling.GetProperty("guest_spread_point").(float64)
		hsp := gambling.GetProperty("host_spread_point").(float64)

		gsc := float64(game.Guest.Score) + gsp
		hsc := float64(game.Host.Score) + hsp

		if gsc < hsc {
			result = structs.BetHost
		} else if gsc > hsc {
			result = structs.BetGuest
		} else {
			result = structs.BetTie
		}

	case structs.GamblingTypeTotalScore:
		threshold := gambling.GetProperty("threshold").(float64)
		totalScore := float64(game.Guest.Score + game.Host.Score)
		if totalScore < threshold {
			result = structs.BetUnder
		} else if totalScore > threshold {
			result = structs.BetOver
		} else {
			result = structs.BetEqual
		}
	default:
		log.Warn("unhandled gambling type: ", gambling.Type)
		err = fmt.Errorf("unhandled gambling type: %s", gambling.Type)
		return
	}

	odds = gambling.GetOdds(result)
	return
}
