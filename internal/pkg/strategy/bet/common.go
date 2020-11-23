package bet

import (
	"KaiJi-Casino/internal/pkg/db/collection"
	log "github.com/sirupsen/logrus"
	"reflect"
)

func betable(game collection.SportsData, gambleType string) bool {

	r := reflect.ValueOf(game.GambleInfo)
	//f := reflect.Indirect(r).FieldByName(bet.GambleType)
	f := reflect.Indirect(r).FieldByName(gambleType)
	if reflect.ValueOf(f).IsZero() {
		log.Debug("unbetable gamble: ", gambleType, ", game id: ", game.Id.Hex())
		return false
	}

	// it's betable if response ration exist
	resp := reflect.Indirect(f).FieldByName("Response")
	if reflect.Indirect(resp).NumField() > 0 {
		return reflect.Indirect(resp).Field(0).Float() > 0
	} else {
		return false
	}
}

var gambleTypeMap = map[string]string{
	collection.GambleTypeTotalPoint:  "TotalPoint",
	collection.GambleTypeSpreadPoint: "SpreadPoint",
	collection.GambleTypeOriginal:    "Original",
}
