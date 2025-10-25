package model

import(
	"time"
)

type PDFEntry struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	FilePath string    `json:"file_path"`
	AddedOn        time.Time       `json:"added_on"`
	TotalTimeSpent time.Duration   `json:"total_time_spent"`
}

type Library struct {
	PDFs   []PDFEntry `json:"pdfs"`
	NextID int        `json:"next_id"`
}