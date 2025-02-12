package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
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
	app := cli.NewApp()
	app.Name = "cwm"
	app.Usage = "Manage code-workspace"
	app.Version = fmt.Sprintf("%s (rev:%s)", version, revision)
	app.Commands = commands
	return app
}
