package api_v1

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/kimhieu153255/first-go/internal/config/db/sqlc"
	handlers "github.com/kimhieu153255/first-go/pkg/handlers"
	"github.com/kimhieu153255/first-go/pkg/utils"
)

type createUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	FullName string `json:"full_name" binding:"required,min=3,max=100"`
	Password string `json:"password" binding:"required"`
}

type createUserResponse struct {
	ID         int64  `json:"id"`
	Email      string `json:"email"`
	FullName   string `json:"full_name"`
	Created_at string `json:"created_at"`
}

func newCreatedUserResponse(user db.User) createUserResponse {
	return createUserResponse{
		ID:         user.ID,
		Email:      user.Email,
		FullName:   user.FullName,
		Created_at: user.CreatedAt.String(),
	}
}

// Create user
func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, handlers.NewBadRequestError(err.Error()))
		return
	}

	newHashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, handlers.NewInternalServerError(err.Error()))
		return
	}

	arg := db.CreateUserParams{
		Email:    req.Email,
		FullName: req.FullName,
		Password: newHashPassword,
	}

	user, err := server.Store.CreateUser(ctx, arg)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, handlers.NewForbiddenError(err.Error()))
			return
		}
		ctx.JSON(http.StatusInternalServerError, handlers.NewInternalServerError(err.Error()))
		return
	}

	rsp := newCreatedUserResponse(user)
	ctx.JSON(http.StatusOK, handlers.NewSuccessResponse(rsp, "Create user successfully"))
}

type GetUserRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

// Get user by id
func (server *Server) getUserByID(ctx *gin.Context) {
	var req GetUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, handlers.NewBadRequestError(err.Error()))
		return
	}

	user, err := server.Store.GetUserById(ctx, req.ID)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, handlers.NewNotFoundError(err.Error()))
			return
		}
		ctx.JSON(http.StatusInternalServerError, handlers.NewInternalServerError(err.Error()))
		return
	}

	rsp := newCreatedUserResponse(user)
	ctx.JSON(http.StatusOK, handlers.NewSuccessResponse(rsp, "Get user successfully"))
}

// Health check
func (server *Server) healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, handlers.NewSuccessResponse(nil, "Health check successfully"))
}

// Get List user
func (server *Server) getListUser(ctx *gin.Context) {
	users, err := server.Store.GetListUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, handlers.NewInternalServerError(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, handlers.NewSuccessResponse(users, "Get list user successfully"))
}

type deleteUserByIDRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

// Delete user by id
func (server *Server) deleteUserByID(ctx *gin.Context) {
	var req deleteUserByIDRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, handlers.NewBadRequestError(err.Error()))
		return
	}

	user, err := server.Store.DeleteUserByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, handlers.NewNotFoundError(err.Error()))
			return
		}
		ctx.JSON(http.StatusInternalServerError, handlers.NewInternalServerError(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, handlers.NewSuccessResponse(user, "Delete user successfully"))
}
