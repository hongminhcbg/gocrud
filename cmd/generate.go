package main

import (
	"fmt"

	"github.com/urfave/cli/v3"
)

func generate(ctx *cli.Context) error {
	fmt.Println("this generate func", fileInput)
	return nil
}
