package cmd

import (

  "github.com/spf13/cobra"
  . "github.com/pool-beta/pool-server/pool-cli/api"
)

func init() {
  rootCmd.AddCommand(drainCmd)

  drainCmd.AddCommand(drainCreateCmd)
}

var drainCmd = &cobra.Command{
  Use:   "drain",
  Short: "drain specific request",
  Long:  `Contains commands for drain`,
}

var drainCreateCmd = &cobra.Command{
  Use:   "create <username> <password> <tankname>",
  Short: "Create drain",
  Long:  `Create drain for user`,
  Args: cobra.MinimumNArgs(3),
  Run: func(cmd *cobra.Command, args []string) {
      ctx, err := NewContext(url)
      if err != nil {
        return
      }

      RunPoolCreate(ctx, args[0], args[1], args[2])
  },
}