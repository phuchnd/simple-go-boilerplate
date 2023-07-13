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
		cfgProvider := config.NewConfig()
		logger := logging.NewLogger(cfgProvider)

		if serverType == "http" {
			serviceServer, err := http.NewServer(logger, cfgProvider)
			if err != nil {
				logger.Fatalf("HTTP server initialization failed: %s", err.Error())
			}
			if err := serviceServer.Start(); err != nil && err != nethttp.ErrServerClosed {
				logger.Fatalf("HTTP listen start failed %s", err)
			}
		} else {
			serviceServer, err := grpc.NewServer(logger, cfgProvider)
			if err != nil {
				logger.Fatalf("GRPC server initialization failed: %s", err.Error())
			}

			if err := serviceServer.Start(); err != nil {
				logger.Fatalf("GRPC server starting failed: %s", err.Error())
			}
		}
	},
}
