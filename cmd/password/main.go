package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

type Entry struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Note     string `json:"note"`
}

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	passwordCollection := client.Database("myDatabase").Collection("passwords")
	notesCollection := client.Database("myDatabase").Collection("notes")

	router := gin.Default()

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

		_, err := passwordCollection.InsertOne(context.Background(), passwordDocument)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while inserting into password collection"})
			return
		}

		notesDocument := map[string]interface{}{
			"1": map[string]string{
				"note": entry.Note,
			},
		}

		_, err = notesCollection.InsertOne(context.Background(), notesDocument)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while inserting into notes collection"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "Data inserted successfully"})
	})

	router.GET("/entries", func(c *gin.Context) {
		passwords, err := passwordCollection.Find(context.Background(), bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while retrieving from password collection"})
			return
		}

		var passwordDocuments []map[string]interface{}
		if err = passwords.All(context.TODO(), &passwordDocuments); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while decoding password documents"})
			return
		}

		notes, err := notesCollection.Find(context.Background(), bson.M{})
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

	router.Static("/static", "./web")

	log.Println("Server runs at :8080")
	router.Run(":8080")
}
