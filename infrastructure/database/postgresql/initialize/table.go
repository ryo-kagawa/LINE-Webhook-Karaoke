package initialize

import (
	"database/sql"

	"github.com/ryo-kagawa/LINE-Webhook-Karaoke/infrastructure/database/postgresql/table"
)

func InitializeTable(db *sql.DB) error {
	err := table.CreateTableArtist(db)
	if err != nil {
		return err
	}
	err = table.CreateTableSong(db)
	if err != nil {
		return err
	}

	return nil
}
