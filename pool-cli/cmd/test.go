package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(testCmd)

  testCmd.AddCommand(testSetupCmd)
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
    fmt.Println("Running test setup...")
  },
}