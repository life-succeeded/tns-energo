package image

import "mime/multipart"

type UploadRequest struct {
	Address      string
	DeviceNumber string
	File         *multipart.FileHeader
}
