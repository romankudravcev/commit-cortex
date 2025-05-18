package cmd

import (
	"os"

	"github.com/romankudravcev/commit-cortex/pkg/core"
	"github.com/spf13/cobra"
)

// scanCmd represents the add command
var scanCmd = &cobra.Command{
	Use:   "scan [path]",
	Short: "Scan a repository",
	Long: `Scan a repository to commit-cortex:
cc scan path`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path, err := os.UserHomeDir()
		if err != nil {
			cobra.CheckErr(err)
		}
		if len(args) > 0 {
			path = args[0]
		}

		err = core.Scan(path)
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}
