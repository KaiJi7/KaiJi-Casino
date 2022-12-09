package cli

import "github.com/urfave/cli/v2"

var (
	pureBetSimulation = &cli.Command{
		Name:    "pure-bet",
		Usage:   "Simulate based on probability only",
		Aliases: []string{"pb"},
		Flags:   pureBetFlag,
	}

	pureBetFlag = []cli.Flag{
		&cli.Float64Flag{
			Name:    "win-ratio",
			Aliases: []string{"w"},
			Usage:   "Win probability",
			Value:   0.5,
		},
		&cli.Float64Flag{
			Name:    "response",
			Aliases: []string{"r"},
			Usage:   "Response ratio",
			Value:   1.75,
		},
		&cli.IntFlag{
			Name:    "times",
			Aliases: []string{"t"},
			Usage:   "Bet times",
			Value:   100,
		},
		&cli.IntFlag{
			Name:    "player",
			Aliases: []string{"p"},
			Usage:   "Number of players",
			Value:   100,
		},
		&cli.IntFlag{
			Name:    "money",
			Aliases: []string{"m"},
			Usage:   "Initial money for each player",
			Value:   0,
		},
		&cli.IntFlag{
			Name:    "bet",
			Aliases: []string{"b"},
			Usage:   "Bet per game",
			Value:   1,
		},
	}
)
