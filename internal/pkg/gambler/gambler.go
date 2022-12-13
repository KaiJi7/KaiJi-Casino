package gambler

import (
	"KaiJi-Casino/internal/pkg/banker"
	"KaiJi-Casino/internal/pkg/strategy/common"
	"fmt"
	"github.com/KaiJi7/common/structs"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

const (
	BrokenThreshold = 1.0
)

type Gambler struct {
	structs.GamblerData
	Strategy common.Strategy
	IsBroken bool
}

// PlaySince plays gamble with latest n days game
func (g *Gambler) PlaySince(wg *sync.WaitGroup, days int) {
	defer wg.Done()

	// for each day since n days before
	for d := -days; d < 0; d++ {
		if g.IsBroken {
			log.Info(fmt.Sprintf("gamble %s was broken", g.Id.Hex()))
			return
		}
		gamesToGamble := make([]structs.SportsGameInfo, 0)
		for _, tg := range g.Strategy.TargetGameType() {
			gamesToGamble = append(gamesToGamble, banker.New().GetGames(tg, time.Now().AddDate(0, 0, d), time.Now().AddDate(0, 0, d+1))...)
		}
		//select {
		//case <-brokenChan:
		//	log.Info(fmt.Sprintf("gambler %s was broken, stop playing and return", g.Id.Hex()))
		//	return
		//default:
		//
		//}

		g.play(gamesToGamble)
	}
}

func (g *Gambler) play(games []structs.SportsGameInfo) {
	log.Debug(fmt.Sprintf("gambler [%s] play %d games", g.Id.Hex(), len(games)))

	// get gambles info
	gambles := make([]structs.Gambling, 0)
	for _, game := range games {
		gambles = append(gambles, banker.New().GetGambles(game.Id, true)...)
	}

	// make decision
	decisions := g.makeDecision(gambles)

	// handle decision
	g.handleDecision(decisions)
}

func (g *Gambler) OnBroken() {
	log.Info("gambler: ", g.Id.Hex(), ". was broken.")
	g.IsBroken = true
	return
}
