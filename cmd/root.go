package cmd

import (
	"os"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/pool-beta/pool-server/daemon"
)

var port string

// Requires subcommand since no Run defined
var rootCmd = &cobra.Command{
  Use:   "pool-server",
  Short: "POOL Server contains ther server implementation for POOL",
  Long: `POOL Server features commands for running the server`,
}

var runCmd =  &cobra.Command{
  Use:   "run",
  Short: "POOL Server contains ther server implementation for POOL",
  Long: `POOL Server features commands for running the server`,
  Run: func(cmd *cobra.Command, args []string) {
  		daemon.Run(port)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}


func init() {

	rootCmd.PersistentFlags().StringVarP(&port, "port", "p", "8000", "port for POOL Server")
	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
	
	rootCmd.AddCommand(runCmd)
}