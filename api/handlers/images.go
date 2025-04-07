package handlers

import (
	"net/http"
	"tns-energo/lib/http/router"
	"tns-energo/service/image"
)

func UploadImageHandler(imageService *image.Service) router.Handler {
	return func(c router.Context) error {
		log := c.Log()

		payload, err := c.ReadBytes()
		if err != nil {
			log.Errorf("failed to read: %v", err)
			return err
		}

		response, err := imageService.Upload(c.Ctx(), log, payload)
		if err != nil {
			log.Errorf("failed to upload: %v", err)
			return err
		}

		return c.WriteJson(http.StatusOK, response)
	}
}
