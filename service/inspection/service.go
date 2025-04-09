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
	// TODO: обновление по номеру счета
	now := time.Now()
	user, err := s.users.GetLightById(ctx, ctx.Authorize.UserId)
	if err != nil {
		return "", fmt.Errorf("could not get user: %w", err)
	}

	inspection.InspectorId = ctx.Authorize.UserId

	consumerName := fmt.Sprintf("%s %s", inspection.Consumer.Surname, inspection.Consumer.Name)
	if len(inspection.Consumer.Patronymic) != 0 {
		consumerName = fmt.Sprintf("%s %s", consumerName, inspection.Consumer.Patronymic)
	}
	if len(consumerName) == 0 {
		consumerName = inspection.Consumer.LegalEntityName
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
		"act_number":            inspection.ActNumber,
		"act_date_day":          inspection.InspectionDate.Day(),
		"act_date_month":        russianMonth(inspection.InspectionDate.Month()),
		"act_date_year":         inspection.InspectionDate.Year(),
		"consumer_name":         consumerName,
		"inspector_position":    inspectorPosition,
		"inspector_name":        inspectorName,
		"consumer_agent_name":   consumerName,
		"account_number":        inspection.AccountNumber,
		"contract_number":       inspection.Contract.Number,
		"contract_date":         inspection.Contract.Date.Format("02.01.2006"),
		"object":                inspection.Object,
		"reason":                inspection.Reason,
		"method":                inspection.Method,
		"seal_number":           inspection.SealNumber,
		"action_date_hours":     inspection.ActionDate.Hour(),
		"action_date_minutes":   inspection.ActionDate.Minute(),
		"action_date_day":       inspection.ActionDate.Day(),
		"action_date_month":     russianMonth(inspection.ActionDate.Month()),
		"action_date_year":      inspection.ActionDate.Year(),
		"have_automaton":        haveAutomaton,
		"no_automaton":          noAutomaton,
		"automaton_seal_number": inspection.AutomatonSealNumber,
		"device_type":           inspection.Device.Type,
		"device_number":         inspection.Device.Number,
		"voltage":               inspection.Device.Voltage,
		"amperage":              inspection.Device.Amperage,
		"valency_before_dot":    inspection.Device.ValencyBeforeDot,
		"valency_after_dot":     inspection.Device.ValencyAfterDot,
		"manufacture_year":      inspection.Device.ManufactureYear,
		"device_value":          inspection.Device.Value,
		"verification_quarter":  inspection.Device.VerificationQuarter,
		"verification_year":     inspection.Device.VerificationYear,
		"accuracy_class":        inspection.Device.AccuracyClass,
		"tariffs_count":         inspection.Device.TariffsCount,
		"deployment_place":      inspection.Device.DeploymentPlace,
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
