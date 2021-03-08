package cmd

import (

  "github.com/spf13/cobra"
  . "github.com/pool-beta/pool-server/pool-cli/api"
)

func init() {
  rootCmd.AddCommand(poolCmd)

  poolCmd.AddCommand(poolGetCmd)
  poolCmd.AddCommand(poolCreateCmd)
  

}

var poolCmd = &cobra.Command{
  Use:   "pool",
  Short: "pool specific request",
  Long:  `Contains commands for pool`,
}

var poolGetCmd = &cobra.Command{
  Use:   "get <username> <password> <pool_id>",
  Short: "Get pool",
  Long:  `Give back pool if user has access`,
  Args: cobra.MinimumNArgs(3),
  Run: func(cmd *cobra.Command, args []string) {
      ctx, err := NewContext(url)
      if err != nil {
        return
      }

      RunPoolGet(ctx, args[0], args[1], args[2])
  },
}

var poolCreateCmd = &cobra.Command{
  Use:   "create <username> <password> <poolname>",
  Short: "Create pool",
  Long:  `Create pool for user`,
  Args: cobra.MinimumNArgs(3),
  Run: func(cmd *cobra.Command, args []string) {
      ctx, err := NewContext(url)
      if err != nil {
        return
      }

      RunPoolCreate(ctx, args[0], args[1], args[2])
  },
}
