package database

import (
	"database/sql"

	"github.com/ryo-kagawa/LINE-Webhook-Karaoke/domain/model"
	"github.com/ryo-kagawa/LINE-Webhook-Karaoke/domain/repository"
	"github.com/ryo-kagawa/LINE-Webhook-Karaoke/infrastructure/database/table"
)

type karaokeSongDatabase struct {
	db *sql.DB
}

func NewKaraokeSongDatabase(db *sql.DB) repository.KaraokeSongRepository {
	return karaokeSongDatabase{db}
}

func (k karaokeSongDatabase) Dam() ([]model.KaraokeSong, error) {
	rows, err := k.db.Query(
		`
SELECT
	artist.name AS artistName,
	song.name AS songName,
	song.lyrics AS lyrics,
	song.damId AS damId
FROM
	(
		SELECT
			artistId,
			name,
			lyrics,
			damId
		FROM
			song
		WHERE
			damId IS NOT NULL
		ORDER BY
			RAND()
		LIMIT 5
	) AS song
	JOIN artist ON(
		artist.id = song.artistId
	)
ORDER BY
	RAND()
`,
	)
	if err != nil {
		return nil, err
	}
	karaokeSongs := make([]model.KaraokeSong, 0, 5)
	for rows.Next() {
		res := &model.KaraokeSong{}
		err = rows.Scan(
			&res.ArtistName,
			&res.SongName,
			&res.Lyrics,
			&res.DamId,
		)
		if err != nil {
			return nil, err
		}
		karaokeSongs = append(karaokeSongs, *res)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return karaokeSongs, nil
}

func (k karaokeSongDatabase) Joysound() ([]model.KaraokeSong, error) {
	rows, err := k.db.Query(
		`
SELECT
	artist.name AS artistName,
	song.name AS songName,
	song.lyrics AS lyrics,
	song.joysoundId AS joysoundId
FROM
	(
		SELECT
			artistId,
			name,
			lyrics,
			joysoundId
		FROM
			song
		WHERE
			joysoundId IS NOT NULL
		ORDER BY
			RAND()
		LIMIT 5
	) AS song
	JOIN artist ON(
		artist.id = song.artistId
	)
ORDER BY
	RAND()
`,
	)
	if err != nil {
		return nil, err
	}
	karaokeSongs := make([]model.KaraokeSong, 0, 5)
	for rows.Next() {
		res := &model.KaraokeSong{}
		err = rows.Scan(
			&res.ArtistName,
			&res.SongName,
			&res.Lyrics,
			&res.JoysoundId,
		)
		if err != nil {
			return nil, err
		}
		karaokeSongs = append(karaokeSongs, *res)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return karaokeSongs, nil
}

func (k karaokeSongDatabase) Ramdom() ([]model.KaraokeSong, error) {
	rows, err := k.db.Query(
		`
SELECT
	artist.name AS artistName,
	song.name AS songName,
	song.lyrics AS lyrics,
	song.damId AS damId,
	song.joysoundId AS joysoundId
FROM
	(
		SELECT
			artistId,
			name,
			lyrics,
			damId,
			joysoundId
		FROM
			song
		ORDER BY
			RAND()
		LIMIT 5
	) AS song
	JOIN artist ON(
		artist.id = song.artistId
	)
ORDER BY
	RAND()
`,
	)
	if err != nil {
		return nil, err
	}
	karaokeSongs := make([]model.KaraokeSong, 0, 5)
	for rows.Next() {
		type KaraokeSong struct {
			Artist table.Artist
			Song   table.Song
		}

		res := &KaraokeSong{}
		err = rows.Scan(
			&res.Artist.Name,
			&res.Song.Name,
			&res.Song.Lyrics,
			&res.Song.DamId,
			&res.Song.JoysoundId,
		)
		if err != nil {
			return nil, err
		}
		karaokeSongs = append(
			karaokeSongs,
			model.KaraokeSong{
				ArtistName: res.Artist.Name,
				SongName:   res.Song.Name,
				Lyrics:     res.Song.Lyrics,
				DamId:      res.Song.DamId.V,
				JoysoundId: res.Song.JoysoundId.V,
			},
		)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return karaokeSongs, nil
}
