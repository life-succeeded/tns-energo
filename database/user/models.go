package user

import "time"

type User struct {
	Id                    int        `db:"id"`
	Email                 string     `db:"email"`
	Surname               string     `db:"surname"`
	Name                  string     `db:"name"`
	Patronymic            *string    `db:"patronymic"`
	Position              *string    `db:"position"`
	RefreshToken          *string    `db:"refresh_token"`
	RefreshTokenExpiresAt *time.Time `db:"refresh_token_expires_at"`
	RoleId                int        `db:"role_id"`
}

type Role struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}
