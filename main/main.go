package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func startBackend(c *cli.Context) error {
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
