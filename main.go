package main

import (
	"os"

	server "app/src"
	"app/src/database"
	"app/src/lib/logger"

	"github.com/urfave/cli"
)

var (
	client *cli.App
)

func init() {
	client = cli.NewApp()
	client.Name = ""
	client.Usage = ""
	client.Version = "0.0.0"
}

func main() {
	var logger = logger.NewLogger("main")

	client.Commands = []cli.Command{
		// RUN: server
		server.StartServer(),

		// RUN: run migrate
		database.Migration(),

		// RUN: rollback last migration
		database.Rollback(),

		// RUN: drop database
		database.DropDatabase(),
	}

	// Run the CLI app
	err := client.Run(os.Args)
	if err != nil {
		logger.Fatalf(err.Error())
	}
}
