package service

import (
	"github.com/Owoade/go-uber/sql"
	"github.com/Owoade/go-uber/token"
)

type UserService struct {
	repo  *sql.Queries
	token *token.PacetoMaker
}

func NewUserService(q *sql.Queries, t *token.PacetoMaker) UserService {
	return UserService{
		repo:  q,
		token: t,
	}
}
