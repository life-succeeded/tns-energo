package report

import (
	domain "tns-energo/service/analytics"
)

func mapToDb(r domain.Report) Report {
	return Report{
		Id:        r.Id,
		Type:      int(r.Type),
		Name:      r.Name,
		URL:       r.URL,
		ForDate:   r.ForDate,
		CreatedAt: r.CreatedAt,
	}
}

func mapToDomain(r Report) domain.Report {
	return domain.Report{
		Id:        r.Id,
		Type:      domain.ReportType(r.Type),
		Name:      r.Name,
		URL:       r.URL,
		ForDate:   r.ForDate,
		CreatedAt: r.CreatedAt,
	}
}

func mapSliceToDomain(reports []Report) []domain.Report {
	domainReports := make([]domain.Report, 0, len(reports))
	for _, r := range reports {
		domainReports = append(domainReports, mapToDomain(r))
	}

	return domainReports
}
