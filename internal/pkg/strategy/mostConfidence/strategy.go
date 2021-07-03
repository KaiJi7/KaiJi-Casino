package mostConfidence

import (
	"KaiJi-Casino/internal/pkg/banker"
	"KaiJi-Casino/internal/pkg/strategy/common"
	"github.com/KaiJi7/common/structs"
	log "github.com/sirupsen/logrus"
	"sort"
)

type Strategy struct {
	structs.StrategyData
	Threshold float64
	Limit     int
}

func New(data structs.StrategyData) common.Strategy {

	return Strategy{
		StrategyData: data,
		Threshold:    data.Properties["threshold"].(float64),
		Limit:        int(data.Properties["limit"].(int32)),
	}
}

func (s Strategy) TargetGameType() []structs.GameType {
	return []structs.GameType{structs.GameTypeAll}
}

func (s Strategy) MakeDecision(gambles []structs.Gambling) []structs.Decision {
	decisions := make([]structs.Decision, 0, s.Limit)
	confidenceData := make([]float64, 0, s.Limit)
	for _, gamble := range gambles {
		betsInfo, err := banker.New().GetBettings(gamble.Id)
		if err != nil {
			log.Error("fail to get bets: ", err.Error())
			continue
		}

		for _, bets := range betsInfo {
			side, confidence := common.GetConfidence(bets, common.ConfidenceTypeLinear)

			decision := structs.Decision{
				StrategyId: s.Id,
				GambleId:   gamble.Id,
				Bet:        side,
				Put:        1,
			}

			if len(decisions) < s.Limit  {
				decisions = append(decisions, decision)
				confidenceData = append(confidenceData, confidence)
			} else if confidenceData[0] < confidence {
				decisions[0] = decision
				confidenceData[0] = confidence
				decisions, confidenceData = sortDecisionByConfidence(decisions, confidenceData)
			}
		}
	}
	return decisions
}

func (s Strategy) OnWin(decision structs.Decision) {

}

func (s Strategy) OnLose(decision structs.Decision) {

}

func (s Strategy) OnTie(decision structs.Decision) {

}

type dc struct {
	decision   structs.Decision
	confidence float64
}

func sortDecisionByConfidence(decisions []structs.Decision, confidenceData []float64) (sortedDecisions []structs.Decision, sortedConfidence []float64) {
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
