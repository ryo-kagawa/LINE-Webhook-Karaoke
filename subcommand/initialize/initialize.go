package initialize

import (
	"github.com/ryo-kagawa/LINE-Webhook-Karaoke/subcommand/initialize/environment"
)

type Initialize struct {
}

func (Initialize) Execute(arguments []string) (string, error) {
	environment := environment.GetEnvironment()
	err := environment.Validate()
	if err != nil {
		return "", err
	}

	db, err := NewDatabase(environment)
	if err != nil {
		return "", err
	}

	if err := DatabaseInitialize(db, environment); err != nil {
		return "", err
	}

	return "initialize finish", nil
}

func (Initialize) Name() string {
	return "initialize"
}
