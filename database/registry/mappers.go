package registry

import "tns-energo/service/registry"

func mapToDb(item registry.Item) Item {
	return Item{
		Id:            item.Id,
		AccountNumber: item.AccountNumber,
		Surname:       item.Surname,
		Name:          item.Name,
		Patronymic:    item.Patronymic,
		Object:        item.Object,
		HaveAutomaton: item.HaveAutomaton,
		CreatedAt:     item.CreatedAt,
		UpdatedAt:     item.UpdatedAt,
	}
}

func mapToDomain(item Item) registry.Item {
	return registry.Item{
		Id:            item.Id,
		AccountNumber: item.AccountNumber,
		Surname:       item.Surname,
		Name:          item.Name,
		Patronymic:    item.Patronymic,
		Object:        item.Object,
		HaveAutomaton: item.HaveAutomaton,
		CreatedAt:     item.CreatedAt,
		UpdatedAt:     item.UpdatedAt,
	}
}

func mapSliceToDomain(items []Item) []registry.Item {
	domainItems := make([]registry.Item, 0, len(items))
	for _, item := range items {
		domainItems = append(domainItems, mapToDomain(item))
	}

	return domainItems
}
