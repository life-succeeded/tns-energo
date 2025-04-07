package user

import "tns-energo/service/user"

func mapToDb(user user.User) User {
	return User{
		Id:                    user.Id,
		Email:                 user.Email,
		Surname:               user.Surname,
		Name:                  user.Name,
		Patronymic:            user.Patronymic,
		Position:              user.Position,
		PasswordHash:          user.PasswordHash,
		RefreshToken:          user.RefreshToken,
		RefreshTokenExpiresAt: user.RefreshTokenExpiresAt,
		RoleId:                user.RoleId,
	}
}

func mapToDomain(dbUser User) user.User {
	return user.User{
		Id:                    dbUser.Id,
		Email:                 dbUser.Email,
		Surname:               dbUser.Surname,
		Name:                  dbUser.Name,
		Patronymic:            dbUser.Patronymic,
		Position:              dbUser.Position,
		PasswordHash:          dbUser.PasswordHash,
		RefreshToken:          dbUser.RefreshToken,
		RefreshTokenExpiresAt: dbUser.RefreshTokenExpiresAt,
		RoleId:                dbUser.RoleId,
	}
}
