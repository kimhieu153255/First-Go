package api_v1

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
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

func TestCreateUserApi(t *testing.T) {
	createUserParams := db.CreateUserParams{
		Password: utils.RandomString(10),
		FullName: utils.RandomString(10),
		Email:    utils.RandomString(10) + "@gmail.com",
	}

	setupAuth := func(t *testing.T, req *http.Request, tokenMaker token.Maker) {
		authPram := addAuthParams{
			email:    createUserParams.Email,
			fullname: createUserParams.FullName,
		}
		addAuthorization(t, req, tokenMaker, authorizationTypeBearer, authPram, time.Minute*15)
	}

	testCases := []struct {
		name          string
		body          gin.H
		setupAuth     func(t *testing.T, req *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"password":  createUserParams.Password,
				"full_name": createUserParams.FullName,
				"email":     createUserParams.Email,
			},
			setupAuth: setupAuth,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.User{}, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "DuplicateEmail",
			body: gin.H{
				"password":  createUserParams.Password,
				"full_name": createUserParams.FullName,
				"email":     createUserParams.Email,
			},
			setupAuth: setupAuth,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.User{}, db.ErrUniqueViolation)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusForbidden, recorder.Code)
			},
		},
		{
			name: "badRequest",
			body: gin.H{
				"password":  createUserParams.Password,
				"full_name": createUserParams.FullName,
				"emai":      createUserParams.Email,
			},
			setupAuth: setupAuth,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InternalServiceError",
			body: gin.H{
				"password":  createUserParams.Password,
				"full_name": createUserParams.FullName,
				"email":     createUserParams.Email,
			},
			setupAuth: setupAuth,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.User{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := NewTestServer(t, store)
			recorder := httptest.NewRecorder()

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/v1/users"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			tc.setupAuth(t, request, server.TokenMaker)
			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func TestGetUserById(t *testing.T) {
	setupAuth := func(t *testing.T, req *http.Request, tokenMaker token.Maker) {
		authPram := addAuthParams{
			email:    "testGet",
			fullname: "testGet",
		}
		addAuthorization(t, req, tokenMaker, authorizationTypeBearer, authPram, time.Minute*15)
	}
	testCases := []struct {
		name          string
		userID        any
		setupAuth     func(t *testing.T, req *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			userID:    int64(1),
			setupAuth: setupAuth,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUserById(gomock.Any(), gomock.Eq(int64(1))).
					Times(1).
					Return(db.User{}, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:      "ID is not a number",
			userID:    "a",
			setupAuth: setupAuth,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUserById(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:      "Record not found",
			userID:    int64(-1),
			setupAuth: setupAuth,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUserById(gomock.Any(), gomock.Eq(int64(-1))).
					Times(1).
					Return(db.User{}, db.ErrRecordNotFound)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:      "InternalServiceError",
			userID:    int64(1),
			setupAuth: setupAuth,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUserById(gomock.Any(), gomock.Eq(int64(1))).
					Times(1).
					Return(db.User{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := NewTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := "/v1/users/" + fmt.Sprintf("%v", tc.userID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.TokenMaker)

			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func TestGetListUser(t *testing.T) {
	setupAuth := func(t *testing.T, req *http.Request, tokenMaker token.Maker) {
		authPram := addAuthParams{
			email:    "testGet",
			fullname: "testGet",
		}
		addAuthorization(t, req, tokenMaker, authorizationTypeBearer, authPram, time.Minute*15)
	}
	testCases := []struct {
		name          string
		setupAuth     func(t *testing.T, req *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			setupAuth: setupAuth,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetListUsers(gomock.Any()).
					Times(1).
					Return([]db.User{}, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:      "InternalServiceError",
			setupAuth: setupAuth,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetListUsers(gomock.Any()).
					Times(1).
					Return([]db.User{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := NewTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := "/v1/users"
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.TokenMaker)

			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func TestDeleteUserById(t *testing.T) {
	setupAuth := func(t *testing.T, req *http.Request, tokenMaker token.Maker) {
		authPram := addAuthParams{
			email:    "testGet",
			fullname: "testGet",
		}
		addAuthorization(t, req, tokenMaker, authorizationTypeBearer, authPram, time.Minute*15)
	}

	testCases := []struct {
		name          string
		userID        any
		setupAuth     func(t *testing.T, req *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			userID:    int64(1),
			setupAuth: setupAuth,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeleteUserByID(gomock.Any(), gomock.Eq(int64(1))).
					Times(1).
					Return(db.User{}, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:      "ID not valid",
			userID:    "a",
			setupAuth: setupAuth,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeleteUserByID(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:      "InternalServiceError",
			userID:    int64(1),
			setupAuth: setupAuth,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeleteUserByID(gomock.Any(), gomock.Eq(int64(1))).
					Times(1).
					Return(db.User{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:      "Record not found",
			userID:    int64(1),
			setupAuth: setupAuth,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeleteUserByID(gomock.Any(), gomock.Eq(int64(1))).
					Times(1).
					Return(db.User{}, db.ErrRecordNotFound)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := NewTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := "/v1/users/" + fmt.Sprintf("%v", tc.userID)
			request, err := http.NewRequest(http.MethodDelete, url, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.TokenMaker)

			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}
