package user

import libctx "tns-energo/lib/ctx"

type Repository interface {
	CreateUser(ctx libctx.Context, user User) error
}
