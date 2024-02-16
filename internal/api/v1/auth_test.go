package api_v1

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	mockdb "github.com/kimhieu153255/first-go/internal/config/db/mock"
	db "github.com/kimhieu153255/first-go/internal/config/db/sqlc"
	"github.com/kimhieu153255/first-go/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestLoginApi(t *testing.T) {
	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"email":    "kimhieu@gmail.com",
				"password": "123456",
			},
			buildStubs: func(store *mockdb.MockStore) {
				hashPassword, _ := utils.HashPassword("123456")
				store.EXPECT().
					GetUserByEmail(gomock.Any(), gomock.Eq("kimhieu@gmail.com")).
					Times(1).
					Return(db.User{Password: hashPassword}, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
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

			url := "/v1/auth/login"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}

}
