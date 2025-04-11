package inspection

import (
	"bytes"
	"fmt"
	"strings"
	"time"
	"tns-energo/config"
	libctx "tns-energo/lib/ctx"
	liblog "tns-energo/lib/log"
	"tns-energo/service/brigade"
	"tns-energo/service/device"
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
	brigades    BrigadeStorage
}

func NewService(settings config.Settings, inspections Storage, documents DocumentStorage, registry RegistryStorage, tasks TaskStorage, brigades BrigadeStorage) *Service {
	return &Service{
		settings:    settings,
		inspections: inspections,
		documents:   documents,
		registry:    registry,
		tasks:       tasks,
		brigades:    brigades,
	}
}

func (s *Service) InspectUniversal(ctx libctx.Context, log liblog.Logger, request InspectUniversalRequest) (file.File, error) {
	now := time.Now()
	inspection := Inspection{
		TaskId:              request.TaskId,
		BrigadeId:           request.BrigadeId,
		ActNumber:           request.ActNumber,
		Resolution:          request.Resolution,
		Address:             request.Address,
		Consumer:            request.Consumer,
		HaveAutomaton:       request.HaveAutomaton,
		AccountNumber:       request.AccountNumber,
		IsIncompletePayment: request.IsIncompletePayment,
		OtherReason:         request.OtherReason,
		MethodBy:            request.MethodBy,
		Method:              request.Method,
		Device:              request.Device,
		ReasonType:          request.ReasonType,
		Reason:              request.Reason,
		ActCopies:           request.ActCopies,
		Images:              request.Images,
		InspectionDate:      now,
	}

	brig, err := s.brigades.GetById(ctx, inspection.BrigadeId)
	if err != nil {
		return file.File{}, fmt.Errorf("could not find brigade: %w", err)
	}

	buf, err := s.generateAct(inspection, brig)
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

	inspection.ResolutionFile.URL = url
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

func (s *Service) GetByBrigadeId(ctx libctx.Context, log liblog.Logger, brigadeId string) ([]Inspection, error) {
	inspections, err := s.inspections.GetByBrigadeId(ctx, log, brigadeId)
	if err != nil {
		return nil, fmt.Errorf("could not get inspections: %w", err)
	}

	return inspections, nil
}

func (s *Service) generateAct(inspection Inspection, brig brigade.Brigade) (*bytes.Buffer, error) {
	isLimitation := "■"
	isResumption := "□"
	if inspection.Resolution == ResumedResolution {
		isLimitation = "□"
		isResumption = "■"
	}

	consumerFIO := fmt.Sprintf("%s %s", inspection.Consumer.Surname, inspection.Consumer.Name)
	if len(inspection.Consumer.Patronymic) != 0 {
		consumerFIO = fmt.Sprintf("%s %s", consumerFIO, inspection.Consumer.Patronymic)
	}

	haveAutomaton := "□"
	noAutomaton := "■"
	if inspection.HaveAutomaton {
		haveAutomaton = "■"
		noAutomaton = "□"
	}

	isIncomplete := "□"
	if inspection.IsIncompletePayment {
		isIncomplete = "■"
	}

	isOtherReason := "□"
	if len(inspection.OtherReason) != 0 {
		isOtherReason = "■"
	}

	isByConsumer := "□"
	isByInspector := "■"
	if inspection.MethodBy == Consumer {
		isByConsumer = "■"
		isByInspector = "□"
	}

	isEnergyLimited := "□"
	isEnergyStopped := "□"
	isEnergyResumed := "□"
	switch inspection.Resolution {
	case LimitedResolution:
		isEnergyLimited = "■"
	case StoppedResolution:
		isEnergyStopped = "■"
	case ResumedResolution:
		isEnergyResumed = "■"
	}

	isInside := "□"
	isOutside := "□"
	switch inspection.Device.DeploymentPlace {
	case device.Inside:
		isInside = "■"
	case device.Outside:
		isOutside = "■"
	}

	seals := make([]string, 0, len(inspection.Device.Seals))
	for _, seal := range inspection.Device.Seals {
		seals = append(seals, fmt.Sprintf("№%s %s", seal.Number, seal.Place))
	}

	isConsumerLimited := "□"
	isInspectorLimited := "□"
	isNotIntroduced := "□"
	switch inspection.ReasonType {
	case NotIntroduced:
		isNotIntroduced = "■"
	case ConsumerLimited:
		isConsumerLimited = "■"
	case InspectorLimited:
		isInspectorLimited = "■"
	}

	firstInspector := fmt.Sprintf("%s %s.", brig.FirstInspector.Surname, string([]rune(brig.FirstInspector.Name)[0]))
	if len(brig.FirstInspector.Patronymic) != 0 {
		firstInspector = fmt.Sprintf("%s%s.", firstInspector, string([]rune(brig.FirstInspector.Patronymic)[0]))
	}

	secondInspector := fmt.Sprintf("%s %s.", brig.SecondInspector.Surname, string([]rune(brig.SecondInspector.Name)[0]))
	if len(brig.SecondInspector.Patronymic) != 0 {
		secondInspector = fmt.Sprintf("%s%s.", secondInspector, string([]rune(brig.SecondInspector.Patronymic)[0]))
	}

	replaceMap := docx.PlaceholderMap{
		"act_number":               inspection.ActNumber,
		"is_limitation":            isLimitation,
		"is_resumption":            isResumption,
		"act_day":                  inspection.InspectionDate.Format("02"),
		"act_month":                russianMonth(inspection.InspectionDate.Month()),
		"act_year":                 inspection.InspectionDate.Year(),
		"act_hour":                 inspection.InspectionDate.Format("15"),
		"act_minute":               inspection.InspectionDate.Format("04"),
		"act_place":                inspection.Address,
		"consumer_fio":             consumerFIO,
		"address":                  inspection.Address,
		"have_automaton":           haveAutomaton,
		"no_automaton":             noAutomaton,
		"account_number":           inspection.AccountNumber,
		"is_incomplete_payment":    isIncomplete,
		"is_other_reason":          isOtherReason,
		"other_reason":             inspection.OtherReason,
		"is_energy_limited":        isEnergyLimited,
		"is_energy_stopped":        isEnergyStopped,
		"is_energy_resumed":        isEnergyResumed,
		"energy_hour":              inspection.InspectionDate.Format("15"),
		"energy_minute":            inspection.InspectionDate.Format("04"),
		"energy_day":               inspection.InspectionDate.Format("02"),
		"energy_month":             russianMonth(inspection.InspectionDate.Month()),
		"energy_year":              inspection.InspectionDate.Year(),
		"is_by_consumer":           isByConsumer,
		"is_by_inspector":          isByInspector,
		"method":                   inspection.Method,
		"is_inside":                isInside,
		"is_outside":               isOutside,
		"other_place":              inspection.Device.OtherPlace,
		"device_type":              inspection.Device.Type,
		"device_number":            inspection.Device.Number,
		"device_value":             inspection.Device.Value,
		"seals":                    strings.Join(seals, ", "),
		"is_consumer_limited":      isConsumerLimited,
		"is_inspector_limited":     isInspectorLimited,
		"is_not_introduced":        isNotIntroduced,
		"is_not_introduced_reason": inspection.Reason,
		"act_copies":               inspection.ActCopies,
		"inspector1_initials":      firstInspector,
		"inspector2_initials":      secondInspector,
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
