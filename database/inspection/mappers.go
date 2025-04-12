package inspection

import (
	"tns-energo/database/consumer"
	"tns-energo/database/device"
	"tns-energo/database/file"
	domain "tns-energo/service/inspection"
)

func MapToDb(i domain.Inspection) Inspection {
	return Inspection{
		Id:                      i.Id,
		TaskId:                  i.TaskId,
		BrigadeId:               i.BrigadeId,
		Type:                    int(i.Type),
		ActNumber:               i.ActNumber,
		Resolution:              int(i.Resolution),
		Address:                 i.Address,
		Consumer:                consumer.MapToDb(i.Consumer),
		HaveAutomaton:           i.HaveAutomaton,
		AccountNumber:           i.AccountNumber,
		IsIncompletePayment:     i.IsIncompletePayment,
		OtherReason:             i.OtherReason,
		MethodBy:                int(i.MethodBy),
		Method:                  i.Method,
		Device:                  device.MapToDb(i.Device),
		ReasonType:              int(i.ReasonType),
		Reason:                  i.Reason,
		ActCopies:               i.ActCopies,
		Images:                  file.MapSliceToDb(i.Images),
		InspectionDate:          i.InspectionDate,
		ResolutionFile:          file.MapToDb(i.ResolutionFile),
		IsChecked:               i.IsChecked,
		IsViolationDetected:     i.IsViolationDetected,
		IsExpenseAvailable:      i.IsExpenseAvailable,
		OtherViolation:          i.OtherViolation,
		IsUnauthorizedConsumers: i.IsUnauthorizedConsumers,
		UnauthorizedDescription: i.UnauthorizedDescription,
		OldDeviceValue:          i.OldDeviceValue,
		OldDeviceValueDate:      i.OldDeviceValueDate,
		UnauthorizedExplanation: i.UnauthorizedExplanation,
		EnergyActionDate:        i.EnergyActionDate,
	}
}

func MapToDomain(i Inspection) domain.Inspection {
	return domain.Inspection{
		Id:                      i.Id,
		TaskId:                  i.TaskId,
		BrigadeId:               i.BrigadeId,
		Type:                    domain.Type(i.Type),
		ActNumber:               i.ActNumber,
		Resolution:              domain.Resolution(i.Resolution),
		Address:                 i.Address,
		Consumer:                consumer.MapToDomain(i.Consumer),
		HaveAutomaton:           i.HaveAutomaton,
		AccountNumber:           i.AccountNumber,
		IsIncompletePayment:     i.IsIncompletePayment,
		OtherReason:             i.OtherReason,
		MethodBy:                domain.MethodBy(i.MethodBy),
		Method:                  i.Method,
		Device:                  device.MapToDomain(i.Device),
		ReasonType:              domain.ReasonType(i.ReasonType),
		Reason:                  i.Reason,
		ActCopies:               i.ActCopies,
		Images:                  file.MapSliceToDomain(i.Images),
		InspectionDate:          i.InspectionDate,
		ResolutionFile:          file.MapToDomain(i.ResolutionFile),
		IsChecked:               i.IsChecked,
		IsViolationDetected:     i.IsViolationDetected,
		IsExpenseAvailable:      i.IsExpenseAvailable,
		OtherViolation:          i.OtherViolation,
		IsUnauthorizedConsumers: i.IsUnauthorizedConsumers,
		UnauthorizedDescription: i.UnauthorizedDescription,
		OldDeviceValue:          i.OldDeviceValue,
		OldDeviceValueDate:      i.OldDeviceValueDate,
		UnauthorizedExplanation: i.UnauthorizedExplanation,
		EnergyActionDate:        i.EnergyActionDate,
	}
}

func MapSliceToDomain(inspections []Inspection) []domain.Inspection {
	result := make([]domain.Inspection, 0, len(inspections))
	for _, i := range inspections {
		result = append(result, MapToDomain(i))
	}

	return result
}
