package main

import (
	"fmt"
	"log"
	"net/http"
	"ping-pong/database"

	"github.com/go-pg/pg/v10"
)

var pgDb *pg.DB

func fetchPongs() (pongs database.Ping) {
	err := pgDb.Model(&pongs).First()
	if err != nil {
		panic(err)
	}
	return
}

func updatePongs(pong *database.Ping) {
	pong.Counter = pong.Counter + 1
	_, err := pgDb.Model(pong).WherePK().Update()
	if err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	pongs := fetchPongs()
	fmt.Fprintf(w, "pong %d", pongs.Counter)
	// writeToFile(counter)
	updatePongs(&pongs)
}

func pongHandler(w http.ResponseWriter, r *http.Request) {
	pongs := fetchPongs()
	fmt.Fprintf(w, "Ping / Pongs: %d", pongs.Counter)
}

func main() {
	pgDb = database.InitDb()

	http.HandleFunc("/", handler)
	http.HandleFunc("/pongs", pongHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
