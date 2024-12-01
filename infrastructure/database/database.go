package database

import (
	"database/sql"

	"github.com/ryo-kagawa/LINE-Webhook-Karaoke/domain/repository"
)

type Database struct {
	db *sql.DB
}

var _ = (repository.KaraokeSongRepository)(Database{})
