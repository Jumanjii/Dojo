package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type status struct {
	Status string
}

func main() {
	fmt.Println("Start Duckland API")

	router := gin.Default()
	router.GET("healthcheck", func(ctx *gin.Context) {
		ctx.JSON(200, status{Status: "ok"})
	})

	router.Run(":3000")
}
