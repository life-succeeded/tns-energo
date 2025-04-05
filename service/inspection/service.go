package inspection

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
	"tns-energo/config"
	"tns-energo/database/document"
	"tns-energo/database/inspection"
	libctx "tns-energo/lib/ctx"
	liblog "tns-energo/lib/log"

	"github.com/google/uuid"
	"github.com/lukasjarosch/go-docx"
	"github.com/shopspring/decimal"
	"github.com/xuri/excelize/v2"
)

type Service interface {
	Inspect(ctx libctx.Context, log liblog.Logger) (string, error)
	ParseExcelRegistry(ctx libctx.Context, log liblog.Logger, fileBytes []byte) error
}

type Impl struct {
	settings    config.Settings
	documents   document.Repository
	inspections inspection.Repository
}

func NewService(settings config.Settings, documents document.Repository, inspections inspection.Repository) *Impl {
	return &Impl{
		settings:    settings,
		documents:   documents,
		inspections: inspections,
	}
}

func (s *Impl) Inspect(ctx libctx.Context, log liblog.Logger) (string, error) {
	now := time.Now()

	replaceMap := docx.PlaceholderMap{
		"act_number":            "123",
		"act_date_day":          now.Day(),
		"act_date_month":        russianMonth(now.Month()),
		"act_date_year":         now.Year(),
		"consumer_name":         `ООО "Рога и Копыта"`,
		"inspector_position":    "главный инспектор",
		"inspector_name":        "Пресняков Артем Дмитриевич",
		"consumer_agent_name":   "Каипов Шамиль Артемович",
		"account_number":        "228Z",
		"contract_number":       "69",
		"contract_date":         now.Format("02.01.2006"),
		"object_city":           "Пенза",
		"object_street":         "Ворошилова",
		"object_house":          "13,",
		"object_apartment":      "кв. 77",
		"reason":                "непоправимых невероятных просто последствий, на которые мы никак не можем повлиять вообще, а также длинного текста без конца и края",
		"method":                "ограничения режима потребления электрической энергии навсегда - иными словами - бан",
		"seal_number":           "348957",
		"action_date_hours":     now.Hour(),
		"action_date_minutes":   now.Minute(),
		"action_date_day":       now.Day(),
		"action_date_month":     russianMonth(now.Month()),
		"action_date_year":      now.Year(),
		"has_automaton":         "■",
		"no_automaton":          "□",
		"automaton_seal_number": "2296923",
		"device_type":           "счетчик",
		"device_number":         "2843620",
		"voltage":               decimal.NewFromFloat(220),
		"amperage":              decimal.NewFromFloat(0.5),
		"valency_before_dot":    "100000",
		"valency_after_dot":     "00",
		"manufacture_year":      now.Year() - 5,
		"device_value":          decimal.NewFromFloat(348625.12),
		"verification_quarter":  2,
		"verification_year":     now.Year() - 1,
		"accuracy_class":        "A",
		"tariffs_count":         7,
		"deployment_place":      "г. Пенза, ул. Ворошилова, д. 13, кв. 77",
	}

	doc, err := docx.Open(s.settings.Templates.Limitation)
	if err != nil {
		return "", fmt.Errorf("could not open document: %w", err)
	}

	err = doc.ReplaceAll(replaceMap)
	if err != nil {
		return "", fmt.Errorf("could not replace: %w", err)
	}

	buf := &bytes.Buffer{}
	err = doc.Write(buf)
	if err != nil {
		return "", fmt.Errorf("could not write: %w", err)
	}

	url, err := s.documents.Create(ctx, fmt.Sprintf("%s.docx", uuid.New()), buf, int64(buf.Len()))
	if err != nil {
		return "", fmt.Errorf("could not create document: %w", err)
	}

	return url, nil
}

func russianMonth(month time.Month) string {
	switch month {
	case time.January:
		return "января"
	case time.February:
		return "февраля"
	case time.March:
		return "марта"
	case time.April:
		return "апреля"
	case time.May:
		return "мая"
	case time.June:
		return "июня"
	case time.July:
		return "июля"
	case time.August:
		return "августа"
	case time.September:
		return "сентября"
	case time.October:
		return "октября"
	case time.November:
		return "ноября"
	case time.December:
		return "декабря"
	}

	return ""
}

func (s *Impl) ParseExcelRegistry(ctx libctx.Context, log liblog.Logger, fileBytes []byte) error {
	file, err := excelize.OpenReader(bytes.NewReader(fileBytes))
	if err != nil {
		return fmt.Errorf("could not open file: %w", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Errorf("failed to close file: %v", err)
		}
	}()

	sheet := file.GetSheetName(0)
	rows, err := file.GetRows(sheet)
	if err != nil {
		return fmt.Errorf("could not get rows: %w", err)
	}

	inspections := make([]inspection.Inspection, 0, len(rows)-1)
	for _, row := range rows[1:] {
		if len(row) != 5 {
			log.Errorf("invalid row length: %d", len(row))
			continue
		}

		tariffsCount, err := strconv.Atoi(row[4])
		if err != nil {
			log.Errorf("could not convert tariffs count: %w", err)
			continue
		}

		inspections = append(inspections, inspection.Inspection{
			Surname:      row[0],
			Name:         row[1],
			Patronymic:   row[2],
			Position:     row[3],
			TariffsCount: tariffsCount,
		})
	}

	err = s.inspections.CreateMany(ctx, inspections)
	if err != nil {
		return fmt.Errorf("could not create inspections: %w", err)
	}

	return nil
}
