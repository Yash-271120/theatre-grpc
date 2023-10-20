package router

import (
	"net/http"

	response "github.com/Yash-271120/theatre-grpc/backend/server/response"
	"github.com/gin-gonic/gin"
)

func addTicketRoutes(rg *gin.RouterGroup) {
	rg.POST("/ticket", func(c *gin.Context) {
		response.SuccessResponse(c, http.StatusOK, "Ticket booked")
	})
}
