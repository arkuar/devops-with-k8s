package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"project/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func checkTodoErr(c *gin.Context, err error) {
	if err != nil {
		c.String(500, err.Error())
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
	pgdb := db.GetDB()

	var todos []db.Todo

	err := pgdb.Model(&todos).Select()

	checkTodoErr(c, err)

	c.JSON(200, gin.H{
		"todos": todos,
	})
}

func addTodo(c *gin.Context) {
	var todo db.Todo

	c.BindJSON(&todo)

	pgdb := db.GetDB()

	_, err := pgdb.Model(&todo).Insert()
	checkTodoErr(c, err)

	c.JSON(201, gin.H{
		"todo": todo,
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

	db.InitDb()

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
