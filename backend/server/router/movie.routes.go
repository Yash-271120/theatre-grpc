package router

import (
	"net/http"

	response "github.com/Yash-271120/theatre-grpc/backend/server/response"
	"github.com/gin-gonic/gin"
)

func addMovieRoutes(rg *gin.RouterGroup) {
	rg.GET("/movies", func(c *gin.Context) {
		response.SuccessResponse(c, http.StatusOK, "Movies")
	})
	rg.POST("/movie", func(c *gin.Context) {
		response.SuccessResponse(c, http.StatusOK, "Movie")
	})
}
