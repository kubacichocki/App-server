package models

// ConversionRequest represents the request body for the conversion endpoint.
type ConversionRequest struct {
	Type  string  `json:"type"`  // Type of conversion (e.g., "kgToPounds", "poundsToKg")
	Value float64 `json:"value"` // Value to be converted
}

// ConversionResponse represents the response body for the conversion endpoint.
type ConversionResponse struct {
	Value string `json:"value"` // Converted value
}
