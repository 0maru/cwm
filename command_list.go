package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func doList(ctx *cli.Context) error {
	fmt.Println("list")
	return nil
}
