package nba

import (
	"fmt"
	"net/http"
	"testing"
)

func TestAllTimeLeadersGrids_ActualCall(t *testing.T) {
	tests := []struct {
		leagueID   string
		perMode    string
		seasonType string
		topX       int
		expectErr  bool
	}{
		{"00", "Totals", "Regular Season", 10, false},     // Valid request
		{"99", "Totals", "Regular Season", 10, true},      // Invalid LeagueID
		{"00", "InvalidMode", "Regular Season", 10, true}, // Invalid perMode
		{"00", "Totals", "Unknown Season", 10, true},      // Invalid SeasonType
		{"00", "Totals", "Regular Season", -1, true},      // Invalid TopX
	}

	for _, test := range tests {
		resp, err := AllTimeLeadersGrids(test.leagueID, test.perMode, test.seasonType, test.topX)

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
