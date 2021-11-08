package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func fetchImage() {
	if _, err := os.Stat("/usr/src/app/cache/image.jpg"); errors.Is(err, os.ErrNotExist) {
		resp, err := http.Get("https://picsum.photos/1200")
		check(err)

		defer resp.Body.Close()

		file, err := os.Create("/usr/src/app/cache/image.jpg")
		check(err)

		defer file.Close()

		_, err = io.Copy(file, resp.Body)
		check(err)

		log.Println("Image fetched and cached")
	}
}

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	router := gin.Default()
	router.StaticFile("image.jpg", "/usr/src/app/cache/image.jpg")
	router.LoadHTMLGlob("public/*.html")
	router.GET("/", func(c *gin.Context) {
		fetchImage()
		c.HTML(http.StatusOK, "index.html", nil)
	})

	fmt.Printf("Server started in port %s\n", port)
	router.Run(":" + port)
}
