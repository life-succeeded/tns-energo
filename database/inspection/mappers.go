package inspection

import (
	"tns-energo/database/consumer"
	"tns-energo/database/contract"
	"tns-energo/database/device"
	"tns-energo/database/file"
	domain "tns-energo/service/inspection"
)

func MapToDb(i domain.Inspection) Inspection {
	return Inspection{
		Id:                  i.Id,
		InspectorId:         i.InspectorId,
		AccountNumber:       i.AccountNumber,
		Consumer:            consumer.MapToDb(i.Consumer),
		Address:             i.Address,
		Resolution:          int(i.Resolution),
		Reason:              i.Reason,
		Method:              i.Method,
		IsReviewRefused:     i.IsReviewRefused,
		ActionDate:          i.ActionDate,
		HaveAutomaton:       i.HaveAutomaton,
		AutomatonSealNumber: i.AutomatonSealNumber,
		Images:              file.MapSliceToDb(i.Images),
		Device:              device.MapToDb(i.Device),
		InspectionDate:      i.InspectionDate,
		ResolutionFile:      file.MapToDb(i.ResolutionFile),
		ActNumber:           i.ActNumber,
		Contract:            contract.MapToDb(i.Contract),
		SealNumber:          i.SealNumber,
		CreatedAt:           i.CreatedAt,
		UpdatedAt:           i.UpdatedAt,
	}
}

func MapToDomain(i Inspection) domain.Inspection {
	return domain.Inspection{
		Id:                  i.Id,
		InspectorId:         i.InspectorId,
		AccountNumber:       i.AccountNumber,
		Consumer:            consumer.MapToDomain(i.Consumer),
		Address:             i.Address,
		Resolution:          domain.Resolution(i.Resolution),
		Reason:              i.Reason,
		Method:              i.Method,
		IsReviewRefused:     i.IsReviewRefused,
		ActionDate:          i.ActionDate,
		HaveAutomaton:       i.HaveAutomaton,
		AutomatonSealNumber: i.AutomatonSealNumber,
		Images:              file.MapSliceToDomain(i.Images),
		Device:              device.MapToDomain(i.Device),
		InspectionDate:      i.InspectionDate,
		ResolutionFile:      file.MapToDomain(i.ResolutionFile),
		ActNumber:           i.ActNumber,
		Contract:            contract.MapToDomain(i.Contract),
		SealNumber:          i.SealNumber,
		CreatedAt:           i.CreatedAt,
		UpdatedAt:           i.UpdatedAt,
	}
}

func MapSliceToDomain(inspections []Inspection) []domain.Inspection {
	result := make([]domain.Inspection, 0, len(inspections))
	for _, i := range inspections {
		result = append(result, MapToDomain(i))
	}

	return result
}
