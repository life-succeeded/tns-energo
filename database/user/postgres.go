package user

import (
	_ "embed"
	"time"
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

func (r Postgres) Create(ctx libctx.Context, user User) (id int, err error) {
	rows, err := r.db.NamedQueryContext(ctx, createSql, user)
	if err != nil {
		return 0, err
	}

	defer func(rows *sqlx.Rows) {
		if tempErr := rows.Close(); tempErr != nil {
			err = tempErr
		}
	}(rows)

	if rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return 0, err
		}
	}

	return id, nil
}

//go:embed sql/get_by_email.sql
var getByEmailSql string

func (r Postgres) GetByEmail(ctx libctx.Context, email string) (User, error) {
	var user User
	err := r.db.GetContext(ctx, &user, getByEmailSql, email)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

//go:embed sql/get_by_refresh_token.sql
var getByRefreshTokenSql string

func (r Postgres) GetByRefreshToken(ctx libctx.Context, refreshToken string) (User, error) {
	var user User
	err := r.db.GetContext(ctx, &user, getByRefreshTokenSql, refreshToken)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

//go:embed sql/update_refresh_token.sql
var updateRefreshTokenSql string

func (r Postgres) UpdateRefreshToken(ctx libctx.Context, userId int, newRefreshToken string, newRefreshTokenExpiresAt time.Time) error {
	_, err := r.db.ExecContext(ctx, updateRefreshTokenSql, userId, newRefreshToken, newRefreshTokenExpiresAt)

	return err
}
