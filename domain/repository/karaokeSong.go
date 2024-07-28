package repository

import "github.com/ryo-kagawa/LINE-Webhook-Karaoke/domain/model"

type KaraokeSongRepository interface {
	Dam() ([]model.KaraokeSong, error)
	Joysound() ([]model.KaraokeSong, error)
	Ramdom() ([]model.KaraokeSong, error)
}
