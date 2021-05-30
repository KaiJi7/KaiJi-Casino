package cli

import (
	"KaiJi-Casino/internal/pkg/casino"
	"KaiJi-Casino/internal/pkg/configs"
	"KaiJi-Casino/internal/pkg/db"
	"KaiJi-Casino/internal/pkg/strategy"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
	"os"
)

var (
	initGambler = &cli.Command{
		Name:    "init-gambler",
		Aliases: []string{"ig"},
		Flags:   initGamblerFlag,
		Action: func(c *cli.Context) (err error) {
			configs.New()
			log.Debug("read strategy schema file at: ", c.Path("strategy-schema"))

			simulation, err := db.New().CreateSimulation(readSchema(c.Path("strategy-schema")), c.Float64("initial-money"))
			if err != nil {
				log.Error("fail to create simulation: ", err.Error())
				return
			}

			if err := casino.InitGamblers(simulation); err != nil {
				log.Error("fail to init gamblers: ", err.Error())
				return
			}

			return
		},
	}

	initGamblerFlag = []cli.Flag{
		&cli.Float64Flag{
			Name: "initial-money",
			Aliases: []string{"m"},
			Usage: "Initial money for each gambler",
			Value: 100,
		},
		&cli.PathFlag{
			Name:     "strategy-schema",
			Aliases:  []string{"f"},
			Usage:    "Strategy schema file",
			Required: true,
		},
	}
)

func readSchema(path string) (schema map[strategy.Name]int){
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	d := yaml.NewDecoder(file)

	if err := d.Decode(&schema); err != nil {
		panic(err)
	}
	return
}
