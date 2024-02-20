package api_v1

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	handlers "github.com/kimhieu153255/first-go/pkg/handlers"
	"github.com/kimhieu153255/first-go/pkg/utils"
)

type loginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type userRes struct {
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
}

type loginResponse struct {
	AccessToken         string    `json:"access_token"`
	ExpiredAccessToken  time.Time `json:"access_token_expires_at"`
	RefreshToken        string    `json:"refresh_token"`
	ExpiredRefreshToken time.Time `json:"refresh_token_expires_at"`
	User                userRes   `json:"user"`
}

func (server *Server) login(ctx *gin.Context) {
	var req loginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, handlers.NewBadRequestError(err.Error()))
		return
	}

	user, err := server.Store.GetUserByEmail(ctx, req.Email)
	if err != nil {
		ctx.JSON(http.StatusForbidden, handlers.NewForbiddenError(err.Error()))
		return
	}

	if !utils.CheckPassword(user.Password, req.Password) {
		ctx.JSON(http.StatusUnauthorized, handlers.NewForbiddenError("Invalid password"))
		return
	}

	accessToken, accessPayload, err := server.TokenMaker.CreateToken(user.Email, user.Role, user.FullName, server.Config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, handlers.NewInternalServerError(err.Error()))
		return
	}

	refeshToken, refeshPayload, err := server.TokenMaker.CreateToken(user.Email, user.Role, user.FullName, server.Config.RefreshTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, handlers.NewInternalServerError(err.Error()))
		return
	}

	res := loginResponse{
		User:                userRes{Email: user.Email, Fullname: user.FullName},
		AccessToken:         accessToken,
		ExpiredAccessToken:  accessPayload.Expire_at,
		RefreshToken:        refeshToken,
		ExpiredRefreshToken: refeshPayload.Expire_at,
	}

	ctx.JSON(http.StatusOK, handlers.NewSuccessResponse(res, "Login successfully"))

}

func (server *Server) register(ctx *gin.Context) {
	server.createUser(ctx)
}
