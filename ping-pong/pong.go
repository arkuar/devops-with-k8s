package main

import (
	"fmt"
	"log"
	"net/http"
)

var counter int = 0

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong %d", counter)
	counter++
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
