//go:build mysql

package rootcommand

import (
	"github.com/ryo-kagawa/LINE-Webhook-Karaoke/infrastructure/database"
	"github.com/ryo-kagawa/LINE-Webhook-Karaoke/rootcommand/environment"
)

func NewDatabase(environment environment.Environment) (database.Database, error) {
	return database.NewDatabase(
		environment.Database.DATABASE_MYSQL_ADDRESS,
		environment.Database.DATABASE_MYSQL_USER,
		environment.Database.DATABASE_MYSQL_PASSWORD,
		environment.Database.DATABASE_MYSQL_DATABASE,
	)
}
