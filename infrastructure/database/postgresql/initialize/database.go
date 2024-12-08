package initialize

import (
	"database/sql"
)

func InitializeDatabase(db *sql.DB, database string) error {
	_, err := db.Exec("DROP DATABASE IF EXISTS " + database)
	if err != nil {
		return err
	}
	_, err = db.Exec("CREATE DATABASE " + database)
	if err != nil {
		return err
	}

	return nil
}

func InitializeSchema(db *sql.DB, schema string) error {
	_, err := db.Exec("DROP SCHEMA IF EXISTS " + schema)
	if err != nil {
		return err
	}
	_, err = db.Exec("CREATE SCHEMA " + schema)
	if err != nil {
		return err
	}

	return nil
}
