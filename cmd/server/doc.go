package server

import (
	"fmt"
	"github.com/phuchnd/simple-go-boilerplate/internal/config"
	"github.com/phuchnd/simple-go-boilerplate/server/http/doc"
	"github.com/spf13/cobra"
)

var docCmd = &cobra.Command{
	Use:   "doc",
	Short: "run the document server.",
	Run: func(cmd *cobra.Command, args []string) {
		cfgProvider := config.NewConfig()
		server := doc.NewDocumentServer(cfgProvider)

		if err := server.Start(); err != nil {
			panic(fmt.Errorf("server failed to start: %w", err))
		}
	},
}
