package main

import (
	"context"
	"log"
	"net/http"

	"github.com/Teddy55Codes/GibbMongoDB/internal/store"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type Entry struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Note     string `json:"note"`
}

func main() {
	database := *store.Connect()

	router := gin.Default()
	router.Use(CORSMiddleware())

	router.POST("/entries", func(c *gin.Context) {
		var entry Entry
		if err := c.BindJSON(&entry); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		passwordDocument := map[string]interface{}{
			"1": map[string]string{
				"name":     entry.Name,
				"password": entry.Password,
			},
		}

		_, err := database.PasswordCollection.InsertOne(context.Background(), passwordDocument)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while inserting into password collection"})
			return
		}

		notesDocument := map[string]interface{}{
			"1": map[string]string{
				"note": entry.Note,
			},
		}

		_, err = database.NotesCollection.InsertOne(context.Background(), notesDocument)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while inserting into notes collection"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "Data inserted successfully"})
	})

	router.GET("/entries", func(c *gin.Context) {
		passwords, err := database.PasswordCollection.Find(context.Background(), bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while retrieving from password collection"})
			return
		}

		var passwordDocuments []map[string]interface{}
		if err = passwords.All(context.TODO(), &passwordDocuments); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while decoding password documents"})
			return
		}

		notes, err := database.NotesCollection.Find(context.Background(), bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while retrieving from notes collection"})
			return
		}

		var notesDocuments []map[string]interface{}
		if err = notes.All(context.TODO(), &notesDocuments); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while decoding notes documents"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"passwords": passwordDocuments, "notes": notesDocuments})
	})

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
