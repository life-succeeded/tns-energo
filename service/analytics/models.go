package analytics

import (
	"time"
	"tns-energo/service/file"
)

type Report struct {
	Id        string     `json:"id"`
	Type      ReportType `json:"type"`
	File      file.File  `json:"file"`
	ForDate   time.Time  `json:"for_date"`
	CreatedAt time.Time  `json:"created_at"`
}
