package server

import (
	"context"
	"fmt"
	"github.com/phuchnd/simple-go-boilerplate/internal/config"
	"github.com/phuchnd/simple-go-boilerplate/internal/logging"
	"github.com/phuchnd/simple-go-boilerplate/server/grpc"
	"github.com/phuchnd/simple-go-boilerplate/server/http"
	"github.com/spf13/cobra"
	nethttp "net/http"
	"os/signal"
	"syscall"
	"time"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run the server.",
	Run: func(cmd *cobra.Command, args []string) {
		serverType, _ := cmd.Flags().GetString("type")
		cfg := config.GetServerConfig()
		logger := logging.NewLogger(cfg)
		ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer stop()

		if serverType == "http" {
			serviceServer := http.NewHTTPServer()
			go func() {
				if err := serviceServer.Start(); err != nil && err != nethttp.ErrServerClosed {
					logger.Fatalf("listen: %s\n", err)
				}
			}()

			logger.Infof("service start listen on port %s", cfg.HTTPPort)

			<-ctx.Done()

			stop()
			logger.Info("shutting down gracefully, press Ctrl+C again to force")

			_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := serviceServer.Stop(); err != nil {
				logger.Fatalf("Server forced to shutdown: ", err)
			}

			logger.Info("Server exiting")
		} else {
			serviceServer, err := grpc.NewGRPCServer()
			if err != nil {
				panic(fmt.Errorf("server initialization failed: %w \n", err))
			}

			if err := serviceServer.Start(); err != nil {
				panic(fmt.Errorf("server starting failed: %w \n", err))
			}
		}
	},
}
