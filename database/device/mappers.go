package device

import (
	domain "tns-energo/service/device"

	"github.com/shopspring/decimal"
)

func MapToDb(d domain.Device) Device {
	return Device{
		Type:                d.Type,
		Number:              d.Number,
		Voltage:             d.Voltage,
		Amperage:            d.Amperage,
		AccuracyClass:       d.AccuracyClass,
		TariffsCount:        d.TariffsCount,
		DeploymentPlace:     d.DeploymentPlace,
		ValencyBeforeDot:    d.ValencyBeforeDot,
		ValencyAfterDot:     d.ValencyAfterDot,
		ManufactureYear:     d.ManufactureYear,
		VerificationQuarter: d.VerificationQuarter,
		VerificationYear:    d.VerificationYear,
		Value:               d.Value.String(),
	}
}

func MapToDomain(d Device) domain.Device {
	return domain.Device{
		Type:                d.Type,
		Number:              d.Number,
		Voltage:             d.Voltage,
		Amperage:            d.Amperage,
		AccuracyClass:       d.AccuracyClass,
		TariffsCount:        d.TariffsCount,
		DeploymentPlace:     d.DeploymentPlace,
		ValencyBeforeDot:    d.ValencyBeforeDot,
		ValencyAfterDot:     d.ValencyAfterDot,
		ManufactureYear:     d.ManufactureYear,
		VerificationQuarter: d.VerificationQuarter,
		VerificationYear:    d.VerificationYear,
		Value:               decimal.RequireFromString(d.Value),
	}
}
