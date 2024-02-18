package service

import (
	"github.com/Owoade/go-uber/sql"
	"github.com/Owoade/go-uber/token"
)

type DriverService struct {
	repo  *sql.Queries
	token *token.PacetoMaker
}

func NewDriverService(q *sql.Queries, t *token.PacetoMaker) DriverService {
	return DriverService{
		repo:  q,
		token: t,
	}
}
