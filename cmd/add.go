package cmd

import (
	"github.com/romankudravcev/commit-cortex/pkg/core"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [path]",
	Short: "Add a repository",
	Long: `Add a repository to commit-cortex:
cc add path`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := "."
		if len(args) > 0 {
			path = args[0]
		}

		err := core.Add(path)
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
