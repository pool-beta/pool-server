package cmd

import (

  "github.com/spf13/cobra"
  . "github.com/pool-beta/pool-server/pool-cli/api"
)

func init() {
  rootCmd.AddCommand(userCmd)

  userCmd.AddCommand(userListCmd)
  userCmd.AddCommand(userCreateCmd)

}

var userCmd = &cobra.Command{
  Use:   "user",
  Short: "User specific request",
  Long:  `Contains commands for user`,
}

var userListCmd = &cobra.Command{
  Use:   "list",
  Short: "Generate user list",
  Long:  `Creates a list of all users`,
  Run: func(cmd *cobra.Command, args []string) {
      ctx, err := NewContext(url)
      if err != nil {
        return
      }

      RunUserList(ctx)
  },
}

var userCreateCmd = &cobra.Command{
  Use:   "create <username> <password>",
  Short: "Create user",
  Long:  `Creates a new user`,
  Args: cobra.MinimumNArgs(2),
  Run: func(cmd *cobra.Command, args []string) {
      ctx, err := NewContext(url)
      if err != nil {
        return
      }

      RunUserCreate(ctx, args[0], args[1])
  },
}