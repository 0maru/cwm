package main

import (
	"github.com/urfave/cli/v2"
)

var commands = []*cli.Command{
	commandOpen,
	commandList,
}
