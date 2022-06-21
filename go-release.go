package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	go_release "go-release/lib"
	"log"
	"os"
	"strings"
)

const ChangelogFilename = "CHANGELOG.md"

func main() {
	app := &cli.App{
		Name:  "git-release",
		Usage: "Release shit",
		Action: func(c *cli.Context) error {
			repo := go_release.GetCurrentRepo()
			if !go_release.IsWorkingTreeClean(repo) {
				fmt.Println("There are uncommited changes. Please commit and merge these before running go-release")
				os.Exit(1)
			}
			rootPath, err := go_release.GetProjectRoot()
			if err != nil {
				fmt.Errorf("encountered exception getting project root path: %s", err)
			}

			ChangelogPath := strings.Join([]string{rootPath, ChangelogFilename}, "/")
			if !FileExists(ChangelogPath) {
				fmt.Printf("File does not exist, creating: %s\n", ChangelogPath)
				CreateFile(ChangelogPath)
			}

			go_release.GetCommitMessages(repo)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
