package token

import (
	"task/utils"
	"time"

	"github.com/google/uuid"
)

// payload
type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	pauload := &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return pauload, err
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return utils.ErrExpiredToken
	}
	return nil
}
