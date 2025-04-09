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

	err := s.reports.AddOne(ctx, report)
	if err != nil {
		return Report{}, fmt.Errorf("failed to add daily report: %w", err)
	}

	return report, nil
}
