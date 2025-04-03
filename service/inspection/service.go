package inspection

import (
	"bytes"
	"fmt"
	"time"
	"tns-energo/config"
	libctx "tns-energo/lib/ctx"
	"tns-energo/lib/db/minio"
	liblog "tns-energo/lib/log"

	"github.com/google/uuid"
	"github.com/lukasjarosch/go-docx"
	"github.com/shopspring/decimal"
)

type Service interface {
	Inspect(ctx libctx.Context, log liblog.Logger) (string, error)
	ParseExcelRegistry(ctx libctx.Context, log liblog.Logger) error
}

type Impl struct {
	minio    minio.Client
	settings config.Settings
}

func NewService(minio minio.Client, settings config.Settings) *Impl {
	return &Impl{
		minio:    minio,
		settings: settings,
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

	doc, err := docx.Open("../templates/limitation.docx")
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

	url, err := s.minio.CreateOne(ctx, s.settings.Databases.Minio.DocumentsBucket, minio.File{
		Name: uuid.New().String(),
		Data: buf,
	})
	if err != nil {
		return "", fmt.Errorf("could not create document in minio: %w", err)
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

func (s *Impl) ParseExcelRegistry(ctx libctx.Context, log liblog.Logger) error {
	//TODO implement me
	panic("implement me")
}
