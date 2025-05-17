package core

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type Repo struct {
	Path string
	Name string
}

func Add(path string) error {
	path, err := filepath.Abs(path)
	gitPath := fmt.Sprintf("%s/.git", path)

	if err != nil {
		return fmt.Errorf("error getting absolute path: %v", err)
	}

	viper.SetDefault("repos", []Repo{})

	var repos []Repo
	err = viper.UnmarshalKey("repos", &repos)
	if err != nil {
		return fmt.Errorf("error unmarshalling repos: %v", err)
	}

	if err := isExisting(path, gitPath); err != nil {
		return err
	}

	if err := isAdded(repos, gitPath); err != nil {
		return err
	}

	newRepo := Repo{
		Path: gitPath,
		Name: getRepoName(path),
	}

	if err := addRepo(newRepo, repos); err != nil {
		return err
	}

	return nil
}

func addRepo(newRepo Repo, repos []Repo) error {
	repos = append(repos, newRepo)

	viper.Set("repos", repos)
	err := viper.WriteConfig()
	if err != nil {
		return fmt.Errorf("error writing config: %v", err)
	}
	return nil
}

func isExisting(path string, gitPath string) error {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("path does not exist")
		}
	}

	if _, err := os.Stat(gitPath); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("path is not a git repository")
		}
	}
	return nil
}

func isAdded(repos []Repo, gitPath string) error {
	for _, repo := range repos {
		if repo.Path == gitPath {
			return fmt.Errorf("repo already added")
		}
	}
	return nil
}

func getRepoName(path string) string {
	return filepath.Base(path)
}
