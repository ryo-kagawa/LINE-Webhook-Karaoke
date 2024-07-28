package initialize

import (
	"database/sql"
)

func InitializeDatabase(db *sql.DB) error {
	_, err := db.Exec("DROP DATABASE karaoke")
	if err != nil {
		return err
	}
	_, err = db.Exec("CREATE DATABASE karaoke")
	if err != nil {
		return err
	}

	return nil
}
