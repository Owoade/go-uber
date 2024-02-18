package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

type PacetoMaker struct {
	paseto      *paseto.V2
	symetricKey []byte
}

func NewPasetomaker(key string) (*PacetoMaker, error) {

	FUNC_ZERO_VALUE := *new(*PacetoMaker)

	if len(key) != chacha20poly1305.KeySize {
		return FUNC_ZERO_VALUE, fmt.Errorf("invalid key size")
	}

	maker := &PacetoMaker{
		paseto:      paseto.NewV2(),
		symetricKey: []byte(key),
	}

	fmt.Println(maker)

	return maker, nil
}

func (p *PacetoMaker) CreateToken(id int64, duration time.Duration) (string, error) {

	payload, err := NewPayload(id, duration)

	if err != nil {
		fmt.Println("error creating token payload")
		return "", err
	}

	return p.paseto.Encrypt(p.symetricKey, payload, nil)

}

func (p *PacetoMaker) VerifyToken(token string) (*Payload, error) {

	payload := &Payload{}

	err := p.paseto.Decrypt(token, p.symetricKey, payload, nil)

	if err != nil {
		return nil, fmt.Errorf("invalid token")
	}

	return payload, nil

}
