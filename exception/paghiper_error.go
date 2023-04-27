package exception

// PaghiperError represents an error response from the Paghiper API.
type PaghiperError struct {
	// ResponseMessage contains a readable error message returned by the API.
	ResponseMessage string `json:"response_message"`
	// HTTPCode contains the HTTP status code of the error response.
	HTTPCode int `json:"http_code"`
}
