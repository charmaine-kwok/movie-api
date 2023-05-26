package controllers

import (
	"context"
	"go-crud/initializers"
	"go-crud/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllWrapper(collectionName string) gin.HandlerFunc {
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
		var movies []models.Movie
		totalItem, err := initializers.DB.Collection(collectionName).CountDocuments(ctx, bson.M{})
		totalPage := (totalItem-1)/limit + 1

		if err != nil {
			c.JSON(500, gin.H{
				"error": "Error getting totla number of movies",
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
			var movie models.Movie
			err := cursor.Decode(&movie)
			if err != nil {
				c.JSON(500, gin.H{
					"error": "Error decoding movie document",
				})
				return
			}
			movies = append(movies, movie)
		}

		// Respond with the movies
		c.JSON(200, gin.H{
			"movies":      movies,
			"totalItem":   totalItem,
			"totalPage":   totalPage,
			"currentPage": page,
		})
	}
}
