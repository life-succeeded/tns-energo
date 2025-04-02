package user

import (
	"fmt"
	"time"
	"tns-energo/config"
	"tns-energo/database/user"
	"tns-energo/lib/authorize"
	libctx "tns-energo/lib/ctx"
	liblog "tns-energo/lib/log"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(ctx libctx.Context, log liblog.Logger, request RegisterRequest) (AuthResponse, error)
	Login(ctx libctx.Context, log liblog.Logger, request LoginRequest) (AuthResponse, error)
	RefreshToken(ctx libctx.Context, log liblog.Logger, refreshToken string) (AuthResponse, error)
}

const (
	InspectorRoleId = 1
	AdminRoleId     = 2
)

type Impl struct {
	users    user.Repository
	settings config.Settings
}

func NewService(users user.Repository, settings config.Settings) *Impl {
	return &Impl{
		users:    users,
		settings: settings,
	}
}

func (s *Impl) Register(ctx libctx.Context, log liblog.Logger, request RegisterRequest) (AuthResponse, error) {
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

	id, err := s.users.Create(ctx, user.User{
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

func (s *Impl) Login(ctx libctx.Context, log liblog.Logger, request LoginRequest) (AuthResponse, error) {
	dbUser, err := s.users.GetByEmail(ctx, request.Email)
	if err != nil {
		return AuthResponse{}, fmt.Errorf("could not get user by email: %w", err)
	}

	tokenTTL := time.Duration(s.settings.Auth.TokenExpiresAfterHours) * time.Hour
	accessToken, err := authorize.NewAccessToken(dbUser.Id, dbUser.Email, isAdmin(dbUser.RoleId), tokenTTL, s.settings.Auth.Secret)
	if err != nil {
		return AuthResponse{}, fmt.Errorf("could not create access token: %w", err)
	}

	if dbUser.RefreshToken != nil && time.Now().Before(*dbUser.RefreshTokenExpiresAt) {
		return AuthResponse{
			AccessToken:  accessToken,
			RefreshToken: *dbUser.RefreshToken,
		}, nil
	}

	newRefreshToken, err := s.updateToken(ctx, dbUser)
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

func (s *Impl) RefreshToken(ctx libctx.Context, log liblog.Logger, refreshToken string) (AuthResponse, error) {
	dbUser, err := s.users.GetByRefreshToken(ctx, refreshToken)
	if err != nil {
		return AuthResponse{}, fmt.Errorf("could not get user by refresh token: %w", err)
	}

	if time.Now().After(*dbUser.RefreshTokenExpiresAt) {
		return AuthResponse{}, fmt.Errorf("refresh token expired")
	}

	tokenTTL := time.Duration(s.settings.Auth.TokenExpiresAfterHours) * time.Hour
	newAccessToken, err := authorize.NewAccessToken(dbUser.Id, dbUser.Email, isAdmin(dbUser.RoleId), tokenTTL, s.settings.Auth.Secret)
	if err != nil {
		return AuthResponse{}, fmt.Errorf("could not create access token: %w", err)
	}

	newRefreshToken, err := s.updateToken(ctx, dbUser)
	if err != nil {
		return AuthResponse{}, fmt.Errorf("could not update refresh token: %w", err)
	}

	return AuthResponse{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}, nil
}

func (s *Impl) updateToken(ctx libctx.Context, dbUser user.User) (string, error) {
	tokenTTL := time.Duration(s.settings.Auth.TokenExpiresAfterHours) * time.Hour
	exp := time.Now().Add(tokenTTL)
	newRefreshToken, err := authorize.NewRefreshToken(tokenTTL, s.settings.Auth.Secret)
	if err != nil {
		return "", fmt.Errorf("could not create refresh token: %w", err)
	}

	err = s.users.UpdateRefreshToken(ctx, dbUser.Id, newRefreshToken, exp)
	if err != nil {
		return "", fmt.Errorf("could not update refresh token: %w", err)
	}

	return newRefreshToken, nil
}
