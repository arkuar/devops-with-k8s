package main

import (
	"fmt"
	"log"
	"net/http"
)

var counter int = 0

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong %d", counter)
	// writeToFile(counter)
	counter++
}

func pongHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ping / Pongs: %d", counter)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/pongs", pongHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
