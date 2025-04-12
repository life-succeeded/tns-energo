package handlers

import (
	"fmt"
	"net/http"
	"tns-energo/lib/http/router"
	"tns-energo/service/registry"
)

func ParseRegistryHandler(registryService *registry.Service) router.Handler {
	return func(c router.Context) error {
		form, err := c.FormData()
		if err != nil {
			return fmt.Errorf("could not parse form data: %w", err)
		}

		if len(form.File["payload"]) != 1 {
			return fmt.Errorf("invalid form data")
		}

		err = registryService.Parse(c.Ctx(), c.Log(), form.File["payload"][0])
		if err != nil {
			return fmt.Errorf("failed to parse: %w", err)
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
		var vars getItemByAccountNumberVars
		if err := c.Vars(&vars); err != nil {
			return fmt.Errorf("failed to read vars: %w", err)
		}

		item, err := registryService.GetItemByAccountNumber(c.Ctx(), c.Log(), vars.AccountNumber)
		if err != nil {
			return fmt.Errorf("failed to get item by account number: %w", err)
		}

		return c.WriteJson(http.StatusOK, item)
	}
}

func GetItemsByAccountNumberRegularHandler(registryService *registry.Service) router.Handler {
	return func(c router.Context) error {
		var vars getItemByAccountNumberVars
		if err := c.Vars(&vars); err != nil {
			return fmt.Errorf("failed to read vars: %w", err)
		}

		items, err := registryService.GetItemsByAccountNumberRegular(c.Ctx(), c.Log(), vars.AccountNumber)
		if err != nil {
			return fmt.Errorf("failed to get items by account number: %w", err)
		}

		return c.WriteJson(http.StatusOK, items)
	}
}
