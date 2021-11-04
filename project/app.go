package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	fmt.Printf("Server started in port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
