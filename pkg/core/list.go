package core

import (
	"fmt"
	"github.com/spf13/viper"
	"strconv"
	"strings"
)

var (
	Reset     = "\033[0m"
	Black     = "\033[30m"
	Red       = "\033[31m"
	Green     = "\033[32m"
	Yellow    = "\033[33m"
	Blue      = "\033[34m"
	Magenta   = "\033[35m"
	Cyan      = "\033[36m"
	Gray      = "\033[37m"
	White     = "\033[97m"
	Bold      = "\033[1m"
	Italic    = "\033[3m"
	Underline = "\033[4m"
	Invert    = "\033[7m"
)

func Color(input interface{}, color ...string) string {
	var s string
	c := ""
	for i := range color {
		c = c + color[i]
	}
	switch v := input.(type) {
	case int:
		s = c + strconv.Itoa(v) + Reset
	case bool:
		s = c + strconv.FormatBool(v) + Reset
	case []string:
		s = c + strings.Join(v, ", ") + Reset
	case string:
		s = c + v + Reset
	default:
		fmt.Printf("Unsupported type provided to Color func - %T\n", v)
	}
	return s
}

func List() error {
	viper.SetDefault("repos", []Repo{})

	var repos []Repo
	err := viper.UnmarshalKey("repos", &repos)
	if err != nil {
		return fmt.Errorf("error unmarshalling repos: %v", err)
	}

	if len(repos) == 0 {
		output := Color("Currently no repositories are tracked.", Red, Bold)
		fmt.Println(output)
		return nil
	}

	fmt.Println(Color("Tracked repositories:", Green, Bold))
	for _, repo := range repos {
		prefix := Color(fmt.Sprintf("[%s]: ", repo.Name), Blue, Bold)
		path := Color(repo.Path, Cyan)
		fmt.Printf(prefix + path + "\n")
	}

	return nil
}
