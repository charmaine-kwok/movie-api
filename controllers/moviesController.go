package controllers

import (
	"context"
	"go-crud/initializers"
	"go-crud/models"
	"reflect"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// swagger:model MovieInformationResponse
type MovieInformationResponse struct {
	Items       []models.Movie `json:"items"`
	TotalItem   int            `json:"totalItem"`
	TotalPage   int            `json:"totalPage"`
	CurrentPage int            `json:"currentPage"`
}

// @Summary Get a list of movie information by type
// @Tags Get All
// @Description Get a list of movie information by type
// @Accept json
// @Produce json
// @Param type path string true "Movies"
// @Success 200 {object} MovieInformationResponse "Movie Information"
// @Failure 400 {string} string "Invalid type"
// @Failure 500 {string} string "Internal server error"
// @Router /{type} [get]
func GetAllWrapper(collectionName string, model models.Model) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a context with a timeout
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Get the page number from the query parameters
		page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
		if err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid page number",
			})
			return
		}
		// Calculate the number of documents to skip
		var limit int64 = 10
		var skip int64 = (int64(page) - 1) * limit

		// Set the find options to limit the search results
		findOptions := options.Find()
		findOptions.SetLimit(limit)
		findOptions.SetSkip(skip)

		// Get the movies
		// var movies []models.Movie
		var items []models.Model
		totalItem, err := initializers.DB.Collection(collectionName).CountDocuments(ctx, bson.M{})
		totalPage := (totalItem-1)/limit + 1

		if err != nil {
			c.JSON(500, gin.H{
				"error": "Error getting total number of movies",
			})
			return
		}
		cursor, err := initializers.DB.Collection(collectionName).Find(ctx, bson.M{}, findOptions)
		if err != nil {
			c.JSON(500, gin.H{
				"error": "Error fetching movies",
			})
			return
		}
		defer cursor.Close(ctx)

		// Iterate through the cursor and decode each document into the movies slice
		for cursor.Next(ctx) {
			// var movie models.Movie
			// err := cursor.Decode(&movie)
			item := reflect.New(reflect.TypeOf(model).Elem()).Interface().(models.Model)
			err := cursor.Decode(item)
			if err != nil {
				c.JSON(500, gin.H{
					"error": "Error decoding movie document",
				})
				return
			}
			// movies = append(movies, movie)
			items = append(items, item)
		}

		// Respond with the movies
		c.JSON(200, gin.H{
			// "items":       movies,
			"items":       items,
			"totalItem":   totalItem,
			"totalPage":   totalPage,
			"currentPage": page,
		})
	}
}
