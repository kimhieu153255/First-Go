package handler_error

import "github.com/gin-gonic/gin"

const (
	BadRequest = 400
	Forbidden  = 403
	NotFound   = 404
	Internal   = 500
)

func NewBadRequestError(message string) gin.H {
	return gin.H{
		"error": gin.H{
			"code":    BadRequest,
			"message": message,
		},
	}
}

func NewNotFoundError(message string) gin.H {
	return gin.H{
		"error": gin.H{
			"code":    NotFound,
			"message": message,
		},
	}
}

func NewForbiddenError(message string) gin.H {
	return gin.H{
		"error": gin.H{
			"code":    Forbidden,
			"message": message,
		},
	}
}

func NewInternalServerError() gin.H {
	return gin.H{
		"error": gin.H{
			"code":    Internal,
			"message": "Internal Server Error",
		},
	}
}
