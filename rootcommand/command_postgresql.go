//go:build postgresql

package rootcommand

import (
	"strconv"

	"github.com/ryo-kagawa/LINE-Webhook-Karaoke/infrastructure/database"
	"github.com/ryo-kagawa/LINE-Webhook-Karaoke/rootcommand/environment"
)

func NewDatabase(environment environment.Environment) (database.Database, error) {
	port, err := strconv.Atoi(environment.Database.DATABASE_POSTGRESQL_PORT)
	if err != nil {
		return database.Database{}, err
	}
	return database.NewDatabase(
		environment.Database.DATABASE_POSTGRESQL_HOST,
		uint16(port),
		environment.Database.DATABASE_POSTGRESQL_USER,
		environment.Database.DATABASE_POSTGRESQL_PASSWORD,
		environment.Database.DATABASE_POSTGRESQL_DATABASE,
	)
}
