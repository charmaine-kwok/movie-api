package controllers

import (
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

// swagger:model NonMoviesListResponse
type NonMoviesListResponse ListResponse[models.NonMovie]

// swagger:model NonMovieItemResponse
type NonMovieItemResponse struct {
	Item models.NonMovie `json:"item"`
}

// @Summary Get a list of non-movie information
// @Tags Non-movies
// @Description Get a list of non-movie information
// @Accept json
// @Produce json
// @Param Authorization	header string true "Server JWT Token"
// @Param page query string false "Page Number"
// @Success 200 {object} NonMoviesListResponse "Non-Movie Information"
// @Failure 400  "Invalid user_id"
// @Failure 500  "Internal server error"
// @Router /non-movies [get]
func GetAllNonMovies() gin.HandlerFunc {
	return GetAllWrapper[models.NonMovie]()
}

// @Summary Get non-movie information by item id
// @Tags Non-movies
// @Description Get non-movie information by item id
// @Accept json
// @Produce json
// @Param Authorization	header string true "Server JWT Token"
// @Param itemId path string true "Item id"
// @Success 200 {object} NonMovieItemResponse "Non-Movie Information"
// @Failure 404  "No item found"
// @Failure 500  "Internal server error"
// @Router /non-movies/{itemId} [get]
func GetNonMoviesByItemId() gin.HandlerFunc {
	return GetByItemId(&models.NonMovie{})
}

// @Summary Create non-movie entry
// @Tags Non-movies
// @Description Create non-movie entry
// @Accept json
// @Produce json
// @Param Authorization	header string true "Server JWT Token"
// @Param body body NonMovieCreatorItem true "Non-Movie details"
// @Success 201 {object} NonMovieItemResponse "Non-Movie Information"
// @Failure 400  "Invalid request body"
// @Failure 500  "Internal server error"
// @Router /non-movies [post]
func CreateNonMovies() gin.HandlerFunc {
	return CreateGeneric(NonMovieCreator{})
}
