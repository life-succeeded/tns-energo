package handlers

import (
	"fmt"
	"net/http"
	"tns-energo/lib/http/router"
	"tns-energo/service/image"
)

func UploadImageHandler(imageService *image.Service) router.Handler {
	return func(c router.Context) error {
		var request image.UploadRequest
		if err := c.ReadJson(&request); err != nil {
			return fmt.Errorf("failed to parse request body: %w", err)
		}

		response, err := imageService.Upload(c.Ctx(), c.Log(), request)
		if err != nil {
			return fmt.Errorf("failed to upload: %w", err)
		}

		return c.WriteJson(http.StatusOK, response)
	}
}
