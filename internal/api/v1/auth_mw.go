package api_v1

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kimhieu153255/first-go/pkg/handlers"
	"github.com/kimhieu153255/first-go/pkg/token"
)

const (
	authorizationHeaderKey       = "authorization"
	authorizationTypeBearer      = "bearer"
	authorizationPayloadKey      = "authorization_payload"
	authorizationHeaderErr       = "authorization header is not provided"
	authorizationHeaderFormatErr = "authorization header is not provided in the right format"
)

func AuthMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, handlers.NewUnAuthorizedError(authorizationHeaderErr))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, handlers.NewUnAuthorizedError(authorizationHeaderFormatErr))
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, handlers.NewUnAuthorizedError(authorizationHeaderFormatErr))
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, handlers.NewUnAuthorizedError(err.Error()))
			return
		}

		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}
