package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func doList(ctx *cli.Context) error {
	fmt.Println("list")
	return nil
}
