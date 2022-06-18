package main

import (
	"github.com/urfave/cli/v2"
	go_release "go-release/lib"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "git-release",
		Usage: "Release shit",
		Action: func(c *cli.Context) error {
			go_release.GetCurrentRepo()
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}