package main

import (
	"KaiJi-Casino/internal/pkg/cli"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	err := cli.InitCli().Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
