package main

import (
	"errors"
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
	pgdb, _ := db.GetDB()

	var todos []db.Todo

	err := pgdb.Model(&todos).Select()

	if err != nil {
		c.String(500, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"todos": todos,
	})
}

func addTodo(c *gin.Context) {
	var todo db.Todo

	err := c.ShouldBindJSON(&todo)
	if err != nil {
		c.Error(err)
		c.String(500, err.Error())
		return
	}

	err = todo.Validate()

	if err != nil {
		c.Error(err)
		c.String(400, err.Error())
		return
	}

	pgdb, _ := db.GetDB()

	_, err = pgdb.Model(&todo).Insert()
	if err != nil {
		c.Error(err)
		c.String(500, err.Error())
		return
	}

	log.Println("Stored todo to database")

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

	go db.InitDb()

	config := cors.DefaultConfig()

	config.AllowOrigins = []string{allowedOrigin}

	router := gin.Default()
	log.SetOutput(os.Stdout)
	router.Use(cors.New(config))

	router.GET("/", func(c *gin.Context) {
		// Health check
		c.Status(http.StatusOK)
	})

	router.GET("/healthz", func(c *gin.Context) {
		_, isReady := db.GetDB()
		if isReady {
			c.Status(http.StatusOK)
		} else {
			c.Status(http.StatusInternalServerError)
		}
	})

	router.GET("/api/image", fetchImage)

	router.GET("/api/todos", fetchTodos)

	router.POST("/api/todos", TodoLogger(), addTodo)

	router.Run(":" + port)
}
