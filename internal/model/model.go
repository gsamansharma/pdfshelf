package model

type PDFEntry struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	FilePath string    `json:"file_path"`
}

type Library struct {
	PDFs   []PDFEntry `json:"pdfs"`
	NextID int        `json:"next_id"`
}