/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/romankudravcev/commit-cortex/pkg/core"

	"github.com/spf13/cobra"
)

// tidyCmd represents the tidy command
var tidyCmd = &cobra.Command{
	Use:   "tidy",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := core.Tidy()
		if err != nil {
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(tidyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tidyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tidyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
