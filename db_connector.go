package main

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

var db *pg.DB

func connectDB(port string, username string, pword string, dbName string) {
	db = pg.Connect(&pg.Options{
		Addr:     port,
		User:     username,
		Password: pword,
		Database: dbName,
	})

	models := []interface{}{
		(*user)(nil),
		(*transaction)(nil),
	}

	for _, model := range models {
		db.Model(model).CreateTable(&orm.CreateTableOptions{})
	}
	// defer db.Close()
}
