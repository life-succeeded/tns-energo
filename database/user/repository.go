package user

import (
	"time"
	libctx "tns-energo/lib/ctx"
)

type Repository interface {
	Create(ctx libctx.Context, user User) (int, error)
	GetByEmail(ctx libctx.Context, email string) (User, error)
	GetByRefreshToken(ctx libctx.Context, refreshToken string) (User, error)
	UpdateRefreshToken(ctx libctx.Context, userId int, newRefreshToken string, newRefreshTokenExpiresAt time.Time) error
}
