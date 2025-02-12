package main

import (
	"fmt"

	"github.com/urfave/cli"
)

var commands = []cli.Command{
	commandOpen,
}

var commandOpen = cli.Command{
	Name:  "open",
	Usage: "Open a code-workspace",
	Action: func(c *cli.Context) error {
		fmt.Println("open")
		return nil
	},
}
