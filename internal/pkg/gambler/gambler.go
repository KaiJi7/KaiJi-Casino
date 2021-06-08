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

var (
	brokenChan = make(chan bool, 1)
)

type Gambler struct {
	collection.GamblerData
	Strategy common.Strategy
}

// PlaySince plays gamble with latest n days game
func (g Gambler) PlaySince(wg *sync.WaitGroup, days int) {
	defer wg.Done()

	// get games which use to gamble
	gamesToGamble := make([]collection.SportsGameInfo, 0)

	// for each day since n days before
	for d := -days; d < 0; d++ {

		select {
		case <-brokenChan:
			log.Info(fmt.Sprintf("gambler %s was broken, stop playing and return", g.Id.Hex()))
			return
		default:
			for _, tg := range g.Strategy.TargetGameType() {
				gamesToGamble = append(gamesToGamble, banker.New().GetGames(tg, time.Now().AddDate(0, 0, d), time.Now())...)
			}
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

func (g Gambler) OnBroken() {
	log.Info("gambler: ", g.Id.Hex(), ". was broken.")
	brokenChan <- true
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
		strategyData, dbErr := db.New().GetStrategyData(gamblerData.Id)
		if dbErr != nil {
			log.Error("fail to get strategyData: ", dbErr.Error())
			err = dbErr
			return
		}
		stg, sErr := strategy.GenStrategy(strategyData)
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
