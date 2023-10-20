package response

import (
	"github.com/gin-gonic/gin"
)

func getResponse(code int, message string, data interface{}) gin.H {
	return gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	}
}

func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(code, getResponse(code, "success", data))
}

func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(code, getResponse(code, message, nil))
}
