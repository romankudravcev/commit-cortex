package core

import (
	"fmt"
	"github.com/romankudravcev/commit-cortex/internal/components"
	"github.com/spf13/viper"
)

func Tidy() error {
	var repos []components.Repo
	err := viper.UnmarshalKey("repos", &repos)
	if err != nil {
		return fmt.Errorf("error unmarshalling repos: %v", err)
	}

	notAvailableRepositories, err := components.GetUnavailableRepositories(repos)
	if err != nil {
		return fmt.Errorf("error getting unavailable repositories: %v", err)
	}

	repos = removeSubset(repos, notAvailableRepositories)

	viper.Set("repos", repos)
	err = viper.WriteConfig()
	if err != nil {
		return fmt.Errorf("error writing config: %v", err)
	}

	fmt.Println("Removed unavailable repositories from the config.")
	return nil
}

func removeSubset(fullset, subset []components.Repo) []components.Repo {
	for _, s := range subset {
		for i, f := range fullset {
			if s.Path == f.Path {
				fullset = append(fullset[:i], fullset[i+1:]...)
				break
			}
		}
	}
	return fullset
}
