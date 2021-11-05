package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("/usr/src/app/files/output")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "%s", string(data))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
