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
	}
}

func MapToDomain(b Brigade) domain.Brigade {
	return domain.Brigade{
		Id:              b.Id,
		FirstInspector:  inspector.MapToDomain(b.FirstInspector),
		SecondInspector: inspector.MapToDomain(b.SecondInspector),
		CreatedAt:       b.CreatedAt,
	}
}
