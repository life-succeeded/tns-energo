package handlers

import (
	"net/http"
	"tns-energo/lib/http/router"
	"tns-energo/service/registry"
)

func ParseRegistryHandler(registryService *registry.Service) router.Handler {
	return func(c router.Context) error {
		log := c.Log()

		payload, err := c.ReadBytes()
		if err != nil {
			log.Errorf("failed to read: %v", err)
			return err
		}

		err = registryService.Parse(c.Ctx(), log, payload)
		if err != nil {
			log.Errorf("failed to parse: %v", err)
			return err
		}

		c.Write(http.StatusOK)
		return nil
	}
}

type getItemByAccountNumberVars struct {
	AccountNumber string `path:"account_number"`
}

func GetItemByAccountNumberHandler(registryService *registry.Service) router.Handler {
	return func(c router.Context) error {
		log := c.Log()

		var vars getItemByAccountNumberVars
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

func GetItemsByAccountNumberRegularHandler(registryService *registry.Service) router.Handler {
	return func(c router.Context) error {
		log := c.Log()

		var vars getItemByAccountNumberVars
		if err := c.Vars(&vars); err != nil {
			log.Errorf("failed to read vars: %v", err)
			return err
		}

		items, err := registryService.GetItemsByAccountNumberRegular(c.Ctx(), log, vars.AccountNumber)
		if err != nil {
			log.Errorf("failed to get items by account number: %v", err)
			return err
		}

		return c.WriteJson(http.StatusOK, items)
	}
}
