package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "scratchpad",
	Short: "A CLI tool for audio scratchpad operations",
	Long:  `scratchpad is a CLI application for manipulating audio files with automation.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(playCmd)
}
