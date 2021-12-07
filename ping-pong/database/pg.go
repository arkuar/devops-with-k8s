package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

var pgDb *pg.DB
var ctx = context.Background()
var isReady = false

func InitDb() {
	pgDb := pg.Connect(&pg.Options{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
	})

	for {
		err := pgDb.Ping(ctx)
		if err != nil {
			fmt.Println("Error connecting to database, retrying")
			time.Sleep(5 * time.Second)
			continue
		} else {
			err := createSchema(pgDb)
			if err != nil {
				log.Fatal(err)
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
			isReady = true
			return
		}
	}
}

func createSchema(db *pg.DB) error {
	err := db.Model((*Ping)(nil)).CreateTable(&orm.CreateTableOptions{
		Temp:        false,
		IfNotExists: true,
	})
	return err
}

func GetDB() (*pg.DB, bool) {
	return pgDb, isReady
}
