package server

import (
	"fmt"
	"github.com/phuchnd/simple-go-boilerplate/internal/config"
	"github.com/phuchnd/simple-go-boilerplate/internal/db/migrations"
	"github.com/phuchnd/simple-go-boilerplate/internal/db/mysql"
	"log"

	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate database.",
	Run: func(cmd *cobra.Command, args []string) {
		cfgProvider := config.NewConfig()
		db, err := mysql.NewDB(cfgProvider.GetDBConfig().MySQL)
		if err != nil {
			panic(fmt.Errorf("failed to create db: %w \n", err))
		}
		migrator, err := migrations.NewMigrator(db)
		if err != nil {
			panic(fmt.Errorf("failed to create db migrator: %w \n", err))
		}

		direction, _ := cmd.Flags().GetString("direction")

		switch direction {
		case "up":
			_, err = migrator.Up()
		case "down":
			_, err = migrator.Down()
		default:
			log.Fatalln("unsupported migration direction")
		}

		if err != nil {
			panic(fmt.Errorf("failed to execute db migrator: %w \n", err))
		}
	},
}
