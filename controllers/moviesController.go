package controllers

import (
	"context"
	"go-crud/initializers"
	"go-crud/models"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// swagger:model MovieInformationResponse
type MoviesInformationResponse struct {
	Items       []models.Movie `json:"items"`
	TotalItem   int            `json:"totalItem"`
	TotalPage   int            `json:"totalPage"`
	CurrentPage int            `json:"currentPage"`
}
type MovieInformationResponse struct {
	Item models.Movie `json:"item"`
}

// @Summary Get a list of movie information by type
// @Tags Movies
// @Description Get a list of movie information by type
// @Accept json
// @Produce json
// @Param type path string true "Type" Enums(movies, others)
//
// @Param page query string false "Page Number"
//
// @Success 200 {object} MoviesInformationResponse "Movies Information"
// @Failure 400  "Invalid type"
// @Failure 500  "Internal server error"
// @Router /movies/{type} [get]
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

		var items []models.Model
		totalItem, err := initializers.DB.Collection(collectionName).CountDocuments(ctx, bson.M{})
		totalPage := (totalItem-1)/limit + 1

		if err != nil {
			c.JSON(500, gin.H{
				"error": "Error getting total number of items",
			})
			return
		}
		cursor, err := initializers.DB.Collection(collectionName).Find(ctx, bson.M{}, findOptions)
		if err != nil {
			c.JSON(500, gin.H{
				"error": "Error fetching items",
			})
			return
		}
		defer cursor.Close(ctx)

		// Iterate through the cursor and decode each document into the movies slice
		for cursor.Next(ctx) {
			item := reflect.New(reflect.TypeOf(model).Elem()).Interface().(models.Model)
			err := cursor.Decode(item)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Error decoding document",
				})
				return
			}
			// movies = append(movies, movie)
			items = append(items, item)
		}

		// Respond with the movies
		c.JSON(http.StatusOK, gin.H{
			// "items":       movies,
			"items":       items,
			"totalItem":   totalItem,
			"totalPage":   totalPage,
			"currentPage": page,
		})
	}
}

// @Summary Get movie information by Title
// @Tags Movies
// @Description Get movie information by Title
// @Accept json
// @Produce json
// @Param type path string true "Type" Enums(movies, others)
// @Param title path string true "Title"
//
// @Success 200 {object} MovieInformationResponse "Movie Information"
// @Failure 400  "Invalid type"
// @Failure 500  "Internal server error"
// @Router /movies/{type}/details/{title} [get]
func GetByTitle(collectionName string, model models.Model) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a context with a timeout
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		titleName := c.Param("title")

		// Determine the filter based on the model type
		var filter bson.M
		switch model.(type) {
		case *models.NonMovie:
			filter = bson.M{"title": titleName}
		case *models.Movie:
			filter = bson.M{"title_en": titleName}
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Invalid model type",
			})
			return
		}
		err := initializers.DB.Collection(collectionName).FindOne(ctx, filter).Decode(model)

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
			"item": model,
		})
	}
}

// @Summary Create movie entry
// @Tags Movies
// @Description Create movie entry
// @Accept json
// @Produce json
// @Param type path string true "Type" Enums(movies, others)
//
// @Success 200 {object} MovieInformationResponse "Movie Information"
// @Failure 400  "Invalid request body"
// @Failure 500  "Internal server error"
// @Router /movies/{type} [post]
func CreateMovie(collectionName string, model models.Model) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a context with a timeout
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Bind the request body to the movie model
		if err := c.ShouldBindJSON(model); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request body",
			})
			return
		}

		// Insert the movie into the database
		_, err := initializers.DB.Collection(collectionName).InsertOne(ctx, model)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
			return
		}

		// Respond with the created movie
		c.JSON(http.StatusCreated, gin.H{
			"item": model,
		})
	}
}
