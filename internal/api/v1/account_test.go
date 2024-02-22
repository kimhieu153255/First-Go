package api_v1

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	mockdb "github.com/kimhieu153255/first-go/internal/config/db/mock"
	db "github.com/kimhieu153255/first-go/internal/config/db/sqlc"
	"github.com/kimhieu153255/first-go/pkg/token"
	"github.com/kimhieu153255/first-go/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	setupAuth := func(t *testing.T, req *http.Request, tokenMaker token.Maker) {
		authPram := addAuthParams{
			userID:   utils.RandomInt(1, 1000),
			email:    utils.RandomString(10) + "@gmail.com",
			role:     "test",
			fullname: utils.RandomString(10),
		}
		addAuthorization(t, req, tokenMaker, authorizationTypeBearer, authPram, time.Minute*15)
	}

	body := gin.H{
		"balance":  utils.RandomInt(1, 1000),
		"currency": "USD",
	}

	testCases := []struct {
		name          string
		body          gin.H
		setupAuth     func(t *testing.T, req *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			body:      body,
			setupAuth: setupAuth,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateAccount(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Account{}, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:      "InternalError",
			body:      body,
			setupAuth: setupAuth,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateAccount(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Account{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidCurrency",
			body: gin.H{"balance": utils.RandomInt(1, 1000), "currency": "invalid"},
			setupAuth: func(t *testing.T, req *http.Request, tokenMaker token.Maker) {
				authPram := addAuthParams{
					userID:   utils.RandomInt(1, 1000),
					email:    utils.RandomString(10) + "@gmail.com",
					role:     "test",
					fullname: utils.RandomString(10),
				}
				addAuthorization(t, req, tokenMaker, authorizationTypeBearer, authPram, time.Minute*15)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateAccount(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := NewTestServer(t, store)
			recorder := httptest.NewRecorder()

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/v1/accounts"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			tc.setupAuth(t, request, server.TokenMaker)
			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}
