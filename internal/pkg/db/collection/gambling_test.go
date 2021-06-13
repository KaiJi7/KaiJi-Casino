package collection

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
	"testing"
)

func TestGambling_SortedOdds(t *testing.T) {
	odds := []float64{
		1.5, 1.6, 1.7,
	}
	type fields struct {
		Id         *primitive.ObjectID
		Type       GamblingType
		GameId     *primitive.ObjectID
		Odds       []Odds
		Properties []struct {
			Name  string      `json:"name" bson:"name"`
			Value interface{} `json:"value" bson:"value"`
		}
	}
	tests := []struct {
		name   string
		fields fields
		want   []Odds
	}{
		{
			name: "normal",
			fields: fields{
				Odds: []Odds{
					{
						Bet: BetGuest,
						Odds: &odds[0],
					},
					{
						Bet: BetHost,
						Odds: &odds[1],
					},
				},
			},
			want: []Odds{
				{
					Bet: BetGuest,
					Odds: &odds[0],
				},
				{
					Bet: BetHost,
					Odds: &odds[1],
				},
			},
		},
		{
			name: "reverse",
			fields: fields{
				Odds: []Odds{
					{
						Bet: BetGuest,
						Odds: &odds[1],
					},
					{
						Bet: BetHost,
						Odds: &odds[0],
					},
				},
			},
			want: []Odds{
				{
					Bet: BetHost,
					Odds: &odds[0],
				},
				{
					Bet: BetGuest,
					Odds: &odds[1],
				},
			},
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Gambling{
				Id:         tt.fields.Id,
				Type:       tt.fields.Type,
				GameId:     tt.fields.GameId,
				Odds:       tt.fields.Odds,
				Properties: tt.fields.Properties,
			}
			if got := g.SortedOdds(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortedOdds() = %v, want %v", got, tt.want)
			}
		})
	}
}
