package controllers

import (
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

type ListResponse[T models.Model] struct {
	Items       *[]T  `json:"items"`
	TotalItem   int64 `json:"totalItem" example:"1"`
	TotalPage   int   `json:"totalPage" example:"1"`
	CurrentPage int   `json:"currentPage" example:"1"`
}

type ItemResponse[T models.Model] struct {
	Item T `json:"item"`
}

type MoviesListResponse ListResponse[*models.Movie]

type MovieItemResponse ItemResponse[*models.Movie]

// GetAllMovies gets a list of movie information
//
//	@Summary		Get a list of movie information
//	@Tags			Movies
//	@Description	Get a list of movie information
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string				true	"Server JWT Token"
//	@Param			page			query		string				false	"Page Number"
//	@Param			order_by		query		string				false	"Order by"
//	@Success		200				{object}	MoviesListResponse	"Movies Information"
//	@Failure		400				"Invalid user_id"
//	@Failure		500				"Internal server error"
//	@Router			/movies [get]
func GetAllMovies() gin.HandlerFunc {
	return GetAllWrapper[*models.Movie]()
}

// GetMoviesByItemId gets movie information by item id
//
//	@Summary		Get movie information by item id
//	@Tags			Movies
//	@Description	Get movie information by item id
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string				true	"Server JWT Token"
//	@Param			itemId			path		number				true	"Item id"
//	@Success		200				{object}	MovieItemResponse	"Movie Information"
//	@Failure		404				"No item found"
//	@Failure		500				"Internal server error"
//	@Router			/movies/{itemId} [get]
func GetMoviesByItemId() gin.HandlerFunc {
	return GetByItemId[*models.Movie]()
}

// CreateMovies creates movie entry
//
//	@Summary		Create movie entry
//	@Tags			Movies
//	@Description	Create movie entry
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string				true	"Server JWT Token"
//	@Param			body			body		MovieCreatorItem	true	"Movie details"
//	@Success		201				{object}	MovieItemResponse	"Movie Information"
//	@Failure		400				"Invalid request body"
//	@Failure		500				"Internal server error"
//	@Router			/movies [post]
func CreateMovies() gin.HandlerFunc {
	return CreateGeneric[MovieCreator]()
}
