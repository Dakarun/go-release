package lib

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"os"
)

func GetCurrentRepo() {
	dir, err := os.Getwd()
	repo, err := git.PlainOpenWithOptions(dir, &git.PlainOpenOptions{DetectDotGit: true})
	if err != nil {
		fmt.Errorf("Failed to open git repository at %s: %s", dir, err)
	}

	ref, err := repo.Head()
	fmt.Println(ref.Name().String())
}
