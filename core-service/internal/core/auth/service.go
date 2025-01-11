package auth

import (
	"context"
	"gym-core-service/internal/core/user"
)

type AuthService struct {
	ctx        context.Context
	repository AuthRepository
}

func NewAuthService(ctx context.Context, repository AuthRepository) *AuthService {
	return &AuthService{ctx, repository}
}

func (s *AuthService) LoginUser(email string, password string) (*user.UserEntity, error) {
	user, err := s.repository.GetUserByEmailAndPassword(s.ctx, email, password)
	return &user, err
}

func (s *AuthService) LoginAdminUser(email string, password string) (*user.UserEntity, error) {
	user, err := s.repository.GetAdminUserByEmailAndPassword(s.ctx, email, password)
	return &user, err
}
