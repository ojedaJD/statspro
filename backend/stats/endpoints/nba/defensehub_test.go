package nba

import (
	"fmt"
	"net/http"
	"testing"
)

func TestDefenseHub_ActualCall(t *testing.T) {
	tests := []struct {
		options   DefenseHubOptions
		expectErr bool
	}{
		{DefenseHubOptions{"Season", "00", "Team", "All Players", "2019-20", "Regular Season"}, false},   // Valid request
		{DefenseHubOptions{"Finals", "00", "Player", "Rookies", "2019-20", "Playoffs"}, false},           // Valid request
		{DefenseHubOptions{"Last 10", "99", "Team", "All Players", "2019-20", "Regular Season"}, true},   // Invalid LeagueID
		{DefenseHubOptions{"Season", "00", "Unknown", "All Players", "2019-20", "Regular Season"}, true}, // Invalid PlayerOrTeam
		{DefenseHubOptions{"Yesterday", "00", "Team", "All Players", "20-19", "Regular Season"}, true},   // Invalid Season format
	}

	for _, test := range tests {
		resp, err := DefenseHub(test.options)

		if test.expectErr {
			if err == nil {
				t.Errorf("Expected error but got nil for input: %+v", test.options)
			} else {
				fmt.Printf("Expected error received for input: %+v\n", test.options)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error: %v for input: %+v", err, test.options)
			} else if resp == nil {
				t.Errorf("Received nil response from API for input: %+v", test.options)
			} else if resp.StatusCode != http.StatusOK {
				t.Errorf("Unexpected response code: got %d, expected %d for input: %+v", resp.StatusCode, http.StatusOK, test.options)
			} else {
				fmt.Printf("API call succeeded with HTTP 200 for input: %+v\n", test.options)
			}
		}
	}
}
