package bet

import (
	"KaiJi-Casino/internal/pkg/db"
	"KaiJi-Casino/internal/pkg/db/collection"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"reflect"
	"testing"
)

func Test_random_GetDecisions(t *testing.T) {
	_ = os.Setenv("CONFIG_PATH", "../../../../configs/config.yaml")
	type fields struct {
		Name string
	}
	type args struct {
		games []collection.SportsData
	}
	oId, _ := primitive.ObjectIDFromHex("5f8daae32df55becef3317e5")
	filter := bson.M{
		"_id": oId,
	}
	games, _ := db.New().GetGames(filter, nil)
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []collection.BetInfo
	}{
		{
			name: "valid gamble",
			fields: fields{
				Name: "Random",
			},
			args: args{
				games: games,
			},
			want: []collection.BetInfo{},
		},
		{
			name: "no gamble",
			fields: fields{
				Name: "Random",
			},
			args: args{
				games: []collection.SportsData{},
			},
			want: nil,
		},
	}

	t.Run(tests[0].name, func(t *testing.T) {
		r := &random{
			Name: tests[0].fields.Name,
		}
		if got := r.GetDecisions(tests[0].args.games); got == nil {
			t.Errorf("GetDecisions() = %v, want non nil", got)
		}
	})

	for _, tt := range tests[1:] {
		t.Run(tt.name, func(t *testing.T) {
			r := &random{
				Name: tt.fields.Name,
			}
			if got := r.GetDecisions(tt.args.games); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDecisions() = %v, want %v", got, tt.want)
			}
		})
	}
}
