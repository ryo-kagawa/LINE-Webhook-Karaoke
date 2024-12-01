//go:build mysql

package environment

import (
	"errors"
	"os"
)

type EnvironmentDatabase struct {
	DATABASE_MYSQL_ADDRESS  string
	DATABASE_MYSQL_USER     string
	DATABASE_MYSQL_PASSWORD string
	DATABASE_MYSQL_DATABASE string
}

func GetEnvironmentDatabase() EnvironmentDatabase {
	return EnvironmentDatabase{
		DATABASE_MYSQL_ADDRESS:  os.Getenv("DATABASE_MYSQL_ADDRESS"),
		DATABASE_MYSQL_DATABASE: os.Getenv("DATABASE_MYSQL_DATABASE"),
		DATABASE_MYSQL_PASSWORD: os.Getenv("DATABASE_MYSQL_PASSWORD"),
		DATABASE_MYSQL_USER:     os.Getenv("DATABASE_MYSQL_USER"),
	}
}

func (e EnvironmentDatabase) Validate() error {
	if e.DATABASE_MYSQL_ADDRESS == "" {
		return errors.New("環境変数DATABASE_MYSQL_ADDRESSが未設定です")
	}
	if e.DATABASE_MYSQL_DATABASE == "" {
		return errors.New("環境変数DATABASE_MYSQL_DATABASEが未設定です")
	}
	if e.DATABASE_MYSQL_PASSWORD == "" {
		return errors.New("環境変数DATABASE_MYSQL_PASSWORDが未設定です")
	}
	if e.DATABASE_MYSQL_USER == "" {
		return errors.New("環境変数DATABASE_MYSQL_USERが未設定です")
	}

	return nil
}
