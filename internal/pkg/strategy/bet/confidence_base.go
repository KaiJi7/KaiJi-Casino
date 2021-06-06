package bet

import (
	"KaiJi-Casino/internal/pkg/db/collection"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"sync"
)

type confidenceBase struct {
	Name string
}

const (
	defaultThreshold = 512
)

var (
	confidenceBaseOnce     sync.Once
	confidenceBaseInstance *confidenceBase
)

func NewConfidenceBase() *confidenceBase {
	confidenceBaseOnce.Do(func() {
		confidenceBaseInstance = &confidenceBase{
			Name: "ConfidenceBase",
		}
		log.Debug("confidence base bet strategy initialized")
	})
	return confidenceBaseInstance
}

func (c confidenceBase) GetDecisions(games []collection.SportsData, parameters map[string]interface{}) (decisions []collection.BetInfo) {
	log.Debug("get confidence base decision")

	t, exist := parameters["threshold"]
	if !exist {
		t = defaultThreshold
	}
	threshold := t.(float64)
	for _, game := range games {
		if betable(game, collection.GambleTypeTotalPoint) {
			if threshold < getConfidence(game.GambleInfo.TotalPoint.Prediction.Over.Population, game.GambleInfo.TotalPoint.Prediction.Under.Population) {
				decisions = append(decisions, c.makeTotalPointDecision(game))
			}
		}

		if betable(game, collection.GambleTypeSpreadPoint) {
			if threshold < getConfidence(game.GambleInfo.SpreadPoint.Prediction.Guest.Population, game.GambleInfo.SpreadPoint.Prediction.Host.Population) {
				decisions = append(decisions, c.makeSpreadPointDecision(game))
			}
		}

		if betable(game, collection.GambleTypeOriginal) {
			if threshold < getConfidence(game.GambleInfo.Original.Prediction.Guest.Population, game.GambleInfo.Original.Prediction.Host.Population) {
				decisions = append(decisions, c.makeOriginalDecision(game))
			}
		}
	}
	return
}

func (c confidenceBase) makeTotalPointDecision(game collection.SportsData) collection.BetInfo {
	bet := collection.BetInfo{
		GameId:     game.Id,
		Banker:     collection.LocalBanker,
		GambleType: collection.GambleTypeTotalPoint,
	}
	if game.GambleInfo.TotalPoint.Prediction.Under.Population < game.GambleInfo.TotalPoint.Prediction.Over.Population {
		bet.BetSide = collection.BetSideOver
		bet.Response = game.GambleInfo.TotalPoint.Response.Over
	} else if game.GambleInfo.TotalPoint.Prediction.Under.Population > game.GambleInfo.TotalPoint.Prediction.Over.Population {
		bet.BetSide = collection.BetSideUnder
		bet.Response = game.GambleInfo.TotalPoint.Response.Under
	} else {
		log.Warn("made tie population on confidence decision")
		candidate := []string{collection.BetSideUnder, collection.BetSideOver}
		bet.BetSide = candidate[rand.Intn(len(candidate))]
	}
	return bet
}

func (c confidenceBase) makeSpreadPointDecision(game collection.SportsData) collection.BetInfo {
	bet := collection.BetInfo{
		GameId:     game.Id,
		Banker:     collection.LocalBanker,
		GambleType: collection.GambleTypeSpreadPoint,
	}
	if game.GambleInfo.SpreadPoint.Prediction.Guest.Population < game.GambleInfo.SpreadPoint.Prediction.Host.Population {
		bet.BetSide = collection.BetSideHost
		bet.Response = game.GambleInfo.SpreadPoint.Response.Host
	} else if game.GambleInfo.SpreadPoint.Prediction.Guest.Population > game.GambleInfo.SpreadPoint.Prediction.Host.Population {
		bet.BetSide = collection.BetSideGuest
		bet.Response = game.GambleInfo.SpreadPoint.Response.Guest
	} else {
		log.Warn("made tie population on confidence decision")
		candidate := []string{collection.BetSideGuest, collection.BetSideHost}
		bet.BetSide = candidate[rand.Intn(len(candidate))]
	}
	return bet
}

func (c confidenceBase) makeOriginalDecision(game collection.SportsData) collection.BetInfo {
	bet := collection.BetInfo{
		GameId:     game.Id,
		Banker:     collection.LocalBanker,
		GambleType: collection.GambleTypeOriginal,
	}
	if game.GambleInfo.Original.Prediction.Guest.Population < game.GambleInfo.Original.Prediction.Host.Population {
		bet.BetSide = collection.BetSideHost
		bet.Response = game.GambleInfo.SpreadPoint.Response.Host
	} else if game.GambleInfo.Original.Prediction.Guest.Population > game.GambleInfo.Original.Prediction.Host.Population {
		bet.BetSide = collection.BetSideGuest
		bet.Response = game.GambleInfo.SpreadPoint.Response.Guest
	} else {
		log.Warn("made tie population on confidence decision")
		candidate := []string{collection.BetSideGuest, collection.BetSideHost}
		bet.BetSide = candidate[rand.Intn(len(candidate))]
	}
	return bet
}

func getConfidence(voteA, voteB int) float64 {
	var major, minor float64
	if voteA < voteB {
		major = float64(voteB)
		minor = float64(voteA)
	} else {
		major = float64(voteA)
		minor = float64(voteB)
	}

	// avoid divide by 0
	if minor == 0 {
		minor = 0.1
	}

	return major / minor * (major - minor)
}
