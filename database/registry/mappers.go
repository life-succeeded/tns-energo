package registry

import domain "tns-energo/service/registry"

func MapToDb(i domain.Item) Item {
	return Item{
		Id:            i.Id,
		AccountNumber: i.AccountNumber,
		Surname:       i.Surname,
		Name:          i.Name,
		Patronymic:    i.Patronymic,
		Address:       i.Address,
		HaveAutomaton: i.HaveAutomaton,
		CreatedAt:     i.CreatedAt,
		UpdatedAt:     i.UpdatedAt,
	}
}

func MapToDomain(i Item) domain.Item {
	return domain.Item{
		Id:            i.Id,
		AccountNumber: i.AccountNumber,
		Surname:       i.Surname,
		Name:          i.Name,
		Patronymic:    i.Patronymic,
		Address:       i.Address,
		HaveAutomaton: i.HaveAutomaton,
		CreatedAt:     i.CreatedAt,
		UpdatedAt:     i.UpdatedAt,
	}
}

func MapSliceToDomain(items []Item) []domain.Item {
	result := make([]domain.Item, 0, len(items))
	for _, i := range items {
		result = append(result, MapToDomain(i))
	}

	return result
}
