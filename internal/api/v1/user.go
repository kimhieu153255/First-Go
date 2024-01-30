package api_v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/kimhieu153255/first-go/internal/config/db/sqlc"
	error "github.com/kimhieu153255/first-go/pkg/handlers"
)

type createUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	FullName string `json:"full_name" binding:"required,min=3,max=100"`
	Password string `json:"password" binding:"required,min=6,max=100"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, error.NewBadRequestError(err.Error()))
		return
	}

	arg := db.CreateUserParams{
		Email:    req.Email,
		FullName: req.FullName,
		Password: req.Password,
	}

	user, err := server.Store.CreateUser(ctx, arg)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, error.NewForbiddenError(err.Error()))
			return
		}
		ctx.JSON(http.StatusInternalServerError, error.NewInternalServerError())
		return
	}

	ctx.JSON(http.StatusOK, user)
}
