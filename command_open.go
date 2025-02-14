package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func doOpen(ctx *cli.Context) error {
	fmt.Println("open")
	return nil
}
