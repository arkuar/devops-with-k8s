package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("/usr/src/app/files/output")
	checkErr(err)

	pongs, err := os.ReadFile("/usr/src/app/files/pongs")
	checkErr(err)

	fmt.Fprintf(w, "%s\n%s", string(data), string(pongs))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
