package service

import (
	"context"
	"fmt"
	"time"

	"github.com/Owoade/go-uber/sql"
	"github.com/Owoade/go-uber/utils"
	"github.com/jackc/pgx/v5/pgtype"
)

type AuthResponse struct {
	Token  string
	UserId int64
}

func (s *DriverService) Login(ctx context.Context, email string, password string) (AuthResponse, error) {

	Email := pgtype.Text{
		String: email,
		Valid:  true,
	}

	existingDriver, err := s.repo.GetDriver(ctx, Email)

	if err != nil {
		return *new(AuthResponse), fmt.Errorf(err.Error())
	}

	PASSWORD_IS_INVALID := utils.CompareHashedPassword(password, existingDriver.Password.String)

	if PASSWORD_IS_INVALID {
		return *new(AuthResponse), fmt.Errorf("invalid password")
	}

	token, err := s.token.CreateToken(int64(existingDriver.ID), time.Hour)

	if err != nil {
		fmt.Println("error creating token", err)
	}

	response := AuthResponse{
		Token:  token,
		UserId: int64(existingDriver.ID),
	}

	return response, nil

}

func (s *DriverService) SignUp(ctx context.Context, email string, password string) (AuthResponse, error) {

	Email := pgtype.Text{
		String: email,
		Valid:  true,
	}

	existingDriver, err := s.repo.GetDriver(ctx, Email)

	if existingDriver != *new(sql.Driver) {
		return *new(AuthResponse), fmt.Errorf("user with this email already exists")
	}

	if err != nil {
		return *new(AuthResponse), fmt.Errorf(err.Error())
	}

	password, _ = utils.HashPassword(password)

	Password := pgtype.Text{
		String: password,
		Valid:  true,
	}

	newDriver, err := s.repo.CreateUser(ctx, sql.CreateUserParams{
		Email:    Email,
		Password: Password,
	})

	if err != nil {
		return *new(AuthResponse), fmt.Errorf("error creating user", err.Error())
	}

	token, err := s.token.CreateToken(int64(existingDriver.ID), time.Hour)

	if err != nil {
		fmt.Println("error creating token", err)
	}

	response := AuthResponse{
		Token:  token,
		UserId: int64(newDriver.ID),
	}

	return response, nil

}

func (s *DriverService) AuthorizeUser(ctx context.Context, token string) (int32, error) {

	payload, err := s.token.VerifyToken(token)

	if err != nil {
		return *new(int32), fmt.Errorf("invalid token")
	}

	user, err := s.repo.GetDriverById(ctx, int32(payload.UserId))

	if err != nil {
		return *new(int32), fmt.Errorf("Authorization falied")
	}

	return user.ID, nil
}
