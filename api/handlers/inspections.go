package handlers

import (
	"net/http"
	"tns-energo/lib/http/router"
	"tns-energo/service/inspection"
)

func InspectHandler(inspectionsService *inspection.Service) router.Handler {
	return func(c router.Context) error {
		log := c.Log()
		response, err := inspectionsService.Inspect(c.Ctx(), log)
		if err != nil {
			log.Errorf("failed to inspect: %v", err)
			return err
		}

		return c.WriteJson(http.StatusOK, response)
	}
}
