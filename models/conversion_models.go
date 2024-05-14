package models

// ConversionRequest represents the request body for the conversion endpoint.
type ConversionRequest struct {
	Kilograms float64 `json:"kilograms"`
}

// ConversionResponse represents the response body for the conversion endpoint.
type ConversionResponse struct {
	Pounds float64 `json:"pounds"`
}
