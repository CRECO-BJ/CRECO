package main

import (
	"log"
	"os"

	"github.com/creco/console"
	"github.com/creco/node"
	"github.com/urfave/cli"
)

func startBackend(c *cli.Context) error {
	n, err := node.NewNode()
	if err != nil {
		log.Fatal(err)
	}
	if err = n.Run(); err != nil {
		log.Fatal(err)
	}
	return nil
}

func startConsole(c *cli.Context) error {
	console.NewConsole()
}

func start(c *cli.Context) error {
	if err := startBackend(c); err != nil {
		log.Fatal(err)
		return nil
	}

	if err := startConsole(c); err != nil {
		log.Fatal(err)
		return nil
	}
	return nil
}

func main() {
	app := cli.NewApp()

	app.Flags = flags
	app.EnableBashCompletion = true
	app.Commands = Commands
	app.Action = start

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
