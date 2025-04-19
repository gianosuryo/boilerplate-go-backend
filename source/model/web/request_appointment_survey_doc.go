package web

type RequestAppointmentSurveyDoc struct {
	Description   string `json:"description" validate:"required" example:"Description"`
	FileUrl       string `json:"file_url" validate:"required" example:"https://example.com/file.pdf"`
	FileExtension string `json:"file_extension" validate:"required" example:"pdf"`
	FileSizeBytes uint64 `json:"file_size_bytes" validate:"required" example:"1000"`
}
