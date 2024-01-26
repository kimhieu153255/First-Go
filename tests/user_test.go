package tests

import (
	"context"
	"testing"

	db "github.com/kimhieu/first-go/internal/config/db/sqlc"
	"github.com/kimhieu/first-go/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	arg := db.CreateUserParams{
		Email:    "test@gmail.com",
		FullName: utils.RandomString(10),
		Password: utils.RandomString(10),
	}

	user, err := testStore.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Password, user.Password)
	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
	require.WithinDuration(t, user.CreatedAt, user.CreatedAt, 2)
}
