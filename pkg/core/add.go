package core

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/romankudravcev/commit-cortex/internal/components"
	"github.com/romankudravcev/commit-cortex/internal/os"
	"github.com/spf13/viper"
)

func Add(path string) error {
	path, err := filepath.Abs(path)
	gitPath := fmt.Sprintf("%s/.git", path)
	if err != nil {
		return fmt.Errorf("error getting absolute path: %v", err)
	}

	viper.SetDefault("repos", []components.Repo{})

	var repos []components.Repo
	err = viper.UnmarshalKey("repos", &repos)
	if err != nil {
		return fmt.Errorf("error unmarshalling repos: %v", err)
	}

	if err := os.PathExists(path); err != nil {
		return err
	}

	if err := os.PathExists(gitPath); err != nil {
		return err
	}

	if err := isAdded(repos, gitPath); err != nil {
		return err
	}

	// TODO handle error correctly
	remoteUrl, _ := getRemoteUrl(path)

	newRepo := components.Repo{
		Path:      gitPath,
		Name:      getRepoName(path),
		RemoteUrl: remoteUrl,
	}

	if err := addRepo(newRepo, repos); err != nil {
		return err
	}

	return nil
}

func addRepo(newRepo components.Repo, repos []components.Repo) error {
	repos = append(repos, newRepo)

	viper.Set("repos", repos)
	err := viper.WriteConfig()
	if err != nil {
		return fmt.Errorf("error writing config: %v", err)
	}
	return nil
}

func isAdded(repos []components.Repo, gitPath string) error {
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

func getRemoteUrl(path string) (string, error) {
	// TODO: Refactor, maybe read the git config file instead of spawning bash
	cmd := fmt.Sprintf("git -C %s remote get-url origin", path)
	out, err := exec.Command("bash", "-c", cmd).Output()

	url := string(out)

	url = strings.Replace(url, "git@github.com:", "https://github.com/", 1)
	url = strings.TrimSuffix(url, ".git\n")

	if err != nil {
		return "", fmt.Errorf("error getting remote url: %v", err)
	}

	return url, nil
}
