package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Teddy55Codes/GibbMongoDB/internal/store"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

type password struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type note struct {
	Note string `json:"note"`
}

func (r *Router) PostEntry(c *gin.Context) {
	var entry entry
	if err := c.BindJSON(&entry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	passwd := password{
		Name:     entry.Name,
		Password: entry.Password,
	}

	_, err := r.database.PasswordCollection.InsertOne(context.Background(), passwd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while inserting into password collection"})
		return
	}

	_, err = r.database.NotesCollection.InsertOne(context.Background(), note{Note: entry.Note})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while inserting into notes collection"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Data inserted successfully"})
}

func (r *Router) GetEntry(c *gin.Context) {
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

	fmt.Println(passwordDocuments)
	c.JSON(http.StatusOK, gin.H{"passwords": passwordDocuments, "notes": notesDocuments})
}

func (r *Router) PutEntry(c *gin.Context) {
	// Get the ID of the document to update from the request URL
	id := c.Param("id")

	// Convert the ID string to an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Get the updated data from the request body
	var updatedData map[string]interface{}
	if err := c.BindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a filter to find the document by its ID
	filter := bson.M{"_id": objectID}

	// Create an update to specify the changes
	update := bson.M{"$set": updatedData}

	// Perform the update operation
	result, err := r.database.PasswordCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Check if any document was updated
	if result.ModifiedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Document not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Document updated successfully"})
}

func (r *Router) DeleteEntry(c *gin.Context) {
	// Get the ID of the document to delete from the request URL
	id := c.Param("id")

	// Create a filter to find the document by its ID
	filter := bson.M{"_id": id}

	// Perform the delete operation
	result, err := r.database.PasswordCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Check if any document was deleted
	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Document not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Document deleted successfully"})
}
