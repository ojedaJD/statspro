package nba

import (
	"fmt"
	"net/http"
	"testing"
)

func TestFranchisePlayers(t *testing.T) {
	validLeagueID := "00"
	validPerMode := "Totals"
	validSeasonType := "Regular Season"
	validTeamID := "1610612739"

	invalidLeagueID := "99"
	invalidPerMode := "InvalidMode"
	invalidSeasonType := "Spring League"
	invalidTeamID := ""

	tests := []struct {
		leagueID   string
		perMode    string
		seasonType string
		teamID     string
		expectErr  bool
	}{
		{validLeagueID, validPerMode, validSeasonType, validTeamID, false},  // Valid request
		{invalidLeagueID, validPerMode, validSeasonType, validTeamID, true}, // Invalid LeagueID
		{validLeagueID, invalidPerMode, validSeasonType, validTeamID, true}, // Invalid PerMode
		{validLeagueID, validPerMode, invalidSeasonType, validTeamID, true}, // Invalid SeasonType
		{validLeagueID, validPerMode, validSeasonType, invalidTeamID, true}, // Missing TeamID
	}

	for _, test := range tests {
		resp, err := FranchisePlayers(test.leagueID, test.perMode, test.seasonType, test.teamID)

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
