package core

import (
	"fmt"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/romankudravcev/commit-cortex/internal/components"
	"github.com/romankudravcev/commit-cortex/internal/output"
	"github.com/spf13/viper"
	"time"
)

func CreateReport() error {

	viper.SetDefault("repos", []components.Repo{})
	var repos []components.Repo
	err := viper.UnmarshalKey("repos", &repos)
	if err != nil {
		return fmt.Errorf("error unmarshalling repos: %v", err)
	}

	for _, repo := range repos {
		createRepositoryOverview(repo)
	}
	return nil
}

func createRepositoryOverview(repo components.Repo) error {
	gitRepo, err := git.PlainOpen(repo.Path)
	if err != nil {
		return fmt.Errorf("error opening git repository: %v", err)
	}

	refs, err := gitRepo.References()
	if err != nil {
		return fmt.Errorf("error getting references: %v", err)
	}

	report := components.Report{
		Repository: repo,
	}

	err = refs.ForEach(func(ref *plumbing.Reference) error {
		if ref.Type() != plumbing.HashReference || !ref.Name().IsBranch() {
			return nil
		}
		branchName := ref.Name().Short()
		commits, err := gitRepo.Log(&git.LogOptions{From: ref.Hash()})
		if err != nil {
			return fmt.Errorf("error getting commit history for branch %s: %v", branchName, err)
		}
		for {
			commit, err := commits.Next()
			if err != nil {
				break
			}
			currentTime := time.Now()
			if commit.Committer.When.Before(currentTime.Add(-24 * time.Hour)) {
				break
			}
			reportItem := components.ReportItem{
				Author: commit.Author.Name,
				Branch: branchName,
				Commit: commit.Message,
				Time:   commit.Committer.When,
			}
			report.ReportItems = append(report.ReportItems, reportItem)
		}
		return nil
	})
	if err != nil {
		return err
	}

	output.PrintReport(report)

	return nil
}
