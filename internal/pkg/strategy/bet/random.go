package bet

import (
	"KaiJi-Casino/internal/pkg/db/collection"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"sync"
	"time"
)

type random struct {
	Name string
}

var (
	randomOnce     sync.Once
	randomInstance *random
	candidate      = []collection.BetInfo{
		{
			Banker:     collection.LocalBanker,
			GambleType: collection.GambleTypeTotalPoint,
			BetSide:    collection.BetSideUnder,
			Confidence: 0,
		},
		{
			Banker:     collection.LocalBanker,
			GambleType: collection.GambleTypeTotalPoint,
			BetSide:    collection.BetSideOver,
			Confidence: 0,
		},
		{
			Banker:     collection.LocalBanker,
			GambleType: collection.GambleTypeSpreadPoint,
			BetSide:    collection.BetSideGuest,
			Confidence: 0,
		},
		{
			Banker:     collection.LocalBanker,
			GambleType: collection.GambleTypeSpreadPoint,
			BetSide:    collection.BetSideHost,
			Confidence: 0,
		},
		{
			Banker:     collection.LocalBanker,
			GambleType: collection.GambleTypeOriginal,
			BetSide:    collection.BetSideGuest,
			Confidence: 0,
		},
		{
			Banker:     collection.LocalBanker,
			GambleType: collection.GambleTypeOriginal,
			BetSide:    collection.BetSideHost,
			Confidence: 0,
		},
	}
)

func NewRandom() *random {
	randomOnce.Do(func() {
		randomInstance = &random{
			Name: "Random",
		}
		log.Debug("random bet strategy initialized")
	})
	return randomInstance
}

func (r *random) GetDecisions(games []collection.SportsData, _ map[string]interface{}) []collection.BetInfo {
	log.Debug("get random decisions")

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(games), func(i, j int) { games[i], games[j] = games[j], games[i] })
	shuffleBet()
	for _, game := range games {
		for _, bet := range candidate {
			if betable(game, gambleTypeMap[bet.GambleType]) {
				return []collection.BetInfo{
					{
						GameId:     game.Id,
						Banker:     bet.Banker,
						GambleType: bet.GambleType,
						BetSide:    bet.BetSide,
						//Unit: bet.Unit,
						//Confidence: bet.Confidence
					},
				}
			} else {
				log.Debug("unbetable game:", bet.GambleType)
			}
		}
	}

	log.Warn("no valid gamble, make no decisions")
	return nil
}

func shuffleBet() {
	rand.Shuffle(len(candidate), func(i, j int) { candidate[i], candidate[j] = candidate[j], candidate[i] })
}
