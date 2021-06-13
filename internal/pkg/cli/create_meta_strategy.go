package cli

import (
	"KaiJi-Casino/internal/pkg/configs"
	"KaiJi-Casino/internal/pkg/db/collection"
	"KaiJi-Casino/internal/pkg/strategy"
	"github.com/urfave/cli/v2"
)

var (
	metaStrategies = []collection.StrategyMeta{
		{
			Name:        collection.StrategyNameLowerResponse,
			Description: "Bet each games with lower odds.",
			Properties:  nil,
		},
		{
			Name:        collection.StrategyNameLowestResponse,
			Description: "Bet a game with the lowest odds.",
			Properties:  nil,
		},
		{
			Name:        collection.StrategyNameConfidenceBase,
			Description: "Bet games based on confidence, where the confidence was based on vote quantity.",
			Properties: []struct {
				Name string `json:"name" bson:"name"`
				Type string `json:"type" bson:"type"` // int, float, string
			}{
				{
					Name: "confidence_type",
					Type: "string",
				},
				{
					Name: "threshold",
					Type: "float",
				},
			},
		},
		{
			Name:        collection.StrategyNameMostConfidence,
			Description: "Bet games based on confidence, where the confidence was based on vote quantity.",
			Properties: []struct {
				Name string `json:"name" bson:"name"`
				Type string `json:"type" bson:"type"` // int, float, string
			}{
				{
					Name: "threshold",
					Type: "float",
				},
				{
					Name: "limit",
					Type: "int",
				},
			},
		},
	}

	createMetaStrategy = &cli.Command{
		Name:    "create-meta-strategy",
		Usage:   "Create meta strategy",
		Aliases: []string{"cm"},
		Action: func(c *cli.Context) error {
			configs.New()
			for _, m := range metaStrategies {
				if err := strategy.CreateMetaStrategy(m); err != nil {
					return err
				}
			}
			return nil
		},
	}
)
