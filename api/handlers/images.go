package handlers

import (
	"fmt"
	"net/http"
	"tns-energo/lib/http/router"
	"tns-energo/service/image"
)

func UploadImageHandler(imageService *image.Service) router.Handler {
	return func(c router.Context) error {
		form, err := c.FormData()
		if err != nil {
			return fmt.Errorf("could not parse form data: %w", err)
		}

		if len(form.Value) != 2 ||
			len(form.Value["address"]) != 1 ||
			len(form.Value["device_number"]) != 1 ||
			len(form.File["file"]) != 1 {
			return fmt.Errorf("invalid form data")
		}

		request := image.UploadRequest{
			Address:      form.Value["address"][0],
			DeviceNumber: form.Value["device_number"][0],
			File:         form.File["file"][0],
		}

		response, err := imageService.Upload(c.Ctx(), c.Log(), request)
		if err != nil {
			return fmt.Errorf("failed to upload: %w", err)
		}

		return c.WriteJson(http.StatusOK, response)
	}
}
