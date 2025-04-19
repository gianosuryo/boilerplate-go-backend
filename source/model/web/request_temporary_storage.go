package web

type RequestTemporaryStorage struct {
	Data string `json:"data" validate:"required"`
}
