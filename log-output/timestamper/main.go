package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

var (
	dir         = "files"
	output      uuid.UUID
	currentTime string
)

func main() {
	output, _ = uuid.NewRandom()
	ticker := time.NewTicker(time.Second * 5)

	// Create directory if it is not present
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	for ; true; <-ticker.C {
		f, err := os.Create(dir + "/output")
		if err != nil {
			log.Fatal("Error creating file", err)
		}
		defer f.Close()

		currentTime = time.Now().Format(time.RFC1123)
		fmt.Fprintf(f, "%s: %s\n", currentTime, output.String())
	}
}
