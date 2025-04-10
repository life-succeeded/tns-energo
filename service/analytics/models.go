package analytics

import "time"

type Report struct {
	Id        string     `json:"id"`
	Type      ReportType `json:"type"`
	Name      string     `json:"name"`
	URL       string     `json:"url"`
	ForDate   time.Time  `json:"for_date"`
	CreatedAt time.Time  `json:"created_at"`
}
