package api_v1

import (
	"testing"

	"github.com/gin-gonic/gin"
	db "github.com/kimhieu153255/first-go/internal/config/db/sqlc"
	config_env "github.com/kimhieu153255/first-go/internal/config/env"
	"github.com/stretchr/testify/require"
)

func NewTestServer(t *testing.T, store db.Store) *Server {

	config, err := config_env.NewConfig("../../../")
	require.NoError(t, err)

	config.AccessTokenDuration = 60
	config.RefreshTokenDuration = 60

	server, err := NewServer(store, config)
	require.NoError(t, err)
	return server
}

func TestMain(t *testing.M) {
	gin.SetMode(gin.TestMode)
	t.Run()
}
