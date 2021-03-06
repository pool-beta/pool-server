package cmd

import (
	"os"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var url string

// Requires subcommand since no Run defined
var rootCmd = &cobra.Command{
  Use:   "pool-cli",
  Short: "POOL Client contains commands for interacting with the POOL Server",
  Long: `POOL Client features commands for creating/viewing/editing 
  pools and users. It is mainly to be used by developers to test and debug API requests.
  For more information, visit https://github.com/pool-beta.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}


func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "pool-cli config file")
	rootCmd.PersistentFlags().StringVarP(&url, "url", "u", "http://localhost:8000", "URL to POOL Server")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
	viper.BindPFlag("url", rootCmd.PersistentFlags().Lookup("url"))

}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}
	}
}