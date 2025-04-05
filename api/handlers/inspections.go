package handlers

import (
	"net/http"
	"tns-energo/lib/http/router"
	"tns-energo/service/inspection"
)

func InspectHandler(inspectionsService inspection.Service) router.Handler {
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

func RegistryHandler(inspectionsService inspection.Service) router.Handler {
	return func(c router.Context) error {
		log := c.Log()

		bytes, err := c.ReadBytes()
		if err != nil {
			log.Errorf("failed to read: %v", err)
			return err
		}

		err = inspectionsService.ParseExcelRegistry(c.Ctx(), log, bytes)
		if err != nil {
			log.Errorf("failed to parse: %v", err)
			return err
		}

		c.Write(http.StatusOK)
		return nil
	}
}
