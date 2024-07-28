package initialize

import (
	"github.com/ryo-kagawa/LINE-Webhook-Karaoke/infrastructure/database"
)

func Initialize(dsn string) error {
	// DSNを指定して接続したとしてもその後にDATABASEを作成した場合にはDATABASEの指定は解除されるのでここだけで使用する
	{
		db, err := database.GetDatabase(dsn)
		if err != nil {
			return err
		}
		defer db.Close()
		err = InitializeDatabase(db)
		if err != nil {
			return err
		}
	}

	db, err := database.GetDatabase(dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	err = InitializeTable(db)
	if err != nil {
		return err
	}

	err = InitializeData(db)
	if err != nil {
		return err
	}

	return nil
}
