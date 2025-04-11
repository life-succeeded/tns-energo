package registry

import (
	"time"
	"tns-energo/service/device"
)

type Item struct {
	Id            string        `json:"id"`
	AccountNumber string        `json:"account_number"`
	Address       string        `json:"address"`
	OldDevice     device.Device `json:"old_device"`
	NewDevice     device.Device `json:"new_device"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
}
