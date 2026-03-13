package service

import (
	"context"
	"errors"

	"github.com/misbahul-alam/notepad-api/internal/config"
	"github.com/misbahul-alam/notepad-api/internal/db/sqlc"
	"github.com/misbahul-alam/notepad-api/internal/util"
)

type AuthService struct {
	queries   *sqlc.Queries
	jwtConfig *config.JWTConfig
}

func NewAuthService(queries *sqlc.Queries, jwtConfig *config.JWTConfig) *AuthService {
	return &AuthService{queries: queries, jwtConfig: jwtConfig}

}

func (s *AuthService) Register(ctx context.Context, first_name, last_name, email, password string) (sqlc.User, error) {
	hash, err := util.HashPassword(password)
	if err != nil {
		return sqlc.User{}, err
	}

	return s.queries.CreateUser(ctx, sqlc.CreateUserParams{
		FirstName: first_name,
		LastName:  last_name,
		Email:     email,
		Password:  hash,
	})

}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	pass := util.CheckPasswordHash(password, user.Password)
	if !pass {
		return "", errors.New("invalid password")
	}
	token, _ := util.GenerateToken(user.ID, s.jwtConfig.Secret, s.jwtConfig.ExpirationHours)
	return token, nil
}
