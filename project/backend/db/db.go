package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

var pgdb *pg.DB
var ctx = context.Background()
var isReady = false

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func InitDb() {
	pgdb = pg.Connect(&pg.Options{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
	})

	for {
		err := pgdb.Ping(ctx)
		if err != nil {
			fmt.Println("Error connecting to database, retrying")
			time.Sleep(5 * time.Second)
			continue
		} else {
			err := createSchema(pgdb)
			checkErr(err)

			todo1 := &Todo{
				ID:      1,
				Content: "TODO 1",
			}
			_, err = pgdb.Model(todo1).OnConflict("(id) DO NOTHING").Insert()
			checkErr(err)

			log.Println("Database initialized")
			isReady = true
			return
		}
	}

}

func GetDB() (*pg.DB, bool) {
	return pgdb, isReady
}

func createSchema(db *pg.DB) error {
	err := db.Model((*Todo)(nil)).CreateTable(&orm.CreateTableOptions{
		Temp:        false,
		IfNotExists: true,
	})
	return err
}
