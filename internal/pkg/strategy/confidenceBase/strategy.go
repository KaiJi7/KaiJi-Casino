package confidenceBase

import (
	"KaiJi-Casino/internal/pkg/banker"
	"KaiJi-Casino/internal/pkg/db/collection"
	"KaiJi-Casino/internal/pkg/strategy/common"
	log "github.com/sirupsen/logrus"
)

type Strategy struct {
	collection.StrategyData
	confidenceType common.ConfidenceType
	threshold      float64
}

func New(data collection.StrategyData) common.Strategy {

	cts := data.Properties["confidence_type"].(string)
	ct := common.ConfidenceType(cts)
	if _, exist := common.Calculator[ct]; !exist {
		log.Error("invalid confidence type: ", ct)
		panic("invalid confidence type")
	}

	return Strategy{
		StrategyData:   data,
		confidenceType: ct,
		threshold:      data.Properties["threshold"].(float64),
	}
}

func (s Strategy) TargetGameType() []collection.GameType {
	return []collection.GameType{collection.GameTypeAll}
}

func (s Strategy) MakeDecision(gambles []collection.Gambling) (decisions []collection.Decision) {
	for _, gamble := range gambles {
		betsInfo, err := banker.New().GetBettings(gamble.Id)
		if err != nil {
			log.Error("fail to get bets: ", err.Error())
			continue
		}

		for _, bets := range betsInfo {
			side, confidence := common.GetConfidence(bets, s.confidenceType)
			if s.threshold < confidence {
				decision := collection.Decision{
					StrategyId: s.Id,
					GambleId:   gamble.Id,
					Bet:        side,
					Put:        1,
				}
				decisions = append(decisions, decision)
			}
		}
	}
	return
}

func (s Strategy) OnWin(decision collection.Decision) {

}

func (s Strategy) OnLose(decision collection.Decision) {

}

func (s Strategy) OnTie(decision collection.Decision) {

}
