package main

import (
	"log"

	"github.com/Teddy55Codes/GibbMongoDB/internal/api"
	"github.com/Teddy55Codes/GibbMongoDB/internal/store"
	"github.com/gin-gonic/gin"
)

func main() {
	database := *store.Connect()

	router := gin.Default()
	router.Use(CORSMiddleware())

	rout := api.Constructor(database)

	router.POST("/entries", rout.PostEntry)
	router.GET("/entries", rout.GetEntry)
	router.PATCH("/entries/:id", rout.PutEntry)

	router.Static("/web", "./web")

	log.Println("Server runs at :8080")
	router.Run(":8080")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
