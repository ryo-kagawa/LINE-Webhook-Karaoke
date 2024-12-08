//go:build postgresql

package environment

import (
	"errors"
	"os"
)

type EnvironmentDatabase struct {
	DATABASE_POSTGRESQL_HOST     string
	DATABASE_POSTGRESQL_PORT     string
	DATABASE_POSTGRESQL_USER     string
	DATABASE_POSTGRESQL_PASSWORD string
	DATABASE_POSTGRESQL_DATABASE string
	DATABASE_POSTGRESQL_SSLMODE  string
	DATABASE_POSTGRESQL_SCHEMA   string
}

func GetEnvironmentDatabase() EnvironmentDatabase {
	return EnvironmentDatabase{
		DATABASE_POSTGRESQL_HOST:     os.Getenv("DATABASE_POSTGRESQL_HOST"),
		DATABASE_POSTGRESQL_PORT:     os.Getenv("DATABASE_POSTGRESQL_PORT"),
		DATABASE_POSTGRESQL_USER:     os.Getenv("DATABASE_POSTGRESQL_USER"),
		DATABASE_POSTGRESQL_PASSWORD: os.Getenv("DATABASE_POSTGRESQL_PASSWORD"),
		DATABASE_POSTGRESQL_DATABASE: os.Getenv("DATABASE_POSTGRESQL_DATABASE"),
		DATABASE_POSTGRESQL_SSLMODE:  os.Getenv("DATABASE_POSTGRESQL_SSLMODE"),
		DATABASE_POSTGRESQL_SCHEMA:   os.Getenv("DATABASE_POSTGRESQL_SCHEMA"),
	}
}

func (e EnvironmentDatabase) Validate() error {
	if e.DATABASE_POSTGRESQL_HOST == "" {
		return errors.New("環境変数DATABASE_POSTGRESQL_HOSTが未設定です")
	}
	if e.DATABASE_POSTGRESQL_PORT == "" {
		return errors.New("環境変数DATABASE_POSTGRESQL_PORTが未設定です")
	}
	if e.DATABASE_POSTGRESQL_USER == "" {
		return errors.New("環境変数DATABASE_POSTGRESQL_USERが未設定です")
	}
	if e.DATABASE_POSTGRESQL_PASSWORD == "" {
		return errors.New("環境変数DATABASE_POSTGRESQL_PASSWORDが未設定です")
	}
	if e.DATABASE_POSTGRESQL_DATABASE == "" {
		return errors.New("環境変数DATABASE_POSTGRESQL_DATABASEが未設定です")
	}
	if e.DATABASE_POSTGRESQL_SSLMODE == "" {
		return errors.New("環境変数DATABASE_POSTGRESQL_SSLMODEが未設定です")
	}
	if e.DATABASE_POSTGRESQL_SCHEMA == "" {
		return errors.New("環境変数DATABASE_POSTGRESQL_SCHEMAが未設定です")
	}

	return nil
}
