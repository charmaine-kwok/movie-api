package controllers

import (
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

type NonMoviesInformationResponse struct {
	Items       []models.NonMovie `json:"items"`
	TotalItem   int               `json:"totalItem"`
	TotalPage   int               `json:"totalPage"`
	CurrentPage int               `json:"currentPage"`
}

// swagger:model NonMovieInformationResponse
type NonMovieInformationResponse struct {
	Item models.NonMovie `json:"item"`
}

// @Summary Get a list of non-movies information
// @Tags Non-movies
// @Description Get a list of non-movies information
// @Accept json
// @Produce json
// @Param page query string false "Page Number"
//
// @Success 200 {object} NonMoviesInformationResponse "Information"
// @Failure 400  "Invalid type"
// @Failure 500  "Internal server error"
// @Router /non-movies [get]
func GetAllNonMovies() gin.HandlerFunc {
	return GetAllWrapper("Non-movies", &models.NonMovie{})
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
func GetNonMovieByTitle() gin.HandlerFunc {
	return GetByTitle("Non-movies", &models.NonMovie{})
}

// @Summary Create non-movie entry
// @Tags Non-movies
// @Description Create non-movie entry
// @Accept json
// @Produce json
// @Param body body models.NonMovie true "Non-Movie details"
// @Success 200 {object} NonMovieInformationResponse "Non-Movie Information"
// @Failure 400  "Invalid request body"
// @Failure 500  "Internal server error"
// @Router /non-movies [post]
func CreateNonMovie(collectionName string) gin.HandlerFunc {
	return CreateGeneric(collectionName, NonMovieCreator{})
}
