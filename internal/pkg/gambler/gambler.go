package gambler

import (
	"KaiJi-Casino/internal/pkg/banker"
	"KaiJi-Casino/internal/pkg/db/collection"
	"KaiJi-Casino/internal/pkg/strategy/bet"
	"KaiJi-Casino/internal/pkg/strategy/put"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Gambler struct {
	Id    *primitive.ObjectID `json:"id"`
	Name  string              `json:"name"`
	Money int                 `json:"money"`

	Capital  float64 `json:"capital"`
	Strategy struct {
		Bet bet.BaseStrategy
		Put put.BaseStrategy
	}
	GambleHistory []collection.GambleHistory
}

func (g Gambler) MakeBet() {
	log.Debug("make bet")

	g.Strategy.Bet.GetDecisions()
}

func (g Gambler) Battle(gambleInfo []collection.SportsData) {
	log.Debug("start battle: ", g.Name)

	decisions := g.Strategy.Bet.GetDecisions(gambleInfo, nil)
	for _, d := range decisions {
		d.Unit = g.Strategy.Put.GetUnit(g.GambleHistory)

		isWin, err := banker.New().Battle(d)
		if err != nil {
			log.Error("fail to battle with banker: ", err.Error())
			continue
		}

		hist := collection.GambleHistory{
			BetInfo:       d,
			Win:           isWin,
			CapitalBefore: g.Capital,
		}
		g.Capital -= float64(d.Unit)

		if isWin {
			log.Debug("gambler ", g.Name, " win the game: ", d.String())
			g.Capital += banker.New().GetResponse(d)
		} else {
			log.Debug("gambler ", g.Name, " lose the game: ", d.String())
		}

		hist.CapitalAfter = g.Capital
		g.GambleHistory = append(g.GambleHistory, hist)
	}
}
