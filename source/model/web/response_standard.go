package web

type ResponseStandard struct {
	StatusCode int             `json:"statusCode"`
	Message    string          `json:"message"`
	Fulfilled  uint8           `json:"fulfilled"`
	Data       any             `json:"data,omitempty"`
	Errors     []ErrorResponse `json:"errors,omitempty"`
	// Meta       any `json:"meta"`
	// JsonApi    any `json:"jsonapi"`
	// Links      any `json:"links"`
	// Included   any `json:"included"`
	// Self       any `json:"self"`
	// Related    any `json:"related"`
	Pagination any `json:"pagination,omitempty"`
}

type ErrorResponse struct {
	Title    string   `json:"title"`
	Messages []string `json:"messages"`
}
