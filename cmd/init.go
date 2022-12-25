package main

import (
	"fmt"

	"github.com/urfave/cli/v3"
)

func initCmd(ctx *cli.Context) error {
	fmt.Println("this init func")
	return nil
}
