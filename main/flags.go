package main

import (
	"github.com/urfave/cli"
)

var test bool
var genesis string

var flags = []cli.Flag{
	cli.BoolTFlag{
		Name:        "test",
		Usage:       "only for test",
		Destination: &test,
	},
	cli.StringFlag{
		Name:        "genesis",
		Usage:       "genesis block JSON file",
		Destination: &genesis,
	},
}
