package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func doOpen(ctx *cli.Context) error {
	fmt.Println("open")
	return nil
}
