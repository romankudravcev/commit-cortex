package output

import (
	"fmt"
	"github.com/romankudravcev/commit-cortex/internal/components"
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

func Link(name, url string) string {
	return fmt.Sprintf("\033]8;;%s\033\\%s\033]8;;\033\\", url, name)
}

func PrintReport(report components.Report) {
	headline := Color(fmt.Sprintf("--------------- %s ----------------", strings.ToUpper(report.Repository.Name)), Green, Bold)
	fmt.Println(headline)
	for _, item := range report.ReportItems {
		commitTime := Color(item.Time.Format("2006-01-02 15:04:06"), Gray)
		branchName := Color(item.Branch, Blue, Bold)
		commitMessage := Color(item.Commit, White)
		fmt.Printf("%s (%s): %s", commitTime, branchName, commitMessage)
	}
	fmt.Println()
}
