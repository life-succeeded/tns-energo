package handlers

import (
	"net/http"
	"time"
	"tns-energo/lib/http/router"
	"tns-energo/service/analytics"
)

type generateDailyReportVars struct {
	Date string `path:"date"`
}

func GenerateDailyReportHandler(analyticsService *analytics.Service) router.Handler {
	return func(c router.Context) error {
		log := c.Log()

		var vars generateDailyReportVars
		if err := c.Vars(&vars); err != nil {
			log.Errorf("failed to read vars: %v", err)
			return err
		}

		date, err := time.Parse("2006-01-02", vars.Date)
		if err != nil {
			log.Errorf("failed to parse date: %v", err)
			return err
		}

		response, err := analyticsService.GenerateDailyReport(c.Ctx(), log, date)
		if err != nil {
			log.Errorf("failed to generate daily report: %v", err)
			return err
		}

		return c.WriteJson(http.StatusOK, response)
	}
}

func GetAllReportsHandler(analyticsService *analytics.Service) router.Handler {
	return func(c router.Context) error {
		log := c.Log()

		response, err := analyticsService.GetAllReports(c.Ctx(), log)
		if err != nil {
			log.Errorf("failed to get all reports: %v", err)
		}

		return c.WriteJson(http.StatusOK, response)
	}
}
