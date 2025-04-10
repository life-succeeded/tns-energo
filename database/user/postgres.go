package user

import (
	_ "embed"
	"time"
	libctx "tns-energo/lib/ctx"
	domain "tns-energo/service/user"

	"github.com/jmoiron/sqlx"
)

type Postgres struct {
	db *sqlx.DB
}

func NewStorage(db *sqlx.DB) *Postgres {
	return &Postgres{
		db: db,
	}
}

//go:embed sql/create.sql
var createSql string

func (s *Postgres) Create(ctx libctx.Context, user domain.User) (id int, err error) {
	rows, err := s.db.NamedQueryContext(ctx, createSql, MapToDb(user))
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

func (s *Postgres) GetByEmail(ctx libctx.Context, email string) (domain.User, error) {
	var user User
	err := s.db.GetContext(ctx, &user, getByEmailSql, email)
	if err != nil {
		return domain.User{}, err
	}

	return MapToDomain(user), nil
}

//go:embed sql/get_by_refresh_token.sql
var getByRefreshTokenSql string

func (s *Postgres) GetByRefreshToken(ctx libctx.Context, refreshToken string) (domain.User, error) {
	var user User
	err := s.db.GetContext(ctx, &user, getByRefreshTokenSql, refreshToken)
	if err != nil {
		return domain.User{}, err
	}

	return MapToDomain(user), nil
}

//go:embed sql/update_refresh_token.sql
var updateRefreshTokenSql string

func (s *Postgres) UpdateRefreshToken(ctx libctx.Context, userId int, newRefreshToken string, newRefreshTokenExpiresAt time.Time) error {
	_, err := s.db.ExecContext(ctx, updateRefreshTokenSql, userId, newRefreshToken, newRefreshTokenExpiresAt)

	return err
}

//go:embed sql/get_light_by_id.sql
var getLightByIdSql string

func (s *Postgres) GetLightById(ctx libctx.Context, userId int) (domain.UserLight, error) {
	var user User
	err := s.db.GetContext(ctx, &user, getLightByIdSql, userId)
	if err != nil {
		return domain.UserLight{}, err
	}

	return MapToDomainLight(user), nil
}
