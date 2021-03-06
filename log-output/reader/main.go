package main

import (
	"fmt"
	"io"
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

	resp, err := http.Get("http://ping-pong-svc/pongs")
	checkErr(err)

	defer resp.Body.Close()

	pongs, err := io.ReadAll(resp.Body)
	checkErr(err)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n%s\n%s", os.Getenv("MESSAGE"), string(data), string(pongs))
}

func healthz(w http.ResponseWriter, r *http.Request) {
	_, err := http.Get("http://ping-pong-svc/")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/healthz", healthz)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
