package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "watch",
		Short: "Entry command for Watchlist CLI",
		Long:  "Watchlist CLI is a tool to have your own Youtube (or any other video service) watchlist via CLI",

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(os.Stdout, "Root Command, use 'watch --help' to see more")
		},
	}

	testCmd = &cobra.Command{
		Use:   "testcmd ...strings",
		Short: "prints the given string to test the tool functionality",
		Args:  cobra.ArbitraryArgs,
		Run: func(cmd *cobra.Command, args []string) {
			output := strings.Join(args, " ")
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

	rootCmd.Version = "0.0.1"
	rootCmd.SetVersionTemplate(fmt.Sprintf("Watchlist CLI version %v", rootCmd.Version))

	rootCmd.AddCommand(testCmd)
	rootCmd.AddCommand(getCmd)
}
