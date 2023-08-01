package controllers

import (
	"context"
	"fmt"
	"go-crud/initializers"
	"go-crud/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func parseAndValidateDate(date string) (time.Time, error) {
	layout := "02-01-2006"
	parsedDate, err := time.Parse(layout, date)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid date format, please enter a date in the format dd-mm-yyyy")
	}
	return parsedDate, nil
}

type Creator interface {
	Create(collectionName string, c *gin.Context) (interface{}, error)
}

type MovieCreator struct{}

func (mc MovieCreator) Create(collectionName string, c *gin.Context) (interface{}, error) {
	var t struct {
		Title_zh string `json:"title_zh" binding:"required"`
		Title_en string `json:"title_en" binding:"required"`
		Desc     string `json:"desc" binding:"required"`
		Location string `json:"location" binding:"required"`
		Date     string `json:"date" binding:"required"`
		Rating   string `json:"rating" binding:"required"`
		Pic      string `json:"pic" binding:"required"`
		Wiki_url string `json:"wiki_url" binding:"required"`
	}

	if err := c.ShouldBindJSON(&t); err != nil {
		return nil, err
	}

	parsedDate, err := parseAndValidateDate(t.Date)
	if err != nil {
		return nil, err
	}

	return &models.Movie{
		Title_zh: t.Title_zh,
		Title_en: t.Title_en,
		Desc:     t.Desc,
		Location: t.Location,
		Date:     parsedDate,
		Rating:   t.Rating,
		Pic:      t.Pic,
		Wiki_url: t.Wiki_url,
	}, nil
}

type NonMovieCreator struct{}

func (nmc NonMovieCreator) Create(collectionName string, c *gin.Context) (interface{}, error) {
	var t struct {
		Title    string `json:"title" binding:"required"`
		Desc     string `json:"desc" binding:"required"`
		Location string `json:"location" binding:"required"`
		Date     string `json:"date" binding:"required"`
		Rating   string `json:"rating" binding:"required"`
		Pic      string `json:"pic" binding:"required"`
	}

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
	}, nil
}

func CreateGeneric(collectionName string, creator Creator) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		model, err := creator.Create(collectionName, c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		_, err = initializers.DB.Collection(collectionName).InsertOne(ctx, model)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"item": model,
		})
	}
}
