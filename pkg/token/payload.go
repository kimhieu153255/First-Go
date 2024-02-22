package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	errExpiredToken = errors.New("token has expired")
	errInvalidToken = errors.New("token is invalid")
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	UserID    int64     `json:"user_id"`
	Email     string    `json:"email"`
	Fullname  string    `json:"fullname"`
	Role      string    `json:"role"`
	Expire_at time.Time `json:"expired_at"`
}

func NewPayload(userID int64, email, fullname, role string, duration time.Duration) (*Payload, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenId,
		UserID:    userID,
		Email:     email,
		Fullname:  fullname,
		Role:      role,
		Expire_at: time.Now().Add(duration),
	}
	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.Expire_at) {
		return errExpiredToken
	}
	return nil
}
