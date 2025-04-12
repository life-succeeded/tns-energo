package report

import (
	"tns-energo/database/file"
	domain "tns-energo/service/analytics"
)

func MapToDb(r domain.Report) Report {
	return Report{
		Id:        r.Id,
		Type:      int(r.Type),
		File:      file.MapToDb(r.File),
		ForDate:   r.ForDate,
		CreatedAt: r.CreatedAt,
	}
}

func MapToDomain(r Report) domain.Report {
	return domain.Report{
		Id:        r.Id,
		Type:      domain.ReportType(r.Type),
		File:      file.MapToDomain(r.File),
		ForDate:   r.ForDate,
		CreatedAt: r.CreatedAt,
	}
}

func MapSliceToDomain(reports []Report) []domain.Report {
	result := make([]domain.Report, 0, len(reports))
	for _, r := range reports {
		result = append(result, MapToDomain(r))
	}

	return result
}
