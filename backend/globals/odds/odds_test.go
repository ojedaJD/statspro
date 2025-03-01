package odds

import (
	"fmt"
	"testing"
)

func TestGetSports(t *testing.T) {
	// Replace with your actual API key
	apiKey := "6712797a115aa05245a64e00077eb993"
	client := NewOddsApiClient(apiKey)

	// Call the /v4/sports endpoint without appending anything
	// No additional sport appended
	// No extra query params

	response, err := client.GetOddsRequest(client.GetSportsURL(), nil, nil)
	if err != nil {
		t.Fatalf("Failed to fetch sports: %v", err)
	}

	// Print the API response
	fmt.Printf("Response: %+v\n", response)

	// Check for a successful response
	if response.StatusCode != 200 {
		t.Errorf("Expected status 200, got %d", response.StatusCode)
	}

	// Ensure the response contains data
	if response.Data == nil {
		t.Errorf("Expected non-nil data, got nil")
	}
}
