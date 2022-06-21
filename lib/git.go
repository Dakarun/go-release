package lib

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"os"
	"os/exec"
	"strings"
)

func GetCurrentRepo() *git.Repository {
	dir, err := os.Getwd()
	repo, err := git.PlainOpenWithOptions(dir, &git.PlainOpenOptions{DetectDotGit: true})
	if err != nil {
		fmt.Errorf("Failed to open git repository at %s: %s", dir, err)
	}
	return repo
}

func GetCommitMessages(repo *git.Repository) {
	ref, err := repo.Head()
	if err != nil {
		fmt.Errorf("Issue getting head from branch %s", err)
	}
	messages, err := repo.Log(&git.LogOptions{From: ref.Hash()})
	i := 0
	messages.ForEach(func(commit *object.Commit) error {
		i++
		return nil
	})
}

func IsWorkingTreeClean(repo *git.Repository) bool {
	worktree, err := repo.Worktree()
	if err != nil {
		fmt.Errorf("couldn't get worktree: %s", err)
	}
	status, err := worktree.Status()
	if err != nil {
		fmt.Errorf("couldn't get worktree status: %s", err)
	}
	return status.IsClean()
}

// GetProjectRoot TODO: Find a more elegant way to get project root
func GetProjectRoot() (string, error) {
	path, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(path)), nil
}
