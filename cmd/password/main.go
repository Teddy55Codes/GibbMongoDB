package main

import (
	"log"
    "github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

    router.Static("/", "./web")

	log.Println("Server runs at :8080")
	router.Run(":8080")
}
