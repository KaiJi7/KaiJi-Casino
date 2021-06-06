package cli

import (
	"KaiJi-Casino/internal/pkg/casino"
	"KaiJi-Casino/internal/pkg/configs"
	"KaiJi-Casino/internal/pkg/db"
	"KaiJi-Casino/internal/pkg/db/collection"
	"KaiJi-Casino/internal/pkg/strategy"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
	"os"
)

var (
	newSimulation = &cli.Command{
		Name:    "new-simulation",
		Usage:   "Create a new simulation",
		Aliases: []string{"ns"},
		Flags:   initGamblerFlag,
		Action: func(c *cli.Context) (err error) {
			configs.New()
			log.Debug("read strategy schema file at: ", c.Path("strategy-schema"))

			simulation, dbErr := db.New().CreateSimulation(readSchema(c.Path("strategy-schema")), c.Float64("initial-money"))
			if dbErr != nil {
				log.Error("fail to create simulation: ", dbErr.Error())
				err = dbErr
				return
			}

			if err = casino.CreateGamblers(simulation); err != nil {
				log.Error("fail to init gamblers: ", err.Error())
				return
			}

			log.Debug("gamblers initialized, simulation id: ", simulation.Id)
			casino.Start(c.Int("days"))

			return
		},
	}

	initGamblerFlag = []cli.Flag{
		&cli.Float64Flag{
			Name:    "initial-money",
			Aliases: []string{"m"},
			Usage:   "Initial money for each gambler",
			Value:   100,
		},
		&cli.PathFlag{
			Name:     "strategy-schema",
			Aliases:  []string{"f"},
			Usage:    "Strategy schema file",
			Required: true,
		},
		&cli.IntFlag{
			Name:     "days",
			Aliases:  []string{"d"},
			Usage:    "How many days each gambler to gamble",
			Required: true,
		},
	}
)

func readSchema(path string) (schema map[collection.StrategyName]int) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	d := yaml.NewDecoder(file)

	if err := d.Decode(&schema); err != nil {
		panic(err)
	}

	if !validateSchema(schema) {
		log.Panic("invalid schema")
	}
	return
}

func validateSchema(schema map[collection.StrategyName]int) bool {
	for s, _ := range schema {
		if _, exist := strategy.NameMap[s]; !exist {
			log.Error("unsupported schema: ", s)
			return false
		}
	}
	return true
}
