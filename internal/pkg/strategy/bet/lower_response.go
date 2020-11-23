package bet

import (
	"KaiJi-Casino/internal/pkg/db/collection"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"sync"
)

// bet side with lower response ratio
type lowerResponse struct {
	Name string
}

var (
	lowerResponseOnce     sync.Once
	lowerResponseInstance *lowerResponse
)

func NewLowerResponse() *lowerResponse {
	lowerResponseOnce.Do(func() {
		lowerResponseInstance = &lowerResponse{
			Name: "LowerResponse",
		}
		log.Debug("lower response bet strategy initialized")
	})
	return lowerResponseInstance
}

func (l lowerResponse) GetDecisions(games []collection.SportsData, _ map[string]interface{}) []collection.BetInfo {
	log.Debug("get lower response decisions")
	var decisions []collection.BetInfo

	for _, game := range games {

		if betable(game, collection.GambleTypeTotalPoint) {
			decisions = append(decisions, l.makeTotalPointDecision(game))
		}

		if betable(game, collection.GambleTypeSpreadPoint) {
			decisions = append(decisions, l.makeSpreadPointDecision(game))
		}

		if betable(game, collection.GambleTypeOriginal) {
			decisions = append(decisions, l.makeOriginalDecision(game))
		}
	}

	return decisions
}

func (l lowerResponse) makeTotalPointDecision(game collection.SportsData) collection.BetInfo {
	bet := collection.BetInfo{
		GameId:     game.Id,
		Banker:     collection.LocalBanker,
		GambleType: collection.GambleTypeTotalPoint,
	}
	if game.GambleInfo.TotalPoint.Response.Under < game.GambleInfo.TotalPoint.Response.Over {
		bet.BetSide = collection.BetSideUnder
		bet.Response = game.GambleInfo.TotalPoint.Response.Under
	} else if game.GambleInfo.TotalPoint.Response.Under > game.GambleInfo.TotalPoint.Response.Over {
		bet.BetSide = collection.BetSideOver
		bet.Response = game.GambleInfo.TotalPoint.Response.Over
	} else {
		candidate := []string{collection.BetSideUnder, collection.BetSideOver}
		bet.BetSide = candidate[rand.Intn(len(candidate))]
	}
	return bet
}

func (l lowerResponse) makeSpreadPointDecision(game collection.SportsData) collection.BetInfo {
	bet := collection.BetInfo{
		GameId:     game.Id,
		Banker:     collection.LocalBanker,
		GambleType: collection.GambleTypeSpreadPoint,
	}
	if game.GambleInfo.SpreadPoint.Response.Guest < game.GambleInfo.SpreadPoint.Response.Host {
		bet.BetSide = collection.BetSideGuest
		bet.Response = game.GambleInfo.SpreadPoint.Response.Guest
	} else if game.GambleInfo.SpreadPoint.Response.Guest > game.GambleInfo.SpreadPoint.Response.Host {
		bet.BetSide = collection.BetSideHost
		bet.Response = game.GambleInfo.SpreadPoint.Response.Host
	} else {
		candidate := []string{collection.BetSideGuest, collection.BetSideHost}
		bet.BetSide = candidate[rand.Intn(len(candidate))]
	}
	return bet
}

func (l lowerResponse) makeOriginalDecision(game collection.SportsData) collection.BetInfo {
	bet := collection.BetInfo{
		GameId:     game.Id,
		Banker:     collection.LocalBanker,
		GambleType: collection.GambleTypeOriginal,
	}
	if game.GambleInfo.Original.Response.Guest < game.GambleInfo.Original.Response.Host {
		bet.BetSide = collection.BetSideGuest
		bet.Response = game.GambleInfo.SpreadPoint.Response.Guest
	} else if game.GambleInfo.Original.Response.Guest > game.GambleInfo.Original.Response.Host {
		bet.BetSide = collection.BetSideHost
		bet.Response = game.GambleInfo.SpreadPoint.Response.Host
	} else {
		candidate := []string{collection.BetSideGuest, collection.BetSideHost}
		bet.BetSide = candidate[rand.Intn(len(candidate))]
	}
	return bet
}
