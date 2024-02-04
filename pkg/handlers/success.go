package handlers

import (
	"github.com/gin-gonic/gin"
)

func NewSuccessResponse(data interface{}, massage string) gin.H {
	return gin.H{
		"code":    200,
		"message": massage,
		"data":    data,
	}
}
