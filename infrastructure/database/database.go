package database

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

func GenerateDsn(user string, password string, url string) string {
	config := mysql.Config{
		User:   user,
		Passwd: password,
		Net:    "tcp",
		Addr:   url,
		DBName: "karaoke",
	}
	return config.FormatDSN()
}

func GetDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open(
		"mysql",
		dsn,
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
