package db

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

var DB *pg.DB

func ConnectDB(port string, username string, pword string, dbName string) {
	DB = pg.Connect(&pg.Options{
		Addr:     port,
		User:     username,
		Password: pword,
		Database: dbName,
	})

	models := []interface{}{
		(*User)(nil),
		(*Transaction)(nil),
	}

	for _, model := range models {
		DB.Model(model).CreateTable(&orm.CreateTableOptions{})
	}
	// defer db.Close()
}
