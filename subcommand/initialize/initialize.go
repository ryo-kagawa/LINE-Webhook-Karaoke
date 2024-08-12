package initialize

import (
	"github.com/ryo-kagawa/LINE-Webhook-Karaoke/infrastructure/database"
	"github.com/ryo-kagawa/LINE-Webhook-Karaoke/infrastructure/database/initialize"
)

type Initialize struct {
}

func (i Initialize) Name() string {
	return "initialize"
}

func (i Initialize) Execute() (string, error) {
	environment := GetEnvironment()
	err := environment.Validate()
	if err != nil {
		return "", err
	}

	dsn := database.GenerateDsn(
		environment.DATABASE_USER,
		environment.DATABASE_PASSWORD,
		environment.DATABASE_URL,
	)

	err = initialize.Initialize(dsn)
	if err != nil {
		return "", err
	}

	return "initialize finish", nil
}
