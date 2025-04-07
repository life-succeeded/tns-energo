package user

import "time"

type User struct {
	Id                    int        `json:"id"`
	Email                 string     `json:"email"`
	Surname               string     `json:"surname"`
	Name                  string     `json:"name"`
	Patronymic            *string    `json:"patronymic"`
	Position              *string    `json:"position"`
	PasswordHash          *string    `json:"password_hash"`
	RefreshToken          *string    `json:"refresh_token"`
	RefreshTokenExpiresAt *time.Time `json:"refresh_token_expires_at"`
	RoleId                int        `json:"role_id"`
}

type Role struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type RegisterRequest struct {
	Email      string  `json:"email"`
	Password   string  `json:"password"`
	Surname    string  `json:"surname"`
	Name       string  `json:"name"`
	Patronymic *string `json:"patronymic"`
	Position   *string `json:"position"`
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLight struct {
	Id         int     `json:"id"`
	Email      string  `json:"email"`
	Surname    string  `json:"surname"`
	Name       string  `json:"name"`
	Patronymic *string `json:"patronymic,omitempty"`
	Position   *string `json:"position,omitempty"`
}
