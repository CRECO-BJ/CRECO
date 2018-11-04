package main

import (
	"github.com/urfave/cli"
)

var test bool

var flags = []cli.Flag{
	cli.BoolTFlag{
		Name:        "test",
		Usage:       "only for test",
		Destination: &test,
	},
}
