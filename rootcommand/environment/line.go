package environment

import (
	"errors"
	"os"
)

type EnvironmentLine struct {
	LINE_CHANNEL_SECRET string
	LINE_CHANNEL_TOKEN  string
}

func GetEnvironmentLine() EnvironmentLine {
	return EnvironmentLine{
		LINE_CHANNEL_SECRET: os.Getenv("LINE_CHANNEL_SECRET"),
		LINE_CHANNEL_TOKEN:  os.Getenv("LINE_CHANNEL_TOKEN"),
	}
}

func (env EnvironmentLine) Validate() error {
	if env.LINE_CHANNEL_SECRET == "" {
		return errors.New("環境変数LINE_CHANNEL_SECRETが未設定です")
	}
	if env.LINE_CHANNEL_TOKEN == "" {
		return errors.New("環境変数LINE_CHANNEL_SECRETが未設定です")
	}

	return nil
}
