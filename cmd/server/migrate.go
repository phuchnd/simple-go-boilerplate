/*
 * Copyright (C) 2023 by Enterprise Technology, Viet Thai International
 * All Rights Reserved.
 *
 * This source code is protected under international copyright law.  All rights
 * reserved and protected by the copyright holders.
 * This file is confidential and only available to authorized individuals with the
 * permission of the copyright holders.  If you encounter this file and do not have
 * permission, please contact the copyright holders and delete this file.
 */

package server

import (
	"fmt"
	"github.com/phuchnd/core-product-management/internal/db/migrations"
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
