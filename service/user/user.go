package user

import (
	"tns-energo/database/user"
	libctx "tns-energo/lib/ctx"
	liblog "tns-energo/lib/log"
)

type Service interface {
	Register(ctx libctx.Context, log liblog.Logger, request RegisterRequest) (AuthResponse, error)
	Login(ctx libctx.Context, log liblog.Logger, request LoginRequest) (AuthResponse, error)
	RefreshToken(ctx libctx.Context, log liblog.Logger, refreshToken string) (AuthResponse, error)
}

type Impl struct {
	users user.Repository
}

func NewService(users user.Repository) *Impl {
	return &Impl{
		users: users,
	}
}

func (s *Impl) Register(ctx libctx.Context, log liblog.Logger, request RegisterRequest) (AuthResponse, error) {
	return AuthResponse{}, nil
}

func (s *Impl) Login(ctx libctx.Context, log liblog.Logger, request LoginRequest) (AuthResponse, error) {
	return AuthResponse{}, nil
}

func (s *Impl) RefreshToken(ctx libctx.Context, log liblog.Logger, refreshToken string) (AuthResponse, error) {
	return AuthResponse{}, nil
}
