package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var counter int = 0
var dir = "/usr/src/app/files"

func writeToFile(count int) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
	f, err := os.Create(dir + "/pongs")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Fprintf(f, "Ping / Pongs: %d", count)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong %d", counter)
	writeToFile(counter)
	counter++
}

func main() {
	// Initial write
	writeToFile(counter)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
