package api_v1

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/kimhieu153255/first-go/internal/config/db/sqlc"
	"github.com/kimhieu153255/first-go/pkg/token"
	"github.com/kimhieu153255/first-go/pkg/utils"
	"github.com/stretchr/testify/require"
)

type addAuthParams struct {
	userID   int64
	email    string
	role     string
	fullname string
}

func addAuthorization(
	t *testing.T,
	request *http.Request,
	tokenMaker token.Maker,
	authorizationType string,
	authPram addAuthParams,
	duration time.Duration,
) {
	token, payload, err := tokenMaker.CreateToken(authPram.userID, authPram.email, authPram.role, authPram.fullname, duration)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	authorizationHeader := fmt.Sprintf("%s %s", authorizationType, token)
	request.Header.Set(authorizationHeaderKey, authorizationHeader)
}

func TestAuthmiddleware(t *testing.T) {
	user := db.User{
		Email:    utils.RandomString(10) + "@gmail.com",
		Role:     "test",
		FullName: utils.RandomString(10),
		Password: utils.RandomString(10),
	}

	testCases := []struct {
		name          string
		setupAuth     func(t *testing.T, req *http.Request, tokenMaker token.Maker)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			setupAuth: func(t *testing.T, req *http.Request, tokenMaker token.Maker) {
				authPram := addAuthParams{
					email:    user.Email,
					role:     user.Role,
					fullname: user.FullName,
				}
				addAuthorization(t, req, tokenMaker, authorizationTypeBearer, authPram, time.Minute*15)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:      "No Authorization field",
			setupAuth: func(t *testing.T, req *http.Request, tokenMaker token.Maker) {},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "Authorization field less than 2",
			setupAuth: func(t *testing.T, req *http.Request, tokenMaker token.Maker) {
				authPram := addAuthParams{
					email:    user.Email,
					role:     user.Role,
					fullname: user.FullName,
				}
				addAuthorization(t, req, tokenMaker, "", authPram, time.Minute*15)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "Invalid Authorization type",
			setupAuth: func(t *testing.T, req *http.Request, tokenMaker token.Maker) {
				authPram := addAuthParams{
					email:    user.Email,
					role:     user.Role,
					fullname: user.FullName,
				}
				addAuthorization(t, req, tokenMaker, "invalid", authPram, time.Minute*15)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "Invalid token",
			setupAuth: func(t *testing.T, req *http.Request, tokenMaker token.Maker) {
				authPram := addAuthParams{
					email:    user.Email,
					role:     user.Role,
					fullname: user.FullName,
				}
				addAuthorization(t, req, tokenMaker, authorizationTypeBearer, authPram, -time.Minute)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			server := NewTestServer(t, nil)

			// Setup a test endpoint
			authPath := "/auth"
			server.Router.GET(
				authPath,
				AuthMiddleware(server.TokenMaker),
				func(ctx *gin.Context) {
					ctx.JSON(http.StatusOK, gin.H{})
				},
			)

			// Create a new recorder
			recorder := httptest.NewRecorder()
			request, err := http.NewRequest(http.MethodGet, authPath, nil)
			require.NoError(t, err)

			// Setup the request
			tc.setupAuth(t, request, server.TokenMaker)
			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}
