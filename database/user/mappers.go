package user

import (
	domain "tns-energo/service/user"
)

func MapToDb(u domain.User) User {
	return User{
		Id:                    u.Id,
		Email:                 u.Email,
		Surname:               u.Surname,
		Name:                  u.Name,
		Patronymic:            u.Patronymic,
		Position:              u.Position,
		PasswordHash:          u.PasswordHash,
		RefreshToken:          u.RefreshToken,
		RefreshTokenExpiresAt: u.RefreshTokenExpiresAt,
		RoleId:                u.RoleId,
	}
}

func MapToDomain(u User) domain.User {
	return domain.User{
		Id:                    u.Id,
		Email:                 u.Email,
		Surname:               u.Surname,
		Name:                  u.Name,
		Patronymic:            u.Patronymic,
		Position:              u.Position,
		PasswordHash:          u.PasswordHash,
		RefreshToken:          u.RefreshToken,
		RefreshTokenExpiresAt: u.RefreshTokenExpiresAt,
		RoleId:                u.RoleId,
	}
}

func MapToDomainLight(u User) domain.UserLight {
	return domain.UserLight{
		Id:         u.Id,
		Email:      u.Email,
		Surname:    u.Surname,
		Name:       u.Name,
		Patronymic: u.Patronymic,
		Position:   u.Position,
	}
}
