package server

import (
	"fmt"
	"github.com/phuchnd/simple-go-boilerplate/internal/db/migrations"
	"log"

	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate database.",
	Run: func(cmd *cobra.Command, args []string) {
		migrator, err := migrations.NewMigrator()
		if err != nil {
			panic(fmt.Errorf("failed to create db migrator: %w \n", err))
		}

		direction, _ := cmd.Flags().GetString("direction")

		switch direction {
		case "up":
			_, _ = migrator.Up()
		case "down":
			_, _ = migrator.Down()
		default:
			log.Fatalln("unsupported migration direction")
		}
	},
}
