package inspection

import (
	"bytes"
	"fmt"
	"time"
	"tns-energo/config"
	libctx "tns-energo/lib/ctx"
	liblog "tns-energo/lib/log"
	"tns-energo/service/file"
	"tns-energo/service/task"

	"github.com/google/uuid"
	"github.com/lukasjarosch/go-docx"
)

type Service struct {
	settings    config.Settings
	inspections Storage
	documents   DocumentStorage
	registry    RegistryStorage
	tasks       TaskStorage
}

func NewService(settings config.Settings, inspections Storage, documents DocumentStorage, registry RegistryStorage, tasks TaskStorage) *Service {
	return &Service{
		settings:    settings,
		inspections: inspections,
		documents:   documents,
		registry:    registry,
		tasks:       tasks,
	}
}

func (s *Service) Inspect(ctx libctx.Context, log liblog.Logger, request InspectRequest) (file.File, error) {
	inspection := Inspection{
		InspectorId:         ctx.Authorize.UserId,
		TaskId:              request.TaskId,
		AccountNumber:       request.AccountNumber,
		Consumer:            request.Consumer,
		Address:             request.Address,
		Resolution:          request.Resolution,
		Reason:              request.Reason,
		Method:              request.Method,
		IsReviewRefused:     request.IsReviewRefused,
		ActionDate:          request.ActionDate,
		HaveAutomaton:       request.HaveAutomaton,
		AutomatonSealNumber: request.AutomatonSealNumber,
		Images:              request.Images,
		Device:              request.Device,
		InspectionDate:      request.InspectionDate,
		ActNumber:           request.ActNumber,
		Contract:            request.Contract,
		SealNumber:          request.SealNumber,
	}

	buf, err := s.generateAct(inspection)
	if err != nil {
		return file.File{}, fmt.Errorf("could not generate act: %w", err)
	}

	inspection.ResolutionFile = file.File{
		Name: fmt.Sprintf("%s.docx", uuid.New()),
	}
	url, err := s.documents.Add(ctx, inspection.ResolutionFile.Name, buf, buf.Len())
	if err != nil {
		return file.File{}, fmt.Errorf("could not create object: %w", err)
	}

	now := time.Now()
	inspection.ResolutionFile.URL = url
	inspection.CreatedAt = now
	inspection.UpdatedAt = now
	err = s.inspections.AddOne(ctx, inspection)
	if err != nil {
		return file.File{}, fmt.Errorf("could not add inspection: %w", err)
	}

	err = s.tasks.UpdateStatus(ctx, request.TaskId, task.Done)
	if err != nil {
		return file.File{}, fmt.Errorf("could not update task status: %w", err)
	}

	/*item, err := s.registry.GetByAccountNumber(ctx, inspection.AccountNumber)
	if err != nil {
		if !errors.Is(err, registry.ErrItemNotFound) {
			return file.File{}, fmt.Errorf("could not get registry item: %w", err)
		}

		addErr := s.registry.AddOne(ctx, registry.Item{
			AccountNumber: inspection.AccountNumber,
			Address:       inspection.Address,
			CreatedAt:     now,
			UpdatedAt:     now,
		})
		if addErr != nil {
			return file.File{}, fmt.Errorf("could not add registry item: %w", err)
		}

		return inspection.ResolutionFile, nil
	}

	err = s.registry.UpdateOne(ctx, registry.Item{
		AccountNumber: inspection.AccountNumber,
		Address:       inspection.Address,
		CreatedAt:     item.CreatedAt,
		UpdatedAt:     now,
	})
	if err != nil {
		return file.File{}, fmt.Errorf("could not update registry item: %w", err)
	}*/

	return inspection.ResolutionFile, nil
}

func (s *Service) GetByInspectorId(ctx libctx.Context, log liblog.Logger, inspectorId int) ([]Inspection, error) {
	inspections, err := s.inspections.GetByInspectorId(ctx, log, inspectorId)
	if err != nil {
		return nil, fmt.Errorf("could not get inspections: %w", err)
	}

	return inspections, nil
}

func (s *Service) generateAct(inspection Inspection) (*bytes.Buffer, error) {
	consumerName := fmt.Sprintf("%s %s", inspection.Consumer.Surname, inspection.Consumer.Name)
	if len(inspection.Consumer.Patronymic) != 0 {
		consumerName = fmt.Sprintf("%s %s", consumerName, inspection.Consumer.Patronymic)
	}
	if len(consumerName) == 0 {
		consumerName = inspection.Consumer.LegalEntityName
	}

	var (
		haveAutomaton = "□"
		noAutomaton   = "■"
	)
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
		"consumer_agent_name":   consumerName,
		"account_number":        inspection.AccountNumber,
		"contract_number":       inspection.Contract.Number,
		"contract_date":         inspection.Contract.Date.Format("02.01.2006"),
		"object":                inspection.Address,
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

	doc, err := docx.Open(s.settings.Templates.Universal)
	if err != nil {
		return nil, fmt.Errorf("could not open act template: %w", err)
	}

	err = doc.ReplaceAll(replaceMap)
	if err != nil {
		return nil, fmt.Errorf("could not replace vars in act: %w", err)
	}

	buf := &bytes.Buffer{}
	err = doc.Write(buf)
	if err != nil {
		return nil, fmt.Errorf("could not write act to buffer: %w", err)
	}

	return buf, nil
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
