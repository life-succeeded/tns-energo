package handlers

import (
	"fmt"
	"net/http"
	"tns-energo/lib/http/router"
	"tns-energo/service/brigade"
)

func CreateBrigadeHandler(brigadeService *brigade.Service) router.Handler {
	return func(c router.Context) error {
		var request brigade.CreateRequest
		if err := c.ReadJson(&request); err != nil {
			return fmt.Errorf("failed to read json: %w", err)
		}

		response, err := brigadeService.Create(c.Ctx(), c.Log(), request)
		if err != nil {
			return fmt.Errorf("failed to create brigade: %w", err)
		}

		return c.WriteJson(http.StatusOK, response)
	}
}

type getBrigadeByIdVars struct {
	BrigadeId string `path:"brigade_id"`
}

func GetBrigadeByIdHandler(brigadeService *brigade.Service) router.Handler {
	return func(c router.Context) error {
		var vars getBrigadeByIdVars
		if err := c.Vars(&vars); err != nil {
			return fmt.Errorf("failed to get vars: %w", err)
		}

		response, err := brigadeService.GetById(c.Ctx(), c.Log(), vars.BrigadeId)
		if err != nil {
			return fmt.Errorf("failed to get brigade: %w", err)
		}

		return c.WriteJson(http.StatusOK, response)
	}
}
