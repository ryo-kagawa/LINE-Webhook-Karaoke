//go:build postgresql

package initialize

import (
	"strconv"

	"github.com/ryo-kagawa/LINE-Webhook-Karaoke/infrastructure/database"
	"github.com/ryo-kagawa/LINE-Webhook-Karaoke/subcommand/initialize/environment"
)

func NewDatabase(environment environment.Environment) (database.Database, error) {
	port, err := strconv.Atoi(environment.DATABASE_POSTGRESQL_PORT)
	if err != nil {
		return database.Database{}, err
	}
	return database.NewDatabase(
		environment.DATABASE_POSTGRESQL_HOST,
		uint16(port),
		environment.DATABASE_POSTGRESQL_USER,
		environment.DATABASE_POSTGRESQL_PASSWORD,
		environment.DATABASE_POSTGRESQL_DATABASE,
		environment.DATABASE_POSTGRESQL_SSLMODE,
		"",
	)
}

func DatabaseInitialize(db database.Database, environment environment.Environment) error {
	if environment.DATABASE_INITIALIZE_DATABASE == "true" {
		if err := db.InitializeDatabase(environment.DATABASE_POSTGRESQL_DATABASE); err != nil {
			return err
		}
	}
	if err := db.InitializeSchema(environment.DATABASE_POSTGRESQL_SCHEMA); err != nil {
		return err
	}
	port, _ := strconv.Atoi(environment.DATABASE_POSTGRESQL_PORT)
	db, err := database.NewDatabase(
		environment.DATABASE_POSTGRESQL_HOST,
		uint16(port),
		environment.DATABASE_POSTGRESQL_USER,
		environment.DATABASE_POSTGRESQL_PASSWORD,
		environment.DATABASE_POSTGRESQL_DATABASE,
		environment.DATABASE_POSTGRESQL_SSLMODE,
		environment.DATABASE_POSTGRESQL_SCHEMA,
	)
	if err != nil {
		return err
	}
	if err := db.InitializeTable(); err != nil {
		return err
	}

	return nil
}
