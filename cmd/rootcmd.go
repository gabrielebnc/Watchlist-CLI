package cmd

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var persistence_path string

var rootCmd = &cobra.Command{
	Use:   "watch",
	Short: "Entry command for Watchlist CLI",
	Long:  "Watchlist CLI is a tool to have your own Youtube (or any other video service) watchlist via CLI",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintln(os.Stdout, "Root Command, use 'watch --help' to see more")
		fmt.Fprintln(os.Stdout, "\nConfig file used: ", viper.ConfigFileUsed())
	},
}

var testCmd = &cobra.Command{
	Use:   "testcmd ...strings",
	Short: "prints the given string to test the tool functionality",
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		output := strings.Join(args, " ")
		fmt.Fprintln(os.Stdout, output)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		cobra.CheckErr(err)
	}
}

func InitFunc() {

	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".watchcli.yaml")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stdout, "Config file not found")
		cobra.CheckErr(err)
	}

	configsBytes, err := os.ReadFile(viper.ConfigFileUsed())
	cobra.CheckErr(err)
	cobra.CheckErr(viper.ReadConfig(bytes.NewBuffer(configsBytes)))

	rootCmd.PersistentFlags().StringVar(&persistence_path, "persistencePath", viper.GetString("watchcli.configs.persistencePath"), "Path of the persistence file")
}

func init() {

	rootCmd.Version = "0.0.1"
	rootCmd.SetVersionTemplate(fmt.Sprintf("Watchlist CLI version %v", rootCmd.Version))

	InitFunc()

	rootCmd.AddCommand(testCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(addCmd)
}
