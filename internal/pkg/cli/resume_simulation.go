package cli

import (
	"KaiJi-Casino/internal/pkg/casino"
	"KaiJi-Casino/internal/pkg/configs"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var (
	resumeSimulation = &cli.Command{
		Name:    "resume-simulation",
		Usage:   "Resume simulation to gamble with the latest (today) game",
		Aliases: []string{"rs"},
		Flags:   resumeSimulationFlag,
		Action: func(c *cli.Context) (err error) {
			configs.New()

			if err = casino.LoadGamblers(c.String("simulation-id")); err != nil {
				log.Error("fail to init gamblers: ", err.Error())
				return
			}

			casino.Start(c.Int("days"))
			return
		},
	}

	resumeSimulationFlag = []cli.Flag{
		&cli.StringFlag{
			Name:     "simulation-id",
			Aliases:  []string{"s"},
			Usage:    "Simulation ID",
			Required: true,
		},
	}
)
