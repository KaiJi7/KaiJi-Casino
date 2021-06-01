package gambler

import (
	"KaiJi-Casino/internal/pkg/banker"
	"KaiJi-Casino/internal/pkg/db"
	"KaiJi-Casino/internal/pkg/db/collection"
	"KaiJi-Casino/internal/pkg/strategy"
	"KaiJi-Casino/internal/pkg/strategy/common"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	BrokenThreshold = 1.0
)

type Gambler struct {
	collection.GamblerData
	Strategy common.Strategy
}

func (g *Gambler) MakeDecision(gambles []collection.Gambling) (decisions []collection.Decision) {
	allDecisions := g.Strategy.MakeDecision(gambles)

	for _, decision := range allDecisions {
		if g.MoneyCurrent > decision.Put {
			g.MoneyCurrent -= decision.Put
			decisions = append(decisions, decision)
		} else {
			log.Info("not enough money to bet decision: ", decision.String())
			continue
		}
	}

	return
}

func (g *Gambler) HandleDecision(decisions []collection.Decision) {
	log.Debug("handle decision")

	bk := banker.New()
	for _, decision := range decisions {
		judge, err := bk.Judge(decision)
		if err != nil {
			log.Error("fail to judge decision: ", err.Error())
			continue
		}

		if judge.Winner == collection.GambleWinnerGambler {
			g.MoneyCurrent += judge.Reward
			g.Strategy.OnWin(decision)
			continue
		}

		if judge.Winner == collection.GambleWinnerBanker {
			if g.MoneyCurrent < BrokenThreshold {
				g.OnBroken()
				continue
			}
			g.Strategy.OnLose(decision)
			continue
		}

		if judge.Winner == collection.GambleWinnerTie {
			g.Strategy.OnTie(decision)
			continue
		}
	}
}

func (g Gambler) OnBroken() {
	log.Info("gambler: ", g.Id.Hex(), ". was broken.")

	// TODO: append to broken chan to stop gambling
	return
}

func GetGamblers(simulationId *primitive.ObjectID) (gamblers []Gambler, err error) {

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
		stg, sErr := strategy.InitStrategy(strategyData.Id)
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
