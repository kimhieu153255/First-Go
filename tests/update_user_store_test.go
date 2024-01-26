package tests

import (
	"context"
	"testing"

	db "github.com/kimhieu/first-go/internal/config/db/sqlc"
	"github.com/kimhieu/first-go/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestUpdateUserUseStore(t *testing.T) {
	arg := db.UpdateUserTxParams{
		Email:    "test@gmail.com",
		FullName: utils.RandomString(10),
		Password: utils.RandomString(10),
	}
	result, err := testStore.UpdateUserUseStore(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, arg.Email, result.User.Email)
	require.Equal(t, arg.FullName, result.User.FullName)
	require.Equal(t, arg.Password, result.User.Password)
}
