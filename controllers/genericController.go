package controllers

import (
	"context"
	"errors"
	"fmt"
	"go-crud/initializers"
	"go-crud/middleware"
	"go-crud/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllWrapper[T modelType]() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a context with a timeout
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Get the page number from the query parameters
		page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid page number",
			})
			return
		}

		var user models.User
		result := initializers.DB.Where("id = ?", middleware.USER_ID).First(&user)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Invalid user_id",
				})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Database error",
				})
			}
			return
		}

		// Calculate the number of documents to skip
		var limit int = 10
		var offset int = (int(page) - 1) * limit

		var items []T
		var totalItem int64

		// Get number of items and pages
		initializers.DB.Where("user_id = ?", middleware.USER_ID).Find(&items).Count(&totalItem)
		totalPage := (int(totalItem)-1)/limit + 1

		// Retrieve the movies with pagination and sorting
		err = initializers.DB.Limit(10).Offset(offset).Where("user_id = ?", "1").Order("date").Find(&items).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to retrieve items",
			})
			return
		}
		// Respond with the items
		c.JSON(http.StatusOK, gin.H{
			"items":       &items,
			"totalItem":   totalItem,
			"totalPage":   totalPage,
			"currentPage": page,
		})
	}
}

func GetByItemId(item models.Model) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a context with a timeout
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Get the title from the title parameter
		itemId := c.Param("item_id")

		// Retrieve the movies with pagination and sorting
		err := initializers.DB.Where("id = ?", itemId).First(&item).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{
					"error": "No item found",
				})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Failed to retrieve item",
				})
			}
			return
		}

		// Respond with the movie
		c.JSON(http.StatusOK, gin.H{
			"item": &item,
		})
	}
}

func parseAndValidateDate(date string) (time.Time, error) {
	layout := "02-01-2006"
	parsedDate, err := time.Parse(layout, date)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid date format, please enter a date in the format dd-mm-yyyy")
	}
	return parsedDate, nil
}

type modelType interface {
	models.Movie | models.NonMovie
}

type Creator interface {
	Create(c *gin.Context) (interface{}, error)
}

type MovieCreator struct{}

type MovieCreatorItem struct {
	Title_zh string `json:"title_zh" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Desc     string `json:"desc" binding:"required"`
	Location string `json:"location" binding:"required"`
	Date     string `json:"date" binding:"required"`
	Rating   string `json:"rating" binding:"required"`
	Pic      string `json:"pic" binding:"required"`
	Wiki_url string `json:"wiki_url" binding:"required"`
}

func (mc MovieCreator) Create(c *gin.Context) (interface{}, error) {
	var t MovieCreatorItem

	if err := c.ShouldBindJSON(&t); err != nil {
		return nil, err
	}

	parsedDate, err := parseAndValidateDate(t.Date)
	if err != nil {
		return nil, err
	}

	return &models.Movie{
		Title_zh: t.Title_zh,
		Title:    t.Title,
		Desc:     t.Desc,
		Location: t.Location,
		Date:     parsedDate,
		Rating:   t.Rating,
		Pic:      t.Pic,
		Wiki_url: t.Wiki_url,
		User_id:  middleware.USER_ID,
	}, nil
}

type NonMovieCreator struct{}

type NonMovieCreatorItem struct {
	Title    string `json:"title" binding:"required"`
	Desc     string `json:"desc" binding:"required"`
	Location string `json:"location" binding:"required"`
	Date     string `json:"date" binding:"required"`
	Rating   string `json:"rating" binding:"required"`
	Pic      string `json:"pic" binding:"required"`
}

func (nmc NonMovieCreator) Create(c *gin.Context) (interface{}, error) {
	var t NonMovieCreatorItem

	if err := c.ShouldBindJSON(&t); err != nil {
		return nil, err
	}

	parsedDate, err := parseAndValidateDate(t.Date)
	if err != nil {
		return nil, err
	}

	return &models.NonMovie{
		Title:    t.Title,
		Desc:     t.Desc,
		Location: t.Location,
		Date:     parsedDate,
		Rating:   t.Rating,
		Pic:      t.Pic,
		User_id:  middleware.USER_ID,
	}, nil
}

func CreateGeneric(creator Creator) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		item, err := creator.Create(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = initializers.DB.Create(item).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"item": item,
		})
	}
}
