package bitkub

import "fmt"

// StatusResponse represents the response from /api/status endpoint
type StatusResponse struct {
	Name    string `json:"name"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

// GetStatus gets endpoint status
// GET /api/status
func GetStatus() ([]StatusResponse, error) {
	var result []StatusResponse

	if err := FetchNonSecure("GET", "/api/status", nil, &result); err != nil {
		return nil, fmt.Errorf("get status: %w", err)
	}

	return result, nil
}

// GetServerTime gets server timestamp in milliseconds
// GET /api/v3/servertime
func GetServerTime() (int64, error) {
	var result int64

	if err := FetchNonSecure("GET", "/api/v3/servertime", nil, &result); err != nil {
		return 0, fmt.Errorf("get server time: %w", err)
	}

	return result, nil
}
