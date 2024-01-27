package handler_error

import "github.com/gin-gonic/gin"

func NewBadRequestError(message string) gin.H {
	return gin.H{
		"error": gin.H{
			"code":    400,
			"message": message,
		},
	}
}

func NewNotFoundError(message string) gin.H {
	return gin.H{
		"error": gin.H{
			"code":    404,
			"message": message,
		},
	}
}

func NewInternalServerError() gin.H {
	return gin.H{
		"error": gin.H{
			"code":    500,
			"message": "Internal Server Error",
		},
	}
}
