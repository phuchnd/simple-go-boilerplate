package server

import (
	"fmt"
	"github.com/phuchnd/simple-go-boilerplate/internal/config"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "app",
	Version: "0.0.0",
	Short:   "Order service application.",
}

func init() {
	rootCmd.AddCommand(docCmd)
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(migrateCmd)

	runCmd.Flags().StringP("type", "t", "http", "Server Type.")
	migrateCmd.Flags().StringP("direction", "d", "up", "Migration direction.")
}

func Execute() {
	cfgProvider := config.NewConfig()
	rootCmd.Short = cfgProvider.GetServerConfig().Name
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
