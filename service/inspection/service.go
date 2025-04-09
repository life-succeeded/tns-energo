package inspection

import (
	"bytes"
	"fmt"
	"time"
	"tns-energo/config"
	libctx "tns-energo/lib/ctx"
	liblog "tns-energo/lib/log"

	"github.com/google/uuid"
	"github.com/lukasjarosch/go-docx"
)

type Service struct {
	settings    config.Settings
	inspections Storage
	documents   DocumentStorage
	users       UserStorage
}

func NewService(settings config.Settings, inspections Storage, documents DocumentStorage, users UserStorage) *Service {
	return &Service{
		settings:    settings,
		inspections: inspections,
		documents:   documents,
		users:       users,
	}
}

func (s *Service) Inspect(ctx libctx.Context, log liblog.Logger, inspection Inspection) (string, error) {
	// TODO: обновление по номеру счета, добавить недостающие поля
	now := time.Now()
	user, err := s.users.GetLightById(ctx, ctx.Authorize.UserId)
	if err != nil {
		return "", fmt.Errorf("could not get user: %w", err)
	}

	inspection.InspectorId = ctx.Authorize.UserId

	consumerName := fmt.Sprintf("%s %s", inspection.ConsumerSurname, inspection.ConsumerName)
	if inspection.ConsumerPatronymic != nil {
		consumerName = fmt.Sprintf("%s %s", consumerName, *inspection.ConsumerPatronymic)
	}

	inspectorPosition := ""
	if user.Position != nil {
		inspectorPosition = *user.Position
	}

	inspectorName := fmt.Sprintf("%s %s", user.Surname, user.Name)
	if user.Patronymic != nil {
		inspectorName = fmt.Sprintf("%s %s", inspectorName, *user.Patronymic)
	}

	haveAutomaton := "□"
	noAutomaton := "■"
	if inspection.HaveAutomaton {
		haveAutomaton = "■"
		noAutomaton = "□"
	}

	replaceMap := docx.PlaceholderMap{
		"act_number":            "123",
		"act_date_day":          inspection.InspectionDate.Day(),
		"act_date_month":        russianMonth(inspection.InspectionDate.Month()),
		"act_date_year":         inspection.InspectionDate.Year(),
		"consumer_name":         consumerName,
		"inspector_position":    inspectorPosition,
		"inspector_name":        inspectorName,
		"consumer_agent_name":   consumerName,
		"account_number":        inspection.AccountNumber,
		"contract_number":       "69",
		"contract_date":         now.Format("02.01.2006"),
		"object":                inspection.Object,
		"reason":                inspection.Reason,
		"method":                inspection.Method,
		"seal_number":           "348957",
		"action_date_hours":     inspection.ActionDate.Hour(),
		"action_date_minutes":   inspection.ActionDate.Minute(),
		"action_date_day":       inspection.ActionDate.Day(),
		"action_date_month":     russianMonth(inspection.ActionDate.Month()),
		"action_date_year":      inspection.ActionDate.Year(),
		"have_automaton":        haveAutomaton,
		"no_automaton":          noAutomaton,
		"automaton_seal_number": "2296923",
		"device_type":           inspection.DeviceType,
		"device_number":         inspection.DeviceNumber,
		"voltage":               inspection.Voltage,
		"amperage":              inspection.Amperage,
		"valency_before_dot":    "100000",
		"valency_after_dot":     "00",
		"manufacture_year":      now.Year() - 5,
		"device_value":          inspection.DeviceValue,
		"verification_quarter":  2,
		"verification_year":     now.Year() - 1,
		"accuracy_class":        inspection.AccuracyClass,
		"tariffs_count":         inspection.TariffsCount,
		"deployment_place":      inspection.DeploymentPlace,
	}

	path := s.settings.Templates.Limitation
	if inspection.Resolution == Resumption {
		path = s.settings.Templates.Resumption
	}

	doc, err := docx.Open(path)
	if err != nil {
		return "", fmt.Errorf("could not open object: %w", err)
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

	inspection.ResolutionFileName = fmt.Sprintf("%s.docx", uuid.New())
	url, err := s.documents.Add(ctx, inspection.ResolutionFileName, buf, buf.Len())
	if err != nil {
		return "", fmt.Errorf("could not create object: %w", err)
	}

	inspection.CreatedAt = now
	inspection.UpdatedAt = now
	err = s.inspections.AddOne(ctx, inspection)
	if err != nil {
		return "", fmt.Errorf("could not add inspection: %w", err)
	}

	return url, nil
}

func (s *Service) GetByInspectorId(ctx libctx.Context, log liblog.Logger, inspectorId int) ([]Inspection, error) {
	inspections, err := s.inspections.GetByInspectorId(ctx, log, inspectorId)
	if err != nil {
		return nil, fmt.Errorf("could not get inspections: %w", err)
	}

	return inspections, nil
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
