package main

import (
	"errors"
	"os"
)

type Environment struct {
	DATABASE_PASSWORD   string
	DATABASE_URL        string
	DATABASE_USER       string
	LINE_CHANNEL_SECRET string
	LINE_CHANNEL_TOKEN  string
}

func GetEnvironment() Environment {
	return Environment{
		DATABASE_PASSWORD:   os.Getenv("DATABASE_PASSWORD"),
		DATABASE_URL:        os.Getenv("DATABASE_URL"),
		DATABASE_USER:       os.Getenv("DATABASE_USER"),
		LINE_CHANNEL_SECRET: os.Getenv("LINE_CHANNEL_SECRET"),
		LINE_CHANNEL_TOKEN:  os.Getenv("LINE_CHANNEL_TOKEN"),
	}
}

func (env Environment) Validate() error {
	if env.DATABASE_PASSWORD == "" {
		return errors.New("環境変数DATABASE_PASSWORDが未設定です")
	}
	if env.DATABASE_URL == "" {
		return errors.New("環境変数DATABASE_URLが未設定です")
	}
	if env.DATABASE_USER == "" {
		return errors.New("環境変数DATABASE_USERが未設定です")
	}
	if env.LINE_CHANNEL_SECRET == "" {
		return errors.New("環境変数LINE_CHANNEL_SECRETが未設定です")
	}
	if env.LINE_CHANNEL_TOKEN == "" {
		return errors.New("環境変数LINE_CHANNEL_SECRETが未設定です")
	}

	return nil
}
