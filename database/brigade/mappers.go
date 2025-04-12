package brigade

import (
	"tns-energo/database/inspector"
	domain "tns-energo/service/brigade"
)

func MapToDb(b domain.Brigade) Brigade {
	return Brigade{
		Id:              b.Id,
		FirstInspector:  inspector.MapToDb(b.FirstInspector),
		SecondInspector: inspector.MapToDb(b.SecondInspector),
		CreatedAt:       b.CreatedAt,
		UpdatedAt:       b.UpdatedAt,
	}
}

func MapToDomain(b Brigade) domain.Brigade {
	return domain.Brigade{
		Id:              b.Id,
		FirstInspector:  inspector.MapToDomain(b.FirstInspector),
		SecondInspector: inspector.MapToDomain(b.SecondInspector),
		CreatedAt:       b.CreatedAt,
		UpdatedAt:       b.UpdatedAt,
	}
}

func MapSliceToDomain(brigades []Brigade) []domain.Brigade {
	result := make([]domain.Brigade, 0, len(brigades))
	for _, b := range brigades {
		result = append(result, MapToDomain(b))
	}

	return result
}
