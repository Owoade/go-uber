package token

import (
	"time"

	"github.com/google/uuid"
)

type Payload struct {
	Id        string `json:"id"`
	UserId    int64  `json:"user_id"`
	IssuedAt  string `json:"issued_at"`
	ExpiresAt string `json:"expires_at"`
}

func NewPayload(userId int64, duration time.Duration) (Payload, error) {

	tokenId, err := uuid.NewRandom()

	if err != nil {
		return *new(Payload), err
	}

	payload := Payload{
		Id:        tokenId.String(),
		UserId:    userId,
		ExpiresAt: time.Now().Add(duration).String(),
		IssuedAt:  time.Now().String(),
	}

	return payload, nil

}
