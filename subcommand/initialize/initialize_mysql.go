//go:build mysql

package initialize

import (
	"github.com/ryo-kagawa/LINE-Webhook-Karaoke/infrastructure/database"
	"github.com/ryo-kagawa/LINE-Webhook-Karaoke/subcommand/initialize/environment"
)

func NewDatabase(environment environment.Environment) (database.Database, error) {
	return database.NewDatabase(
		environment.DATABASE_MYSQL_ADDRESS,
		environment.DATABASE_MYSQL_USER,
		environment.DATABASE_MYSQL_PASSWORD,
		"",
	)
}

func DatabaseInitialize(db database.Database, environment environment.Environment) error {
	return db.Initialize(environment.DATABASE_MYSQL_DATABASE)
}
