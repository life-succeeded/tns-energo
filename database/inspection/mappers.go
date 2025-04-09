package inspection

import (
	"tns-energo/service/image"
	domain "tns-energo/service/inspection"

	"github.com/shopspring/decimal"
)

func mapToDb(inspection domain.Inspection) Inspection {
	return Inspection{
		Id:                  inspection.Id,
		InspectorId:         inspection.InspectorId,
		AccountNumber:       inspection.AccountNumber,
		Consumer:            mapConsumerToDb(inspection.Consumer),
		Object:              inspection.Object,
		Resolution:          int(inspection.Resolution),
		Reason:              inspection.Reason,
		Method:              inspection.Method,
		IsReviewRefused:     inspection.IsReviewRefused,
		ActionDate:          inspection.ActionDate,
		HaveAutomaton:       inspection.HaveAutomaton,
		AutomatonSealNumber: inspection.AutomatonSealNumber,
		Images:              mapImageSliceToDb(inspection.Images),
		Device:              mapDeviceToDb(inspection.Device),
		InspectionDate:      inspection.InspectionDate,
		ResolutionFileName:  inspection.ResolutionFileName,
		ActNumber:           inspection.ActNumber,
		Contract:            mapContractToDb(inspection.Contract),
		SealNumber:          inspection.SealNumber,
		CreatedAt:           inspection.CreatedAt,
		UpdatedAt:           inspection.UpdatedAt,
	}
}

func mapConsumerToDb(consumer domain.Consumer) Consumer {
	return Consumer{
		Surname:         consumer.Surname,
		Name:            consumer.Name,
		Patronymic:      consumer.Patronymic,
		LegalEntityName: consumer.LegalEntityName,
	}
}

func mapImageToDb(image image.Image) Image {
	return Image{
		Name: image.Name,
		URL:  image.URL,
	}
}

func mapImageSliceToDb(images []image.Image) []Image {
	result := make([]Image, 0, len(images))
	for _, img := range images {
		result = append(result, mapImageToDb(img))
	}

	return result
}

func mapDeviceToDb(device domain.Device) Device {
	return Device{
		Type:                device.Type,
		Number:              device.Number,
		Voltage:             device.Voltage,
		Amperage:            device.Amperage,
		AccuracyClass:       device.AccuracyClass,
		TariffsCount:        device.TariffsCount,
		DeploymentPlace:     device.DeploymentPlace,
		ValencyBeforeDot:    device.ValencyBeforeDot,
		ValencyAfterDot:     device.ValencyAfterDot,
		ManufactureYear:     device.ManufactureYear,
		VerificationQuarter: device.VerificationQuarter,
		VerificationYear:    device.VerificationYear,
		Value:               device.Value.String(),
	}
}

func mapContractToDb(contract domain.Contract) Contract {
	return Contract{
		Number: contract.Number,
		Date:   contract.Date,
	}
}

func mapToDomain(inspection Inspection) domain.Inspection {
	return domain.Inspection{
		Id:                  inspection.Id,
		InspectorId:         inspection.InspectorId,
		AccountNumber:       inspection.AccountNumber,
		Consumer:            mapConsumerToDomain(inspection.Consumer),
		Object:              inspection.Object,
		Resolution:          domain.Resolution(inspection.Resolution),
		Reason:              inspection.Reason,
		Method:              inspection.Method,
		IsReviewRefused:     inspection.IsReviewRefused,
		ActionDate:          inspection.ActionDate,
		HaveAutomaton:       inspection.HaveAutomaton,
		AutomatonSealNumber: inspection.AutomatonSealNumber,
		Images:              mapImageSliceToDomain(inspection.Images),
		Device:              mapDeviceToDomain(inspection.Device),
		InspectionDate:      inspection.InspectionDate,
		ResolutionFileName:  inspection.ResolutionFileName,
		ActNumber:           inspection.ActNumber,
		Contract:            mapContractToDomain(inspection.Contract),
		SealNumber:          inspection.SealNumber,
		CreatedAt:           inspection.CreatedAt,
		UpdatedAt:           inspection.UpdatedAt,
	}
}

func mapSliceToDomain(inspections []Inspection) []domain.Inspection {
	domainInspections := make([]domain.Inspection, 0, len(inspections))
	for _, inspection := range inspections {
		domainInspections = append(domainInspections, mapToDomain(inspection))
	}

	return domainInspections
}

func mapConsumerToDomain(consumer Consumer) domain.Consumer {
	return domain.Consumer{
		Surname:         consumer.Surname,
		Name:            consumer.Name,
		Patronymic:      consumer.Patronymic,
		LegalEntityName: consumer.LegalEntityName,
	}
}

func mapImageToDomain(img Image) image.Image {
	return image.Image{
		Name: img.Name,
		URL:  img.URL,
	}
}

func mapImageSliceToDomain(images []Image) []image.Image {
	result := make([]image.Image, 0, len(images))
	for _, img := range images {
		result = append(result, mapImageToDomain(img))
	}

	return result
}

func mapDeviceToDomain(device Device) domain.Device {
	return domain.Device{
		Type:                device.Type,
		Number:              device.Number,
		Voltage:             device.Voltage,
		Amperage:            device.Amperage,
		AccuracyClass:       device.AccuracyClass,
		TariffsCount:        device.TariffsCount,
		DeploymentPlace:     device.DeploymentPlace,
		ValencyBeforeDot:    device.ValencyBeforeDot,
		ValencyAfterDot:     device.ValencyAfterDot,
		ManufactureYear:     device.ManufactureYear,
		VerificationQuarter: device.VerificationQuarter,
		VerificationYear:    device.VerificationYear,
		Value:               decimal.RequireFromString(device.Value),
	}
}

func mapContractToDomain(contract Contract) domain.Contract {
	return domain.Contract{
		Number: contract.Number,
		Date:   contract.Date,
	}
}
