package user

import (
	"fmt"
	"time"
	"tns-energo/config"
	"tns-energo/lib/authorize"
	libctx "tns-energo/lib/ctx"
	liblog "tns-energo/lib/log"

	"golang.org/x/crypto/bcrypt"
)

const (
	InspectorRoleId = 1
	AdminRoleId     = 2
)

type Service struct {
	settings config.Settings
	users    UserStorage
}

func NewService(settings config.Settings, users UserStorage) *Service {
	return &Service{
		settings: settings,
		users:    users,
	}
}

func (s *Service) Register(ctx libctx.Context, log liblog.Logger, request RegisterRequest) (AuthResponse, error) {
	tokenTTL := time.Duration(s.settings.Auth.TokenExpiresAfterHours) * time.Hour
	exp := time.Now().Add(tokenTTL)
	refreshToken, err := authorize.NewRefreshToken(tokenTTL, s.settings.Auth.Secret)
	if err != nil {
		return AuthResponse{}, fmt.Errorf("could not create refresh token: %w", err)
	}

	rawPasswordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return AuthResponse{}, fmt.Errorf("could not hash password: %w", err)
	}

	passwordHash := string(rawPasswordHash)

	id, err := s.users.Create(ctx, User{
		Email:                 request.Email,
		Surname:               request.Surname,
		Name:                  request.Name,
		Patronymic:            request.Patronymic,
		Position:              request.Position,
		PasswordHash:          &passwordHash,
		RefreshToken:          &refreshToken,
		RefreshTokenExpiresAt: &exp,
		RoleId:                InspectorRoleId,
	})
	if err != nil {
		return AuthResponse{}, fmt.Errorf("could not create user: %w", err)
	}

	accessToken, err := authorize.NewAccessToken(id, request.Email, false, tokenTTL, s.settings.Auth.Secret)
	if err != nil {
		return AuthResponse{}, fmt.Errorf("could not create access token: %w", err)
	}

	return AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *Service) Login(ctx libctx.Context, log liblog.Logger, request LoginRequest) (AuthResponse, error) {
	user, err := s.users.GetByEmail(ctx, request.Email)
	if err != nil {
		return AuthResponse{}, fmt.Errorf("could not get user by email: %w", err)
	}

	tokenTTL := time.Duration(s.settings.Auth.TokenExpiresAfterHours) * time.Hour
	accessToken, err := authorize.NewAccessToken(user.Id, user.Email, isAdmin(user.RoleId), tokenTTL, s.settings.Auth.Secret)
	if err != nil {
		return AuthResponse{}, fmt.Errorf("could not create access token: %w", err)
	}

	if user.RefreshToken != nil && time.Now().Before(*user.RefreshTokenExpiresAt) {
		return AuthResponse{
			AccessToken:  accessToken,
			RefreshToken: *user.RefreshToken,
		}, nil
	}

	newRefreshToken, err := s.updateToken(ctx, user)
	if err != nil {
		return AuthResponse{}, fmt.Errorf("could not update refresh token: %w", err)
	}

	return AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}, nil
}

func isAdmin(roleId int) bool {
	return roleId == AdminRoleId
}

func (s *Service) RefreshToken(ctx libctx.Context, log liblog.Logger, refreshToken string) (AuthResponse, error) {
	user, err := s.users.GetByRefreshToken(ctx, refreshToken)
	if err != nil {
		return AuthResponse{}, fmt.Errorf("could not get user by refresh token: %w", err)
	}

	if time.Now().After(*user.RefreshTokenExpiresAt) {
		return AuthResponse{}, fmt.Errorf("refresh token expired")
	}

	tokenTTL := time.Duration(s.settings.Auth.TokenExpiresAfterHours) * time.Hour
	newAccessToken, err := authorize.NewAccessToken(user.Id, user.Email, isAdmin(user.RoleId), tokenTTL, s.settings.Auth.Secret)
	if err != nil {
		return AuthResponse{}, fmt.Errorf("could not create access token: %w", err)
	}

	newRefreshToken, err := s.updateToken(ctx, user)
	if err != nil {
		return AuthResponse{}, fmt.Errorf("could not update refresh token: %w", err)
	}

	return AuthResponse{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}, nil
}

func (s *Service) updateToken(ctx libctx.Context, user User) (string, error) {
	tokenTTL := time.Duration(s.settings.Auth.TokenExpiresAfterHours) * time.Hour
	exp := time.Now().Add(tokenTTL)
	newRefreshToken, err := authorize.NewRefreshToken(tokenTTL, s.settings.Auth.Secret)
	if err != nil {
		return "", fmt.Errorf("could not create refresh token: %w", err)
	}

	err = s.users.UpdateRefreshToken(ctx, user.Id, newRefreshToken, exp)
	if err != nil {
		return "", fmt.Errorf("could not update refresh token: %w", err)
	}

	return newRefreshToken, nil
}
