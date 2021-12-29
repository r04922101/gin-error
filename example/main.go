package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	ginerror "github.com/r04922101/gin-error"
)

func main() {
	r := gin.Default()
	r.Use(ginerror.RespondError)

	r.POST("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.POST("/error", func(c *gin.Context) {
		c.AbortWithError(http.StatusInternalServerError, errors.New("error occurs"))
	})

	if err := r.Run(); err != nil {
		log.Fatalf("failed to run gin: %v", err)
	}
}
