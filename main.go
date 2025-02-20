package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

const version = "0.0.1"

var revision = "HEAD"

func main() {
	if err := newApp().Run(os.Args); err != nil {
		exitCode := 1
		if excoder, ok := err.(cli.ExitCoder); ok {
			exitCode = excoder.ExitCode()
		}
		os.Exit(exitCode)
	}
}

func newApp() *cli.App {
	app := &cli.App{
		Name:     "cwm",
		Usage:    "Manage code-workspace",
		Version:  fmt.Sprintf("%s (rev:%s)", version, revision),
		Commands: commands,
		Action: func(ctx *cli.Context) error {
			LoadConfig(ctx)
			return nil
		},
	}
	return app
}
