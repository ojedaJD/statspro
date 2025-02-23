package nba

import (
	"fmt"
	"net/http"
	"testing"
)

func TestAssistLeaders_ActualCall(t *testing.T) {
	tests := []struct {
		leagueID   string
		season     string
		seasonType string
		perMode    string
		topX       int
		expectErr  bool
	}{
		{"00", "2023-24", "Regular Season", "PerGame", 10, false},      // Valid request
		{"99", "2023-24", "Regular Season", "PerGame", 10, true},       // Invalid LeagueID
		{"00", "2023-24", "Regular Season", "InvalidMode", 10, true},   // Invalid PerMode
		{"00", "2023-24", "Unknown Season", "PerGame", 10, true},       // Invalid SeasonType
		{"00", "InvalidSeason", "Regular Season", "PerGame", 10, true}, // Invalid Season format
		{"00", "2023-24", "Regular Season", "PerGame", -1, true},       // Invalid TopX
	}

	for _, test := range tests {
		resp, err := AssistLeaders(test.leagueID, test.season, test.seasonType, test.perMode, test.topX)

		if test.expectErr {
			// If we expect an error, ensure an error is returned
			if err == nil {
				t.Errorf("Expected error but got nil for input: %v", test)
			} else {
				fmt.Printf("Expected error received for input: %v\n", test)
			}
		} else {
			// If no error is expected, check the response
			if err != nil {
				t.Errorf("Unexpected error: %v for input: %v", err, test)
			} else if resp == nil {
				t.Errorf("Received nil response from API for input: %v", test)
			} else if resp.StatusCode != http.StatusOK {
				t.Errorf("Unexpected response code: got %d, expected %d for input: %v", resp.StatusCode, http.StatusOK, test)
			} else {
				fmt.Printf("API call succeeded with HTTP 200 for input: %v\n", test)
			}
		}
	}
}
