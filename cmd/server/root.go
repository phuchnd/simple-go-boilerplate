package server

import (
	"fmt"
	"github.com/phuchnd/core-product-management/internal/config"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "app",
	Version: "0.0.0",
	Short:   "Order service application.",
}

func init() {
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(migrateCmd)

	migrateCmd.Flags().StringP("direction", "d", "up", "Migration direction.")
}

func Execute() {
	cfg := config.GetServerConfig()
	rootCmd.Short = cfg.Name
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
