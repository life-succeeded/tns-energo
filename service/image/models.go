package image

type UploadRequest struct {
	Address      string `json:"address"`
	DeviceNumber string `json:"device_number"`
	Payload      string `json:"payload"`
}
