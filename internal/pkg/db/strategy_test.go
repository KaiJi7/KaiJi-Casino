package db

import (
	"KaiJi-Casino/internal/pkg/configs"
	"KaiJi-Casino/internal/pkg/db/collection"
	"testing"
)

func TestClient_GetStrategyMetaData(t *testing.T) {
	configPath := "../../../configs/config.yaml"
	configs.SetConfigPath(configPath)
	client := New()
	m, e := client.GetStrategyMetaData(collection.StrategyNameLowerResponse)
	if e != nil {
		t.Error(e.Error())
	}
	t.Log(m.Name)
}
