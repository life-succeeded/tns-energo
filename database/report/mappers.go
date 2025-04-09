package report

import (
	domain "tns-energo/service/analytics"
)

func mapToDb(r domain.Report) Report {
	return Report{
		Type:      int(r.Type),
		Name:      r.Name,
		URL:       r.URL,
		ForDate:   r.ForDate,
		CreatedAt: r.CreatedAt,
	}
}

func mapToDomain(r Report) domain.Report {
	return domain.Report{
		Type:      domain.ReportType(r.Type),
		Name:      r.Name,
		URL:       r.URL,
		ForDate:   r.ForDate,
		CreatedAt: r.CreatedAt,
	}
}
