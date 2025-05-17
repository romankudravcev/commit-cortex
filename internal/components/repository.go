package components

import "github.com/romankudravcev/commit-cortex/internal/os"

type Repo struct {
	Path      string
	Name      string
	RemoteUrl string
}

func GetUnavailableRepositories(repos []Repo) ([]Repo, error) {
	var unavailableRepos []Repo
	for _, repo := range repos {
		if err := os.PathExists(repo.Path); err != nil {
			unavailableRepos = append(unavailableRepos, repo)
		}
	}
	return unavailableRepos, nil
}
