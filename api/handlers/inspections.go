package handlers

import (
	"fmt"
	"net/http"
	"tns-energo/service/inspection"

	"github.com/sunshineOfficial/golib/gohttp/gorouter"
)

func InspectHandler(inspectionsService *inspection.Service) gorouter.Handler {
	return func(c gorouter.Context) error {
		var request inspection.InspectRequest
		if err := c.ReadJson(&request); err != nil {
			return fmt.Errorf("failed to read json: %w", err)
		}

		response, err := inspectionsService.Inspect(c.Ctx(), c.Log(), request)
		if err != nil {
			return fmt.Errorf("failed to inspect: %w", err)
		}

		return c.WriteJson(http.StatusOK, response)
	}
}

type getInspectionsByBrigadeIdVars struct {
	BrigadeId string `path:"brigade_id"`
}

func GetInspectionsByBrigadeId(inspectionService *inspection.Service) gorouter.Handler {
	return func(c gorouter.Context) error {
		var vars getInspectionsByBrigadeIdVars
		if err := c.Vars(&vars); err != nil {
			return fmt.Errorf("failed to read vars: %w", err)
		}

		inspections, err := inspectionService.GetByBrigadeId(c.Ctx(), c.Log(), vars.BrigadeId)
		if err != nil {
			return fmt.Errorf("failed to get inspections: %w", err)
		}

		return c.WriteJson(http.StatusOK, inspections)
	}
}
