package handlers

import (
	"net/http"
	"tns-energo/lib/http/router"
	"tns-energo/service/registry"
)

func ParseRegistryHandler(registryService registry.Service) router.Handler {
	return func(c router.Context) error {
		log := c.Log()

		bytes, err := c.ReadBytes()
		if err != nil {
			log.Errorf("failed to read: %v", err)
			return err
		}

		err = registryService.Parse(c.Ctx(), log, bytes)
		if err != nil {
			log.Errorf("failed to parse: %v", err)
			return err
		}

		c.Write(http.StatusOK)
		return nil
	}
}

type GetItemByAccountNumberVars struct {
	AccountNumber string `path:"account_number"`
}

func GetItemByAccountNumberHandler(registryService registry.Service) router.Handler {
	return func(c router.Context) error {
		log := c.Log()

		var vars GetItemByAccountNumberVars
		if err := c.Vars(&vars); err != nil {
			log.Errorf("failed to read vars: %v", err)
			return err
		}

		item, err := registryService.GetItemByAccountNumber(c.Ctx(), log, vars.AccountNumber)
		if err != nil {
			log.Errorf("failed to get item by account number: %v", err)
			return err
		}

		return c.WriteJson(http.StatusOK, item)
	}
}
