package banker

import (
	"KaiJi-Casino/internal/pkg/db/collection"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"reflect"
	"testing"
)

func Test_banker_GetGambleInfo(t *testing.T) {
	_ = os.Setenv("CONFIG_PATH", "../../../configs/config.yaml")
	type args struct {
		gameId primitive.ObjectID
	}
	oId, _ := primitive.ObjectIDFromHex("5f8daae32df55becef3317e5")
	tests := []struct {
		name string
		args args
		want *collection.SportsData
	}{
		{
			name: "test",
			args: args{
				gameId: oId,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &banker{}
			if got := b.GetGambleInfo(tt.args.gameId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGambleInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
