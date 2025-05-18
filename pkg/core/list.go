package core

import (
	"fmt"
	"slices"

	"github.com/romankudravcev/commit-cortex/internal/components"
	"github.com/romankudravcev/commit-cortex/internal/output"
	"github.com/spf13/viper"
)

func List() error {
	viper.SetDefault("repos", []components.Repo{})

	var repos []components.Repo
	err := viper.UnmarshalKey("repos", &repos)
	if err != nil {
		return fmt.Errorf("error unmarshalling repos: %v", err)
	}

	if len(repos) == 0 {
		line := output.Color("Currently no repositories are tracked.", output.Red, output.Bold)
		fmt.Println(line)
		return nil
	}

	notAvailableRepositories, err := components.GetUnavailableRepositories(repos)
	if err != nil {
		return fmt.Errorf("error getting unavailable repositories: %v", err)
	}

	fmt.Println(output.Color("Tracked repositories:", output.Green, output.Bold))
	for _, repo := range repos {
		if slices.Contains(notAvailableRepositories, repo) {
			continue
		}

		prefix := output.Color(fmt.Sprintf("[%s]: ", output.Link(repo.Name, repo.RemoteUrl)), output.Blue, output.Bold)
		path := output.Color(repo.Path, output.Cyan)
		fmt.Println(prefix + path)
	}

	if len(notAvailableRepositories) > 0 {
		fmt.Println()
		fmt.Println(output.Color("Not found repositories: (remove with `cc tidy` if not needed anymore)", output.Red, output.Bold))

		for _, repo := range notAvailableRepositories {
			prefix := output.Color(fmt.Sprintf("[%s]: ", output.Link(repo.Name, repo.RemoteUrl)), output.Blue, output.Bold)
			path := output.Color(repo.Path, output.Cyan)
			fmt.Println(prefix + path)
		}
	}

	return nil
}
