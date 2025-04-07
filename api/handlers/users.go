package handlers

import (
	"net/http"
	"tns-energo/lib/http/router"
	"tns-energo/service/user"
)

func RegisterHandler(userService *user.Service) router.Handler {
	return func(c router.Context) error {
		log := c.Log()

		var request user.RegisterRequest
		if err := c.ReadJson(&request); err != nil {
			log.Errorf("failed to read json: %v", err)
			return err
		}

		response, err := userService.Register(c.Ctx(), log, request)
		if err != nil {
			log.Errorf("failed to register user: %v", err)
			return err
		}

		return c.WriteJson(http.StatusOK, response)
	}
}

func LoginHandler(userService *user.Service) router.Handler {
	return func(c router.Context) error {
		log := c.Log()

		var request user.LoginRequest
		if err := c.ReadJson(&request); err != nil {
			log.Errorf("failed to read json: %v", err)
			return err
		}

		response, err := userService.Login(c.Ctx(), log, request)
		if err != nil {
			log.Errorf("failed to login user: %v", err)
			return err
		}

		return c.WriteJson(http.StatusOK, response)
	}
}

type refreshTokenVars struct {
	RefreshToken string `path:"refresh_token"`
}

func RefreshTokenHandler(userService *user.Service) router.Handler {
	return func(c router.Context) error {
		log := c.Log()

		var vars refreshTokenVars
		if err := c.Vars(&vars); err != nil {
			log.Errorf("failed to read vars: %v", err)
			return err
		}

		response, err := userService.RefreshToken(c.Ctx(), log, vars.RefreshToken)
		if err != nil {
			log.Errorf("failed to refresh token: %v", err)
			return err
		}

		return c.WriteJson(http.StatusOK, response)
	}
}

type getUserByIdVars struct {
	UserId int `path:"user_id"`
}

func GetUserByIdHandler(userService *user.Service) router.Handler {
	return func(c router.Context) error {
		log := c.Log()

		var vars getUserByIdVars
		if err := c.Vars(&vars); err != nil {
			log.Errorf("failed to read vars: %v", err)
			return err
		}

		response, err := userService.GetById(c.Ctx(), vars.UserId)
		if err != nil {
			log.Errorf("failed to retrieve user: %v", err)
			return err
		}

		return c.WriteJson(http.StatusOK, response)
	}
}
