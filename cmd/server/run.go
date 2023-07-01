package server

import (
	"github.com/phuchnd/simple-go-boilerplate/internal/config"
	"github.com/phuchnd/simple-go-boilerplate/internal/logging"
	"github.com/phuchnd/simple-go-boilerplate/server/grpc"
	"github.com/phuchnd/simple-go-boilerplate/server/http"
	"github.com/spf13/cobra"
	nethttp "net/http"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run the server.",
	Run: func(cmd *cobra.Command, args []string) {
		serverType, _ := cmd.Flags().GetString("type")
		cfg := config.GetServerConfig()
		logger := logging.NewLogger(cfg)

		if serverType == "http" {
			serviceServer, err := http.NewServer(logger)
			if err != nil {
				logger.Fatalf("HTTP server initialization failed: %w \n", err)
			}
			if err := serviceServer.Start(); err != nil && err != nethttp.ErrServerClosed {
				logger.Fatalf("HTTP listen start failed %s", err)
			}
		} else {
			serviceServer, err := grpc.NewServer(logger)
			if err != nil {
				logger.Fatalf("GRPC server initialization failed: %w \n", err)
			}

			if err := serviceServer.Start(); err != nil {
				logger.Fatalf("GRPC server starting failed: %w \n", err)
			}
		}
	},
}
