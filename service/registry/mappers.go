package registry

import dbregistry "tns-energo/database/registry"

func MapToDomain(dbItem dbregistry.Item) Item {
	return Item{
		Id:            dbItem.Id,
		AccountNumber: dbItem.AccountNumber,
		Surname:       dbItem.Surname,
		Name:          dbItem.Name,
		Patronymic:    dbItem.Patronymic,
		Object:        dbItem.Object,
		HaveAutomaton: dbItem.HaveAutomaton,
	}
}
