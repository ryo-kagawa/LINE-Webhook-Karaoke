package database

import (
	"database/sql"

	"github.com/ryo-kagawa/LINE-Webhook-Karaoke/domain/repository"
)

type Database struct {
	DB *sql.DB
}

var _ = (repository.KaraokeSongRepository)(Database{})
