package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func main() {
	output, _ := uuid.NewRandom()
	ticker := time.NewTicker(time.Second * 5)
	for ; true; <-ticker.C {
		fmt.Printf("%s: %s\n", time.Now().Format(time.RFC1123), output.String())
	}
}
