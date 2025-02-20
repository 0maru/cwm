package main

import (
	"github.com/urfave/cli/v2"
)

var commands = []*cli.Command{
	commandOpen,
	commandList,
}

var commandOpen = &cli.Command{
	Name:   "open",
	Usage:  "Open a code-workspace",
	Action: doOpen,
}

var commandList = &cli.Command{
	Name:   "list",
	Usage:  "List all code-workspaces",
	Action: doList,
}
