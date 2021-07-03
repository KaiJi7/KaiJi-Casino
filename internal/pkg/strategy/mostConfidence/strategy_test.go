package mostConfidence

import (
	"KaiJi-Casino/internal/pkg/db/collection"
	"reflect"
	"testing"
)

func Test_sortDecisionByConfidence(t *testing.T) {
	decisions := []structs.Decision{
		{
			Put: 1.0,
		},
		{
			Put: 2.0,
		},
		{
			Put: 3.0,
		},
	}
	confidenceData := []float64{3, 2, 1}

	expected := []structs.Decision{
		{
			Put: 3.0,
		},
		{
			Put: 2.0,
		},
		{
			Put: 1.0,
		},
	}

	sortedDecisions, _ := sortDecisionByConfidence(decisions, confidenceData)
	for i, expect := range expected {
		if sortedDecisions[i].Put != expect.Put {
			t.Errorf("fail")
		}
	}

}

func Test_sortDecisionByConfidence1(t *testing.T) {
	type args struct {
		decisions      []structs.Decision
		confidenceData []float64
	}
	tests := []struct {
		name                 string
		args                 args
		wantSortedDecisions  []structs.Decision
		wantSortedConfidence []float64
	}{
		{
			name: "normal",
			args: args{
				decisions: []structs.Decision{
					{
						Put: 1.0,
					},
					{
						Put: 2.0,
					},
					{
						Put: 3.0,
					},
				},
				confidenceData: []float64{3, 2, 1},
			},
			wantSortedDecisions: []structs.Decision{
				{
					Put: 3.0,
				},
				{
					Put: 2.0,
				},
				{
					Put: 1.0,
				},
			},
			wantSortedConfidence: []float64{1, 2, 3},
		},
		{
			name: "with nil decision",
			args: args{
				decisions: []structs.Decision{
					{
						Put: 1.0,
					},
					{},
					{
						Put: 3.0,
					},
				},
				confidenceData: []float64{3, 0, 1},
			},
			wantSortedDecisions: []structs.Decision{
				{},
				{
					Put: 3.0,
				},
				{
					Put: 1.0,
				},
			},
			wantSortedConfidence: []float64{0, 1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSortedDecisions, gotSortedConfidence := sortDecisionByConfidence(tt.args.decisions, tt.args.confidenceData)
			if !reflect.DeepEqual(gotSortedDecisions, tt.wantSortedDecisions) {
				t.Errorf("sortDecisionByConfidence() gotSortedDecisions = %v, want %v", gotSortedDecisions, tt.wantSortedDecisions)
			}
			if !reflect.DeepEqual(gotSortedConfidence, tt.wantSortedConfidence) {
				t.Errorf("sortDecisionByConfidence() gotSortedConfidence = %v, want %v", gotSortedConfidence, tt.wantSortedConfidence)
			}
		})
	}
}