package api_v1

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/kimhieu153255/first-go/internal/config/db/sqlc"
	"github.com/kimhieu153255/first-go/pkg/handlers"
	"github.com/kimhieu153255/first-go/pkg/token"
)

type createAccountRequest struct {
	Balance  int64  `json:"balance" binding:"required"`
	Currency string `json:"currency" binding:"required,currency"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, handlers.NewBadRequestError(err.Error()))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.CreateAccountParams{
		Balance:  req.Balance,
		Currency: req.Currency,
		UserID:   authPayload.UserID,
	}

	account, err := server.Store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, handlers.NewInternalServerError(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, handlers.NewSuccessResponse(account, "Create account successfully"))
}

type GetAccountRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

func (server *Server) getAccountByID(ctx *gin.Context) {
	var req GetAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, handlers.NewBadRequestError(err.Error()))
		return
	}

	account, err := server.Store.GetAccountByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, handlers.NewNotFoundError(err.Error()))
			return
		}
		ctx.JSON(http.StatusInternalServerError, handlers.NewInternalServerError(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, handlers.NewSuccessResponse(account, "Get account successfully"))
}

type deleteAccountRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

func (server *Server) deleteAccountByID(ctx *gin.Context) {
	var req deleteAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, handlers.NewBadRequestError(err.Error()))
		return
	}

	account, err := server.Store.DeleteAccountByID(ctx, req.ID)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, handlers.NewNotFoundError(err.Error()))
			return
		}
		ctx.JSON(http.StatusInternalServerError, handlers.NewInternalServerError(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, handlers.NewSuccessResponse(account, "Delete account successfully"))
}

func (server *Server) getAccounts(ctx *gin.Context) {
	accounts, err := server.Store.GetListAccounts(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, handlers.NewInternalServerError(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, handlers.NewSuccessResponse(accounts, "Get accounts successfully"))
}
