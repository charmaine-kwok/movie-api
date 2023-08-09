package controllers

import (
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

type ListResponse[T modelType] struct {
	Items       []T `json:"items"`
	TotalItem   int `json:"totalItem"`
	TotalPage   int `json:"totalPage"`
	CurrentPage int `json:"currentPage"`
}

// swagger:model MoviesListResponse
type MoviesListResponse ListResponse[models.Movie]

// swagger:model MovieItemResponse
type MovieItemResponse struct {
	Item models.Movie `json:"item"`
}

// @Summary Get a list of movie information
// @Tags Movies
// @Description Get a list of movie information
// @Accept json
// @Produce json
// @Param Authorization	header string true "Server JWT Token"
// @Param page query string false "Page Number"
// @Success 200 {object} MoviesListResponse "Movies Information"
// @Failure 400  "Invalid user_id"
// @Failure 500  "Internal server error"
// @Router /movies [get]
func GetAllMovies() gin.HandlerFunc {
	return GetAllWrapper[models.Movie]()
}

// @Summary Get movie information by item id
// @Tags Movies
// @Description Get movie information by item id
// @Accept json
// @Produce json
// @Param Authorization	header string true "Server JWT Token"
// @Param itemId path number true "Item id"
// @Success 200 {object} MovieItemResponse "Movie Information"
// @Failure 404  "No item found"
// @Failure 500  "Internal server error"
// @Router /movies/{itemId} [get]
func GetMoviesByItemId() gin.HandlerFunc {
	return GetByItemId(&models.Movie{})
}

// @Summary Create movie entry
// @Tags Movies
// @Description Create movie entry
// @Accept json
// @Produce json
// @Param Authorization	header string true "Server JWT Token"
// @Param body body MovieCreatorItem true "Movie details"
// @Success 201 {object} MovieItemResponse "Movie Information"
// @Failure 400  "Invalid request body"
// @Failure 500  "Internal server error"
// @Router /movies [post]
func CreateMovies() gin.HandlerFunc {
	return CreateGeneric(MovieCreator{})
}
