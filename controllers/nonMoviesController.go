package controllers

import (
	"context"
	"go-crud/initializers"
	"go-crud/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type NonMovieListInformationResponse struct {
	Items       []models.NonMovie `json:"items"`
	TotalItem   int               `json:"totalItem"`
	TotalPage   int               `json:"totalPage"`
	CurrentPage int               `json:"currentPage"`
}

// swagger:model NonMovieInformationResponse
type NonMovieInformationResponse struct {
	Item models.NonMovie `json:"item"`
}

// @Summary Get a list of non-movies information by type
// @Tags Non-movies
// @Description Get a list of non-movies information by type
// @Accept json
// @Produce json
// @Param page query string false "Page Number"
//
// @Success 200 {object} NonMovieListInformationResponse "Information"
// @Failure 400  "Invalid type"
// @Failure 500  "Internal server error"
// @Router /non-movies [get]
func GetAllNonMovies(c *gin.Context) {
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

	var items []models.NonMovie
	totalItem, err := initializers.DB.Collection("Non-movies").CountDocuments(ctx, bson.M{})
	totalPage := (totalItem-1)/limit + 1

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Error getting total number of items",
		})
		return
	}
	cursor, err := initializers.DB.Collection("Non-movies").Find(ctx, bson.M{}, findOptions)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Error fetching items",
		})
		return
	}
	defer cursor.Close(ctx)

	// Iterate through the cursor and decode each document into the movies slice
	for cursor.Next(ctx) {
		var item models.NonMovie

		err := cursor.Decode(&item)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error decoding document",
			})
			return
		}

		items = append(items, item)
	}

	c.JSON(http.StatusOK, gin.H{
		"items":       items,
		"totalItem":   totalItem,
		"totalPage":   totalPage,
		"currentPage": page,
	})
}

// @Summary Get non-movie information by Title
// @Tags Non-movies
// @Description Get non-movie information by Title
// @Accept json
// @Produce json
// @Param title path string true "Title"
//
// @Success 200 {object} NonMovieInformationResponse "Non-Movie Information"
// @Failure 400  "Invalid type"
// @Failure 500  "Internal server error"
// @Router /non-movies/details/{title} [get]
func GetNonMovieByTitle(c *gin.Context) {
	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	titleName := c.Param("title")

	// Determine the filter based on the model type

	filter := bson.M{"title": titleName}

	var item models.NonMovie

	err := initializers.DB.Collection("Non-movies").FindOne(ctx, filter).Decode(&item)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "No item found with the given title",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error fetching item",
			})
		}
		return
	}

	// Respond with the movie
	c.JSON(http.StatusOK, gin.H{
		"item": item,
	})
}
