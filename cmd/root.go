package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	rootCmd = &cobra.Command{
		Use:   "commit-cortex",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {

	home, err := os.UserHomeDir()
	if err != nil {
		err = fmt.Errorf("error getting home dir: %v", err)
		cobra.CheckErr(err)
	}

	viper.AddConfigPath(home)
	fileType := "json"
	fileName := ".commit-cortex"
	viper.SetConfigType(fileType)
	viper.SetConfigName(fileName)

	configPath := home + string(os.PathSeparator) + fileName + "." + fileType

	if _, err := os.Stat(configPath); err != nil {
		if os.IsNotExist(err) {
			err = viper.SafeWriteConfig()
			if err != nil {
				err = fmt.Errorf("failed to write: %v", err)
				cobra.CheckErr(err)
			}
		}
	}

	if err := viper.ReadInConfig(); err != nil {
		err = fmt.Errorf("error reading config file: %v", err)
		cobra.CheckErr(err)
	}
}
