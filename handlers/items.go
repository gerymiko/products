package handlers

import (
	"context"
	"net/http"

	"products/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetItemsHandler(db *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		collection := db.Database("test").Collection("items")
		cursor, err := collection.Find(context.Background(), bson.D{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch items"})
			return
		}

		var items []models.Item
		if err := cursor.All(context.Background(), &items); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode items"})
			return
		}

		c.JSON(http.StatusOK, items)
	}
}

func GetItemDetailHandler(db *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		collection := db.Database("test").Collection("items")

		var item models.Item
		err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&item)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
			return
		}

		c.JSON(http.StatusOK, item)
	}
}
