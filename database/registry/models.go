package registry

import (
	"time"
	"tns-energo/database/device"
)

type Item struct {
	Id            string        `bson:"_id,omitempty" json:"_id,omitempty"`
	AccountNumber string        `bson:"account_number" json:"account_number"`
	Address       string        `bson:"address" json:"address"`
	OldDevice     device.Device `bson:"old_device" json:"old_device"`
	NewDevice     device.Device `bson:"new_device" json:"new_device"`
	CreatedAt     time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time     `bson:"updated_at" json:"updated_at"`
}
