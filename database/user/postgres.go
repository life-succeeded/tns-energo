package user

import (
	_ "embed"
	libctx "tns-energo/lib/ctx"

	"github.com/jmoiron/sqlx"
)

type Postgres struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Postgres {
	return Postgres{
		db: db,
	}
}

//go:embed sql/create.sql
var createSql string

func (r Postgres) CreateUser(ctx libctx.Context, user User) error {
	if _, err := r.db.NamedExecContext(ctx, createSql, user); err != nil {
		return err
	}

	return nil
}
