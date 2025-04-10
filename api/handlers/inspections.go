package handlers

import (
	"net/http"
	"tns-energo/lib/http/router"
	"tns-energo/service/inspection"
)

func InspectHandler(inspectionsService *inspection.Service) router.Handler {
	return func(c router.Context) error {
		log := c.Log()

		var request inspection.InspectRequest
		if err := c.ReadJson(&request); err != nil {
			log.Errorf("failed to read json: %v", err)
			return err
		}

		response, err := inspectionsService.Inspect(c.Ctx(), log, request)
		if err != nil {
			log.Errorf("failed to inspect: %v", err)
			return err
		}

		return c.WriteJson(http.StatusOK, response)
	}
}

type getInspectionsByInspectorIdVars struct {
	InspectorId int `path:"inspector_id"`
}

func GetInspectionsByInspectorId(inspectionService *inspection.Service) router.Handler {
	return func(c router.Context) error {
		log := c.Log()

		var vars getInspectionsByInspectorIdVars
		if err := c.Vars(&vars); err != nil {
			log.Errorf("failed to read vars: %v", err)
			return err
		}

		inspections, err := inspectionService.GetByInspectorId(c.Ctx(), log, vars.InspectorId)
		if err != nil {
			log.Errorf("failed to get inspections: %v", err)
			return err
		}

		return c.WriteJson(http.StatusOK, inspections)
	}
}
