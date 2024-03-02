package service

import (
	"context"
	"fmt"
	"time"

	"github.com/Owoade/go-uber/sql"
	"github.com/Owoade/go-uber/utils"
)

type AuthResponse struct {
	Token  string
	UserId int64
}

func (s *UserService) Login(ctx context.Context, email string, password string) (AuthResponse, error) {

	Email := utils.SqlTypeText(email)

	existingUser, err := s.repo.GetUser(ctx, Email)

	if err != nil {
		return *new(AuthResponse), fmt.Errorf(err.Error())
	}

	PASSWORD_IS_INVALID := utils.CompareHashedPassword(password, existingUser.Password.String)

	if PASSWORD_IS_INVALID {
		return *new(AuthResponse), fmt.Errorf("invalid password")
	}

	token, err := s.token.CreateToken(int64(existingUser.ID), time.Hour)

	if err != nil {
		fmt.Println("error creating token", err)
	}

	response := AuthResponse{
		Token:  token,
		UserId: int64(existingUser.ID),
	}

	return response, nil

}

func (s *UserService) SignUp(ctx context.Context, email string, password string) (AuthResponse, error) {

	Email := utils.SqlTypeText(email)

	existingUser, err := s.repo.GetUser(ctx, Email)

	if existingUser != *new(sql.User) {
		return *new(AuthResponse), fmt.Errorf("user with this email already exists")
	}

	if err != nil {
		return *new(AuthResponse), fmt.Errorf(err.Error())
	}

	password, _ = utils.HashPassword(password)

	Password := utils.SqlTypeText(password)

	newUser, err := s.repo.CreateUser(ctx, sql.CreateUserParams{
		Email:    Email,
		Password: Password,
	})

	if err != nil {
		return *new(AuthResponse), fmt.Errorf("error creating user", err.Error())
	}

	token, err := s.token.CreateToken(int64(existingUser.ID), time.Hour)

	if err != nil {
		fmt.Println("error creating token", err)
	}

	response := AuthResponse{
		Token:  token,
		UserId: int64(newUser.ID),
	}

	return response, nil

}

func (s *UserService) AuthorizeUser(ctx context.Context, token string) (int32, error) {

	payload, err := s.token.VerifyToken(token)

	if err != nil {
		return *new(int32), fmt.Errorf("invalid token")
	}

	user, err := s.repo.GetUserById(ctx, int32(payload.UserId))

	if err != nil {
		return *new(int32), fmt.Errorf("Authorization falied")
	}

	return user.ID, nil

}
