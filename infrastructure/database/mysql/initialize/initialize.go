package initialize

import (
	"database/sql"
)

func InitializeDB(db *sql.DB, database string) error {
	if err := InitializeDatabase(db, database); err != nil {
		return err
	}
	if err := InitializeTable(db); err != nil {
		return err
	}
	if err := InitializeData(db); err != nil {
		return err
	}

	return nil
}
