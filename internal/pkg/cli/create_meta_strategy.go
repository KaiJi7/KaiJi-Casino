package cli

import (
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
	}

	createMetaStrategy = &cli.Command{
		Name: "create-meta-strategy",
		Usage: "Create meta strategy",
		Aliases: []string{"cm"},
		Action: func(c *cli.Context) error {
			for _, m := range metaStrategies {
				if err := strategy.CreateMetaStrategy(m); err != nil {
					return err
				}
			}
			return nil
		},
	}
)
