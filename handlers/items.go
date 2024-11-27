package handlers

import (
	"context"
	"net/http"

	"products/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetItemsHandler retrieves a list of items
// @Summary      Get Items
// @Description  Fetch a list of all items
// @Tags         Items
// @Produce      json
// @Success      200   {array}   models.Item
// @Failure      500   {object}  map[string]string
// @Router       /items [get]
func GetItemsHandler(db *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		collection := db.Database("products").Collection("items")
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

// GetItemDetailHandler retrieves the details of a specific item
// @Summary      Get Item by ID
// @Description  Fetch details of a specific item by ID
// @Tags         Items
// @Produce      json
// @Param        id    path      string  true  "Item ID"
// @Success      200   {object}  models.Item
// @Failure      404   {object}  map[string]string
// @Router       /items/{id} [get]
func GetItemDetailHandler(db *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		collection := db.Database("products").Collection("items")

		var item models.Item
		err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&item)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
			return
		}

		c.JSON(http.StatusOK, item)
	}
}

// CreateItemHandler add items
// @Summary      Add Item
// @Description  add a new item
// @Tags         Items
// @Produce      json
// @Param        item  body      models.Item  true  "Item Data"
// @Success      200   {object}  models.Item
// @Failure      404   {object}  map[string]string
// @Router       /items [post]
func CreateItemHandler(db *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var item models.Item
		if err := c.ShouldBindJSON(&item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		collection := db.Database("products").Collection("items")
		_, err := collection.InsertOne(context.Background(), item)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create item"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Item created successfully"})
	}
}
