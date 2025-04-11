package handlers

import (
	"fmt"
	"net/http"
	"tns-energo/lib/http/router"
	"tns-energo/service/inspection"
)

func InspectHandler(inspectionsService *inspection.Service) router.Handler {
	return func(c router.Context) error {
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

type getInspectionsByInspectorIdVars struct {
	InspectorId int `path:"inspector_id"`
}

func GetInspectionsByInspectorId(inspectionService *inspection.Service) router.Handler {
	return func(c router.Context) error {
		var vars getInspectionsByInspectorIdVars
		if err := c.Vars(&vars); err != nil {
			return fmt.Errorf("failed to read vars: %w", err)
		}

		inspections, err := inspectionService.GetByInspectorId(c.Ctx(), c.Log(), vars.InspectorId)
		if err != nil {
			return fmt.Errorf("failed to get inspections: %w", err)
		}

		return c.WriteJson(http.StatusOK, inspections)
	}
}
