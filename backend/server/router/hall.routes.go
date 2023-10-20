package router

import (
	"net/http"

	response "github.com/Yash-271120/theatre-grpc/backend/server/response"
	"github.com/gin-gonic/gin"
)

func addHallRoutes(rg *gin.RouterGroup) {
	rg.GET("/hall/:movieId", func(c *gin.Context) {
		movieId := c.Param("movieId")
		response.SuccessResponse(c, http.StatusOK, "Hall for movie"+movieId)
	})
}
