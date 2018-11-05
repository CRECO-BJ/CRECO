package main

import (
	"fmt"
	"log"
	"os"

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

func start(c *cli.Context) error {
	if err := startBackend(c); err != nil {
		return nil
	}
	fmt.Println("Started")
	return nil
}

func main() {
	app := cli.NewApp()

	app.Flags = flags
	app.Commands = Commands
	app.Action = start

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
