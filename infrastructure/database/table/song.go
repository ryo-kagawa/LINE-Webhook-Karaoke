package table

import "database/sql"

type Song struct {
	Id         int64
	ArtistId   int64
	Name       string
	Lyrics     string
	DamId      sql.Null[string]
	JoysoundId sql.Null[string]
}

func CreateTableSong(db *sql.DB) error {
	_, err := db.Exec(
		`
CREATE TABLE song(
id         BIGINT        NOT NULL AUTO_INCREMENT,
artistId   BIGINT        NOT NULL,
name       VARCHAR(50)   NOT NULL,
lyrics     VARCHAR(1000) NOT NULL,
damId      VARCHAR(7)        NULL                UNIQUE,
joysoundId VARCHAR(6)        NULL                UNIQUE,


PRIMARY KEY(id),
FOREIGN KEY (artistId) REFERENCES artist(id)
)
`,
	)

	return err
}

func (s Song) Insert(db *sql.DB) error {
	_, err := db.Exec(
		`
INSERT INTO song
(
	artistId,
	name,
	lyrics,
	damId,
	joysoundId
)VALUES(
	?,
	?,
	?,
	?,
	?
)
`,
		s.ArtistId,
		s.Name,
		s.Lyrics,
		s.DamId,
		s.JoysoundId,
	)
	if err != nil {
		return err
	}

	return nil
}
