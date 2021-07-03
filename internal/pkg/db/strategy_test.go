package db

import (
	"github.com/KaiJi7/common/structs"
	"github.com/spf13/viper"
	"testing"
)

func TestClient_GetStrategyMetaData(t *testing.T) {
	viper.AddConfigPath("../../../configs")
	client := New()
	m, e := client.GetStrategyMetaData(structs.StrategyNameLowerResponse)
	if e != nil {
		t.Error(e.Error())
	}
	t.Log(m.Name)
}
