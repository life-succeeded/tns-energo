package analytics

import (
	"fmt"
	"time"
	"tns-energo/config"
	libctx "tns-energo/lib/ctx"
	liblog "tns-energo/lib/log"
	libtime "tns-energo/lib/time"
	"tns-energo/service/file"
	"tns-energo/service/inspection"

	"github.com/xuri/excelize/v2"
)

type Service struct {
	settings    config.Settings
	reports     ReportStorage
	inspections InspectionStorage
	documents   DocumentStorage
}

func NewService(settings config.Settings, reports ReportStorage, inspections InspectionStorage, documents DocumentStorage) *Service {
	return &Service{
		settings:    settings,
		reports:     reports,
		inspections: inspections,
		documents:   documents,
	}
}

func (s *Service) GenerateDailyReport(ctx libctx.Context, log liblog.Logger, date time.Time) (Report, error) {
	inspections, err := s.inspections.GetByInspectionDate(ctx, log, date)
	if err != nil {
		return Report{}, fmt.Errorf("could not get inspections: %w", err)
	}

	report := Report{
		Type: Daily,
		File: file.File{
			Name: fmt.Sprintf("Отчет за %s.xlsx", date.Format("02.01.2006")),
		},
		ForDate:   date,
		CreatedAt: time.Now(),
	}

	f, err := excelize.OpenFile(s.settings.Templates.Report)
	if err != nil {
		return Report{}, fmt.Errorf("failed to open template file: %w", err)
	}

	sheet := f.GetSheetName(0)
	for i, ins := range inspections {
		cell, err := excelize.CoordinatesToCellName(1, i+2)
		if err != nil {
			return Report{}, fmt.Errorf("could not create cell: %w", err)
		}

		work := ""
		result := ""
		if ins.Type == inspection.Resumption {
			work = "Возобновление"
			if ins.Resolution == inspection.ResumedResolution {
				result = "Возобновление"
			} else {
				result = "Недопуск"
			}
		} else if ins.Type == inspection.Limitation {
			work = "Отключение"
			if ins.Resolution != inspection.ResumedResolution {
				result = "Отключение"
			} else {
				result = "Недопуск"
			}
		} else {
			work = "Контроль ранее введенного ограничения"
			if ins.IsViolationDetected {
				result = "Нарушено"
			} else {
				result = "Не нарушено"
			}
		}

		brig := ins.Brigade

		firstInspector := fmt.Sprintf("%s %s", brig.FirstInspector.Surname, brig.FirstInspector.Name)
		if len(brig.FirstInspector.Patronymic) != 0 {
			firstInspector = fmt.Sprintf("%s %s", firstInspector, brig.FirstInspector.Patronymic)
		}

		secondInspector := fmt.Sprintf("%s %s", brig.SecondInspector.Surname, brig.SecondInspector.Name)
		if len(brig.SecondInspector.Patronymic) != 0 {
			secondInspector = fmt.Sprintf("%s %s", secondInspector, brig.SecondInspector.Patronymic)
		}

		values := []any{
			i + 1,
			ins.Address,
			fmt.Sprintf("%s №%s", ins.Device.Type, ins.Device.Number),
			ins.InspectionDate.In(libtime.MoscowLocation()).Format("02.01.2006"),
			ins.InspectionDate.In(libtime.MoscowLocation()).Format("15:04"),
			work,
			result,
			firstInspector,
			secondInspector,
			len(ins.Images),
		}

		err = f.SetSheetRow(sheet, cell, &values)
		if err != nil {
			return report, fmt.Errorf("failed to set sheet row: %w", err)
		}
	}

	buf, err := f.WriteToBuffer()
	if err != nil {
		return Report{}, fmt.Errorf("failed to write report: %w", err)
	}

	url, err := s.documents.Add(ctx, report.File.Name, buf, int64(buf.Len()))
	if err != nil {
		return Report{}, fmt.Errorf("failed to add report: %w", err)
	}

	report.File.URL = url

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
