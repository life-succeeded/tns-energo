package image

import "mime/multipart"

type UploadRequest struct {
	Address      string                `json:"address"`
	DeviceNumber string                `json:"device_number"`
	Payload      *multipart.FileHeader `json:"payload"`
}
