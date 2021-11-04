package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

var (
	output      uuid.UUID
	currentTime string
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s: %s", currentTime, output.String())
}

func main() {
	output, _ = uuid.NewRandom()
	ticker := time.NewTicker(time.Second * 5)

	http.HandleFunc("/", handler)

	go func() {
		for ; true; <-ticker.C {
			currentTime = time.Now().Format(time.RFC1123)
			fmt.Printf("%s: %s\n", currentTime, output.String())
		}
	}()
	log.Fatal(http.ListenAndServe(":3000", nil))
}
