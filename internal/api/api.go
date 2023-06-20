package api

import (
	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"github.com/Teddy55Codes/GibbMongoDB/internal/store"
	"github.com/gin-gonic/gin"
)


type Router struct {
    database store.Database
}

func Constructor(database store.Database) *Router {
	return &Router{database: database}
}

type entry struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Note     string `json:"note"`
}

func (r *Router) PostEntities(c *gin.Context) {
		var entry entry
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

		_, err := r.database.PasswordCollection.InsertOne(context.Background(), passwordDocument)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while inserting into password collection"})
			return
		}

		notesDocument := map[string]interface{}{
			"1": map[string]string{
				"note": entry.Note,
			},
		}

		_, err = r.database.NotesCollection.InsertOne(context.Background(), notesDocument)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while inserting into notes collection"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "Data inserted successfully"})
}

func (r *Router) GetEntities (c *gin.Context) {
		passwords, err := r.database.PasswordCollection.Find(context.Background(), bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while retrieving from password collection"})
			return
		}

		var passwordDocuments []map[string]interface{}
		if err = passwords.All(context.TODO(), &passwordDocuments); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while decoding password documents"})
			return
		}

		notes, err := r.database.NotesCollection.Find(context.Background(), bson.M{})
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
	}
