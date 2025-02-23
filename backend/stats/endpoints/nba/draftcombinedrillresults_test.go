package nba

import (
	"fmt"
	"testing"
)

func TestDraftCombineDrillResults(t *testing.T) {
	tests := []struct {
		leagueID   string
		seasonYear string
		expectErr  bool
	}{
		{"00", "2019", false}, // Valid request
		{"10", "2021", false}, // Valid WNBA request
		{"99", "2020", true},  // Invalid LeagueID
		{"00", "20", true},    // Invalid SeasonYear format
		{"00", "abcd", true},  // Invalid SeasonYear (non-numeric)
	}

	for _, test := range tests {
		resp, err := DraftCombineDrillResults(test.leagueID, test.seasonYear)

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
			} else {
				fmt.Printf("API call succeeded with HTTP 200 for input: %+v\n", test)
			}
		}
	}
}
