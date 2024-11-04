package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCommand = &cobra.Command{
	Use:   "dbsync",
	Short: "A CLI utility to manage database backups and restores",
	Long:  "A Command Line Interface Utility for managing database backups and restores",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Your args are %s", args)
	},
}

func Execute() {
	rootCommand.AddCommand(backupCommand)
	if err := rootCommand.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
