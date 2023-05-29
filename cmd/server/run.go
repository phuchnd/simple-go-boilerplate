package server

import (
	"fmt"
	"github.com/phuchnd/core-product-management/server"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run the server.",
	Run: func(cmd *cobra.Command, args []string) {
		serviceServer, err := server.NewServer()
		if err != nil {
			panic(fmt.Errorf("server initialization failed: %w \n", err))
		}

		if err := serviceServer.Start(); err != nil {
			panic(fmt.Errorf("server starting failed: %w \n", err))
		}
	},
}
