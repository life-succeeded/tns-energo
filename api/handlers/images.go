package handlers

import (
	"net/http"
	"tns-energo/lib/http/router"
	"tns-energo/service/image"
)

func UploadImageHandler(imageService *image.Service) router.Handler {
	return func(c router.Context) error {
		log := c.Log()

		var request image.UploadRequest
		if err := c.ReadJson(&request); err != nil {
			log.Errorf("failed to parse request body: %w", err)
			return err
		}

		response, err := imageService.Upload(c.Ctx(), log, request)
		if err != nil {
			log.Errorf("failed to upload: %v", err)
			return err
		}

		return c.WriteJson(http.StatusOK, response)
	}
}
