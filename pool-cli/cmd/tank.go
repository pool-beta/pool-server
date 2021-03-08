package cmd

import (

  "github.com/spf13/cobra"
  . "github.com/pool-beta/pool-server/pool-cli/api"
)

func init() {
  rootCmd.AddCommand(tankCmd)

  tankCmd.AddCommand(tankCreateCmd)
}

var tankCmd = &cobra.Command{
  Use:   "tank",
  Short: "tank specific request",
  Long:  `Contains commands for tank`,
}

var tankCreateCmd = &cobra.Command{
  Use:   "create <username> <password> <tankname>",
  Short: "Create tank",
  Long:  `Create tank for user`,
  Args: cobra.MinimumNArgs(3),
  Run: func(cmd *cobra.Command, args []string) {
      ctx, err := NewContext(url)
      if err != nil {
        return
      }

      RunPoolCreate(ctx, args[0], args[1], args[2])
  },
}