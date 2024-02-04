package handlers

import "github.com/gin-gonic/gin"

const (
	Created        = 201
	BadRequest     = 400
	UnAuthorized   = 401
	Forbidden      = 403
	NotFound       = 404
	Internal       = 500
	NetwordTimeOut = 504
)

func NewBadRequestError(message string) gin.H {
	return gin.H{
		"code":    BadRequest,
		"message": message,
	}
}

func NewNotFoundError(message string) gin.H {
	return gin.H{
		"code":    NotFound,
		"message": message,
	}
}

func NewForbiddenError(message string) gin.H {
	return gin.H{
		"code":    Forbidden,
		"message": message,
	}
}

func NewInternalServerError(message string) gin.H {
	return gin.H{
		"code":    Internal,
		"message": message,
	}
}

func NewUnAuthorizedError(message string) gin.H {
	return gin.H{
		"code":    UnAuthorized,
		"message": message,
	}
}

func NewNetwordTimeOutError(message string) gin.H {
	return gin.H{
		"code":    NetwordTimeOut,
		"message": message,
	}
}

func NewCreatedError(message string) gin.H {
	return gin.H{
		"code":    Created,
		"message": message,
	}
}
