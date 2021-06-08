package gambler

import (
	"KaiJi-Casino/internal/pkg/banker"
	"KaiJi-Casino/internal/pkg/db"
	"KaiJi-Casino/internal/pkg/db/collection"
	log "github.com/sirupsen/logrus"
)

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
