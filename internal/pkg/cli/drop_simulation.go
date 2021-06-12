package cli

import (
	"KaiJi-Casino/internal/pkg/configs"
	"KaiJi-Casino/internal/pkg/db"
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

var (
	dropSimulation = &cli.Command{
		Name:    "drop-simulation",
		Usage:   "Drop all simulation-related collections.",
		Aliases: []string{"ds"},
		Action: func(c *cli.Context) (err error) {
			fmt.Printf("It's danger action, type y to confirm drop all the simulation data: ")

			text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
			text = strings.Replace(text, "\n", "", -1)
			if text == "y" {
				configs.New()
				log.Debug("drop all the simulation data")
				_ = db.New().Simulation.Drop(nil)
				_ = db.New().Gambler.Drop(nil)
				_ = db.New().Strategy.Drop(nil)
				_ = db.New().StrategyMeta.Drop(nil)
				_ = db.New().Decision.Drop(nil)
				_ = db.New().GambleHistory.Drop(nil)
			} else {
				log.Debug("no drop")
			}
			return
		},
	}
)
