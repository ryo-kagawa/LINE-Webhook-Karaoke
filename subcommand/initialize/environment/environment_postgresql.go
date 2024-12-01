//go:build postgresql

package environment

import (
	"errors"
	"os"
)

type Environment struct {
	DATABASE_POSTGRESQL_HOST     string
	DATABASE_POSTGRESQL_PORT     string
	DATABASE_POSTGRESQL_USER     string
	DATABASE_POSTGRESQL_PASSWORD string
	DATABASE_POSTGRESQL_DATABASE string
}

func GetEnvironment() Environment {
	return Environment{
		DATABASE_POSTGRESQL_HOST:     os.Getenv("DATABASE_POSTGRESQL_HOST"),
		DATABASE_POSTGRESQL_PORT:     os.Getenv("DATABASE_POSTGRESQL_PORT"),
		DATABASE_POSTGRESQL_USER:     os.Getenv("DATABASE_POSTGRESQL_USER"),
		DATABASE_POSTGRESQL_PASSWORD: os.Getenv("DATABASE_POSTGRESQL_PASSWORD"),
		DATABASE_POSTGRESQL_DATABASE: os.Getenv("DATABASE_POSTGRESQL_DATABASE"),
	}
}

func (e Environment) Validate() error {
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

	return nil
}
