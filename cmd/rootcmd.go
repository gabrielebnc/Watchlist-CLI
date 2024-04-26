package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	softwareVersion = "0.1"
)

var (
	rootCmd = &cobra.Command{
		Use:   "watch",
		Short: "Entry command for Watchlist CLI",
		Long:  "Watchlist CLI is a tool to have your own Youtube (or any other video service) watchlist via CLI",

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(os.Stdout, "Roooooot")
		},
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Prints the version of Watchlist CLI",
		Run: func(cmd *cobra.Command, args []string) {
			output := fmt.Sprintf("Watchlist CLI: v%v", softwareVersion)
			fmt.Fprintln(os.Stdout, output)
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
