package device

import (
	"tns-energo/database/seal"
	domain "tns-energo/service/device"
)

func MapToDb(d domain.Device) Device {
	return Device{
		Type:            d.Type,
		Number:          d.Number,
		DeploymentPlace: int(d.DeploymentPlace),
		OtherPlace:      d.OtherPlace,
		Seals:           seal.MapSliceToDb(d.Seals),
		Value:           d.Value,
		Consumption:     d.Consumption,
	}
}

func MapToDomain(d Device) domain.Device {
	return domain.Device{
		Type:            d.Type,
		Number:          d.Number,
		DeploymentPlace: domain.DeploymentPlace(d.DeploymentPlace),
		OtherPlace:      d.OtherPlace,
		Seals:           seal.MapSliceToDomain(d.Seals),
		Value:           d.Value,
		Consumption:     d.Consumption,
	}
}
