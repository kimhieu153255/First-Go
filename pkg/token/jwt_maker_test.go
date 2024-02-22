package token

import (
	"testing"
	"time"

	"github.com/kimhieu153255/first-go/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	secretStr := utils.RandomString(32)
	newMaker, err := NewJWTMaker(secretStr)
	require.NoError(t, err)

	userID := utils.RandomInt(1, 1000)
	email := utils.RandomString(20) + "@gmail.com"
	fullname := utils.RandomString(20)
	role := "test"
	duration := time.Duration(time.Minute * 15)

	token, payload, err := newMaker.CreateToken(userID, email, role, fullname, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotZero(t, payload)

	// Verify the token
	payload, err = newMaker.VerifyToken(token)
	require.NoError(t, err)
	require.NotZero(t, payload)
}
