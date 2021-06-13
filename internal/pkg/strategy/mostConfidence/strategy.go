package mostConfidence

import (
	"KaiJi-Casino/internal/pkg/banker"
	"KaiJi-Casino/internal/pkg/db/collection"
	"KaiJi-Casino/internal/pkg/strategy/common"
	log "github.com/sirupsen/logrus"
	"sort"
)

type Strategy struct {
	collection.StrategyData
	Threshold float64
	Limit     int
}

func New(data collection.StrategyData) common.Strategy {

	return Strategy{
		StrategyData: data,
		Threshold:    data.Properties["threshold"].(float64),
		Limit:        int(data.Properties["limit"].(int32)),
	}
}

func (s Strategy) TargetGameType() []collection.GameType {
	return []collection.GameType{collection.GameTypeAll}
}

func (s Strategy) MakeDecision(gambles []collection.Gambling) []collection.Decision {
	decisions := make([]collection.Decision, 0, s.Limit)
	confidenceData := make([]float64, 0, s.Limit)
	for _, gamble := range gambles {
		betsInfo, err := banker.New().GetBettings(gamble.Id)
		if err != nil {
			log.Error("fail to get bets: ", err.Error())
			continue
		}

		for _, bets := range betsInfo {
			side, confidence := common.GetConfidence(bets, common.ConfidenceTypeLinear)

			if len(decisions) < s.Limit || confidenceData[0] < confidence {
				decision := collection.Decision{
					StrategyId: s.Id,
					GambleId:   gamble.Id,
					Bet:        side,
					Put:        1,
				}
				//decisions[0] = decision
				//confidenceData[0] = confidence
				decisions = append(decisions, decision)
				confidenceData = append(confidenceData, confidence)
			}
			decisions, confidenceData = sortDecisionByConfidence(decisions, confidenceData)

		}
	}
	return decisions
}

func (s Strategy) OnWin(decision collection.Decision) {

}

func (s Strategy) OnLose(decision collection.Decision) {

}

func (s Strategy) OnTie(decision collection.Decision) {

}

type dc struct {
	decision   collection.Decision
	confidence float64
}

func sortDecisionByConfidence(decisions []collection.Decision, confidenceData []float64) (sortedDecisions []collection.Decision, sortedConfidence []float64) {
	comb := make([]dc, len(decisions))
	for i, decision := range decisions {
		comb[i] = dc{
			decision:   decision,
			confidence: confidenceData[i],
		}
	}
	sort.SliceStable(comb, func(i, j int) bool {
		return comb[i].confidence < comb[j].confidence
	})

	for _, dc := range comb {
		sortedDecisions = append(sortedDecisions, dc.decision)
		sortedConfidence = append(sortedConfidence, dc.confidence)
	}
	return
}
