package table

import "database/sql"

type Artist struct {
	Id         int64
	Name       string
	DamId      sql.Null[string]
	JoysoundId sql.Null[string]
}

func CreateTableArtist(db *sql.DB) error {
	_, err := db.Exec(
		`
CREATE TABLE artist(
id         BIGINT      NOT NULL AUTO_INCREMENT,
name       VARCHAR(50) NOT NULL,
damId      VARCHAR(6)      NULL                UNIQUE,
joysoundId VARCHAR(6)      NULL                UNIQUE,

PRIMARY KEY(id)
)
`,
	)

	return err
}

func (a Artist) Insert(db *sql.DB) error {
	_, err := db.Exec(
		`
INSERT INTO artist
(
	name,
	damId,
	joysoundId
)VALUES(
	?,
	?,
	?
)
`,
		a.Name,
		a.DamId,
		a.JoysoundId,
	)
	if err != nil {
		return err
	}

	return nil
}

func (a Artist) GetIdFromName(db *sql.DB) (int64, error) {
	id := int64(0)
	err := db.QueryRow(`SELECT id FROM artist WHERE name=?`, a.Name).Scan(&id)
	return id, err
}
