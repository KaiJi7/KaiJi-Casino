package common

import (
	"github.com/KaiJi7/common/structs"
	log "github.com/sirupsen/logrus"
	"math"
)

const (
	ConfidenceTypeLinear   ConfidenceType = "linear"
	ConfidenceTypeWeighted ConfidenceType = "weighted"
)

var (
	Calculator = map[ConfidenceType]func(betting structs.Betting) (structs.Bet, float64){
		ConfidenceTypeLinear:   linearConfidence,
		ConfidenceTypeWeighted: weightedConfidence,
	}
)

type ConfidenceType string

func GetConfidence(betsInfo structs.Betting, confidenceType ConfidenceType) (side structs.Bet, confidence float64) {
	log.Debug("get confidence of bets: ", betsInfo.Id.Hex(), ", type: ", confidenceType)
	return Calculator[confidenceType](betsInfo)
}

// Calculate bets confidence by the quantity ratio of all, 0 < confidence < 1
func linearConfidence(betsInfo structs.Betting) (side structs.Bet, confidence float64) {
	sum := 0.0
	for _, b := range betsInfo.Bet {
		sum += float64(b.Quantity)
	}

	for _, b := range betsInfo.Bet {
		if c := float64(b.Quantity) / sum; confidence < c {
			confidence = c
			side = b.Side
		}
	}
	return
}

func weightedConfidence(betsInfo structs.Betting) (side structs.Bet, confidence float64) {
	sum := 0.0
	for _, b := range betsInfo.Bet {
		sum += float64(b.Quantity)
	}

	for _, b := range betsInfo.Bet {
		if c := math.Pow(float64(b.Quantity), 2) / sum; confidence < c {
			confidence = c
			side = b.Side
		}
	}
	return
}
