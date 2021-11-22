package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var todos = []string{"TODO 1", "TODO 2"}

type Data struct {
	Todo string `json:"todo"`
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func fetchImage(c *gin.Context) {
	if _, err := os.Stat("/usr/src/app/cache/image.jpg"); errors.Is(err, os.ErrNotExist) {
		resp, err := http.Get("https://picsum.photos/1200")
		check(err)

		defer resp.Body.Close()

		file, err := os.Create("/usr/src/app/cache/image.jpg")
		check(err)

		defer file.Close()

		_, err = io.Copy(file, resp.Body)
		check(err)
		log.Println("Cached image")
	}

	c.File("/usr/src/app/cache/image.jpg")
}

func fetchTodos(c *gin.Context) {
	c.JSON(200, gin.H{
		"todos": todos,
	})
}

func addTodo(c *gin.Context) {
	var data Data
	c.BindJSON(&data)

	todos = append(todos, data.Todo)

	c.JSON(201, gin.H{
		"todo": data.Todo,
	})
}

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	allowedOrigin := os.Getenv("REQUEST_ORIGIN")
	if len(allowedOrigin) == 0 {
		allowedOrigin = "http://localhost"
	}

	config := cors.DefaultConfig()

	config.AllowOrigins = []string{allowedOrigin}

	router := gin.Default()
	router.Use(cors.New(config))

	router.GET("/image", fetchImage)

	router.GET("/todos", fetchTodos)

	router.POST("/todos", addTodo)

	fmt.Printf("Server started in port %s\n", port)
	router.Run(":" + port)
}
