package cmd

import (

  "github.com/spf13/cobra"
  . "github.com/pool-beta/pool-server/pool-cli/api"
)

func init() {
  rootCmd.AddCommand(testCmd)

  testCmd.AddCommand(testSetupCmd)
  testCmd.AddCommand(testResetCmd)
}

var testCmd = &cobra.Command{
  Use:   "test",
  Short: "Dev Test Specific Requests",
  Long:  `Contains commands for setting up a mock POOL Space for testing`,
}

var testSetupCmd = &cobra.Command{
  Use:   "setup",
  Short: "Quick Pool Server Setup for Testing",
  Long:  `Creates dummy users and pools`,
  Run: func(cmd *cobra.Command, args []string) {
      ctx, err := NewContext(url)
      if err != nil {
        return
      }

      RunTestSetup(ctx)
  },
}

var testResetCmd = &cobra.Command{
  Use:   "reset",
  Short: "Pool Server Reset for Testing",
  Long:  `Resets dummy users and pools`,
  Run: func(cmd *cobra.Command, args []string) {
      ctx, err := NewContext(url)
      if err != nil {
        return
      }

      RunTestReset(ctx)
  },
}