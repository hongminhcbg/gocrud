package main

import (
	"fmt"

	"github.com/urfave/cli/v3"
)

var Commit string
var BuildDate string
var Builder string

func version(ctx *cli.Context) error {
	fmt.Printf("Version: %s\nBuildDate: %s\nBuilder: %s", Commit, BuildDate, Builder)
	return nil
}
