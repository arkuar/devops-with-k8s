package database

import (
	"fmt"
	"log"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func InitDb() (pgDb *pg.DB) {
	db := pg.Connect(&pg.Options{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
	})

	pgDb = db

	err := createSchema(db)
	if err != nil {
		panic(err)
	}
	pong := &Ping{
		ID:      1,
		Counter: 0,
	}
	_, err = pgDb.Model(pong).OnConflict("DO NOTHING").Insert()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database initialized")

	return
}

func createSchema(db *pg.DB) error {
	err := db.Model((*Ping)(nil)).CreateTable(&orm.CreateTableOptions{
		Temp:        false,
		IfNotExists: true,
	})
	return err
}
