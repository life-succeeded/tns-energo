package analytics

import (
	"fmt"
	"time"
	"tns-energo/config"
	libctx "tns-energo/lib/ctx"
	liblog "tns-energo/lib/log"

	"github.com/google/uuid"
)

type Service struct {
	settings config.Settings
	reports  ReportStorage
}

func NewService(settings config.Settings, reports ReportStorage) *Service {
	return &Service{
		settings: settings,
		reports:  reports,
	}
}

func (s *Service) GenerateDailyReport(ctx libctx.Context, log liblog.Logger, date time.Time) (Report, error) {
	report := Report{
		Type:      Daily,
		Name:      fmt.Sprintf("Daily Report %s.docx", uuid.New()),
		URL:       "test url",
		ForDate:   date,
		CreatedAt: time.Now(),
	}

	id, err := s.reports.AddOne(ctx, report)
	if err != nil {
		return Report{}, fmt.Errorf("failed to add daily report: %w", err)
	}

	report.Id = id

	return report, nil
}

func (s *Service) GetAllReports(ctx libctx.Context, log liblog.Logger) ([]Report, error) {
	reports, err := s.reports.GetAll(ctx, log)
	if err != nil {
		return nil, fmt.Errorf("failed to get all reports: %w", err)
	}

	return reports, nil
}
