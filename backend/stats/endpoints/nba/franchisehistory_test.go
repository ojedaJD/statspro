package nba

import (
	"fmt"
	"net/http"
	"testing"
)

func TestFranchiseHistory(t *testing.T) {
	validLeagueID := "00"
	invalidLeagueID := "99"

	tests := []struct {
		leagueID  string
		expectErr bool
	}{
		{validLeagueID, false},  // Valid request
		{invalidLeagueID, true}, // Invalid LeagueID
		{"", true},              // Missing LeagueID
	}

	for _, test := range tests {
		resp, err := FranchiseHistory(test.leagueID)

		if test.expectErr {
			if err == nil {
				t.Errorf("Expected error but got nil for input: %+v", test)
			} else {
				fmt.Printf("Expected error received for input: %+v\n", test)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error: %v for input: %+v", err, test)
			} else if resp == nil {
				t.Errorf("Received nil response from API for input: %+v", test)
			} else if resp.StatusCode != http.StatusOK {
				t.Errorf("Unexpected response code: got %d, expected %d for input: %+v", resp.StatusCode, http.StatusOK, test)
			} else {
				fmt.Printf("API call succeeded with HTTP 200 for input: %+v\n", test)
			}
		}
	}
}
