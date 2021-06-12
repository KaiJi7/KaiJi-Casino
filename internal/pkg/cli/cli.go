package cli

import (
	"github.com/urfave/cli/v2"
)

const version = "0.1.0"

func InitCli() (app *cli.App) {

	app = &cli.App{
		Name: "KaiJI Casino",
		//Usage:   "Casino Simulator",
		Version: version,
	}

	//app.Flags = initFlag()
	app.Commands = initCommand()
	return
}

func initCommand() []*cli.Command {
	return []*cli.Command{
		createMetaStrategy,
		newSimulation,
		resumeSimulation,
		dropSimulation,
	}
}
