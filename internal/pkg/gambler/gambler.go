package gambler

import (
	"KaiJi-Casino/internal/pkg/banker"
	"KaiJi-Casino/internal/pkg/db"
	"KaiJi-Casino/internal/pkg/db/collection"
	"KaiJi-Casino/internal/pkg/strategy"
	"KaiJi-Casino/internal/pkg/strategy/common"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sync"
	"time"
)

const (
	BrokenThreshold = 1.0
)

type Gambler struct {
	collection.GamblerData
	Strategy common.Strategy
}

//func (g Gambler) Play() {
//
//	// get games which use to gamble
//	gamesToGamble := make([]collection.SportsGameInfo, 0)
//	for _, tg := range g.Strategy.TargetGameType() {
//		gamesToGamble = append(gamesToGamble, banker.New().GetTodayGames(tg)...)
//	}
//
//	g.play(gamesToGamble)
//}

// PlaySince plays gamble with latest n days game
func (g Gambler) PlaySince(wg *sync.WaitGroup, days int) {
	defer wg.Done()

	// get games which use to gamble
	gamesToGamble := make([]collection.SportsGameInfo, 0)

	// for each day since n days before
	for d := -days; d < 0; d++ {

		for _, tg := range g.Strategy.TargetGameType() {
			gamesToGamble = append(gamesToGamble, banker.New().GetGames(tg, time.Now().AddDate(0, 0, d), time.Now())...)
		}

		g.play(gamesToGamble)
	}
}

func (g Gambler) play(games []collection.SportsGameInfo) {
	log.Debug(fmt.Sprintf("gambler [%s] play %d games", g.Id.Hex(), len(games)))

	// get gambles info
	gambles := make([]collection.Gambling, 0)
	for _, game := range games {
		gambles = append(gambles, banker.New().GetGambles(game.Id)...)
	}

	// make decision
	decisions := g.makeDecision(gambles)

	// handle decision
	g.handleDecision(decisions)
}

func (g *Gambler) makeDecision(gambles []collection.Gambling) (decisions []collection.Decision) {
	allDecisions := g.Strategy.MakeDecision(gambles)

	for _, decision := range allDecisions {
		if g.MoneyCurrent > decision.Put {
			g.MoneyCurrent -= decision.Put

			// TODO: consider error handling
			decision, err := db.New().SaveDecision(decision)
			if err != nil {
				log.Error("fail to save decision: ", err.Error())
				continue
			}
			decisions = append(decisions, decision)
		} else {
			log.Info("not enough money to bet decision: ", decision.String())
			continue
		}
	}

	return
}

func (g *Gambler) handleDecision(decisions []collection.Decision) {
	log.Debug("handle decision")

	bk := banker.New()
	for _, decision := range decisions {
		judge, err := bk.Judge(decision)
		if err != nil {
			log.Error("fail to judge decision: ", err.Error())
			continue
		}

		hist := collection.GambleHistory{
			DecisionId: decision.Id,
			Winner: judge.Winner,
			MoneyBefore: g.MoneyCurrent,
		}

		switch judge.Winner {
		case collection.GambleWinnerGambler:
			g.MoneyCurrent += judge.Reward
			g.Strategy.OnWin(decision)

		case collection.GambleWinnerBanker:
			if g.MoneyCurrent < BrokenThreshold {
				g.OnBroken()
				continue
			}
			g.Strategy.OnLose(decision)

		case collection.GambleWinnerTie:
			g.Strategy.OnTie(decision)

		default:
			log.Warn("unhandled judge result: ", judge.Winner)
		}

		// set moneyAfter after handled decision
		hist.MoneyAfter = g.MoneyCurrent

		if err := db.New().SaveHistory(hist); err != nil {
			log.Error("fail to save gamble history: ", err.Error())
		}
	}
	return
}

func (g Gambler) OnBroken() {
	log.Info("gambler: ", g.Id.Hex(), ". was broken.")

	// TODO: append to broken chan to stop gambling
	return
}

func GetGamblers(simulationId *primitive.ObjectID) (gamblers []Gambler, err error) {
	log.Debug("get gamblers by simulation id: ", simulationId.Hex())

	gamblersData, dbErr := db.New().ListGambler(simulationId)
	if dbErr != nil {
		log.Error("fail to load gambler: ", dbErr.Error())
		err = dbErr
		return
	}

	for _, gamblerData := range gamblersData {
		strategyData, dbErr := db.New().GetStrategy(gamblerData.Id)
		if dbErr != nil {
			log.Error("fail to get strategyData: ", dbErr.Error())
			err = dbErr
			return
		}
		stg, sErr := strategy.GetStrategy(strategyData)
		if sErr != nil {
			log.Error("gail to init strategy: ", sErr.Error())
			err = sErr
			return
		}

		gamblers = append(gamblers, Gambler{
			GamblerData: gamblerData,
			Strategy:    stg,
		})
	}
	return
}
