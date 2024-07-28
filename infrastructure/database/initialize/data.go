package initialize

import (
	"database/sql"

	"github.com/ryo-kagawa/LINE-Webhook-Karaoke/infrastructure/database/table"
	"github.com/ryo-kagawa/LINE-Webhook-Karaoke/infrastructure/json/model"
)

func createSqlNull[T any](value *T) sql.Null[T] {
	valid := value != nil
	v := *new(T)
	if valid {
		v = *value
	}
	return sql.Null[T]{
		V:     v,
		Valid: valid,
	}
}

func InitializeData(db *sql.DB) error {
	data, err := model.Load()
	if err != nil {
		return err
	}
	for _, jsonArtist := range data {
		artist := table.Artist{
			Name:       jsonArtist.Name,
			DamId:      createSqlNull(jsonArtist.DamId),
			JoysoundId: createSqlNull(jsonArtist.JoysoundId),
		}
		err = artist.Insert(db)
		if err != nil {
			return err
		}
		artistId, err := artist.GetIdFromName(db)
		if err != nil {
			return err
		}
		for _, jsonSong := range jsonArtist.Songs {
			song := table.Song{
				ArtistId:   artistId,
				Name:       jsonSong.Name,
				Lyrics:     jsonSong.Lyrics,
				DamId:      createSqlNull(jsonSong.DamId),
				JoysoundId: createSqlNull(jsonSong.JoysoundId),
			}
			err = song.Insert(db)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
