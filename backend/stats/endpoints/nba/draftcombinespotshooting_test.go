package nba

import (
	"fmt"
	"net/http"
	"testing"
)

func TestDraftCombineSpotShooting(t *testing.T) {
	validLeagueID := "00"
	validSeasonYear := "2019"
	invalidLeagueID := "99"
	invalidSeasonYear := "19"

	tests := []struct {
		leagueID   string
		seasonYear string
		expectErr  bool
	}{
		{validLeagueID, validSeasonYear, false},  // Valid request
		{invalidLeagueID, validSeasonYear, true}, // Invalid LeagueID
		{validLeagueID, invalidSeasonYear, true}, // Invalid SeasonYear
		{"", validSeasonYear, true},              // Missing LeagueID
		{validLeagueID, "", true},                // Missing SeasonYear
	}

	for _, test := range tests {
		resp, err := DraftCombineSpotShooting(test.leagueID, test.seasonYear)
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
