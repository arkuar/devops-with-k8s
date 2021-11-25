package main

import (
	"bytes"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

func TodoLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		body, err1 := ioutil.ReadAll(c.Request.Body)
		if err1 != nil {
			log.Printf("Logger error parsing request body %s", err1.Error())
			return
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewReader(body))

		log.Printf("Received todo: %s\n", string(body))

	}
}
