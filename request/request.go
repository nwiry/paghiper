package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// makeRequest sends a HTTP POST request with a JSON payload to the specified endpoint and returns
// the response as a map[string]interface{}, along with the HTTP status code and any errors encountered.
func Make(endpoint, action string, payload interface{}) (map[string]interface{}, int, error) {
	// Convert the payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, 0, err
	}

	// Create a new HTTP request with the JSON payload
	request, err := http.NewRequest(http.MethodPost, endpoint+action, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, 0, err
	}

	// Set headers for the request
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	// Send the request and get the response
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, 0, err
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, response.StatusCode, err
	}

	// Unmarshal the response body to a map[string]interface{}
	responseData := make(map[string]interface{})
	if err = json.Unmarshal(body, &responseData); err != nil {
		return nil, 0, err
	}

	var ref string
	// If the endpoint is the PIX endpoint, set a reference string to "pix_"
	if strings.Contains(endpoint, "pix") {
		ref = "pix_"
	}

	// Check if the "pix_create_request" (or "create_request" if not PIX endpoint) key is present in the response
	if responseData[(ref+"create_request")] == nil {
		return nil, response.StatusCode, fmt.Errorf("%screate_request not found in response", ref)
	}

	// If the "pix_create_request" (or "create_request" if not PIX endpoint) key is present, return its value
	if val, ok := responseData[(ref + "create_request")].(map[string]interface{}); ok {
		return val, response.StatusCode, nil
	}

	// If the "pix_create_request" (or "create_request" if not PIX endpoint) key is not a map[string]interface{}, return an error
	return nil, response.StatusCode, fmt.Errorf("unable to parse %screate_request from response", ref)
}

// validateRequestResponse checks if the "result" key is "reject" or if the HTTP status code
// is greater than http.StatusCreated, and returns an error with the "response_message" key
// from the response.
func ValidateRequestResponse(r map[string]interface{}, c int) error {
	if r["result"] == "reject" || c > http.StatusCreated {
		return fmt.Errorf("%v", r["response_message"])
	}

	return fmt.Errorf("map: %v", r)
}

// setResponse converts a map[string]interface{} to JSON and decodes it into the specified struct.
func SetResponse(r map[string]interface{}, s interface{}) error {
	// Convert the map to JSON
	j, err := json.Marshal(r)
	if err != nil {
		return err
	}

	// Decode the JSON into the output struct
	if err := json.Unmarshal(j, s); err != nil {
		return err
	}

	return nil
}
