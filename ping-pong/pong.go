package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"ping-pong/database"
)

var ready = make(chan bool, 1)

func fetchPongs() (pongs database.Ping) {
	db, _ := database.GetDB()
	err := db.Model(&pongs).First()
	if err != nil {
		panic(err)
	}
	return
}

func updatePongs(pong *database.Ping) {
	db, _ := database.GetDB()
	pong.Counter = pong.Counter + 1
	_, err := db.Model(pong).WherePK().Update()
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

func healthz(w http.ResponseWriter, r *http.Request) {
	_, isReady := database.GetDB()
	if isReady {
		w.WriteHeader(http.StatusOK)
	} else {
		fmt.Println("Database connection not yet ready")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func main() {
	database.InitDb()
	port, ok := os.LookupEnv("PINGPONG_PORT")
	if !ok {
		port = "3000"
	}
	http.HandleFunc("/", healthz)
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/pingpong", handler)
	http.HandleFunc("/pongs", pongHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
