package api_v1

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/kimhieu153255/first-go/internal/config/db/sqlc"
	"github.com/kimhieu153255/first-go/pkg/handlers"
	"github.com/kimhieu153255/first-go/pkg/token"
)

type transferRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required"`
	ToAccountID   int64  `json:"to_account_id" binding:"required"`
	Amount        int64  `json:"amount" binding:"required"`
	Currency      string `json:"currency" binding:"required,currency"`
}

func (server *Server) Transfer(ctx *gin.Context) {
	var req transferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, handlers.NewBadRequestError(err.Error()))
		return
	}

	fromAccount, ok := server.ValidAccount(ctx, req.FromAccountID, req.Currency)
	if !ok {
		return
	}

	_, ok = server.ValidAccount(ctx, req.ToAccountID, req.Currency)
	if !ok {
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if fromAccount.UserID != authPayload.UserID {
		ctx.JSON(http.StatusUnauthorized, handlers.NewUnAuthorizedError("Account doesn't belong to the user"))
		return
	}

	transferResult, err := server.Store.TransferTx(ctx, db.AddTransferParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
		Currency:      req.Currency,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, handlers.NewInternalServerError(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, handlers.NewSuccessResponse(transferResult, "Transfer successfully"))
}

func (server *Server) ValidAccount(ctx *gin.Context, accountID int64, currency string) (db.Account, bool) {
	account, err := server.Store.GetAccountByID(ctx, accountID)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, handlers.NewNotFoundError(err.Error()))
			return db.Account{}, false
		}
		ctx.JSON(http.StatusInternalServerError, handlers.NewInternalServerError(err.Error()))
		return db.Account{}, false
	}

	if account.Currency != currency {
		errStr := fmt.Sprintf("Account [%d] currency is not match", accountID)
		ctx.JSON(http.StatusBadRequest, handlers.NewBadRequestError(errStr))
		return db.Account{}, false
	}

	return account, true
}
