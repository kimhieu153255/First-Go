package api_v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/kimhieu153255/first-go/internal/config/db/sqlc"
	"github.com/kimhieu153255/first-go/pkg/handlers"
)

type createAccountRequest struct {
	Balance  int64  `json:"balance" binding:"required"`
	Currency string `json:"currency" binding:"required"`
}

func (s *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, handlers.NewBadRequestError(err.Error()))
		return
	}

	arg := db.CreateAccountParams{
		Balance:  req.Balance,
		Currency: req.Currency,
	}

	account, err := s.Store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, handlers.NewInternalServerError(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, account)
}
