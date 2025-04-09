package inspection

import (
	domain "tns-energo/service/inspection"
)

func mapToDb(inspection domain.Inspection) Inspection {
	return Inspection{
		Id:                 inspection.Id,
		InspectorId:        inspection.InspectorId,
		AccountNumber:      inspection.AccountNumber,
		ConsumerSurname:    inspection.ConsumerSurname,
		ConsumerName:       inspection.ConsumerName,
		ConsumerPatronymic: inspection.ConsumerPatronymic,
		Object:             inspection.Object,
		Resolution:         int(inspection.Resolution),
		Reason:             inspection.Reason,
		Method:             inspection.Method,
		IsReviewRefused:    inspection.IsReviewRefused,
		ActionDate:         inspection.ActionDate,
		HaveAutomaton:      inspection.HaveAutomaton,
		Images:             inspection.Images,
		DeviceType:         inspection.DeviceType,
		DeviceNumber:       inspection.DeviceNumber,
		Voltage:            inspection.Voltage,
		Amperage:           inspection.Amperage,
		DeviceValue:        inspection.DeviceValue,
		InspectionDate:     inspection.InspectionDate,
		AccuracyClass:      inspection.AccuracyClass,
		TariffsCount:       inspection.TariffsCount,
		DeploymentPlace:    inspection.DeploymentPlace,
		ResolutionFileName: inspection.ResolutionFileName,
		CreatedAt:          inspection.CreatedAt,
		UpdatedAt:          inspection.UpdatedAt,
	}
}

func mapToDomain(inspection Inspection) domain.Inspection {
	return domain.Inspection{
		Id:                 inspection.Id,
		InspectorId:        inspection.InspectorId,
		AccountNumber:      inspection.AccountNumber,
		ConsumerSurname:    inspection.ConsumerSurname,
		ConsumerName:       inspection.ConsumerName,
		ConsumerPatronymic: inspection.ConsumerPatronymic,
		Object:             inspection.Object,
		Resolution:         domain.Resolution(inspection.Resolution),
		Reason:             inspection.Reason,
		Method:             inspection.Method,
		IsReviewRefused:    inspection.IsReviewRefused,
		ActionDate:         inspection.ActionDate,
		HaveAutomaton:      inspection.HaveAutomaton,
		Images:             inspection.Images,
		DeviceType:         inspection.DeviceType,
		DeviceNumber:       inspection.DeviceNumber,
		Voltage:            inspection.Voltage,
		Amperage:           inspection.Amperage,
		DeviceValue:        inspection.DeviceValue,
		InspectionDate:     inspection.InspectionDate,
		AccuracyClass:      inspection.AccuracyClass,
		TariffsCount:       inspection.TariffsCount,
		DeploymentPlace:    inspection.DeploymentPlace,
		ResolutionFileName: inspection.ResolutionFileName,
		CreatedAt:          inspection.CreatedAt,
		UpdatedAt:          inspection.UpdatedAt,
	}
}

func mapSliceToDomain(inspections []Inspection) []domain.Inspection {
	domainInspections := make([]domain.Inspection, 0, len(inspections))
	for _, inspection := range inspections {
		domainInspections = append(domainInspections, mapToDomain(inspection))
	}

	return domainInspections
}
