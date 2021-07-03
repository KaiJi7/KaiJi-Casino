package confidenceBase

import (
	"KaiJi-Casino/internal/pkg/db/collection"
	"KaiJi-Casino/internal/pkg/strategy/common"
	"testing"
)

func TestStrategy_MakeDecision(t *testing.T) {
	data := structs.StrategyData{
		Name: structs.StrategyNameConfidenceBase,
		Properties: map[string]interface{}{
			"confidence_type": common.ConfidenceTypeLinear,
			"threshold": 0.5,
		},
	}

	_ = data
	// TODO: implement test
	//gambles := structs.Gambling{
	//
	//}
	//
	//s := New(data)

}
