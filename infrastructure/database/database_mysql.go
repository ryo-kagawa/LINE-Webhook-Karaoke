//go:build mysql

package database

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"

	"github.com/ryo-kagawa/LINE-Webhook-Karaoke/domain/model"
	"github.com/ryo-kagawa/LINE-Webhook-Karaoke/infrastructure/database/mysql/initialize"
	"github.com/ryo-kagawa/LINE-Webhook-Karaoke/infrastructure/database/mysql/table"
)

func NewDatabase(address string, user string, password string, database string) (Database, error) {
	config := mysql.Config{
		User:   user,
		Passwd: password,
		Net:    "tcp",
		Addr:   address,
		DBName: database,
	}
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		return Database{}, err
	}

	return Database{
		DB: db,
	}, nil
}

func (d Database) Initialize(database string) error {
	return initialize.InitializeDB(d.DB, database)
}

func (d Database) Dam() ([]model.KaraokeSong, error) {
	rows, err := d.DB.Query(
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
			LEFT(lyrics, 50) AS lyrics,
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
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return karaokeSongs, nil
}

func (d Database) Joysound() ([]model.KaraokeSong, error) {
	rows, err := d.DB.Query(
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
			LEFT(lyrics, 50) AS lyrics,
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
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return karaokeSongs, nil
}

func (d Database) Ramdom() ([]model.KaraokeSong, error) {
	rows, err := d.DB.Query(
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
			LEFT(lyrics, 50) AS lyrics,
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
