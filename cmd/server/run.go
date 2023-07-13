package server

import (
	"github.com/phuchnd/simple-go-boilerplate/internal/config"
	mysqldb "github.com/phuchnd/simple-go-boilerplate/internal/db/mysql"
	"github.com/phuchnd/simple-go-boilerplate/internal/logging"
	http2 "github.com/phuchnd/simple-go-boilerplate/internal/service/http"
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

		dbConfig := cfgProvider.GetDBConfig()
		db, err := mysqldb.NewDB(dbConfig.MySQL)
		if err != nil {
			panic("failed to init db")
		}

		if serverType == "http" {
			httpServerHandler, err := http2.NewHTTPService(cfgProvider)
			if err != nil {
				logger.Fatalf("HTTP server initialization halder failed : %s", err.Error())
			}
			serviceServer, err := http.NewServer(logger, cfgProvider, db, httpServerHandler)
			if err != nil {
				logger.Fatalf("HTTP server initialization failed: %s", err.Error())
			}
			if err := serviceServer.Start(); err != nil && err != nethttp.ErrServerClosed {
				logger.Fatalf("HTTP listen start failed %s", err)
			}
		} else {
			serviceServer, err := grpc.NewServer(logger, cfgProvider, db)
			if err != nil {
				logger.Fatalf("GRPC server initialization failed: %s", err.Error())
			}

			if err := serviceServer.Start(); err != nil {
				logger.Fatalf("GRPC server starting failed: %s", err.Error())
			}
		}
	},
}
