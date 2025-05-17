package core

import (
	"fmt"
	. "github.com/romankudravcev/commit-cortex/internal/components"
	"github.com/romankudravcev/commit-cortex/internal/output"
	"github.com/spf13/viper"
)

func List() error {
	viper.SetDefault("repos", []Repo{})

	var repos []Repo
	err := viper.UnmarshalKey("repos", &repos)
	if err != nil {
		return fmt.Errorf("error unmarshalling repos: %v", err)
	}

	if len(repos) == 0 {
		line := output.Color("Currently no repositories are tracked.", output.Red, output.Bold)
		fmt.Println(line)
		return nil
	}

	fmt.Println(output.Color("Tracked repositories:", output.Green, output.Bold))
	for _, repo := range repos {
		prefix := output.Color(fmt.Sprintf("[%s]: ", output.Link(repo.Name, repo.RemoteUrl)), output.Blue, output.Bold)
		path := output.Color(repo.Path, output.Cyan)
		fmt.Printf(prefix + path + "\n")
	}

	return nil
}
