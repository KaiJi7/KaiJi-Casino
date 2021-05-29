package gambler

import (
	"KaiJi-Casino/internal/pkg/banker"
	"KaiJi-Casino/internal/pkg/db"
	"KaiJi-Casino/internal/pkg/db/collection"
	"KaiJi-Casino/internal/pkg/strategy"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	BrokenThreshold = 1.0
)

type Gambler struct {
	collection.GamblerData
	Strategy strategy.Strategy
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

		if judge.Winner == banker.WinnerGambler {
			g.MoneyCurrent += judge.Reward
			g.Strategy.OnWin(decision)
			continue
		}

		if judge.Winner == banker.WinnerBanker {
			if g.MoneyCurrent < BrokenThreshold {
				g.OnBroken()
				continue
			}
			g.Strategy.OnLose(decision)
			continue
		}

		if judge.Winner == banker.WinnerTie {
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

func GetGambler(simulationId *primitive.ObjectID) (gamblers []Gambler, err error) {

	gamblersData, err := db.New().ListGambler(simulationId)
	if err != nil {
		log.Error("fail to load gambler: ", err.Error())
		return
	}

	for _, gamblerData := range gamblersData {
		strategyData, err := db.New().GetStrategy(gamblerData.SimulationId)
		if err != nil {
			log.Error("fail to get strategyData: ", err.Error())
			return
		}
		stg, err := strategy.InitStrategy(strategyData.Id)
		if err != nil {
			log.Error("gail to init strategy: ", err.Error())
			return
		}

		gamblers = append(gamblers, Gambler{
			GamblerData: gamblerData,
			Strategy:    stg,
		})
	}
	return
}
