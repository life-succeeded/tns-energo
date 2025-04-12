package handlers

import (
	"fmt"
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
		var vars generateDailyReportVars
		if err := c.Vars(&vars); err != nil {
			return fmt.Errorf("failed to read vars: %w", err)
		}

		date, err := time.Parse("2006-01-02", vars.Date)
		if err != nil {
			return fmt.Errorf("failed to parse date: %w", err)
		}

		response, err := analyticsService.GenerateDailyReport(c.Ctx(), c.Log(), date)
		if err != nil {
			return fmt.Errorf("failed to generate daily report: %w", err)
		}

		return c.WriteJson(http.StatusOK, response)
	}
}

func GetAllReportsHandler(analyticsService *analytics.Service) router.Handler {
	return func(c router.Context) error {
		response, err := analyticsService.GetAllReports(c.Ctx(), c.Log())
		if err != nil {
			return fmt.Errorf("failed to get all reports: %w", err)
		}

		return c.WriteJson(http.StatusOK, response)
	}
}
