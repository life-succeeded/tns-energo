package registry

import (
	"tns-energo/database/device"
	domain "tns-energo/service/registry"
)

func MapToDb(i domain.Item) Item {
	return Item{
		Id:            i.Id,
		AccountNumber: i.AccountNumber,
		Address:       i.Address,
		OldDevice:     device.MapToDb(i.OldDevice),
		NewDevice:     device.MapToDb(i.NewDevice),
		CreatedAt:     i.CreatedAt,
		UpdatedAt:     i.UpdatedAt,
	}
}

func MapToDomain(i Item) domain.Item {
	return domain.Item{
		Id:            i.Id,
		AccountNumber: i.AccountNumber,
		Address:       i.Address,
		OldDevice:     device.MapToDomain(i.OldDevice),
		NewDevice:     device.MapToDomain(i.NewDevice),
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
