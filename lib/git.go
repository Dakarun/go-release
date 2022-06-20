package lib

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"os"
)

func GetCurrentRepo() *git.Repository {
	dir, err := os.Getwd()
	repo, err := git.PlainOpenWithOptions(dir, &git.PlainOpenOptions{DetectDotGit: true})
	if err != nil {
		fmt.Errorf("Failed to open git repository at %s: %s", dir, err)
	}

	ref, err := repo.Head()
	fmt.Println(ref.Name())
	return repo
}

func GetCommitMessages(repo *git.Repository) {
	ref, err := repo.Head()
	if err != nil {
		fmt.Errorf("Issue getting head from branch %s", err)
	}
	messages, err := repo.Log(&git.LogOptions{From: ref.Hash()})
	messages.ForEach(func(commit *object.Commit) error {
		fmt.Println(commit.Message)
		return nil
	})
}
